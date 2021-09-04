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

```
modgoule example.com/hello

go 1.16

replace example.com/greetings => ../greetings

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


