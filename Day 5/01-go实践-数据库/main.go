package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	//dsn格式
	dsn := "root:0307@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//设置数据库连接池
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	//测试数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}

	fmt.Println("数据库连接成功!")

	//创建表
	createTable(db)

	//插入数据
	insterUser(db, User{Name: "张三", Age: 18})
	insterUser(db, User{Name: "李四", Age: 20})
	insterUser(db, User{Name: "王五", Age: 22})

	//查询用户
	users, err := queryUsers(db)
	if err != nil {
		log.Fatal("查询用户失败", err)
	}
	fmt.Println("查询结果:", users)

	//更新数据
	updateUser(db, 1, "张三丰", 25)

	//删除数据
	deleteUser(db, 3)

	//再次查询用户
	users, err = queryUsers(db)
	if err != nil {
		log.Fatal("查询用户失败", err)
	}
	fmt.Println("更新后的结果:", users)
}

func insterUser(db *sql.DB, user User) {
	result, err := db.Exec("INSERT into users(name,age) VALUES (?,?)", user.Name, user.Age)
	if err != nil {
		log.Printf("插入数据失败", err)
		return
	}

	id, _ := result.LastInsertId()
	fmt.Printf("插入成功，ID为: %d\n", id)
}

func createTable(db *sql.DB) {
	sql := "create table if not exists users(id int primary key auto_increment,name varchar(20),age int)"

	_, err := db.Exec(sql)
	if err != nil {
		log.Fatal("创建表失败", err)
	}
	fmt.Println("创建表成功!")
}

func queryUsers(db *sql.DB) ([]User, error) {
	var users []User
	rows, err := db.Query("select id,name,age from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//遍历查询结果
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func updateUser(db *sql.DB, id int, name string, age int) {
	result, err := db.Exec("UPDATE users SET name=?, age=? WHERE id=?", name, age, id)
	if err != nil {
		log.Printf("更新数据失败: %v", err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("更新成功，影响了 %d 行\n", rowsAffected)
}

func deleteUser(db *sql.DB, id int) {
	result, err := db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		log.Printf("删除数据失败: %v", err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("删除成功，影响了 %d 行\n", rowsAffected)
}
