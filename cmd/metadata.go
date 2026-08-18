package cmd

type FlagMetadata[T any] struct {
	Name          string
	DefaultValue  T
	Usage         string
	IsInheritable bool
}

var commonStringFlagMetadata = map[string]FlagMetadata[string]{
	"name": {
		Name:         "name",
		DefaultValue: "",
		Usage:        "the item name",
	},
	"comment": {
		Name:         "comment",
		DefaultValue: "",
		Usage:        "free form text description",
	},
}

var commonStringSliceFlagMetadata = map[string]FlagMetadata[[]string]{
	"owners": {
		Name:          "owners",
		DefaultValue:  []string{},
		Usage:         "owners list for authorization.ownership (comma delimited)",
		IsInheritable: true,
	},
}

var copyRenameStringFlagMetadata = map[string]FlagMetadata[string]{
	"newname": {
		Name:         "newname",
		DefaultValue: "",
		Usage:        "the new item name",
	},
}

var distroStringFlagMetadata = map[string]FlagMetadata[string]{
	"kernel": {
		Name:         "kernel",
		DefaultValue: "",
		Usage:        "Kernel (absolute path on filesystem)",
	},
	"initrd": {
		Name:         "initrd",
		DefaultValue: "",
		Usage:        "Initrd (absolute path on filesystem)",
	},
	"arch": {
		Name:         "arch",
		DefaultValue: "x86_64",
		Usage:        "Architecture",
	},
	"breed": {
		Name:         "breed",
		DefaultValue: "",
		Usage:        "Breed (what is the type of the distribution?)",
	},
	"os-version": {
		Name:         "os-version",
		DefaultValue: "",
		Usage:        "OS version (needed for some virtualization optimizations)",
	},
	"remote-boot-kernel": {
		Name:         "remote-boot-kernel",
		DefaultValue: "",
		Usage:        "remote boot kernel (URL the bootloader directly retrieves and boots from)",
	},
	"remote-boot-initrd": {
		Name:         "remote-boot-initrd",
		DefaultValue: "",
		Usage:        "remote boot initrd (URL the bootloader directly retrieves and boots from)",
	},
	"redhat-management-key": {
		Name:         "redhat-management-key",
		DefaultValue: "",
		Usage:        "RedHat management key (registration key for RHN, Spacewalk, or Satellite)",
	},
	"source-tree-path": {
		Name:         "source-tree-path",
		DefaultValue: "",
		Usage:        "the original location of the distro's source tree on disk (for use by the dynamic_httpd manager)",
	},
}

var distroStringSliceFlagMetadata = map[string]FlagMetadata[[]string]{
	"boot-loaders": {
		Name:          "boot-loaders",
		DefaultValue:  []string{},
		Usage:         "boot loaders (network installation boot loaders)",
		IsInheritable: true,
	},
}

var distroMapFlagMetadata = map[string]FlagMetadata[map[string]string]{
	"autoinstall-meta": {
		Name:          "autoinstall-meta",
		DefaultValue:  map[string]string{},
		Usage:         "automatic installation template metadata",
		IsInheritable: true,
	},
	"kernel-options": {
		Name:          "kernel-options",
		DefaultValue:  map[string]string{},
		Usage:         "kernel options (e.g. selinux=permissive)",
		IsInheritable: true,
	},
	"kernel-options-post": {
		Name:          "kernel-options-post",
		DefaultValue:  map[string]string{},
		Usage:         "post install kernel options (e.g. clocksource=pit noapic)",
		IsInheritable: true,
	},
	"template-files": {
		Name:         "template-files",
		DefaultValue: map[string]string{},
		Usage:        "template files (file mappings for built-in config management)",
	},
}

