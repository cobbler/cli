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

func updateRepoFromFlags(cmd *cobra.Command, repo *cobbler.Repo) error {
	// This object doesn't have the in-place flag
	var err error
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		if err != nil {
			// If one of the previous flags has had an error just directly return.
			return
		}
		switch flag.Name {
		case "comment":
			var repoNewComment string
			repoNewComment, err = cmd.Flags().GetString("comment")
			if err != nil {
				return
			}
			repo.Comment = repoNewComment
		case "arch":
			var repoNewArch string
			repoNewArch, err = cmd.Flags().GetString("arch")
			if err != nil {
				return
			}
			repo.Arch = repoNewArch
		case "breed":
			var repoNewBreed string
			repoNewBreed, err = cmd.Flags().GetString("breed")
			if err != nil {
				return
			}
			repo.Breed = repoNewBreed
		case "owners":
			fallthrough
		case "owners-inherit":
			if cmd.Flags().Lookup("owners-inherit").Changed {
				repo.Owners.Data = []string{}
				repo.Owners.IsInherited, err = cmd.Flags().GetBool("owners-inherit")
				if err != nil {
					return
				}
			} else {
				var repoNewOwners []string
				repoNewOwners, err = cmd.Flags().GetStringSlice("owners")
				if err != nil {
					return
				}
				repo.Owners.IsInherited = false
				repo.Owners.Data = repoNewOwners
			}
		case "apt-components":
			var repoNewAptComponents []string
			repoNewAptComponents, err = cmd.Flags().GetStringSlice("apt-components")
			if err != nil {
				return
			}
			repo.AptComponents = repoNewAptComponents
		case "apt-dists":
			var repoNewAptDists []string
			repoNewAptDists, err = cmd.Flags().GetStringSlice("apt-dists")
			if err != nil {
				return
			}
			repo.AptDists = repoNewAptDists
		case "createrepo-flags":
			fallthrough
		case "createrepo-flags-inherit":
			if cmd.Flags().Lookup("createrepo-flags-inherit").Changed {
				repo.CreateRepoFlags.Data = ""
				repo.CreateRepoFlags.IsInherited, err = cmd.Flags().GetBool("createrepo-flags-inherit")
				if err != nil {
					return
				}
			} else {
				var repoNewCreatrepoFlags string
				repoNewCreatrepoFlags, err = cmd.Flags().GetString("createrepo-flags")
				if err != nil {
					return
				}
				repo.CreateRepoFlags.IsInherited = false
				repo.CreateRepoFlags.Data = repoNewCreatrepoFlags
			}
		case "environment":
			var repoNewEnvironment map[string]string
			repoNewEnvironment, err = cmd.Flags().GetStringToString("environment")
			if err != nil {
				return
			}
			repo.Environment = repoNewEnvironment
		case "keep-updated":
			var repoNewKeepUpdated bool
			repoNewKeepUpdated, err = cmd.Flags().GetBool("keep-updated")
			if err != nil {
				return
			}
			repo.KeepUpdated = repoNewKeepUpdated
		case "mirror":
			var repoNewMirror string
			repoNewMirror, err = cmd.Flags().GetString("mirror")
			if err != nil {
				return
			}
			repo.Mirror = repoNewMirror
		case "mirror-type":
			var repoNewMirrorType string
			repoNewMirrorType, err = cmd.Flags().GetString("mirror-type")
			if err != nil {
				return
			}
			repo.MirrorType = repoNewMirrorType
		case "priority":
			var repoNewPriority int
			repoNewPriority, err = cmd.Flags().GetInt("priority")
			if err != nil {
				return
			}
			repo.Priority = repoNewPriority
		case "proxy":
			fallthrough
		case "proxy-inherit":
			if cmd.Flags().Lookup("proxy-inherit").Changed {
				repo.Proxy.Data = ""
				repo.Proxy.IsInherited, err = cmd.Flags().GetBool("proxy-inherit")
				if err != nil {
					return
				}
			} else {
				var repoNewProxy string
				repoNewProxy, err = cmd.Flags().GetString("proxy")
				if err != nil {
					return
				}
				repo.Proxy.IsInherited = false
				repo.Proxy.Data = repoNewProxy
			}
		case "rpm-list":
			var repoNewRpmList []string
			repoNewRpmList, err = cmd.Flags().GetStringSlice("rpm-list")
			if err != nil {
				return
			}
			repo.RpmList = repoNewRpmList
		case "yumopts":
			var repoNewYumOpts map[string]string
			repoNewYumOpts, err = cmd.Flags().GetStringToString("yumopts")
			if err != nil {
				return
			}
			repo.YumOpts = repoNewYumOpts
		case "rsyncopts":
			var repoNewRsyncOpts map[string]string
			repoNewRsyncOpts, err = cmd.Flags().GetStringToString("rsyncopts")
			if err != nil {
				return
			}
			repo.RsyncOpts = repoNewRsyncOpts
		}
	})
	return err
}

