# AGENTS.md

## Skills first (mandatory)

- Before any execution (code, build, scaffold, template change), list available skills via the skill tool and load every skill relevant to the task (e.g. `caveman`/`ponytail` for style constraints, `improve-codebase-architecture` for structural changes).
- Do not skip this even for "trivial" tasks. Re-check skills when the task changes direction.

## What this repo is

- `fndn`: Go CLI scaffolding tool that generates Go backends (clean architecture). Module `github.com/daffadon/fndn`, Go >= 1.26 (see `README.MD`).
- Entrypoint: `main.go` -> `cmd.Execute()` (`cmd/init.go`). Commands: `fndn init [.]`, `fndn generate [framework|database|mq|cache|storage]` (wired in `cmd/init.go:20-28`, defined in `cmd/scaffold.go`, `cmd/generate.go`).
- Stack: `cobra` (CLI) + `bubbletea`/`bubbles`/`lipgloss` (interactive UI in `internal/ui/`). Generation logic: `internal/app/`, `internal/domain/`, output templates: `internal/template/`.

## Commands

- Run locally: `go run . --help`, `go run . init .`
- Build all platforms: `make build` -> `script/build.sh` (CGO_ENABLED=0, `linux/windows/darwin x amd64/arm64` into `bin/dist/`, runs `upx --best --lzma` except darwin + windows/arm64). Requires `upx` installed; `make` is the only task runner.
- No test/lint/typecheck config in repo — verify with `go build ./...` and `go vet ./...`; don't invent test commands.
- Release: push tag `v*` -> `.github/workflows/releaser.yml` runs GoReleaser (`goreleaser/goreleaser-action@v6`, Go 1.26 on CI). Do not hand-build release artifacts.

## Conventions / gotchas

- Templates live under `internal/template/`; UI flows under `internal/ui/cli/`. Keep template edits and prompt/machine logic in sync.
- Local skills live in `.agents/skills/` and are pinned in `skills-lock.json` — don't edit skills by hand; update via the skills mechanism.
- Windows/WSL note from README: generated projects using Air hot-reload have path issues on WSL — not this repo's code, but don't "fix" template `.air.toml` paths without checking that context.
