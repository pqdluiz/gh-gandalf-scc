package analyzer

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/pqdluiz/cli-command/models"
)

const complexityLimit = 1

// AnalyzeAllFiles analisa a complexidade dos arquivos em diretórios específicos.
func AnalyzeAllFiles(directories []string) error {
	fmt.Println("Info: Start process", time.Now().Format(time.RFC3339))

	directoriesToAnalyze, err := getDirectoriesToAnalyze(".", directories)
	if err != nil {
		return err
	}

	if len(directoriesToAnalyze) == 0 {
		return fmt.Errorf("nenhum diretório na lista fornecida encontrado na raiz do projeto")
	}

	hasHighComplexity, err := analyzeDirectories(directoriesToAnalyze)
	if err != nil {
		return err
	}

	fmt.Printf("Cyclomatic complexity analysis for repository with: %d files.\n", len(directoriesToAnalyze))
	fmt.Println("Info: End process", time.Now().Format(time.RFC3339))

	if hasHighComplexity {
		fmt.Println("⚠️ Complexidade alta detectada. Push abortado.")
		return fmt.Errorf("complexidade alta detectada")
	}

	return nil
}

// getDirectoriesToAnalyze retorna a lista de diretórios a serem analisados com base na raiz e na lista fornecida.
func getDirectoriesToAnalyze(rootDir string, directories []string) ([]string, error) {
	dirEntries, err := os.ReadDir(rootDir)
	if err != nil {
		return nil, fmt.Errorf("error reading root directory: %v", err)
	}

	var directoriesToAnalyze []string
	for _, entry := range dirEntries {
		if entry.IsDir() && contains(directories, entry.Name()) {
			directoriesToAnalyze = append(directoriesToAnalyze, filepath.Join(rootDir, entry.Name()))
		}
	}

	return directoriesToAnalyze, nil
}

// contains verifica se um slice contém um valor específico.
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// analyzeDirectories analisa a complexidade dos arquivos em todos os diretórios fornecidos.
func analyzeDirectories(directories []string) (bool, error) {
	hasHighComplexity := false

	for _, dir := range directories {
		files, err := GetAllFiles(dir)
		if err != nil {
			return false, fmt.Errorf("error getting files from %s: %v", dir, err)
		}

		for _, file := range files {
			result, err := AnalyzeFile(file)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if result.Complexity > complexityLimit {
				printFileAnalysis(file, result)
				hasHighComplexity = true
			}
		}
	}

	return hasHighComplexity, nil
}

// printFileAnalysis imprime os detalhes da análise de um arquivo.
func printFileAnalysis(file string, result *models.FileAnalysis) {
	fmt.Printf("Arquivo: %s\n", file)
	fmt.Printf("Linhas: %d\n", result.Lines)
	fmt.Printf("Linhas de Código: %d\n", result.Code)
	fmt.Printf("Linhas de Comentário: %d\n", result.Comment)
	fmt.Printf("Linhas em Branco: %d\n", result.Blank)
	fmt.Printf("Complexidade: %d\n", result.Complexity)
	fmt.Println("--------------------------------------------")
	fmt.Printf("⚠️ Alta complexidade detectada! Por favor, refatore esse arquivo:\n%s\n", file)
	fmt.Printf("O limite de nível de complexidade permitido é de %d\n", complexityLimit)
	fmt.Println("Para saber mais detalhes sobre a verificação ciclomática de código:")
	fmt.Println("https://github.com/boyter/scc?tab=readme-ov-file#complexity-estimates")
}
