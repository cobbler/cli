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

func updateDistroFromFlags(cmd *cobra.Command, distro *cobbler.Distro) error {
	var inPlace bool
	var err error
	if cmd.Flags().Lookup("in-place") != nil {
		inPlace, err = cmd.Flags().GetBool("in-place")
		if err != nil {
			return err
		}
	}
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		if err != nil {
			// If one of the previous flags has had an error just directly return.
			return
		}
		switch flag.Name {
		// The rename & copy operations are special operations as such we cannot blindly set this inside here.
		// Any rename & copy operation must be handled outside of this method.
		case "arch":
			var distroNewArch string
			distroNewArch, err = cmd.Flags().GetString("arch")
			if err != nil {
				return
			}
			distro.Arch = distroNewArch
		case "autoinstall-meta":
			fallthrough
		case "autoinstall-meta-inherit":
			var distroNewAutoinstallMeta map[string]string
			distroNewAutoinstallMeta, err = cmd.Flags().GetStringToString("autoinstall-meta")
			if err != nil {
				return
			}
			if inPlace {
				err = Client.ModifyItemInPlace(
					"distro",
					distro.Name,
					"autoinstall_meta",
					convertMapStringToMapInterface(distroNewAutoinstallMeta),
				)
				if err != nil {
					return
				}
			} else {
				distro.AutoinstallMeta.IsInherited = false
				distro.AutoinstallMeta.Data = convertMapStringToMapInterface(distroNewAutoinstallMeta)
			}
		case "boot-files":
			fallthrough
		case "boot-files-inherit":
			var distroNewBootFiles map[string]string
			distroNewBootFiles, err = cmd.Flags().GetStringToString("boot-files")
			if err != nil {
				return
			}
			if cmd.Flags().Lookup("boot-files-inherit").Changed {
				distro.BootFiles.Data = make(map[string]interface{})
				distro.BootFiles.IsInherited, err = cmd.Flags().GetBool("boot-files-inherit")
				if err != nil {
					return
				}
			} else {
				distro.BootFiles.IsInherited = false
				distro.BootFiles.Data = convertMapStringToMapInterface(distroNewBootFiles)
			}
		case "boot-loaders":
			fallthrough
		case "boot-loaders-inherit":
			var distroNewBootLoaders []string
			distroNewBootLoaders, err = cmd.Flags().GetStringSlice("boot-loaders")
			if err != nil {
				return
			}
			if cmd.Flags().Lookup("boot-loaders-inherit").Changed {
				distro.BootLoaders.Data = []string{}
				distro.BootLoaders.IsInherited, err = cmd.Flags().GetBool("boot-loaders-inherit")
				if err != nil {
					return
				}
			} else {
				distro.BootLoaders.IsInherited = false
				distro.BootLoaders.Data = distroNewBootLoaders
			}
		case "breed":
			var distroNewBreed string
			distroNewBreed, err = cmd.Flags().GetString("breed")
			if err != nil {
				return
			}
			distro.Breed = distroNewBreed
		case "comment":
			var distroNewComment string
			distroNewComment, err = cmd.Flags().GetString("comment")
			if err != nil {
				return
			}
			distro.Comment = distroNewComment
		case "fetchable-files":
			fallthrough
		case "fetchable-files-inherit":
			var newFetchableFiles map[string]string
			newFetchableFiles, err = cmd.Flags().GetStringToString("fetchable-files")
			if err != nil {
				return
			}
			if cmd.Flags().Lookup("fetchable-files-inherit").Changed {
				distro.FetchableFiles.Data = make(map[string]interface{})
				distro.FetchableFiles.IsInherited, err = cmd.Flags().GetBool("fetchable-files-inherit")
				if err != nil {
					return
				}
			} else {
				if inPlace {
					err = Client.ModifyItemInPlace(
						"distro",
						distro.Name,
						"fetchable_files",
						convertMapStringToMapInterface(newFetchableFiles),
					)
					if err != nil {
						return
					}
				} else {
					distro.FetchableFiles.IsInherited = false
					distro.FetchableFiles.Data = convertMapStringToMapInterface(newFetchableFiles)
				}
			}
		case "initrd":
			var distroNewInitrd string
			distroNewInitrd, err = cmd.Flags().GetString("initrd")
			if err != nil {
				return
			}
			distro.Initrd = distroNewInitrd
		case "remote-boot-initrd":
			var distroNewRemoteBootInitrd string
			distroNewRemoteBootInitrd, err = cmd.Flags().GetString("remote-boot-initrd")
			if err != nil {
				return
			}
			distro.RemoteBootInitrd = distroNewRemoteBootInitrd
		case "kernel":
			var distroNewKernel string
			distroNewKernel, err = cmd.Flags().GetString("kernel")
			if err != nil {
				return
			}
			distro.Kernel = distroNewKernel
		case "remote-boot-kernel":
			var distroNewRemoteBootKernel string
			distroNewRemoteBootKernel, err = cmd.Flags().GetString("remote-boot-kernel")
			if err != nil {
				return
			}
			distro.RemoteBootKernel = distroNewRemoteBootKernel
		case "kernel-options":
			fallthrough
		case "kernel-options-inherit":
			if cmd.Flags().Lookup("kernel-options-inherit").Changed {
				distro.KernelOptions.Data = make(map[string]interface{})
				distro.KernelOptions.IsInherited, err = cmd.Flags().GetBool("kernel-options-inherit")
				if err != nil {
					return
				}
			} else {
				var newKernelOptions map[string]string
				newKernelOptions, err = cmd.Flags().GetStringToString("kernel-options")
				if err != nil {
					return
				}
				if inPlace {
					err = Client.ModifyItemInPlace(
						"distro",
						distro.Name,
						"kernel_options",
						convertMapStringToMapInterface(newKernelOptions),
					)
					if err != nil {
						return
					}
				} else {
					distro.KernelOptions.IsInherited = false
					distro.KernelOptions.Data = convertMapStringToMapInterface(newKernelOptions)
				}
			}
		case "kernel-options-post":
			fallthrough
		case "kernel-options-post-inherit":
			if cmd.Flags().Lookup("kernel-options-post-inherit").Changed {
				distro.KernelOptionsPost.Data = make(map[string]interface{})
				distro.KernelOptionsPost.IsInherited, err = cmd.Flags().GetBool("kernel-options-post-inherit")
				if err != nil {
					return
				}
			} else {
				var newKernelOptionsPost map[string]string
				newKernelOptionsPost, err = cmd.Flags().GetStringToString("kernel-options-post")
				if err != nil {
					return
				}
				if inPlace {
					err = Client.ModifyItemInPlace(
						"distro",
						distro.Name,
						"kernel_options_post",
						convertMapStringToMapInterface(newKernelOptionsPost),
					)
					if err != nil {
						return
					}
				} else {
					distro.KernelOptionsPost.IsInherited = false
					distro.KernelOptions.Data = convertMapStringToMapInterface(newKernelOptionsPost)
				}
			}
		case "mgmt-classes":
			fallthrough
		case "mgmt-classes-inherit":
			var distroNewMgmtClasses []string
			distroNewMgmtClasses, err = cmd.Flags().GetStringSlice("mgmt-classes")
			if err != nil {
				return
			}
			if cmd.Flags().Lookup("mgmt-classes-inherit").Changed {
				distro.MgmtClasses.Data = []string{}
				distro.MgmtClasses.IsInherited, err = cmd.Flags().GetBool("mgmt-classes-inherit")
				if err != nil {
					return
				}
			} else {
				distro.MgmtClasses.IsInherited = false
				distro.MgmtClasses.Data = distroNewMgmtClasses
			}
		case "os-version":
			var distroNewOsVersion string
			distroNewOsVersion, err = cmd.Flags().GetString("os-version")
			if err != nil {
				return
			}
			distro.OSVersion = distroNewOsVersion
		case "owners":
			fallthrough
		case "owners-inherit":
			var distroNewOwners []string
			distroNewOwners, err = cmd.Flags().GetStringSlice("owners")
			if err != nil {
				return
			}
			if cmd.Flags().Lookup("owners-inherit").Changed {
				distro.Owners.Data = []string{}
				distro.Owners.IsInherited, err = cmd.Flags().GetBool("owners-inherit")
				if err != nil {
					return
				}
			} else {
				distro.Owners.IsInherited = false
				distro.Owners.Data = distroNewOwners
			}
		case "redhat-management-key":
			var distroNewRedhatManagementKey string
			distroNewRedhatManagementKey, err = cmd.Flags().GetString("redhat-management-key")
			if err != nil {
				return
			}
			distro.RedhatManagementKey = distroNewRedhatManagementKey
		case "template-files":
			fallthrough
		case "template-files-inherit":
			if cmd.Flags().Lookup("template-files-inherit").Changed {
				distro.TemplateFiles.Data = make(map[string]interface{})
				distro.TemplateFiles.IsInherited, err = cmd.Flags().GetBool("template-files-inherit")
				if err != nil {
					return
				}
			} else {
				var newTemplateFiles map[string]string
				newTemplateFiles, err = cmd.Flags().GetStringToString("template-files")
				if err != nil {
					return
				}
				if inPlace {
					err = Client.ModifyItemInPlace(
						"distro",
						distro.Name,
						"template_files",
						convertMapStringToMapInterface(newTemplateFiles),
					)
					if err != nil {
						return
					}
				} else {
					distro.TemplateFiles.IsInherited = false
					distro.TemplateFiles.Data = convertMapStringToMapInterface(newTemplateFiles)
				}
			}
		}
	})
	if inPlace {
		// Update distro in case we did modify the distro
		distro.Meta.IsDirty = true
	}
	// Don't blindly return nil because maybe one of the flags had an issue retrieving an argument.
	return err
}

