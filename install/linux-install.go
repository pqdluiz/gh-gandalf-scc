package install

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func installSCCLinux() error {
	// URL do arquivo para download e checksum esperado
	downloadURL := "https://github.com/boyter/scc/releases/download/v3.3.5/scc-v3.3.5-linux-amd64.tar.gz"
	expectedChecksum := "checksum_aqui" // Substitua com o checksum SHA256 fornecido pelo fabricante

	// Baixar o arquivo
	tarFilePath := "scc-linux-amd64.tar.gz"
	resp, err := http.Get(downloadURL)
	if err != nil {
		return fmt.Errorf("erro ao baixar o arquivo: %v", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(tarFilePath)
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("erro ao salvar o arquivo: %v", err)
	}

	// Calcular o checksum do arquivo baixado
	calculatedChecksum, err := checksum.calculateChecksum(tarFilePath)
	if err != nil {
		return fmt.Errorf("erro ao calcular o checksum: %v", err)
	}

	// Comparar o checksum calculado com o checksum esperado
	if calculatedChecksum != expectedChecksum {
		return fmt.Errorf("checksum do arquivo não corresponde ao esperado")
	}

	// Extrair o arquivo TAR.GZ
	cmd := exec.Command("tar", "-xzf", tarFilePath)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("erro ao extrair o arquivo TAR.GZ: %v", err)
	}

	// Mover o binário para o diretório PATH
	err = exec.Command("sudo", "mv", "scc", "/usr/local/bin/scc").Run()
	if err != nil {
		return fmt.Errorf("erro ao mover o binário para o diretório PATH: %v", err)
	}

	// Limpar o arquivo TAR.GZ
	err = os.Remove(tarFilePath)
	if err != nil {
		return fmt.Errorf("erro ao remover o arquivo TAR.GZ: %v", err)
	}

	return nil
}
