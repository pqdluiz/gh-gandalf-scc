#!/usr/bin/env bash
set -e

# Obtenha o diretório onde o script está localizado
SCRIPT_DIR="$(dirname "$(realpath "$0")")"

# Caminho relativo para o binário dentro do repositório
BIN_PATH="$SCRIPT_DIR/cli-command-gandalf"

# Verifique se o binário existe e é executável
if [ ! -x "$BIN_PATH" ]; then
  echo "O binário do Go não foi encontrado ou não é executável: $BIN_PATH" >&2
  exit 1
fi

$BIN_PATH "$@"
