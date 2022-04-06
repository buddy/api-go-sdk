package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func testProjectGroupCreate(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, group *buddy.Group, permission *buddy.Permission) func(t *testing.T) {
	return func(t *testing.T) {
		ops := buddy.ProjectGroupOps{
			Id: &group.Id,
			PermissionSet: &buddy.ProjectGroupOps{
				Id: &permission.Id,
			},
		}
		pg, _, err := client.ProjectGroupService.CreateProjectGroup(workspace.Domain, project.Name, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectGroupService.CreateProjectGroup", err))
		}
		err = CheckProjectGroup(pg, group, permission)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testProjectGroupUpdate(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, group *buddy.Group, permission *buddy.Permission) func(t *testing.T) {
	return func(t *testing.T) {
		ops := buddy.ProjectGroupOps{
			PermissionSet: &buddy.ProjectGroupOps{
				Id: &permission.Id,
			},
		}
		pg, _, err := client.ProjectGroupService.UpdateProjectGroup(workspace.Domain, project.Name, group.Id, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectGroupService.UpdateProjectGroup", err))
		}
		err = CheckProjectGroup(pg, group, permission)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testProjectGroupGet(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, group *buddy.Group, permission *buddy.Permission) func(t *testing.T) {
	return func(t *testing.T) {
		pg, _, err := client.ProjectGroupService.GetProjectGroup(workspace.Domain, project.Name, group.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectMemberService.GetProjectMember", err))
		}
		err = CheckProjectGroup(pg, group, permission)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testProjectGroupGetList(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, count int) func(t *testing.T) {
	return func(t *testing.T) {
		groups, _, err := client.ProjectGroupService.GetProjectGroups(workspace.Domain, project.Name)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectMemberService.GetProjectMembers", err))
		}
		err = CheckGroups(groups, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestProjectGroup(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace:  true,
		project:    true,
		permission: true,
		group:      true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	t.Run("Create", testProjectGroupCreate(seed.Client, seed.Workspace, seed.Project, seed.Group, seed.Permission))
	t.Run("Update", testProjectGroupUpdate(seed.Client, seed.Workspace, seed.Project, seed.Group, seed.Permission2))
	t.Run("Get", testProjectGroupGet(seed.Client, seed.Workspace, seed.Project, seed.Group, seed.Permission2))
	t.Run("GetList", testProjectGroupGetList(seed.Client, seed.Workspace, seed.Project, 1))
}
