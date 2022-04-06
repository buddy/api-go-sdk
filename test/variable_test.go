package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

const (
	SshKey = `-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAABFwAAAAdzc2gtcn
NhAAAAAwEAAQAAAQEA573Yk5eUJh2mVKYByEQqFt6mSQsES6HhlT9H4iDjQWtK2g78vDm1
Ni9MmUfMyJ1IYcNhOZW5CfPlueqXsUobZyo1eAMI1/0Ju5bSwZwRkdR2htydCJOX0mG0KZ
zT6NKNf7lg0R3bHoQcCcbHnvzE1F8wIXWsDJyeo0iXb8kilNr0kFAkMmjZGKjeXbRKR374
48TcrSaeGvSlOpsAa6YvBGoRnfhSkFysE+FxTFRhF7iknrvhjLqwbK7BR5Pf3j8hifiy3i
tMJM230COXwCAg2BoHyzH6xefP4TE6Po2qVfAcNmUzp+ktVbqf2HH44aFiZJgZYXJCTct/
RopyD0Uq3QAAA9D3L7tx9y+7cQAAAAdzc2gtcnNhAAABAQDnvdiTl5QmHaZUpgHIRCoW3q
ZJCwRLoeGVP0fiIONBa0raDvy8ObU2L0yZR8zInUhhw2E5lbkJ8+W56pexShtnKjV4AwjX
/Qm7ltLBnBGR1HaG3J0Ik5fSYbQpnNPo0o1/uWDRHdsehBwJxsee/MTUXzAhdawMnJ6jSJ
dvySKU2vSQUCQyaNkYqN5dtEpHfvjjxNytJp4a9KU6mwBrpi8EahGd+FKQXKwT4XFMVGEX
uKSeu+GMurBsrsFHk9/ePyGJ+LLeK0wkzbfQI5fAICDYGgfLMfrF58/hMTo+japV8Bw2ZT
On6S1Vup/YcfjhoWJkmBlhckJNy39GinIPRSrdAAAAAwEAAQAAAQAG860BkHSDSDRrKae4
CENy+C7o1gnE8xA/V+yiHfZzSfKu4/A0/U4wV+7mUj8UbZN0S1YpUhKA9+4WS7FNQjncOG
nuNbkYMaEPHZEo+bOVOlhr50ZWsYbGauPqs6evvlE8WaVL4KdoHPJyYKIwZMjKzig1eMA2
iKRBpbXVRqVg7bn0+opBUdv5FpsDkSa+ijJKLA7szSpM7yq03sZfx9u1/WTvh1Qa375mFC
O/8NUJoGii5Tacp+QeIHi2IEl+38eBLx1mal0AM68mkeJvU6Pfa+f/aUG4Z9RfQByD2/ql
JmrLVzaEv1Jy3n4lFnTsDT11dLmdQ7WjCNa/NKvVOrVJAAAAgCX4+x7xjcKAxEdprXpKkN
m4L+4Ciy1I03x/fk0GNDi316IiEUylTzTCy7HFAg0RUaxyN3iyFbwe1kN6DPV4r877WiOV
lt7v7hIjS1eeOwSGpsxijPxbWEIzOnu35t23YUROCOWaYXP9EAk1YEqEB34UKVq9jeB2t+
7zh/0ORkbwAAAAgQD8c9R/g6+Jx5c09MY0nX8IiTp1YBws6WVXDkaPPY//t1nKSfYEZy4p
ojOgWAS8vDMd7g24gr0Fm0lEKBXhgPVJTuClIH1P6IKRacLNtpp9uzVjBn/hi3f60GXo8y
A8ZbizwLkSpECLjFUer+ZlbrlcZ2Oq6o8DufiIjFAiqH2RMwAAAIEA6v+Cxku/IZ5XLrnm
wItoQx8eLu/Ly88wCSe1rgwzazgV4zQz1Y6B8ONqY03SgnE1im0/zaVmhMsRyE8FNGHPZP
8dclO3b8qod57BiY3PLpBKi9swhNQlxQn1zDhaF5cDEcTXYNL6xXC6iD6aa0bffjdBkuzz
A10SiVeaAxv3c68AAAAYTWljaGFsQE1pY2hhcy1pTWFjLmxvY2FsAQID
-----END OPENSSH PRIVATE KEY-----`
	SshKey2 = `-----BEGIN OPENSSH PRIVATE KEY-----
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
-----END OPENSSH PRIVATE KEY-----`
)

func testVariableCreate(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, typ string, out *buddy.Variable) func(t *testing.T) {
	return func(t *testing.T) {
		key := RandString(10)
		val := RandString(10)
		desc := RandString(10)
		set := true
		enc := false
		fileChmod := ""
		filePath := ""
		fileName := ""
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
			val = SshKey
			fileChmod = "666"
			filePath = "~/.ssh/" + RandString(6)
			fileName = RandString(6)
			filePlace = buddy.VariableSshKeyFilePlaceContainer
			ops.FileChmod = &fileChmod
			ops.FileName = &fileName
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
		err = CheckVariable(variable, key, val, typ, desc, set, enc, fileName, filePath, fileChmod, filePlace, 0)
		if err != nil {
			t.Fatal(err)
		}
		*out = *variable
	}
}

func testVariableUpdate(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Variable) func(t *testing.T) {
	return func(t *testing.T) {
		val := RandString(10)
		desc := ""
		set := false
		enc := true
		fileName := ""
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
			val = SshKey2
			fileName = RandString(10)
			filePath = "/bec/" + RandString(5)
			filePlace = buddy.VariableSshKeyFilePlaceNone
			fileChmod = "600"
			ops.FileName = &fileName
			ops.FilePath = &filePath
			ops.FilePlace = &filePlace
			ops.FileChmod = &fileChmod
			ops.Value = &val
		}
		variable, _, err := client.VariableService.Update(workspace.Domain, out.Id, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("VariableService.Patch", err))
		}
		err = CheckVariable(variable, out.Key, val, out.Type, desc, set, enc, fileName, filePath, fileChmod, filePlace, out.Id)
		if err != nil {
			t.Fatal(err)
		}
		*out = *variable
	}
}

func testVariableGet(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Variable) func(t *testing.T) {
	return func(t *testing.T) {
		variable, _, err := client.VariableService.Get(workspace.Domain, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("VariableService.Get", err))
		}
		err = CheckVariable(variable, out.Key, out.Value, out.Type, out.Description, out.Settable, out.Encrypted, out.FileName, out.FilePath, out.FileChmod, out.FilePlace, out.Id)
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
	t.Run("Create", testVariableCreate(seed.Client, seed.Workspace, nil, buddy.VariableTypeVar, &variable))
	t.Run("Update", testVariableUpdate(seed.Client, seed.Workspace, &variable))
	t.Run("Get", testVariableGet(seed.Client, seed.Workspace, &variable))
	t.Run("GetList", testVariableGetList(seed.Client, seed.Workspace, nil, 1))
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
	t.Run("Create", testVariableCreate(seed.Client, seed.Workspace, seed.Project, buddy.VariableTypeSshKey, &variable))
	t.Run("Update", testVariableUpdate(seed.Client, seed.Workspace, &variable))
	t.Run("Get", testVariableGet(seed.Client, seed.Workspace, &variable))
	t.Run("GetList", testVariableGetList(seed.Client, seed.Workspace, nil, 0))
	t.Run("GetListInProject", testVariableGetList(seed.Client, seed.Workspace, seed.Project, 2))
	t.Run("Delete", testVariableDelete(seed.Client, seed.Workspace, &variable))
}
