package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func testPublicKeyCreate(client *buddy.Client, out *buddy.PublicKey) func(t *testing.T) {
	return func(t *testing.T) {
		publicKey, _, err := GenerateRsaKeyPair()
		if err != nil {
			t.Fatal(ErrorFormatted("testPublicKeyCreate", err))
		}
		content := publicKey
		title := RandString(10)
		ops := buddy.PublicKeyOps{
			Content: &content,
			Title:   &title,
		}
		key, _, err := client.PublicKeyService.Create(&ops)
		if err != nil {
			t.Fatal(ErrorFormatted("PublicKeyService.Create", err))
		}
		err = CheckPublicKey(key, title, content, 0)
		if err != nil {
			t.Fatal(err)
		}
		*out = *key
	}
}

func testPublicKeyGet(client *buddy.Client, out *buddy.PublicKey) func(t *testing.T) {
	return func(t *testing.T) {
		key, _, err := client.PublicKeyService.Get(out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("PublicKeyService.Get", err))
		}
		err = CheckPublicKey(key, out.Title, out.Content, out.Id)
		if err != nil {
			t.Fatal(err)
		}
		*out = *key
	}
}

func testPublicKeyDelete(client *buddy.Client, out *buddy.PublicKey) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.PublicKeyService.Delete(out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("PublicKeyService.Delete", err))
		}
	}
}

func TestPublicKey(t *testing.T) {
	seed, err := SeedInitialData(nil)
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var key buddy.PublicKey
	t.Run("Create", testPublicKeyCreate(seed.Client, &key))
	t.Run("Get", testPublicKeyGet(seed.Client, &key))
	t.Run("Delete", testPublicKeyDelete(seed.Client, &key))
}
