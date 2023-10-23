-- CREATING AND USING DATABASE
CREATE DATABASE IF NOT EXISTS golang_fiber_mvc;
USE golang_fiber_mvc;


-- CREATING TABLES
DROP TABLE IF EXISTS roles;
CREATE TABLE IF NOT EXISTS roles (
	role_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	role_name VARCHAR(100) NOT NULL,
	code VARCHAR(20) NOT NULL
);

DROP TABLE IF EXISTS users;
CREATE TABLE users (
	user_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	role_id INT NOT NULL,
	user_name VARCHAR(100) NOT NULL,
	password VARCHAR(150) NOT NULL,
	image VARCHAR(100),
	active ENUM('Yes', 'No') DEFAULT 'Yes',
	unique_id VARCHAR(50) UNIQUE,
	created_at DATETIME DEFAULT NOW(),
	updated_at DATETIME DEFAULT NOW(),
	CONSTRAINT user_role FOREIGN KEY(role_id) REFERENCES roles(role_id)
);

DROP TABLE IF EXISTS task_status;
CREATE TABLE task_status (
	status_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	status_name VARCHAR(100) NOT NULL,
	code VARCHAR(20) NOT NULL
);

DROP TABLE IF EXISTS task_complexity;
CREATE TABLE task_complexity (
	complexity_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	complexity_name VARCHAR(100) NOT NULL,
	code VARCHAR(20) NOT NULL
);

DROP TABLE IF EXISTS tasks;
CREATE TABLE IF NOT EXISTS tasks (
	task_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	user_id INT NOT NULL,
	status_id INT NOT NULL,
	complexity_id INT NOT NULL,
	task_name VARCHAR(100) NOT NULL,
	description VARCHAR(200) NOT NULL,
	start_date DATE,
	end_date DATE,
	attachment VARCHAR(100),
	unique_id VARCHAR(50) UNIQUE,
	created_at DATETIME DEFAULT NOW(),
	updated_at DATETIME DEFAULT NOW(),
	CONSTRAINT fk_task_user FOREIGN KEY(user_id) REFERENCES users(user_id),
	CONSTRAINT fk_task_status FOREIGN KEY(status_id) REFERENCES task_status(status_id),
	CONSTRAINT fk_task_complexity FOREIGN KEY(complexity_id) REFERENCES task_complexity(complexity_id)
);


-- VIEWS --------------------------------------
-- View view_user_data
DROP VIEW IF EXISTS view_user_data;
CREATE VIEW view_user_data AS 
SELECT us.user_id, us.unique_id,
	us.user_name, us.password, 
	us.active, us.image,
	DATE_FORMAT(us.created_at, '%Y-%m-%d %H:%i:%s') AS created_at,
	DATE_FORMAT(us.updated_at, '%Y-%m-%d %H:%i:%s') AS updated_at,
	ro.role_id, ro.role_name,
	ro.code AS role_code
FROM users us
JOIN roles ro ON(ro.role_id = us.role_id)
ORDER BY us.created_at DESC;

-- View view_task_data ta.start_date, ta.end_date,
DROP VIEW IF EXISTS view_task_data;
CREATE VIEW view_task_data AS
SELECT ta.task_id, ta.unique_id,
	ta.task_name, ta.description,
	DATE_FORMAT(ta.start_date, '%Y-%m-%d') AS start_date,
	DATE_FORMAT(ta.end_date, '%Y-%m-%d') AS end_date,
	ta.attachment, 
	DATE_FORMAT(ta.created_at, '%Y-%m-%d %H:%i:%s') AS created_at,
	DATE_FORMAT(ta.updated_at, '%Y-%m-%d %H:%i:%s') AS updated_at,
	us.user_id, us.user_name,
	st.status_id, st.status_name,
	st.code AS status_code,
	co.complexity_id, co.complexity_name
FROM tasks ta
JOIN users us ON(us.user_id = ta.user_id)
JOIN task_status st ON(st.status_id = ta.status_id)
JOIN task_complexity co ON(co.complexity_id = ta.complexity_id)
ORDER BY created_at DESC;


-- INSERTS ----------------------------------------------------------------------------
INSERT INTO roles(role_name, code) VALUES("Administrator", "admin");
INSERT INTO roles(role_name, code) VALUES("Normal User", "normal");

INSERT INTO task_status(status_name, code) VALUES("Pending", "pending");
INSERT INTO task_status(status_name, code) VALUES("In Progress", "in-progress");
INSERT INTO task_status(status_name, code) VALUES("Complete", "complete");
INSERT INTO task_status(status_name, code) VALUES("Blocked", "blocked");
INSERT INTO task_status(status_name, code) VALUES("Canceled", "canceled");

INSERT INTO task_complexity(complexity_name, code) VALUES("Easy", "easy");
INSERT INTO task_complexity(complexity_name, code) VALUES("Medium", "medium");
INSERT INTO task_complexity(complexity_name, code) VALUES("Hard", "Hard");
INSERT INTO task_complexity(complexity_name, code) VALUES("Very Hard", "very-hard");
INSERT INTO task_complexity(complexity_name, code) VALUES("Extremely Hard", "extreme-hard");


-- INDEXES ------------------------------

CREATE INDEX idx_role_name ON roles(role_name);
CREATE INDEX idx_role_code ON roles(code);

CREATE INDEX idx_user_name ON users(user_name);
CREATE INDEX idx_active ON users(active);

CREATE INDEX idx_task_name ON tasks(task_name);
CREATE INDEX idx_start_end_date ON tasks(start_date, end_date);

CREATE INDEX idx_status_name ON task_status(status_name);
CREATE INDEX idx_status_code ON task_status(code);

CREATE INDEX idx_complexity_name ON task_complexity(complexity_name);
CREATE INDEX idx_complexity_code ON task_complexity(code);
