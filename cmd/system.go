// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
	"reflect"

	cobbler "github.com/cobbler/cobblerclient"
)

func updateSystemFromFlags(cmd *cobra.Command, system *cobbler.System) error {
	// TODO: Implementation for more interfaces
	// See https://github.com/cobbler/cli/issues/38
	systemNewInterface, err := cmd.Flags().GetString("interface")
	if err != nil {
		return err
	}
	systemInterface, keyInMap := system.Interfaces[systemNewInterface]
	if !keyInMap {
		// Interface doesn't exist, so add a new one.
		// We cannot call CreateInterface because the system might not exist.
		system.Interfaces[systemNewInterface] = cobbler.Interface{}
		systemInterface = system.Interfaces[systemNewInterface]
	}
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		if err != nil {
			// If one of the previous flags has had an error just directly return.
			return
		}
		switch flag.Name {
		// The rename & copy operations are special operations as such we cannot blindly set this inside here.
		// Any rename & copy operation must be handled outside of this method.
		case "autoinstall":
			var systemNewAutoinstall string
			systemNewAutoinstall, err = cmd.Flags().GetString("autoinstall")
			if err != nil {
				return
			}
			system.Autoinstall = systemNewAutoinstall
		case "autoinstall-meta":
			fallthrough
		case "autoinstall-meta-inherit":
			if cmd.Flags().Lookup("autoinstall-meta-inherit").Changed {
				system.AutoinstallMeta.Data = make(map[string]interface{})
				system.AutoinstallMeta.IsInherited, err = cmd.Flags().GetBool("autoinstall-meta-inherit")
				if err != nil {
					return
				}
			} else {
				var systemNewAutoinstallMeta map[string]string
				systemNewAutoinstallMeta, err = cmd.Flags().GetStringToString("autoinstall-meta")
				if err != nil {
					return
				}
				system.AutoinstallMeta.IsInherited = false
				system.AutoinstallMeta.Data = convertMapStringToMapInterface(systemNewAutoinstallMeta)
			}
		case "boot-files":
			fallthrough
		case "boot-files-inherit":
			if cmd.Flags().Lookup("boot-files-inherit").Changed {
				system.BootFiles.Data = make(map[string]interface{})
				system.BootFiles.IsInherited, err = cmd.Flags().GetBool("boot-files-inherit")
				if err != nil {
					return
				}
			} else {
				var systemNewBootFiles map[string]string
				systemNewBootFiles, err = cmd.Flags().GetStringToString("boot-files")
				if err != nil {
					return
				}
				system.BootFiles.IsInherited = false
				system.BootFiles.Data = convertMapStringToMapInterface(systemNewBootFiles)
			}
		case "boot-loaders":
			fallthrough
		case "boot-loaders-inherit":
			if cmd.Flags().Lookup("boot-loaders-inherit").Changed {
				system.BootLoaders.Data = []string{}
				system.BootLoaders.IsInherited, err = cmd.Flags().GetBool("boot-loaders-inherit")
				if err != nil {
					return
				}
			} else {
				var systemNewBootLoaders []string
				systemNewBootLoaders, err = cmd.Flags().GetStringSlice("boot-loaders")
				if err != nil {
					return
				}
				system.BootLoaders.IsInherited = false
				system.BootLoaders.Data = systemNewBootLoaders
			}
		case "comment":
			var systemNewComment string
			systemNewComment, err = cmd.Flags().GetString("comment")
			if err != nil {
				return
			}
			system.Comment = systemNewComment
		case "fetchable-files":
			fallthrough
		case "fetchable-files-inherit":
			if cmd.Flags().Lookup("fetchable-files-inherit").Changed {
				system.FetchableFiles.Data = make(map[string]interface{})
				system.FetchableFiles.IsInherited, err = cmd.Flags().GetBool("fetchable-files-inherit")
				if err != nil {
					return
				}
			} else {
				var systemNewFetchableFiles map[string]string
				systemNewFetchableFiles, err = cmd.Flags().GetStringToString("fetchable-files")
				if err != nil {
					return
				}
				system.FetchableFiles.IsInherited = false
				system.FetchableFiles.Data = convertMapStringToMapInterface(systemNewFetchableFiles)
			}
		case "kernel-options":
			fallthrough
		case "kernel-options-inherit":
			if cmd.Flags().Lookup("kernel-options-inherit").Changed {
				system.KernelOptions.Data = make(map[string]interface{})
				system.KernelOptions.IsInherited, err = cmd.Flags().GetBool("kernel-options-inherit")
				if err != nil {
					return
				}
			} else {
				var systemNewKernelOptions map[string]string
				systemNewKernelOptions, err = cmd.Flags().GetStringToString("kernel-options")
				if err != nil {
					return
				}
				system.KernelOptions.IsInherited = false
				system.KernelOptions.Data = convertMapStringToMapInterface(systemNewKernelOptions)
			}
		case "kernel-options-post":
			fallthrough
		case "kernel-options-post-inherit":
			if cmd.Flags().Lookup("kernel-options-post-inherit").Changed {
				system.KernelOptionsPost.Data = make(map[string]interface{})
				system.KernelOptionsPost.IsInherited, err = cmd.Flags().GetBool("kernel-options-post-inherit")
				if err != nil {
					return
				}
			} else {
				var systemNewKernelOptionsPost map[string]string
				systemNewKernelOptionsPost, err = cmd.Flags().GetStringToString("kernel-options-post")
				if err != nil {
					return
				}
				system.KernelOptionsPost.IsInherited = false
				system.KernelOptionsPost.Data = convertMapStringToMapInterface(systemNewKernelOptionsPost)
			}
		case "mgmt-classes":
			fallthrough
		case "mgmt-classes-inherit":
			if cmd.Flags().Lookup("mgmt-classes-inherit").Changed {
				system.MgmtClasses.Data = []string{}
				system.MgmtClasses.IsInherited, err = cmd.Flags().GetBool("mgmt-classes-inherit")
				if err != nil {
					return
				}
			} else {
				var systemNewMgmtClasses []string
				systemNewMgmtClasses, err = cmd.Flags().GetStringSlice("mgmt-classes")
				if err != nil {
					return
				}
				system.MgmtClasses.IsInherited = false
				system.MgmtClasses.Data = systemNewMgmtClasses
			}
		case "owners":
			fallthrough
		case "owners-inherit":
			if cmd.Flags().Lookup("owners-inherit").Changed {
				system.Owners.Data = []string{}
				system.Owners.IsInherited, err = cmd.Flags().GetBool("owners-inherit")
				if err != nil {
					return
				}
			} else {
				var systemNewOwners []string
				systemNewOwners, err = cmd.Flags().GetStringSlice("owners")
				if err != nil {
					return
				}
				system.Owners.IsInherited = false
				system.Owners.Data = systemNewOwners
			}
		case "redhat-management-key":
			var systemNewRedhatManagementKey string
			systemNewRedhatManagementKey, err = cmd.Flags().GetString("redhat-management-key")
			if err != nil {
				return
			}
			system.RedhatManagementKey = systemNewRedhatManagementKey
		case "template-files-post":
			fallthrough
		case "template-files-inherit":
			if cmd.Flags().Lookup("template-files-inherit").Changed {
				system.TemplateFiles.Data = make(map[string]interface{})
				system.TemplateFiles.IsInherited, err = cmd.Flags().GetBool("template-files-inherit")
				if err != nil {
					return
				}
			} else {
				var systemNewTemplateFiles map[string]string
				systemNewTemplateFiles, err = cmd.Flags().GetStringToString("template-files")
				if err != nil {
					return
				}
				system.TemplateFiles.IsInherited = false
				system.TemplateFiles.Data = convertMapStringToMapInterface(systemNewTemplateFiles)
			}
		case "dhcp-tag":
			var systemNewDhcpTag string
			systemNewDhcpTag, err = cmd.Flags().GetString("dhcp-tag")
			if err != nil {
				return
			}
			systemInterface.DHCPTag = systemNewDhcpTag
		case "enable-ipxe":
			fallthrough
		case "enable-ipxe-inherit":
			if cmd.Flags().Lookup("enable-ipxe-inherit").Changed {
				system.EnableIPXE.Data = false
				system.EnableIPXE.IsInherited, err = cmd.Flags().GetBool("enable-ipxe-inherit")
				if err != nil {
					return
				}
			} else {
				var systemNewEnableIpxe bool
				systemNewEnableIpxe, err = cmd.Flags().GetBool("enable-ipxe")
				if err != nil {
					return
				}
				system.EnableIPXE.IsInherited = false
				system.EnableIPXE.Data = systemNewEnableIpxe
			}
		case "enable-menu":
			fallthrough
		case "enable-menu-inherit":
			if cmd.Flags().Lookup("enable-menu-inherit").Changed {
				system.EnableMenu.Data = false
				system.EnableMenu.IsInherited, err = cmd.Flags().GetBool("enable-menu-inherit")
				if err != nil {
					return
				}
			} else {
				var systemNewEnableMenu bool
				systemNewEnableMenu, err = cmd.Flags().GetBool("enable-menu")
				if err != nil {
					return
				}
				system.EnableMenu.IsInherited = false
				system.EnableMenu.Data = systemNewEnableMenu
			}
		case "mgmt-parameters":
			fallthrough
		case "mgmt-parameters-inherit":
			if cmd.Flags().Lookup("mgmt-parameters-inherit").Changed {
				system.MgmtParameters.Data = make(map[string]interface{})
				system.MgmtParameters.IsInherited, err = cmd.Flags().GetBool("mgmt-parameters-inherit")
				if err != nil {
					return
				}
			} else {
				var systemNewMgmtParameters map[string]string
				systemNewMgmtParameters, err = cmd.Flags().GetStringToString("mgmt-parameters")
				if err != nil {
					return
				}
				system.MgmtParameters.IsInherited = false
				system.MgmtParameters.Data = convertMapStringToMapInterface(systemNewMgmtParameters)
			}
		case "name-servers":
			var systemNewNameServers []string
			systemNewNameServers, err = cmd.Flags().GetStringSlice("name-servers")
			if err != nil {
				return
			}
			system.NameServers = systemNewNameServers
		case "name-servers-search":
			var systemNewNameServersSearch []string
			systemNewNameServersSearch, err = cmd.Flags().GetStringSlice("name-servers-search")
			if err != nil {
				return
			}
			system.NameServersSearch = systemNewNameServersSearch
		case "next-server-v4":
			var systemNewNextServerV4 string
			systemNewNextServerV4, err = cmd.Flags().GetString("next-server-v4")
			if err != nil {
				return
			}
			system.NextServerv4 = systemNewNextServerV4
		case "next-server-v6":
			var systemNewNextServerV6 string
			systemNewNextServerV6, err = cmd.Flags().GetString("next-server-v6")
			if err != nil {
				return
			}
			system.NextServerv6 = systemNewNextServerV6
		case "filename":
			var systemNewFilename string
			systemNewFilename, err = cmd.Flags().GetString("filename")
			if err != nil {
				return
			}
			system.Filename = systemNewFilename
		case "parent":
			var systemNewParent string
			systemNewParent, err = cmd.Flags().GetString("parent")
			if err != nil {
				return
			}
			system.Parent = systemNewParent
		case "proxy":
			var systemNewProxy string
			systemNewProxy, err = cmd.Flags().GetString("proxy")
			if err != nil {
				return
			}
			system.Proxy = systemNewProxy
		case "server":
			var systemNewServer string
			systemNewServer, err = cmd.Flags().GetString("server")
			if err != nil {
				return
			}
			system.Server = systemNewServer
		case "menu":
			var systemNewMenu string
			systemNewMenu, err = cmd.Flags().GetString("menu")
			if err != nil {
				return
			}
			system.Menu = systemNewMenu
		case "virt-auto-boot":
			fallthrough
		case "virt-auto-boot-inherit":
			if cmd.Flags().Lookup("virt-auto-boot-inherit").Changed {
				system.VirtAutoBoot.Data = false
				system.VirtAutoBoot.IsInherited = true
			} else {
				var systemNewVirtAutoBoot bool
				systemNewVirtAutoBoot, err = cmd.Flags().GetBool("virt-auto-boot")
				if err != nil {
					return
				}
				system.VirtAutoBoot.Data = systemNewVirtAutoBoot
				system.VirtAutoBoot.IsInherited = false
			}
		case "virt-cpus":
			fallthrough
		case "virt-cpus-inherit":
			if cmd.Flags().Lookup("virt-cpus-inherit").Changed {
				system.VirtCPUs.IsInherited = true
			} else {
				var systemNewVirtCpus int
				systemNewVirtCpus, err = cmd.Flags().GetInt("virt-cpus")
				if err != nil {
					return
				}
				system.VirtCPUs.Data = systemNewVirtCpus
				system.VirtCPUs.IsInherited = false
			}
		case "virt-disk-driver":
			var systemNewVirtDiskDriver string
			systemNewVirtDiskDriver, err = cmd.Flags().GetString("virt-disk-driver")
			if err != nil {
				return
			}
			system.VirtDiskDriver = systemNewVirtDiskDriver
		case "virt-file-size":
			fallthrough
		case "virt-file-size-inherit":
			if cmd.Flags().Lookup("virt-file-size-inherit").Changed {
				system.VirtFileSize.IsInherited = true
			} else {
				var systemNewVirtFileSize float64
				systemNewVirtFileSize, err = cmd.Flags().GetFloat64("virt-file-size")
				if err != nil {
					return
				}
				system.VirtFileSize.Data = systemNewVirtFileSize
				system.VirtFileSize.IsInherited = false
			}
		case "virt-path":
			var systemNewVirtPath string
			systemNewVirtPath, err = cmd.Flags().GetString("virt-path")
			if err != nil {
				return
			}
			system.VirtPath = systemNewVirtPath
		case "virt-ram":
			fallthrough
		case "virt-ram-inherit":
			if cmd.Flags().Lookup("virt-ram-inherit").Changed {
				system.VirtRAM.IsInherited = true
			} else {
				var systemNewVirtRam int
				systemNewVirtRam, err = cmd.Flags().GetInt("virt-ram")
				if err != nil {
					return
				}
				system.VirtRAM.Data = systemNewVirtRam
				system.VirtRAM.IsInherited = false
			}
		case "virt-type":
			var systemNewVirtType string
			systemNewVirtType, err = cmd.Flags().GetString("virt-type")
			if err != nil {
				return
			}
			system.VirtType = systemNewVirtType
		case "gateway":
			var systemNewGateway string
			systemNewGateway, err = cmd.Flags().GetString("gateway")
			if err != nil {
				return
			}
			system.Gateway = systemNewGateway
		case "hostname":
			var systemNewHostname string
			systemNewHostname, err := cmd.Flags().GetString("hostname")
			if err != nil {
				return
			}
			system.Hostname = systemNewHostname
		case "image":
			var systemNewImage string
			systemNewImage, err = cmd.Flags().GetString("image")
			if err != nil {
				return
			}
			system.Image = systemNewImage
		case "ipv6-default-device":
			var systemNewIpv6DefaultDevice string
			systemNewIpv6DefaultDevice, err = cmd.Flags().GetString("ipv6-default-device")
			if err != nil {
				return
			}
			system.IPv6DefaultDevice = systemNewIpv6DefaultDevice
		case "netboot-enabled":
			var systemNewNetbootEnabled bool
			systemNewNetbootEnabled, err = cmd.Flags().GetBool("netboot-enabled")
			if err != nil {
				return
			}
			system.NetbootEnabled = systemNewNetbootEnabled
		case "power-address":
			var systemNewPowerAddress string
			systemNewPowerAddress, err = cmd.Flags().GetString("power-address")
			if err != nil {
				return
			}
			system.PowerAddress = systemNewPowerAddress
		case "power-id":
			var systemNewPowerId string
			systemNewPowerId, err = cmd.Flags().GetString("power-id")
			if err != nil {
				return
			}
			system.PowerID = systemNewPowerId
		case "power-pass":
			var systemNewPowerPass string
			systemNewPowerPass, err = cmd.Flags().GetString("power-pass")
			if err != nil {
				return
			}
			system.PowerPass = systemNewPowerPass
		case "power-type":
			var systemNewPowerType string
			systemNewPowerType, err = cmd.Flags().GetString("power-type")
			if err != nil {
				return
			}
			system.PowerType = systemNewPowerType
		case "power-user":
			var systemNewPowerUser string
			systemNewPowerUser, err = cmd.Flags().GetString("power-user")
			if err != nil {
				return
			}
			system.PowerUser = systemNewPowerUser
		case "power-options":
			var systemNewPowerOptions string
			systemNewPowerOptions, err = cmd.Flags().GetString("power-options")
			if err != nil {
				return
			}
			system.PowerOptions = systemNewPowerOptions
		case "power-identity-file":
			var systemNewPowerIdentityFile string
			systemNewPowerIdentityFile, err = cmd.Flags().GetString("power-identity-file")
			if err != nil {
				return
			}
			system.PowerIdentityFile = systemNewPowerIdentityFile
		case "profile":
			var systemNewProfile string
			systemNewProfile, err = cmd.Flags().GetString("profile")
			if err != nil {
				return
			}
			system.Profile = systemNewProfile
		case "status":
			var systemNewStatus string
			systemNewStatus, err = cmd.Flags().GetString("status")
			if err != nil {
				return
			}
			system.Status = systemNewStatus
		case "virt-pxe-boot":
			var systemNewVirtPxeBoot bool
			systemNewVirtPxeBoot, err = cmd.Flags().GetBool("virt-pxe-boot")
			if err != nil {
				return
			}
			system.VirtPXEBoot = systemNewVirtPxeBoot
		case "serial-device":
			var systemNewSerialDevice int
			systemNewSerialDevice, err = cmd.Flags().GetInt("serial-device")
			if err != nil {
				return
			}
			system.SerialDevice = systemNewSerialDevice
		case "serial-baud-rate":
			var systemNewSerialBaudRate int
			systemNewSerialBaudRate, err = cmd.Flags().GetInt("serial-baud-rate")
			if err != nil {
				return
			}
			system.SerialBaudRate = systemNewSerialBaudRate
		case "bonding-opts":
			var systemNewBondingOpts string
			systemNewBondingOpts, err = cmd.Flags().GetString("bonding-opts")
			if err != nil {
				return
			}
			systemInterface.BondingOpts = systemNewBondingOpts
		case "bridge-opts":
			var systemNewBridgeOpts string
			systemNewBridgeOpts, err = cmd.Flags().GetString("bridge-opts")
			if err != nil {
				return
			}
			systemInterface.BridgeOpts = systemNewBridgeOpts
		case "cnames":
			var systemNewCNames []string
			systemNewCNames, err = cmd.Flags().GetStringSlice("cnames")
			if err != nil {
				return
			}
			systemInterface.CNAMEs = systemNewCNames
		case "connected-mode":
			var systemNewConnectedMode bool
			systemNewConnectedMode, err = cmd.Flags().GetBool("connected-mode")
			if err != nil {
				return
			}
			systemInterface.ConnectedMode = systemNewConnectedMode
		case "dns-name":
			var systemNewDnsName string
			systemNewDnsName, err = cmd.Flags().GetString("dns-name")
			if err != nil {
				return
			}
			systemInterface.DNSName = systemNewDnsName
		case "if-gateway":
			var systemNewIfGateway string
			systemNewIfGateway, err = cmd.Flags().GetString("if-gateway")
			if err != nil {
				return
			}
			systemInterface.Gateway = systemNewIfGateway
		case "interface-master":
			var systemNewInterfaceMaster string
			systemNewInterfaceMaster, err = cmd.Flags().GetString("interface-master")
			if err != nil {
				return
			}
			systemInterface.InterfaceMaster = systemNewInterfaceMaster
		case "interface-type":
			var systemNewInterfaceType string
			systemNewInterfaceType, err = cmd.Flags().GetString("interface-type")
			if err != nil {
				return
			}
			systemInterface.InterfaceType = systemNewInterfaceType
		case "ip-address":
			var systemNewIpAddress string
			systemNewIpAddress, err := cmd.Flags().GetString("ip-address")
			if err != nil {
				return
			}
			systemInterface.IPAddress = systemNewIpAddress
		case "ipv6-address":
			var systemNewIpv6Address string
			systemNewIpv6Address, err = cmd.Flags().GetString("ipv6-address")
			if err != nil {
				return
			}
			systemInterface.IPv6Address = systemNewIpv6Address
		case "ipv6-default-gateway":
			var systemNewIpv6DefaultGateway string
			systemNewIpv6DefaultGateway, err = cmd.Flags().GetString("ipv6-default-gateway")
			if err != nil {
				return
			}
			systemInterface.IPv6DefaultGateway = systemNewIpv6DefaultGateway
		case "ipv6-mtu":
			var systemNewIpv6Mtu string
			systemNewIpv6Mtu, err = cmd.Flags().GetString("ipv6-mtu")
			if err != nil {
				return
			}
			systemInterface.IPv6MTU = systemNewIpv6Mtu
		case "ipv6-prefix":
			var systemNewIpv6Prefix string
			systemNewIpv6Prefix, err = cmd.Flags().GetString("ipv6-prefix")
			if err != nil {
				return
			}
			systemInterface.IPv6Prefix = systemNewIpv6Prefix
		case "ipv6-secondaries":
			var systemNewIpv6Secondaries []string
			systemNewIpv6Secondaries, err = cmd.Flags().GetStringSlice("ipv6-secondaries")
			if err != nil {
				return
			}
			systemInterface.IPv6Secondaries = systemNewIpv6Secondaries
		case "ipv6-static-routes":
			var systemNewIpv6StaticRoutes []string
			systemNewIpv6StaticRoutes, err = cmd.Flags().GetStringSlice("ipv6-static-routes")
			if err != nil {
				return
			}
			systemInterface.IPv6StaticRoutes = systemNewIpv6StaticRoutes
		case "mac-address":
			var systemNewMacAddress string
			systemNewMacAddress, err = cmd.Flags().GetString("mac-address")
			if err != nil {
				return
			}
			systemInterface.MACAddress = systemNewMacAddress
		case "management":
			var systemNewManagement bool
			systemNewManagement, err = cmd.Flags().GetBool("management")
			if err != nil {
				return
			}
			systemInterface.Management = systemNewManagement
		case "mtu":
			var systemNewMtu string
			systemNewMtu, err = cmd.Flags().GetString("mtu")
			if err != nil {
				return
			}
			systemInterface.MTU = systemNewMtu
		case "netmask":
			var systemNewNetmask string
			systemNewNetmask, err = cmd.Flags().GetString("netmask")
			if err != nil {
				return
			}
			systemInterface.Netmask = systemNewNetmask
		case "static":
			var systemNewStatic bool
			systemNewStatic, err = cmd.Flags().GetBool("static")
			if err != nil {
				return
			}
			systemInterface.Static = systemNewStatic
		case "static-routes":
			var systemNewStaticRoutes []string
			systemNewStaticRoutes, err = cmd.Flags().GetStringSlice("static-routes")
			if err != nil {
				return
			}
			systemInterface.StaticRoutes = systemNewStaticRoutes
		case "virt-bridge":
			var systemNewVirtBridge string
			systemNewVirtBridge, err = cmd.Flags().GetString("virt-bridge")
			if err != nil {
				return
			}
			systemInterface.VirtBridge = systemNewVirtBridge
		}
	})
	// Don't blindly return nil because maybe one of the flags had an issue retrieving an argument.
	return err
}

