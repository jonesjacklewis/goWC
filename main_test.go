package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
)

func Test_is_valid_switch(t *testing.T) {
	t.Run("Valid Switch", func(t *testing.T) {
		switch_chosen := "-c"
		is_valid := is_valid_switch(switch_chosen)

		if !is_valid {
			t.Errorf("Got false should be true")
		}
	})

	t.Run("Invlid Switch", func(t *testing.T) {
		switch_chosen := "-hello"
		is_valid := is_valid_switch(switch_chosen)

		if is_valid {
			t.Errorf("Got true should be false")
		}
	})
}

func Test_get_size_in_bytes_by_filename(t *testing.T) {
	t.Run("File does not exist", func(t *testing.T) {
		random_file_name := uuid.New().String()
		_, err := get_size_in_bytes_by_filename(fmt.Sprintf("%s.txt", random_file_name))

		if err == nil {
			t.Errorf("Expected error got nil")
		}
	})

	t.Run("File is empty", func(t *testing.T) {
		random_file_name := uuid.New().String()

		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		file.Close()

		size, err := get_size_in_bytes_by_filename(filename)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		if size != 0 {
			t.Errorf("Expected 0 got %d", size)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}
	})

	t.Run("File is not empty", func(t *testing.T) {
		random_file_name := uuid.New().String()

		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		content_to_write := "Hello, World!"

		_, err = file.WriteString(content_to_write)

		if err != nil {
			t.Errorf("Error writing to file: %s", filename)
		}

		file.Close()

		size, err := get_size_in_bytes_by_filename(filename)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		if size != int64(len(content_to_write)) {
			t.Errorf("Expected %d got %d", len(content_to_write), size)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}
	})
}

func Test_get_number_of_lines_by_filename(t *testing.T) {
	t.Run("File does not exist", func(t *testing.T) {
		random_file_name := uuid.New().String()
		_, err := get_number_of_lines_by_filename(fmt.Sprintf("%s.txt", random_file_name))

		if err == nil {
			t.Errorf("Expected error got nil")
		}
	})

	t.Run("File is empty", func(t *testing.T) {
		random_file_name := uuid.New().String()

		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		file.Close()

		number_of_lines, err := get_number_of_lines_by_filename(filename)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		if number_of_lines != 0 {
			t.Errorf("Expected 0 got %d", number_of_lines)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}
	})

	t.Run("File is not empty", func(t *testing.T) {
		random_file_name := uuid.New().String()

		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		content_to_write := "Hello, World!\nHello, World!\nHello, World!\n"

		_, err = file.WriteString(content_to_write)

		if err != nil {
			t.Errorf("Error writing to file: %s", filename)
		}

		file.Close()

		number_of_lines, err := get_number_of_lines_by_filename(filename)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		if number_of_lines != 3 {
			t.Errorf("Expected 3 got %d", number_of_lines)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}
	})

	t.Run("File contains one line", func(t *testing.T) {
		random_file_name := uuid.New().String()

		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		content_to_write := "Hello, World!"

		_, err = file.WriteString(content_to_write)

		if err != nil {
			t.Errorf("Error writing to file: %s", filename)
		}

		file.Close()

		number_of_lines, err := get_number_of_lines_by_filename(filename)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		if number_of_lines != 1 {
			t.Errorf("Expected 1 got %d", number_of_lines)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}
	})
}

func Test_get_number_of_words_by_filename(t *testing.T) {
	t.Run("File does not exist", func(t *testing.T) {
		random_file_name := uuid.New().String()
		_, err := get_number_of_words_by_filename(fmt.Sprintf("%s.txt", random_file_name))

		if err == nil {
			t.Errorf("Expected error got nil")
		}
	})

	t.Run("File is empty", func(t *testing.T) {
		random_file_name := uuid.New().String()

		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		file.Close()

		number_of_words, err := get_number_of_words_by_filename(filename)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		if number_of_words != 0 {
			t.Errorf("Expected 0 got %d", number_of_words)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}
	})

	t.Run("File is not empty", func(t *testing.T) {
		random_file_name := uuid.New().String()

		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		content_to_write := "Hello World"

		_, err = file.WriteString(content_to_write)

		if err != nil {
			t.Errorf("Error writing to file: %s", filename)
		}

		file.Close()

		number_of_words, err := get_number_of_words_by_filename(filename)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		if number_of_words != 2 {
			t.Errorf("Expected 2 got %d", number_of_words)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}
	})

}

