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

func updateImageFromFlags(cmd *cobra.Command, image *cobbler.Image) error {
	inPlace, err := cmd.Flags().GetBool("in-place")
	if err != nil {
		return err
	}
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		switch flag.Name {
		// The rename & copy operations are special operations as such we cannot blindly set this inside here.
		// Any rename & copy operation must be handled outside of this method.
		case "kernel-options":
			fallthrough
		case "kernel-options-inherit":
			if cmd.Flags().Lookup("kernel-options-inherit").Changed {
				image.KernelOptions.Data = make(map[string]interface{})
				image.KernelOptions.IsInherited, err = cmd.Flags().GetBool("kernel-options-inherit")
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
						"image",
						image.Name,
						"kernel_options",
						convertMapStringToMapInterface(newKernelOptions),
					)
					if err != nil {
						return
					}
				} else {
					image.KernelOptions.IsInherited = false
					image.KernelOptions.Data = convertMapStringToMapInterface(newKernelOptions)
				}
			}
		case "kernel-options-post":
			fallthrough
		case "kernel-options-post-inherit":
			if cmd.Flags().Lookup("kernel-options-post-inherit").Changed {
				image.KernelOptionsPost.Data = make(map[string]interface{})
				image.KernelOptionsPost.IsInherited, err = cmd.Flags().GetBool("kernel-options-post-inherit")
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
						"image",
						image.Name,
						"kernel_options_post",
						convertMapStringToMapInterface(newKernelOptionsPost),
					)
					if err != nil {
						return
					}
				} else {
					image.KernelOptionsPost.IsInherited = false
					image.KernelOptions.Data = convertMapStringToMapInterface(newKernelOptionsPost)
				}
			}
		case "autoinstall-meta":
			fallthrough
		case "autoinstall-meta-inherit":
			if cmd.Flags().Lookup("boot-loaders-inherit").Changed {
				image.AutoinstallMeta.Data = make(map[string]interface{})
				image.AutoinstallMeta.IsInherited, err = cmd.Flags().GetBool("boot-loaders-inherit")
				if err != nil {
					return
				}
			} else {
				var profileNewAutoinstallMeta map[string]string
				profileNewAutoinstallMeta, err = cmd.Flags().GetStringToString("autoinstall-meta")
				if err != nil {
					return
				}
				image.AutoinstallMeta.IsInherited = false
				image.AutoinstallMeta.Data = convertMapStringToMapInterface(profileNewAutoinstallMeta)
			}
		case "template-files":
			fallthrough
		case "template-files-inherit":
			if cmd.Flags().Lookup("template-files-inherit").Changed {
				image.TemplateFiles.Data = make(map[string]interface{})
				image.TemplateFiles.IsInherited, err = cmd.Flags().GetBool("template-files-inherit")
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
						"image",
						image.Name,
						"template_files",
						convertMapStringToMapInterface(newTemplateFiles),
					)
					if err != nil {
						return
					}
				} else {
					image.TemplateFiles.IsInherited = false
					image.TemplateFiles.Data = convertMapStringToMapInterface(newTemplateFiles)
				}
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
				image.BootFiles.Data = make(map[string]interface{})
				image.BootFiles.IsInherited, err = cmd.Flags().GetBool("boot-files-inherit")
				if err != nil {
					return
				}
			} else {
				image.BootFiles.IsInherited = false
				image.BootFiles.Data = convertMapStringToMapInterface(distroNewBootFiles)
			}
		case "fetchable-files":
			fallthrough
		case "fetchable-files-inherit":
			var newFetchableFiles map[string]string
			newFetchableFiles, err = cmd.Flags().GetStringToString("fetchable-files")
			if err != nil {
				return
			}
			if cmd.Flags().Lookup("fetchable-files-inherit").Changed {
				image.FetchableFiles.Data = make(map[string]interface{})
				image.FetchableFiles.IsInherited, err = cmd.Flags().GetBool("fetchable-files-inherit")
				if err != nil {
					return
				}
			} else {
				if inPlace {
					err = Client.ModifyItemInPlace(
						"image",
						image.Name,
						"fetchable_files",
						convertMapStringToMapInterface(newFetchableFiles),
					)
					if err != nil {
						return
					}
				} else {
					image.FetchableFiles.IsInherited = false
					image.FetchableFiles.Data = convertMapStringToMapInterface(newFetchableFiles)
				}
			}
		case "comment":
			var imageNewComment string
			imageNewComment, err = cmd.Flags().GetString("comment")
			if err != nil {
				return
			}
			image.Comment = imageNewComment
		case "arch":
			var imageNewArch string
			imageNewArch, err = cmd.Flags().GetString("arch")
			if err != nil {
				return
			}
			image.Arch = imageNewArch
		case "breed":
			var imageNewBreed string
			imageNewBreed, err = cmd.Flags().GetString("breed")
			if err != nil {
				return
			}
			image.Breed = imageNewBreed
		case "owners":
			fallthrough
		case "owners-inherit":
			if cmd.Flags().Lookup("owners-inherit").Changed {
				image.Owners.Data = []string{}
				image.Owners.IsInherited, err = cmd.Flags().GetBool("owners-inherit")
				if err != nil {
					return
				}
			} else {
				var imageNewOwners []string
				imageNewOwners, err = cmd.Flags().GetStringSlice("owners")
				if err != nil {
					return
				}
				image.Owners.IsInherited = false
				image.Owners.Data = imageNewOwners
			}
		case "parent":
			var imageNewParent string
			imageNewParent, err = cmd.Flags().GetString("parent")
			if err != nil {
				return
			}
			image.Parent = imageNewParent
		case "file":
			var imageNewFile string
			imageNewFile, err = cmd.Flags().GetString("file")
			if err != nil {
				return
			}
			image.File = imageNewFile
		case "image-type":
			var imageNewImageType string
			imageNewImageType, err = cmd.Flags().GetString("image-type")
			if err != nil {
				return
			}
			image.ImageType = imageNewImageType
		case "network-count":
			var imageNewNetworkCount int
			imageNewNetworkCount, err = cmd.Flags().GetInt("network-count")
			if err != nil {
				return
			}
			image.NetworkCount = imageNewNetworkCount
		case "os-version":
			var imageNewOsVersion string
			imageNewOsVersion, err = cmd.Flags().GetString("os-version")
			if err != nil {
				return
			}
			image.OsVersion = imageNewOsVersion
		case "menu":
			var imageNewMenu string
			imageNewMenu, err = cmd.Flags().GetString("menu")
			if err != nil {
				return
			}
			image.Menu = imageNewMenu
		case "boot-loaders":
			var imageNewBootLoaders []string
			imageNewBootLoaders, err = cmd.Flags().GetStringSlice("boot-loaders")
			if err != nil {
				return
			}
			image.BootLoaders = imageNewBootLoaders
		case "virt-auto-boot":
			var imageNewVirtAutoBoot bool
			imageNewVirtAutoBoot, err = cmd.Flags().GetBool("virt-auto-boot")
			if err != nil {
				return
			}
			image.VirtAutoBoot = imageNewVirtAutoBoot
		case "virt-bridge":
			var imageNewVirtBridge string
			imageNewVirtBridge, err = cmd.Flags().GetString("virt-bridge")
			if err != nil {
				return
			}
			image.VirtBridge = imageNewVirtBridge
		case "virt-cpus":
			var imageNewVirtCpus int
			imageNewVirtCpus, err = cmd.Flags().GetInt("virt-cpus")
			if err != nil {
				return
			}
			image.VirtCpus = imageNewVirtCpus
		case "virt-disk-driver":
			var imageNewVirtDiskDriver string
			imageNewVirtDiskDriver, err = cmd.Flags().GetString("virt-disk-driver")
			if err != nil {
				return
			}
			image.VirtDiskDriver = imageNewVirtDiskDriver
		case "virt-file-size":
			fallthrough
		case "virt-file-size-inherit":
			if cmd.Flags().Lookup("owners-inherit").Changed {
				image.VirtFileSize.Data = 0
				image.VirtFileSize.IsInherited, err = cmd.Flags().GetBool("owners-inherit")
				if err != nil {
					return
				}
			} else {
				var imageNewVirtFileSize float64
				imageNewVirtFileSize, err = cmd.Flags().GetFloat64("virt-file-size")
				if err != nil {
					return
				}
				image.VirtFileSize.IsInherited = false
				image.VirtFileSize.Data = imageNewVirtFileSize
			}
		case "virt-path":
			var imageNewVirtPath string
			imageNewVirtPath, err = cmd.Flags().GetString("virt-path")
			if err != nil {
				return
			}
			image.VirtPath = imageNewVirtPath
		case "virt-ram":
			fallthrough
		case "virt-ram-inherit":
			if cmd.Flags().Lookup("owners-inherit").Changed {
				image.VirtRam.Data = 0
				image.VirtRam.IsInherited, err = cmd.Flags().GetBool("owners-inherit")
				if err != nil {
					return
				}
			} else {
				var imageNewVirtRam int
				imageNewVirtRam, err = cmd.Flags().GetInt("virt-ram")
				if err != nil {
					return
				}
				image.VirtRam.IsInherited = false
				image.VirtRam.Data = imageNewVirtRam
			}
		case "virt-type":
			var imageNewVirtType string
			imageNewVirtType, err = cmd.Flags().GetString("virt-type")
			if err != nil {
				return
			}
			image.VirtType = imageNewVirtType
		}
	})
	return err
}

