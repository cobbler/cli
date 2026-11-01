// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	cobbler "github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
)

var templateStringFlagMetadata = map[string]FlagMetadata[string]{
	"template-type": {
		Name:         "template-type",
		DefaultValue: "jinja2",
		Usage:        "template engine (jinja2, cheetah)",
	},
	"uri-schema": {
		Name:         "uri-schema",
		DefaultValue: "file",
		Usage:        "URI schema (file, importlib, environment)",
	},
	"uri-path": {
		Name:         "uri-path",
		DefaultValue: "",
		Usage:        "URI path (file path, importlib path, or env var name)",
	},
	"content": {
		Name:         "content",
		DefaultValue: "",
		Usage:        "inline template content (mutually exclusive with --content-file)",
	},
	"content-file": {
		Name:         "content-file",
		DefaultValue: "",
		Usage:        "path to a file whose contents become the template content",
	},
}

var templateStringSliceFlagMetadata = map[string]FlagMetadata[[]string]{
	"tags": {
		Name:         "tags",
		DefaultValue: []string{},
		Usage:        "template tags (comma delimited)",
	},
}

// parseTemplateSchema converts the CLI string form to the typed enum.
func parseTemplateSchema(s string) (cobbler.TemplateSchema, error) {
	switch strings.ToLower(s) {
	case "", "file":
		return cobbler.TemplateSchemaFile, nil
	case "importlib":
		return cobbler.TemplateSchemaImportlib, nil
	case "environment":
		return cobbler.TemplateSchemaEnvironment, nil
	}
	return cobbler.TemplateSchemaFile, fmt.Errorf("unknown URI schema %q", s)
}

// updateTemplateFromFlags applies any set --flags to t.
func updateTemplateFromFlags(cmd *cobra.Command, t *cobbler.Template) error {
	var err error
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		if err != nil {
			return
		}
		switch flag.Name {
		case "comment":
			t.Comment, err = cmd.Flags().GetString("comment")
		case "template-type":
			t.TemplateType, err = cmd.Flags().GetString("template-type")
		case "uri-schema":
			var v string
			v, err = cmd.Flags().GetString("uri-schema")
			if err != nil {
				return
			}
			t.URI.Schema, err = parseTemplateSchema(v)
		case "uri-path":
			t.URI.Path, err = cmd.Flags().GetString("uri-path")
		case "content":
			t.Content, err = cmd.Flags().GetString("content")
		case "content-file":
			var path string
			path, err = cmd.Flags().GetString("content-file")
			if err != nil || path == "" {
				return
			}
			var data []byte
			data, err = os.ReadFile(path)
			if err != nil {
				return
			}
			t.Content = string(data)
		case "tags":
			t.Tags, err = cmd.Flags().GetStringSlice("tags")
		}
	})
	return err
}

func addTemplateFlagSet(cmd *cobra.Command) {
	addCommonArgs(cmd)
	addStringFlags(cmd, templateStringFlagMetadata)
	addStringSliceFlags(cmd, templateStringSliceFlagMetadata)
}

// NewTemplateCmd builds the `cobbler template` command tree.
func NewTemplateCmd() (*cobra.Command, error) {
	templateCmd := &cobra.Command{
		Use:   "template",
		Short: "Manage autoinstall templates",
		Long:  `Manage Cobbler 4.0.0 Template items.`,
	}
	templateCmd.AddCommand(NewTemplateAddCmd())
	templateCmd.AddCommand(NewTemplateCopyCmd())
	templateCmd.AddCommand(NewTemplateEditCmd())
	templateCmd.AddCommand(NewTemplateFindCmd())
	templateCmd.AddCommand(NewTemplateListCmd())
	templateCmd.AddCommand(NewTemplateRemoveCmd())
	templateCmd.AddCommand(NewTemplateRenameCmd())
	templateCmd.AddCommand(NewTemplateReportCmd())
	templateCmd.AddCommand(NewTemplateExportCmd())
	templateCmd.AddCommand(NewTemplateContentCmd())
	templateCmd.AddCommand(NewTemplateRefreshCmd())
	return templateCmd, nil
}

func NewTemplateAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add a template",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			t := cobbler.NewTemplate()
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			t.Name = name
			if err := updateTemplateFromFlags(cmd, &t); err != nil {
				return err
			}
			var created *cobbler.Template
			err = withTransaction(func() error {
				var txErr error
				created, txErr = Client.CreateTemplate(t)
				return txErr
			})
			if err != nil {
				return err
			}
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Template %s created\n", created.Name)
			return nil
		},
	}
	addTemplateFlagSet(cmd)
	return cmd
}

func NewTemplateEditCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "edit a template",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			uid, err := cmd.Flags().GetString("uid")
			if err != nil {
				return err
			}
			resolvedUID, err := resolveUID(&Client, "template", name, uid)
			if err != nil {
				return err
			}
			t, err := Client.GetTemplate(resolvedUID, false, false)
			if err != nil {
				return err
			}
			if err := updateTemplateFromFlags(cmd, t); err != nil {
				return err
			}
			return withTransaction(func() error {
				return Client.UpdateTemplate(t)
			})
		},
	}
	addTemplateFlagSet(cmd)
	addUIDFlag(cmd, "template")
	return cmd
}

func NewTemplateCopyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "copy",
		Short: "copy a template",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			uid, err := cmd.Flags().GetString("uid")
			if err != nil {
				return err
			}
			newName, err := cmd.Flags().GetString("newname")
			if err != nil {
				return err
			}
			handle, err := resolveUID(&Client, "template", name, uid)
			if err != nil {
				return err
			}
			return withTransaction(func() error {
				return Client.CopyTemplate(handle, newName)
			})
		},
	}
	addStringFlags(cmd, commonStringFlagMetadata)
	addStringFlags(cmd, copyRenameStringFlagMetadata)
	addUIDFlag(cmd, "template")
	return cmd
}

func NewTemplateRenameCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rename",
		Short: "rename a template",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			uid, err := cmd.Flags().GetString("uid")
			if err != nil {
				return err
			}
			newName, err := cmd.Flags().GetString("newname")
			if err != nil {
				return err
			}
			handle, err := resolveUID(&Client, "template", name, uid)
			if err != nil {
				return err
			}
			return withTransaction(func() error {
				return Client.RenameTemplate(handle, newName)
			})
		},
	}
	addStringFlags(cmd, commonStringFlagMetadata)
	addStringFlags(cmd, copyRenameStringFlagMetadata)
	addUIDFlag(cmd, "template")
	return cmd
}

// NOTE: unlike every other item type's remove command, Template's remove below is
// deliberately NOT wrapped in withTransaction. This is out of scope for this task (which
// only covers add/edit/copy/rename); see the task brief for Template's add/edit/copy/rename
// cleanup for background on the now-fixed server-side crash this guarded against for
// mutating calls. Whether remove_item is similarly safe under a transaction for a non-
// InheritableItem type like Template was not verified here.
func NewTemplateRemoveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "remove a template",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			return RemoveItemRecursive(cmd, args, "template")
		},
	}
	cmd.Flags().String("name", "", "the template name")
	cmd.Flags().Bool("recursive", false, "also delete child objects")
	addUIDFlag(cmd, "template")
	return cmd
}

func NewTemplateFindCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "find",
		Short: "find templates",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			return FindItemNames(cmd, args, "template")
		},
	}
	addStringFlags(cmd, commonStringFlagMetadata)
	addStringFlags(cmd, templateStringFlagMetadata)
	addStringSliceFlags(cmd, templateStringSliceFlagMetadata)
	addPaginationFlags(cmd)
	return cmd
}

func NewTemplateListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all templates",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			names, err := Client.ListTemplateNames()
			if err != nil {
				return err
			}
			listItems(cmd, "templates", names)
			return nil
		},
	}
	return cmd
}

// templateReportContentIndicator returns the value `template report` should
// print for a Template's content field. Content can be arbitrarily large
// (it's a whole autoinstall template body), so report deliberately never
// prints it inline -- that would make report's output unusable for exactly
// the templates where content's presence matters most. When content is
// empty it's returned unchanged, matching every other field's existing
// convention of printing an empty value rather than omitting the field.
// When content is non-empty, a short indicator is returned instead,
// pointing at the dedicated `template content` subcommand that already
// prints the real (and possibly very long) text.
func templateReportContentIndicator(t *cobbler.Template) string {
	if t.Content == "" {
		return t.Content
	}
	return fmt.Sprintf(
		"<%d bytes, not shown -- see 'cobbler template content --name %s'>",
		len(t.Content), t.Name,
	)
}

func NewTemplateReportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report",
		Short: "show template details",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			uid, err := cmd.Flags().GetString("uid")
			if err != nil {
				return err
			}
			var templates []*cobbler.Template
			if name == "" && uid == "" {
				templates, err = Client.GetTemplates()
				if err != nil {
					return err
				}
			} else {
				resolvedUID, err := resolveUID(&Client, "template", name, uid)
				if err != nil {
					return err
				}
				t, err := Client.GetTemplate(resolvedUID, false, false)
				if err != nil {
					return err
				}
				templates = []*cobbler.Template{t}
			}
			for _, t := range templates {
				// Print a shallow copy with content replaced by a short indicator
				// (see templateReportContentIndicator) rather than the real
				// struct, so the full -- potentially huge -- content is never
				// dumped inline by report.
				display := *t
				display.Content = templateReportContentIndicator(t)
				printStructured(cmd, &display)
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "")
			}
			return nil
		},
	}
	cmd.Flags().String("name", "", "the template name")
	addUIDFlag(cmd, "template")
	return cmd
}

func NewTemplateExportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "export templates",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			format, err := cmd.Flags().GetString("format")
			if err != nil {
				return err
			}
			if format != "json" && format != "yaml" {
				return fmt.Errorf("format must be json or yaml")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			uid, err := cmd.Flags().GetString("uid")
			if err != nil {
				return err
			}
			format, err := cmd.Flags().GetString("format")
			if err != nil {
				return err
			}
			var templates []*cobbler.Template
			if name == "" && uid == "" {
				templates, err = Client.GetTemplates()
				if err != nil {
					return err
				}
			} else {
				resolvedUID, err := resolveUID(&Client, "template", name, uid)
				if err != nil {
					return err
				}
				t, err := Client.GetTemplate(resolvedUID, false, false)
				if err != nil {
					return err
				}
				templates = []*cobbler.Template{t}
			}
			for _, t := range templates {
				switch format {
				case "json":
					out, err := json.Marshal(t)
					if err != nil {
						return err
					}
					_, _ = fmt.Fprintln(cmd.OutOrStdout(), string(out))
				case "yaml":
					out, err := yaml.Marshal(t)
					if err != nil {
						return err
					}
					_, _ = fmt.Fprintln(cmd.OutOrStdout(), "---")
					_, _ = fmt.Fprintln(cmd.OutOrStdout(), string(out))
				}
			}
			return nil
		},
	}
	cmd.Flags().String("name", "", "the template name")
	cmd.Flags().String(exportStringMetadata["format"].Name, exportStringMetadata["format"].DefaultValue, exportStringMetadata["format"].Usage)
	addUIDFlag(cmd, "template")
	return cmd
}

// NewTemplateContentCmd dumps the resolved template content to stdout.
func NewTemplateContentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "content",
		Short: "print resolved template content",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			uid, err := cmd.Flags().GetString("uid")
			if err != nil {
				return err
			}
			resolvedUID, err := resolveUID(&Client, "template", name, uid)
			if err != nil {
				return err
			}
			content, err := Client.GetTemplateContent(resolvedUID)
			if err != nil {
				return err
			}
			_, _ = fmt.Fprint(cmd.OutOrStdout(), content)
			return nil
		},
	}
	cmd.Flags().String("name", "", "the template name")
	addUIDFlag(cmd, "template")
	return cmd
}

// NewTemplateRefreshCmd forces a backend re-read of one or more templates.
// With no --name flag, all templates are refreshed.
func NewTemplateRefreshCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "refresh",
		Short: "refresh template content from disk",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			names, err := cmd.Flags().GetStringSlice("name")
			if err != nil {
				return err
			}
			return Client.TemplatesRefreshContent(names)
		},
	}
	cmd.Flags().StringSlice("name", []string{}, "the template name(s) to refresh; omit for all")
	return cmd
}
