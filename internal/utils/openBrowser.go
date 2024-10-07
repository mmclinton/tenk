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
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		return fmt.Errorf("cannot open url on unsupported platform")
	}

	return cmd.Start()
}
