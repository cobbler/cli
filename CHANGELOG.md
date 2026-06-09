# Changelog

## v1.0.0 (unreleased)

First release against the Cobbler 4.0.0 backend and cobblerclient `v1.x`.
Clean break — talks only to Cobbler 4.0.0+. For 3.3.x servers, stay on the
`v0.x` line.

### Breaking changes

- Removed `cobbler mgmtclass`, `cobbler package`, `cobbler file` (item types
  removed from the backend).
- Removed all interface-related flags from `cobbler system add/edit/copy/
  rename/find`. Network interfaces are now managed via the dedicated
  `cobbler interface` command, with nested `--ipv4-*` / `--ipv6-*` / `--dns-*`
  flag families that mirror the new `NetworkInterface` value objects.
- Removed `--mgmt-classes` and `--mgmt-parameters` flags from every item
  command.
- `cobbler list` and `cobbler report` no longer include `mgmtclasses`,
  `packages`, `files` sections.
- `cobbler setting edit` now auto-detects the setting's type via a
  `GetSettings` round-trip and rejects unknown names.
- `cobbler system dumpvars` and `cobbler profile dumpvars` now call the
  backend's `dump_vars` (formerly `get_blended_data`, which was removed in
  Cobbler 3.4.0).
- `--distro` (`profile`), `--parent`/`--menu` (`profile`), `--profile`/`--image`
  (`system`), and `--menu` (`image`) now require the referenced item's Cobbler
  UID, not its name (matches cobblerclient `v1.0.0`'s field semantics — see
  its changelog). Look up the UID first, e.g. `cobbler distro report --name
  <name>` prints a `uid` line.

### Added

- `cobbler template` (add/edit/copy/remove/rename/find/list/report/export +
  `content` + `refresh`).
- `cobbler distro-group`, `cobbler profile-group`, `cobbler system-group`.
- `--page` and `--items-per-page` flags on every `find` subcommand
  (routes through `find_items_paged`).
- `cobbler setting export`.

### Removed (no longer accepted)

- The flags listed in the migration guide.

### Deferred to a later release

- `cobbler transaction` subcommand. The cobblerclient `v1.x` wraps
  `TransactionBegin/Commit/Abort`, but per-token transactions don't span
  multiple shell invocations naturally. Scripted batches need a Go program
  today; a session-aware CLI mode is on the post-1.0 roadmap.

See [MIGRATION.md](MIGRATION.md) for upgrade guidance.
