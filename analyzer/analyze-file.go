package analyzer

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/pqdluiz/cli-command/models"
)

func AnalyzeFile(filePath string) (*models.FileAnalysis, error) {
	cmd := exec.Command("scc", "--by-file", "--format", "json", filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error executing scc on file %s: %v", filePath, err)
	}

	var results []models.FileAnalysis
	if err := json.Unmarshal(output, &results); err != nil {
		return nil, fmt.Errorf("error parsing JSON output for file %s: %v", filePath, err)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no results found for file %s", filePath)
	}

	return &results[0], nil
}
