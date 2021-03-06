package dev

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const etcDir = "/etc/resolver"

func ConfigureResolver(domains []string, port int) error {
	err := os.MkdirAll(etcDir, 0755)
	if err != nil {
		return err
	}

	body := fmt.Sprintf(
		"# Generated by puma-dev\nnameserver 127.0.0.1\nport %d\n", port)

	for _, domain := range domains {
		path := filepath.Join(etcDir, domain)
		err := ioutil.WriteFile(path, []byte(body), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
