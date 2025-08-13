package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func testSsoGet404(client *buddy.Client, workspace *buddy.Workspace) func(t *testing.T) {
	return func(t *testing.T) {
		_, resp, err := client.SsoService.Get(workspace.Domain)
		if err == nil || resp.StatusCode != 404 {
			t.Fatal("Get SSO should throw 404 error")
		}
	}
}

func testSsoEnable(client *buddy.Client, workspace *buddy.Workspace) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.SsoService.Enable(workspace.Domain)
		if err != nil {
			t.Fatal(ErrorFormatted("SsoService.Enable", err))
		}
	}
}

func testSsoDisable(client *buddy.Client, workspace *buddy.Workspace) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.SsoService.Disable(workspace.Domain)
		if err != nil {
			t.Fatal(ErrorFormatted("SsoService.Disable", err))
		}
	}
}

func testSsoGet(client *buddy.Client, workspace *buddy.Workspace, sso *buddy.Sso) func(t *testing.T) {
	return func(t *testing.T) {
		ssoGet, _, err := client.SsoService.Get(workspace.Domain)
		if err != nil {
			t.Fatal(ErrorFormatted("SsoService.Get", err))
		}
		err = CheckSso(ssoGet, sso.Type, sso.SsoUrl, sso.Issuer, sso.Certificate, sso.SignatureMethod, sso.DigestMethod, sso.RequireSsoForAllMembers)
		if err != nil {
			t.Fatal(err)
		}
		*sso = *ssoGet
	}
}

func testSsoUpdateOidc(client *buddy.Client, workspace *buddy.Workspace, require bool, out *buddy.Sso) func(t *testing.T) {
	return func(t *testing.T) {
		issuer := "https://test.com/" + UniqueString()
		clientId := UniqueString()
		clientSecret := UniqueString()
		typ := buddy.SsoTypeOidc
		ops := buddy.SsoUpdateOps{
			Type:                    &typ,
			Issuer:                  &issuer,
			ClientId:                &clientId,
			ClientSecret:            &clientSecret,
			RequireSsoForAllMembers: &require,
		}
		sso, _, err := client.SsoService.Update(workspace.Domain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("SsoService.Update", err))
		}
		err = CheckSso(sso, typ, "", issuer, "", "", "", require)
		if err != nil {
			t.Fatal(err)
		}
		*out = *sso
	}
}

func testSsoUpdate(client *buddy.Client, workspace *buddy.Workspace, require bool, out *buddy.Sso) func(t *testing.T) {
	return func(t *testing.T) {
		ssoUrl := "https://login.microsoftonline.com/" + UniqueString() + "/saml2"
		issuer := "https://sts.windows.net/" + UniqueString()
		signature := buddy.SignatureMethodSha256
		digest := buddy.DigestMethodSha256
		cert, err := GenerateCertificate()
		if err != nil {
			t.Fatal(ErrorFormatted("GenerateCertificate", err))
		}
		ops := buddy.SsoUpdateOps{
			SsoUrl:                  &ssoUrl,
			Issuer:                  &issuer,
			SignatureMethod:         &signature,
			DigestMethod:            &digest,
			Certificate:             &cert,
			RequireSsoForAllMembers: &require,
		}
		sso, _, err := client.SsoService.Update(workspace.Domain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("SsoService.Update", err))
		}
		err = CheckSso(sso, buddy.SsoTypeSaml, ssoUrl, issuer, cert, signature, digest, require)
		if err != nil {
			t.Fatal(err)
		}
		*out = *sso
	}
}

func TestSso(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var sso buddy.Sso
	t.Run("404", testSsoGet404(seed.Client, seed.Workspace))
	t.Run("Enable", testSsoEnable(seed.Client, seed.Workspace))
	t.Run("Update", testSsoUpdate(seed.Client, seed.Workspace, false, &sso))
	t.Run("Get", testSsoGet(seed.Client, seed.Workspace, &sso))
	t.Run("Update", testSsoUpdate(seed.Client, seed.Workspace, true, &sso))
	t.Run("Get", testSsoGet(seed.Client, seed.Workspace, &sso))
	t.Run("Disable", testSsoDisable(seed.Client, seed.Workspace))
	t.Run("404", testSsoGet404(seed.Client, seed.Workspace))
	t.Run("Enable OIDC", testSsoEnable(seed.Client, seed.Workspace))
	t.Run("Update OIDC", testSsoUpdateOidc(seed.Client, seed.Workspace, false, &sso))
	t.Run("Get OIDC", testSsoGet(seed.Client, seed.Workspace, &sso))
	t.Run("Update OIDC", testSsoUpdateOidc(seed.Client, seed.Workspace, true, &sso))
	t.Run("Get OIDC", testSsoGet(seed.Client, seed.Workspace, &sso))
	t.Run("Disable OIDC", testSsoDisable(seed.Client, seed.Workspace))
	t.Run("404 OIDC", testSsoGet404(seed.Client, seed.Workspace))
}
