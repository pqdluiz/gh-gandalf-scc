package install

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func installSCCMac() error {
	// URL do arquivo para download e checksum esperado
	downloadURL := "https://github.com/boyter/scc/releases/download/v3.3.5/scc-macos-amd64.zip"
	expectedChecksum := "checksum_aqui" // Substitua com o checksum SHA256 fornecido pelo fabricante

	// Baixar o arquivo
	zipFilePath := "scc.zip"
	resp, err := http.Get(downloadURL)
	if err != nil {
		return fmt.Errorf("erro ao baixar o arquivo: %v", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(zipFilePath)
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("erro ao salvar o arquivo: %v", err)
	}

	// Calcular o checksum do arquivo baixado
	calculatedChecksum, err := checksum.calculateChecksum(zipFilePath)
	if err != nil {
		return fmt.Errorf("erro ao calcular o checksum: %v", err)
	}

	// Comparar o checksum calculado com o checksum esperado
	if calculatedChecksum != expectedChecksum {
		return fmt.Errorf("checksum do arquivo não corresponde ao esperado")
	}

	// Expandir o arquivo ZIP
	cmd := exec.Command("unzip", zipFilePath)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("erro ao expandir o arquivo ZIP: %v", err)
	}

	// Mover o binário para o diretório PATH
	err = exec.Command("sudo", "mv", "scc", "/usr/local/bin/").Run()
	if err != nil {
		return fmt.Errorf("erro ao mover o binário para o diretório PATH: %v", err)
	}

	// Limpar o arquivo ZIP
	err = os.Remove(zipFilePath)
	if err != nil {
		return fmt.Errorf("erro ao remover o arquivo ZIP: %v", err)
	}

	return nil
}
