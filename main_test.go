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

func Test_handle_dash_c(t *testing.T) {
	t.Run("Invalid number of arguments", func(t *testing.T) {
		os_args := []string{"gowc", "-c"}
		_, err := handle_dash_c(os_args)

		if err == nil {
			t.Errorf("Expected error got nil")
		}
	})

	t.Run("File does not exist", func(t *testing.T) {
		random_file_name := uuid.New().String()
		os_args := []string{"gowc", "-c", fmt.Sprintf("%s.txt", random_file_name)}
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
		fmt.Println(expected_message)

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
