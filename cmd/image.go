// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"github.com/spf13/cobra"
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Image management",
	Long: `Let you manage images.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-image for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var imageAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add image",
	Long:  `Adds a image.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var imageCopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy image",
	Long:  `Copies a image.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var imageEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit image",
	Long:  `Edits a image.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var imageFindCmd = &cobra.Command{
	Use:   "find",
	Short: "find image",
	Long:  `Finds a given image.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var imageListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all images",
	Long:  `Lists all available images.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var imageRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove image",
	Long:  `Removes a given image.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var imageRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename image",
	Long:  `Renames a given image.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var imageimagertCmd = &cobra.Command{
	Use:   "imagert",
	Short: "list all images in detail",
	Long:  `Shows detailed information about all images.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
	imageCmd.AddCommand(imageAddCmd)
	imageCmd.AddCommand(imageCopyCmd)
	imageCmd.AddCommand(imageEditCmd)
	imageCmd.AddCommand(imageFindCmd)
	imageCmd.AddCommand(imageListCmd)
	imageCmd.AddCommand(imageRemoveCmd)
	imageCmd.AddCommand(imageRenameCmd)
	imageCmd.AddCommand(imageimagertCmd)

	// local flags for image add
	imageAddCmd.Flags().String("name", "", "the image name")
	imageAddCmd.Flags().String("ctime", "", "")
	imageAddCmd.Flags().String("depth", "", "")
	imageAddCmd.Flags().String("mtime", "", "")
	imageAddCmd.Flags().String("uid", "", "UID")
	imageAddCmd.Flags().String("arch", "", "Architecture")
	imageAddCmd.Flags().String("breed", "", "Breed (valid options: none,rsync,rhn,yum,apt,wget)")
	imageAddCmd.Flags().String("comment", "", "free form text description")
	imageAddCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	imageAddCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	imageAddCmd.Flags().String("parent", "", "")
	imageAddCmd.Flags().String("file", "", "path to local file or nfs://user@host:path")
	imageAddCmd.Flags().String("image-type", "", "image type. Valid options: iso,direct,memdisk,virt-image")
	imageAddCmd.Flags().String("network-count", "", "")
	imageAddCmd.Flags().String("os-version", "", "OS version (needed for some virtualization optimizations)")
	imageAddCmd.Flags().String("menu", "", "parent boot menu")
	imageAddCmd.Flags().String("boot-loaders", "", "boot loaders (network installation boot loaders)")
	imageAddCmd.Flags().Bool("virt-auto-boot", false, "auto boot this VM?")
	imageAddCmd.Flags().String("virt-bridge", "", "virt bridge")
	imageAddCmd.Flags().String("virt-cpus", "", "virt CPUs")
	imageAddCmd.Flags().String("virt-disk-driver", "", "the on-disk format for the virtualization disk. Valid options: <<inherit>>,raw,qcow2,qed,vdi,vdmk")
	imageAddCmd.Flags().String("virt-file-size", "", "virt file size in GB")
	imageAddCmd.Flags().String("virt-path", "", "virt Path (e.g. /directory or VolGroup00)")
	imageAddCmd.Flags().String("virt-ram", "", "virt RAM size in MB")
	imageAddCmd.Flags().String("virt-type", "", "virtualization technology to use. Valid options: xenpv,xenfv,qemu,kvm,vmware")

	// local flags for image copy
	imageCopyCmd.Flags().String("name", "", "the image name")
	imageCopyCmd.Flags().String("newname", "", "the new image name")
	imageCopyCmd.Flags().String("ctime", "", "")
	imageCopyCmd.Flags().String("depth", "", "")
	imageCopyCmd.Flags().String("mtime", "", "")
	imageCopyCmd.Flags().String("uid", "", "UID")
	imageCopyCmd.Flags().String("arch", "", "Architecture")
	imageCopyCmd.Flags().String("breed", "", "Breed (valid options: none,rsync,rhn,yum,apt,wget)")
	imageCopyCmd.Flags().String("comment", "", "free form text description")
	imageCopyCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	imageCopyCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	imageCopyCmd.Flags().String("parent", "", "")
	imageCopyCmd.Flags().String("file", "", "path to local file or nfs://user@host:path")
	imageCopyCmd.Flags().String("image-type", "", "image type. Valid options: iso,direct,memdisk,virt-image")
	imageCopyCmd.Flags().String("network-count", "", "")
	imageCopyCmd.Flags().String("os-version", "", "OS version (needed for some virtualization optimizations)")
	imageCopyCmd.Flags().String("menu", "", "parent boot menu")
	imageCopyCmd.Flags().String("boot-loaders", "", "boot loaders (network installation boot loaders)")
	imageCopyCmd.Flags().Bool("virt-auto-boot", false, "auto boot this VM?")
	imageCopyCmd.Flags().String("virt-bridge", "", "virt bridge")
	imageCopyCmd.Flags().String("virt-cpus", "", "virt CPUs")
	imageCopyCmd.Flags().String("virt-disk-driver", "", "the on-disk format for the virtualization disk. Valid options: <<inherit>>,raw,qcow2,qed,vdi,vdmk")
	imageCopyCmd.Flags().String("virt-file-size", "", "virt file size in GB")
	imageCopyCmd.Flags().String("virt-path", "", "virt Path (e.g. /directory or VolGroup00)")
	imageCopyCmd.Flags().String("virt-ram", "", "virt RAM size in MB")
	imageCopyCmd.Flags().String("virt-type", "", "virtualization technology to use. Valid options: xenpv,xenfv,qemu,kvm,vmware")

	// local flags for image edit
	imageEditCmd.Flags().String("name", "", "the image name")
	imageEditCmd.Flags().String("ctime", "", "")
	imageEditCmd.Flags().String("depth", "", "")
	imageEditCmd.Flags().String("mtime", "", "")
	imageEditCmd.Flags().String("uid", "", "UID")
	imageEditCmd.Flags().String("arch", "", "Architecture")
	imageEditCmd.Flags().String("breed", "", "Breed (valid options: none,rsync,rhn,yum,apt,wget)")
	imageEditCmd.Flags().String("comment", "", "free form text description")
	imageEditCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	imageEditCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	imageEditCmd.Flags().String("parent", "", "")
	imageEditCmd.Flags().String("file", "", "path to local file or nfs://user@host:path")
	imageEditCmd.Flags().String("image-type", "", "image type. Valid options: iso,direct,memdisk,virt-image")
	imageEditCmd.Flags().String("network-count", "", "")
	imageEditCmd.Flags().String("os-version", "", "OS version (needed for some virtualization optimizations)")
	imageEditCmd.Flags().String("menu", "", "parent boot menu")
	imageEditCmd.Flags().String("boot-loaders", "", "boot loaders (network installation boot loaders)")
	imageEditCmd.Flags().Bool("virt-auto-boot", false, "auto boot this VM?")
	imageEditCmd.Flags().String("virt-bridge", "", "virt bridge")
	imageEditCmd.Flags().String("virt-cpus", "", "virt CPUs")
	imageEditCmd.Flags().String("virt-disk-driver", "", "the on-disk format for the virtualization disk. Valid options: <<inherit>>,raw,qcow2,qed,vdi,vdmk")
	imageEditCmd.Flags().String("virt-file-size", "", "virt file size in GB")
	imageEditCmd.Flags().String("virt-path", "", "virt Path (e.g. /directory or VolGroup00)")
	imageEditCmd.Flags().String("virt-ram", "", "virt RAM size in MB")
	imageEditCmd.Flags().String("virt-type", "", "virtualization technology to use. Valid options: xenpv,xenfv,qemu,kvm,vmware")

	// local flags for image find
	imageFindCmd.Flags().String("name", "", "the image name")
	imageFindCmd.Flags().String("ctime", "", "")
	imageFindCmd.Flags().String("depth", "", "")
	imageFindCmd.Flags().String("mtime", "", "")
	imageFindCmd.Flags().String("uid", "", "UID")
	imageFindCmd.Flags().String("arch", "", "Architecture")
	imageFindCmd.Flags().String("breed", "", "Breed (valid options: none,rsync,rhn,yum,apt,wget)")
	imageFindCmd.Flags().String("comment", "", "free form text description")
	imageFindCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	imageFindCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	imageFindCmd.Flags().String("parent", "", "")
	imageFindCmd.Flags().String("file", "", "path to local file or nfs://user@host:path")
	imageFindCmd.Flags().String("image-type", "", "image type. Valid options: iso,direct,memdisk,virt-image")
	imageFindCmd.Flags().String("network-count", "", "")
	imageFindCmd.Flags().String("os-version", "", "OS version (needed for some virtualization optimizations)")
	imageFindCmd.Flags().String("menu", "", "parent boot menu")
	imageFindCmd.Flags().String("boot-loaders", "", "boot loaders (network installation boot loaders)")
	imageFindCmd.Flags().Bool("virt-auto-boot", false, "auto boot this VM?")
	imageFindCmd.Flags().String("virt-bridge", "", "virt bridge")
	imageFindCmd.Flags().String("virt-cpus", "", "virt CPUs")
	imageFindCmd.Flags().String("virt-disk-driver", "", "the on-disk format for the virtualization disk. Valid options: <<inherit>>,raw,qcow2,qed,vdi,vdmk")
	imageFindCmd.Flags().String("virt-file-size", "", "virt file size in GB")
	imageFindCmd.Flags().String("virt-path", "", "virt Path (e.g. /directory or VolGroup00)")
	imageFindCmd.Flags().String("virt-ram", "", "virt RAM size in MB")
	imageFindCmd.Flags().String("virt-type", "", "virtualization technology to use. Valid options: xenpv,xenfv,qemu,kvm,vmware")

	// local flags for image remove
	imageRemoveCmd.Flags().String("name", "", "the image name")
	imageRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")

	// local flags for image rename
	imageRenameCmd.Flags().String("name", "", "the image name")
	imageRenameCmd.Flags().String("newname", "", "the new image name")
	imageRenameCmd.Flags().String("ctime", "", "")
	imageRenameCmd.Flags().String("depth", "", "")
	imageRenameCmd.Flags().String("mtime", "", "")
	imageRenameCmd.Flags().String("uid", "", "UID")
	imageRenameCmd.Flags().String("arch", "", "Architecture")
	imageRenameCmd.Flags().String("breed", "", "Breed (valid options: none,rsync,rhn,yum,apt,wget)")
	imageRenameCmd.Flags().String("comment", "", "free form text description")
	imageRenameCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	imageRenameCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	imageRenameCmd.Flags().String("parent", "", "")
	imageRenameCmd.Flags().String("file", "", "path to local file or nfs://user@host:path")
	imageRenameCmd.Flags().String("image-type", "", "image type. Valid options: iso,direct,memdisk,virt-image")
	imageRenameCmd.Flags().String("network-count", "", "")
	imageRenameCmd.Flags().String("os-version", "", "OS version (needed for some virtualization optimizations)")
	imageRenameCmd.Flags().String("menu", "", "parent boot menu")
	imageRenameCmd.Flags().String("boot-loaders", "", "boot loaders (network installation boot loaders)")
	imageRenameCmd.Flags().Bool("virt-auto-boot", false, "auto boot this VM?")
	imageRenameCmd.Flags().String("virt-bridge", "", "virt bridge")
	imageRenameCmd.Flags().String("virt-cpus", "", "virt CPUs")
	imageRenameCmd.Flags().String("virt-disk-driver", "", "the on-disk format for the virtualization disk. Valid options: <<inherit>>,raw,qcow2,qed,vdi,vdmk")
	imageRenameCmd.Flags().String("virt-file-size", "", "virt file size in GB")
	imageRenameCmd.Flags().String("virt-path", "", "virt Path (e.g. /directory or VolGroup00)")
	imageRenameCmd.Flags().String("virt-ram", "", "virt RAM size in MB")
	imageRenameCmd.Flags().String("virt-type", "", "virtualization technology to use. Valid options: xenpv,xenfv,qemu,kvm,vmware")

	// local flags for image imagert
	imageimagertCmd.Flags().String("name", "", "the image name")
}
