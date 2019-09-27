package onepassword

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetOnePassword(t *testing.T) {
	itemName := "itemName"
	itemValue := "itemValue"

	Convey("Given an OPERATOR_CONFIG env var", t, func() {
		backend := &OnePassword{}
		backend.Cli = &FakeCli{
			ItemName:  itemName,
			ItemValue: itemValue,
		}

		Convey("When retrieving a secret", func() {
			actualValue, err := backend.Get(itemName)
			Convey("Then no error is returned", func() {
				So(err, ShouldBeNil)
				So(actualValue, ShouldEqual, itemValue)
			})
		})
	})
}

func TestInitOnePassword(t *testing.T) {
	domain := "https://externalsecretoperator.1password.com"
	email := "externalsecretoperator@example.com"
	secretKey := "AA-BB-CC-DD-EE-FF-GG-HH-II-JJ"
	masterPassword := "MasterPassword12346!"
	vault := "production"

	backend := &OnePassword{
		Cli: &FakeCli{SignInOK: true},
	}

	params := map[string]string{
		"domain":         domain,
		"email":          email,
		"secretKey":      secretKey,
		"masterPassword": masterPassword,
		"vault":          vault,
	}

	err := backend.Init(params)
	if err != nil {
		t.Fail()
		fmt.Println("expected signin to succeed")
	}
}

func TestInitOnePassword_ErrSigninFailed(t *testing.T) {
	domain := "https://externalsecretoperator.1password.com"
	email := "externalsecretoperator@example.com"
	secretKey := "AA-BB-CC-DD-EE-FF-GG-HH-II-JJ"
	masterPassword := "MasterPassword12346!"
	vault := "production"

	backend := &OnePassword{
		Cli: &FakeCli{
			SignInOK: false,
		},
	}

	params := map[string]string{
		"domain":         domain,
		"email":          email,
		"secretKey":      secretKey,
		"masterPassword": masterPassword,
		"vault":          vault,
	}

	err := backend.Init(params)
	switch err.(type) {
	case *ErrSigninFailed:
	default:
		t.Fail()
		fmt.Println("expected signin failed error")
	}
}

func TestInitOnePassword_MissingEmail(t *testing.T) {
	Convey("Given a OnePasswordBackend", t, func() {
		domain := "https://externalsecretoperator.1password.com"
		secretKey := "AA-BB-CC-DD-EE-FF-GG-HH-II-JJ"
		masterPassword := "MasterPassword12346!"

		backend := NewBackend()

		Convey("When initializing", func() {
			params := map[string]string{
				"domain":         domain,
				"secretKey":      secretKey,
				"masterPassword": masterPassword,
			}

			So(backend.Init(params).Error(), ShouldEqual, "error reading 1password backend parameters: invalid init parameters: expected `email` not found")
		})
	})
}

func TestInitOnePassword_MissingDomain(t *testing.T) {
	Convey("Given a OnePasswordBackend", t, func() {
		email := "externalsecretoperator@example.com"
		secretKey := "AA-BB-CC-DD-EE-FF-GG-HH-II-JJ"
		masterPassword := "MasterPassword12346!"

		backend := NewBackend()

		Convey("When initializing", func() {
			params := map[string]string{
				"email":          email,
				"secretKey":      secretKey,
				"masterPassword": masterPassword,
			}

			So(backend.Init(params).Error(), ShouldEqual, "error reading 1password backend parameters: invalid init parameters: expected `domain` not found")
		})
	})
}

func TestInitOnePassword_MissingSecretKey(t *testing.T) {
	Convey("Given a OnePasswordBackend", t, func() {
		domain := "https://externalsecretoperator.1password.com"
		email := "externalsecretoperator@example.com"
		masterPassword := "MasterPassword12346!"

		backend := NewBackend()

		Convey("When initializing", func() {
			params := map[string]string{
				"email":          email,
				"domain":         domain,
				"masterPassword": masterPassword,
			}

			So(backend.Init(params).Error(), ShouldEqual, "error reading 1password backend parameters: invalid init parameters: expected `secretKey` not found")
		})
	})
}

func TestInitOnePassword_MissingMasterPassword(t *testing.T) {
	Convey("Given a OnePasswordBackend", t, func() {
		domain := "https://externalsecretoperator.1password.com"
		email := "externalsecretoperator@example.com"
		secretKey := "AA-BB-CC-DD-EE-FF-GG-HH-II-JJ"

		backend := NewBackend()

		Convey("When initializing", func() {
			params := map[string]string{
				"email":     email,
				"domain":    domain,
				"secretKey": secretKey,
			}

			So(backend.Init(params).Error(), ShouldEqual, "error reading 1password backend parameters: invalid init parameters: expected `masterPassword` not found")
		})
	})
}

func TestNewBackend(t *testing.T) {
	backend := NewBackend()

	if backend.(*OnePassword).Cli == nil {
		t.Fail()
		fmt.Println("expected backend to have a 1password cli")
	}

	expectedVault := "Personal"

	if backend.(*OnePassword).Vault != expectedVault {
		t.Fail()
		fmt.Printf("expected vault to be equal to '%s'", expectedVault)
	}
}
