package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func testMemberCreate(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Member) func(t *testing.T) {
	return func(t *testing.T) {
		email := RandEmail()
		note := RandString(10)
		agentNote := RandString(10)
		ops := buddy.MemberCreateOps{
			Email:     &email,
			Note:      &note,
			AgentNote: &agentNote,
		}
		member, _, err := client.MemberService.Create(workspace.Domain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("MemberService.Create", err))
		}
		err = CheckMember(member, email, "", note, agentNote, false, 0, false, false, 0, "")
		if err != nil {
			t.Fatal(err)
		}
		*out = *member
	}
}

func testMemberUpdateAssignToProject(client *buddy.Client, workspace *buddy.Workspace, permission *buddy.Permission, out *buddy.Member) func(t *testing.T) {
	return func(t *testing.T) {
		assign := true
		newNote := RandString(10)
		newAgentNote := RandString(10)
		ops := buddy.MemberUpdateOps{
			AutoAssignToNewProjects:   &assign,
			AutoAssignPermissionSetId: &permission.Id,
			Note:                      &newNote,
			AgentNote:                 &newAgentNote,
		}
		memberUpdated, _, err := client.MemberService.Update(workspace.Domain, out.Id, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("MemberService.Update", err))
		}
		err = CheckMember(memberUpdated, out.Email, out.Name, newNote, newAgentNote, assign, permission.Id, out.Admin, out.WorkspaceOwner, out.Id, "")
		if err != nil {
			t.Fatal(err)
		}
		*out = *memberUpdated
	}
}

func testMemberUpdateAdmin(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Member) func(t *testing.T) {
	return func(t *testing.T) {
		admin := true
		note := ""
		agentNote := RandString(10)
		ops := buddy.MemberUpdateOps{
			Admin:     &admin,
			Note:      &note,
			AgentNote: &agentNote,
		}
		memberUpdated, _, err := client.MemberService.Update(workspace.Domain, out.Id, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("MemberService.Update", err))
		}
		err = CheckMember(memberUpdated, out.Email, out.Name, note, agentNote, out.AutoAssignToNewProjects, out.AutoAssignPermissionSetId, admin, out.WorkspaceOwner, out.Id, "")
		if err != nil {
			t.Fatal(err)
		}
		*out = *memberUpdated
	}
}

func testMemberGet(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Member) func(t *testing.T) {
	return func(t *testing.T) {
		memberGet, _, err := client.MemberService.Get(workspace.Domain, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("MemberService.Get", err))
		}
		err = CheckMember(memberGet, out.Email, out.Name, out.Note, out.AgentNote, out.AutoAssignToNewProjects, out.AutoAssignPermissionSetId, out.Admin, out.WorkspaceOwner, out.Id, "")
		if err != nil {
			t.Fatal(err)
		}
		*out = *memberGet
	}
}

func testMemberGetListAll(client *buddy.Client, workspace *buddy.Workspace, count int) func(t *testing.T) {
	return func(t *testing.T) {
		members, _, err := client.MemberService.GetListAll(workspace.Domain)
		if err != nil {
			t.Fatal(ErrorFormatted("MemberService.GetListAll", err))
		}
		err = CheckMembers(members, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testMemberGetList(client *buddy.Client, workspace *buddy.Workspace, count int) func(t *testing.T) {
	return func(t *testing.T) {
		members, _, err := client.MemberService.GetList(workspace.Domain, &buddy.PageQuery{
			Page:    1,
			PerPage: 20,
		})
		if err != nil {
			t.Fatal(ErrorFormatted("MemberService.GetList", err))
		}
		err = CheckMembers(members, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testMemberDelete(client *buddy.Client, workspace *buddy.Workspace, member *buddy.Member) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.MemberService.Delete(workspace.Domain, member.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("MemberService.Delete", err))
		}
	}
}

func TestMember(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace:  true,
		permission: true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var member buddy.Member
	t.Run("Create", testMemberCreate(seed.Client, seed.Workspace, &member))
	t.Run("UpdateAssignToProject", testMemberUpdateAssignToProject(seed.Client, seed.Workspace, seed.Permission, &member))
	t.Run("UpdateAdmin", testMemberUpdateAdmin(seed.Client, seed.Workspace, &member))
	t.Run("Get", testMemberGet(seed.Client, seed.Workspace, &member))
	t.Run("GetList", testMemberGetList(seed.Client, seed.Workspace, 3))
	t.Run("GetListAll", testMemberGetListAll(seed.Client, seed.Workspace, 3))
	t.Run("Delete", testMemberDelete(seed.Client, seed.Workspace, &member))
}
