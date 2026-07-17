// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	cobbler "github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
)

// parseNetworkInterfaceType converts the CLI string form of an interface type
// (na, bond, bond_slave, bridge, ...) to the typed enum used by cobblerclient.
func parseNetworkInterfaceType(s string) (cobbler.NetworkInterfaceType, error) {
	switch strings.ToLower(s) {
	case "", "na":
		return cobbler.NetworkInterfaceTypeNA, nil
	case "bond":
		return cobbler.NetworkInterfaceTypeBond, nil
	case "bond_slave":
		return cobbler.NetworkInterfaceTypeBondSlave, nil
	case "bridge":
		return cobbler.NetworkInterfaceTypeBridge, nil
	case "bridge_slave":
		return cobbler.NetworkInterfaceTypeBridgeSlave, nil
	case "bonded_bridge_slave":
		return cobbler.NetworkInterfaceTypeBondedBridgeSlave, nil
	case "infiniband":
		return cobbler.NetworkInterfaceTypeInfiniband, nil
	}
	return cobbler.NetworkInterfaceTypeNA, fmt.Errorf("unknown interface type %q", s)
}

// updateNetworkInterfaceFromFlags applies any --flags the user supplied on the
// command line to iface. Flag names mirror NetworkInterface field names with
// IPv4/IPv6/DNS prefixes used for the nested value objects.
func updateNetworkInterfaceFromFlags(cmd *cobra.Command, iface *cobbler.NetworkInterface) error {
	var err error
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		if err != nil {
			return
		}
		switch flag.Name {
		// Top-level link-layer
		case "mac-address":
			iface.MacAddress, err = cmd.Flags().GetString("mac-address")
		case "interface-type":
			var v string
			v, err = cmd.Flags().GetString("interface-type")
			if err != nil {
				return
			}
			iface.InterfaceType, err = parseNetworkInterfaceType(v)
		case "interface-master":
			iface.InterfaceMaster, err = cmd.Flags().GetString("interface-master")
		case "bonding-opts":
			iface.BondingOpts, err = cmd.Flags().GetString("bonding-opts")
		case "bridge-opts":
			iface.BridgeOpts, err = cmd.Flags().GetString("bridge-opts")
		case "connected-mode":
			iface.ConnectedMode, err = cmd.Flags().GetBool("connected-mode")
		case "management":
			iface.Management, err = cmd.Flags().GetBool("management")
		case "static":
			iface.Static, err = cmd.Flags().GetBool("static")
		case "dhcp-tag":
			iface.DHCPTag, err = cmd.Flags().GetString("dhcp-tag")
		case "mtu":
			iface.MTU, err = cmd.Flags().GetString("mtu")
		case "virt-bridge":
			fallthrough
		case "virt-bridge-inherit":
			if cmd.Flags().Lookup("virt-bridge-inherit") != nil &&
				cmd.Flags().Lookup("virt-bridge-inherit").Changed {
				iface.VirtBridge.Data = ""
				iface.VirtBridge.IsInherited, err = cmd.Flags().GetBool("virt-bridge-inherit")
			} else {
				iface.VirtBridge.IsInherited = false
				iface.VirtBridge.Data, err = cmd.Flags().GetString("virt-bridge")
			}
		// IPv4
		case "ipv4-address":
			iface.IPv4.Address, err = cmd.Flags().GetString("ipv4-address")
		case "ipv4-netmask":
			iface.IPv4.Netmask, err = cmd.Flags().GetString("ipv4-netmask")
		case "ipv4-gateway":
			iface.IfGateway, err = cmd.Flags().GetString("ipv4-gateway")
		case "ipv4-static-routes":
			iface.IPv4.StaticRoutes, err = cmd.Flags().GetStringSlice("ipv4-static-routes")
		// IPv6
		case "ipv6-address":
			iface.IPv6.Address, err = cmd.Flags().GetString("ipv6-address")
		case "ipv6-prefix":
			iface.IPv6.Prefix, err = cmd.Flags().GetString("ipv6-prefix")
		case "ipv6-secondaries":
			iface.IPv6.Secondaries, err = cmd.Flags().GetStringSlice("ipv6-secondaries")
		case "ipv6-mtu":
			iface.IPv6.MTU, err = cmd.Flags().GetString("ipv6-mtu")
		case "ipv6-static-routes":
			iface.IPv6.StaticRoutes, err = cmd.Flags().GetStringSlice("ipv6-static-routes")
		case "ipv6-default-gateway":
			iface.Ipv6DefaultGateway, err = cmd.Flags().GetString("ipv6-default-gateway")
		// DNS
		case "dns-name":
			iface.DNS.Name, err = cmd.Flags().GetString("dns-name")
		case "dns-cnames":
			iface.DNS.CNames, err = cmd.Flags().GetStringSlice("dns-cnames")
		}
	})
	return err
}

