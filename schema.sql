CREATE DATABASE gorello IF NOT EXISTS;
USE gorello;


CREATE TABLE projects IF NOT EXISTS(
id SERIAL PRIMARY KEY,
name varchar(500) NOT NULL,
description varchar(1000),
);


CREATE TABLE columns IF NOT EXISTS(
id SERIAL PRIMARY KEY,
name varchar(255) NOT NULL UNIQUE,
project_id BIGINT UNSIGNED,
CONSTRAINT project_id_fk FOREIGN KEY (project_id)
    REFERENCES projects(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);


CREATE TABLE tasks IF NOT EXISTS(
id SERIAL PRIMARY KEY,
name varchar(500) NOT NULL,
description text(5000),
priority int NOT NULL,
column_id BIGINT UNSIGNED,
CONSTRAINT column_id_fk FOREIGN KEY (column_id)
    REFERENCES columns(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);


CREATE TABLE comments IF NOT EXISTS(
id SERIAL PRIMARY KEY,
text text(5000) NOT NULL,
task_id BIGINT UNSIGNED,
CONSTRAINT task_id_fk FOREIGN KEY (task_id)
    REFERENCES tasks(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
