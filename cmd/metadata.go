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
}

var distroStringSliceFlagMetadata = map[string]FlagMetadata[[]string]{
	"boot-loaders": {
		Name:          "boot-loaders",
		DefaultValue:  []string{},
		Usage:         "boot loaders (network installation boot loaders)",
		IsInheritable: true,
	},
	"mgmt-classes": {
		Name:         "mgmt-classes",
		DefaultValue: []string{},
		Usage:        "management classes (for external config management)",
	},
}

var distroMapFlagMetadata = map[string]FlagMetadata[map[string]string]{
	"autoinstall-meta": {
		Name:          "autoinstall-meta",
		DefaultValue:  map[string]string{},
		Usage:         "automatic installation template metadata",
		IsInheritable: true,
	},
	"boot-files": {
		Name:          "boot-files",
		DefaultValue:  map[string]string{},
		Usage:         "TFTP boot files (files copied into tftpboot beyond the kernel/initrd)",
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
	"fetchable-files": {
		Name:          "fetchable-files",
		DefaultValue:  map[string]string{},
		Usage:         "fetchable files (templates for tftp, wget or curl)",
		IsInheritable: true,
	},
	"template-files": {
		Name:          "template-files",
		DefaultValue:  map[string]string{},
		Usage:         "template files (file mappings for built-in config management)",
		IsInheritable: true,
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
		Usage:        "the name of a previously defined Cobbler distribution. This value is required",
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
	"menu": {
		Name:         "menu",
		DefaultValue: "",
		Usage:        "parent boot menu",
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
		Name:          "name-servers-search",
		DefaultValue:  []string{},
		Usage:         "name servers search path (comma delimited)",
		IsInheritable: true,
	},
}

var profileMapFlagMetadata = map[string]FlagMetadata[map[string]string]{
	"mgmt-parameters": {
		Name:         "mgmt-parameters",
		DefaultValue: map[string]string{},
		Usage:        "Parameters which will be handed to your management application (must be a valid YAML dictionary))",
	},
}

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
	"menu": {
		Name:         "menu",
		DefaultValue: "",
		Usage:        "parent boot menu",
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
		Usage:        "parent image (if not a profile)",
	},
	"ipv6-default-device": {
		Name:         "ipv6-default-device",
		DefaultValue: "",
		Usage:        "IPv6 default device",
	},
	"profile": {
		Name:         "profile",
		DefaultValue: "",
		Usage:        "Parent profile",
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
	"enable-menu": {
		Name:          "enable-menu",
		DefaultValue:  false,
		Usage:         "enable PXE Menu? (show this profile in the PXE menu?)",
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
	"mgmt-classes": {
		Name:         "mgmt-classes",
		DefaultValue: []string{},
		Usage:        "management classes (for external config management)",
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
	"boot-files": {
		Name:          "boot-files",
		DefaultValue:  map[string]string{},
		Usage:         "TFTP boot files (files copied into tftpboot beyond the kernel/initrd)",
		IsInheritable: true,
	},
	"fetchable-files": {
		Name:          "fetchable-files",
		DefaultValue:  map[string]string{},
		Usage:         "fetchable files (templates for tftp, wget or curl)",
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
		Name:          "template-files",
		DefaultValue:  map[string]string{},
		Usage:         "template files (file mappings for built-in config management)",
		IsInheritable: true,
	},
	"mgmt-parameters": {
		Name:         "mgmt-parameters",
		DefaultValue: map[string]string{},
		Usage:        "Parameters which will be handed to your management application (must be a valid YAML dictionary))",
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

var interfaceStringFlagMetadata = map[string]FlagMetadata[string]{
	"bonding-opts": {
		Name:         "bonding-opts",
		DefaultValue: "",
		Usage:        "bonding opts (should be used with --interface)",
	},
	"bridge-opts": {
		Name:         "bridge-opts",
		DefaultValue: "",
		Usage:        "bridge opts (should be used with --interface)",
	},
	"dhcp-tag": {
		Name:         "dhcp-tag",
		DefaultValue: "",
		Usage:        "DHCP tag (see manpage or leave blank)",
	},
	"dns-name": {
		Name:         "dns-name",
		DefaultValue: "",
		Usage:        "DNS name (should be used with --interface)",
	},
	"if-gateway": {
		Name:         "if-gateway",
		DefaultValue: "",
		Usage:        "per-Interface Gateway (should be used with --interface)",
	},
	"interface-master": {
		Name:         "interface-master",
		DefaultValue: "",
		Usage:        "master interface (Should be used with --interface)",
	},
	"interface-type": {
		Name:         "interface-type",
		DefaultValue: "",
		Usage: `interface Type. Valid options: na,bond,bond_slave,bridge,bridge_slave,bonded_bridge_slave,bmc,infiniband.
	(should be used with --interface)`,
	},
	"ip-address": {
		Name:         "ip-address",
		DefaultValue: "",
		Usage:        "IPv4 address (should be used with --interface)",
	},
	"ipv6-address": {
		Name:         "ipv6-address",
		DefaultValue: "",
		Usage:        "IPv6 address (should be used with --interface)",
	},
	"ipv6-default-gateway": {
		Name:         "ipv6-default-gateway",
		DefaultValue: "",
		Usage:        "IPv6 Default Gateway (should be used with --interface)",
	},
	"ipv6-mtu": {
		Name:         "ipv6-mtu",
		DefaultValue: "",
		Usage:        "IPv6 MTU",
	},
	"ipv6-prefix": {
		Name:         "ipv6-prefix",
		DefaultValue: "",
		Usage:        "IPv6 Prefix (should be used with --interface)",
	},
	"mac-address": {
		Name:         "mac-address",
		DefaultValue: "",
		Usage:        "MAC Address (place 'random' in this field for a random MAC Address.)",
	},
	"mtu": {
		Name:         "mtu",
		DefaultValue: "",
		Usage:        "MTU (should be used with --interface)",
	},
	"netmask": {
		Name:         "netmask",
		DefaultValue: "",
		Usage:        "Subnet mask (should be used with --interface)",
	},
}

var interfaceBoolFlagMetadata = map[string]FlagMetadata[bool]{
	"connected-mode": {
		Name:         "connected-mode",
		DefaultValue: false,
		Usage:        "InfiniBand connected mode (should be used with --interface)",
	},
	"management": {
		Name:         "management",
		DefaultValue: false,
		Usage:        "declares the interface as management interface (should be used with --interface)",
	},
	"static": {
		Name:         "static",
		DefaultValue: false,
		Usage:        "Is this interface static? (should be used with --interface)",
	},
}

var interfaceStringSliceFlagMetadata = map[string]FlagMetadata[[]string]{
	"cnames": {
		Name:         "cnames",
		DefaultValue: []string{},
		Usage:        "Cannonical Name Records, should be used with --interface (comma delimited)",
	},
	"ipv6-secondaries": {
		Name:         "ipv6-secondaries",
		DefaultValue: []string{},
		Usage:        "IPv6 Secondaries (should be used with --interface)",
	},
	"ipv6-static-routes": {
		Name:         "ipv6-static-routes",
		DefaultValue: []string{},
		Usage:        "IPv6 Static Routes (should be used with --interface)",
	},
	"static-routes": {
		Name:         "static-routes",
		DefaultValue: []string{},
		Usage:        "static routes (should be used with --interface)",
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
		Usage:        "parent boot menu",
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

var fileStringFlagMetadata = map[string]FlagMetadata[string]{
	"action": {
		Name:         "action",
		DefaultValue: "",
		Usage:        "create or remove file resource",
	},
	"mode": {
		Name:         "mode",
		DefaultValue: "",
		Usage:        "file modes",
	},
	"template": {
		Name:         "template",
		DefaultValue: "",
		Usage:        "the template for the file",
	},
	"path": {
		Name:         "path",
		DefaultValue: "",
		Usage:        "the path of the file",
	},
	"group": {
		Name:         "group",
		DefaultValue: "",
		Usage:        "file owner group in file system",
	},
	"owner": {
		Name:         "owner",
		DefaultValue: "",
		Usage:        "file owner user in file system",
	},
}

var fileBoolFlagMetadata = map[string]FlagMetadata[bool]{
	"is-dir": {
		Name:         "is-dir",
		DefaultValue: false,
		Usage:        "treat file resource as a directory",
	},
}

var packageStringFlagMetadata = map[string]FlagMetadata[string]{
	"action": {
		Name:         "action",
		DefaultValue: "",
		Usage:        "install or remove package resource",
	},
	"installer": {
		Name:         "installer",
		DefaultValue: "",
		Usage:        "package manager",
	},
	"version": {
		Name:         "version",
		DefaultValue: "",
		Usage:        "package version",
	},
}

var mgmtclassStringFlagMetadata = map[string]FlagMetadata[string]{
	"params": {
		Name:         "params",
		DefaultValue: "",
		Usage:        "list of parameters/variables",
	},
	"class-name": {
		Name:         "class-name",
		DefaultValue: "",
		Usage:        "actual class name (leave blank to use the name field)",
	},
}

var mgmtclassBoolFlagMetadata = map[string]FlagMetadata[bool]{
	"is-definition": {
		Name:         "is-definition",
		DefaultValue: false,
		Usage:        "is Definition? Treat this class as a definition (puppet only)",
	},
}

var mgmtclassStringSliceFlagMetadata = map[string]FlagMetadata[[]string]{
	"files": {
		Name:         "files",
		DefaultValue: []string{},
		Usage:        "file resources",
	},
	"packages": {
		Name:         "packages",
		DefaultValue: []string{},
		Usage:        "package resources",
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
