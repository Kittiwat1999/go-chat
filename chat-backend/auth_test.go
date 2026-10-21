package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"chat-backend/database"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func TestLoadJWTSecretRequiresAtLeast32Bytes(t *testing.T) {
	if _, err := loadJWTSecret(""); err == nil {
		t.Fatal("loadJWTSecret() accepted an empty secret")
	}
	if _, err := loadJWTSecret("too-short"); err == nil {
		t.Fatal("loadJWTSecret() accepted a short secret")
	}

	secret := "0123456789abcdef0123456789abcdef"
	loaded, err := loadJWTSecret(secret)
	if err != nil {
		t.Fatalf("loadJWTSecret() error = %v", err)
	}
	if string(loaded) != secret {
		t.Fatal("loadJWTSecret() returned a different secret")
	}
}

func TestCreateAuthTokenContainsUserClaimsAndExpires(t *testing.T) {
	now := time.Date(2026, time.October, 21, 0, 0, 0, 0, time.UTC)
	secret := []byte("0123456789abcdef0123456789abcdef")
	user := database.User{ID: "user_123", Username: "alice"}

	signedToken, err := createAuthToken(user, secret, now)
	if err != nil {
		t.Fatalf("createAuthToken() error = %v", err)
	}

	claims := &authClaims{}
	token, err := jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (any, error) {
		return secret, nil
	}, jwt.WithTimeFunc(func() time.Time { return now }))
	if err != nil {
		t.Fatalf("jwt.ParseWithClaims() error = %v", err)
	}
	if !token.Valid {
		t.Fatal("created token is not valid")
	}
	if claims.UserID != user.ID || claims.Username != user.Username || claims.Subject != user.ID {
		t.Fatalf("unexpected token user claims: %+v", claims)
	}
	if !claims.ExpiresAt.Time.Equal(now.Add(authTokenLifetime)) {
		t.Fatalf("token expiry = %v, want %v", claims.ExpiresAt.Time, now.Add(authTokenLifetime))
	}

	_, err = jwt.ParseWithClaims(signedToken, &authClaims{}, func(token *jwt.Token) (any, error) {
		return []byte("abcdef0123456789abcdef0123456789"), nil
	}, jwt.WithTimeFunc(func() time.Time { return now }))
	if err == nil {
		t.Fatal("token verified with the wrong secret")
	}
}

func TestAuthenticateTokenRejectsExpiredToken(t *testing.T) {
	now := time.Now()
	secret := []byte("0123456789abcdef0123456789abcdef")
	signedToken, err := createAuthToken(database.User{ID: "user_123", Username: "alice"}, secret, now.Add(-authTokenLifetime-time.Minute))
	if err != nil {
		t.Fatalf("createAuthToken() error = %v", err)
	}

	_, err = authenticateToken(signedToken, secret)
	if err == nil {
		t.Fatal("authenticateToken() accepted an expired token")
	}
}

func TestRequireAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	secret := []byte("0123456789abcdef0123456789abcdef")
	now := time.Now()
	signedToken, err := createAuthToken(database.User{ID: "user_123", Username: "alice"}, secret, now)
	if err != nil {
		t.Fatalf("createAuthToken() error = %v", err)
	}

	router := gin.New()
	router.GET("/protected", requireAuth(secret), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"userId": c.GetString("userId")})
	})

	tests := []struct {
		name       string
		authHeader string
		wantStatus int
	}{
		{name: "missing token", wantStatus: http.StatusUnauthorized},
		{name: "invalid token", authHeader: "Bearer invalid", wantStatus: http.StatusUnauthorized},
		{name: "valid token", authHeader: "Bearer " + signedToken, wantStatus: http.StatusOK},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/protected", nil)
			if test.authHeader != "" {
				request.Header.Set("Authorization", test.authHeader)
			}
			response := httptest.NewRecorder()
			router.ServeHTTP(response, request)
			if response.Code != test.wantStatus {
				t.Fatalf("status = %d, want %d", response.Code, test.wantStatus)
			}
		})
	}
}
