CREATE DATABASE IF NOT EXISTS mydb;

use mydb;

CREATE TABLE Obj(
    id VARCHAR(30) PRIMARY KEY,
    val VARCHAR(50),
    updated TIMESTAMP
);

insert into Obj (id, val, updated) values ("a", "def", "2019-12-02 06:53:32")

