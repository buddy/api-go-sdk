package test

import (
	"api-go-sdk/buddy"
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
		err = CheckGroup(group, name, desc, 0)
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
			t.Fatal(ErrorFormatted("GroupService.Create", err))
		}
		err = CheckGroup(groupUpdated, name, desc, groupId)
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
		err = CheckGroup(groupGet, group.Name, group.Description, group.Id)
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

func testGroupMemberAdd(client *buddy.Client, workspace *buddy.Workspace, group *buddy.Group, member *buddy.Member) func(t *testing.T) {
	return func(t *testing.T) {
		ops := buddy.GroupMemberOps{
			Id: &member.Id,
		}
		memberAdded, _, err := client.GroupService.AddGroupMember(workspace.Domain, group.Id, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("GroupService.AddGroupMember", err))
		}
		err = CheckMember(memberAdded, member.Email, member.Name, member.Admin, member.WorkspaceOwner, member.Id)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testGroupMemberGet(client *buddy.Client, workspace *buddy.Workspace, group *buddy.Group, member *buddy.Member) func(t *testing.T) {
	return func(t *testing.T) {
		memberGet, _, err := client.GroupService.GetGroupMember(workspace.Domain, group.Id, member.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("GroupService.GetGroupMember", err))
		}
		err = CheckMember(memberGet, member.Email, member.Name, member.Admin, member.WorkspaceOwner, member.Id)
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
		workspace: true,
		member:    true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var group buddy.Group
	t.Run("Create", testGroupCreate(seed.client, seed.workspace, &group))
	t.Run("Update", testGroupUpdate(seed.client, seed.workspace, &group))
	t.Run("Get", testGroupGet(seed.client, seed.workspace, &group))
	t.Run("GetList", testGroupsGet(seed.client, seed.workspace, 1))
	t.Run("MemberAdd", testGroupMemberAdd(seed.client, seed.workspace, &group, seed.member))
	t.Run("MemberGet", testGroupMemberGet(seed.client, seed.workspace, &group, seed.member))
	t.Run("MemberGetList", testGroupMembersGet(seed.client, seed.workspace, &group, 1))
	t.Run("MemberDelete", testGroupMemberDelete(seed.client, seed.workspace, &group, seed.member))
	t.Run("Delete", testGroupDelete(seed.client, seed.workspace, &group))
}
