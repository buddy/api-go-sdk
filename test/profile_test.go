package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func testProfileUpdate(client *buddy.Client, name string) func(t *testing.T) {
	return func(t *testing.T) {
		ops := buddy.ProfileOps{
			Name: &name,
		}
		profile, _, err := client.ProfileService.Update(&ops)
		if err != nil {
			t.Fatal(ErrorFormatted("ProfileService.Update", err))
		}
		err = CheckProfile(profile, name)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testProfileGet(client *buddy.Client, name string) func(t *testing.T) {
	return func(t *testing.T) {
		profile, _, err := client.ProfileService.Get()
		if err != nil {
			t.Fatal(ErrorFormatted("ProfileService.Get", err))
		}
		err = CheckProfile(profile, name)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestProfile(t *testing.T) {
	seed, err := SeedInitialData(nil)
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	name := RandString(10)
	t.Run("Update", testProfileUpdate(seed.Client, name))
	t.Run("Get", testProfileGet(seed.Client, name))
}
