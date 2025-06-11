-- name: CreateUser :one
INSERT INTO users (username, password) VALUES (?, ?) RETURNING id;

-- name: GetUser :one
SELECT id, password FROM users WHERE username = ? LIMIT 1;