-- ========================================
-- MySQL 小练习
-- ========================================

-- 1. 建表
CREATE DATABASE IF NOT EXISTS practice;
USE practice;

CREATE TABLE departments (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE employees (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    salary DECIMAL(10,2),
    department_id INT,
    FOREIGN KEY (department_id) REFERENCES departments(id)
);

-- 2. 修改表结构
ALTER TABLE employees ADD COLUMN age INT;
ALTER TABLE employees DROP COLUMN age;
ALTER TABLE employees MODIFY COLUMN name VARCHAR(100);

-- 3. 插入数据
INSERT INTO departments (name) VALUES ('研发部'), ('市场部'), ('人事部');
INSERT INTO employees (name, salary, department_id) VALUES
('张三', 8000, 1), ('李四', 12000, 1), ('王五', 9500, 2),
('赵六', 15000, 2), ('孙七', 7000, 3), ('周八', 11000, 3);

-- 4. 查询数据
SELECT * FROM employees;
SELECT * FROM employees WHERE salary > 10000;
SELECT * FROM employees ORDER BY salary DESC;
SELECT * FROM employees LIMIT 3;
SELECT COUNT(*), AVG(salary) FROM employees;

-- 5. 分组查询
SELECT department_id, COUNT(*), AVG(salary) 
FROM employees GROUP BY department_id;

SELECT department_id, AVG(salary) AS avg_sal
FROM employees GROUP BY department_id HAVING avg_sal > 10000;

-- 6. 联表查询
SELECT e.name, d.name 
FROM employees e INNER JOIN departments d ON e.department_id = d.id;

SELECT e.name, d.name 
FROM employees e LEFT JOIN departments d ON e.department_id = d.id;

-- 7. 更新数据
UPDATE employees SET salary = 9000 WHERE id = 1;
UPDATE employees SET salary = salary * 1.1 WHERE department_id = 1;

-- 8. 删除数据
DELETE FROM employees WHERE id = 6;
DELETE FROM employees WHERE salary < 8000;
