package test

import (
	"api-go-sdk/buddy"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	CharSetAlpha = "abcdefghijklmnopqrstuvwxyz"
)

var (
	workspace *buddy.Workspace
	project   *buddy.Project
	group     *buddy.Group
)

func GetClient() (*buddy.Client, error) {
	return buddy.NewClient(os.Getenv("BUDDY_TOKEN"), os.Getenv("BUDDY_BASE_URL"), os.Getenv("BUDDY_INSECURE") == "true")
}

func RandStringFromCharSet(strlen int, charSet string) string {
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(result)
}

func RandInt() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int()
}

func RandString(strlen int) string {
	return RandStringFromCharSet(strlen, CharSetAlpha)
}

func RandEmail() string {
	return fmt.Sprintf("%s@0zxc.com", UniqueString())
}

func UniqueString() string {
	return fmt.Sprintf("%s%d", RandString(5), time.Now().UnixNano())
}

func ErrorFormatted(msg string, err error) error {
	return fmt.Errorf("%s: %s", msg, err.Error())
}

func CheckFieldEqual(field string, got string, want string) error {
	if got != want {
		return ErrorFieldFormatted(field, got, want)
	}
	return nil
}

func CheckFieldEqualAndSet(field string, got string, want string) error {
	if err := CheckFieldEqual(field, got, want); err != nil {
		return err
	}
	return CheckFieldSet(field, got)
}

func CheckFieldSet(field string, got string) error {
	if got == "" {
		return ErrorFieldEmpty(field)
	}
	return nil
}

func CheckBoolFieldEqual(field string, got bool, want bool) error {
	if got != want {
		return ErrorFieldFormatted(field, strconv.FormatBool(got), strconv.FormatBool(want))
	}
	return nil
}

func CheckIntFieldEqual(field string, got int, want int) error {
	if got != want {
		return ErrorFieldFormatted(field, strconv.Itoa(got), strconv.Itoa(want))
	}
	return nil
}

func CheckIntFieldEqualAndSet(field string, got int, want int) error {
	if err := CheckIntFieldEqual(field, got, want); err != nil {
		return err
	}
	return CheckIntFieldSet(field, got)
}

func CheckIntFieldSet(field string, got int) error {
	if got == 0 {
		return ErrorFieldEmpty(field)
	}
	return nil
}

func ErrorFieldFormatted(field string, got string, want string) error {
	return fmt.Errorf("got %q %q; want %q", field, got, want)
}

func ErrorFieldEmpty(field string) error {
	return fmt.Errorf("expected %q not to be empty", field)
}

func SeedInitialData(client *buddy.Client) (*buddy.Workspace, *buddy.Project, *buddy.Group, error) {
	if project == nil || workspace == nil {
		domain := UniqueString()
		w := buddy.WorkspaceOperationOptions{
			Domain: &domain,
		}
		var err error
		workspace, _, err = client.WorkspaceService.Create(&w)
		if err != nil {
			return nil, nil, nil, err
		}
		projectDisplayName := UniqueString()
		p := buddy.ProjectCreateOptions{
			DisplayName: &projectDisplayName,
		}
		project, _, err = client.ProjectService.Create(domain, &p)
		if err != nil {
			return nil, nil, nil, err
		}
		groupName := UniqueString()
		g := buddy.GroupOperationOptions{
			Name: &groupName,
		}
		group, _, err = client.GroupService.Create(domain, &g)
		if err != nil {
			return nil, nil, nil, err
		}
	}
	return workspace, project, group, nil
}
