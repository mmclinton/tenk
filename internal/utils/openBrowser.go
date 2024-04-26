package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OpenDefaultBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return fmt.Errorf("cannot open url on unsupported platform")
	}
	return cmd.Run()
}
