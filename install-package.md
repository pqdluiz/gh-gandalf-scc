# CLI Command Integration

Este guia fornece instruções sobre como baixar e configurar o [CLI Command](https://github.com/pqdluiz/cli-command/blob/main/cli-command) para ser executado como parte de um hook de pre-push em seu projeto.

## Passo 1: Baixar e Instalar o Binário

1. **Baixe o binário para o seu sistema operacional**:

   Para sistemas Linux:

   ```sh
   curl -L -o cli-command https://github.com/pqdluiz/cli-command/releases/download/v1.0.0/cli-command-linux-amd64
   ```

Para sistemas macOS:

```sh
curl -L -o cli-command https://github.com/pqdluiz/cli-command/releases/download/v1.0.0/cli-command-darwin-amd64
```

Para sistemas Windows (a ser convertido para `.exe`):

```sh
curl -L -o cli-command.exe https://github.com/pqdluiz/cli-command/releases/download/v1.0.0/cli-command-windows-amd64.zip
```

2. **Torne o binário executável**:

   Para sistemas Unix (Linux/macOS):

   ```sh
   chmod +x cli-command
   ```

3. **Coloque o binário na raiz do seu projeto**:

   Mova o binário para a raiz do seu projeto (ou para um diretório de sua escolha).

   ```sh
   mv cli-command ./cli-command
   ```

## Passo 2: Configurar o Pre-Push Hook

1. **Crie um hook de pre-push usando Husky**:

   Se você ainda não tiver o Husky instalado, adicione-o ao seu projeto:

   ```sh
   npm install husky --save-dev
   npx husky install
   ```

2. **Adicione o script de pre-push**:

   Crie um arquivo chamado `pre-push` no diretório `.husky/` do seu projeto com o seguinte conteúdo:

   ```sh
   #!/bin/sh
   . "$(dirname "$0")/_/husky.sh"

   npm install
   npm run lint
   npm run test
   ./cli-command

   # Verifique o código de saída do script de análise
   if [ $? -ne 0 ]; then
     echo "Push abortado devido a alta complexidade no código."
     exit 1
   fi
   ```

3. **Torne o script executável**:

   ```sh
   chmod +x .husky/pre-push
   ```

## Considerações

- Certifique-se de substituir os links e caminhos para o binário conforme necessário, dependendo do seu sistema operacional.
- Verifique se o binário `cli-command` está na mesma pasta que o script `pre-push` ou ajuste o caminho no script conforme necessário.
- Este guia pressupõe que você já tem o [Node.js](https://nodejs.org/) e o [npm](https://www.npmjs.com/) instalados.

Com estas etapas, você terá configurado o binário `cli-command` para ser executado automaticamente antes de cada push, garantindo que o código do seu projeto esteja dentro dos padrões definidos.

Se você encontrar problemas ou tiver dúvidas adicionais, consulte a [documentação do CLI Command](https://pkg.go.dev/github.com/pqdluiz/cli-command/pkg/analyzer) ou entre em contato com o mantenedor do projeto.

```

### Notas

- **Links e Caminhos**: Certifique-se de que os links de download do binário e os caminhos para o binário estejam corretos e atualizados conforme a versão e o sistema operacional que você está usando.
- **Permissões**: Assegure-se de que o binário e o script do Husky têm permissões de execução.
- **Sistema Operacional**: Adapte os comandos de acordo com o sistema operacional do usuário final.

Se precisar de mais alguma coisa ou ajustes específicos, me avise!
```