// systemCmd represents the system command
var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "System management",
	Long: `Let you manage systems.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-system for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var systemAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add system",
	Long:  `Adds a system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()

		newSystem := cobbler.NewSystem()
		var err error

		// internal fields (ctime, mtime, depth, uid, repos-enabled, ipv6-autoconfiguration) cannot be modified
		newSystem.Name, err = cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		// Update system in-memory
		err = updateSystemFromFlags(cmd, &newSystem)
		if err != nil {
			return err
		}
		// No create the system via XML-RPC
		// FIXME: Call modify_interface in the client when getting to the interfaces objects
		system, err := Client.CreateSystem(newSystem)
		if err != nil {
			return err
		}
		fmt.Printf("System %s created\n", system.Name)
		return nil
	},
}

var systemCopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy system",
	Long:  `Copies a system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		systemName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		systemNewName, err := cmd.Flags().GetString("newname")
		if err != nil {
			return err
		}

		systemHandle, err := Client.GetSystemHandle(systemName)
		if err != nil {
			return err
		}
		err = Client.CopySystem(systemHandle, systemNewName)
		if err != nil {
			return err
		}
		newSystem, err := Client.GetSystem(systemNewName, false, false)
		if err != nil {
			return err
		}
		// Update the system in-memory
		err = updateSystemFromFlags(cmd, newSystem)
		if err != nil {
			return err
		}
		// Update the system via XML-RPC
		// FIXME: Call modify_interface in the client when getting to the interfaces objects
		return Client.UpdateSystem(newSystem)
	},
}

var systemDumpVarsCmd = &cobra.Command{
	Use:   "dumpvars",
	Short: "dump system variables",
	Long:  `Prints all system variables to stdout.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()

		// Get CLI flags
		systemName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		// Now retrieve data
		blendedData, err := Client.GetBlendedData("", systemName)
		if err != nil {
			return err
		}
		// Print data
		// TODO: Deduplicate with profile
		for key, value := range blendedData {
			if value == nil {
				fmt.Printf("%s:\n", key)
				continue
			}
			valueType := reflect.TypeOf(value).Kind()
			switch valueType {
			case reflect.Bool:
				fmt.Printf("%s: %t\n", key, value.(bool))
			case reflect.Int64:
				fmt.Printf("%s: %d\n", key, value.(int64))
			case reflect.Int32:
				fmt.Printf("%s: %d\n", key, value.(int32))
			case reflect.Int16:
				fmt.Printf("%s: %d\n", key, value.(int16))
			case reflect.Int8:
				fmt.Printf("%s: %d\n", key, value.(int8))
			case reflect.Int:
				fmt.Printf("%s: %d\n", key, value.(int))
			case reflect.Float32:
				fmt.Printf("%s: %f\n", key, value.(float32))
			case reflect.Float64:
				fmt.Printf("%s: %f\n", key, value.(float64))
			case reflect.Slice, reflect.Array:
				arr := reflect.ValueOf(value)
				fmt.Printf("%s: [", key)
				for i := 0; i < arr.Len(); i++ {
					if i+1 != arr.Len() {
						fmt.Printf("'%v', ", arr.Index(i).Interface())
					} else {
						fmt.Printf("'%v'", arr.Index(i).Interface())
					}
				}
				fmt.Printf("]\n")
			case reflect.Map:
				res2B, _ := json.Marshal(value)
				fmt.Printf("%s: %s\n", key, string(res2B))
			default:
				fmt.Printf("%s: %s\n", key, value)
			}
		}
		return err
	},
}

var systemEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit system",
	Long:  `Edits a system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()

		// find profile through its name
		systemName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		updateSystem, err := Client.GetSystem(systemName, false, false)
		if err != nil {
			return err
		}

		// Update the system in-memory
		err = updateSystemFromFlags(cmd, updateSystem)
		if err != nil {
			return err
		}
		// Update the system via XML-RPC
		// FIXME: Call modify_interface in the client when getting to the interfaces objects
		return Client.UpdateSystem(updateSystem)
	},
}

var systemFindCmd = &cobra.Command{
	Use:   "find",
	Short: "find system",
	Long:  `Finds a given system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		return FindItemNames(cmd, args, "system")
	},
}

var systemGetAutoinstallCmd = &cobra.Command{
	Use:   "get-autoinstall",
	Short: "dump autoinstall XML",
	Long:  `Prints the autoinstall XML file of the given system to stdout.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		systemName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		systemExists, err := Client.HasItem("system", systemName)
		if err != nil {
			return err
		}
		if !systemExists {
			fmt.Println("System does not exist!")
			os.Exit(1)
		}
		autoinstallRendered, err := Client.GenerateAutoinstall("", systemName)
		if err != nil {
			return err
		}
		fmt.Println(autoinstallRendered)
		return nil
	},
}

var systemListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all systems",
	Long:  `Lists all available systems.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		systemNames, err := Client.ListSystemNames()
		if err != nil {
			return err
		}
		listItems("systems", systemNames)
		return nil
	},
}

