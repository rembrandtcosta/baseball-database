-- query.sql

-- name: GetPlayer :one
SELECT * FROM players
WHERE playerID = $1 LIMIT 1;

-- name: ListPlayers :many
SELECT * FROM players
ORDER BY playerID;

-- name: CreatePlayer :one
INSERT INTO players (
  playerID, birthYear, nameFirst, nameLast
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM players
WHERE playerID = $1;
