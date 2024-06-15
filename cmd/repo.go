// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	cobbler "github.com/cobbler/cobblerclient"
)

var repo *cobbler.Repo //nolint:golint,unused
var repos []*cobbler.Repo

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Repository management",
	Long: `Let you manage repositories.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-repo for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var repoAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add repository",
	Long:  `Adds a repository.`,
	Run: func(cmd *cobra.Command, args []string) {

		var newRepo cobbler.Repo

		// internal fields (ctime, mtime, depth, uid, parent, tree-build-time) cannot be modified
		newRepo.AptComponents, _ = cmd.Flags().GetStringArray("apt-components")
		newRepo.AptDists, _ = cmd.Flags().GetStringArray("apt-dists")
		newRepo.Arch, _ = cmd.Flags().GetString("arch")
		newRepo.Breed, _ = cmd.Flags().GetString("breed")
		newRepo.Comment, _ = cmd.Flags().GetString("comment")
		newRepo.CreateRepoFlags, _ = cmd.Flags().GetString("createrepo-flags")
		newRepo.Environment, _ = cmd.Flags().GetStringArray("environment")
		newRepo.KeepUpdated, _ = cmd.Flags().GetBool("keep-updated")
		newRepo.Mirror, _ = cmd.Flags().GetString("mirror")
		// not implemented in Cobbler yet
		// newRepo.MirrorLocally, _ = cmd.Flags().GetBool("mirror-locally")
		newRepo.Name, _ = cmd.Flags().GetString("name")
		newRepo.Owners, _ = cmd.Flags().GetStringArray("owners")
		newRepo.Proxy, _ = cmd.Flags().GetString("proxy")
		newRepo.RpmList, _ = cmd.Flags().GetStringArray("rpm-list")

		repo, err = Client.CreateRepo(newRepo)

		if checkError(err) == nil {
			fmt.Printf("Repo %s created", newRepo.Name)
		} else {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

var repoAutoAddCmd = &cobra.Command{
	Use:   "autoadd",
	Short: "add repository automatically",
	Long:  `Automatically adds a repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var repoCopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy repository",
	Long:  `Copies a repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var repoEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit repository",
	Long:  `Edits a repository.`,
	Run: func(cmd *cobra.Command, args []string) {

		// find repo through its name
		rname, _ := cmd.Flags().GetString("name")
		var updateRepo, err = Client.GetRepo(rname)

		if checkError(err) != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		// internal fields (ctime, mtime, depth, uid, parent, tree-build-time) cannot be modified
		var tmpArgs, _ = cmd.Flags().GetString("apt-components")
		if tmpArgs != "" {
			updateRepo.AptDists, _ = cmd.Flags().GetStringArray("apt-components")
		}
		tmpArgs, _ = cmd.Flags().GetString("apt-dists")
		if tmpArgs != "" {
			updateRepo.AptDists, _ = cmd.Flags().GetStringArray("apt-dists")
		}
		tmpArgs, _ = cmd.Flags().GetString("arch")
		if tmpArgs != "" {
			updateRepo.Arch, _ = cmd.Flags().GetString("arch")
		}
		tmpArgs, _ = cmd.Flags().GetString("breed")
		if tmpArgs != "" {
			updateRepo.Breed, _ = cmd.Flags().GetString("breed")
		}
		tmpArgs, _ = cmd.Flags().GetString("comment")
		if tmpArgs != "" {
			updateRepo.Comment, _ = cmd.Flags().GetString("comment")
		}
		tmpArgs, _ = cmd.Flags().GetString("createrepo-flags")
		if tmpArgs != "" {
			updateRepo.CreateRepoFlags, _ = cmd.Flags().GetString("createrepo-flags")
		}
		var tmpArgsArray, _ = cmd.Flags().GetStringArray("environment")
		if len(tmpArgsArray) > 0 {
			updateRepo.Environment, _ = cmd.Flags().GetStringArray("environment")
		}
		// TODO
		/* 		tmpArgs, _ = cmd.Flags().GetBool("keep-updated")
		   		if tmpArgs != "" {
		   			updateRepo.KeepUpdated, _ = cmd.Flags().GetBool("keep-updated")
		   		}
		*/
		tmpArgs, _ = cmd.Flags().GetString("mirror")
		if tmpArgs != "" {
			updateRepo.Mirror, _ = cmd.Flags().GetString("mirror")
		}
		// not implemented in Cobbler yet
		/* 		tmpArgs, _ = cmd.Flags().GetBool("mirror-locally")
		   		if tmpArgs != "" {
		   			updateRepo.KeepUpdated, _ = cmd.Flags().GetBool("mirror-locally")
		   		}
		*/
		tmpArgs, _ = cmd.Flags().GetString("name")
		if tmpArgs != "" {
			updateRepo.Name, _ = cmd.Flags().GetString("name")
		}
		tmpArgsArray, _ = cmd.Flags().GetStringArray("owners")
		if len(tmpArgsArray) > 0 {
			updateRepo.Owners, _ = cmd.Flags().GetStringArray("owners")
		}
		tmpArgs, _ = cmd.Flags().GetString("proxy")
		if tmpArgs != "" {
			updateRepo.Proxy, _ = cmd.Flags().GetString("proxy")
		}
		tmpArgsArray, _ = cmd.Flags().GetStringArray("rpm-list")
		if len(tmpArgsArray) > 0 {
			updateRepo.RpmList, _ = cmd.Flags().GetStringArray("rpm-list")
		}

		err = Client.UpdateRepo(updateRepo)

		if checkError(err) != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

var repoFindCmd = &cobra.Command{
	Use:   "find",
	Short: "find repository",
	Long:  `Finds a given repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var repoListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all repositorys",
	Long:  `Lists all available repositories.`,
	Run: func(cmd *cobra.Command, args []string) {

		repos, err = Client.GetRepos()

		if checkError(err) == nil {
			fmt.Println(repos)
		} else {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	},
}

var repoRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove repository",
	Long:  `Removes a given repository.`,
	Run: func(cmd *cobra.Command, args []string) {

		rname, _ := cmd.Flags().GetString("name")
		err := Client.DeleteRepo(rname)
		if checkError(err) != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	},
}

var repoRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename repository",
	Long:  `Renames a given repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var repoReportCmd = &cobra.Command{
	Use:   "report",
	Short: "list all repositorys in detail",
	Long:  `Shows detailed information about all repositories.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
	repoCmd.AddCommand(repoAddCmd)
	repoCmd.AddCommand(repoAutoAddCmd)
	repoCmd.AddCommand(repoCopyCmd)
	repoCmd.AddCommand(repoEditCmd)
	repoCmd.AddCommand(repoFindCmd)
	repoCmd.AddCommand(repoListCmd)
	repoCmd.AddCommand(repoRemoveCmd)
	repoCmd.AddCommand(repoRenameCmd)
	repoCmd.AddCommand(repoReportCmd)

	// local flags for repo add
	repoAddCmd.Flags().String("name", "", "the repo name")
	repoAddCmd.Flags().String("arch", "", "Architecture")
	repoAddCmd.Flags().String("breed", "", "Breed (valid options: none,rsync,rhn,yum,apt,wget)")
	repoAddCmd.Flags().String("comment", "", "free form text description")
	repoAddCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	repoAddCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	repoAddCmd.Flags().String("apt-components", "", "APT components (e.g. main restricted universe)")
	repoAddCmd.Flags().String("apt-dists", "", "APT dist names (e.g. precise,bullseye,buster)")
	repoAddCmd.Flags().String("createrepo-flags", "", "flags to use with createrepo")
	repoAddCmd.Flags().String("environment", "", "environment variables (use these environment variables during commands (key=value, space delimited)")
	repoAddCmd.Flags().Bool("keep-updated", false, "update this repo on next 'cobbler reposync'?")
	repoAddCmd.Flags().String("mirror", "", "address of yum or rsync repo to mirror")
	repoAddCmd.Flags().String("mirror-type", "", "mirror type. Valid options: metalink,mirrorlist,baseurl)")
	repoAddCmd.Flags().String("priority", "", "value for yum priorities plugin, if installed")
	repoAddCmd.Flags().String("proxy", "", "proxy URL (<<inherit>> to use proxy_url_ext from settings, blank or <<None>> for no proxy)")
	repoAddCmd.Flags().String("rpm-list", "", "mirror just these RPMs (yum only)")
	repoAddCmd.Flags().String("yumopts", "", "options to write to yum config file")
	repoAddCmd.Flags().String("rsyncopts", "", "options to use with rsync repo")

	// local flags for repo autoadd
	repoAutoAddCmd.Flags().String("name", "", "the repo name")
	repoAutoAddCmd.Flags().String("arch", "", "Architecture")
	repoAutoAddCmd.Flags().String("breed", "", "Breed (valid options: none,rsync,rhn,yum,apt,wget)")
	repoAutoAddCmd.Flags().String("comment", "", "free form text description")
	repoAutoAddCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	repoAutoAddCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	repoAutoAddCmd.Flags().String("apt-components", "", "APT components (e.g. main restricted universe)")
	repoAutoAddCmd.Flags().String("apt-dists", "", "APT dist names (e.g. precise,bullseye,buster)")
	repoAutoAddCmd.Flags().String("createrepo-flags", "", "flags to use with createrepo")
	repoAutoAddCmd.Flags().String("environment", "", "environment variables (use these environment variables during commands (key=value, space delimited)")
	repoAutoAddCmd.Flags().Bool("keep-updated", false, "update this repo on next 'cobbler reposync'?")
	repoAutoAddCmd.Flags().String("mirror", "", "address of yum or rsync repo to mirror")
	repoAutoAddCmd.Flags().String("mirror-type", "", "mirror type. Valid options: metalink,mirrorlist,baseurl)")
	repoAutoAddCmd.Flags().String("priority", "", "value for yum priorities plugin, if installed")
	repoAutoAddCmd.Flags().String("proxy", "", "proxy URL (<<inherit>> to use proxy_url_ext from settings, blank or <<None>> for no proxy)")
	repoAutoAddCmd.Flags().String("rpm-list", "", "mirror just these RPMs (yum only)")
	repoAutoAddCmd.Flags().String("yumopts", "", "options to write to yum config file")
	repoAutoAddCmd.Flags().String("rsyncopts", "", "options to use with rsync repo")

	// local flags for repo copy
	repoCopyCmd.Flags().String("name", "", "the repo name")
	repoCopyCmd.Flags().String("newname", "", "the new repo name")
	repoCopyCmd.Flags().String("arch", "", "Architecture")
	repoCopyCmd.Flags().String("breed", "", "Breed (valid options: none,rsync,rhn,yum,apt,wget)")
	repoCopyCmd.Flags().String("comment", "", "free form text description")
	repoCopyCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	repoCopyCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	repoCopyCmd.Flags().String("apt-components", "", "APT components (e.g. main restricted universe)")
	repoCopyCmd.Flags().String("apt-dists", "", "APT dist names (e.g. precise,bullseye,buster)")
	repoCopyCmd.Flags().String("createrepo-flags", "", "flags to use with createrepo")
	repoCopyCmd.Flags().String("environment", "", "environment variables (use these environment variables during commands (key=value, space delimited)")
	repoCopyCmd.Flags().Bool("keep-updated", false, "update this repo on next 'cobbler reposync'?")
	repoCopyCmd.Flags().String("mirror", "", "address of yum or rsync repo to mirror")
	repoCopyCmd.Flags().String("mirror-type", "", "mirror type. Valid options: metalink,mirrorlist,baseurl)")
	repoCopyCmd.Flags().String("priority", "", "value for yum priorities plugin, if installed")
	repoCopyCmd.Flags().String("proxy", "", "proxy URL (<<inherit>> to use proxy_url_ext from settings, blank or <<None>> for no proxy)")
	repoCopyCmd.Flags().String("rpm-list", "", "mirror just these RPMs (yum only)")
	repoCopyCmd.Flags().String("yumopts", "", "options to write to yum config file")
	repoCopyCmd.Flags().String("rsyncopts", "", "options to use with rsync repo")

	// local flags for repo edit
	repoEditCmd.Flags().String("name", "", "the repo name")
	repoEditCmd.Flags().String("arch", "", "Architecture")
	repoEditCmd.Flags().String("breed", "", "Breed (valid options: none,rsync,rhn,yum,apt,wget)")
	repoEditCmd.Flags().String("comment", "", "free form text description")
	repoEditCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	repoEditCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	repoEditCmd.Flags().String("apt-components", "", "APT components (e.g. main restricted universe)")
	repoEditCmd.Flags().String("apt-dists", "", "APT dist names (e.g. precise,bullseye,buster)")
	repoEditCmd.Flags().String("createrepo-flags", "", "flags to use with createrepo")
	repoEditCmd.Flags().String("environment", "", "environment variables (use these environment variables during commands (key=value, space delimited)")
	repoEditCmd.Flags().Bool("keep-updated", false, "update this repo on next 'cobbler reposync'?")
	repoEditCmd.Flags().String("mirror", "", "address of yum or rsync repo to mirror")
	repoEditCmd.Flags().String("mirror-type", "", "mirror type. Valid options: metalink,mirrorlist,baseurl)")
	repoEditCmd.Flags().String("priority", "", "value for yum priorities plugin, if installed")
	repoEditCmd.Flags().String("proxy", "", "proxy URL (<<inherit>> to use proxy_url_ext from settings, blank or <<None>> for no proxy)")
	repoEditCmd.Flags().String("rpm-list", "", "mirror just these RPMs (yum only)")
	repoEditCmd.Flags().String("yumopts", "", "options to write to yum config file")
	repoEditCmd.Flags().String("rsyncopts", "", "options to use with rsync repo")

	// local flags for repo find
	repoFindCmd.Flags().String("name", "", "the repo name")
	repoFindCmd.Flags().String("ctime", "", "")
	repoFindCmd.Flags().String("depth", "", "")
	repoFindCmd.Flags().String("mtime", "", "")
	repoFindCmd.Flags().String("uid", "", "UID")
	repoFindCmd.Flags().String("arch", "", "Architecture")
	repoFindCmd.Flags().String("breed", "", "Breed (valid options: none,rsync,rhn,yum,apt,wget)")
	repoFindCmd.Flags().String("comment", "", "free form text description")
	repoFindCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	repoFindCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	repoFindCmd.Flags().String("parent", "", "")
	repoFindCmd.Flags().String("apt-components", "", "APT components (e.g. main restricted universe)")
	repoFindCmd.Flags().String("apt-dists", "", "APT dist names (e.g. precise,bullseye,buster)")
	repoFindCmd.Flags().String("createrepo-flags", "", "flags to use with createrepo")
	repoFindCmd.Flags().String("environment", "", "environment variables (use these environment variables during commands (key=value, space delimited)")
	repoFindCmd.Flags().Bool("keep-updated", false, "update this repo on next 'cobbler reposync'?")
	repoFindCmd.Flags().String("mirror", "", "address of yum or rsync repo to mirror")
	repoFindCmd.Flags().String("mirror-type", "", "mirror type. Valid options: metalink,mirrorlist,baseurl)")
	repoFindCmd.Flags().String("priority", "", "value for yum priorities plugin, if installed")
	repoFindCmd.Flags().String("proxy", "", "proxy URL (<<inherit>> to use proxy_url_ext from settings, blank or <<None>> for no proxy)")
	repoFindCmd.Flags().String("rpm-list", "", "mirror just these RPMs (yum only)")
	repoFindCmd.Flags().String("yumopts", "", "options to write to yum config file")
	repoFindCmd.Flags().String("rsyncopts", "", "options to use with rsync repo")

	// local flags for repo remove
	repoRemoveCmd.Flags().String("name", "", "the repo name")
	repoRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")

	// local flags for repo rename
	repoRenameCmd.Flags().String("name", "", "the repo name")
	repoRenameCmd.Flags().String("newname", "", "the new repo name")
	repoRenameCmd.Flags().String("arch", "", "Architecture")
	repoRenameCmd.Flags().String("breed", "", "Breed (valid options: none,rsync,rhn,yum,apt,wget)")
	repoRenameCmd.Flags().String("comment", "", "free form text description")
	repoRenameCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	repoRenameCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	repoRenameCmd.Flags().String("apt-components", "", "APT components (e.g. main restricted universe)")
	repoRenameCmd.Flags().String("apt-dists", "", "APT dist names (e.g. precise,bullseye,buster)")
	repoRenameCmd.Flags().String("createrepo-flags", "", "flags to use with createrepo")
	repoRenameCmd.Flags().String("environment", "", "environment variables (use these environment variables during commands (key=value, space delimited)")
	repoRenameCmd.Flags().Bool("keep-updated", false, "update this repo on next 'cobbler reposync'?")
	repoRenameCmd.Flags().String("mirror", "", "address of yum or rsync repo to mirror")
	repoRenameCmd.Flags().String("mirror-type", "", "mirror type. Valid options: metalink,mirrorlist,baseurl)")
	repoRenameCmd.Flags().String("priority", "", "value for yum priorities plugin, if installed")
	repoRenameCmd.Flags().String("proxy", "", "proxy URL (<<inherit>> to use proxy_url_ext from settings, blank or <<None>> for no proxy)")
	repoRenameCmd.Flags().String("rpm-list", "", "mirror just these RPMs (yum only)")
	repoRenameCmd.Flags().String("yumopts", "", "options to write to yum config file")
	repoRenameCmd.Flags().String("rsyncopts", "", "options to use with rsync repo")

	// local flags for repo report
	repoReportCmd.Flags().String("name", "", "the repo name")
}