// distroCmd represents the distro command
var distroCmd = &cobra.Command{
	Use:   "distro",
	Short: "Distribution management",
	Long: `Let you manage distributions.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-distro for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var distroAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add distribution",
	Long:  `Adds a distribution.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		newDistro := cobbler.NewDistro()
		var err error

		// internal fields (ctime, mtime, depth, uid, source-repos, tree-build-time) cannot be modified
		newDistro.Name, err = cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		// Update distro in-memory
		err = updateDistroFromFlags(cmd, &newDistro)
		if err != nil {
			return err
		}
		// Now create the distro via XML-RPC
		distro, err := Client.CreateDistro(newDistro)
		if err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Distro %s created\n", distro.Name)
		return nil
	},
}

var distroCopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy distribution",
	Long:  `Copies a distribution.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		dname, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		distroNewName, err := cmd.Flags().GetString("newname")
		if err != nil {
			return err
		}

		dhandle, err := Client.GetDistroHandle(dname)
		if err != nil {
			return err
		}
		err = Client.CopyDistro(dhandle, distroNewName)
		if err != nil {
			return err
		}
		newDistro, err := Client.GetDistro(distroNewName, false, false)
		if err != nil {
			return err
		}
		// Update distro in-memory
		err = updateDistroFromFlags(cmd, newDistro)
		if err != nil {
			return err
		}
		// Update the distro via XML-RPC
		return Client.UpdateDistro(newDistro)
	},
}

var distroEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit distribution",
	Long:  `Edits a distribution.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()

		// find distro through its name
		dname, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		// Get distro from the API
		updateDistro, err := Client.GetDistro(dname, false, false)
		if err != nil {
			return err
		}
		// Update distro in-memory
		err = updateDistroFromFlags(cmd, updateDistro)
		if err != nil {
			return err
		}
		if updateDistro.Meta.IsDirty {
			updateDistro, err = Client.GetDistro(
				updateDistro.Name,
				updateDistro.Meta.IsFlattened,
				updateDistro.Meta.IsResolved,
			)
			if err != nil {
				return err
			}
		}
		// Now update distro via XML-RPC
		return Client.UpdateDistro(updateDistro)
	},
}

