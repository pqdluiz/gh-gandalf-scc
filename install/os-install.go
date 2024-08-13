package install

import (
	"fmt"
	"os/exec"
)

func installSCCMac() error {
	// Instala SCC no macOS usando o Homebrew
	cmd := exec.Command("brew", "install", "scc")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("erro ao instalar o scc no macOS: %v", err)
	}
	return nil
}
