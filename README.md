# Buddy SDK for Go

`api-go-sdk` is the Buddy SDK for the Go programming language.

The SDK requires a minimum version of `Go 1.19`.

Checkout out [our blog](https://buddy.works/blog) for information about the latest changes

## Getting started
To get started working with the SDK setup your project for Go modules, and retrieve the SDK dependencies with `go get`.
This example shows how you can use the SDK to make an API request using the SDK's client.

###### Initialize Project
```sh
$ mkdir ~/buddytest
$ cd ~/buddytest
$ go mod init buddytest
```
###### Add SDK Dependencies
```sh
$ go get github.com/buddy/api-go-sdk
```

###### Write Code
In your preferred editor add the following content to `main.go`

```go
package main

import (
  "fmt"
  "github.com/buddy/api-go-sdk/buddy"
  "os"
)

func main() {
  client, err := buddy.NewDefaultClient(os.Getenv("BUDDY_TOKEN"))
  if err != nil {
    fmt.Printf("Error: %s", err.Error())
    os.Exit(1)
  }
  profile, _, err := client.ProfileService.Get()
  if err != nil {
    fmt.Printf("Error: %s", err.Error())
    os.Exit(1)
  } else {
    fmt.Printf("My name is %s", profile.Name)
  }
}
```

## Examples
For more examples go to [examples directory](https://github.com/buddy/api-go-sdk/tree/main/examples)

## API docs
Full API docs can be found [here](https://buddy.works/docs/api)