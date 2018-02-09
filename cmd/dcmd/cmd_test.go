package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestUsage(t *testing.T) {
	prepareTest(t)
	defer os.Remove("$GOPATH/bin/dcmd")
	cmd := exec.Command("dcmd", "-h")
	_, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal("Expected exit code 1 but 0")
	}
}

func TestInvalidArgs(t *testing.T) {
	prepareTest(t)
	expectString := "No help topic for 'fqefhfjaa'\n"
	defer os.Remove("$GOPATH/bin/dcmd")
	cmd := exec.Command("dcmd", "fqefhfjaa")
	b, _ := cmd.CombinedOutput()

	if expectString != string(b) {
		t.Fatalf("Expected string is : %s", expectString)
	}
}

func prepareTest(t *testing.T) {
	runCmd(t, "go", "install")
}

func runCmd(t *testing.T, cmd string, args ...string) []byte {
	b, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		t.Fatalf("Expected %v, but %v: %v", nil, err, string(b))
	}
	return b
}
