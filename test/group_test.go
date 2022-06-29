package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func testGroupCreate(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Group) func(t *testing.T) {
	return func(t *testing.T) {
		name := RandString(10)
		desc := RandString(10)
		ops := buddy.GroupOps{
			Name:        &name,
			Description: &desc,
		}
		group, _, err := client.GroupService.Create(workspace.Domain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("GroupService.Create", err))
		}
		err = CheckGroup(group, name, desc, false, 0, 0)
		if err != nil {
			t.Fatal(err)
		}
		*out = *group
	}
}

func testGroupUpdate(client *buddy.Client, workspace *buddy.Workspace, group *buddy.Group) func(t *testing.T) {
	return func(t *testing.T) {
		name := RandString(10)
		desc := RandString(10)
		ops := buddy.GroupOps{
			Name:        &name,
			Description: &desc,
		}
		groupId := group.Id
		var err error
		groupUpdated, _, err := client.GroupService.Update(workspace.Domain, groupId, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("GroupService.Update", err))
		}
		err = CheckGroup(groupUpdated, name, desc, group.AutoAssignToNewProjects, group.AutoAssignPermissionSetId, groupId)
		if err != nil {
			t.Fatal(err)
		}
		*group = *groupUpdated
	}
}

func testGroupUpdateAssignToProjects(client *buddy.Client, workspace *buddy.Workspace, permission *buddy.Permission, group *buddy.Group) func(t *testing.T) {
	return func(t *testing.T) {
		assign := true
		ops := buddy.GroupOps{
			AutoAssignToNewProjects:   &assign,
			AutoAssignPermissionSetId: &permission.Id,
		}
		groupId := group.Id
		var err error
		groupUpdated, _, err := client.GroupService.Update(workspace.Domain, groupId, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("GroupService.Update", err))
		}
		err = CheckGroup(groupUpdated, group.Name, group.Description, assign, permission.Id, groupId)
		if err != nil {
			t.Fatal(err)
		}
		*group = *groupUpdated
	}
}

