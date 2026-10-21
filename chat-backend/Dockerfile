FROM golang:1.27.1-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /chat-backend .

FROM alpine:3.22

COPY --from=build /chat-backend /usr/local/bin/chat-backend

EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/chat-backend"]
