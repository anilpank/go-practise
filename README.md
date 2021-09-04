# go-practise
Set of Go Code that I am practising

### Enable Dependency Tracking for your code
go mod init example.com/hello

### Your First Go Code
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Running your code
go run .

### Getting Help on Go
go help

### Call code in an external package 
```go
import "rsc.io/quote"
```

### Creating a Module
Start your module using go mod init command
```go mod init example.com/greetings```

### Declare a message variable to hold your greeting
```go
var message string
message = fmt.Sprintf("Hi, %v. Welcome!", name)
```

