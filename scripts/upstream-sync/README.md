# Sincronização com o upstream

Este repositório deriva de [rorkai/App-Store-Connect-CLI](https://github.com/rorkai/App-Store-Connect-CLI) (licença MIT).

| Sync | Upstream | Data |
|------|----------|------|
| Fork inicial | `5.1.0` (`ca759a3`) | 2026-09-08 |
| Atualização | `5.9.1` + `main` (`8ba978af9`) | 2026-10-02 |

## O que é nosso (manter em todo sync)

- Rebrand: módulo `github.com/Izaiaspertrelly/apple-store-cli`, URLs, nomes de pacote (winget `Pertrelly.ASC`, orb `pertrelly/asc`), identificador de assinatura `com.pertrelly.asc`.
- Telemetria desligada: `internal/telemetry/client.go` usa `DefaultEndpoint = ""` e não envia nada sem endpoint.
- Contexto de runtime "Rork" renomeado para "hosted" (`ASC_HOSTED_SANDBOX_ID`, `RuntimeHostedSandbox`, `SourceHostedAgent`).
- Dados pessoais do autor original trocados por dados de teste (`Test Phone`, `com.example.*`).
- `README.md` próprio em PT-BR.

## Como sincronizar

1. Normalizar o caminho do módulo no upstream base e no upstream novo:
   `sed 's#rudrankriyam/App-Store-Connect-CLI#Izaiaspertrelly/apple-store-cli#g'`.
2. Fazer merge de 3 vias: base = upstream do último sync normalizado, lado A = este repo, lado B = upstream novo normalizado.
3. Rodar `perl -pi scripts/upstream-sync/rebrand.pl` nos arquivos que ainda citam `rorkai`, `rudrankriyam` ou `App-Store-Connect-CLI` (fora de `README.md`, `docs/openapi/` e `scripts/upstream-sync/`).
4. `git grep -niE 'rorkai|rudrankriyam|rork'` deve voltar vazio (exceto referências intencionais).
5. `go build ./... && go test ./...`.

## Exceção: skills

`asc install-skills` continua apontando para `rorkai/app-store-connect-cli-skills`. O instalador fixa um commit e os hashes das árvores desse repositório, e não existe `Izaiaspertrelly/apple-store-cli-skills`. Por isso `rebrand.pl` não mexe em `app-store-connect-cli-skills`.
