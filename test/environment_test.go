package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func testEnvironmentCreate(client *buddy.Client, workspace *buddy.Workspace, ops *buddy.EnvironmentOps, out *buddy.Environment) func(t *testing.T) {
	return func(t *testing.T) {
		environment, _, err := client.EnvironmentService.Create(workspace.Domain, ops)
		if err != nil {
			t.Fatal(ErrorFormatted("EnvironmentService.Create", err))
		}
		err = CheckEnvironment(environment, out, ops)
		if err != nil {
			t.Fatal(err)
		}
		*out = *environment
	}
}

func testEnvironmentGet(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Environment) func(t *testing.T) {
	return func(t *testing.T) {
		environment, _, err := client.EnvironmentService.Get(workspace.Domain, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("EnvironmentService.Get", err))
		}
		err = CheckEnvironment(environment, out, nil)
		if err != nil {
			t.Fatal(err)
		}
		*out = *environment
	}
}

func testEnvironmentDelete(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Environment) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.EnvironmentService.Delete(workspace.Domain, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("EnvironmentService.Delete", err))
		}
	}
}

func testEnvironmentGetList(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, count int) func(t *testing.T) {
	return func(t *testing.T) {
		projectName := ""
		if project != nil {
			projectName = project.Name
		}
		environments, _, err := client.EnvironmentService.GetList(workspace.Domain, projectName)
		if err != nil {
			t.Fatal(ErrorFormatted("EnvironmentService.GetList", err))
		}
		err = CheckEnvironments(environments, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testEnvironmentUpdate(client *buddy.Client, workspace *buddy.Workspace, ops *buddy.EnvironmentOps, out *buddy.Environment) func(t *testing.T) {
	return func(t *testing.T) {
		environment, _, err := client.EnvironmentService.Update(workspace.Domain, out.Id, ops)
		if err != nil {
			t.Fatal(ErrorFormatted("EnvironmentService.Update", err))
		}
		err = CheckEnvironment(environment, out, ops)
		if err != nil {
			t.Fatal(err)
		}
		*out = *environment
	}
}

func TestEnvironmentWorkspace(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	name := RandString(10)
	newName := RandString(10)
	identifier := UniqueString()
	newIdentifier := UniqueString()
	publicUrl := "https://google.com"
	newPublicUrl := "https://bing.com"
	tags := []string{"aaa"}
	newTags := []string{} // must be like that to proper json encode
	allPipAllowed := buddy.EnvironmentAccessLevelUseOnly
	newAllPipAllowed := buddy.EnvironmentAccessLevelDenied
	allEnvsAllowed := buddy.EnvironmentAccessLevelUseOnly
	newAllEnvsAllowed := buddy.EnvironmentAccessLevelDenied
	baseOnly := true
	newBaseOnly := false
	scope := buddy.EnvironmentScopeWorkspace
	ops := buddy.EnvironmentOps{
		Name:                    &name,
		Identifier:              &identifier,
		PublicUrl:               &publicUrl,
		Tags:                    &tags,
		PipelinesAccessLevel:    &allPipAllowed,
		EnvironmentsAccessLevel: &allEnvsAllowed,
		BaseOnly:                &baseOnly,
		Scope:                   &scope,
	}
	updOps := buddy.EnvironmentOps{
		Name:                    &newName,
		Identifier:              &newIdentifier,
		PublicUrl:               &newPublicUrl,
		Tags:                    &newTags,
		BaseOnly:                &newBaseOnly,
		PipelinesAccessLevel:    &newAllPipAllowed,
		EnvironmentsAccessLevel: &newAllEnvsAllowed,
		Scope:                   &scope,
	}
	var environment buddy.Environment
	t.Run("Create", testEnvironmentCreate(seed.Client, seed.Workspace, &ops, &environment))
	t.Run("Update", testEnvironmentUpdate(seed.Client, seed.Workspace, &updOps, &environment))
	t.Run("Get", testEnvironmentGet(seed.Client, seed.Workspace, &environment))
	t.Run("GetList", testEnvironmentGetList(seed.Client, seed.Workspace, seed.Project, 1))
	t.Run("Delete", testEnvironmentDelete(seed.Client, seed.Workspace, &environment))
}

func TestEnvironment(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace:  true,
		project:    true,
		group:      true,
		member:     true,
		permission: true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	pm := buddy.ProjectMemberOps{
		Id:            &seed.Member.Id,
		PermissionSet: &buddy.ProjectMemberOps{Id: &seed.Permission.Id},
	}
	_, _, err = seed.Client.ProjectMemberService.CreateProjectMember(seed.Workspace.Domain, seed.Project.Name, &pm)
	if err != nil {
		t.Fatal(ErrorFormatted("ProjectMemberService.CreateProjectMember", err))
	}
	pg := buddy.ProjectGroupOps{
		Id:            &seed.Group.Id,
		PermissionSet: &buddy.ProjectGroupOps{Id: &seed.Permission.Id},
	}
	_, _, err = seed.Client.ProjectGroupService.CreateProjectGroup(seed.Workspace.Domain, seed.Project.Name, &pg)
	if err != nil {
		t.Fatal(ErrorFormatted("ProjectGroupService.CreateProjectGroup", err))
	}
	name := RandString(10)
	newName := RandString(10)
	identifier := UniqueString()
	icon := RandString(10)
	newIcon := RandString(10)
	newIdentifier := UniqueString()
	publicUrl := "https://google.com"
	newPublicUrl := "https://bing.com"
	tags := []string{"aaa", "bbb"}
	newTags := []string{"ccc"}
	user := buddy.EnvironmentResourcePermissions{
		Id:          seed.Member.Id,
		AccessLevel: buddy.EnvironmentPermissionAccessLevelDenied,
	}
	group := buddy.EnvironmentResourcePermissions{
		Id:          seed.Group.Id,
		AccessLevel: buddy.EnvironmentPermissionAccessLevelDefault,
	}
	perms := buddy.EnvironmentPermissions{
		Others: buddy.EnvironmentPermissionAccessLevelUseOnly,
		Users:  []*buddy.EnvironmentResourcePermissions{&user},
		Groups: []*buddy.EnvironmentResourcePermissions{&group},
	}
	newPerms := buddy.EnvironmentPermissions{
		Others: buddy.EnvironmentPermissionAccessLevelManage,
	}
	allPipAllowed := buddy.EnvironmentAccessLevelDenied
	newAllPipAllowed := buddy.EnvironmentAccessLevelUseOnly
	allEnvsAllowed := buddy.EnvironmentAccessLevelUseOnly
	newAllEnvsAllowed := buddy.EnvironmentAccessLevelDenied
	scope := buddy.EnvironmentScopeProject
	ops := buddy.EnvironmentOps{
		Name:                    &name,
		Identifier:              &identifier,
		PublicUrl:               &publicUrl,
		Icon:                    &icon,
		Tags:                    &tags,
		Permissions:             &perms,
		PipelinesAccessLevel:    &allPipAllowed,
		EnvironmentsAccessLevel: &allEnvsAllowed,
		Project: &buddy.ProjectSimple{
			Name: seed.Project.Name,
		},
		Scope: &scope,
	}
	updOps := buddy.EnvironmentOps{
		Name:                    &newName,
		Identifier:              &newIdentifier,
		PublicUrl:               &newPublicUrl,
		Icon:                    &newIcon,
		Tags:                    &newTags,
		Permissions:             &newPerms,
		PipelinesAccessLevel:    &newAllPipAllowed,
		EnvironmentsAccessLevel: &newAllEnvsAllowed,
		Scope:                   &scope,
	}
	var environment buddy.Environment
	t.Run("Create", testEnvironmentCreate(seed.Client, seed.Workspace, &ops, &environment))
	t.Run("Update", testEnvironmentUpdate(seed.Client, seed.Workspace, &updOps, &environment))
	t.Run("Get", testEnvironmentGet(seed.Client, seed.Workspace, &environment))
	t.Run("GetList", testEnvironmentGetList(seed.Client, seed.Workspace, seed.Project, 1))
	t.Run("Delete", testEnvironmentDelete(seed.Client, seed.Workspace, &environment))
}
