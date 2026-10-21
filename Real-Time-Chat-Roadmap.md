# Real-Time Chat Server Roadmap (Go + Vue + MongoDB)

## Goal
สร้าง Real-Time Chat Server เพื่อเรียนรู้ Go Backend, WebSocket, Docker และ MongoDB โดยพัฒนาเป็น Phase

---

## Tech Stack

### Frontend
- Vue.js
- Vite
- Axios
- Native WebSocket

### Backend
- Go
- Gin
- Gorilla WebSocket

### Database
- MongoDB
- MongoDB Compass (GUI)

### Infrastructure
- Docker
- Docker Compose

## Local Docker Setup

Start the Go API, Vuetify web app, MongoDB, and Mongo Express from the project root:

```bash
docker compose up --build
```

- Web app: http://localhost:3000
- Go API: http://localhost:8080/ping
- Mongo Express: http://localhost:8081 (default login: `admin` / `pass`)
- MongoDB: `mongodb://chatadmin:chatpass@localhost:27017/?authSource=admin`

These credentials are for local development only. Override them with a root `.env`
file before exposing the services outside your machine:

```dotenv
MONGO_INITDB_ROOT_USERNAME=chatadmin
MONGO_INITDB_ROOT_PASSWORD=replace-this
MONGO_EXPRESS_USERNAME=admin
MONGO_EXPRESS_PASSWORD=replace-this-too
```

MongoDB data is persisted in the `mongo_data` volume. The API receives its
connection string in `MONGODB_URI`; the current API does not yet connect to MongoDB.
If port 3000 is already in use, add `WEB_PORT=3001` to the root `.env` file and
open the web app at http://localhost:3001. Stop the services with
`docker compose down`. To also delete persisted database data, run
`docker compose down -v`.

### Future Add-ons
- Redis
- JWT Authentication
- Nginx
- Kubernetes

---

## Architecture V1

```text
Vue Client
    |
    | HTTP / WebSocket
    v
Go Backend
    |
    v
MongoDB
```

Message Flow

```text
Vue
  -> WebSocket
  -> Go Server
  -> MongoDB
  -> Broadcast
  -> All Connected Clients
```

---

# Phase 1: MVP Chat

## Features
- Create Account
- Join chat
- Send message
- Receive message realtime
- Persist messages in MongoDB

## REST API

```http
GET  /messages
POST /messages
```

## WebSocket

```http
/ws
```

## Learning Goals
- Docker Compose
- Go API
- MongoDB Connection
- WebSocket
- Goroutine

---

# Phase 2: Chat Rooms

## Features
- Username
- Multiple rooms
- Join room
- Leave room

Example Rooms

```text
/general
/golang
/random
```

## Learning Goals
- Room Management
- In-memory State
- Client Tracking

---

# Phase 3: Redis Pub/Sub

## Features
- Multi-instance chat
- Distributed message broadcasting

Architecture

```text
Vue
  |
  v
Go Server #1
             Redis Pub/Sub
      /
Go Server #2
```

## Learning Goals
- Scaling
- Event Driven Architecture
- Pub/Sub Pattern

---

# Phase 4: Authentication

## Features
- Register
- Login
- JWT Access Token
- Protected Routes

## Learning Goals
- Authentication
- Authorization
- Token Validation

---

# Phase 5: Deployment

## Stack
- Nginx
- Go
- MongoDB
- Redis
- Docker Compose

## Goals
- Deploy on VPS
- Environment Variables
- Reverse Proxy
- Production Setup

---

# Suggested Project Structure

```text
chat-app/
|
|-- backend/
|   |-- cmd/
|   |-- handlers/
|   |-- websocket/
|   |-- repository/
|   |-- models/
|   `-- config/
|
|-- frontend/
|   `-- src/
|
|-- docker-compose.yml
`-- README.md
```

---

# Milestone Checklist

## Phase 1
- [ ] Docker Compose running
- [ ] Vue frontend connected
- [ ] Go API working
- [ ] MongoDB connected
- [ ] WebSocket working

## Phase 2
- [ ] Username support
- [ ] Room support
- [ ] Room switching

## Phase 3
- [ ] Redis installed
- [ ] Pub/Sub integrated
- [ ] Multi-instance broadcasting

## Phase 4
- [ ] Register API
- [ ] Login API
- [ ] JWT middleware

## Phase 5
- [ ] Production Docker Compose
- [ ] Nginx setup
- [ ] VPS deployment

---

## End Goal

Portfolio-ready project demonstrating:

- Go Backend
- REST API
- WebSocket
- Goroutine & Concurrency
- MongoDB
- Redis Pub/Sub
- JWT
- Docker
- Deployment
