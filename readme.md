# Documentação da Extensão `gh-gandalf-scc`

## Descrição

A extensão `gh-gandalf-scc` é uma ferramenta desenvolvida para o GitHub CLI (`gh`) que realiza análise de complexidade ciclomática no código-fonte de um repositório. Utiliza a ferramenta `scc` (Source Code Complexity) para gerar relatórios detalhados sobre a complexidade do código.

## Funcionalidades

- **Análise de Complexidade Ciclomática:** Executa uma análise detalhada da complexidade ciclomática dos arquivos de código-fonte.
- **Integração com GitHub CLI:** Funciona como uma extensão do GitHub CLI, permitindo que você execute a análise diretamente de seus comandos `gh`.

## Instalação

A extensão `gh-gandalf-scc` inclui um fluxo para instalação automática do `scc` em Linux, macOS e Windows. Siga os passos abaixo para instalar e configurar a extensão.

### Passos para Instalação

1. **Instale o GitHub CLI:**

   Siga as instruções na [documentação oficial do GitHub CLI](https://cli.github.com/) para instalar o `gh`.

2. **Instale a Extensão `gh-gandalf-scc`:**

   No terminal, execute:

   ```sh
   gh extension install pqdluiz/gh-gandalf-scc
   ```

   A extensão se encarregará de verificar e instalar o `scc` automaticamente.

## Uso

### Comandos

1. **Executar Análise de Complexidade Ciclomática:**

   Após a instalação, você pode usar a extensão para realizar a análise de complexidade ciclomática com o comando:

   ```sh
   gh gandalf-scc
   ```

   Esse comando irá:

   - Analisar os arquivos de código fonte nos diretórios especificados.
   - Gerar um relatório de complexidade ciclomática usando o `scc`.

### Opções

- **Diretórios a serem analisados:** A extensão está configurada para analisar diretórios como `web`, `src`, `scripts`, `internal`, `api`, `pages`, e `cmd`. Você pode modificar esses diretórios no código da extensão se necessário.

## Desenvolvimento

### Estrutura do Projeto

- **Diretório `cmd/`:** Contém o código principal da extensão, incluindo o arquivo `main.go` que define a lógica e os comandos.
- **Diretório `analyzer/`:** Contém o código para análise de complexidade ciclomática usando o `scc`.

### Compilação

Para compilar a extensão localmente, execute:

```sh
go build -o gh-gandalf-scc cmd/main.go
```

### Testes

Para testar a extensão localmente, execute:

```sh
gh extension install https://github.com/pqdluiz/gh-gandalf-scc
gh gandalf-scc
```

## Contribuição

Se você deseja contribuir para o desenvolvimento da extensão, siga as diretrizes abaixo:

1. **Faça um Fork do Repositório:**

   - Crie um fork do repositório no GitHub.

2. **Clone o Repositório:**

   - Clone o repositório forkado para seu ambiente local.

3. **Faça suas Alterações:**

   - Implemente as alterações desejadas e teste a extensão localmente.

4. **Envie um Pull Request:**

   - Envie um pull request para o repositório principal com suas alterações.

## Licença

Esta extensão está licenciada sob a [Licença MIT](https://opensource.org/licenses/MIT).

---

### **Dúvidas e Suporte**

Para dúvidas ou suporte, você pode abrir uma issue no [repositório do GitHub](https://github.com/pqdluiz/gh-gandalf-scc/issues).

---

## Código da Instalação Automática do `scc`

Aqui está o trecho de código da extensão responsável pela instalação automática do `scc`:

```sh
#!/usr/bin/env bash

# Determine the OS and architecture
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Download and install the correct version of scc
case "$OS" in
  linux)
    URL="https://github.com/boyter/scc/releases/download/v3.3.5/scc-v3.3.5-linux-amd64.tar.gz"
    curl -LO $URL
    tar -xzf scc-v3.3.5-linux-amd64.tar.gz
    sudo mv scc /usr/local/bin/
    ;;
  darwin)
    URL="https://github.com/boyter/scc/releases/download/v3.3.5/scc-v3.3.5-darwin-amd64.tar.gz"
    curl -LO $URL
    tar -xzf scc-v3.3.5-darwin-amd64.tar.gz
    sudo mv scc /usr/local/bin/
    ;;
  *)
    echo "Unsupported OS: $OS"
    exit 1
    ;;
esac
```

Para Windows, a instalação do `scc` é feita de maneira similar com o uso do PowerShell para baixar e extrair o arquivo ZIP.
