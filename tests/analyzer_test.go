package analyzer

import (
	"errors"
	"os/exec"
	"path/filepath"
	"testing"

	pkg "github.com/pqdluiz/cli-command/analyzer"
	"github.com/stretchr/testify/assert"
)

// MockWalker é um mock da interface Walker para testes
type MockWalker struct {
	WalkDirFunc func(root string, walkFn filepath.WalkFunc) error
}

func (mw *MockWalker) WalkDir(root string, walkFn filepath.WalkFunc) error {
	return mw.WalkDirFunc(root, walkFn)
}

// Run é um método mockado para simular o comportamento de exec.Command
func (m *MockExecutor) Run() error {
	args := m.Called()
	return args.Error(0)
}

// CombinedOutput é um método mockado para simular o comportamento de exec.Command
func (m *MockExecutor) CombinedOutput() ([]byte, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Error(1)
}

// TestAnalyzeFileSuccess simula a análise de arquivo com sucesso
/*func TestAnalyzeFileSuccess(t *testing.T) {
	mockExec := new(MockExecutor)
	ExecCommand = func(name string, arg ...string) *exec.Cmd {
		cmd := &exec.Cmd{}
		mockExec.On("CombinedOutput").Return([]byte(`[{"lines":10,"code":5,"comment":2,"blank":3,"complexity":1}]`), nil)
		return cmd
	}

	result, err := analyzer.AnalyzeFile("testfile.go")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.Complexity)

	mockExec.AssertExpectations(t)
}*/

// TestAnalyzeFileCommandError simula um erro ao executar o comando
func TestAnalyzeFileCommandError(t *testing.T) {
	mockExec := new(MockExecutor)
	ExecCommand = func(name string, arg ...string) *exec.Cmd {
		cmd := &exec.Cmd{}
		mockExec.On("CombinedOutput").Return(nil, errors.New("command error"))
		return cmd
	}

	result, err := pkg.AnalyzeFile("testfile.go")
	assert.Error(t, err)
	assert.Nil(t, result)

	mockExec.AssertExpectations(t)
}

// TestAnalyzeFileJSONError simula um erro ao analisar o JSON
func TestAnalyzeFileJSONError(t *testing.T) {
	mockExec := new(MockExecutor)
	ExecCommand = func(name string, arg ...string) *exec.Cmd {
		cmd := &exec.Cmd{}
		mockExec.On("CombinedOutput").Return([]byte(`invalid json`), nil)
		return cmd
	}

	result, err := pkg.AnalyzeFile("testfile.go")
	assert.Error(t, err)
	assert.Nil(t, result)

	mockExec.AssertExpectations(t)
}

// TestAnalyzeAllFilesHighComplexity simula a análise de todos os arquivos com alta complexidade
func TestAnalyzeAllFilesHighComplexity(t *testing.T) {
	mockExec := new(MockExecutor)
	ExecCommand = func(name string, arg ...string) *exec.Cmd {
		cmd := &exec.Cmd{}
		mockExec.On("CombinedOutput").Return([]byte(`[{"lines":10,"code":5,"comment":2,"blank":3,"complexity":3}]`), nil)
		return cmd
	}

	directories := []string{"web", "src", "scripts", "internal", "api", "pages", "cmd"}

	err := pkg.AnalyzeAllFiles(directories)
	assert.Error(t, err)
	assert.Equal(t, "complexidade alta detectada", err.Error())

	mockExec.AssertExpectations(t)
}

// TestAnalyzeAllFilesNoHighComplexity simula a análise de todos os arquivos sem alta complexidade
/*func TestAnalyzeAllFilesNoHighComplexity(t *testing.T) {
	mockExec := new(MockExecutor)
	ExecCommand = func(name string, arg ...string) *exec.Cmd {
		cmd := &exec.Cmd{}
		mockExec.On("CombinedOutput").Return([]byte(`[{"lines":10,"code":5,"comment":2,"blank":3,"complexity":1}]`), nil)
		return cmd
	}

	err := analyzer.AnalyzeAllFiles(".")
	assert.NoError(t, err)

	mockExec.AssertExpectations(t)
}*/

// TestAnalyzeAllFilesErr
