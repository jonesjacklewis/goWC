package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func is_valid_switch(switch_chosen string) bool {
	valid_switch_map := map[string]bool{
		"-c": true,
		"-l": true,
		"-w": true,
		"-m": true,
	}

	return valid_switch_map[switch_chosen]
}

func get_size_in_bytes_by_filename(filename string) (int64, error) {
	file, err := os.Open(filename)

	if err != nil {
		return -1, fmt.Errorf("Error opening file: %s", filename)
	}

	defer file.Close()

	fi, err := file.Stat()

	if err != nil {
		return -1, fmt.Errorf("Error getting file info: %s", filename)
	}

	return fi.Size(), nil
}

func get_number_of_lines_by_filename(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return -1, fmt.Errorf("Error opening file: %s", filename)
	}

	defer file.Close()

	number_of_lines := 0

	buf := make([]byte, 1024)
	has_content := false

	for {
		n, err := file.Read(buf)

		if n == 0 {
			break
		}

		for i := 0; i < n; i++ {
			if buf[i] == '\n' {
				number_of_lines++
			} else {
				if number_of_lines == 0 {
					has_content = true
				}
			}
		}

		if err != nil {
			break
		}
	}

	if number_of_lines == 0 && has_content {
		number_of_lines = 1
	}

	return number_of_lines, nil
}

func get_number_of_words_by_filename(filename string) (int, error) {

	file, err := os.Open(filename)

	if err != nil {
		return -1, fmt.Errorf("Error opening file: %s", filename)
	}

	defer file.Close()

	number_of_words := 0

	buf := make([]byte, 1024)
	in_word := false

	for {
		n, err := file.Read(buf)

		if n == 0 {
			break
		}

		for i := 0; i < n; i++ {
			at_position := buf[i]

			if unicode.IsSpace(rune(at_position)) {
				if in_word {
					number_of_words++
					in_word = false
				}
			} else {
				in_word = true
			}
		}

		if err != nil {
			break
		}

	}

	if in_word {
		number_of_words++
	}

	return number_of_words, nil
}

func get_number_of_characters_by_filename(filename string) (int, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return -1, fmt.Errorf("Error reading file: %s", filename)
	}

	return len(data), nil
}

func handle_dash_c(os_args []string) (string, error) {

	if len(os_args) != 3 {
		return "", fmt.Errorf("Invalid number of arguments")
	}

	filename := os_args[2]

	file_size, err := get_size_in_bytes_by_filename(filename)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d %s", file_size, filename), nil
}

func handle_dash_l(os_args []string) (string, error) {
	if len(os_args) != 3 {
		return "", fmt.Errorf("Invalid number of arguments")
	}

	filename := os_args[2]

	number_of_lines, err := get_number_of_lines_by_filename(filename)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d %s", number_of_lines, filename), nil
}

func handle_dash_w(os_args []string) (string, error) {
	if len(os_args) != 3 {
		return "", fmt.Errorf("Invalid number of arguments")
	}

	filename := os_args[2]

	number_of_words, err := get_number_of_words_by_filename(filename)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d %s", number_of_words, filename), nil
}

func handle_dash_m(os_args []string) (string, error) {
	if len(os_args) != 3 {
		return "", fmt.Errorf("Invalid number of arguments")
	}

	filename := os_args[2]

	number_of_characters, err := get_number_of_characters_by_filename(filename)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d %s", number_of_characters, filename), nil
}

func handle_default(filename string) (string, error) {

	number_of_lines, err := get_number_of_lines_by_filename(filename)

	if err != nil {
		return "", err
	}

	number_of_words, err := get_number_of_words_by_filename(filename)

	if err != nil {
		return "", err
	}

	size_in_bytes, err := get_size_in_bytes_by_filename(filename)

	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	parts := []string{
		fmt.Sprintf("%d", number_of_lines),
		fmt.Sprintf("%d", number_of_words),
		fmt.Sprintf("%d", size_in_bytes),
		filename,
	}

	tabbed_parts := strings.Join(parts, "\t")

	return tabbed_parts, nil
}

// func main() that takes command line arguments
func main() {
	number_of_args := len(os.Args)

	if number_of_args <= 1 {
		fmt.Println("No arguments provided")
		return
	}

	first_arg := os.Args[1]

	// if switch_chosen starts with a dash
	if first_arg[0] != '-' {
		message, err := handle_default(first_arg)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(message)

		return
	}

	if !is_valid_switch(first_arg) {
		fmt.Printf("Invalid Switch: %s\n", first_arg)
		return
	}

	switch first_arg {
	case "-c":
		message, err := handle_dash_c(os.Args)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(message)

		return
	case "-l":
		message, err := handle_dash_l(os.Args)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(message)

		return
	case "-w":
		message, err := handle_dash_w(os.Args)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(message)

		return
	case "-m":
		message, err := handle_dash_m(os.Args)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(message)
	}

}