var profileStringFlagMetadata = map[string]FlagMetadata[string]{
	"autoinstall": {
		Name:         "autoinstall",
		DefaultValue: "",
		Usage:        "path to automatic installation template",
	},
	"distro": {
		Name:         "distro",
		DefaultValue: "",
		Usage:        "the UID of a previously defined Cobbler distribution (see 'distro report'). This value is required",
	},
	"redhat-management-key": {
		Name:         "redhat-management-key",
		DefaultValue: "",
		Usage:        "RedHat management key (registration key for RHN, Spacewalk, or Satellite)",
	},
	"dhcp-tag": {
		Name:         "dhcp-tag",
		DefaultValue: "",
		Usage:        "DHCP tag (see manpage or leave blank)",
	},
	"next-server-v4": {
		Name:         "next-server-v4",
		DefaultValue: "",
		Usage:        "next server (IPv4) override (see manpage or leave blank)",
	},
	"next-server-v6": {
		Name:         "next-server-v6",
		DefaultValue: "",
		Usage:        "next server (IPv6) override (see manpage or leave blank)",
	},
	"filename": {
		Name:         "filename",
		DefaultValue: "",
		Usage:        "DHCP filename override (used to boot non-default bootloaders)",
	},
	"parent": {
		Name:         "parent",
		DefaultValue: "",
		Usage:        "the UID of the parent profile (see 'profile report')",
	},
	"proxy": {
		Name:         "proxy",
		DefaultValue: "",
		Usage:        "proxy server URL",
	},
	"server": {
		Name:         "server",
		DefaultValue: "",
		Usage:        "server override",
	},
	"menu": {
		Name:         "menu",
		DefaultValue: "",
		Usage:        "the UID of the parent boot menu (see 'menu report')",
	},
	"virt-bridge": {
		Name:         "virt-bridge",
		DefaultValue: "",
		Usage:        "virt bridge",
	},
	"virt-disk-driver": {
		Name:         "virt-disk-driver",
		DefaultValue: "",
		Usage:        "the on-disk format for the virtualization disk. Valid options: <<inherit>>,raw,qcow2,qed,vdi,vdmk",
	},
	"virt-path": {
		Name:         "virt-path",
		DefaultValue: "",
		Usage:        "virt Path (e.g. /directory or VolGroup00)",
	},
	"virt-type": {
		Name:         "virt-type",
		DefaultValue: "",
		Usage:        "virtualization technology to use. Valid options: <<inherit>>,qemu,kvm,xenpv,xenfv,vmware,vmwarew,openvz,auto)",
	},
}

var profileBoolFlagMetadata = map[string]FlagMetadata[bool]{
	"enable-ipxe": {
		Name:         "enable-ipxe",
		DefaultValue: false,
		Usage:        "enable iPXE? (use iPXE instead of PXELINUX for advanced booting options)",
	},
	"enable-menu": {
		Name:         "enable-menu",
		DefaultValue: false,
		Usage:        "enable PXE Menu? (show this profile in the PXE menu?)",
	},
	"virt-auto-boot": {
		Name:         "virt-auto-boot",
		DefaultValue: false,
		Usage:        "auto boot this VM?",
	},
	"virt-uefi": {
		Name:         "virt-uefi",
		DefaultValue: false,
		Usage:        "boot this VM via UEFI firmware instead of legacy BIOS?",
	},
}

var profileIntFlagMetadata = map[string]FlagMetadata[int]{
	"virt-cpus": {
		Name:         "virt-cpus",
		DefaultValue: 0,
		Usage:        "virt CPUs",
	},
	"virt-ram": {
		Name:         "virt-ram",
		DefaultValue: 0,
		Usage:        "virt RAM size in MB",
	},
}

var profileFloatFlagMetadata = map[string]FlagMetadata[float64]{
	"virt-file-size": {
		Name:         "virt-file-size",
		DefaultValue: float64(0),
		Usage:        "virt file size in GB",
	},
}

var profileStringSliceFlagMetadata = map[string]FlagMetadata[[]string]{
	"repos": {
		Name:         "repos",
		DefaultValue: []string{},
		Usage:        "repos to auto-assign to this profile",
	},
	"name-servers": {
		Name:          "name-servers",
		DefaultValue:  []string{},
		Usage:         "name servers (comma delimited)",
		IsInheritable: true,
	},
	"name-servers-search": {
		Name:         "name-servers-search",
		DefaultValue: []string{},
		Usage:        "name servers search path (comma delimited)",
	},
}

var profileMapFlagMetadata = map[string]FlagMetadata[map[string]string]{}

