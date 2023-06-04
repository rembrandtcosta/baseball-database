-- query.sql

-- name: GetPlayer :one
SELECT * FROM players
WHERE playerID = $1 LIMIT 1;

-- name: ListPlayers :many
SELECT * FROM players
ORDER BY playerID;

-- name: CreatePlayer :one
INSERT INTO players (
  playerID,
  birthYear,
  birthMonth,
  birthDay,
  birthCountry,
  birthState,
  birthCity,
  deathYear,
  deathMonth,
  deathDay,
  deathCountry,
  deathState,
  deathCity,
  nameFirst, 
  nameLast,
  nameGiven,
  weight,
  height,
  bats,
  throws,
  debut,
  finalGame,
  retroID,
  bbrefID
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24)
RETURNING *;

-- name: DeletePlayer :exec
DELETE FROM players
WHERE playerID = $1;

-- name: GetFranchise :one
SELECT * FROM franchises
WHERE franchID = $1 LIMIT 1;

-- name: ListFranchises :many
SELECT * FROM franchises
ORDER BY franchID;

-- name: CreateFranchise :one
INSERT INTO franchises (
  franchID,
  franchName,
  active,
  NAassoc
) VALUES ($1, $2, $3, $4) 
RETURNING *;

-- name: DeleteFranchise :exec
DELETE FROM franchises
WHERE franchID = $1;