// NewImageCmd builds a new command that represents the image action
func NewImageCmd() *cobra.Command {
	imageCmd := &cobra.Command{
		Use:   "image",
		Short: "Image management",
		Long: `Let you manage images.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-image for more information.`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	imageCmd.AddCommand(NewImageAddCmd())
	imageCmd.AddCommand(NewImageCopyCmd())
	imageCmd.AddCommand(NewImageEditCmd())
	imageCmd.AddCommand(NewImageFindCmd())
	imageCmd.AddCommand(NewImageListCmd())
	imageCmd.AddCommand(NewImageRemoveCmd())
	imageCmd.AddCommand(NewImageRenameCmd())
	imageCmd.AddCommand(NewImageReportCmd())
	return imageCmd
}

func NewImageAddCmd() *cobra.Command {
	imageAddCmd := &cobra.Command{
		Use:   "add",
		Short: "add image",
		Long:  `Adds a image.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			newImage := cobbler.NewImage()
			newImage.Name, err = cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			// Update image in-memory
			err = updateImageFromFlags(cmd, &newImage)
			if err != nil {
				return err
			}
			// Now create the image via XML-RPC
			system, err := Client.CreateImage(newImage)
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Image %s created\n", system.Name)
			return nil
		},
	}
	addCommonArgs(imageAddCmd)
	addStringFlags(imageAddCmd, imageStringFlagMetadata)
	addIntFlags(imageAddCmd, imageIntFlagMetadata)
	addFloatFlags(imageAddCmd, imageFloatFlagMetadata)
	addBoolFlags(imageAddCmd, imageBoolFlagMetadata)
	addStringSliceFlags(imageAddCmd, imageStringSliceFlagMetadata)
	imageAddCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	return imageAddCmd
}

func NewImageCopyCmd() *cobra.Command {
	imageCopyCmd := &cobra.Command{
		Use:   "copy",
		Short: "copy image",
		Long:  `Copies a image.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			imageName, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			imageNewName, err := cmd.Flags().GetString("newname")
			if err != nil {
				return err
			}

			imageHandle, err := Client.GetImageHandle(imageName)
			if err != nil {
				return err
			}
			err = Client.CopyImage(imageHandle, imageNewName)
			if err != nil {
				return err
			}
			copiedImage, err := Client.GetImage(imageNewName, false, false)
			if err != nil {
				return err
			}
			// Update image in-memory
			err = updateImageFromFlags(cmd, copiedImage)
			if err != nil {
				return err
			}
			return Client.UpdateImage(copiedImage)
		},
	}
	addCommonArgs(imageCopyCmd)
	addStringFlags(imageCopyCmd, imageStringFlagMetadata)
	addIntFlags(imageCopyCmd, imageIntFlagMetadata)
	addFloatFlags(imageCopyCmd, imageFloatFlagMetadata)
	addBoolFlags(imageCopyCmd, imageBoolFlagMetadata)
	addStringSliceFlags(imageCopyCmd, imageStringSliceFlagMetadata)
	imageCopyCmd.Flags().String("newname", "", "the new image name")
	imageCopyCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	return imageCopyCmd
}

