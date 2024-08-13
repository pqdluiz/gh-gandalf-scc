# Documentação do Fluxo de Análise de Código

## Visão Geral

Este documento descreve o fluxo de análise de código implementado no projeto Go. O fluxo utiliza o pacote `scc` (Source Code Counter) para avaliar a complexidade ciclomática dos arquivos de código e garantir que eles não excedam um limite definido. A instalação e execução do `scc` são automatizadas e adaptadas para diferentes sistemas operacionais.

## Estrutura do Pacote

O pacote `analyzer` contém as seguintes funções e responsabilidades:

1. **Instalação do `scc`**:

   - Instala o `scc` com base no sistema operacional (Windows, Linux, macOS).

2. **Análise de Arquivos**:
   - **`AnalyzeFile(filePath string) (*FileAnalysis, error)`**: Executa o `scc` em um arquivo e retorna a análise em formato JSON.
   - **`GetAllFiles(dirPath string) ([]string, error)`**: Obtém todos os arquivos de um diretório especificado.
   - **`AnalyzeAllFiles(pathToAnalyze string) error`**: Analisa todos os arquivos em um diretório e verifica a complexidade ciclomática.

## Instalação do `scc`

### Instalação no Windows

Para instalar o `scc` no Windows, o pacote `analyzer` executa um comando PowerShell que baixa e descompacta o arquivo binário do `scc`.

```go
func installSCCWindows() error {
	cmd := exec.Command("powershell", "-Command", "Invoke-WebRequest -Uri https://github.com/boyter/scc/releases/download/v1.0.0/scc-windows-amd64.zip -OutFile scc.zip; Expand-Archive scc.zip -DestinationPath .")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("erro ao instalar o scc no Windows: %v", err)
	}
	return nil
}
```

### Instalação no Linux

Para Linux, o `scc` é baixado, extraído e movido para o diretório `/usr/local/bin`:

```go
func installSCCLinux() error {
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
```

### Instalação no macOS

No macOS, o `scc` é instalado usando o Homebrew:

```go
func installSCCMac() error {
	cmd := exec.Command("brew", "install", "scc")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("erro ao instalar o scc no macOS: %v", err)
	}
	return nil
}
```

## Funções de Análise

### `AnalyzeFile`

Executa o `scc` em um arquivo específico e analisa o resultado JSON para retornar a análise detalhada.

```go
func AnalyzeFile(filePath string) (*FileAnalysis, error) {
	cmd := exec.Command("scc", "--by-file", "--format", "json", filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error executing scc on file %s: %v", filePath, err)
	}

	var results []FileAnalysis
	if err := json.Unmarshal(output, &results); err != nil {
		return nil, fmt.Errorf("error parsing JSON output for file %s: %v", filePath, err)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no results found for file %s", filePath)
	}

	return &results[0], nil
}
```

### `GetAllFiles`

Obtém todos os arquivos de um diretório, ignorando subdiretórios.

```go
func GetAllFiles(dirPath string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(dirPath, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !entry.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
```

### `AnalyzeAllFiles`

Analisa todos os arquivos no diretório especificado e verifica se a complexidade ciclomática excede o limite definido.

```go
func AnalyzeAllFiles(pathToAnalyze string) error {
	fmt.Println("Info: Start process", time.Now().Format(time.RFC3339))

	files, err := GetAllFiles(pathToAnalyze)
	if err != nil {
		return fmt.Errorf("error getting files: %v", err)
	}

	hasHighComplexity := false

	for _, file := range files {
		result, err := AnalyzeFile(file)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if result.Complexity > complexityLimit {
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

			hasHighComplexity = true
		}
	}

	fmt.Printf("Cyclomatic complexity analysis for repository with: %d files.\n", len(files))
	fmt.Println("Info: End process", time.Now().Format(time.RFC3339))

	if hasHighComplexity {
		fmt.Println("⚠️ Complexidade alta detectada. Push abortado.")
		return fmt.Errorf("complexidade alta detectada")
	}

	return nil
}
```

## Testes Unitários

Os testes unitários devem cobrir os seguintes cenários:

- **Teste de Instalação**:

  - Verificar se a função de instalação do `scc` executa corretamente para cada sistema operacional (Windows, Linux, macOS).

- **Teste de Análise de Arquivo**:

  - Simular a execução do comando `scc` e verificar se a análise de um arquivo retorna os resultados esperados.

- **Teste de Obtenção de Arquivos**:

  - Verificar se a função de obtenção de arquivos lista corretamente todos os arquivos de um diretório.

- **Teste de Análise de Todos os Arquivos**:
  - Testar se a função de análise de todos os arquivos verifica corretamente a complexidade ciclomática e lida com casos de erro.

## Desenvolvimento

### Build

Para compilar o projeto, execute:

```bash
go build -o cli-command cmd/main.go
```

### Run

Para executar o projeto, utilize:

```bash
./cli-command
```

### Test

Para rodar os testes, execute:

```bash
go test ./...
```
