package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"os"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}
type Task struct {
	ID         string
	Title      string
	DueDate    string
	CategoryID string
	IsDone     bool
	UserID     string
}
type Category struct {
	ID     string
	Title  string
	Color  string
	UserID string
}

var userStorage []User
var taskStorage []Task
var categoryStorage []Category

var authenticatedUser *User

func main() {
	fmt.Println("Hello TODO app")
	command := flag.String("command", "no command", "command to run")
	flag.Parse()

	for {
		runCommand(*command)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		*command = scanner.Text()
	}

}

func runCommand(command string) {
	if command != "register" && command != "exit" && authenticatedUser == nil {
		login()
		if authenticatedUser == nil {
			return
		}
	}

	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "list-task":
		listTask()
	case "register":
		register()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Command not found")
	}

}

func createTask() {

	scanner := bufio.NewScanner(os.Stdin)
	id := uuid.New()
	var title, dueDate, category string

	fmt.Println("Please enter the task title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the task dueDate")
	scanner.Scan()
	dueDate = scanner.Text()

	fmt.Println("please enter the task category ID")
	scanner.Scan()
	category = scanner.Text()
	isFound := false
	for _, c := range categoryStorage {
		if c.ID == category && c.UserID == authenticatedUser.ID {
			isFound = true
			break
		}
	}
	if !isFound {
		fmt.Println("category ID is not found\n")

		return
	}
	task := Task{
		ID:         id.String(),
		Title:      title,
		DueDate:    dueDate,
		CategoryID: category,
		IsDone:     false,
		UserID:     authenticatedUser.ID,
	}

	taskStorage = append(taskStorage, task)

	fmt.Println("task", title, category, dueDate)
}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	id := uuid.New()
	var title, color string

	fmt.Println("Please enter the category title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("Please enter the category color")
	scanner.Scan()
	color = scanner.Text()

	category := Category{
		ID:     id.String(),
		Title:  title,
		Color:  color,
		UserID: authenticatedUser.ID,
	}

	categoryStorage = append(categoryStorage, category)
	fmt.Println("category", title, color, id.String())
}

func register() {
	scanner := bufio.NewScanner(os.Stdin)
	var email, name, password string
	id := uuid.New()

	fmt.Println("Please enter the user email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("Please enter the user name")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("Please enter the user password")
	scanner.Scan()
	password = scanner.Text()

	user := User{
		ID:       id.String(),
		Name:     name,
		Email:    email,
		Password: password,
	}

	userStorage = append(userStorage, user)

	fmt.Println("User", id.String(), email, password)
}

func login() {
	fmt.Println("You must log in first")

	scn := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the email")
	scn.Scan()
	email := scn.Text()

	fmt.Println("Please enter the password")
	scn.Scan()
	password := scn.Text()

	for _, user := range userStorage {
		if user.Email == email && user.Password == password {
			authenticatedUser = &user

			break
		}
	}

	if authenticatedUser == nil {
		fmt.Println("the email or password is  not correct")
	}

}
func listTask() {
	for _, task := range taskStorage {
		if task.UserID == authenticatedUser.ID {
			fmt.Println(task)
		}
	}
}
