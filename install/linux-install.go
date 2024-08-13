package install

import (
	"fmt"
	"os/exec"
)

func installSCCLinux() error {
	// Instala SCC no Linux
	cmd := exec.Command("curl", "-LO", "https://github.com/boyter/scc/releases/download/v1.0.0/scc-linux-amd64.tar.gz")
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
