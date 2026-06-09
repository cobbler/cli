// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"encoding/json"
	"fmt"
	cobbler "github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
)

func updateSystemFromFlags(cmd *cobra.Command, system *cobbler.System) error {
	// Network interfaces are first-class items in Cobbler 4.0.0 and managed via
	// the dedicated `cobbler interface` command. The legacy interface flags on
	// `cobbler system add/edit/copy/rename/find` have been removed in v1.0.0.
	var err error
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
		case "template-files":
			var systemNewTemplateFiles map[string]string
			systemNewTemplateFiles, err = cmd.Flags().GetStringToString("template-files")
			if err != nil {
				return
			}
			system.TemplateFiles = systemNewTemplateFiles
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
		case "name-servers":
			var systemNewNameServers []string
			systemNewNameServers, err = cmd.Flags().GetStringSlice("name-servers")
			if err != nil {
				return
			}
			system.DNS.NameServers.IsInherited = false
			system.DNS.NameServers.Data = systemNewNameServers
		case "name-servers-search":
			var systemNewNameServersSearch []string
			systemNewNameServersSearch, err = cmd.Flags().GetStringSlice("name-servers-search")
			if err != nil {
				return
			}
			system.DNS.NameServersSearch = systemNewNameServersSearch
		case "next-server-v4":
			var systemNewNextServerV4 string
			systemNewNextServerV4, err = cmd.Flags().GetString("next-server-v4")
			if err != nil {
				return
			}
			system.TFTP.NextServerV4 = systemNewNextServerV4
		case "next-server-v6":
			var systemNewNextServerV6 string
			systemNewNextServerV6, err = cmd.Flags().GetString("next-server-v6")
			if err != nil {
				return
			}
			system.TFTP.NextServerV6 = systemNewNextServerV6
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
		case "virt-auto-boot":
			fallthrough
		case "virt-auto-boot-inherit":
			if cmd.Flags().Lookup("virt-auto-boot-inherit").Changed {
				system.Virt.AutoBoot.Data = false
				system.Virt.AutoBoot.IsInherited = true
			} else {
				var systemNewVirtAutoBoot bool
				systemNewVirtAutoBoot, err = cmd.Flags().GetBool("virt-auto-boot")
				if err != nil {
					return
				}
				system.Virt.AutoBoot.Data = systemNewVirtAutoBoot
				system.Virt.AutoBoot.IsInherited = false
			}
		case "virt-cpus":
			fallthrough
		case "virt-cpus-inherit":
			if cmd.Flags().Lookup("virt-cpus-inherit").Changed {
				system.Virt.Cpus.IsInherited = true
			} else {
				var systemNewVirtCpus int
				systemNewVirtCpus, err = cmd.Flags().GetInt("virt-cpus")
				if err != nil {
					return
				}
				system.Virt.Cpus.Data = systemNewVirtCpus
				system.Virt.Cpus.IsInherited = false
			}
		case "virt-disk-driver":
			var systemNewVirtDiskDriver string
			systemNewVirtDiskDriver, err = cmd.Flags().GetString("virt-disk-driver")
			if err != nil {
				return
			}
			system.Virt.DiskDriver = systemNewVirtDiskDriver
		case "virt-file-size":
			fallthrough
		case "virt-file-size-inherit":
			if cmd.Flags().Lookup("virt-file-size-inherit").Changed {
				system.Virt.FileSize.IsInherited = true
			} else {
				var systemNewVirtFileSize float64
				systemNewVirtFileSize, err = cmd.Flags().GetFloat64("virt-file-size")
				if err != nil {
					return
				}
				system.Virt.FileSize.Data = systemNewVirtFileSize
				system.Virt.FileSize.IsInherited = false
			}
		case "virt-path":
			var systemNewVirtPath string
			systemNewVirtPath, err = cmd.Flags().GetString("virt-path")
			if err != nil {
				return
			}
			system.Virt.Path = systemNewVirtPath
		case "virt-ram":
			fallthrough
		case "virt-ram-inherit":
			if cmd.Flags().Lookup("virt-ram-inherit").Changed {
				system.Virt.Ram.IsInherited = true
			} else {
				var systemNewVirtRam int
				systemNewVirtRam, err = cmd.Flags().GetInt("virt-ram")
				if err != nil {
					return
				}
				system.Virt.Ram.Data = systemNewVirtRam
				system.Virt.Ram.IsInherited = false
			}
		case "virt-type":
			var systemNewVirtType string
			systemNewVirtType, err = cmd.Flags().GetString("virt-type")
			if err != nil {
				return
			}
			system.Virt.Type = systemNewVirtType
		case "gateway":
			var systemNewGateway string
			systemNewGateway, err = cmd.Flags().GetString("gateway")
			if err != nil {
				return
			}
			system.Gateway = systemNewGateway
		case "hostname":
			var systemNewHostname string
			systemNewHostname, err = cmd.Flags().GetString("hostname")
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
			system.Power.Address = systemNewPowerAddress
		case "power-id":
			var systemNewPowerId string
			systemNewPowerId, err = cmd.Flags().GetString("power-id")
			if err != nil {
				return
			}
			system.Power.ID = systemNewPowerId
		case "power-pass":
			var systemNewPowerPass string
			systemNewPowerPass, err = cmd.Flags().GetString("power-pass")
			if err != nil {
				return
			}
			system.Power.Password = systemNewPowerPass
		case "power-type":
			var systemNewPowerType string
			systemNewPowerType, err = cmd.Flags().GetString("power-type")
			if err != nil {
				return
			}
			system.Power.Type = systemNewPowerType
		case "power-user":
			var systemNewPowerUser string
			systemNewPowerUser, err = cmd.Flags().GetString("power-user")
			if err != nil {
				return
			}
			system.Power.User = systemNewPowerUser
		case "power-options":
			var systemNewPowerOptions string
			systemNewPowerOptions, err = cmd.Flags().GetString("power-options")
			if err != nil {
				return
			}
			system.Power.Options = systemNewPowerOptions
		case "power-identity-file":
			var systemNewPowerIdentityFile string
			systemNewPowerIdentityFile, err = cmd.Flags().GetString("power-identity-file")
			if err != nil {
				return
			}
			system.Power.IdentityFile = systemNewPowerIdentityFile
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
		}
	})
	// Don't blindly return nil because maybe one of the flags had an issue retrieving an argument.
	return err
}