func testGroupGet(client *buddy.Client, workspace *buddy.Workspace, group *buddy.Group) func(t *testing.T) {
	return func(t *testing.T) {
		groupGet, _, err := client.GroupService.Get(workspace.Domain, group.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("GroupService.Get", err))
		}
		err = CheckGroup(groupGet, group.Name, group.Description, group.AutoAssignToNewProjects, group.AutoAssignPermissionSetId, group.Id)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testGroupsGet(client *buddy.Client, workspace *buddy.Workspace, count int) func(t *testing.T) {
	return func(t *testing.T) {
		groups, _, err := client.GroupService.GetList(workspace.Domain)
		if err != nil {
			t.Fatal(ErrorFormatted("GroupService.GetList", err))
		}
		err = CheckGroups(groups, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testGroupDelete(client *buddy.Client, workspace *buddy.Workspace, group *buddy.Group) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.GroupService.Delete(workspace.Domain, group.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("GroupService.Delete", err))
		}
	}
}

func testGroupMemberUpdate(client *buddy.Client, workspace *buddy.Workspace, group *buddy.Group, member *buddy.Member, status string) func(t *testing.T) {
	return func(t *testing.T) {
		memberUpdated, _, err := client.GroupService.UpdateGroupMember(workspace.Domain, group.Id, member.Id, status)
		if err != nil {
			t.Fatal(ErrorFormatted("GroupService.UpdateGroupMember", err))
		}
		err = CheckMember(memberUpdated, member.Email, member.Name, false, 0, member.Admin, member.WorkspaceOwner, member.Id, status)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testGroupMemberAdd(client *buddy.Client, workspace *buddy.Workspace, group *buddy.Group, member *buddy.Member, status string) func(t *testing.T) {
	return func(t *testing.T) {
		ops := buddy.GroupMemberOps{
			Id: &member.Id,
		}
		if status != "" {
			ops.Status = &status
		} else {
			status = buddy.GroupMemberStatusMember
		}
		memberAdded, _, err := client.GroupService.AddGroupMember(workspace.Domain, group.Id, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("GroupService.AddGroupMember", err))
		}
		err = CheckMember(memberAdded, member.Email, member.Name, false, 0, member.Admin, member.WorkspaceOwner, member.Id, status)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testGroupMemberGet(client *buddy.Client, workspace *buddy.Workspace, group *buddy.Group, member *buddy.Member, status string) func(t *testing.T) {
	return func(t *testing.T) {
		memberGet, _, err := client.GroupService.GetGroupMember(workspace.Domain, group.Id, member.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("GroupService.GetGroupMember", err))
		}
		err = CheckMember(memberGet, member.Email, member.Name, false, 0, member.Admin, member.WorkspaceOwner, member.Id, status)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testGroupMemberDelete(client *buddy.Client, workspace *buddy.Workspace, group *buddy.Group, member *buddy.Member) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.GroupService.DeleteGroupMember(workspace.Domain, group.Id, member.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("GroupService.GetGroupMember", err))
		}
	}
}

func testGroupMembersGet(client *buddy.Client, workspace *buddy.Workspace, group *buddy.Group, count int) func(t *testing.T) {
	return func(t *testing.T) {
		members, _, err := client.GroupService.GetGroupMembers(workspace.Domain, group.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("GroupService.GetGroupMembers", err))
		}
		err = CheckMembers(members, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestGroup(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace:  true,
		member:     true,
		permission: true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var group buddy.Group
	var member buddy.Member
	t.Run("CreateManager", testMemberCreate(seed.Client, seed.Workspace, &member))
	t.Run("Create", testGroupCreate(seed.Client, seed.Workspace, &group))
	t.Run("Update", testGroupUpdate(seed.Client, seed.Workspace, &group))
	t.Run("UpdateProjectAssign", testGroupUpdateAssignToProjects(seed.Client, seed.Workspace, seed.Permission, &group))
	t.Run("Get", testGroupGet(seed.Client, seed.Workspace, &group))
	t.Run("GetList", testGroupsGet(seed.Client, seed.Workspace, 1))
	t.Run("MemberAdd", testGroupMemberAdd(seed.Client, seed.Workspace, &group, seed.Member, ""))
	t.Run("ManagerAdd", testGroupMemberAdd(seed.Client, seed.Workspace, &group, &member, buddy.GroupMemberStatusManager))
	t.Run("MemberGet", testGroupMemberGet(seed.Client, seed.Workspace, &group, seed.Member, buddy.GroupMemberStatusMember))
	t.Run("ManagerGet", testGroupMemberGet(seed.Client, seed.Workspace, &group, &member, buddy.GroupMemberStatusManager))
	t.Run("MemberUpdateToManager", testGroupMemberUpdate(seed.Client, seed.Workspace, &group, seed.Member, buddy.GroupMemberStatusManager))
	t.Run("ManagerUpdateToMember", testGroupMemberUpdate(seed.Client, seed.Workspace, &group, &member, buddy.GroupMemberStatusMember))
	t.Run("NewManagerGet", testGroupMemberGet(seed.Client, seed.Workspace, &group, seed.Member, buddy.GroupMemberStatusManager))
	t.Run("NewMemberGet", testGroupMemberGet(seed.Client, seed.Workspace, &group, &member, buddy.GroupMemberStatusMember))
	t.Run("MemberGetList", testGroupMembersGet(seed.Client, seed.Workspace, &group, 2))
	t.Run("MemberDelete", testGroupMemberDelete(seed.Client, seed.Workspace, &group, seed.Member))
	t.Run("MemberDelete", testGroupMemberDelete(seed.Client, seed.Workspace, &group, &member))
	t.Run("Delete", testGroupDelete(seed.Client, seed.Workspace, &group))
}
