// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	cobbler "github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func updateProfileFromFlags(cmd *cobra.Command, profile *cobbler.Profile) error {
	inPlace, err := cmd.Flags().GetBool("in-place")
	if err != nil {
		return err
	}
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		if err != nil {
			// If one of the previous flags has had an error just directly return.
			return
		}
		switch flag.Name {
		case "name":
			return
		case "newname":
			return
		case "repos":
			var profileNewRepos []string
			profileNewRepos, err = cmd.Flags().GetStringSlice("repos")
			if err != nil {
				return
			}
			profile.Repos = profileNewRepos
		case "autoinstall":
			var profileNewAutoinstall string
			profileNewAutoinstall, err = cmd.Flags().GetString("autoinstall")
			if err != nil {
				return
			}
			profile.Autoinstall = profileNewAutoinstall
		case "autoinstall-meta":
			fallthrough
		case "autoinstall-meta-inherit":
			if cmd.Flags().Lookup("boot-loaders-inherit").Changed {
				profile.AutoinstallMeta.Data = make(map[string]interface{})
				profile.AutoinstallMeta.IsInherited, err = cmd.Flags().GetBool("boot-loaders-inherit")
				if err != nil {
					return
				}
			} else {
				var profileNewAutoinstallMeta map[string]string
				profileNewAutoinstallMeta, err = cmd.Flags().GetStringToString("autoinstall-meta")
				if err != nil {
					return
				}
				if inPlace {
					err = Client.ModifyItemInPlace(
						"profile",
						profile.Name,
						"autoinstall_meta",
						convertMapStringToMapInterface(profileNewAutoinstallMeta),
					)
					if err != nil {
						return
					}
				} else {
					profile.AutoinstallMeta.IsInherited = false
					profile.AutoinstallMeta.Data = convertMapStringToMapInterface(profileNewAutoinstallMeta)
				}
			}
		case "boot-files":
			fallthrough
		case "boot-files-inherit":
			if cmd.Flags().Lookup("boot-files-inherit").Changed {
				profile.BootFiles.Data = make(map[string]interface{})
				profile.BootFiles.IsInherited, err = cmd.Flags().GetBool("boot-files-inherit")
				if err != nil {
					return
				}
			} else {
				var profileNewBootFiles map[string]string
				profileNewBootFiles, err = cmd.Flags().GetStringToString("boot-files")
				if err != nil {
					return
				}
				profile.BootFiles.IsInherited = false
				profile.BootFiles.Data = convertMapStringToMapInterface(profileNewBootFiles)
			}
		case "boot-loaders":
			fallthrough
		case "boot-loaders-inherit":
			if cmd.Flags().Lookup("boot-loaders-inherit").Changed {
				profile.BootLoaders.Data = []string{}
				profile.BootLoaders.IsInherited, err = cmd.Flags().GetBool("boot-loaders-inherit")
				if err != nil {
					return
				}
			} else {
				var profileNewBootLoaders []string
				profileNewBootLoaders, err = cmd.Flags().GetStringSlice("boot-loaders")
				if err != nil {
					return
				}
				profile.BootLoaders.IsInherited = false
				profile.BootLoaders.Data = profileNewBootLoaders
			}
		case "distro":
			var profileNewDistro string
			profileNewDistro, err = cmd.Flags().GetString("distro")
			if err != nil {
				return
			}
			profile.Distro = profileNewDistro
		case "comment":
			var profileNewComment string
			profileNewComment, err = cmd.Flags().GetString("comment")
			if err != nil {
				return
			}
			profile.Comment = profileNewComment
		case "fetchable-files":
			fallthrough
		case "fetchable-files-inherit":
			if cmd.Flags().Lookup("fetchable-files-inherit").Changed {
				profile.FetchableFiles.Data = make(map[string]interface{})
				profile.FetchableFiles.IsInherited, err = cmd.Flags().GetBool("fetchable-files-inherit")
				if err != nil {
					return
				}
			} else {
				var profileNewFetchableFiles map[string]string
				profileNewFetchableFiles, err = cmd.Flags().GetStringToString("fetchable-files")
				if err != nil {
					return
				}
				if inPlace {
					err = Client.ModifyItemInPlace(
						"profile",
						profile.Name,
						"fetchable_files",
						convertMapStringToMapInterface(profileNewFetchableFiles),
					)
					if err != nil {
						return
					}
				} else {
					profile.FetchableFiles.IsInherited = false
					profile.FetchableFiles.Data = convertMapStringToMapInterface(profileNewFetchableFiles)
				}
			}
		case "kernel-options":
			fallthrough
		case "kernel-options-inherit":
			if cmd.Flags().Lookup("kernel-options-inherit").Changed {
				profile.KernelOptions.Data = make(map[string]interface{})
				profile.KernelOptions.IsInherited, err = cmd.Flags().GetBool("kernel-options-inherit")
				if err != nil {
					return
				}
			} else {
				var profileNewKernelOptions map[string]string
				profileNewKernelOptions, err = cmd.Flags().GetStringToString("kernel-options")
				if err != nil {
					return
				}
				if inPlace {
					err = Client.ModifyItemInPlace(
						"profile",
						profile.Name,
						"kernel_options",
						convertMapStringToMapInterface(profileNewKernelOptions),
					)
					if err != nil {
						return
					}
				} else {
					profile.KernelOptions.IsInherited = false
					profile.KernelOptions.Data = convertMapStringToMapInterface(profileNewKernelOptions)
				}
			}
		case "kernel-options-post":
			fallthrough
		case "kernel-options-post-inherit":
			if cmd.Flags().Lookup("kernel-options-post-inherit").Changed {
				profile.KernelOptionsPost.Data = make(map[string]interface{})
				profile.KernelOptionsPost.IsInherited, err = cmd.Flags().GetBool("kernel-options-post-inherit")
				if err != nil {
					return
				}
			} else {
				var profileNewKernelOptionsPost map[string]string
				profileNewKernelOptionsPost, err = cmd.Flags().GetStringToString("kernel-options-post")
				if err != nil {
					return
				}
				if inPlace {
					err = Client.ModifyItemInPlace(
						"profile",
						profile.Name,
						"kernel_options_post",
						convertMapStringToMapInterface(profileNewKernelOptionsPost),
					)
					if err != nil {
						return
					}
				} else {
					profile.KernelOptionsPost.IsInherited = false
					profile.KernelOptionsPost.Data = convertMapStringToMapInterface(profileNewKernelOptionsPost)
				}
			}
		case "mgmt-classes":
			fallthrough
		case "mgmt-classes-inherit":
			if cmd.Flags().Lookup("mgmt-classes-inherit").Changed {
				profile.MgmtClasses.Data = []string{}
				profile.MgmtClasses.IsInherited, err = cmd.Flags().GetBool("mgmt-classes-inherit")
				if err != nil {
					return
				}
			} else {
				var profileNewMgmtClasses []string
				profileNewMgmtClasses, err = cmd.Flags().GetStringSlice("mgmt-classes")
				if err != nil {
					return
				}
				profile.MgmtClasses.IsInherited = false
				profile.MgmtClasses.Data = profileNewMgmtClasses
			}
		case "owners":
			fallthrough
		case "owners-inherit":
			if cmd.Flags().Lookup("owners-inherit").Changed {
				profile.Owners.Data = []string{}
				profile.Owners.IsInherited, err = cmd.Flags().GetBool("owners-inherit")
				if err != nil {
					return
				}
			} else {
				var profileNewOwners []string
				profileNewOwners, err = cmd.Flags().GetStringSlice("owners")
				if err != nil {
					return
				}
				profile.Owners.IsInherited = false
				profile.Owners.Data = profileNewOwners
			}
		case "redhat-management-key":
			var profileNewRedhatManagementKey string
			profileNewRedhatManagementKey, err = cmd.Flags().GetString("redhat-management-key")
			if err != nil {
				return
			}
			profile.RedhatManagementKey = profileNewRedhatManagementKey
		case "template-files-post":
			fallthrough
		case "template-files-inherit":
			if cmd.Flags().Lookup("template-files-inherit").Changed {
				profile.TemplateFiles.Data = make(map[string]interface{})
				profile.TemplateFiles.IsInherited, err = cmd.Flags().GetBool("template-files-inherit")
				if err != nil {
					return
				}
			} else {
				var profileNewTemplateFiles map[string]string
				profileNewTemplateFiles, err = cmd.Flags().GetStringToString("template-files")
				if err != nil {
					return
				}
				if inPlace {
					err = Client.ModifyItemInPlace(
						"profile",
						profile.Name,
						"template_files",
						convertMapStringToMapInterface(profileNewTemplateFiles),
					)
					if err != nil {
						return
					}
				} else {
					profile.TemplateFiles.IsInherited = false
					profile.TemplateFiles.Data = convertMapStringToMapInterface(profileNewTemplateFiles)
				}
			}
		case "dhcp-tag":
			var profileNewDhcpTag string
			profileNewDhcpTag, err = cmd.Flags().GetString("dhcp-tag")
			if err != nil {
				return
			}
			profile.DHCPTag = profileNewDhcpTag
		case "enable-ipxe":
			fallthrough
		case "enable-ipxe-inherit":
			if cmd.Flags().Lookup("enable-ipxe-inherit").Changed {
				profile.EnableIPXE.Data = false
				profile.EnableIPXE.IsInherited, err = cmd.Flags().GetBool("enable-ipxe-inherit")
				if err != nil {
					return
				}
			} else {
				var profileNewEnableIpxe bool
				profileNewEnableIpxe, err = cmd.Flags().GetBool("enable-ipxe")
				if err != nil {
					return
				}
				profile.EnableIPXE.IsInherited = false
				profile.EnableIPXE.Data = profileNewEnableIpxe
			}
		case "enable-menu":
			fallthrough
		case "enable-menu-inherit":
			if cmd.Flags().Lookup("enable-menu-inherit").Changed {
				profile.EnableMenu.Data = false
				profile.EnableMenu.IsInherited, err = cmd.Flags().GetBool("enable-menu-inherit")
				if err != nil {
					return
				}
			} else {
				var profileNewEnableMenu bool
				profileNewEnableMenu, err = cmd.Flags().GetBool("enable-menu")
				if err != nil {
					return
				}
				profile.EnableMenu.IsInherited = false
				profile.EnableMenu.Data = profileNewEnableMenu
			}
		case "mgmt-parameters":
			fallthrough
		case "mgmt-parameters-inherit":
			if cmd.Flags().Lookup("mgmt-parameters-inherit").Changed {
				profile.MgmtParameters.Data = make(map[string]interface{})
				profile.MgmtParameters.IsInherited, err = cmd.Flags().GetBool("mgmt-parameters-inherit")
				if err != nil {
					return
				}
			} else {
				var profileNewMgmtParameters map[string]string
				profileNewMgmtParameters, err = cmd.Flags().GetStringToString("mgmt-parameters")
				if err != nil {
					return
				}
				profile.MgmtParameters.IsInherited = false
				profile.MgmtParameters.Data = convertMapStringToMapInterface(profileNewMgmtParameters)
			}
		case "name-servers":
			fallthrough
		case "name-servers-inherit":
			if cmd.Flags().Lookup("name-servers-inherit").Changed {
				profile.NameServers.Data = make([]string, 0)
				profile.NameServers.IsInherited = true
			} else {
				var profileNewNameServers []string
				profileNewNameServers, err = cmd.Flags().GetStringSlice("name-servers")
				if err != nil {
					return
				}
				profile.NameServers.Data = profileNewNameServers
				profile.NameServers.IsInherited = false
			}
		case "name-servers-search":
			fallthrough
		case "name-servers-search-inherit":
			if cmd.Flags().Lookup("name-servers-search-inherit").Changed {
				profile.NameServersSearch.Data = make([]string, 0)
				profile.NameServersSearch.IsInherited = true
			} else {
				var profileNewNameServersSearch []string
				profileNewNameServersSearch, err = cmd.Flags().GetStringSlice("name-servers-search")
				if err != nil {
					return
				}
				profile.NameServersSearch.Data = profileNewNameServersSearch
				profile.NameServersSearch.IsInherited = false
			}
		case "next-server-v4":
			var profileNewNextServerV4 string
			profileNewNextServerV4, err = cmd.Flags().GetString("next-server-v4")
			if err != nil {
				return
			}
			profile.NextServerv4 = profileNewNextServerV4
		case "next-server-v6":
			var profileNewNextServerV6 string
			profileNewNextServerV6, err = cmd.Flags().GetString("next-server-v6")
			if err != nil {
				return
			}
			profile.NextServerv6 = profileNewNextServerV6
		case "filename":
			var profileNewFilename string
			profileNewFilename, err = cmd.Flags().GetString("filename")
			if err != nil {
				return
			}
			profile.Filename = profileNewFilename
		case "parent":
			var profileNewParent string
			profileNewParent, err = cmd.Flags().GetString("parent")
			if err != nil {
				return
			}
			profile.Parent = profileNewParent
		case "proxy":
			var profileNewProxy string
			profileNewProxy, err = cmd.Flags().GetString("proxy")
			if err != nil {
				return
			}
			profile.Proxy = profileNewProxy
		case "server":
			var profileNewServer string
			profileNewServer, err = cmd.Flags().GetString("server")
			if err != nil {
				return
			}
			profile.Server = profileNewServer
		case "menu":
			var profileNewMenu string
			profileNewMenu, err = cmd.Flags().GetString("menu")
			if err != nil {
				return
			}
			profile.Menu = profileNewMenu
		case "virt-auto-boot":
			fallthrough
		case "virt-auto-boot-inherit":
			if cmd.Flags().Lookup("virt-auto-boot-inherit").Changed {
				profile.VirtAutoBoot.IsInherited = true
			} else {
				var profileNewVirtAutoBoot bool
				profileNewVirtAutoBoot, err = cmd.Flags().GetBool("virt-auto-boot")
				if err != nil {
					return
				}
				profile.VirtAutoBoot.Data = profileNewVirtAutoBoot
				profile.VirtAutoBoot.IsInherited = false
			}
		case "virt-bridge":
			var profileNewVirtBridge string
			profileNewVirtBridge, err = cmd.Flags().GetString("virt-bridge")
			if err != nil {
				return
			}
			profile.VirtBridge = profileNewVirtBridge
		case "virt-cpus":
			var profileNewVirtCpus int
			profileNewVirtCpus, err = cmd.Flags().GetInt("virt-cpus")
			if err != nil {
				return
			}
			profile.VirtCPUs = profileNewVirtCpus
		case "virt-disk-driver":
			var profileNewVirtDiskDriver string
			profileNewVirtDiskDriver, err = cmd.Flags().GetString("virt-disk-driver")
			if err != nil {
				return
			}
			profile.VirtDiskDriver = profileNewVirtDiskDriver
		case "virt-file-size":
			fallthrough
		case "virt-file-size-inherit":
			if cmd.Flags().Lookup("virt-auto-boot-inherit").Changed {
				profile.VirtAutoBoot.IsInherited = true
			} else {
				var profileNewVirtFileSize float64
				profileNewVirtFileSize, err = cmd.Flags().GetFloat64("virt-file-size")
				if err != nil {
					return
				}
				profile.VirtFileSize.Data = profileNewVirtFileSize
				profile.VirtFileSize.IsInherited = false
			}
		case "virt-path":
			var profileNewVirtPath string
			profileNewVirtPath, err = cmd.Flags().GetString("virt-path")
			if err != nil {
				return
			}
			profile.VirtPath = profileNewVirtPath
		case "virt-ram":
			fallthrough
		case "virt-ram-inherit":
			if cmd.Flags().Lookup("virt-auto-boot-inherit").Changed {
				profile.VirtRAM.IsInherited = true
			} else {
				var profileNewVirtRam int
				profileNewVirtRam, err = cmd.Flags().GetInt("virt-ram")
				if err != nil {
					return
				}
				profile.VirtRAM.Data = profileNewVirtRam
				profile.VirtRAM.IsInherited = false
			}
		case "virt-type":
			var profileNewVirtType string
			profileNewVirtType, err = cmd.Flags().GetString("virt-type")
			if err != nil {
				return
			}
			profile.VirtType = profileNewVirtType
		}
	})
	return nil
}

