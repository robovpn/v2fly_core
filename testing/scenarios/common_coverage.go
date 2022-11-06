//go:build coverage
// +build coverage

package scenarios

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/robovpn/v2fly_core/common/uuid"
)

func Buildv2fly() error {
	genTestBinaryPath()
	if _, err := os.Stat(testBinaryPath); err == nil {
		return nil
	}

	cmd := exec.Command("go", "test", "-tags", "coverage coveragemain", "-coverpkg", "github.com/robovpn/v2fly_core/...", "-c", "-o", testBinaryPath, GetSourcePath())
	return cmd.Run()
}

func Runv2flyProtobuf(config []byte) *exec.Cmd {
	genTestBinaryPath()

	covDir := os.Getenv("v2fly_COV")
	os.MkdirAll(covDir, os.ModeDir)
	randomID := uuid.New()
	profile := randomID.String() + ".out"
	proc := exec.Command(testBinaryPath, "run", "-format=pb", "-test.run", "TestRunMainForCoverage", "-test.coverprofile", profile, "-test.outputdir", covDir)
	proc.Stdin = bytes.NewBuffer(config)
	proc.Stderr = os.Stderr
	proc.Stdout = os.Stdout

	return proc
}