func NewImageEditCmd() *cobra.Command {
	imageEditCmd := &cobra.Command{
		Use:   "edit",
		Short: "edit image",
		Long:  `Edits a image.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			imageName, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}

			imageToEdit, err := Client.GetImage(imageName, false, false)
			if err != nil {
				return err
			}
			// Update image in-memory
			err = updateImageFromFlags(cmd, imageToEdit)
			if err != nil {
				return err
			}
			return Client.UpdateImage(imageToEdit)
		},
	}
	addCommonArgs(imageEditCmd)
	addStringFlags(imageEditCmd, imageStringFlagMetadata)
	addIntFlags(imageEditCmd, imageIntFlagMetadata)
	addFloatFlags(imageEditCmd, imageFloatFlagMetadata)
	addBoolFlags(imageEditCmd, imageBoolFlagMetadata)
	addStringSliceFlags(imageEditCmd, imageStringSliceFlagMetadata)
	imageEditCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	return imageEditCmd
}

func NewImageFindCmd() *cobra.Command {
	imageFindCmd := &cobra.Command{
		Use:   "find",
		Short: "find image",
		Long:  `Finds a given image.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}
			return FindItemNames(cmd, args, "image")
		},
	}
	addCommonArgs(imageFindCmd)
	addStringFlags(imageFindCmd, imageStringFlagMetadata)
	addIntFlags(imageFindCmd, imageIntFlagMetadata)
	addFloatFlags(imageFindCmd, imageFloatFlagMetadata)
	addBoolFlags(imageFindCmd, imageBoolFlagMetadata)
	addStringSliceFlags(imageFindCmd, imageStringSliceFlagMetadata)
	addStringFlags(imageFindCmd, findStringFlagMetadata)
	addIntFlags(imageFindCmd, findIntFlagMetadata)
	addFloatFlags(imageFindCmd, findFloatFlagMetadata)
	return imageFindCmd
}

