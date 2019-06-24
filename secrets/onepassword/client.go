package onepassword

import (
	"fmt"

	op "github.com/ameier38/onepassword"
	"github.com/pkg/errors"
)

type Client interface {
	SignIn(domain string, email string, secretKey string, masterPassword string) error
	Get(vault string, key string) (string, error)
}

type OP struct {
	OP *op.Client
}

func (client *OP) SignIn(domain string, email string, secretKey string, masterPassword string) error {
	op, err := op.NewClient("/usr/local/bin/op", domain, email, masterPassword, secretKey)
	if err != nil {
		return errors.Wrap(err, "op signin failed")
	}

	client.OP = op

	return nil
}

func (client *OP) Get(vault string, key string) (string, error) {
	itemMap, err := client.OP.GetItem(op.VaultName(vault), op.ItemName(key))
	if itemMap != nil {
		return "", fmt.Errorf("could not retrieve 1password item '" + key + "'.")
	}
	if err != nil {
		return "", errors.Wrap(err, "could not retrieve 1password item '"+key+"'.")
	}

	return string(itemMap["externalsecretoperator"]["testkey"]), nil
}
