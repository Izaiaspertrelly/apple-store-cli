# CI/CD Integrations

## GitHub Actions

Install `asc` using the official setup action:

```yaml
- uses: Izaiaspertrelly/setup-asc@v1
  with:
    version: latest

- run: asc --help
```

For end-to-end examples, see:
https://github.com/Izaiaspertrelly/setup-asc

## GitLab CI/CD Components

Use the official `asc-ci-components` repository:

```yaml
include:
  - component: github.com/Izaiaspertrelly/apple-store-cli/run@main
    inputs:
      stage: deploy
      job_prefix: release
      asc_version: latest
      command: asc --help
```

For install/run templates and self-managed examples:
https://github.com/Izaiaspertrelly/apple-store-cli

## Bitrise

Use the official `setup-asc` Bitrise step repository:

```yaml
workflows:
  primary:
    steps:
    - git::https://github.com/Izaiaspertrelly/setup-asc.git@main:
        inputs:
        - mode: run
        - version: latest
        - command: asc --help
```

## CircleCI

Use the official CircleCI orb repository:
https://github.com/Izaiaspertrelly/apple-store-cli
