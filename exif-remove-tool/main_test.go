package main

import (
	"bytes"
	"os"
	"os/exec"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestImages(t *testing.T) {
	cmd := exec.Command(
		"go", "run", "main.go")

	b := new(bytes.Buffer)
	cmd.Stdout = b
	cmd.Stderr = b

	err := cmd.Run()
	actualRaw := b.Bytes()

	assert.Nil(t, err)

	var re = regexp.MustCompile(`0 fail`)
	matches := re.FindStringSubmatch(string(actualRaw))
	assert.True(t, len(matches) == 1, "Should parse all files in img folder")
}