func NewImageListCmd() *cobra.Command {
	imageListCmd := &cobra.Command{
		Use:   "list",
		Short: "list all images",
		Long:  `Lists all available images.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			imageNames, err := Client.ListImageNames()
			if err != nil {
				return err
			}
			listItems(cmd, "images", imageNames)
			return nil
		},
	}
	return imageListCmd
}

func NewImageRemoveCmd() *cobra.Command {
	imageRemoveCmd := &cobra.Command{
		Use:   "remove",
		Short: "remove image",
		Long:  `Removes a given image.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			return RemoveItemRecursive(cmd, args, "image")
		},
	}
	imageRemoveCmd.Flags().String("name", "", "the image name")
	imageRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")
	return imageRemoveCmd
}

func NewImageRenameCmd() *cobra.Command {
	imageRenameCmd := &cobra.Command{
		Use:   "rename",
		Short: "rename image",
		Long:  `Renames a given image.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			imageName, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			imageNewName, err := cmd.Flags().GetString("newname")
			if err != nil {
				return err
			}

			imageHandle, err := Client.GetImageHandle(imageName)
			if err != nil {
				return err
			}
			err = Client.RenameImage(imageHandle, imageNewName)
			if err != nil {
				return err
			}
			renamedImage, err := Client.GetImage(imageNewName, false, false)
			if err != nil {
				return err
			}
			// Update image in-memory
			err = updateImageFromFlags(cmd, renamedImage)
			if err != nil {
				return err
			}
			return Client.UpdateImage(renamedImage)
		},
	}
	addCommonArgs(imageRenameCmd)
	addStringFlags(imageRenameCmd, imageStringFlagMetadata)
	addIntFlags(imageRenameCmd, imageIntFlagMetadata)
	addFloatFlags(imageRenameCmd, imageFloatFlagMetadata)
	addBoolFlags(imageRenameCmd, imageBoolFlagMetadata)
	addStringSliceFlags(imageRenameCmd, imageStringSliceFlagMetadata)
	imageRenameCmd.Flags().String("newname", "", "the new image name")
	imageRenameCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	return imageRenameCmd
}

func reportImages(cmd *cobra.Command, imageNames []string) error {
	for _, itemName := range imageNames {
		system, err := Client.GetImage(itemName, false, false)
		if err != nil {
			return err
		}
		printStructured(cmd, system)
		fmt.Fprintln(cmd.OutOrStdout(), "")
	}
	return nil
}

func NewImageReportCmd() *cobra.Command {
	imageReportCmd := &cobra.Command{
		Use:   "report",
		Short: "list all images in detail",
		Long:  `Shows detailed information about all images.`,
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
				itemNames, err = Client.ListImageNames()
				if err != nil {
					return err
				}
			} else {
				itemNames = append(itemNames, name)
			}
			return reportImages(cmd, itemNames)
		},
	}
	imageReportCmd.Flags().String("name", "", "the image name")
	return imageReportCmd
}
