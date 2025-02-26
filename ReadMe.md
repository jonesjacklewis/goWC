# 📝 gowc

A simple command-line utility in Go that replicates the basic functionality of `wc` (word count) for counting file size in bytes and the number of lines in a file.

## 🚀 Features
- Count the size (in bytes) of a given file using the `-c` flag.
- Count the number of lines in a file using the `-l` flag.
- Count the number of words in a file using the `-w` flag.
- Count the number of characters in a file using the `-m` flag.
- When no flags are provided, the program defaults to behaving like -l, -w, and -c.
- Handles errors gracefully for missing files and invalid arguments.
- Includes comprehensive unit tests.

## 🛠️ Installation

Ensure you have [Go installed](https://go.dev/doc/install) (version 1.23 or later).

Clone the repository:
```sh
$ git clone https://github.com/yourusername/gowc.git
$ cd gowc
```

Build the project:
```sh
$ go build -o gowc
```

## 🔧 Usage

```sh
$ ./gowc -c test.txt   # Get file size in bytes
$ ./gowc -l test.txt   # Get the number of lines
$ ./gowc -w test.txt   # Get the number of words
$ ./gowc -m test.txt   # Get the number of characters
$ ./gowc test.txt      # Get the number of lines, words, and bytes
```

### Example Output:
```sh
$ ./gowc -c test.txt
342190 test.txt

$ ./gowc -l test.txt
7145 test.txt

$ ./gowc -w test.txt
58164 test.txt

$ ./gowc -m test.txt
342190 test.txt

$ ./gowc test.txt
7145	58164	342190	test.txt
```

## 🧪 Running Tests

Run the included tests using:
```sh
$ go test ./...
```

## 📂 Project Structure
```
 gowc/
 ├── main.go         # Main program logic
 ├── main_test.go    # Unit tests
 ├── go.mod          # Go module file
```

---
Made with ❤️ using Go.