var distroFindCmd = &cobra.Command{
	Use:   "find",
	Short: "find distribution",
	Long:  `Finds a given distribution.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		return FindItemNames(cmd, args, "distro")
	},
}

var distroListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all distributions",
	Long:  `Lists all available distributions.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		distroNames, err := Client.ListDistroNames()
		if err != nil {
			return err
		}
		listItems(cmd, "distros", distroNames)
		return nil
	},
}

var distroRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove distribution",
	Long:  `Removes a given distribution.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()

		dname, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		recursiveDelete, err := cmd.Flags().GetBool("recursive")
		if err != nil {
			return err
		}
		return Client.DeleteDistroRecursive(dname, recursiveDelete)
	},
}

var distroRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename distribution",
	Long:  `Renames a given distribution.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()

		// Get the name and newname flags
		distroName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		distroNewName, err := cmd.Flags().GetString("newname")
		if err != nil {
			return err
		}

		// Get the distro API handle
		distroHandle, err := Client.GetDistroHandle(distroName)
		if err != nil {
			return err
		}
		// Perform the rename operation server-side
		err = Client.RenameDistro(distroHandle, distroNewName)
		if err != nil {
			return err
		}
		// Retrieve the renamed distro from the API
		newDistro, err := Client.GetDistro(distroNewName, false, false)
		if err != nil {
			return err
		}
		// Now edit the distro in-memory
		err = updateDistroFromFlags(cmd, newDistro)
		if err != nil {
			return err
		}
		// Now update the distro via XML-RPC
		return Client.UpdateDistro(newDistro)
	},
}