var systemStringFlagMetadata = map[string]FlagMetadata[string]{
	"autoinstall": {
		Name:         "autoinstall",
		DefaultValue: "",
		Usage:        "path to automatic installation template",
	},
	"redhat-management-key": {
		Name:         "redhat-management-key",
		DefaultValue: "<<inherit>>",
		Usage:        "RedHat management key (registration key for RHN, Spacewalk, or Satellite)",
	},
	"next-server-v4": {
		Name:         "next-server-v4",
		DefaultValue: "",
		Usage:        "next server (IPv4) override (see manpage or leave blank)",
	},
	"next-server-v6": {
		Name:         "next-server-v6",
		DefaultValue: "",
		Usage:        "next server (IPv6) override (see manpage or leave blank)",
	},
	"filename": {
		Name:         "filename",
		DefaultValue: "",
		Usage:        "DHCP filename override (used to boot non-default bootloaders)",
	},
	"parent": {
		Name:         "parent",
		DefaultValue: "",
		Usage:        "parent profile",
	},
	"proxy": {
		Name:         "proxy",
		DefaultValue: "",
		Usage:        "proxy server URL",
	},
	"server": {
		Name:         "server",
		DefaultValue: "",
		Usage:        "server override",
	},
	"virt-bridge": {
		Name:         "virt-bridge",
		DefaultValue: "",
		Usage:        "virt bridge",
	},
	"virt-disk-driver": {
		Name:         "virt-disk-driver",
		DefaultValue: "",
		Usage:        "the on-disk format for the virtualization disk. Valid options: <<inherit>>,raw,qcow2,qed,vdi,vdmk",
	},
	"virt-path": {
		Name:         "virt-path",
		DefaultValue: "",
		Usage:        "virt Path (e.g. /directory or VolGroup00)",
	},
	"virt-type": {
		Name:         "virt-type",
		DefaultValue: "",
		Usage:        "virtualization technology to use. Valid options: <<inherit>>,qemu,kvm,xenpv,xenfv,vmware,vmwarew,openvz,auto)",
	},
	"gateway": {
		Name:         "gateway",
		DefaultValue: "",
		Usage:        "gateway",
	},
	"hostname": {
		Name:         "hostname",
		DefaultValue: "",
		Usage:        "hostname",
	},
	"image": {
		Name:         "image",
		DefaultValue: "",
		Usage:        "the UID of the parent image (if not a profile; see 'image report')",
	},
	"ipv6-default-device": {
		Name:         "ipv6-default-device",
		DefaultValue: "",
		Usage:        "IPv6 default device",
	},
	"profile": {
		Name:         "profile",
		DefaultValue: "",
		Usage:        "the UID of the parent profile (see 'profile report')",
	},
	"status": {
		Name:         "status",
		DefaultValue: "",
		Usage:        "system status. Valid options: development,testing,acceptance,production",
	},
}

var systemBoolFlagMetadata = map[string]FlagMetadata[bool]{
	"enable-ipxe": {
		Name:          "enable-ipxe",
		DefaultValue:  false,
		Usage:         "enable iPXE? (use iPXE instead of PXELINUX for advanced booting options)",
		IsInheritable: true,
	},
	"virt-auto-boot": {
		Name:          "virt-auto-boot",
		DefaultValue:  false,
		Usage:         "auto boot this VM?",
		IsInheritable: true,
	},
	"netboot-enabled": {
		Name:         "netboot-enabled",
		DefaultValue: false,
		Usage:        "PXE (re)install this machine at next boot?",
	},
	"virt-pxe-boot": {
		Name:         "virt-pxe-boot",
		DefaultValue: false,
		Usage:        "use PXE to build this VM?",
	},
	"virt-uefi": {
		Name:         "virt-uefi",
		DefaultValue: false,
		Usage:        "boot this VM via UEFI firmware instead of legacy BIOS?",
	},
}

var systemIntFlagMetadata = map[string]FlagMetadata[int]{
	"virt-cpus": {
		Name:         "virt-cpus",
		DefaultValue: 0,
		Usage:        "virt CPUs",
	},
	"virt-ram": {
		Name:         "virt-ram",
		DefaultValue: 0,
		Usage:        "virt RAM size in MB",
	},
	"serial-device": {
		Name:         "serial-device",
		DefaultValue: 0,
		Usage:        "serial device number",
	},
	"serial-baud-rate": {
		Name:         "serial-baud-rate",
		DefaultValue: 0,
		Usage:        "serial Baud Rate. Valid options: 2400,4800,9600,19200,38400,57600,115200",
	},
}

var systemFloatFlagMetadata = map[string]FlagMetadata[float64]{
	"virt-file-size": {
		Name:         "virt-file-size",
		DefaultValue: float64(0),
		Usage:        "virt file size in GB",
	},
}

