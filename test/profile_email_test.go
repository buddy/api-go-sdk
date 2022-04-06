package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func testProfileEmailCreate(client *buddy.Client, email string) func(t *testing.T) {
	return func(t *testing.T) {
		ops := buddy.ProfileEmailOps{
			Email: &email,
		}
		pe, _, err := client.ProfileEmailService.Create(&ops)
		if err != nil {
			t.Fatal(ErrorFormatted("ProfileEmailService.Create", err))
		}
		if err := CheckFieldEqualAndSet("ProfileEmail.Email", pe.Email, email); err != nil {
			t.Fatal(err)
		}
		if err := CheckBoolFieldEqual("ProfileEmail.Confirmed", pe.Confirmed, false); err != nil {
			t.Fatal(err)
		}
	}
}

func testProfileEmailGetList(client *buddy.Client, count int) func(t *testing.T) {
	return func(t *testing.T) {
		list, _, err := client.ProfileEmailService.GetList()
		if err != nil {
			t.Fatal(ErrorFormatted("ProfileEmailService.GetList", err))
		}
		if err := CheckIntFieldEqual("len(Emails)", len(list.Emails), count); err != nil {
			t.Fatal(err)
		}
	}
}

func testProfileEmailDelete(client *buddy.Client, email string) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.ProfileEmailService.Delete(email)
		if err != nil {
			t.Fatal(ErrorFormatted("ProfileEmailService.Delete", err))
		}
	}
}

func TestProfileEmail(t *testing.T) {
	seed, err := SeedInitialData(nil)
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	email := RandEmail()
	t.Run("Create", testProfileEmailCreate(seed.Client, email))
	t.Run("GetList", testProfileEmailGetList(seed.Client, 2))
	t.Run("Delete", testProfileEmailDelete(seed.Client, email))
}
