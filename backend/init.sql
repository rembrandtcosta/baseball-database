-- schema.sql

CREATE TABLE IF NOT EXISTS players (
  playerID VARCHAR PRIMARY KEY,
  birthYear INTEGER,
  birthMonth INTEGER,
  birthDay INTEGER,
  birthCountry VARCHAR,
  birthState VARCHAR,
  birthCity VARCHAR,
  deathYear INTEGER,
  deathMonth INTEGER,
  deathDay INTEGER,
  deathCountry VARCHAR,
  deathState VARCHAR,
  deathCity VARCHAR,
  nameFirst VARCHAR, 
  nameLast VARCHAR,
  nameGiven VARCHAR,
  weight INTEGER,
  height INTEGER,
  bats VARCHAR,
  throws VARCHAR,
  debut DATE,
  finalGame DATE,
  retroID VARCHAR,
  bbrefID VARCHAR
);

CREATE TABLE IF NOT EXISTS franchises (
  franchID VARCHAR PRIMARY KEY,
  franchName VARCHAR,
  active VARCHAR,
  NAassoc VARCHAR
);
  

INSERT INTO players 
VALUES ('tt', 2001, 11, 26, 'BRA', 'PB', 'Campina Grande', 2111, 13, 32, 'TRA', 'TR', 'Tra',  'zz', 'xx', 'Zz Xx', 10, 11, 'L', 'L', '2023-01-01', '2023-01-02', 'tt', 'tt');

INSERT INTO franchises
VALUES ('1', '1', '1', '1');
