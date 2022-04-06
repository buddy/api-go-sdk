package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

const (
	ssh1 = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC5h1SgFvq45BGpYIDowIlaWiGe24kZg2DJ8NYqFo003PcAGdk30oJNvBqJfooGaI7GUkVoCzx9w3Oz/CYmC/NKsz45yUafJRwOBAQ4Gtt1o5RfLNIgj4GlfP1WmCFXIe4cuzureUkPCUIx2K+i1oAOdbEVorzfR3zqPIN/0u3Jwq3nLGmLYS8xCTq3odJT7GvyAj1jyOnXo+dpYZRm6LIteAkhtrnIAI+Le87Bp7JivPZwov/DG7HjW4IuStlTCJOQoYSUtTTu/zBWfSIbmZFakNqIBpiw8vmCeOOgBeOA5u4/JdfNMxH3CP0zxspjPplxkl3DiK/bBs1EGL0zvJrf test1"
)

func testPublicKeyCreate(client *buddy.Client, out *buddy.PublicKey) func(t *testing.T) {
	return func(t *testing.T) {
		content := ssh1
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
