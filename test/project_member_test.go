package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func testProjectMemberCreate(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, member *buddy.Member, permission *buddy.Permission) func(t *testing.T) {
	return func(t *testing.T) {
		ops := buddy.ProjectMemberOps{
			Id: &member.Id,
			PermissionSet: &buddy.ProjectMemberOps{
				Id: &permission.Id,
			},
		}
		pm, _, err := client.ProjectMemberService.CreateProjectMember(workspace.Domain, project.Name, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectMemberService.CreateProjectMember", err))
		}
		err = CheckProjectMember(pm, member, permission)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testProjectMemberUpdate(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, member *buddy.Member, permission *buddy.Permission) func(t *testing.T) {
	return func(t *testing.T) {
		ops := buddy.ProjectMemberOps{
			PermissionSet: &buddy.ProjectMemberOps{
				Id: &permission.Id,
			},
		}
		pm, _, err := client.ProjectMemberService.UpdateProjectMember(workspace.Domain, project.Name, member.Id, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectMemberService.CreateProjectMember", err))
		}
		err = CheckProjectMember(pm, member, permission)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testProjectMemberGet(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, member *buddy.Member, permission *buddy.Permission) func(t *testing.T) {
	return func(t *testing.T) {
		pm, _, err := client.ProjectMemberService.GetProjectMember(workspace.Domain, project.Name, member.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectMemberService.GetProjectMember", err))
		}
		err = CheckProjectMember(pm, member, permission)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testProjectMemberGetList(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, count int) func(t *testing.T) {
	return func(t *testing.T) {
		query := buddy.PageQuery{
			Page:    1,
			PerPage: 30,
		}
		members, _, err := client.ProjectMemberService.GetProjectMembers(workspace.Domain, project.Name, &query)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectMemberService.GetProjectMembers", err))
		}
		err = CheckMembers(members, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testProjectMemberGetListAll(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, count int) func(t *testing.T) {
	return func(t *testing.T) {
		members, _, err := client.ProjectMemberService.GetProjectMembersAll(workspace.Domain, project.Name)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectMemberService.GetProjectMembers", err))
		}
		err = CheckMembers(members, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestProjectMember(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace:  true,
		member:     true,
		project:    true,
		permission: true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	t.Run("Create", testProjectMemberCreate(seed.Client, seed.Workspace, seed.Project, seed.Member, seed.Permission))
	t.Run("Update", testProjectMemberUpdate(seed.Client, seed.Workspace, seed.Project, seed.Member, seed.Permission2))
	t.Run("Get", testProjectMemberGet(seed.Client, seed.Workspace, seed.Project, seed.Member, seed.Permission2))
	t.Run("GetList", testProjectMemberGetList(seed.Client, seed.Workspace, seed.Project, 2))
	t.Run("GetListAll", testProjectMemberGetListAll(seed.Client, seed.Workspace, seed.Project, 2))
}