// NewProfileCmd builds a new command that represents the profile action
func NewProfileCmd() *cobra.Command {
	profileCmd := &cobra.Command{
		Use:   "profile",
		Short: "Profile management",
		Long: `Let you manage profiles.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-profile for more information.`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	profileCmd.AddCommand(NewProfileAddCmd())
	profileCmd.AddCommand(NewProfileCopyCmd())
	profileCmd.AddCommand(NewProfileDumpVars())
	profileCmd.AddCommand(NewProfileEditCmd())
	profileCmd.AddCommand(NewProfileFindCmd())
	profileCmd.AddCommand(NewProfileGetAutoinstallCmd())
	profileCmd.AddCommand(NewProfileListCmd())
	profileCmd.AddCommand(NewProfileRemoveCmd())
	profileCmd.AddCommand(NewProfileRenameCmd())
	profileCmd.AddCommand(NewProfileReportCmd())
	return profileCmd
}

func NewProfileAddCmd() *cobra.Command {
	profileAddCmd := &cobra.Command{
		Use:   "add",
		Short: "add profile",
		Long:  `Adds a profile.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			newProfile := cobbler.NewProfile()
			// internal fields (ctime, mtime, uid, depth, repos-enabled) cannot be modified
			newProfile.Name, err = cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			err = updateProfileFromFlags(cmd, &newProfile)
			if err != nil {
				return err
			}
			profile, err := Client.CreateProfile(newProfile)
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Profile %s created\n", profile.Name)
			return nil
		},
	}
	addCommonArgs(profileAddCmd)
	addStringFlags(profileAddCmd, profileStringFlagMetadata)
	addBoolFlags(profileAddCmd, profileBoolFlagMetadata)
	addIntFlags(profileAddCmd, profileIntFlagMetadata)
	addFloatFlags(profileAddCmd, profileFloatFlagMetadata)
	addStringSliceFlags(profileAddCmd, distroStringSliceFlagMetadata)
	addStringSliceFlags(profileAddCmd, profileStringSliceFlagMetadata)
	addMapFlags(profileAddCmd, distroMapFlagMetadata)
	addMapFlags(profileAddCmd, profileMapFlagMetadata)
	profileAddCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	return profileAddCmd
}

func NewProfileCopyCmd() *cobra.Command {
	profileCopyCmd := &cobra.Command{
		Use:   "copy",
		Short: "copy profile",
		Long:  `Copies a profile.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			profileName, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			profileNewName, err := cmd.Flags().GetString("newname")
			if err != nil {
				return err
			}

			profileHandle, err := Client.GetProfileHandle(profileName)
			if err != nil {
				return err
			}
			err = Client.CopyProfile(profileHandle, profileNewName)
			if err != nil {
				return err
			}
			newProfile, err := Client.GetProfile(profileNewName, false, false)
			if err != nil {
				return err
			}
			err = updateProfileFromFlags(cmd, newProfile)
			if err != nil {
				return err
			}
			return Client.UpdateProfile(newProfile)
		},
	}
	addCommonArgs(profileCopyCmd)
	addStringFlags(profileCopyCmd, profileStringFlagMetadata)
	addBoolFlags(profileCopyCmd, profileBoolFlagMetadata)
	addIntFlags(profileCopyCmd, profileIntFlagMetadata)
	addFloatFlags(profileCopyCmd, profileFloatFlagMetadata)
	addStringSliceFlags(profileCopyCmd, distroStringSliceFlagMetadata)
	addStringSliceFlags(profileCopyCmd, profileStringSliceFlagMetadata)
	addMapFlags(profileCopyCmd, distroMapFlagMetadata)
	addMapFlags(profileCopyCmd, profileMapFlagMetadata)
	profileCopyCmd.Flags().String("newname", "", "the new profile name")
	profileCopyCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	return profileCopyCmd
}

