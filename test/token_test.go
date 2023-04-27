package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
	"time"
)

func testTokenCreateAdvanced(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Token) func(t *testing.T) {
	return func(t *testing.T) {
		name := UniqueString()
		expiresIn := 25
		scopes := []string{"WORKSPACE", "USER_INFO"}
		workspaceRestrictions := []string{workspace.Domain}
		ipRestrictions := []string{"127.0.0.1"}
		ops := buddy.TokenOps{
			Name:                  &name,
			ExpiresIn:             &expiresIn,
			Scopes:                &scopes,
			WorkspaceRestrictions: &workspaceRestrictions,
			IpRestrictions:        &ipRestrictions,
		}
		token, _, err := client.TokenService.Create(&ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TokenService.Create", err))
		}
		err = CheckToken(token, name, expiresIn, "", scopes, workspaceRestrictions, ipRestrictions, "")
		if err != nil {
			t.Fatal(err)
		}
		*out = *token
	}
}

func testTokenDelete(client *buddy.Client, token *buddy.Token) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.TokenService.Delete(token.Id)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testTokenGetList(client *buddy.Client, count int) func(t *testing.T) {
	return func(t *testing.T) {
		tokens, _, err := client.TokenService.GetList()
		if err != nil {
			t.Fatal(ErrorFormatted("TokenService.Get", err))
		}
		err = CheckTokens(tokens, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testTokenGet(client *buddy.Client, token *buddy.Token) func(t *testing.T) {
	return func(t *testing.T) {
		getToken, _, err := client.TokenService.Get(token.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("TokenService.Get", err))
		}
		err = CheckToken(getToken, token.Name, 0, token.ExpiresAt, token.Scopes, token.WorkspaceRestrictions, token.IpRestrictions, token.Id)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testTokenCreateBasic(client *buddy.Client) func(t *testing.T) {
	return func(t *testing.T) {
		name := UniqueString()
		dt := time.Now().AddDate(0, 0, 10)
		expiresAt := dt.Format(time.RFC3339)
		scopes := []string{"WORKSPACE"}
		ops := buddy.TokenOps{
			Name:      &name,
			ExpiresAt: &expiresAt,
			Scopes:    &scopes,
		}
		token, _, err := client.TokenService.Create(&ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TokenService.Create", err))
		}
		err = CheckToken(token, name, 0, expiresAt, scopes, nil, nil, "")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestToken(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var token buddy.Token
	t.Run("CreateBasic", testTokenCreateBasic(seed.Client))
	t.Run("CreateAdvanced", testTokenCreateAdvanced(seed.Client, seed.Workspace, &token))
	t.Run("Get", testTokenGet(seed.Client, &token))
	t.Run("GetList", testTokenGetList(seed.Client, 3))
	t.Run("Delete", testTokenDelete(seed.Client, &token))
	t.Run("GetList2", testTokenGetList(seed.Client, 2))
}