func Test_get_number_of_characters_by_filename(t *testing.T) {
	t.Run("File does not exist", func(t *testing.T) {
		random_file_name := uuid.New().String()
		_, err := get_number_of_characters_by_filename(fmt.Sprintf("%s.txt", random_file_name))

		if err == nil {
			t.Errorf("Expected error got nil")
		}
	})

	t.Run("File is empty", func(t *testing.T) {
		random_file_name := uuid.New().String()

		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		file.Close()

		number_of_characters, err := get_number_of_characters_by_filename(filename)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		if number_of_characters != 0 {
			t.Errorf("Expected 0 got %d", number_of_characters)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}
	})

	t.Run("File is not empty", func(t *testing.T) {
		random_file_name := uuid.New().String()

		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		content_to_write := "Hello World"

		_, err = file.WriteString(content_to_write)

		if err != nil {
			t.Errorf("Error writing to file: %s", filename)
		}

		file.Close()

		number_of_characters, err := get_number_of_characters_by_filename(filename)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		if number_of_characters != 11 {
			t.Errorf("Expected 11 got %d", number_of_characters)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}
	})

	t.Run("Multibyte Characters", func(t *testing.T) {
		random_file_name := uuid.New().String()

		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		content_to_write := "Hello, 世界"

		_, err = file.WriteString(content_to_write)

		if err != nil {
			t.Errorf("Error writing to file: %s", filename)
		}

		file.Close()

		number_of_characters, err := get_number_of_characters_by_filename(filename)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		if number_of_characters != 13 {
			t.Errorf("Expected 13 got %d", number_of_characters)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}
	})
}

func Test_handle_dash_c(t *testing.T) {
	t.Run("Invalid number of arguments", func(t *testing.T) {
		os_args := []string{"gowc", "-c"}
		_, err := handle_dash_c(os_args)

		if err == nil {
			t.Errorf("Expected error got nil")
		}
	})

	t.Run("File is empty", func(t *testing.T) {
		random_file_name := uuid.New().String()
		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		file.Close()

		os_args := []string{"gowc", "-c", filename}
		message, err := handle_dash_c(os_args)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		expected_message := fmt.Sprintf("0 %s", filename)

		if message != expected_message {
			t.Errorf("Expected %s got %s", expected_message, message)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}

	})

	t.Run("File has content", func(t *testing.T) {
		content_to_write := "Hello, World!"
		random_file_name := uuid.New().String()
		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		_, err = file.WriteString(content_to_write)

		if err != nil {
			t.Errorf("Error writing to file: %s", filename)
		}

		file.Close()

		os_args := []string{"gowc", "-c", filename}
		message, err := handle_dash_c(os_args)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		expected_message := fmt.Sprintf("%d %s", len(content_to_write), filename)

		if message != expected_message {
			t.Errorf("Expected %s got %s", expected_message, message)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}
	})
}

func Test_handle_dash_l(t *testing.T) {
	t.Run("Invalid number of arguments", func(t *testing.T) {
		os_args := []string{"gowc", "-l"}
		_, err := handle_dash_l(os_args)

		if err == nil {
			t.Errorf("Expected error got nil")
		}
	})

	t.Run("File is empty", func(t *testing.T) {
		random_file_name := uuid.New().String()
		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		file.Close()

		os_args := []string{"gowc", "-l", filename}
		message, err := handle_dash_l(os_args)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		expected_message := fmt.Sprintf("0 %s", filename)

		if message != expected_message {
			t.Errorf("Expected %s got %s", expected_message, message)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}

	})

	t.Run("File is not empty", func(t *testing.T) {
		random_file_name := uuid.New().String()
		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		_, err = file.WriteString("Hello, World!\nHello, World!\nHello, World!\n")

		if err != nil {
			t.Errorf("Error writing to file: %s", filename)
		}

		file.Close()

		os_args := []string{"gowc", "-l", filename}
		message, err := handle_dash_l(os_args)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		expected_message := fmt.Sprintf("3 %s", filename)

		if message != expected_message {
			t.Errorf("Expected %s got %s", expected_message, message)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}

	})

	t.Run("File contains one line", func(t *testing.T) {
		random_file_name := uuid.New().String()
		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		_, err = file.WriteString("Hello, World!")

		if err != nil {
			t.Errorf("Error writing to file: %s", filename)
		}

		file.Close()

		os_args := []string{"gowc", "-l", filename}
		message, err := handle_dash_l(os_args)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		expected_message := fmt.Sprintf("1 %s", filename)

		if message != expected_message {
			t.Errorf("Expected %s got %s", expected_message, message)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}

	})

}

