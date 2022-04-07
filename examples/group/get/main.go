package main

import (
	"flag"
	"fmt"
	"github.com/buddy/api-go-sdk/buddy"
	"os"
)

var (
	token   string
	domain  string
	groupId int
)

func init() {
	flag.StringVar(&token, "token", "", "Buddy Personal API Token")
	flag.StringVar(&domain, "domain", "", "Buddy workspace domain")
	flag.IntVar(&groupId, "group-id", 0, "Group ID to fetch")
}

func main() {
	flag.Parse()
	client, err := buddy.NewDefaultClient(token)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
	group, _, err := client.GroupService.Get(domain, groupId)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Group name: %s", group.Name)
}
