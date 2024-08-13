package analyzer

import (
	"os/exec"
	"runtime"
	"testing"

	"github.com/pqdluiz/cli-command/install"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockExecutor é um mock para exec.Command
type MockExecutor struct {
	mock.Mock
}

// ExecCommand é uma função para facilitar o uso do MockExecutor
var ExecCommand = exec.Command

// TestInstallSCCWindows testa a instalação do SCC no Windows
func TestInstallSCCWindows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skipf("Skipping test for Windows installation on %s", runtime.GOOS)
	}

	mockExec := new(MockExecutor)
	ExecCommand = func(name string, arg ...string) *exec.Cmd {
		cmd := &exec.Cmd{}
		mockExec.On("Run").Return(nil)
		return cmd
	}

	err := install.InstallSCC()
	assert.NoError(t, err)

	mockExec.AssertExpectations(t)
}

// TestInstallSCCLinux testa a instalação do SCC no Linux
func TestInstallSCCLinux(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skipf("Skipping test for Linux installation on %s", runtime.GOOS)
	}

	mockExec := new(MockExecutor)
	ExecCommand = func(name string, arg ...string) *exec.Cmd {
		cmd := &exec.Cmd{}
		mockExec.On("Run").Return(nil)
		return cmd
	}

	err := install.InstallSCC()
	assert.NoError(t, err)

	mockExec.AssertExpectations(t)
}

// TestInstallSCCMac testa a instalação do SCC no macOS
func TestInstallSCCMac(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skipf("Skipping test for macOS installation on %s", runtime.GOOS)
	}

	mockExec := new(MockExecutor)
	ExecCommand = func(name string, arg ...string) *exec.Cmd {
		cmd := &exec.Cmd{}
		mockExec.On("Run").Return(nil)
		return cmd
	}

	err := install.InstallSCC()
	assert.NoError(t, err)

	mockExec.AssertExpectations(t)
}

// TestInstallSCC testa a instalação do SCC dependendo do sistema operacional
func TestInstallSCC(t *testing.T) {
	switch runtime.GOOS {
	case "windows":
		t.Run("windows", TestInstallSCCWindows)
	case "linux":
		t.Run("linux", TestInstallSCCLinux)
	case "darwin":
		t.Run("mac", TestInstallSCCMac)
	default:
		t.Skipf("Skipping test for unsupported OS: %s", runtime.GOOS)
	}
}
