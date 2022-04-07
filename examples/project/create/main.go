package main

import (
	"flag"
	"fmt"
	"github.com/buddy/api-go-sdk/buddy"
	"os"
)

var (
	token  string
	domain string
	name   string
)

func init() {
	flag.StringVar(&token, "token", "", "Buddy Personal API Token")
	flag.StringVar(&domain, "domain", "", "Buddy workspace domain")
	flag.StringVar(&name, "name", "", "Name of the project")
}

func main() {
	flag.Parse()
	client, err := buddy.NewDefaultClient(token)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
	ops := buddy.ProjectCreateOps{
		DisplayName: &name,
	}
	project, _, err := client.ProjectService.Create(domain, &ops)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Project display name: %s\n", project.DisplayName)
	fmt.Printf("Project name: %s\n", project.Name)
}
