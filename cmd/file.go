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

func updateFileFromFlags(cmd *cobra.Command, file *cobbler.File) error {
	// This object type doesn't have the in-place flag
	var err error
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		if err != nil {
			// If one of the previous flags has had an error just directly return.
			return
		}
		switch flag.Name {
		// The rename & copy operations are special operations as such we cannot blindly set this inside here.
		// Any rename & copy operation must be handled outside of this method.
		case "comment":
			var fileNewComment string
			fileNewComment, err = cmd.Flags().GetString("comment")
			if err != nil {
				return
			}
			file.Comment = fileNewComment
		case "owners":
			fallthrough
		case "owners-inherit":
			var fileNewOwners []string
			fileNewOwners, err = cmd.Flags().GetStringSlice("owners")
			if err != nil {
				return
			}
			if cmd.Flags().Lookup("owners-inherit").Changed {
				file.Owners.Data = []string{}
				file.Owners.IsInherited, err = cmd.Flags().GetBool("owners-inherit")
				if err != nil {
					return
				}
			} else {
				file.Owners.IsInherited = false
				file.Owners.Data = fileNewOwners
			}
		case "action":
			var fileNewAction string
			fileNewAction, err = cmd.Flags().GetString("action")
			if err != nil {
				return
			}
			file.Action = fileNewAction
		case "mode":
			var fileNewMode string
			fileNewMode, err = cmd.Flags().GetString("mode")
			if err != nil {
				return
			}
			file.Mode = fileNewMode
		case "template":
			var fileNewTemplate string
			fileNewTemplate, err = cmd.Flags().GetString("template")
			if err != nil {
				return
			}
			file.Template = fileNewTemplate
		case "path":
			var fileNewPath string
			fileNewPath, err = cmd.Flags().GetString("path")
			if err != nil {
				return
			}
			file.Path = fileNewPath
		case "group":
			var fileNewGroup string
			fileNewGroup, err = cmd.Flags().GetString("group")
			if err != nil {
				return
			}
			file.Group = fileNewGroup
		case "owner":
			var fileNewOwner string
			fileNewOwner, err = cmd.Flags().GetString("owner")
			if err != nil {
				return
			}
			file.Owner = fileNewOwner
		case "is-dir":
			var fileNewIsDir bool
			fileNewIsDir, err = cmd.Flags().GetBool("is-dir")
			if err != nil {
				return
			}
			file.IsDir = fileNewIsDir
		}
	})
	// Don't blindly return nil because maybe one of the flags had an issue retrieving an argument.
	return err
}

// NewFileCmd builds a new command that represents the file action
func NewFileCmd() (*cobra.Command, error) {
	fileCmd := &cobra.Command{
		Use:   "file",
		Short: "File management",
		Long: `Let you manage files.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-file for more information.`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	fileCmd.AddCommand(NewFileAddCmd())
	fileCmd.AddCommand(NewFileCopyCmd())
	fileCmd.AddCommand(NewFileEditCmd())
	fileCmd.AddCommand(NewFileFindCmd())
	fileCmd.AddCommand(NewFileListCmd())
	fileCmd.AddCommand(NewFileRemoveCmd())
	fileCmd.AddCommand(NewFileRenameCmd())
	fileCmd.AddCommand(NewFileReportCmd())
	fileCmd.AddCommand(NewFileExportCmd())
	return fileCmd, nil
}

func NewFileAddCmd() *cobra.Command {
	fileAddCmd := &cobra.Command{
		Use:   "add",
		Short: "add file",
		Long:  `Adds a file.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			newFile := cobbler.NewFile()

			// Get special name flag
			newFile.Name, err = cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			// Update with the rest of the flags
			err = updateFileFromFlags(cmd, &newFile)
			if err != nil {
				return err
			}
			// Now create the file via XML-RPC
			file, err := Client.CreateFile(newFile)
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "File %s created\n", file.Name)
			return nil
		},
	}
	addCommonArgs(fileAddCmd)
	addStringFlags(fileAddCmd, fileStringFlagMetadata)
	addBoolFlags(fileAddCmd, fileBoolFlagMetadata)
	return fileAddCmd
}

func NewFileCopyCmd() *cobra.Command {
	fileCopyCmd := &cobra.Command{
		Use:   "copy",
		Short: "copy file",
		Long:  `Copies a file.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			// Get special name and newname flags
			fileName, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			fileNewName, err := cmd.Flags().GetString("newname")
			if err != nil {
				return err
			}

			// Now copy the file
			fileHandle, err := Client.GetFileHandle(fileName)
			if err != nil {
				return err
			}
			err = Client.CopyFile(fileHandle, fileNewName)
			if err != nil {
				return err
			}
			newFile, err := Client.GetFile(fileNewName, false, false)
			if err != nil {
				return err
			}
			err = updateFileFromFlags(cmd, newFile)
			if err != nil {
				return err
			}
			return Client.UpdateFile(newFile)
		},
	}
	addCommonArgs(fileCopyCmd)
	addStringFlags(fileCopyCmd, fileStringFlagMetadata)
	addBoolFlags(fileCopyCmd, fileBoolFlagMetadata)
	fileCopyCmd.Flags().String("newname", "", "the new file name")
	return fileCopyCmd
}

