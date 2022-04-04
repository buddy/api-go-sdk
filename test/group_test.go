package test

import (
	"api-go-sdk/buddy"
	"testing"
)

func TestGroup(t *testing.T) {
	client, err := GetClient()
	if err != nil {
		t.Fatal(ErrorFormatted("GetClient", err))
	}
	workspace, _, _, err := SeedInitialData(client)
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	// CREATE GROUP
	name := RandString(10)
	desc := RandString(10)
	c := buddy.GroupOperationOptions{
		Name:        &name,
		Description: &desc,
	}
	group, _, err := client.GroupService.Create(workspace.Domain, &c)
	if err != nil {
		t.Fatal(ErrorFormatted("GroupService.Create", err))
	}
	err = CheckFieldSet("Group.Url", group.Url)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldSet("Group.HtmlUrl", group.HtmlUrl)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckIntFieldSet("Group.Id", group.Id)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Group.Name", group.Name, name)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Group.Description", group.Description, desc)
	if err != nil {
		t.Fatal(err)
	}
	// EDIT GROUP
	newName := RandString(10)
	newDesc := ""
	u := buddy.GroupOperationOptions{
		Name:        &newName,
		Description: &newDesc,
	}
	group, _, err = client.GroupService.Update(workspace.Domain, group.Id, &u)
	if err != nil {
		t.Fatal(ErrorFormatted("GroupService.Update", err))
	}
	err = CheckFieldSet("Group.Url", group.Url)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldSet("Group.HtmlUrl", group.HtmlUrl)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckIntFieldSet("Group.Id", group.Id)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Group.Name", group.Name, newName)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqual("Group.Description", group.Description, newDesc)
	if err != nil {
		t.Fatal(err)
	}
	// DELETE GROUP
	_, err = client.GroupService.Delete(workspace.Domain, group.Id)
	if err != nil {
		t.Fatal(ErrorFormatted("GroupService.Delete", err))
	}
}