// NewSystemCmd builds a new command that represents the system action
func NewSystemCmd() (*cobra.Command, error) {
	systemCmd := &cobra.Command{
		Use:   "system",
		Short: "System management",
		Long: `Let you manage systems.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-system for more information.`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	systemCmd.AddCommand(NewSystemAddCmd())
	systemCmd.AddCommand(NewSystemCopyCmd())
	systemCmd.AddCommand(NewSystemDumpVarsCmd())
	systemCmd.AddCommand(NewSystemEditCmd())
	systemCmd.AddCommand(NewSystemFindCmd())
	systemCmd.AddCommand(NewSystemGetAutoinstallCmd())
	systemCmd.AddCommand(NewSystemListCmd())
	systemCmd.AddCommand(NewSystemPowerOffCmd())
	systemCmd.AddCommand(NewSystemPowerOnCmd())
	systemCmd.AddCommand(NewSystemPowerStatusCmd())
	systemCmd.AddCommand(NewSystemRebootCmd())
	systemCmd.AddCommand(NewSystemRemoveCmd())
	systemCmd.AddCommand(NewSystemRenameCmd())
	systemCmd.AddCommand(NewSystemReportCmd())
	systemCmd.AddCommand(NewSystemExportCmd())
	return systemCmd, nil
}

func NewSystemAddCmd() *cobra.Command {
	systemAddCmd := &cobra.Command{
		Use:   "add",
		Short: "add system",
		Long:  `Adds a system.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			newSystem := cobbler.NewSystem()

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
			// Now create the system via XML-RPC
			system, err := Client.CreateSystem(newSystem)
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "System %s created\n", system.Name)
			return nil
		},
	}
	addCommonArgs(systemAddCmd)
	addStringFlags(systemAddCmd, systemStringFlagMetadata)
	addStringFlags(systemAddCmd, systemPowerStringFlagMetadata)
	addBoolFlags(systemAddCmd, systemBoolFlagMetadata)
	addIntFlags(systemAddCmd, systemIntFlagMetadata)
	addFloatFlags(systemAddCmd, systemFloatFlagMetadata)
	addStringSliceFlags(systemAddCmd, systemStringSliceFlagMetadata)
	addMapFlags(systemAddCmd, systemMapFlagMetadata)
	systemAddCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	return systemAddCmd
}

func NewSystemCopyCmd() *cobra.Command {
	systemCopyCmd := &cobra.Command{
		Use:   "copy",
		Short: "copy system",
		Long:  `Copies a system.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

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
			if newSystem.Meta.IsDirty {
				newSystem, err = Client.GetSystem(
					newSystem.Name,
					newSystem.Meta.IsFlattened,
					newSystem.Meta.IsResolved,
				)
				if err != nil {
					return err
				}
			}
			// Update the system via XML-RPC
			return Client.UpdateSystem(newSystem)
		},
	}
	addCommonArgs(systemCopyCmd)
	addStringFlags(systemCopyCmd, systemStringFlagMetadata)
	addStringFlags(systemCopyCmd, systemPowerStringFlagMetadata)
	addBoolFlags(systemCopyCmd, systemBoolFlagMetadata)
	addIntFlags(systemCopyCmd, systemIntFlagMetadata)
	addFloatFlags(systemCopyCmd, systemFloatFlagMetadata)
	addStringSliceFlags(systemCopyCmd, systemStringSliceFlagMetadata)
	addMapFlags(systemCopyCmd, systemMapFlagMetadata)
	// Network interface flags
	addStringFlags(systemCopyCmd, copyRenameStringFlagMetadata)
	systemCopyCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	return systemCopyCmd
}

