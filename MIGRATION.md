# Migration guide: cobbler-cli v0.x → v1.0.0

This release re-anchors `cobbler-cli` on the Cobbler 4.0.0 backend and the matching cobblerclient `v1.x` library. It is
a clean break — there is no compatibility shim against Cobbler 3.3.x. If you need to talk to a 3.3.x server, stay on
`cobbler-cli v0.0.1-rc*` (which pulls cobblerclient `v0.5.x`).

## Removed commands

These item types no longer exist in Cobbler 4.0.0 and the corresponding commands have been removed:

- `cobbler mgmtclass …`
- `cobbler package …`
- `cobbler file …`

`cobbler list` and `cobbler report` no longer include sections for these types.

The `--mgmt-classes` and `--mgmt-parameters` flags have been removed from every item command (`distro`, `profile`,
`system`).

## Network interfaces moved to a dedicated command

Network interfaces are first-class items in 4.0.0. Every `--*` flag related  to interfaces has been removed from
`cobbler system add/edit/copy/rename/find`. Use `cobbler interface` instead.

The flag names also reflect the new nested address objects:

| Old (`cobbler system …`)      | New (`cobbler interface …`)       |
| ----------------------------- | --------------------------------- |
| `--interface <name>`          | `--name <name>` plus `--system-name <s>` (on `add`) |
| `--ip-address`                | `--ipv4-address`                  |
| `--netmask`                   | `--ipv4-netmask`                  |
| `--if-gateway`                | `--ipv4-gateway`                  |
| `--static-routes`             | `--ipv4-static-routes`            |
| `--ipv6-address`              | `--ipv6-address` (unchanged)      |
| `--ipv6-prefix`               | `--ipv6-prefix` (unchanged)       |
| `--ipv6-default-gateway`      | `--ipv6-default-gateway` (unchanged) |
| `--ipv6-secondaries`          | `--ipv6-secondaries` (unchanged)  |
| `--ipv6-static-routes`        | `--ipv6-static-routes` (unchanged) |
| `--ipv6-mtu`                  | `--ipv6-mtu` (unchanged)          |
| `--dns-name`                  | `--dns-name` (unchanged)          |
| `--cnames`                    | `--dns-cnames`                    |
| `--mac-address`               | `--mac-address` (unchanged)       |
| `--bonding-opts`              | `--bonding-opts` (unchanged)      |
| `--bridge-opts`               | `--bridge-opts` (unchanged)       |
| `--interface-type`            | `--interface-type` (unchanged)    |
| `--interface-master`          | `--interface-master` (unchanged)  |
| `--connected-mode`            | `--connected-mode` (unchanged)    |
| `--management`                | `--management` (unchanged)        |
| `--static`                    | `--static` (unchanged)            |
| `--mtu`                       | `--mtu` (unchanged)               |
| `--virt-bridge`               | `--virt-bridge` (now inheritable: `--virt-bridge-inherit`) |
| `--delete-interface`          | `cobbler interface remove --name <n>` |
| `--rename-interface`          | `cobbler interface rename --name <old> --newname <new>` |

Lifecycle example:

```bash
cobbler system add --name=server1 --profile=$(cobbler profile report --name=rocky-9-x86_64 | grep '^uid' | awk '{print $NF}')
cobbler interface add --system-name=server1 --name=eth0 \
    --mac-address=aa:bb:cc:dd:ee:ff \
    --ipv4-address=10.0.0.5 --ipv4-netmask=255.255.255.0 --ipv4-gateway=10.0.0.1 \
    --dns-name=server1.lab.local
cobbler interface report --system-name=server1
```

## Cross-reference flags now take a UID, not a name

`--distro` (`profile`), `--parent`/`--menu` (`profile`), `--profile`/`--image` (`system`), and `--menu`
(`image`) now require the referenced item's Cobbler UID rather than its name. This matches
cobblerclient `v1.0.0`, whose `Distro`/`Parent`/`Menu`/`Profile`/`Image` fields hold UIDs directly
(Cobbler 4.0.0's server-side setters for these fields are UID-only regardless — this removes a
name↔UID translation layer that could produce stale references, e.g. after renaming a distro).

Look up the UID first, e.g.:

```bash
cobbler distro report --name=rocky-9-x86_64   # prints a "uid" line among other fields
cobbler profile add --name=rocky-9-x86_64 --distro=<the distro's uid>
```

## New commands

- `cobbler template` — manage Cobbler 4.0.0 `Template` items, including `cobbler template content --name=<n>` (dump
  rendered content to stdout) and `cobbler template refresh [--name=<n>...]` (reload from disk).
- `cobbler distro-group`, `cobbler profile-group`, `cobbler system-group` — manage the three group item types.
- Every `find` subcommand now supports `--page` and `--items-per-page`, routing through the backend's `find_items_paged`
  endpoint and emitting a trailing `# page N of M (T total)` summary.

## Strict setting types

`cobbler setting edit --name=<n> --value=<v>` now performs a round-trip to `get_settings` to discover the existing
setting's type, then parses `--value` accordingly. The 4.0.0 backend rejects mistyped values. Examples:

```bash
cobbler setting edit --name=server --value=cobbler.lab.local   # string
cobbler setting edit --name=http_port --value=8080             # int
cobbler setting edit --name=manage_dhcp --value=true           # bool
cobbler setting edit --name=cheetah_import_whitelist --value=foo,bar  # []string
```

Unknown setting names error out cleanly.

## Transactions: deferred

Cobbler 4.0.0 exposes per-token `transaction_begin/commit/abort`. The Go client (cobblerclient) wraps them as
`TransactionBegin`/`TransactionCommit`/ `TransactionAbort`. There is no `cobbler transaction` subcommand in v1.0.0 —
the CLI logs in fresh on each invocation, so per-token transactions cannot naturally span multiple shell commands. Use
the Go client directly for scripted atomic batches; a session-aware CLI mode may land in a later release.

## Things you don't need to do

- Existing scripts that hit `cobbler distro`, `cobbler profile`, `cobbler system`, `cobbler repo`, `cobbler image`,
  `cobbler menu` keep working, modulo the field removals listed above.
- `cobbler version` now shows the build date and git hash of the connected server (was already there in the
  cobblerclient — surfaced now).