func NewProfileDumpVars() *cobra.Command {
	profileDumpVarsCmd := &cobra.Command{
		Use:   "dumpvars",
		Short: "dump profile variables",
		Long:  `Prints all profile variables to stdout.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			// Get CLI flags
			profileName, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}

			// Now retrieve data
			blendedData, err := Client.GetBlendedData(profileName, "")
			if err != nil {
				return err
			}
			// Print data
			printDumpVars(cmd, blendedData)
			return err
		},
	}
	profileDumpVarsCmd.Flags().String("name", "", "the profile name")
	return profileDumpVarsCmd
}

func NewProfileEditCmd() *cobra.Command {
	profileEditCmd := &cobra.Command{
		Use:   "edit",
		Short: "edit profile",
		Long:  `Edits a profile.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			// find profile through its name
			pname, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			updateProfile, err := Client.GetProfile(pname, false, false)
			if err != nil {
				return err
			}

			// internal fields (ctime, mtime, uid, depth, repos-enabled) cannot be modified
			err = updateProfileFromFlags(cmd, updateProfile)
			if err != nil {
				return err
			}
			return Client.UpdateProfile(updateProfile)
		},
	}
	addCommonArgs(profileEditCmd)
	addStringFlags(profileEditCmd, profileStringFlagMetadata)
	addBoolFlags(profileEditCmd, profileBoolFlagMetadata)
	addIntFlags(profileEditCmd, profileIntFlagMetadata)
	addFloatFlags(profileEditCmd, profileFloatFlagMetadata)
	addStringSliceFlags(profileEditCmd, distroStringSliceFlagMetadata)
	addStringSliceFlags(profileEditCmd, profileStringSliceFlagMetadata)
	addMapFlags(profileEditCmd, distroMapFlagMetadata)
	addMapFlags(profileEditCmd, profileMapFlagMetadata)
	profileEditCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	return profileEditCmd
}

