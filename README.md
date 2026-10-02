# Apple Store CLI

CLI rápida, leve e scriptável para a **App Store Connect API**.

Criada por **Izaias Pertrelly**.

Com o comando `asc` você automatiza o ciclo de apps iOS, macOS, tvOS e visionOS direto do terminal, do IDE ou do CI/CD: autenticação, builds, TestFlight, metadados, screenshots, review, assinatura e publicação.

```bash
asc version
asc --help
```

## Requisitos

- macOS ou Linux
- Go na versão declarada em `go.mod` (só para build a partir do código)
- Uma API Key da App Store Connect (`.p8`)

## Instalação

```bash
git clone https://github.com/Izaiaspertrelly/apple-store-cli.git
cd apple-store-cli
make build
make install
```

O binário vai para `/usr/local/bin/asc` por padrão. Para outro destino:

```bash
make build
INSTALL_PREFIX="$HOME/.local/bin" make install
```

Confirme:

```bash
which asc
asc version
```

## Autenticação

Gere a chave em [App Store Connect → Integrações → API](https://appstoreconnect.apple.com/access/integrations/api).

Chave de equipe:

```bash
asc auth login \
  --name "MeuApp" \
  --key-id "ABC123" \
  --issuer-id "DEF456" \
  --private-key /caminho/para/AuthKey.p8 \
  --network
```

Chave individual (sem issuer ID):

```bash
asc auth login \
  --name "MinhaChave" \
  --key-id "ABC123" \
  --key-type individual \
  --private-key /caminho/para/AuthKey.p8
```

Em CI ou máquina sem Keychain:

```bash
asc auth login \
  --bypass-keychain \
  --name "CI" \
  --key-id "ABC123" \
  --issuer-id "DEF456" \
  --private-key /caminho/para/AuthKey.p8
```

Validar:

```bash
asc auth status --validate
asc auth doctor
```

## Primeiros comandos

```bash
asc apps list --output table
asc apps list --output json --pretty
asc status --app APP_ID
asc search "upload a build" --output json
```

Em terminal interativo o padrão de `--output` é `table`. Em pipe/CI, `json`. A flag `--output` sempre vence.

## Fluxos comuns

### TestFlight, crashes e feedback

```bash
asc testflight feedback list --app "123456789" --paginate
asc testflight crashes list --app "123456789" --sort -createdDate --limit 10
asc testflight crashes log --submission-id "SUBMISSION_ID"
asc testflight groups list --app "123456789" --output table
```

### Builds e distribuição

```bash
asc builds upload --app "123456789" --ipa "/caminho/para/MeuApp.ipa"
asc builds list --app "123456789" --output table
```

macOS (`.pkg`):

```bash
asc builds upload --app "123456789" --pkg "./build/MeuMacApp.pkg" --version "1.2.3" --build-number "42" --wait --output json
asc builds add-groups --app "123456789" --build-number "42" --version "1.2.3" --platform MAC_OS --group "Internal Testers"
```

### Publicar na App Store

```bash
asc release stage --app "123456789" --version "1.2.3" --build-id "BUILD_ID" --copy-metadata-from "1.2.2" --dry-run
asc publish appstore --app "123456789" --ipa "/caminho/para/MeuApp.ipa" --version "1.2.3" --submit --confirm
asc status --app "123456789" --watch
```

Validação antes do envio (o app passaria na review?):

```bash
asc validate --app "123456789"
asc validate --app "123456789" --version "1.2.3" --strict
asc validate --app "123456789" --version "1.2.3" --ipa "./MeuApp.ipa"
asc validate --app "123456789" --version "1.2.3" --deep
asc review status --app "123456789"
asc review doctor --app "123456789"
```

`asc validate` devolve um relatório de prontidão com um plano de correção em ordem: o primeiro item é o próximo a ajustar. Ele confere limites e placeholders de metadados, campos e localizações obrigatórios, dados de review completos, categoria, build anexado e processado, declaração de criptografia, direitos de conteúdo, preço no território base (Free conta), disponibilidade, screenshots e tamanhos, classificação etária e prontidão de assinaturas. Sem `--version`, ele escolhe a versão editável mais nova. Com `--ipa`, ele lê o `UIDeviceFamily` do build e bloqueia o envio se o app roda em iPad e faltam screenshots de iPad. Com `--deep`, ele usa a sessão web já salva para checar App Privacy, contratos e a primeira assinatura, e diz se cada problema se corrige pela API, pela web ou à mão.

### Metadados e localização

```bash
asc localizations list --app "123456789" --type app-info
asc metadata init --dir "./metadata" --version "1.2.3" --locale "pt-BR"
asc metadata apply --app "123456789" --version "1.2.3" --dir "./metadata" --dry-run
asc metadata keywords audit --app "123456789" --version "1.2.3"
```

### Screenshots

```bash
asc screenshots plan --app "123456789" --version "1.2.3" --review-output-dir "./screenshots/review"
asc screenshots apply --app "123456789" --version "1.2.3" --review-output-dir "./screenshots/review" --confirm
asc screenshots upload --version-localization "VERSION_LOCALIZATION_ID" --path "./screenshots/pt-BR" --device-type "IPHONE_65" --replace --max-screenshots 10
```

### Assinatura

```bash
asc bundle-ids capabilities list --bundle "BUNDLE_ID"
asc signing fetch --bundle-id com.exemplo.app --profile-type IOS_APP_STORE --output .asc/signing
```

### Workflow local

```bash
asc workflow validate --output json
asc workflow run --dry-run testflight_beta VERSION:1.2.3
```

Ajuda de qualquer comando:

```bash
asc --help
asc <comando> --help
asc <comando> <subcomando> --help
```

## Skills para agentes

O `asc` instala globalmente as skills de fluxo da App Store Connect, para o seu agente usar em qualquer projeto:

```bash
asc install-skills
```

O comando baixa o commit revisado `f52c4f04323bb2dfb21ca8be82e6494e9cd0b4d8` do repositório [rorkai/app-store-connect-cli-skills](https://github.com/rorkai/app-store-connect-cli-skills) e copia as 25 skills para a pasta global de skills do agente. Ele confere o pacote inteiro e cada arquivo instalado antes de concluir, preserva skills que não são dele e só precisa do `git`.

## Privacidade

A telemetria **não envia nada** a menos que você configure `ASC_TELEMETRY_ENDPOINT`. Para desligar de forma explícita:

```bash
asc telemetry disable
```

Ou `ASC_TELEMETRY_DISABLED=1` / `DO_NOT_TRACK=1`.

O CLI não inclui argumentos, credenciais, IDs de app, bundle IDs, hostnames nem caminhos de arquivo em eventos de telemetria.

## Documentação no repositório

- [docs/COMMANDS.md](docs/COMMANDS.md) — mapa de comandos
- [docs/WORKFLOWS.md](docs/WORKFLOWS.md) — fluxos reutilizáveis, incluindo Xcode → TestFlight
- [docs/CI_CD.md](docs/CI_CD.md) — integração com CI
- [CONTRIBUTING.md](CONTRIBUTING.md) — como contribuir

## Licença

MIT. Veja [LICENSE](LICENSE).

Criado e mantido por [Izaias Pertrelly](https://github.com/Izaiaspertrelly).

---

Ferramenta independente, não afiliada, endossada ou patrocinada pela Apple Inc. App Store Connect, TestFlight, Xcode Cloud e Apple são marcas da Apple Inc.