func NewSystemDumpVarsCmd() *cobra.Command {
	systemDumpVarsCmd := &cobra.Command{
		Use:   "dumpvars",
		Short: "dump system variables",
		Long:  `Prints all system variables to stdout.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			systemName, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			system, err := Client.GetSystem(systemName, false, false)
			if err != nil {
				return err
			}
			blendedData, err := Client.DumpVars(system.Uid, false, false)
			if err != nil {
				return err
			}
			printDumpVars(cmd, blendedData)
			return err
		},
	}
	systemDumpVarsCmd.Flags().String("name", "", "the system name")
	return systemDumpVarsCmd
}

func NewSystemEditCmd() *cobra.Command {
	systemEditCmd := &cobra.Command{
		Use:   "edit",
		Short: "edit system",
		Long:  `Edits a system.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

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
			if updateSystem.Meta.IsDirty {
				updateSystem, err = Client.GetSystem(
					updateSystem.Name,
					updateSystem.Meta.IsFlattened,
					updateSystem.Meta.IsResolved,
				)
				if err != nil {
					return err
				}
			}
			// Update the system via XML-RPC
			return Client.UpdateSystem(updateSystem)
		},
	}
	addCommonArgs(systemEditCmd)
	addStringFlags(systemEditCmd, systemStringFlagMetadata)
	addStringFlags(systemEditCmd, systemPowerStringFlagMetadata)
	addBoolFlags(systemEditCmd, systemBoolFlagMetadata)
	addIntFlags(systemEditCmd, systemIntFlagMetadata)
	addFloatFlags(systemEditCmd, systemFloatFlagMetadata)
	addStringSliceFlags(systemEditCmd, systemStringSliceFlagMetadata)
	addMapFlags(systemEditCmd, systemMapFlagMetadata)
	// Network interface flags
	systemEditCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	return systemEditCmd
}

func NewSystemFindCmd() *cobra.Command {
	systemFindCmd := &cobra.Command{
		Use:   "find",
		Short: "find system",
		Long:  `Finds a given system.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			return FindItemNames(cmd, args, "system")
		},
	}
	addCommonArgs(systemFindCmd)
	addStringFlags(systemFindCmd, systemStringFlagMetadata)
	addStringFlags(systemFindCmd, systemPowerStringFlagMetadata)
	addBoolFlags(systemFindCmd, systemBoolFlagMetadata)
	addIntFlags(systemFindCmd, systemIntFlagMetadata)
	addFloatFlags(systemFindCmd, systemFloatFlagMetadata)
	addStringSliceFlags(systemFindCmd, systemStringSliceFlagMetadata)
	addMapFlags(systemFindCmd, systemMapFlagMetadata)
	// Network interface flags
	addStringFlags(systemFindCmd, findStringFlagMetadata)
	addIntFlags(systemFindCmd, findIntFlagMetadata)
	addFloatFlags(systemFindCmd, findFloatFlagMetadata)
	addPaginationFlags(systemFindCmd)
	return systemFindCmd
}

func NewSystemGetAutoinstallCmd() *cobra.Command {
	systemGetAutoinstallCmd := &cobra.Command{
		Use:   "get-autoinstall",
		Short: "dump autoinstall XML",
		Long:  `Prints the autoinstall XML file of the given system to stdout.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			systemName, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			systemExists, err := Client.HasItem("system", systemName)
			if err != nil {
				return err
			}
			if !systemExists {
				//goland:noinspection GoErrorStringFormat
				return fmt.Errorf("System does not exist")
			}
			autoinstallRendered, err := Client.GenerateAutoinstall(systemName, "system", "name", "", "")
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), autoinstallRendered)
			return nil
		},
	}
	systemGetAutoinstallCmd.Flags().String("name", "", "the system name")
	return systemGetAutoinstallCmd
}

