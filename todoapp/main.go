package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var id int = 0

var DB *sql.DB

type Todo struct {
	title           string
	taskDescription string
	user            *user
}

type user struct {
	username string
}

var Todos []Todo

func initDB() {
	var err error
	DB, err = sql.Open("sqlite3", "file:app.db?_foreign_keys=on")
	if err != nil {
		log.Fatal(err)
	}

	sqluser := `CREATE TABLE IF NOT EXISTS user (
	     id INTEGER PRIMARY KEY AUTOINCREMENT,
	     username TEXT UNIQUE NOT NULL
	);`

	sqltodo := `
	 CREATE TABLE IF NOT EXISTS todos (
	  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	  title TEXT,
	  taskDescription TEXT,
	  user_username TEXT,
	  FOREIGN KEY (user_username) REFERENCES user(username)
	);`

	_, err = DB.Exec(sqluser)
	if err != nil {

		log.Fatal(err)
	} else {
		_, err = DB.Exec(sqltodo)
		if err != nil {
			log.Fatalf("Error creating table: %q: %s\n", err, sqltodo)
		}
	}
}
func deletetodoByID() {

	var taskId int
	fmt.Print("Enter task id to delete:")
	fmt.Scanln(&taskId)

	sqlquery, err := DB.Exec("DELETE FROM todos WHERE id = ?;", taskId)
	if err != nil {
		log.Fatalf("couldnt delete the task: %s", err)
	}
	rowsAffected, err := sqlquery.RowsAffected()
	if rowsAffected > 0 {
		fmt.Println("Task deleted successfully.")
	} else {
		fmt.Println("No task found with the provided ID.")
	}

	// for i, todo := range Todos {
	// 	if todo.id == taskId {
	// 		Todos = append(Todos[:i], Todos[i+1:]...)
	// 		fmt.Println("Task deleted successfuly")
	// 		return
	// 	}
	// }
	// fmt.Printf("Task with id %d not found ", taskId)
}

func updatetodo() {

}

func createNewTodo() {
	var task string
	var userName string
	fmt.Printf("Enter Task name: ")
	fmt.Scanln(&task)
	fmt.Printf("Enter Description: ")
	descriptionReader := bufio.NewReader(os.Stdin)
	description, _ := descriptionReader.ReadString('\n')
	description = strings.TrimSpace(description)

	fmt.Printf("User Name:")
	fmt.Scanln(&userName)
	addUser := &user{username: userName}

	value, err := DB.Query("SELECT username FROM user WHERE username= ?", userName)
	defer value.Close()

	var user string
	for value.Next() {
		if err := value.Scan(&user); err != nil {
			fmt.Println("user not found", err)
		}

	}

	if err != nil {
		log.Fatal(err)
	} else if user == "" {
		fmt.Printf("User %s doesnt exist kindly register the user first\n", userName)
	} else {
		newTodo := Todo{title: task, taskDescription: description, user: addUser}

		Todos = append(Todos, newTodo)
		_, err := DB.Exec("INSERT INTO todos( title, taskDescription, user_username) VALUES( ?, ?, ?)", newTodo.title, newTodo.taskDescription, newTodo.user.username)
		if err != nil {
			fmt.Printf("Error couldnt update in the table: %s", err)
		} else {
			fmt.Println("Task updated successfully")
		}
	}

}

func listAllTodos() {

	for {

		sqlquery, err := DB.Query("SELECT *  FROM todos ")
		fmt.Println(sqlquery)
		if err != nil {
			fmt.Println("err:", err)
		}
		defer sqlquery.Close()
		type Todo_custom struct {
			id              int
			title           string
			taskDescription string
			user            string
		}
		var todos []Todo_custom
		for sqlquery.Next() {
			var todo Todo_custom
			//var userName string
			err := sqlquery.Scan(&todo.id, &todo.title, &todo.taskDescription, &todo.user)
			if err != nil {
				log.Fatal(err)
				return
			}
			todos = append(todos, todo)

		}
		fmt.Println(todos)
		if err := sqlquery.Err(); err != nil {
			log.Fatal(err)
			return
		}

		if len(todos) > 0 {
			for _, item := range todos {
				fmt.Println("Task ID:", item.id)
				fmt.Println("Task title:", item.title)
				fmt.Println("Task Description:", item.taskDescription)
				fmt.Println("Task User:", item.user)
				fmt.Println("------------------------------------------------------------")
			}
			break

		} else {
			fmt.Println("No tasks found.")
			break
		}

	}

}

func registerAuser() {
	var newUser string
	fmt.Print("Enter a username to register: ")
	fmt.Scanln(&newUser)
	newUser = strings.ReplaceAll(newUser, " ", "")

	value, err := DB.Query("SELECT username FROM user WHERE username= ?", newUser)
	defer value.Close()

	var user string
	for value.Next() {
		if err := value.Scan(&user); err != nil {
			fmt.Println("user not found", err)
		}

	}

	if user != "" {
		fmt.Println("User already exist")
	} else if err != nil {
		log.Fatalf("Error accessing DB: %s", err)
	} else {
		_, err := DB.Exec("INSERT INTO user(username) VALUES(?)", newUser)
		if err != nil {
			log.Fatalf("Couldn't Add user. %s", err)
		} else {
			fmt.Printf("User %s registered\n", newUser)
		}

	}

	fmt.Println("------------------------------------------------------------")
}

func main() {
	fmt.Println("Checking DB ")
	initDB()
	defer DB.Close()
	for {
		fmt.Println("")

		fmt.Println("Select option to continue:")
		fmt.Println("1.List all to do task:")
		fmt.Println("2.create a new task")
		fmt.Println("3.Update a task")
		fmt.Println("4.Delete a task")
		fmt.Println("5.Register A user")
		fmt.Println("")
		fmt.Println("------------------------------------------------------------")
		var first string

		fmt.Scanln(&first)

		switch first {
		case "1":
			listAllTodos()
		case "2":
			createNewTodo()
		case "3":
			updatetodo()
		case "4":
			deletetodoByID()
		case "5":
			registerAuser()
		default:
			fmt.Println("insert the correct number to perform operation ")

		}

	}
}