func NewFileEditCmd() *cobra.Command {
	fileEditCmd := &cobra.Command{
		Use:   "edit",
		Short: "edit file",
		Long:  `Edits a file.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			// Get the file name
			fileName, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}

			// Now get the file from the API
			newFile, err := Client.GetFile(fileName, false, false)
			if err != nil {
				return err
			}
			// Update the file in-memory
			err = updateFileFromFlags(cmd, newFile)
			if err != nil {
				return err
			}
			// Now update the file via XML-RPC
			return Client.UpdateFile(newFile)
		},
	}
	addCommonArgs(fileEditCmd)
	addStringFlags(fileEditCmd, fileStringFlagMetadata)
	addBoolFlags(fileEditCmd, fileBoolFlagMetadata)
	return fileEditCmd
}

func NewFileFindCmd() *cobra.Command {
	fileFindCmd := &cobra.Command{
		Use:   "find",
		Short: "find file",
		Long:  `Finds a given file.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			return FindItemNames(cmd, args, "file")
		},
	}
	addCommonArgs(fileFindCmd)
	addStringFlags(fileFindCmd, fileStringFlagMetadata)
	addBoolFlags(fileFindCmd, fileBoolFlagMetadata)
	addStringFlags(fileFindCmd, findStringFlagMetadata)
	addIntFlags(fileFindCmd, findIntFlagMetadata)
	addFloatFlags(fileFindCmd, findFloatFlagMetadata)
	return fileFindCmd
}

func NewFileListCmd() *cobra.Command {
	fileListCmd := &cobra.Command{
		Use:   "list",
		Short: "list all files",
		Long:  `Lists all available files.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			fileNames, err := Client.ListFileNames()
			if err != nil {
				return err
			}
			listItems(cmd, "files", fileNames)
			return nil
		},
	}
	return fileListCmd
}

func NewFileRemoveCmd() *cobra.Command {
	fileRemoveCmd := &cobra.Command{
		Use:   "remove",
		Short: "remove file",
		Long:  `Removes a given file.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			return RemoveItemRecursive(cmd, args, "file")
		},
	}
	fileRemoveCmd.Flags().String("name", "", "the file name")
	fileRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")
	return fileRemoveCmd
}

func NewFileRenameCmd() *cobra.Command {
	fileRenameCmd := &cobra.Command{
		Use:   "rename",
		Short: "rename file",
		Long:  `Renames a given file.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			// Get the special name and newname flags
			fileName, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			fileNewName, err := cmd.Flags().GetString("newname")
			if err != nil {
				return err
			}

			// Get the file handle
			fileHandle, err := Client.GetFileHandle(fileName)
			if err != nil {
				return err
			}
			// Rename the file (server-side)
			err = Client.RenameFile(fileHandle, fileNewName)
			if err != nil {
				return err
			}
			// Get the renamed file from the API
			newFile, err := Client.GetFile(fileNewName, false, false)
			if err != nil {
				return err
			}
			// Update the file in-memory
			err = updateFileFromFlags(cmd, newFile)
			if err != nil {
				return err
			}
			// Update the file via XML-RPC
			return Client.UpdateFile(newFile)
		},
	}
	addCommonArgs(fileRenameCmd)
	addStringFlags(fileRenameCmd, fileStringFlagMetadata)
	addBoolFlags(fileRenameCmd, fileBoolFlagMetadata)
	fileRenameCmd.Flags().String("newname", "", "the new file name")
	return fileRenameCmd
}

func reportFiles(cmd *cobra.Command, fileNames []string) error {
	for _, itemName := range fileNames {
		file, err := Client.GetFile(itemName, false, false)
		if err != nil {
			return err
		}
		printStructured(cmd, file)
		fmt.Fprintln(cmd.OutOrStdout(), "")
	}
	return nil
}

func NewFileReportCmd() *cobra.Command {
	fileReportCmd := &cobra.Command{
		Use:   "report",
		Short: "list all files in detail",
		Long:  `Shows detailed information about all files.`,
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
				itemNames, err = Client.ListFileNames()
				if err != nil {
					return err
				}
			} else {
				itemNames = append(itemNames, name)
			}
			return reportFiles(cmd, itemNames)
		},
	}
	fileReportCmd.Flags().String("name", "", "the file name")
	return fileReportCmd
}

func NewFileExportCmd() *cobra.Command {
	fileExportCmd := &cobra.Command{
		Use:   "export",
		Short: "export files",
		Long:  `Export files.`,
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
				itemNames, err = Client.ListFileNames()
				if err != nil {
					return err
				}
			} else {
				itemNames = append(itemNames, name)
			}

			for _, itemName := range itemNames {
				file, err := Client.GetFile(itemName, false, false)
				if err != nil {
					return err
				}
				if formatOption == "json" {
					jsonDocument, err := json.Marshal(file)
					if err != nil {
						return err
					}
					fmt.Fprintln(cmd.OutOrStdout(), string(jsonDocument))
				}
				if formatOption == "yaml" {
					yamlDocument, err := yaml.Marshal(file)
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
	fileExportCmd.Flags().String("name", "", "the file name")
	fileExportCmd.Flags().String(exportStringMetadata["format"].Name, exportStringMetadata["format"].DefaultValue, exportStringMetadata["format"].Usage)
	return fileExportCmd
}