func reportDistros(cmd *cobra.Command, distroNames []string) error {
	for _, itemName := range distroNames {
		distro, err := Client.GetDistro(itemName, false, false)
		if err != nil {
			return err
		}
		printStructured(cmd, distro)
		fmt.Fprintln(cmd.OutOrStdout(), "")
	}
	return nil
}

var distroReportCmd = &cobra.Command{
	Use:   "report",
	Short: "list all distributions in detail",
	Long:  `Shows detailed information about all distributions.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		itemNames := make([]string, 0)
		if name == "" {
			itemNames, err = Client.ListDistroNames()
			if err != nil {
				return err
			}
		} else {
			itemNames = append(itemNames, name)
		}
		return reportDistros(cmd, itemNames)
	},
}

func init() {
	rootCmd.AddCommand(distroCmd)
	distroCmd.AddCommand(distroAddCmd)
	distroCmd.AddCommand(distroCopyCmd)
	distroCmd.AddCommand(distroEditCmd)
	distroCmd.AddCommand(distroFindCmd)
	distroCmd.AddCommand(distroListCmd)
	distroCmd.AddCommand(distroRemoveCmd)
	distroCmd.AddCommand(distroRenameCmd)
	distroCmd.AddCommand(distroReportCmd)

	// local flags for distro add
	addCommonArgs(distroAddCmd)
	addStringFlags(distroAddCmd, distroStringFlagMetadata)
	addStringSliceFlags(distroAddCmd, distroStringSliceFlagMetadata)
	addMapFlags(distroAddCmd, distroMapFlagMetadata)
	// Required Flags
	err := distroAddCmd.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}

	// local flags for distro copy
	addCommonArgs(distroCopyCmd)
	addStringFlags(distroCopyCmd, distroStringFlagMetadata)
	addStringSliceFlags(distroCopyCmd, distroStringSliceFlagMetadata)
	addMapFlags(distroCopyCmd, distroMapFlagMetadata)
	distroCopyCmd.Flags().String("newname", "", "the new distro name")
	distroCopyCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")

	// local flags for distro edit
	addCommonArgs(distroEditCmd)
	addStringFlags(distroEditCmd, distroStringFlagMetadata)
	addStringSliceFlags(distroEditCmd, distroStringSliceFlagMetadata)
	addMapFlags(distroEditCmd, distroMapFlagMetadata)
	distroEditCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")

	// local flags for distro find
	addCommonArgs(distroFindCmd)
	addStringFlags(distroFindCmd, distroStringFlagMetadata)
	addStringSliceFlags(distroFindCmd, distroStringSliceFlagMetadata)
	addMapFlags(distroFindCmd, distroMapFlagMetadata)
	addStringFlags(distroFindCmd, findStringFlagMetadata)
	addIntFlags(distroFindCmd, findIntFlagMetadata)
	addFloatFlags(distroFindCmd, findFloatFlagMetadata)
	distroFindCmd.Flags().String("source-repos", "", "source repositories")
	distroFindCmd.Flags().String("tree-build-time", "", "tree build time")

	// local flags for distro remove
	distroRemoveCmd.Flags().String("name", "", "the distro name")
	distroRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")

	// local flags for distro rename
	addCommonArgs(distroRenameCmd)
	addStringFlags(distroRenameCmd, distroStringFlagMetadata)
	addStringSliceFlags(distroRenameCmd, distroStringSliceFlagMetadata)
	addMapFlags(distroRenameCmd, distroMapFlagMetadata)
	distroRenameCmd.Flags().String("newname", "", "the new distro name")
	distroRenameCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")

	// local flags for distro report
	distroReportCmd.Flags().String("name", "", "the distro name")
}
