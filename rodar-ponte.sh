#!/bin/zsh
# Uso: rodar-ponte.sh <conta> <porta>
# Roda a ponte de uma conta em contas/<conta>/store e reinicia se ela cair.
conta=$1
porta=$2
if [[ -z $conta || -z $porta ]]; then
  echo "uso: $0 <conta> <porta>"
  exit 1
fi

trap "exit 0" INT TERM

dir="$HOME/whatsapp-mcp/contas/$conta"
mkdir -p "$dir" && cd "$dir" || exit 1

while true; do
  WHATSAPP_BRIDGE_PORT=$porta "$HOME/whatsapp-mcp/whatsapp-bridge/whatsapp-bridge"
  echo "--- ponte $conta caiu, reiniciando em 3s ---"
  sleep 3
done
