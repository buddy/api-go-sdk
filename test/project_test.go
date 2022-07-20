package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"os"
	"testing"
	"time"
)

func testProjectExternalCreate(client *buddy.Client, workspace *buddy.Workspace, integration *buddy.Integration, out *buddy.Project) func(t *testing.T) {
	return func(t *testing.T) {
		displayName := UniqueString()
		i := buddy.ProjectIntegration{
			HashId: integration.HashId,
		}
		n := os.Getenv("BUDDY_GH_PROJECT")
		updateBranch := false
		ops := buddy.ProjectCreateOps{
			DisplayName:                     &displayName,
			Integration:                     &i,
			ExternalProjectId:               &n,
			UpdateDefaultBranchFromExternal: &updateBranch,
		}
		project, _, err := client.ProjectService.Create(workspace.Domain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectService.Create", err))
		}
		err = CheckProject(project, displayName, displayName, false, updateBranch)
		if err != nil {
			t.Fatal(err)
		}
		*out = *project
	}
}

func testProjectCustomCreate(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Project) func(t *testing.T) {
	return func(t *testing.T) {
		variables, _, err := client.VariableService.GetList(workspace.Domain, nil)
		if err != nil {
			t.Fatal(ErrorFormatted("VariableService.GetList", err))
		}
		var variableId int
		for _, v := range variables.Variables {
			if v.Type == buddy.VariableTypeSshKey {
				variableId = v.Id
				break
			}
		}
		repoUrl := "git@github.com:octocat/Hello-World.git"
		displayName := UniqueString()
		ops := buddy.ProjectCreateOps{
			DisplayName:        &displayName,
			CustomRepoUrl:      &repoUrl,
			CustomRepoSshKeyId: &variableId,
		}
		project, _, err := client.ProjectService.Create(workspace.Domain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectService.Create", err))
		}
		err = CheckProject(project, displayName, displayName, false, false)
		if err != nil {
			t.Fatal(err)
		}
		*out = *project
	}
}

func testProjectCreate(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Project) func(t *testing.T) {
	return func(t *testing.T) {
		displayName := UniqueString()
		ops := buddy.ProjectCreateOps{
			DisplayName: &displayName,
		}
		project, _, err := client.ProjectService.Create(workspace.Domain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectService.Create", err))
		}
		err = CheckProject(project, displayName, displayName, false, false)
		if err != nil {
			t.Fatal(err)
		}
		*out = *project
	}
}

func testProjectUpdate(client *buddy.Client, workspace *buddy.Workspace, updateBranch bool, out *buddy.Project) func(t *testing.T) {
	return func(t *testing.T) {
		displayName := RandString(10)
		name := UniqueString()
		ops := buddy.ProjectUpdateOps{
			DisplayName:                     &displayName,
			Name:                            &name,
			UpdateDefaultBranchFromExternal: &updateBranch,
		}
		project, _, err := client.ProjectService.Update(workspace.Domain, out.Name, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectService.Update", err))
		}
		err = CheckProject(project, name, displayName, false, updateBranch)
		if err != nil {
			t.Fatal(err)
		}
		*out = *project
	}
}

func testProjectGet(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Project) func(t *testing.T) {
	return func(t *testing.T) {
		project, _, err := client.ProjectService.Get(workspace.Domain, out.Name)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectService.Get", err))
		}
		err = CheckProject(project, out.Name, out.DisplayName, false, out.UpdateDefaultBranchFromExternal)
		if err != nil {
			t.Fatal(err)
		}
		*out = *project
	}
}

func testProjectGetList(client *buddy.Client, workspace *buddy.Workspace, count int) func(t *testing.T) {
	return func(t *testing.T) {
		pageQuery := buddy.PageQuery{
			Page:    1,
			PerPage: 30,
		}
		query := buddy.ProjectListQuery{
			PageQuery: pageQuery,
			Status:    buddy.ProjectStatusActive,
		}
		projects, _, err := client.ProjectService.GetList(workspace.Domain, &query)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectService.GetList", err))
		}
		err = CheckProjects(projects, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testProjectGetListAll(client *buddy.Client, workspace *buddy.Workspace, count int) func(t *testing.T) {
	return func(t *testing.T) {
		projects, _, err := client.ProjectService.GetListAll(workspace.Domain, nil)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectService.GetListAll", err))
		}
		err = CheckProjects(projects, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestProjectBuddy(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var project buddy.Project
	t.Run("Create", testProjectCreate(seed.Client, seed.Workspace, &project))
	t.Run("Update", testProjectUpdate(seed.Client, seed.Workspace, false, &project))
	t.Run("Get", testProjectGet(seed.Client, seed.Workspace, &project))
	t.Run("GetList", testProjectGetList(seed.Client, seed.Workspace, 1))
	t.Run("GetListAll", testProjectGetListAll(seed.Client, seed.Workspace, 1))
}

func TestProjectCustom(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var project buddy.Project
	t.Run("Create", testProjectCustomCreate(seed.Client, seed.Workspace, &project))
	time.Sleep(20 * time.Second)
	t.Run("Update", testProjectUpdate(seed.Client, seed.Workspace, false, &project))
	t.Run("Get", testProjectGet(seed.Client, seed.Workspace, &project))
	t.Run("GetList", testProjectGetList(seed.Client, seed.Workspace, 2))
	t.Run("GetListAll", testProjectGetListAll(seed.Client, seed.Workspace, 2))
}

func TestProjectExternal(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace:      true,
		gitIntegration: true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var project buddy.Project
	t.Run("Create", testProjectExternalCreate(seed.Client, seed.Workspace, seed.GitIntegration, &project))
	time.Sleep(20 * time.Second)
	t.Run("Update", testProjectUpdate(seed.Client, seed.Workspace, true, &project))
}