var systemPowerOffCmd = &cobra.Command{
	Use:   "poweroff",
	Short: "power off system",
	Long:  `Powers off the selected system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()

		// Get flags
		systemName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		// Perform action
		systemHandle, err := Client.GetSystemHandle(systemName)
		if err != nil {
			return err
		}
		_, err = Client.PowerSystem(systemHandle, "off")
		return err
	},
}

var systemPowerOnCmd = &cobra.Command{
	Use:   "poweron",
	Short: "power on system",
	Long:  `Powers on the selected system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()

		// Get flags
		systemName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		// Perform action
		systemHandle, err := Client.GetSystemHandle(systemName)
		if err != nil {
			return err
		}
		_, err = Client.PowerSystem(systemHandle, "on")
		return err
	},
}

var systemPowerStatusCmd = &cobra.Command{
	Use:   "powerstatus",
	Short: "Power status of the system",
	Long:  `Querys the power status of the selected system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()

		// Get flags
		systemName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		// Perform action
		systemHandle, err := Client.GetSystemHandle(systemName)
		if err != nil {
			return err
		}
		_, err = Client.PowerSystem(systemHandle, "status")
		return err
	},
}

var systemRebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "reboot system",
	Long:  `Reboots the selected system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()

		// Get flags
		systemName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		// Perform action
		systemHandle, err := Client.GetSystemHandle(systemName)
		if err != nil {
			return err
		}
		_, err = Client.PowerSystem(systemHandle, "reboot")
		return err
	},
}

var systemRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove system",
	Long:  `Removes a given system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		return RemoveItemRecursive(cmd, args, "system")
	},
}

var systemRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename system",
	Long:  `Renames a given system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()

		// Get flags
		systemName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		systemNewName, err := cmd.Flags().GetString("newname")
		if err != nil {
			return err
		}

		// Perform action
		systemHandle, err := Client.GetSystemHandle(systemName)
		if err != nil {
			return err
		}
		err = Client.RenameSystem(systemHandle, systemNewName)
		if err != nil {
			return err
		}
		newSystem, err := Client.GetSystem(systemNewName, false, false)
		if err != nil {
			return err
		}
		err = updateSystemFromFlags(cmd, newSystem)
		if err != nil {
			return err
		}
		return Client.UpdateSystem(newSystem)
	},
}

func reportSystems(systemNames []string) error {
	for _, itemName := range systemNames {
		system, err := Client.GetSystem(itemName, false, false)
		if err != nil {
			return err
		}
		printStructured(system)
		fmt.Println("")
	}
	return nil
}

var systemReportCmd = &cobra.Command{
	Use:   "report",
	Short: "list all systems in detail",
	Long:  `Shows detailed information about all systems.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		itemNames := make([]string, 0)
		if name == "" {
			itemNames, err = Client.ListSystemNames()
			if err != nil {
				return err
			}
		} else {
			itemNames = append(itemNames, name)
		}
		return reportSystems(itemNames)
	},
}

func init() {
	rootCmd.AddCommand(systemCmd)
	systemCmd.AddCommand(systemAddCmd)
	systemCmd.AddCommand(systemCopyCmd)
	systemCmd.AddCommand(systemDumpVarsCmd)
	systemCmd.AddCommand(systemEditCmd)
	systemCmd.AddCommand(systemFindCmd)
	systemCmd.AddCommand(systemGetAutoinstallCmd)
	systemCmd.AddCommand(systemListCmd)
	systemCmd.AddCommand(systemPowerOffCmd)
	systemCmd.AddCommand(systemPowerOnCmd)
	systemCmd.AddCommand(systemPowerStatusCmd)
	systemCmd.AddCommand(systemRebootCmd)
	systemCmd.AddCommand(systemRemoveCmd)
	systemCmd.AddCommand(systemRenameCmd)
	systemCmd.AddCommand(systemReportCmd)

	// local flags for system add
	addCommonArgs(systemAddCmd)
	addStringFlags(systemAddCmd, systemStringFlagMetadata)
	addStringFlags(systemAddCmd, systemPowerStringFlagMetadata)
	addBoolFlags(systemAddCmd, systemBoolFlagMetadata)
	addIntFlags(systemAddCmd, systemIntFlagMetadata)
	addFloatFlags(systemAddCmd, systemFloatFlagMetadata)
	addStringSliceFlags(systemAddCmd, systemStringSliceFlagMetadata)
	addMapFlags(systemAddCmd, systemMapFlagMetadata)
	// Network interface flags
	addStringFlags(systemAddCmd, interfaceStringFlagMetadata)
	addBoolFlags(systemAddCmd, interfaceBoolFlagMetadata)
	addStringSliceFlags(systemAddCmd, interfaceStringSliceFlagMetadata)
	// Other
	systemAddCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	systemAddCmd.Flags().String("interface", "", "the interface to operate on")
	systemAddCmd.Flags().Bool("delete-interface", false, "delete the given interface (should be used with --interface)")
	systemAddCmd.Flags().String("rename-interface", "", "rename the given interface (should be used with --interface)")

	// local flags for system copy
	addCommonArgs(systemCopyCmd)
	addStringFlags(systemCopyCmd, systemStringFlagMetadata)
	addStringFlags(systemCopyCmd, systemPowerStringFlagMetadata)
	addBoolFlags(systemCopyCmd, systemBoolFlagMetadata)
	addIntFlags(systemCopyCmd, systemIntFlagMetadata)
	addFloatFlags(systemCopyCmd, systemFloatFlagMetadata)
	addStringSliceFlags(systemCopyCmd, systemStringSliceFlagMetadata)
	addMapFlags(systemCopyCmd, systemMapFlagMetadata)
	// Network interface flags
	addStringFlags(systemCopyCmd, interfaceStringFlagMetadata)
	addBoolFlags(systemCopyCmd, interfaceBoolFlagMetadata)
	addStringSliceFlags(systemCopyCmd, interfaceStringSliceFlagMetadata)
	// Other
	systemCopyCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	systemCopyCmd.Flags().String("interface", "", "the interface to operate on")
	systemCopyCmd.Flags().Bool("delete-interface", false, "delete the given interface (should be used with --interface)")
	systemCopyCmd.Flags().String("rename-interface", "", "rename the given interface (should be used with --interface)")

	// local flags for system dumpvars
	systemDumpVarsCmd.Flags().String("name", "", "the system name")

	// local flags for system edit
	addCommonArgs(systemEditCmd)
	addStringFlags(systemEditCmd, systemStringFlagMetadata)
	addStringFlags(systemEditCmd, systemPowerStringFlagMetadata)
	addBoolFlags(systemEditCmd, systemBoolFlagMetadata)
	addIntFlags(systemEditCmd, systemIntFlagMetadata)
	addFloatFlags(systemEditCmd, systemFloatFlagMetadata)
	addStringSliceFlags(systemEditCmd, systemStringSliceFlagMetadata)
	addMapFlags(systemEditCmd, systemMapFlagMetadata)
	// Network interface flags
	addStringFlags(systemEditCmd, interfaceStringFlagMetadata)
	addBoolFlags(systemEditCmd, interfaceBoolFlagMetadata)
	addStringSliceFlags(systemEditCmd, interfaceStringSliceFlagMetadata)
	// Other
	systemEditCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	systemEditCmd.Flags().String("interface", "", "the interface to operate on")

	// local flags for system find
	addCommonArgs(systemFindCmd)
	addStringFlags(systemFindCmd, systemStringFlagMetadata)
	addStringFlags(systemFindCmd, systemPowerStringFlagMetadata)
	addBoolFlags(systemFindCmd, systemBoolFlagMetadata)
	addIntFlags(systemFindCmd, systemIntFlagMetadata)
	addFloatFlags(systemFindCmd, systemFloatFlagMetadata)
	addStringSliceFlags(systemFindCmd, systemStringSliceFlagMetadata)
	addMapFlags(systemFindCmd, systemMapFlagMetadata)
	// Network interface flags
	addStringFlags(systemFindCmd, interfaceStringFlagMetadata)
	addBoolFlags(systemFindCmd, interfaceBoolFlagMetadata)
	addStringSliceFlags(systemFindCmd, interfaceStringSliceFlagMetadata)
	// Other
	systemFindCmd.Flags().String("ctime", "", "")
	systemFindCmd.Flags().String("depth", "", "")
	systemFindCmd.Flags().String("mtime", "", "")
	systemFindCmd.Flags().String("uid", "", "UID")
	systemFindCmd.Flags().String("interface", "", "the interface to operate on")

	// local flags for system get-autoinstall
	systemGetAutoinstallCmd.Flags().String("name", "", "the system name")

	// local flags for system poweroff
	systemPowerOffCmd.Flags().String("name", "", "the system name")

	// local flags for system poweron
	systemPowerOnCmd.Flags().String("name", "", "the system name")

	// local flags for system powerstatus
	systemPowerStatusCmd.Flags().String("name", "", "the system name")

	// local flags for system reboot
	systemRebootCmd.Flags().String("name", "", "the system name")

	// local flags for system remove
	systemRemoveCmd.Flags().String("name", "", "the system name")
	systemRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")

	// local flags for system rename
	addCommonArgs(systemRenameCmd)
	addStringFlags(systemRenameCmd, systemStringFlagMetadata)
	addStringFlags(systemRenameCmd, systemPowerStringFlagMetadata)
	addBoolFlags(systemRenameCmd, systemBoolFlagMetadata)
	addIntFlags(systemRenameCmd, systemIntFlagMetadata)
	addFloatFlags(systemRenameCmd, systemFloatFlagMetadata)
	addStringSliceFlags(systemRenameCmd, systemStringSliceFlagMetadata)
	addMapFlags(systemRenameCmd, systemMapFlagMetadata)
	// Network interface flags
	addStringFlags(systemRenameCmd, interfaceStringFlagMetadata)
	addBoolFlags(systemRenameCmd, interfaceBoolFlagMetadata)
	addStringSliceFlags(systemRenameCmd, interfaceStringSliceFlagMetadata)
	// Other
	systemRenameCmd.Flags().String("newname", "", "the new system name")
	systemRenameCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	systemRenameCmd.Flags().String("interface", "", "the interface to operate on")

	// local flags for system report
	systemReportCmd.Flags().String("name", "", "the system name")
}
