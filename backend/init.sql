-- schema.sql

CREATE TABLE IF NOT EXISTS players (
  playerID VARCHAR PRIMARY KEY,
  birthYear INTEGER,
  nameFirst VARCHAR,
  nameLast VARCHAR
);
  

INSERT INTO players VALUES ('tt', 2001, 'zz', 'xx');
