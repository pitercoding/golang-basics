package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var tasks []string

	for {
		fmt.Println("\n=== To Do List ===")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Remove Task")
		fmt.Println("0. Exit")

		option := readInt(scanner, "Choose an option: ")

		switch option {
		case 1:
			task := readString(scanner, "Enter task: ")

			if task == "" {
				fmt.Println("Task cannot be empty.")
				continue
			}

			tasks = append(tasks, task)
			fmt.Println("Task added successfully.")

		case 2:
			listTasks(tasks)
		case 3:
			if len(tasks) == 0 {
				fmt.Println("No tasks found.")
				continue
			}

			listTasks(tasks)

			taskNumber := readInt(scanner, "Enter task number to remove: ")

			if taskNumber < 1 || taskNumber > len(tasks) {
				fmt.Println("Invalid task number.")
				continue
			}

			index := taskNumber - 1

			tasks = append(tasks[:index], tasks[index+1:]...)

			fmt.Println("Task removed successfully.")

		case 0:
			fmt.Println("\nAll done! Thank you for using ToDoList.")
			return

		default:
			fmt.Println("\nInvalid option! Try again.")
		}
	}
}

func readInt(scanner *bufio.Scanner, message string) int {
	for {
		fmt.Print(message)

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		value, err := strconv.Atoi(input)
		if err == nil {
			return value
		}

		fmt.Println("Invalid input. Try again.")
	}
}

func readString(scanner *bufio.Scanner, message string) string {
	fmt.Print(message)

	scanner.Scan()

	return strings.TrimSpace(scanner.Text())
}

func listTasks(tasks []string) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found!")
		return
	}

	fmt.Println("\nTasks:")

	for index, task := range tasks {
		fmt.Printf("%d. %s\n", index+1, task)
	}
}
