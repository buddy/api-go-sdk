package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func testEnvironmentCreate(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, ops *buddy.EnvironmentOps, out *buddy.Environment) func(t *testing.T) {
	return func(t *testing.T) {
		environment, _, err := client.EnvironmentService.Create(workspace.Domain, project.Name, ops)
		if err != nil {
			t.Fatal(ErrorFormatted("EnvironmentService.Create", err))
		}
		err = CheckEnvironment(project, environment, out, ops)
		if err != nil {
			t.Fatal(err)
		}
		*out = *environment
	}
}

func testEnvironmentGet(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, out *buddy.Environment) func(t *testing.T) {
	return func(t *testing.T) {
		environment, _, err := client.EnvironmentService.Get(workspace.Domain, project.Name, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("EnvironmentService.Get", err))
		}
		err = CheckEnvironment(project, environment, out, nil)
		if err != nil {
			t.Fatal(err)
		}
		*out = *environment
	}
}

func testEnvironmentDelete(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, out *buddy.Environment) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.EnvironmentService.Delete(workspace.Domain, project.Name, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("EnvironmentService.Delete", err))
		}
	}
}

func testEnvironmentGetList(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, count int) func(t *testing.T) {
	return func(t *testing.T) {
		environments, _, err := client.EnvironmentService.GetList(workspace.Domain, project.Name)
		if err != nil {
			t.Fatal(ErrorFormatted("EnvironmentService.GetList", err))
		}
		err = CheckEnvironments(environments, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testEnvironmentUpdate(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, ops *buddy.EnvironmentOps, out *buddy.Environment) func(t *testing.T) {
	return func(t *testing.T) {
		environment, _, err := client.EnvironmentService.Update(workspace.Domain, project.Name, out.Id, ops)
		if err != nil {
			t.Fatal(ErrorFormatted("EnvironmentService.Update", err))
		}
		err = CheckEnvironment(project, environment, out, ops)
		if err != nil {
			t.Fatal(err)
		}
		*out = *environment
	}
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
	newIdentifier := UniqueString()
	publicUrl := "https://google.com"
	newPublicUrl := "https://bing.com"
	tags := []string{"aaa", "bbb"}
	newTags := []string{"ccc"}
	v := buddy.Variable{
		Key:       RandString(10),
		Value:     RandString(10),
		Type:      buddy.VariableTypeVar,
		FilePlace: buddy.VariableSshKeyFilePlaceNone,
	}
	vars := []*buddy.Variable{&v}
	// musi byc istniejaca tablica bo inaczej jest null
	newVars := []*buddy.Variable{}
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
	allPipAllowed := true
	newAllPipAllowed := false
	ops := buddy.EnvironmentOps{
		Name:                &name,
		Identifier:          &identifier,
		PublicUrl:           &publicUrl,
		Tags:                &tags,
		Variables:           &vars,
		Permissions:         &perms,
		AllPipelinesAllowed: &allPipAllowed,
	}
	updOps := buddy.EnvironmentOps{
		Name:                &newName,
		Identifier:          &newIdentifier,
		PublicUrl:           &newPublicUrl,
		Tags:                &newTags,
		Variables:           &newVars,
		Permissions:         &newPerms,
		AllPipelinesAllowed: &newAllPipAllowed,
	}
	var environment buddy.Environment
	t.Run("Create", testEnvironmentCreate(seed.Client, seed.Workspace, seed.Project, &ops, &environment))
	t.Run("Update", testEnvironmentUpdate(seed.Client, seed.Workspace, seed.Project, &updOps, &environment))
	t.Run("Get", testEnvironmentGet(seed.Client, seed.Workspace, seed.Project, &environment))
	t.Run("GetList", testEnvironmentGetList(seed.Client, seed.Workspace, seed.Project, 1))
	t.Run("Delete", testEnvironmentDelete(seed.Client, seed.Workspace, seed.Project, &environment))
}
