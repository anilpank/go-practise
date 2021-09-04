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

### Call code from another module
Enable dependency tracking for your code
```go
go mod init example.com/hello
```

go: creating new go.mod: module example.com/hello
Your hello.go code

```go
modgoule example.com/hello

go 1.16

replace example.com/greetings => ../greetings
```

```go
import (
    "fmt"

    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
```

Edit the example.com/hello module to use your local example.com/greetings module.
go mod edit -replace example.com/greetings=../greetings

After you run the command, the go.mod file in the hello directory should include a replace directive:
```go
module example.com/hello

go 1.16

replace example.com/greetings => ../greetings
```
From the command prompt in the hello directory, run the go mod tidy command to synchronize the example.com/hello module's dependencies, adding those required by the code, but not yet tracked in the module.
go mod tidy

After the command completes, the example.com/hello module's go.mod file should look like this:
```go
module example.com/hello

go 1.16

replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000
```

### Return and handle an error

Your greetings.go code

```go
package greetings

import (
    "errors"
    "fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }

    // If a name was received, return a value that embeds the name
    // in a greeting message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
```

Your hello.go code
```go
package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Request a greeting message.
    message, err := greetings.Hello("")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}
```

