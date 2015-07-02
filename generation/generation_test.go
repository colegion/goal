package generation

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestStart_IncorrectSubcommand(t *testing.T) {
	defer expectPanic("requested incorrect subcommand and thus expected panic")
	Start("generate", map[string]string{
		"generate": "commandThatDoesNotExist",
	})

	// Remove the directory we have created.
	os.RemoveAll(filepath.Join("./assets"))
}

func TestStart_Handlers(t *testing.T) {
	Start("generate", map[string]string{
		"generate": "handlers",
	})

	// Remove the directory we have created.
	os.RemoveAll("./assets")
}

func TestStart_Listing(t *testing.T) {
	Start("generate", map[string]string{
		"generate":  "listing",
		"--input":   "./testdata/views",
		"--output":  "./testdata/assets/views",
		"--package": "views",
	})

	cmd := exec.Command("go", "run", "testdata/listing/main.go")
	cmd.Stderr = os.Stderr // Show the output of the program we run.
	if err := cmd.Run(); err != nil {
		t.Errorf("There are problems with generated listing, error: '%s'.", err)
	}

	// Remove the directory we have created.
	os.RemoveAll("./testdata/assets")
}

func expectPanic(msg string) {
	if err := recover(); err == nil {
		panic(msg)
	}
}
