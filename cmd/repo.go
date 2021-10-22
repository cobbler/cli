// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"github.com/spf13/cobra"
)

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Repository management",
	Long: `Let you manage repositories.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-repo for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var repoAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add repository",
	Long:  `Adds a repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var repoAutoAddCmd = &cobra.Command{
	Use:   "autoadd",
	Short: "add repository automatically",
	Long:  `Automatically adds a repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var repoCopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy repository",
	Long:  `Copies a repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var repoEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit repository",
	Long:  `Edits a repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var repoFindCmd = &cobra.Command{
	Use:   "find",
	Short: "find repository",
	Long:  `Finds a given repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var repoListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all repositorys",
	Long:  `Lists all available repositories.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var repoRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove repository",
	Long:  `Removes a given repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var repoRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename repository",
	Long:  `Renames a given repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var repoReportCmd = &cobra.Command{
	Use:   "report",
	Short: "list all repositorys in detail",
	Long:  `Shows detailed information about all repositories.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
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