// addInterfaceFlagSet registers the full set of interface configuration flags
// onto a subcommand.
func addInterfaceFlagSet(cmd *cobra.Command) {
	addStringFlags(cmd, interfaceStringFlagMetadata)
	addBoolFlags(cmd, interfaceBoolFlagMetadata)
	addStringSliceFlags(cmd, interfaceStringSliceFlagMetadata)
}

// resolveSystemUid looks up the UID of a system, given either --system-uid or
// --system-name. Returns an error if neither is set.
func resolveSystemUid(cmd *cobra.Command) (string, error) {
	systemUid, err := cmd.Flags().GetString("system-uid")
	if err != nil {
		return "", err
	}
	if systemUid != "" {
		return systemUid, nil
	}
	systemName, err := cmd.Flags().GetString("system-name")
	if err != nil {
		return "", err
	}
	if systemName == "" {
		return "", fmt.Errorf("one of --system-name or --system-uid is required")
	}
	system, err := Client.GetSystem(systemName, false, false)
	if err != nil {
		return "", err
	}
	return system.Uid, nil
}

// NewInterfaceCommand builds the `cobbler interface` command and its subtree.
func NewInterfaceCommand() (*cobra.Command, error) {
	interfaceCmd := &cobra.Command{
		Use:   "interface",
		Short: "Manage network interfaces",
		Long:  `Manage network interfaces as first-class Cobbler 4.0.0 items.`,
	}
	interfaceCmd.AddCommand(NewInterfaceAddCommand())
	interfaceCmd.AddCommand(NewInterfaceCopyCommand())
	interfaceCmd.AddCommand(NewInterfaceEditCommand())
	interfaceCmd.AddCommand(NewInterfaceFindCommand())
	interfaceCmd.AddCommand(NewInterfaceListCommand())
	interfaceCmd.AddCommand(NewInterfaceRemoveCommand())
	interfaceCmd.AddCommand(NewInterfaceRenameCommand())
	interfaceCmd.AddCommand(NewInterfaceReportCommand())
	interfaceCmd.AddCommand(NewInterfaceExportCmd())
	return interfaceCmd, nil
}

func NewInterfaceAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add a network interface",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			if name == "" {
				return fmt.Errorf("--name is required")
			}
			systemUid, err := resolveSystemUid(cmd)
			if err != nil {
				return err
			}
			iface := cobbler.NewNetworkInterface()
			iface.Name = name
			iface.SystemUid = systemUid
			if err := updateNetworkInterfaceFromFlags(cmd, &iface); err != nil {
				return err
			}
			created, err := Client.CreateNetworkInterface(systemUid, iface)
			if err != nil {
				return err
			}
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Network interface %s created\n", created.Name)
			return nil
		},
	}
	cmd.Flags().String("name", "", "the network interface name")
	cmd.Flags().String("system-name", "", "the parent system name (resolved to UID)")
	cmd.Flags().String("system-uid", "", "the parent system UID")
	addInterfaceFlagSet(cmd)
	return cmd
}

func NewInterfaceEditCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "edit a network interface",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			iface, err := Client.GetNetworkInterface(name, false, false)
			if err != nil {
				return err
			}
			if err := updateNetworkInterfaceFromFlags(cmd, iface); err != nil {
				return err
			}
			return Client.UpdateNetworkInterface(iface)
		},
	}
	cmd.Flags().String("name", "", "the network interface name")
	addInterfaceFlagSet(cmd)
	return cmd
}

func NewInterfaceCopyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "copy",
		Short: "copy a network interface",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			newName, err := cmd.Flags().GetString("newname")
			if err != nil {
				return err
			}
			handle, err := Client.GetNetworkInterfaceHandle(name)
			if err != nil {
				return err
			}
			if err := Client.CopyNetworkInterface(handle, newName); err != nil {
				return err
			}
			// Clear identity-bound fields on the copy by default; users can
			// re-set them via --mac-address / --ipv4-address / --ipv6-address.
			fresh, err := Client.GetNetworkInterface(newName, false, false)
			if err != nil {
				return err
			}
			fresh.MacAddress = ""
			fresh.IPv4.Address = ""
			fresh.IPv6.Address = ""
			if err := updateNetworkInterfaceFromFlags(cmd, fresh); err != nil {
				return err
			}
			return Client.UpdateNetworkInterface(fresh)
		},
	}
	cmd.Flags().String("name", "", "the network interface to copy")
	cmd.Flags().String("newname", "", "the new interface name")
	addInterfaceFlagSet(cmd)
	return cmd
}

func NewInterfaceRenameCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rename",
		Short: "rename a network interface",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			newName, err := cmd.Flags().GetString("newname")
			if err != nil {
				return err
			}
			handle, err := Client.GetNetworkInterfaceHandle(name)
			if err != nil {
				return err
			}
			return Client.RenameNetworkInterface(handle, newName)
		},
	}
	cmd.Flags().String("name", "", "the network interface to rename")
	cmd.Flags().String("newname", "", "the new interface name")
	return cmd
}

func NewInterfaceRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "remove a network interface",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			return Client.DeleteNetworkInterface(name)
		},
	}
	cmd.Flags().String("name", "", "the network interface to remove")
	return cmd
}

func NewInterfaceFindCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "find",
		Short: "find network interfaces",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			return FindItemNames(cmd, args, "network_interface")
		},
	}
	addInterfaceFlagSet(cmd)
	cmd.Flags().String("name", "", "match by interface name")
	cmd.Flags().String("system-name", "", "filter by parent system name")
	cmd.Flags().String("system-uid", "", "filter by parent system UID")
	addPaginationFlags(cmd)
	return cmd
}

func NewInterfaceListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list network interfaces grouped by system",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			interfaces, err := Client.GetNetworkInterfaces()
			if err != nil {
				return err
			}
			grouped := make(map[string][]string)
			for _, iface := range interfaces {
				grouped[iface.SystemName] = append(grouped[iface.SystemName], iface.Name)
			}
			systemNames := make([]string, 0, len(grouped))
			for systemName := range grouped {
				systemNames = append(systemNames, systemName)
			}
			sort.Strings(systemNames)
			for _, systemName := range systemNames {
				ifaceNames := grouped[systemName]
				sort.Strings(ifaceNames)
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s:\n", systemName)
				for _, n := range ifaceNames {
					_, _ = fmt.Fprintf(cmd.OutOrStdout(), "    %s\n", n)
				}
			}
			return nil
		},
	}
	return cmd
}

func NewInterfaceReportCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report",
		Short: "show network interface details",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			systemName, err := cmd.Flags().GetString("system-name")
			if err != nil {
				return err
			}

			var interfaces []*cobbler.NetworkInterface
			switch {
			case name != "":
				iface, err := Client.GetNetworkInterface(name, false, false)
				if err != nil {
					return err
				}
				interfaces = []*cobbler.NetworkInterface{iface}
			case systemName != "":
				interfaces, err = Client.FindNetworkInterface(map[string]interface{}{
					"system_name": systemName,
				}, false)
				if err != nil {
					return err
				}
			default:
				interfaces, err = Client.GetNetworkInterfaces()
				if err != nil {
					return err
				}
			}
			for _, iface := range interfaces {
				printStructured(cmd, iface)
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "")
			}
			return nil
		},
	}
	cmd.Flags().String("name", "", "the network interface name")
	cmd.Flags().String("system-name", "", "filter by parent system name")
	return cmd
}

func NewInterfaceExportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "export network interfaces",
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
			systemName, err := cmd.Flags().GetString("system-name")
			if err != nil {
				return err
			}
			format, err := cmd.Flags().GetString("format")
			if err != nil {
				return err
			}

			var interfaces []*cobbler.NetworkInterface
			switch {
			case name != "":
				iface, err := Client.GetNetworkInterface(name, false, false)
				if err != nil {
					return err
				}
				interfaces = []*cobbler.NetworkInterface{iface}
			case systemName != "":
				interfaces, err = Client.FindNetworkInterface(map[string]interface{}{
					"system_name": systemName,
				}, false)
				if err != nil {
					return err
				}
			default:
				interfaces, err = Client.GetNetworkInterfaces()
				if err != nil {
					return err
				}
			}

			for _, iface := range interfaces {
				switch format {
				case "json":
					out, err := json.Marshal(iface)
					if err != nil {
						return err
					}
					_, _ = fmt.Fprintln(cmd.OutOrStdout(), string(out))
				case "yaml":
					out, err := yaml.Marshal(iface)
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
	cmd.Flags().String("name", "", "the network interface name")
	cmd.Flags().String("system-name", "", "filter by parent system name")
	cmd.Flags().String(exportStringMetadata["format"].Name, exportStringMetadata["format"].DefaultValue, exportStringMetadata["format"].Usage)
	return cmd
}