func NewSystemListCmd() *cobra.Command {
	systemListCmd := &cobra.Command{
		Use:   "list",
		Short: "list all systems",
		Long:  `Lists all available systems.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			systemNames, err := Client.ListSystemNames()
			if err != nil {
				return err
			}
			listItems(cmd, "systems", systemNames)
			return nil
		},
	}
	return systemListCmd
}

func NewSystemPowerOffCmd() *cobra.Command {
	systemPowerOffCmd := &cobra.Command{
		Use:   "poweroff",
		Short: "power off system",
		Long:  `Powers off the selected system.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

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
	systemPowerOffCmd.Flags().String("name", "", "the system name")
	return systemPowerOffCmd
}

func NewSystemPowerOnCmd() *cobra.Command {
	systemPowerOnCmd := &cobra.Command{
		Use:   "poweron",
		Short: "power on system",
		Long:  `Powers on the selected system.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

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
	systemPowerOnCmd.Flags().String("name", "", "the system name")
	return systemPowerOnCmd
}

func NewSystemPowerStatusCmd() *cobra.Command {
	systemPowerStatusCmd := &cobra.Command{
		Use:   "powerstatus",
		Short: "Power status of the system",
		Long:  `Querys the power status of the selected system.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

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
	systemPowerStatusCmd.Flags().String("name", "", "the system name")
	return systemPowerStatusCmd
}

func NewSystemRebootCmd() *cobra.Command {
	systemRebootCmd := &cobra.Command{
		Use:   "reboot",
		Short: "reboot system",
		Long:  `Reboots the selected system.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

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
	systemRebootCmd.Flags().String("name", "", "the system name")
	return systemRebootCmd
}

func NewSystemRemoveCmd() *cobra.Command {
	systemRemoveCmd := &cobra.Command{
		Use:   "remove",
		Short: "remove system",
		Long:  `Removes a given system.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			return RemoveItemRecursive(cmd, args, "system")
		},
	}
	systemRemoveCmd.Flags().String("name", "", "the system name")
	systemRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")
	return systemRemoveCmd
}

func NewSystemRenameCmd() *cobra.Command {
	systemRenameCmd := &cobra.Command{
		Use:   "rename",
		Short: "rename system",
		Long:  `Renames a given system.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

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
			if newSystem.Meta.IsDirty {
				newSystem, err = Client.GetSystem(
					newSystem.Name,
					newSystem.Meta.IsFlattened,
					newSystem.Meta.IsResolved,
				)
				if err != nil {
					return err
				}
			}
			return Client.UpdateSystem(newSystem)
		},
	}
	addCommonArgs(systemRenameCmd)
	addStringFlags(systemRenameCmd, systemStringFlagMetadata)
	addStringFlags(systemRenameCmd, systemPowerStringFlagMetadata)
	addBoolFlags(systemRenameCmd, systemBoolFlagMetadata)
	addIntFlags(systemRenameCmd, systemIntFlagMetadata)
	addFloatFlags(systemRenameCmd, systemFloatFlagMetadata)
	addStringSliceFlags(systemRenameCmd, systemStringSliceFlagMetadata)
	addMapFlags(systemRenameCmd, systemMapFlagMetadata)
	// Network interface flags
	addStringFlags(systemRenameCmd, copyRenameStringFlagMetadata)
	systemRenameCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	return systemRenameCmd
}

func reportSystems(cmd *cobra.Command, systemNames []string) error {
	for _, itemName := range systemNames {
		system, err := Client.GetSystem(itemName, false, false)
		if err != nil {
			return err
		}
		printStructured(cmd, system)
		fmt.Fprintln(cmd.OutOrStdout(), "")
	}
	return nil
}

func NewSystemReportCmd() *cobra.Command {
	systemReportCmd := &cobra.Command{
		Use:   "report",
		Short: "list all systems in detail",
		Long:  `Shows detailed information about all systems.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

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
			return reportSystems(cmd, itemNames)
		},
	}
	systemReportCmd.Flags().String("name", "", "the system name")
	return systemReportCmd
}

func NewSystemExportCmd() *cobra.Command {
	systemExportCmd := &cobra.Command{
		Use:   "export",
		Short: "export systems",
		Long:  `Export systems.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			formatOption, err := cmd.Flags().GetString("format")
			if err != nil {
				return err
			}
			if formatOption != "json" && formatOption != "yaml" {
				return fmt.Errorf("format must be json or yaml")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			formatOption, err := cmd.Flags().GetString("format")
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

			for _, itemName := range itemNames {
				system, err := Client.GetSystem(itemName, false, false)
				if err != nil {
					return err
				}
				if formatOption == "json" {
					jsonDocument, err := json.Marshal(system)
					if err != nil {
						return err
					}
					fmt.Fprintln(cmd.OutOrStdout(), string(jsonDocument))
				}
				if formatOption == "yaml" {
					yamlDocument, err := yaml.Marshal(system)
					if err != nil {
						return err
					}
					fmt.Fprintln(cmd.OutOrStdout(), "---")
					fmt.Fprintln(cmd.OutOrStdout(), string(yamlDocument))
				}
			}
			return nil
		},
	}
	systemExportCmd.Flags().String("name", "", "the system name")
	systemExportCmd.Flags().String(exportStringMetadata["format"].Name, exportStringMetadata["format"].DefaultValue, exportStringMetadata["format"].Usage)
	return systemExportCmd
}
