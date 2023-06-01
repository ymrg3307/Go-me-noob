# Go-me-noob

Just exploring Go!!!

1. Installation
    - Installed Go
    - created 01hello and go into it's path and do `go mod init hello` in terminal
    - Installed Go extension in vs code
    - Sample code in main.go and run `go run main.go`
2. Reading docs and GOPATH
    - `go env GOPATH` -> gives the **GOPATH**, location where source code lives.
    - `go help <command name>` -> gives doc on the mentioned command.
3. Lexer
    - Lexer -> reads input code and converts it into a token stream, automatically inserts ';' while converting.
    - Observe that ';' is automatically removed in vs code when you have installed go extension.
4. Types, vars and constants
    - Everything is a type in Go
    - var -> variable, const -> constant(cannot modify value)
    - Better to provide type, if not type will be calculated based on value provided, no var style where keyword need not be used
    - no var style doesn't work outside of function body.
    - variables which starts with **capital letter** are public variables and can be accessed outside that module.
    - if value is not provided, string type has no value, bool has false, int has 0 as garbage values.
    - check 02variables code for some examples
5. Packages and Comma ok syntax
    - [go packages](https://pkg.go.dev/) can be found here.
    - bufio -> provides a number of functions and types that can be used to implement buffered I/O.
    - os -> provides a platform-independent interface to operating system functionality.
    - In the context of the Go programming language, buffering refers to the practice of storing data in a temporary location before it is read or written to a file or other input/output (I/O) device. This can help to improve performance by reducing the number of times that the I/O device needs to be accessed.
    - os.Stdin is a variable. This variable is a pointer to a *os.File object, which we can use to read data from the standard input stream.
    - comma ok idiom, just a fancy name in my opinion -> similar to variables mapping `result,ok = func()` type of ok depends on of the return type from the right hand sider.
6. Type Conversions
    - strconv -> library that helps in conversions to and from string representations of basic data types.
7. Time handling
    - check available [default time format strings and how to make a custom format](https://pkg.go.dev/time#Layout)
    - It's important to remember that time formatting is very different in go, **always check documentation before handling time**.
    - check **05mytime** code for some examples
8. Build for different OS
   - `go env` -> gives list of environment variables.
   - **GOOS** -> contains the the current environment OS.
   - `go build` -> builds a executable file.
   - GOOS="[OS]" go build -> build to a different OS from the current OS.
9. Memory Management
    - Memory management is automatic.
    - new()
      - memory is allocated but no initialization
      - we will get a memory address
      - zeroed storage
    - make()
      - memory is allocated and will be initialized
      - we will get a memory address
      - non zeroed storage
    - Garbage collection happens is automatic when something is out of scope or value is nil
    - [runtime package](https://pkg.go.dev/runtime) -> contains operations that interact with Go's runtime system
10. Pointers
    - A pointer holds the memory address of a value
    - default value is nil for pointers
    - **&** operator generates a pointer
    - **\*** denotes the pointer's underlying value
    - using pointers we can referance and modify the values directly
    - check pointers module for basic example on how to declare and use pointers
11. Arrays
    - Not much used in Golang
    - We should define the size of array explicitely
    - len() will give the complete size of array not the actual number of values present inside.
    - all the array elements are initialized with the zero value of the corresponding array type like '' for string, 0 for int etc.,
    - check array module for some examples.
12. Slices
    - FOUND THESE TO BE TRICKY
    - Comeback and read [slices intro](https://go.dev/blog/slices-intro) and [mechanics of append](https://go.dev/blog/slices) to gain indepth understanding.
    - What I observed is imagine an array in background, slices are in foreground. When you change something array gets update but slices holds the referance to the specific part we are slicing. Observe the last code block of fruitList -> WTH!!! replaced 'Peach' because the previous **fruitList** referance is [:2] so new append adds content after index 1 (as max index is exclusive)
    - **sort** package provides utilities for sorting slices and more.
    - Look at slices module to know how to delete an element at a specific indexes.
13. Maps
    - stores values as key value pairs
    - map[KeyType]ValueType{} or make(map[KeyType]ValueType)-> declaring a map
    - delete(\<mapname>,\<keyvalue>) -> delete key from map
