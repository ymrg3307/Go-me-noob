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