func NewProfileFindCmd() *cobra.Command {
	profileFindCmd := &cobra.Command{
		Use:   "find",
		Short: "find profile",
		Long:  `Finds a given profile.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			return FindItemNames(cmd, args, "profile")
		},
	}
	addCommonArgs(profileFindCmd)
	addStringFlags(profileFindCmd, profileStringFlagMetadata)
	addBoolFlags(profileFindCmd, profileBoolFlagMetadata)
	addIntFlags(profileFindCmd, profileIntFlagMetadata)
	addFloatFlags(profileFindCmd, profileFloatFlagMetadata)
	addStringSliceFlags(profileFindCmd, distroStringSliceFlagMetadata)
	addStringSliceFlags(profileFindCmd, profileStringSliceFlagMetadata)
	addMapFlags(profileFindCmd, distroMapFlagMetadata)
	addMapFlags(profileFindCmd, profileMapFlagMetadata)
	addStringFlags(profileFindCmd, findStringFlagMetadata)
	addIntFlags(profileFindCmd, findIntFlagMetadata)
	addFloatFlags(profileFindCmd, findFloatFlagMetadata)
	return profileFindCmd
}

func NewProfileGetAutoinstallCmd() *cobra.Command {
	profileGetAutoinstallCmd := &cobra.Command{
		Use:   "get-autoinstall",
		Short: "dump autoinstall XML",
		Long:  `Prints the autoinstall XML file of the given profile to stdout.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			profileName, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			profileExists, err := Client.HasItem("profile", profileName)
			if err != nil {
				return err
			}
			if !profileExists {
				return fmt.Errorf("Profile does not exist!")

			}
			autoinstallRendered, err := Client.GenerateAutoinstall(profileName, "")
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), autoinstallRendered)
			return nil
		},
	}
	profileGetAutoinstallCmd.Flags().String("name", "", "the profile name")
	return profileGetAutoinstallCmd
}