var systemStringSliceFlagMetadata = map[string]FlagMetadata[[]string]{
	"boot-loaders": {
		Name:          "boot-loaders",
		DefaultValue:  []string{},
		Usage:         "boot loaders (network installation boot loaders)",
		IsInheritable: true,
	},
	"name-servers": {
		Name:         "name-servers",
		DefaultValue: []string{},
		Usage:        "name servers (comma delimited)",
	},
	"name-servers-search": {
		Name:         "name-servers-search",
		DefaultValue: []string{},
		Usage:        "name servers search path (comma delimited)",
	},
}

var systemMapFlagMetadata = map[string]FlagMetadata[map[string]string]{
	"autoinstall-meta": {
		Name:          "autoinstall-meta",
		DefaultValue:  map[string]string{},
		Usage:         "automatic installation template metadata",
		IsInheritable: true,
	},
	"kernel-options": {
		Name:          "kernel-options",
		DefaultValue:  map[string]string{},
		Usage:         "kernel options (e.g. selinux=permissive)",
		IsInheritable: true,
	},
	"kernel-options-post": {
		Name:          "kernel-options-post",
		DefaultValue:  map[string]string{},
		Usage:         "post install kernel options (e.g. clocksource=pit noapic)",
		IsInheritable: true,
	},
	"template-files": {
		Name:         "template-files",
		DefaultValue: map[string]string{},
		Usage:        "template files (file mappings for built-in config management)",
	},
}

var systemPowerStringFlagMetadata = map[string]FlagMetadata[string]{
	"power-address": {
		Name:         "power-address",
		DefaultValue: "",
		Usage:        "power management address (e.g. power-device.example.org)",
	},
	"power-id": {
		Name:         "power-id",
		DefaultValue: "",
		Usage:        "power management ID (usually a plug number or blade name, if power type requires it)",
	},
	"power-pass": {
		Name:         "power-pass",
		DefaultValue: "",
		Usage:        "power management password",
	},
	"power-type": {
		Name:         "power-type",
		DefaultValue: "",
		Usage:        "power management script to use",
	},
	"power-user": {
		Name:         "power-user",
		DefaultValue: "",
		Usage:        "power management username",
	},
	"power-options": {
		Name:         "power-options",
		DefaultValue: "",
		Usage:        "additional options, to be passed to the fencing agent",
	},
	"power-identity-file": {
		Name:         "power-identity-file",
		DefaultValue: "",
		Usage:        "identity file to be passed to the fencing agent (SSH key)",
	},
}

// interfaceStringFlagMetadata holds the string flags for the dedicated
// `cobbler interface` command. In Cobbler 4.0.0 the IPv4/IPv6/DNS configuration
// lives in nested value objects on NetworkInterface; the flag names are
// dot-separated by convention (--ipv4-address sets IPv4.Address, etc.).
var interfaceStringFlagMetadata = map[string]FlagMetadata[string]{
	"bonding-opts": {
		Name:         "bonding-opts",
		DefaultValue: "",
		Usage:        "bonding opts",
	},
	"bridge-opts": {
		Name:         "bridge-opts",
		DefaultValue: "",
		Usage:        "bridge opts",
	},
	"dhcp-tag": {
		Name:         "dhcp-tag",
		DefaultValue: "",
		Usage:        "DHCP tag (see manpage or leave blank)",
	},
	"dns-name": {
		Name:         "dns-name",
		DefaultValue: "",
		Usage:        "DNS name",
	},
	"interface-master": {
		Name:         "interface-master",
		DefaultValue: "",
		Usage:        "master interface",
	},
	"interface-type": {
		Name:         "interface-type",
		DefaultValue: "na",
		Usage:        "interface type (na,bond,bond_slave,bridge,bridge_slave,bonded_bridge_slave,infiniband)",
	},
	"ipv4-address": {
		Name:         "ipv4-address",
		DefaultValue: "",
		Usage:        "IPv4 address",
	},
	"ipv4-netmask": {
		Name:         "ipv4-netmask",
		DefaultValue: "",
		Usage:        "IPv4 subnet mask",
	},
	"ipv4-gateway": {
		Name:         "ipv4-gateway",
		DefaultValue: "",
		Usage:        "per-interface IPv4 gateway",
	},
	"ipv6-address": {
		Name:         "ipv6-address",
		DefaultValue: "",
		Usage:        "IPv6 address",
	},
	"ipv6-prefix": {
		Name:         "ipv6-prefix",
		DefaultValue: "",
		Usage:        "IPv6 prefix",
	},
	"ipv6-mtu": {
		Name:         "ipv6-mtu",
		DefaultValue: "",
		Usage:        "IPv6 MTU",
	},
	"ipv6-default-gateway": {
		Name:         "ipv6-default-gateway",
		DefaultValue: "",
		Usage:        "IPv6 default gateway",
	},
	"mac-address": {
		Name:         "mac-address",
		DefaultValue: "",
		Usage:        "MAC address (use 'random' for a random MAC address)",
	},
	"mtu": {
		Name:         "mtu",
		DefaultValue: "",
		Usage:        "MTU",
	},
	"virt-bridge": {
		Name:          "virt-bridge",
		DefaultValue:  "",
		Usage:         "virt bridge",
		IsInheritable: true,
	},
}

