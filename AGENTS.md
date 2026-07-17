# AGENTS.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project overview

`cobbler-cli` is an independent Go CLI client for the [Cobbler](https://cobbler.readthedocs.io) provisioning server.
It talks to a Cobbler server over XML-RPC via the `github.com/cobbler/cobblerclient` library. The CLI is built with
`spf13/cobra` (commands) and `spf13/viper` (config file / env var loading).

Current line targets Cobbler **4.0.0** backend and cobblerclient `v1.x` — this was a clean break from the `v0.x` CLI
line, which talks to Cobbler 3.3.x via cobblerclient `v0.5.x`. See `MIGRATION.md` for the full list of breaking changes
(removed item types, moved interface flags, UID-vs-name reference semantics) and `CHANGELOG.md` for the current
unreleased changes.

## Build, lint, test

```bash
make build             # go build -o cobbler main.go, then regenerates shell completions
make run               # build and run the binary
make test              # runs testing/start.sh then `go test -v -coverprofile=coverage.out -covermode=atomic ./...`
make clean             # go clean, remove binary and generated completions
```

Linting is `golangci-lint` (v2.9.0, run via `golangci-lint run` — no repo-local `.golangci.yml`, defaults from CI at
`.github/workflows/lint.yml`). Docs (`docs/`) are linted separately with `rstcheck` and `doc8`.

### Tests are integration tests against a real Cobbler server

There are no mocks — `cmd/*_test.go` files execute the actual cobra commands (`rootCmd.Execute()`) against a live
Cobbler XML-RPC server. `make test` / `./testing/start.sh <server_url>` will:

1. Clone/checkout a pinned commit of `github.com/cobbler/cobbler` into `testing/cobbler_source`.
2. Download and extract a legacy Ubuntu 20.04 ISO into `extracted_iso_image/` (used as fake kernel/initrd fixtures for
   distros).
3. Build a `cobbler-dev` Docker image and bring up `testing/compose.yml`.
4. Poll the server URL until Cobbler answers, then run `go test`.

This requires Docker and `xorriso` locally. To run a single test **after** the server is already up (e.g. iterating on
one file without re-running the whole setup):

```bash
go test -v -run Test_DistroAddCmd ./cmd/...
```

Tests read connection settings from `testing/.cobbler.yaml` (passed via `--config` in test args, or loaded through
`setupClient(t)` in `cmd/testing.go`). Test helpers of note:

- `cobbler.FailOnError(t, err)` — from the cobblerclient package, fails the test on non-nil error.
- `FailOnNonEmptyStream(t, buf)` (`cmd/testing.go`) — asserts stderr stayed empty.
- Every test that creates an object registers a `defer` cleanup that deletes it via the client directly (not through the
  CLI).

### Commit messages

Commit messages are linted with commitlint (`commitlint.config.mjs`, conventional-commits) via a pre-commit hook
(`.pre-commit-config.yaml`). Header ≤ 72 chars, body/footer lines ≤ 120 chars.

## Architecture

### Command structure

All CLI code lives in `cmd/`. Every Cobbler item type (`distro`, `profile`, `system`, `image`, `menu`, `repo`, etc.)
has its own `<item>.go` file following the same shape:

- `New<Item>Cmd() (*cobra.Command, error)` builds the parent command (e.g. `distro`) and wires up its subcommands
  (`add`, `copy`, `edit`, `find`, `list`, `remove`, `rename`, `report`, and often `export`).
- `New<Item>AddCmd`, `New<Item>EditCmd`, etc. build each subcommand. `add`/`copy`/`edit`/`rename` all funnel field
  updates through a shared `update<Item>FromFlags(cmd, item)` function that walks `cmd.Flags().Visit(...)` and only
  touches fields whose flags were explicitly set — this is what makes partial edits work without clobbering unset
  fields.
- Every `RunE` starts with `generateCobblerClient()` (`cmd/utils.go`) to log into the server before doing anything
  else.
- New group item types (`distro_group`, `profile_group`, `system_group` — added for 4.0.0) share logic in
  `cmd/group_common.go` instead of each having their own copy.

`cmd/root.go` (`NewRootCmd`) is the single place every top-level command is registered; add new item commands there.

### Flag metadata (`cmd/metadata.go`)

Flags aren't declared ad hoc — they're defined once as data in `cmd/metadata.go` using the generic `FlagMetadata[T]`
struct (`Name`, `DefaultValue`, `Usage`, `IsInheritable`), grouped into maps like `distroStringFlagMetadata`,
`profileMapFlagMetadata`, `systemBoolFlagMetadata`, etc. `cmd/item.go` has generic registration helpers
(`addStringFlags`, `addBoolFlags`, `addIntFlags`, `addFloatFlags`, `addStringSliceFlags`, `addMapFlags`) that iterate
these maps and call the matching `cobra.Command.Flags()` method.

For fields that support Cobbler's "inherit from parent" semantics (`IsInheritable: true`), the flag registration helpers
automatically also add a `--<flag>-inherit` boolean sibling flag (e.g. `--owners` / `--owners-inherit`) and mark them
mutually exclusive where applicable. `update<Item>FromFlags` checks whether the `-inherit` flag was `.Changed` to decide
whether to set `IsInherited` or the concrete value.

When adding a new flag to an item type: add an entry to the relevant metadata map in `cmd/metadata.go`, then add a
`case` for it in that item's `update<Item>FromFlags` switch in `<item>.go` (or `group_common.go` for group types).

### Value objects and "in-place" edits

Fields like `AutoinstallMeta`, `KernelOptions`, `KernelOptionsPost`, `BootLoaders`, `Owners` are cobblerclient
`Value[T]` wrapper objects with `Data` and `IsInherited`. Some commands (`distro edit/copy/rename`, similarly for other
item types) accept an `--in-place` flag: when set, map-valued fields (kernel options, autoinstall meta, template files)
are updated via `Client.ModifyItemInPlace(itemType, name, field, value)` (merge semantics on the server) instead of
being replaced wholesale via `Client.Update<Item>`. This distinction matters — read the `inPlace` branching in
`update<Item>FromFlags` before changing that logic.

### Cross-cutting helpers (`cmd/utils.go`, `cmd/item.go`)

- `printStructured` / `printField` / `printValueStructured` (`cmd/root.go`) — reflection-based generic printer used by
  every `report` subcommand to dump an item's fields, handling `Value[T]` wrappers, nested `NetworkInterface` maps,
  and `ctime`/`mtime` float→UTC-time conversion.
- `printDumpVars` (`cmd/utils.go`) — separate reflection-based printer for `dumpvars` output (raw
  `map[string]interface{}`, not a struct).
- `FindItemNames` (`cmd/item.go`) — shared implementation backing every `<item> find` subcommand. Builds a criteria
  map from whatever flags the user set, and supports pagination (`--page`/`--items-per-page`, added via
  `addPaginationFlags`) by routing through `Client.FindItemsPaged` instead of `Client.FindItemNames`.
- `RemoveItemRecursive` (`cmd/item.go`) — shared implementation for recursive item deletion (`--recursive`).
- Export subcommands (`<item> export`) marshal items to JSON or YAML (`--format`); shared via `writeExport` in
  `cmd/group_common.go` for group types, inlined per-item elsewhere (e.g. `NewDistroExportCmd`).

### Network interfaces are a separate command

Unlike `distro`/`profile`/`system`, network interfaces are managed by the standalone `cobbler interface` command
(`cmd/interface.go`), not flags on `cobbler system`. This is a 4.0.0-era change — see `MIGRATION.md` for the full
old-flag → new-flag mapping if working on interface-related code.

### Config resolution

`cmd/root.go`'s `initConfig()` (run via `cobra.OnInitialize`) loads `$HOME/.cobbler.yaml` (or `--config <path>`) through
viper, with defaults `server_url=http://127.0.0.1/cobbler_api`, `server_username=cobbler`, `server_password=cobbler`,
overridable by environment variables (`viper.AutomaticEnv()`).
