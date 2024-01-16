DROP TABLE IF EXISTS fiber_storage;
CREATE TABLE IF NOT EXISTS fiber_storage (
  k varchar(64) NOT NULL DEFAULT '',
  v blob NOT NULL,
  e bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (k)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;


INSERT INTO fiber_storage (k, v, e) VALUES
('033b50c0-13cd-486f-a9bc-4140a7f4f582', 0x0d7f040102ff8000010c0110000075ff80000208757365726e616d6506737472696e670c0f000d7465737440757365722e636f6d0870617373776f726406737472696e670c3e003c24326124313024416c51553943363465516769584754636e322f674c75737a4a5766773331566b506b50345449364f70674b6a6d7a53543668312f61, 1700575203);

DROP TABLE IF EXISTS roles;
CREATE TABLE IF NOT EXISTS roles (
  role_id int NOT NULL AUTO_INCREMENT,
  role_name varchar(100) NOT NULL,
  code varchar(20) NOT NULL,
  PRIMARY KEY (role_id),
  KEY idx_role_name (role_name),
  KEY idx_role_code (code)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


INSERT INTO roles (role_id, role_name, code) VALUES
(1, 'Administrator', 'admin'),
(2, 'Normal User', 'normal');


DROP TABLE IF EXISTS tasks;
CREATE TABLE IF NOT EXISTS tasks (
  task_id int NOT NULL AUTO_INCREMENT,
  user_id int NOT NULL,
  status_id int NOT NULL,
  complexity_id int NOT NULL,
  task_name varchar(100) NOT NULL,
  description varchar(200) NOT NULL,
  start_date date DEFAULT NULL,
  end_date date DEFAULT NULL,
  attachment varchar(100) DEFAULT NULL,
  unique_id varchar(50) DEFAULT NULL,
  created_at datetime DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (task_id),
  UNIQUE KEY unique_id (unique_id),
  KEY fk_task_user (user_id),
  KEY fk_task_status (status_id),
  KEY fk_task_complexity (complexity_id),
  KEY idx_task_name (task_name),
  KEY idx_start_end_date (start_date,end_date)
);


INSERT INTO tasks (task_id, user_id, status_id, complexity_id, task_name, description, start_date, end_date, attachment, unique_id, created_at, updated_at) VALUES
(1, 9, 1, 1, 'Testing App', 'aaa', '2023-11-10', '2023-11-16', '', '91eb19e1-4171-41ad-aec2-57822d43a253', '2023-11-20 15:00:23', '2023-11-20 15:00:23'),
(4, 9, 1, 2, 'Upload Code to Github', 'Uploading code on Github', '2023-01-01', '2023-02-04', '', '9e207f57-4bf2-41ed-84d9-ac6a411cf33b', '2023-11-20 15:58:21', '2023-11-20 15:58:21'),
(5, 9, 1, 2, 'Upload Code to Github', 'Uploading code on Github', '2023-01-01', '2023-02-04', '', 'bf763e27-4738-40f1-87c3-ac28d7a8d3c0', '2023-11-20 15:58:21', '2023-11-20 15:58:21'),
(6, 9, 1, 2, 'Upload Code to Github', 'Uploading code on Github', '2023-01-01', '2023-02-04', '', '585ee826-e349-47a5-927e-9c0c8e7dd617', '2023-11-20 15:58:21', '2023-11-20 15:58:21'),
(7, 9, 1, 2, 'Upload Code to Github', 'Uploading code on Github', '2023-01-01', '2023-02-04', '', '4021432e-cb3f-40ec-ab56-5c6541a70a33', '2023-11-20 15:58:21', '2023-11-20 15:58:21'),
(8, 9, 1, 2, 'Upload Code to Github', 'Uploading code on Github', '2023-01-01', '2023-02-04', '', '37fd3123-5ab7-418e-a3e4-70cee78e1c5c', '2023-11-20 15:58:21', '2023-11-20 15:58:21'),
(9, 9, 1, 2, 'Upload Code to Github', 'Uploading code on Github', '2023-01-01', '2023-02-04', '', '3b109ee6-73c6-466d-9695-493d35e22ea3', '2023-11-20 15:58:21', '2023-11-20 15:58:21'),
(10, 9, 1, 2, 'Upload Code to Github', 'Uploading code on Github', '2023-01-01', '2023-02-04', '', 'c2e71b78-eb39-470b-a5cc-439b9a024528', '2023-11-20 15:58:21', '2023-11-20 15:58:21'),
(11, 9, 1, 2, 'Upload Code to Github', 'Uploading code on Github', '2023-01-01', '2023-02-04', '', '277052a0-5178-4a3b-80cd-ee29977e9dcb', '2023-11-20 15:58:21', '2023-11-20 15:58:21'),
(12, 9, 1, 2, 'Upload Code to Github', 'Uploading code on Github', '2023-01-01', '2023-02-04', '', '43c635ad-fa23-4908-916b-6fba745251be', '2023-11-20 15:58:21', '2023-11-20 15:58:21');


DROP TABLE IF EXISTS task_complexity;
CREATE TABLE IF NOT EXISTS task_complexity (
  complexity_id int NOT NULL AUTO_INCREMENT,
  complexity_name varchar(100) NOT NULL,
  code varchar(20) NOT NULL,
  PRIMARY KEY (complexity_id),
  KEY idx_complexity_name (complexity_name),
  KEY idx_status_ode (code)
);


INSERT INTO task_complexity (complexity_id, complexity_name, code) VALUES
(1, 'Easy', 'easy'),
(2, 'Medium', 'medium'),
(3, 'Hard', 'Hard'),
(4, 'Very Hard', 'very-hard'),
(5, 'Extremely Hard', 'extreme-hard');



DROP TABLE IF EXISTS task_status;
CREATE TABLE IF NOT EXISTS task_status (
  status_id int NOT NULL AUTO_INCREMENT,
  status_name varchar(100) NOT NULL,
  code varchar(20) NOT NULL,
  PRIMARY KEY (status_id),
  KEY idx_status_name (status_name),
  KEY idx_status_code (code)
);


INSERT INTO task_status (status_id, status_name, code) VALUES
(1, 'Pending', 'pending'),
(2, 'In Progress', 'in-progress'),
(3, 'Complete', 'complete'),
(4, 'Blocked', 'blocked'),
(5, 'Canceled', 'canceled');


DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users (
  user_id int NOT NULL AUTO_INCREMENT,
  role_id int NOT NULL,
  user_name varchar(100) NOT NULL,
  password varchar(150) NOT NULL,
  image varchar(100) DEFAULT NULL,
  active enum('Yes','No') DEFAULT 'Yes',
  unique_id varchar(50) DEFAULT NULL,
  created_at datetime DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT CURRENT_TIMESTAMP,
  token varchar(150) DEFAULT NULL,
  PRIMARY KEY (user_id),
  UNIQUE KEY unique_id (unique_id),
  KEY user_role (role_id),
  KEY idx_user_name (user_name),
  KEY idx_active (active)
) ;


INSERT INTO users (user_id, role_id, user_name, password, image, active, unique_id, created_at, updated_at, token) VALUES
(7, 1, 'admin@gmail.com', '$2a$10$WK73KU34gno.h1TqJFLrmux5uVIrNwS5TfgKxLcKxeSO15DP.McwO', '', 'Yes', '01511ee6-c7f9-4d9d-8026-0ddc8c114369', '2023-11-20 14:53:01', '2023-11-20 14:59:21', 'wBEO2VSYUQ--1kAzI3pz_vlLz9vPjAS8YTqVU6Qlh4ph2KnC7Rzy0pM192fwJyAxYt9Ypn-m_M8YojOWXGINvoqfgsT5xWYzRjTlhGgRosMoq90llyPr6dUu3KjX5kNbYRbXPw'),
(8, 2, 'normal@gmail.com', '$2a$10$Rb44LaGqdM9R4Lx3zg59Z.bZGAlP05OGU5cR9Vni7W35EksJOuW/a', '', 'Yes', 'e9dce919-35d9-4e5e-8486-612d82a2882f', '2023-11-20 14:55:39', '2023-11-20 14:58:53', 'hxlzIVixinijQwgVioq94oc-4DKvSM-ZolJx7rq3VquLHqgEoBXvSTpUDJTQ0VJSSSupvLqSmDmykxRkgGkkHLkOTRXaaJTXSBU0aI5mWfcHAr1_ondZ_aG1H6ohZ3OsefrU1g'),
(9, 2, 'test@user.com', '$2a$10$AlQU9C64eQgiXGTcn2/gLuszJWfw31VkPkP4TI6OpgKjmzST6h1/a', '', 'Yes', '466a518a-0630-486e-97fb-96581fdfd85e', '2023-11-20 14:59:51', '2023-11-20 15:00:03', 'TfqOwz8uCmyx0YxqIea249gJtkvL3qb2McjXwCjZq1ArSxU8zbCqR0h5bdHTJpmy-V3MWtxeKAYeACAVgPbUtPQ0MKe3M-Pug5rTAR9qYDe_TOGPKQ2LvTXB0yTAyF5ukIi5yA');


ALTER TABLE tasks
  ADD CONSTRAINT fk_task_complexity FOREIGN KEY (complexity_id) REFERENCES task_complexity (complexity_id),
  ADD CONSTRAINT fk_task_status FOREIGN KEY (status_id) REFERENCES task_status (status_id),
  ADD CONSTRAINT fk_task_user FOREIGN KEY (user_id) REFERENCES users (user_id);

ALTER TABLE users
  ADD CONSTRAINT user_role FOREIGN KEY (role_id) REFERENCES roles (role_id);
COMMIT;


DROP VIEW IF EXISTS view_task_data;
CREATE VIEW view_task_data AS 
SELECT 
    ta.task_id AS task_id,
    ta.unique_id AS unique_id,
    ta.task_name AS task_name,
    ta.description AS description,
    DATE_FORMAT(ta.start_date, '%Y-%m-%d') AS start_date,
    DATE_FORMAT(ta.end_date, '%Y-%m-%d') AS end_date,
    ta.attachment AS attachment,
    DATE_FORMAT(ta.created_at, '%Y-%m-%d %H:%i:%s') AS created_at,
    DATE_FORMAT(ta.updated_at, '%Y-%m-%d %H:%i:%s') AS updated_at,
    us.user_id AS user_id,
    us.user_name AS user_name,
    st.status_id AS status_id,
    st.status_name AS status_name,
    st.code AS status_code,
    co.complexity_id AS complexity_id,
    co.complexity_name AS complexity_name
FROM 
    tasks ta 
    JOIN users us ON (us.user_id = ta.user_id) 
    JOIN task_status st ON (st.status_id = ta.status_id)
    JOIN task_complexity co ON (co.complexity_id = ta.complexity_id)
ORDER BY ta.created_at DESC;


DROP VIEW IF EXISTS view_user_data;
CREATE VIEW view_user_data AS 
SELECT 
    us.user_id AS user_id,
    us.unique_id AS unique_id,
    us.user_name AS user_name,
    us.password AS password,
    us.active AS active,
    us.image AS image,
    DATE_FORMAT(us.created_at, '%Y-%m-%d %H:%i:%s') AS created_at,
    DATE_FORMAT(us.updated_at, '%Y-%m-%d %H:%i:%s') AS updated_at,
    ro.role_id AS role_id,
    ro.role_name AS role_name,
    ro.code AS role_code
FROM 
    users us 
    JOIN roles ro ON (ro.role_id = us.role_id)
ORDER BY us.created_at DESC;
