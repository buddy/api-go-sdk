package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func testVariableCreate(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, typ string, enc bool, set bool, out *buddy.Variable) func(t *testing.T) {
	return func(t *testing.T) {
		err, _, privateKey := GenerateRsaKeyPair()
		if err != nil {
			t.Fatal(ErrorFormatted("create rsa key pair", err))
		}
		key := RandString(10)
		val := RandString(10)
		desc := RandString(10)
		fileChmod := ""
		filePath := ""
		filePlace := ""
		ops := buddy.VariableOps{
			Key:         &key,
			Value:       &val,
			Type:        &typ,
			Description: &desc,
			Settable:    &set,
			Encrypted:   &enc,
		}
		if typ == buddy.VariableTypeSshKey {
			val = privateKey
			fileChmod = "666"
			filePath = "~/.ssh/" + RandString(6)
			filePlace = buddy.VariableSshKeyFilePlaceContainer
			ops.FileChmod = &fileChmod
			ops.FilePlace = &filePlace
			ops.FilePath = &filePath
			ops.Value = &val
		}
		if project != nil {
			ops.Project = &buddy.VariableProject{
				Name: project.Name,
			}
		}
		variable, _, err := client.VariableService.Create(workspace.Domain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("VariableService.Create", err))
		}
		err = CheckVariable(variable, key, val, typ, desc, set, enc, filePath, fileChmod, filePlace, 0, project)
		if err != nil {
			t.Fatal(err)
		}
		*out = *variable
	}
}

func testVariableUpdate(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, out *buddy.Variable) func(t *testing.T) {
	return func(t *testing.T) {
		err, _, privateKey := GenerateRsaKeyPair()
		if err != nil {
			t.Fatal(ErrorFormatted("create rsa key pair", err))
		}
		val := RandString(10)
		desc := ""
		set := false
		enc := true
		filePath := ""
		filePlace := ""
		fileChmod := ""
		ops := buddy.VariableOps{
			Value:       &val,
			Description: &desc,
			Settable:    &set,
			Encrypted:   &enc,
			Type:        &out.Type,
		}
		if out.Type == buddy.VariableTypeSshKey {
			val = privateKey
			filePath = "/bec/" + RandString(5)
			filePlace = buddy.VariableSshKeyFilePlaceNone
			fileChmod = "600"
			ops.FilePath = &filePath
			ops.FilePlace = &filePlace
			ops.FileChmod = &fileChmod
			ops.Value = &val
		}
		variable, _, err := client.VariableService.Update(workspace.Domain, out.Id, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("VariableService.Patch", err))
		}
		err = CheckVariable(variable, out.Key, val, out.Type, desc, set, enc, filePath, fileChmod, filePlace, out.Id, project)
		if err != nil {
			t.Fatal(err)
		}
		*out = *variable
	}
}

func testVariableGet(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, out *buddy.Variable) func(t *testing.T) {
	return func(t *testing.T) {
		variable, _, err := client.VariableService.Get(workspace.Domain, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("VariableService.Get", err))
		}
		err = CheckVariable(variable, out.Key, out.Value, out.Type, out.Description, out.Settable, out.Encrypted, out.FilePath, out.FileChmod, out.FilePlace, out.Id, project)
		if err != nil {
			t.Fatal(err)
		}
		*out = *variable
	}
}

func testVariableGetList(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, count int) func(t *testing.T) {
	return func(t *testing.T) {
		query := buddy.VariableGetListQuery{}
		if project != nil {
			query.ProjectName = project.Name
		}
		variables, _, err := client.VariableService.GetList(workspace.Domain, &query)
		if err != nil {
			t.Fatal(ErrorFormatted("VariableService.GetList", err))
		}
		err = CheckVariables(variables, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testVariableDelete(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Variable) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.VariableService.Delete(workspace.Domain, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("VariableService.Delete", err))
		}
	}
}

func TestVariable(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
		project:   true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var variable buddy.Variable
	t.Run("Create", testVariableCreate(seed.Client, seed.Workspace, nil, buddy.VariableTypeVar, false, true, &variable))
	t.Run("Update", testVariableUpdate(seed.Client, seed.Workspace, nil, &variable))
	t.Run("Get", testVariableGet(seed.Client, seed.Workspace, nil, &variable))
	t.Run("GetList", testVariableGetList(seed.Client, seed.Workspace, nil, 2))
	t.Run("GetListInProject", testVariableGetList(seed.Client, seed.Workspace, seed.Project, 1))
	t.Run("Delete", testVariableDelete(seed.Client, seed.Workspace, &variable))
}

func TestVariableSsh(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
		project:   true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var variable buddy.Variable
	t.Run("Create", testVariableCreate(seed.Client, seed.Workspace, seed.Project, buddy.VariableTypeSshKey, true, false, &variable))
	t.Run("Update", testVariableUpdate(seed.Client, seed.Workspace, seed.Project, &variable))
	t.Run("Get", testVariableGet(seed.Client, seed.Workspace, seed.Project, &variable))
	t.Run("GetList", testVariableGetList(seed.Client, seed.Workspace, nil, 1))
	t.Run("GetListInProject", testVariableGetList(seed.Client, seed.Workspace, seed.Project, 2))
	t.Run("Delete", testVariableDelete(seed.Client, seed.Workspace, &variable))
}