func NewProfileListCmd() *cobra.Command {
	profileListCmd := &cobra.Command{
		Use:   "list",
		Short: "list all profiles",
		Long:  `Lists all available profiles.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			profileNames, err := Client.ListProfileNames()
			if err != nil {
				return err
			}
			listItems(cmd, "profiles", profileNames)
			return nil
		},
	}
	return profileListCmd
}

func NewProfileRemoveCmd() *cobra.Command {
	profileRemoveCmd := &cobra.Command{
		Use:   "remove",
		Short: "remove profile",
		Long:  `Removes a given profile.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			pname, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			recursiveDelete, err := cmd.Flags().GetBool("recursive")
			if err != nil {
				return err
			}
			return Client.DeleteProfileRecursive(pname, recursiveDelete)
		},
	}
	profileRemoveCmd.Flags().String("name", "", "the profile name")
	profileRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")
	return profileRemoveCmd
}

func NewProfileRenameCmd() *cobra.Command {
	profileRenameCmd := &cobra.Command{
		Use:   "rename",
		Short: "rename profile",
		Long:  `Renames a given profile.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			profileName, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			profileNewName, err := cmd.Flags().GetString("newname")
			if err != nil {
				return err
			}

			// Now do the real edit
			profileHandle, err := Client.GetProfileHandle(profileName)
			if err != nil {
				return err
			}
			err = Client.RenameProfile(profileHandle, profileNewName)
			if err != nil {
				return err
			}
			newProfile, err := Client.GetProfile(profileNewName, false, false)
			if err != nil {
				return err
			}
			err = updateProfileFromFlags(cmd, newProfile)
			if err != nil {
				return err
			}
			return Client.UpdateProfile(newProfile)
		},
	}
	addCommonArgs(profileRenameCmd)
	addStringFlags(profileRenameCmd, profileStringFlagMetadata)
	addBoolFlags(profileRenameCmd, profileBoolFlagMetadata)
	addIntFlags(profileRenameCmd, profileIntFlagMetadata)
	addFloatFlags(profileRenameCmd, profileFloatFlagMetadata)
	addStringSliceFlags(profileRenameCmd, distroStringSliceFlagMetadata)
	addStringSliceFlags(profileRenameCmd, profileStringSliceFlagMetadata)
	addMapFlags(profileRenameCmd, distroMapFlagMetadata)
	addMapFlags(profileRenameCmd, profileMapFlagMetadata)
	profileRenameCmd.Flags().String("newname", "", "the new profile name")
	profileRenameCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	return profileRenameCmd
}

func reportProfiles(cmd *cobra.Command, profileNames []string) error {
	for _, itemName := range profileNames {
		profile, err := Client.GetProfile(itemName, false, false)
		if err != nil {
			return err
		}
		printStructured(cmd, profile)
		fmt.Fprintln(cmd.OutOrStdout(), "")
	}
	return nil
}

func NewProfileReportCmd() *cobra.Command {
	profileReportCmd := &cobra.Command{
		Use:   "report",
		Short: "list all profiles in detail",
		Long:  `Shows detailed information about all profiles.`,
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
				itemNames, err = Client.ListProfileNames()
				if err != nil {
					return err
				}
			} else {
				itemNames = append(itemNames, name)
			}
			return reportProfiles(cmd, itemNames)
		},
	}
	profileReportCmd.Flags().String("name", "", "the profile name")
	return profileReportCmd
}