func Test_handle_dash_w(t *testing.T) {
	t.Run("Invalid number of arguments", func(t *testing.T) {
		os_args := []string{"gowc", "-w"}
		_, err := handle_dash_w(os_args)

		if err == nil {
			t.Errorf("Expected error got nil")
		}
	})

	t.Run("File is empty", func(t *testing.T) {
		random_file_name := uuid.New().String()
		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		file.Close()

		os_args := []string{"gowc", "-w", filename}
		message, err := handle_dash_w(os_args)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		expected_message := fmt.Sprintf("0 %s", filename)

		if message != expected_message {
			t.Errorf("Expected %s got %s", expected_message, message)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}

	})

	t.Run("File is not empty", func(t *testing.T) {
		random_file_name := uuid.New().String()
		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		_, err = file.WriteString("Hello World")

		if err != nil {
			t.Errorf("Error writing to file: %s", filename)
		}

		file.Close()

		os_args := []string{"gowc", "-w", filename}
		message, err := handle_dash_w(os_args)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		expected_message := fmt.Sprintf("2 %s", filename)

		if message != expected_message {
			t.Errorf("Expected %s got %s", expected_message, message)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}

	})

}

func Test_handle_dash_m(t *testing.T) {
	t.Run("Invalid number of arguments", func(t *testing.T) {
		os_args := []string{"gowc", "-m"}
		_, err := handle_dash_m(os_args)

		if err == nil {
			t.Errorf("Expected error got nil")
		}
	})

	t.Run("File is empty", func(t *testing.T) {
		random_file_name := uuid.New().String()
		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		file.Close()

		os_args := []string{"gowc", "-w", filename}
		message, err := handle_dash_m(os_args)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		expected_message := fmt.Sprintf("0 %s", filename)

		if message != expected_message {
			t.Errorf("Expected %s got %s", expected_message, message)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}

	})

	t.Run("File is not empty", func(t *testing.T) {
		random_file_name := uuid.New().String()
		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		_, err = file.WriteString("Hello World")

		if err != nil {
			t.Errorf("Error writing to file: %s", filename)
		}

		file.Close()

		os_args := []string{"gowc", "-w", filename}
		message, err := handle_dash_m(os_args)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		expected_message := fmt.Sprintf("11 %s", filename)

		if message != expected_message {
			t.Errorf("Expected %s got %s", expected_message, message)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}

	})

	t.Run("Multibyte Characters", func(t *testing.T) {
		random_file_name := uuid.New().String()
		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		_, err = file.WriteString("Hello, 世界")

		if err != nil {
			t.Errorf("Error writing to file: %s", filename)
		}

		file.Close()

		os_args := []string{"gowc", "-w", filename}
		message, err := handle_dash_m(os_args)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		expected_message := fmt.Sprintf("13 %s", filename)

		if message != expected_message {
			t.Errorf("Expected %s got %s", expected_message, message)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}

	})

}

func Test_handle_default(t *testing.T) {
	t.Run("File does not exist", func(t *testing.T) {
		random_file_name := uuid.New().String()
		_, err := handle_default(fmt.Sprintf("%s.txt", random_file_name))

		if err == nil {
			t.Errorf("Expected error got nil")
		}
	})

	t.Run("File is empty", func(t *testing.T) {
		random_file_name := uuid.New().String()
		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		file.Close()

		message, err := handle_default(filename)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		expected_message := fmt.Sprintf("0\t0\t0\t%s", filename)

		if message != expected_message {
			t.Errorf("Expected %s got %s", expected_message, message)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}
	})

	t.Run("File is not empty", func(t *testing.T) {
		random_file_name := uuid.New().String()
		filename := fmt.Sprintf("%s.txt", random_file_name)

		// Create a file with the random_file_name
		file, err := os.Create(filename)

		if err != nil {
			t.Errorf("Error creating file: %s", filename)
		}

		_, err = file.WriteString("Hello, World!\nHello, World!\nHello, World!\n")

		if err != nil {
			t.Errorf("Error writing to file: %s", filename)
		}

		file.Close()

		message, err := handle_default(filename)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		expected_message := fmt.Sprintf("3\t6\t42\t%s", filename)

		if message != expected_message {
			t.Errorf("Expected %s got %s", expected_message, message)
		}

		// Remove the file
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %s", filename)
		}
	})
}
