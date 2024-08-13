package install

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

// Função para calcular o checksum SHA256 de um arquivo
func calculateChecksum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
