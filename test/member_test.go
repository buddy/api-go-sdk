package test

import (
	"api-go-sdk/buddy"
	"testing"
)

func TestMember(t *testing.T) {
	client, err := GetClient()
	if err != nil {
		t.Fatal(ErrorFormatted("GetClient", err))
	}
	workspace, _, _, err := SeedInitialData(client)
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	// CREATE MEMBER
	email := RandEmail()
	c := buddy.MemberOperationOptions{
		Email: &email,
	}
	member, _, err := client.MemberService.Create(workspace.Domain, &c)
	if err != nil {
		t.Fatal(ErrorFormatted("MemberService.Create", err))
	}
	err = CheckFieldSet("Member.HtmlUrl", member.HtmlUrl)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckIntFieldSet("Member.Id", member.Id)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqual("Member.Name", member.Name, "")
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Member.Email", member.Email, email)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldSet("Member.AvatarUrl", member.AvatarUrl)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckBoolFieldEqual("Member.Admin", member.Admin, false)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckBoolFieldEqual("Member.WorkspaceOwner", member.WorkspaceOwner, false)
	if err != nil {
		t.Fatal(err)
	}
	memberId := member.Id
	admin := true
	// UPDATE ADMIN
	u := buddy.MemberAdminOperationOptions{
		Admin: &admin,
	}
	member, _, err = client.MemberService.UpdateAdmin(workspace.Domain, memberId, &u)
	if err != nil {
		t.Fatal(ErrorFormatted("MemberService.UpdateAdmin", err))
	}
	err = CheckFieldSet("Member.HtmlUrl", member.HtmlUrl)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckIntFieldEqualAndSet("Member.Id", member.Id, memberId)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqual("Member.Name", member.Name, "")
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Member.Email", member.Email, email)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldSet("Member.AvatarUrl", member.AvatarUrl)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckBoolFieldEqual("Member.Admin", member.Admin, true)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckBoolFieldEqual("Member.WorkspaceOwner", member.WorkspaceOwner, false)
	if err != nil {
		t.Fatal(err)
	}
	// DELETE
	_, err = client.MemberService.Delete(workspace.Domain, memberId)
	if err != nil {
		t.Fatal(ErrorFormatted("MemberService.Delete", err))
	}
}
