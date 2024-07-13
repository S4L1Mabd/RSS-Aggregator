-- name: CreateUser :one

INSERT INTO users(id,created_time,update_time,name) /* this file used from sqlc to run this queries and give a .go equivalent*/
VALUES ($1,$2,$3,$4)
RETURNING *;