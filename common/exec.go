package common

import (
	"fmt"
	"os"
	"os/exec"
)

func GoInstallCmd(packageName, installPath string) *exec.Cmd {
	execCmd := exec.Command("go", "install", packageName)
	execCmd.Env = append(os.Environ(), fmt.Sprintf("GOPATH=%s", installPath))
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	execCmd.Dir = os.TempDir()

	return execCmd
}

func GoCleanModCacheCmd(installPath string) *exec.Cmd {
	execCmd := exec.Command("go", "clean", "-modcache")
	execCmd.Env = append(os.Environ(), fmt.Sprintf("GOPATH=%s", installPath))
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	execCmd.Dir = os.TempDir()
	return execCmd
}
