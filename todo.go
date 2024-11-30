package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var selection int = -1

	todos := make([]string, 0)

	// todos = append(todos, "Sample")
	fmt.Println("Welcome to ToDo Program")

	todos = init_file()

	for selection != 0 {
		fmt.Println()
		fmt.Println("---------------------")
		fmt.Println("Select Option")
		fmt.Println("1 - save")
		fmt.Println("2 - list")
		fmt.Println("3 - create")
		fmt.Println("4 - delete")
		fmt.Println("0 - exit")
		fmt.Print("select: ")

		_, err := fmt.Scanln(&selection)

		for err != nil {
			fmt.Println("Error getting selection :", err)
			_, err = fmt.Scanln(&selection)
		}

		if selection == 2 {
			todo_list(todos)
		} else if selection == 3 {
			todos = todo_create(todos)
		} else if selection == 4 {
			todos = todo_delete(todos)
		} else if selection == 1 {
			todo_save(todos)
		}
	}
}

func todo_save(todos []string) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		fmt.Println(err)
		return
	}

	path := filepath.Join(homeDir, "list.todo")

	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		println("hello")
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, line := range todos {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing data:", err)
	}
}

func init_file() []string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	path := filepath.Join(homeDir, "list.todo")

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer file.Close()
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		// fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil
	}

	return lines
}

func todo_list(todos []string) {
	fmt.Println()
	for i := 0; i < len(todos); i++ {
		fmt.Println(i+1, ":", todos[i])
	}
}

func todo_create(todos []string) []string {
	fmt.Println("Enter Todo Title : ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	for err != nil {
		fmt.Println("Error reading input:", err)
		input, err = reader.ReadString('\n')
	}

	input = input[:len(input)-1]

	todos = append(todos, input)
	return todos
}

func todo_delete(todos []string) []string {
	todo_list(todos)

	var selection int
	fmt.Println("Select Todo to delete")
	fmt.Println("0 to cancel")

	_, err := fmt.Scanln(&selection)

	if selection == 0 {
		return todos
	}

	for err != nil || selection > len(todos) || selection < 0 {
		if selection > len(todos) || selection < 0 {
			err = errors.New("out of bounds")
		}

		fmt.Println("Error getting Selection :", err)
		_, err = fmt.Scanln(&selection)
	}

	todos = append(todos[:selection-1], todos[selection:]...)

	return todos
}