// NewRepoCmd builds a new command that represents the repo action
func NewRepoCmd() *cobra.Command {
	repoCmd := &cobra.Command{
		Use:   "repo",
		Short: "Repository management",
		Long: `Let you manage repositories.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-repo for more information.`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	repoCmd.AddCommand(NewRepoAddCmd())
	repoCmd.AddCommand(NewRepoAutoAddCmd())
	repoCmd.AddCommand(NewRepoCopyCmd())
	repoCmd.AddCommand(NewRepoEditCmd())
	repoCmd.AddCommand(NewRepoFindCmd())
	repoCmd.AddCommand(NewRepoListCmd())
	repoCmd.AddCommand(NewRepoRemoveCmd())
	repoCmd.AddCommand(NewRepoRenameCmd())
	repoCmd.AddCommand(NewRepoReportCmd())
	return repoCmd
}

func NewRepoAddCmd() *cobra.Command {
	repoAddCmd := &cobra.Command{
		Use:   "add",
		Short: "add repository",
		Long:  `Adds a repository.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			newRepo := cobbler.NewRepo()

			// internal fields (ctime, mtime, depth, uid, parent, tree-build-time) cannot be modified
			newRepo.Name, err = cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			// Update repo in-memory
			err = updateRepoFromFlags(cmd, &newRepo)
			if err != nil {
				return err
			}
			// Now create via XML-RPC
			repo, err := Client.CreateRepo(newRepo)
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Repo %s created\n", repo.Name)
			return nil
		},
	}
	addCommonArgs(repoAddCmd)
	addStringFlags(repoAddCmd, repoStringFlagMetadata)
	addBoolFlags(repoAddCmd, repoBoolFlagMetadata)
	addIntFlags(repoAddCmd, repoIntFlagMetadata)
	addStringSliceFlags(repoAddCmd, repoStringSliceFlagMetadata)
	addMapFlags(repoAddCmd, repoMapFlagMetadata)
	return repoAddCmd
}

func NewRepoAutoAddCmd() *cobra.Command {
	repoAutoAddCmd := &cobra.Command{
		Use:   "autoadd",
		Short: "add repository automatically",
		Long:  `Automatically adds a repository.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			return Client.AutoAddRepos()
		},
	}
	return repoAutoAddCmd
}

func NewRepoCopyCmd() *cobra.Command {
	repoCopyCmd := &cobra.Command{
		Use:   "copy",
		Short: "copy repository",
		Long:  `Copies a repository.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			repoName, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			repoNewName, err := cmd.Flags().GetString("newname")
			if err != nil {
				return err
			}

			repoHandle, err := Client.GetRepoHandle(repoName)
			if err != nil {
				return err
			}
			err = Client.CopyRepo(repoHandle, repoNewName)
			if err != nil {
				return err
			}
			copiedRepo, err := Client.GetRepo(repoNewName, false, false)
			if err != nil {
				return err
			}
			err = updateRepoFromFlags(cmd, copiedRepo)
			if err != nil {
				return err
			}
			return Client.UpdateRepo(copiedRepo)
		},
	}
	addCommonArgs(repoCopyCmd)
	addStringFlags(repoCopyCmd, repoStringFlagMetadata)
	addBoolFlags(repoCopyCmd, repoBoolFlagMetadata)
	addIntFlags(repoCopyCmd, repoIntFlagMetadata)
	addStringSliceFlags(repoCopyCmd, repoStringSliceFlagMetadata)
	addMapFlags(repoCopyCmd, repoMapFlagMetadata)
	repoCopyCmd.Flags().String("newname", "", "the new repo name")
	repoCopyCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	return repoCopyCmd
}

func NewRepoEditCmd() *cobra.Command {
	repoEditCmd := &cobra.Command{
		Use:   "edit",
		Short: "edit repository",
		Long:  `Edits a repository.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			// find repo through its name
			rname, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			// Get repo from API
			updateRepo, err := Client.GetRepo(rname, false, false)
			if err != nil {
				return err
			}
			// Update repo in-memory
			err = updateRepoFromFlags(cmd, updateRepo)
			if err != nil {
				return err
			}
			// Update repo via XML-RPC
			return Client.UpdateRepo(updateRepo)
		},
	}
	addCommonArgs(repoEditCmd)
	addStringFlags(repoEditCmd, repoStringFlagMetadata)
	addBoolFlags(repoEditCmd, repoBoolFlagMetadata)
	addIntFlags(repoEditCmd, repoIntFlagMetadata)
	addStringSliceFlags(repoEditCmd, repoStringSliceFlagMetadata)
	addMapFlags(repoEditCmd, repoMapFlagMetadata)
	repoEditCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	return repoEditCmd
}

