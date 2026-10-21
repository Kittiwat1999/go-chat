package main

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"chat-backend/database"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Hub struct {
	mu      sync.RWMutex
	clients map[*websocket.Conn]bool
}

type RegisterUserRequest struct {
	Username    string `json:"username" binding:"required,min=3,max=30"`
	DisplayName string `json:"displayName" binding:"required,max=50"`
	Password    string `json:"password" binding:"required,min=6"`
}

type LoginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func newHub() *Hub {
	return &Hub{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (h *Hub) addClient(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[conn] = true
}

func (h *Hub) removeClient(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.clients, conn)
	conn.Close()
}

func (h *Hub) broadcast(messageType int, message []byte) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	for conn := range h.clients {
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Printf("Broadcast error: %v", err)
		}
	}
}

func main() {
	jwtSecret, err := loadJWTSecret(os.Getenv("JWT_SECRET"))
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Connect(os.Getenv("MONGODB_URI")); err != nil {
		log.Fatalf("Connect to MongoDB: %v", err)
	}
	defer database.Close()

	hub := newHub()
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.POST("/api/users/register", func(c *gin.Context) {
		var req RegisterUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := hashPassword(req.Password)
		if err != nil {
			log.Printf("Hash password: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
			return
		}

		user := database.User{
			Username:    req.Username,
			DisplayName: req.DisplayName,
			Password:    string(hashedPassword),
			Status:      "offline",
		}
		if err := database.InsertUser(&user); err != nil {
			if mongo.IsDuplicateKeyError(err) {
				c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
				return
			}
			log.Printf("Insert user: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
			return
		}

		c.JSON(http.StatusCreated, user)
	})

	router.POST("/api/users/login", func(c *gin.Context) {
		var req LoginUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := database.FindUserByUsername(req.Username)
		if err != nil {
			if err != mongo.ErrNoDocuments {
				log.Printf("Find user for login: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not log in"})
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		if err := comparePassword(user.Password, req.Password); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		token, err := createAuthToken(user, jwtSecret, time.Now())
		if err != nil {
			log.Printf("Create login token: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not log in"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"user":  user,
		})
	})

	router.GET("/api/users/:username", requireAuth(jwtSecret), func(c *gin.Context) {
		username := c.Param("username")
		var user database.User
		err := database.GetCollection("go-chat", "users").FindOne(c, map[string]interface{}{"username": username}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	router.GET("/ws", func(c *gin.Context) {
		claims, err := authenticateToken(c.Query("token"), jwtSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("Upgrade error: %v", err)
			return
		}
		hub.addClient(conn)
		clientAddr := conn.RemoteAddr().String()
		log.Printf("User %s connected over WebSocket: %s", claims.Username, clientAddr)
		defer func() {
			hub.removeClient(conn)
			log.Printf("User %s disconnected from WebSocket: %s", claims.Username, clientAddr)
		}()

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Read error: %v", err)
				break
			}
			hub.broadcast(messageType, message)
		}
	})

	router.Run(":8080")
}
