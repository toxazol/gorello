CREATE DATABASE IF NOT EXISTS gorello;
USE gorello;


CREATE TABLE IF NOT EXISTS projects(
id SERIAL PRIMARY KEY,
name varchar(500) NOT NULL,
description varchar(1000)
);


CREATE TABLE IF NOT EXISTS columns(
id SERIAL PRIMARY KEY,
name varchar(255) NOT NULL,
priority double, -- index!
project_id BIGINT UNSIGNED,
CONSTRAINT project_id_fk FOREIGN KEY (project_id)
    REFERENCES projects(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);


CREATE TABLE IF NOT EXISTS tasks(
id SERIAL PRIMARY KEY,
name varchar(500) NOT NULL,
description text(5000),
priority double,
column_id BIGINT UNSIGNED,
CONSTRAINT column_id_fk FOREIGN KEY (column_id)
    REFERENCES columns(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);


CREATE TABLE IF NOT EXISTS comments(
id SERIAL PRIMARY KEY,
text text(5000) NOT NULL,
task_id BIGINT UNSIGNED,
createTs TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
CONSTRAINT task_id_fk FOREIGN KEY (task_id)
    REFERENCES tasks(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