var interfaceBoolFlagMetadata = map[string]FlagMetadata[bool]{
	"connected-mode": {
		Name:         "connected-mode",
		DefaultValue: false,
		Usage:        "InfiniBand connected mode",
	},
	"management": {
		Name:         "management",
		DefaultValue: false,
		Usage:        "declares the interface as a management interface",
	},
	"static": {
		Name:         "static",
		DefaultValue: false,
		Usage:        "is this interface static?",
	},
}

var interfaceStringSliceFlagMetadata = map[string]FlagMetadata[[]string]{
	"dns-cnames": {
		Name:         "dns-cnames",
		DefaultValue: []string{},
		Usage:        "canonical name records (comma delimited)",
	},
	"ipv4-static-routes": {
		Name:         "ipv4-static-routes",
		DefaultValue: []string{},
		Usage:        "IPv4 static routes (comma delimited)",
	},
	"ipv6-secondaries": {
		Name:         "ipv6-secondaries",
		DefaultValue: []string{},
		Usage:        "IPv6 secondary addresses (comma delimited)",
	},
	"ipv6-static-routes": {
		Name:         "ipv6-static-routes",
		DefaultValue: []string{},
		Usage:        "IPv6 static routes (comma delimited)",
	},
}

var imageStringFlagMetadata = map[string]FlagMetadata[string]{
	"arch": {
		Name:         "arch",
		DefaultValue: "",
		Usage:        "Architecture",
	},
	"breed": {
		Name:         "breed",
		DefaultValue: "",
		Usage:        "Breed (valid options: none,rsync,rhn,yum,apt,wget)",
	},
	"parent": {
		Name:         "parent",
		DefaultValue: "",
		Usage:        "parent item",
	},
	"file": {
		Name:         "file",
		DefaultValue: "",
		Usage:        "path to local file or nfs://user@host:path",
	},
	"image-type": {
		Name:         "image-type",
		DefaultValue: "",
		Usage:        "image type. Valid options: iso,direct,memdisk,virt-image",
	},
	"os-version": {
		Name:         "os-version",
		DefaultValue: "",
		Usage:        "OS version (needed for some virtualization optimizations)",
	},
	"menu": {
		Name:         "menu",
		DefaultValue: "",
		Usage:        "the UID of the parent boot menu (see 'menu report')",
	},
	"virt-bridge": {
		Name:         "virt-bridge",
		DefaultValue: "",
		Usage:        "virt bridge",
	},
	"virt-disk-driver": {
		Name:         "virt-disk-driver",
		DefaultValue: "<<inherit>>",
		Usage:        "the on-disk format for the virtualization disk. Valid options: <<inherit>>,raw,qcow2,qed,vdi,vdmk",
	},
	"virt-path": {
		Name:         "virt-path",
		DefaultValue: "",
		Usage:        "virt Path (e.g. /directory or VolGroup00)",
	},
	"virt-type": {
		Name:         "virt-type",
		DefaultValue: "",
		Usage:        "virtualization technology to use. Valid options: xenpv,xenfv,qemu,kvm,vmware",
	},
}

var imageIntFlagMetadata = map[string]FlagMetadata[int]{
	"network-count": {
		Name:         "network-count",
		DefaultValue: 0,
		Usage:        "Network Count",
	},
	"virt-cpus": {
		Name:         "virt-cpus",
		DefaultValue: 1,
		Usage:        "virt CPUs",
	},
	"virt-ram": {
		Name:         "virt-ram",
		DefaultValue: 0,
		Usage:        "virt RAM size in MB",
	},
}

var imageFloatFlagMetadata = map[string]FlagMetadata[float64]{
	"virt-file-size": {
		Name:         "virt-file-size",
		DefaultValue: float64(0),
		Usage:        "virt file size in GB",
	},
}