func NewRepoFindCmd() *cobra.Command {
	repoFindCmd := &cobra.Command{
		Use:   "find",
		Short: "find repository",
		Long:  `Finds a given repository.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			return FindItemNames(cmd, args, "repo")
		},
	}
	addCommonArgs(repoFindCmd)
	addStringFlags(repoFindCmd, repoStringFlagMetadata)
	addBoolFlags(repoFindCmd, repoBoolFlagMetadata)
	addIntFlags(repoFindCmd, repoIntFlagMetadata)
	addStringSliceFlags(repoFindCmd, repoStringSliceFlagMetadata)
	addMapFlags(repoFindCmd, repoMapFlagMetadata)
	addStringFlags(repoFindCmd, findStringFlagMetadata)
	addIntFlags(repoFindCmd, findIntFlagMetadata)
	addFloatFlags(repoFindCmd, findFloatFlagMetadata)
	repoFindCmd.Flags().String("parent", "", "")
	return repoFindCmd
}

func NewRepoListCmd() *cobra.Command {
	repoListCmd := &cobra.Command{
		Use:   "list",
		Short: "list all repositorys",
		Long:  `Lists all available repositories.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			repoNames, err := Client.ListRepoNames()
			if err != nil {
				return err
			}
			listItems(cmd, "repos", repoNames)
			return nil
		},
	}
	return repoListCmd
}

func NewRepoRemoveCmd() *cobra.Command {
	repoRemoveCmd := &cobra.Command{
		Use:   "remove",
		Short: "remove repository",
		Long:  `Removes a given repository.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			return RemoveItemRecursive(cmd, args, "repo")
		},
	}
	repoRemoveCmd.Flags().String("name", "", "the repo name")
	repoRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")
	return repoRemoveCmd
}

func NewRepoRenameCmd() *cobra.Command {
	repoRenameCmd := &cobra.Command{
		Use:   "rename",
		Short: "rename repository",
		Long:  `Renames a given repository.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			// Get special name and newname flags
			repoName, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			repoNewName, err := cmd.Flags().GetString("newname")
			if err != nil {
				return err
			}
			// Get repo handle from the API
			repoHandle, err := Client.GetMenuHandle(repoName)
			if err != nil {
				return err
			}
			// Rename the repo server side
			err = Client.RenameRepo(repoHandle, repoNewName)
			if err != nil {
				return err
			}
			// Get the renamed repository from the API
			newRepository, err := Client.GetRepo(repoNewName, false, false)
			if err != nil {
				return err
			}
			// Update the repo in-memory
			err = updateRepoFromFlags(cmd, newRepository)
			if err != nil {
				return err
			}
			// Update the repo via XML-RPC
			return Client.UpdateRepo(newRepository)
		},
	}
	addCommonArgs(repoRenameCmd)
	addStringFlags(repoRenameCmd, repoStringFlagMetadata)
	addBoolFlags(repoRenameCmd, repoBoolFlagMetadata)
	addIntFlags(repoRenameCmd, repoIntFlagMetadata)
	addStringSliceFlags(repoRenameCmd, repoStringSliceFlagMetadata)
	addMapFlags(repoRenameCmd, repoMapFlagMetadata)
	repoRenameCmd.Flags().String("newname", "", "the new repo name")
	repoRenameCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	return repoRenameCmd
}

func reportRepos(cmd *cobra.Command, repoNames []string) error {
	for _, itemName := range repoNames {
		repo, err := Client.GetRepo(itemName, false, false)
		if err != nil {
			return err
		}
		printStructured(cmd, repo)
		fmt.Fprintln(cmd.OutOrStdout(), "")
	}
	return nil
}

func NewRepoReportCmd() *cobra.Command {
	repoReportCmd := &cobra.Command{
		Use:   "report",
		Short: "list all repositorys in detail",
		Long:  `Shows detailed information about all repositories.`,
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
				itemNames, err = Client.ListRepoNames()
				if err != nil {
					return err
				}
			} else {
				itemNames = append(itemNames, name)
			}
			return reportRepos(cmd, itemNames)
		},
	}
	repoReportCmd.Flags().String("name", "", "the repo name")
	return repoReportCmd
}
