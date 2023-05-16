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
		allowPullRequests := true
		ops := buddy.ProjectCreateOps{
			DisplayName:                     &displayName,
			Integration:                     &i,
			ExternalProjectId:               &n,
			UpdateDefaultBranchFromExternal: &updateBranch,
			AllowPullRequests:               &allowPullRequests,
		}
		project, _, err := client.ProjectService.Create(workspace.Domain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectService.Create", err))
		}
		err = CheckProject(project, displayName, displayName, false, updateBranch, allowPullRequests, false, "", buddy.ProjectAccessPrivate, false)
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
		fetchSubmodules := true
		fetchSubmodulesEnvKey := "id_workspace"
		ops := buddy.ProjectCreateOps{
			DisplayName:           &displayName,
			CustomRepoUrl:         &repoUrl,
			CustomRepoSshKeyId:    &variableId,
			FetchSubmodules:       &fetchSubmodules,
			FetchSubmodulesEnvKey: &fetchSubmodulesEnvKey,
		}
		project, _, err := client.ProjectService.Create(workspace.Domain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectService.Create", err))
		}
		err = CheckProject(project, displayName, displayName, false, true, false, fetchSubmodules, fetchSubmodulesEnvKey, buddy.ProjectAccessPrivate, false)
		if err != nil {
			t.Fatal(err)
		}
		*out = *project
	}
}

func testProjectCreate(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Project) func(t *testing.T) {
	return func(t *testing.T) {
		displayName := UniqueString()
		access := buddy.ProjectAccessPublic
		ops := buddy.ProjectCreateOps{
			DisplayName: &displayName,
			Access:      &access,
		}
		project, _, err := client.ProjectService.Create(workspace.Domain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectService.Create", err))
		}
		err = CheckProject(project, displayName, displayName, false, true, false, false, "", access, false)
		if err != nil {
			t.Fatal(err)
		}
		*out = *project
	}
}

func testProjectCreateWithoutRepository(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Project) func(t *testing.T) {
	return func(t *testing.T) {
		displayName := UniqueString()
		access := buddy.ProjectAccessPrivate
		withoutRepo := true
		ops := buddy.ProjectCreateOps{
			DisplayName:       &displayName,
			Access:            &access,
			WithoutRepository: &withoutRepo,
		}
		project, _, err := client.ProjectService.Create(workspace.Domain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectService.Create", err))
		}
		err = CheckProject(project, displayName, displayName, false, true, false, false, "", access, withoutRepo)
		if err != nil {
			t.Fatal(err)
		}
		*out = *project
	}
}

func testProjectUpdate(client *buddy.Client, workspace *buddy.Workspace, updateBranch bool, allowPullRequests bool, fetchSubmodules bool, fetchSubmodulesKey string, access string, withoutRepository bool, out *buddy.Project) func(t *testing.T) {
	return func(t *testing.T) {
		displayName := RandString(10)
		name := UniqueString()
		ops := buddy.ProjectUpdateOps{
			DisplayName:                     &displayName,
			Name:                            &name,
			UpdateDefaultBranchFromExternal: &updateBranch,
			Access:                          &access,
			AllowPullRequests:               &allowPullRequests,
			FetchSubmodules:                 &fetchSubmodules,
			FetchSubmodulesEnvKey:           &fetchSubmodulesKey,
		}
		project, _, err := client.ProjectService.Update(workspace.Domain, out.Name, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectService.Update", err))
		}
		err = CheckProject(project, name, displayName, false, updateBranch, allowPullRequests, fetchSubmodules, fetchSubmodulesKey, access, withoutRepository)
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
		err = CheckProject(project, out.Name, out.DisplayName, false, out.UpdateDefaultBranchFromExternal, out.AllowPullRequests, out.FetchSubmodules, out.FetchSubmodulesEnvKey, out.Access, out.WithoutRepository)
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
	t.Run("Update", testProjectUpdate(seed.Client, seed.Workspace, false, false, false, "", buddy.ProjectAccessPrivate, false, &project))
	t.Run("Get", testProjectGet(seed.Client, seed.Workspace, &project))
	t.Run("GetList", testProjectGetList(seed.Client, seed.Workspace, 1))
	t.Run("GetListAll", testProjectGetListAll(seed.Client, seed.Workspace, 1))
}

func TestProjectWithoutRepository(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var project buddy.Project
	t.Run("Create", testProjectCreateWithoutRepository(seed.Client, seed.Workspace, &project))
	t.Run("Update", testProjectUpdate(seed.Client, seed.Workspace, false, false, false, "", buddy.ProjectAccessPrivate, true, &project))
	t.Run("Get", testProjectGet(seed.Client, seed.Workspace, &project))
	t.Run("GetList", testProjectGetList(seed.Client, seed.Workspace, 1))
	t.Run("GetListAll", testProjectGetListAll(seed.Client, seed.Workspace, 1))
}

func TestProjectCustom(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
		project:   true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var project buddy.Project
	t.Run("Create", testProjectCustomCreate(seed.Client, seed.Workspace, &project))
	time.Sleep(20 * time.Second)
	t.Run("Update", testProjectUpdate(seed.Client, seed.Workspace, false, false, false, "", buddy.ProjectAccessPublic, false, &project))
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
	t.Run("Update", testProjectUpdate(seed.Client, seed.Workspace, true, false, false, "", buddy.ProjectAccessPrivate, false, &project))
}
