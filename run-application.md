# Guia para Rodar a Extensão `gh-gandalf-scc` Localmente

Este guia descreve o processo para executar a extensão `gh-gandalf-scc` localmente, desde a configuração até a execução e testes. A extensão é uma ferramenta para análise de complexidade ciclomática no código-fonte, integrada ao GitHub CLI (`gh`).

## 1. **Pré-requisitos**

Antes de começar, verifique se você tem os seguintes requisitos:

- **GitHub CLI (`gh`):** Certifique-se de que o GitHub CLI está instalado. Você pode instalá-lo seguindo a [documentação oficial](https://cli.github.com/).
- **Go (Golang):** A extensão é desenvolvida em Go. Instale o Go se ainda não estiver instalado. Verifique a instalação com `go version`.
- **scc (Source Code Complexity):** Certifique-se de que o binário `scc` está disponível no PATH. Você pode instalar o `scc` seguindo as instruções na [documentação do scc](https://github.com/boyter/scc#installation).

## 2. **Clone o Repositório**

Clone o repositório da extensão para o seu ambiente local:

```sh
git clone https://github.com/pqdluiz/gh-gandalf-scc.git
cd gh-gandalf-scc
```

## 3. **Compilar o Código**

Compile o código Go para gerar o binário da extensão:

```sh
go build -o gh-gandalf-scc cmd/main.go
```

Isso criará um executável chamado `gh-gandalf-scc` no diretório atual.

## 4. **Instalar a Extensão Localmente**

Para testar a extensão localmente com o GitHub CLI, você pode instalá-la diretamente do diretório onde você compilou o binário:

```sh
gh extension install /path/to/gh-gandalf-scc
```

Substitua `/path/to/gh-gandalf-scc` pelo caminho absoluto onde o binário `gh-gandalf-scc` foi compilado.

## 5. **Executar a Extensão**

Depois de instalar a extensão localmente, você pode executá-la usando o GitHub CLI:

```sh
gh gandalf-scc
```

### **Executar com Parâmetros**

Se a sua extensão aceita parâmetros, você pode passá-los diretamente no comando. Por exemplo:

```sh
gh gandalf-scc --dir src
```

### **Exemplo Completo de Execução**

Aqui está um exemplo completo para garantir que a extensão está funcionando conforme o esperado:

1. **Clone e Entre no Repositório:**

   ```sh
   git clone https://github.com/pqdluiz/gh-gandalf-scc.git
   cd gh-gandalf-scc
   ```

2. **Compile o Código Go:**

   ```sh
   go build -o gh-gandalf-scc cmd/main.go
   ```

3. **Instale a Extensão Localmente:**

   ```sh
   gh extension install $(pwd)
   ```

4. **Execute a Extensão:**

   ```sh
   gh gandalf-scc
   ```

## 6. **Depuração e Testes**

Para depurar a extensão, você pode adicionar mensagens de log e executar o binário diretamente para verificar o comportamento. Para testar a extensão, execute o binário compilado com diferentes argumentos e verifique a saída.

### **Exemplo de Teste Direto**

```sh
./gh-gandalf-scc
```

Ou se preferir usar o GitHub CLI:

```sh
gh gandalf-scc --help
```

Isso deve exibir as opções disponíveis e garantir que a extensão está funcionando corretamente.

## 7. **Desinstalar a Extensão Localmente**

Se precisar desinstalar a extensão localmente, use o comando:

```sh
gh extension remove gh-gandalf-scc
```

## **Conclusão**

Seguindo este guia, você deve ser capaz de compilar, instalar e executar a extensão `gh-gandalf-scc` localmente. Se você encontrar problemas ou tiver dúvidas, verifique os logs e mensagens de erro para diagnosticar o problema ou consulte a documentação do repositório.