var imageBoolFlagMetadata = map[string]FlagMetadata[bool]{
	"virt-auto-boot": {
		Name:         "virt-auto-boot",
		DefaultValue: false,
		Usage:        "auto boot this VM?",
	},
	"virt-uefi": {
		Name:         "virt-uefi",
		DefaultValue: false,
		Usage:        "boot this VM via UEFI firmware instead of legacy BIOS?",
	},
}

var imageStringSliceFlagMetadata = map[string]FlagMetadata[[]string]{
	"boot-loaders": {
		Name:         "boot-loaders",
		DefaultValue: []string{},
		Usage:        "boot loaders (network installation boot loaders)",
	},
}

var menuStringFlagMetadata = map[string]FlagMetadata[string]{
	"parent": {
		Name:         "parent",
		DefaultValue: "",
		Usage:        "parent menu",
	},
	"display-name": {
		Name:         "display-name",
		DefaultValue: "",
		Usage:        "display name",
	},
}

var repoStringFlagMetadata = map[string]FlagMetadata[string]{
	"arch": {
		Name:         "arch",
		DefaultValue: "none",
		Usage:        "Architecture",
	},
	"breed": {
		Name:         "breed",
		DefaultValue: "none",
		Usage:        "Breed (valid options: none,rsync,rhn,yum,apt,wget)",
	},
	"createrepo-flags": {
		Name:         "createrepo-flags",
		DefaultValue: "",
		Usage:        "flags to use with createrepo",
	},
	"mirror": {
		Name:         "mirror",
		DefaultValue: "",
		Usage:        "address of yum or rsync repo to mirror",
	},
	"mirror-type": {
		Name:         "mirror-type",
		DefaultValue: "",
		Usage:        "mirror type. Valid options: metalink,mirrorlist,baseurl",
	},
	"proxy": {
		Name:         "proxy",
		DefaultValue: "",
		Usage:        "proxy URL (<<inherit>> to use proxy_url_ext from settings, blank or <<None>> for no proxy)",
	},
}

var repoBoolFlagMetadata = map[string]FlagMetadata[bool]{
	"keep-updated": {
		Name:         "keep-updated",
		DefaultValue: false,
		Usage:        "update this repo on next 'cobbler reposync'?",
	},
}

var repoIntFlagMetadata = map[string]FlagMetadata[int]{
	"priority": {
		Name:         "priority",
		DefaultValue: 0,
		Usage:        "value for yum priorities plugin, if installed",
	},
}

var repoStringSliceFlagMetadata = map[string]FlagMetadata[[]string]{
	"apt-components": {
		Name:         "apt-components",
		DefaultValue: []string{},
		Usage:        "APT components (e.g. main restricted universe)",
	},
	"apt-dists": {
		Name:         "apt-dists",
		DefaultValue: []string{},
		Usage:        "APT dist names (e.g. precise,bullseye,buster)",
	},
}

var repoMapFlagMetadata = map[string]FlagMetadata[map[string]string]{
	"environment": {
		Name:         "environment",
		DefaultValue: map[string]string{},
		Usage:        "environment variables (use these environment variables during commands (key=value, comma delimited)",
	},
	"yumopts": {
		Name:         "yumopts",
		DefaultValue: map[string]string{},
		Usage:        "options to write to yum config file",
	},
	"rsyncopts": {
		Name:         "rsyncopts",
		DefaultValue: map[string]string{},
		Usage:        "options to use with rsync repo",
	},
	"rpm-list": {
		Name:         "rpm-list",
		DefaultValue: map[string]string{},
		Usage:        "mirror just these RPMs (yum only)",
	},
}

var findStringFlagMetadata = map[string]FlagMetadata[string]{
	"uid": {
		Name:         "uid",
		DefaultValue: "",
		Usage:        "",
	},
}

var findIntFlagMetadata = map[string]FlagMetadata[int]{
	"depth": {
		Name:         "depth",
		DefaultValue: 0,
		Usage:        "",
	},
}

var findFloatFlagMetadata = map[string]FlagMetadata[float64]{
	"ctime": {
		Name:         "ctime",
		DefaultValue: 0.0,
		Usage:        "",
	},
	"mtime": {
		Name:         "mtime",
		DefaultValue: 0.0,
		Usage:        "",
	},
}

var exportStringMetadata = map[string]FlagMetadata[string]{
	"format": {
		Name:         "format",
		DefaultValue: "json",
		Usage:        `the export format, must be one of "JSON" or "YAML"`,
	},
}
