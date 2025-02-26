package main

import (
	"fmt"
	"os"
)

func is_valid_switch(switch_chosen string) bool {
	switch switch_chosen {
	case "-c":
		return true
	}
	return false
}

func handle_dash_c(os_args []string) (string, error) {

	if len(os_args) != 3 {
		return "", fmt.Errorf("Invalid number of arguments")
	}

	filename := os_args[2]

	file, err := os.Open(filename)

	if err != nil {
		return "", fmt.Errorf("Error opening file: %s", filename)
	}

	defer file.Close()

	fi, err := file.Stat()

	if err != nil {
		return "", fmt.Errorf("Error getting file info: %s", filename)
	}

	file_size := fi.Size()

	return fmt.Sprintf("%d %s", file_size, filename), nil
}

// func main() that takes command line arguments
func main() {
	number_of_args := len(os.Args)

	if number_of_args <= 1 {
		fmt.Println("No arguments provided")
		return
	}

	switch_chosen := os.Args[1]

	if !is_valid_switch(switch_chosen) {
		fmt.Printf("Invalid Switch: %s\n", switch_chosen)
		return
	}

	switch switch_chosen {
	case "-c":
		message, err := handle_dash_c(os.Args)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(message)

		return
	}

}
