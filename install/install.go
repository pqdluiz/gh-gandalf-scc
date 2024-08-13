package install

import (
	"fmt"
	"runtime"
)

// InstallSCC instala o SCC com base no sistema operacional.
func InstallSCC() error {
	switch runtime.GOOS {
	case "windows":
		return installSCCWindows()
	case "linux":
		return installSCCLinux()
	case "darwin":
		return installSCCMac()
	default:
		return fmt.Errorf("sistema operacional não suportado: %s", runtime.GOOS)
	}
}
