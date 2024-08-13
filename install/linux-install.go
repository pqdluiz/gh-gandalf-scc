package install

import (
	"fmt"
	"os/exec"
)

func installSCCLinux() error {
	// Instala SCC no Linux
	cmd := exec.Command("wget", "https://github.com/boyter/scc/releases/download/v3.3.5/scc-v3.3.5-linux-amd64.tar.gz")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("erro ao baixar o scc no Linux: %v", err)
	}

	cmd = exec.Command("tar", "-xzf", "scc-linux-amd64.tar.gz")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("erro ao extrair o scc no Linux: %v", err)
	}

	cmd = exec.Command("sudo", "mv", "scc", "/usr/local/bin/scc")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("erro ao mover o scc para /usr/local/bin no Linux: %v", err)
	}

	return nil
}
