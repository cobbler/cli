package cmd

import (
	"encoding/json"
	"fmt"
	cobbler "github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
	"strings"
)

func updateNetworkInterfaceFromFlags(cmd *cobra.Command, networkInterface *cobbler.Interface) error {
	var err error
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		if err != nil {
			// If one of the previous flags has had an error just directly return.
			return
		}
		switch flag.Name {
		// The rename & copy operations are special operations as such we cannot blindly set this inside here.
		// Any rename & copy operation must be handled outside of this method.
		case "bonding-opts":
			var systemNewBondingOpts string
			systemNewBondingOpts, err = cmd.Flags().GetString("bonding-opts")
			if err != nil {
				return
			}
			networkInterface.BondingOpts = systemNewBondingOpts
		case "bridge-opts":
			var systemNewBridgeOpts string
			systemNewBridgeOpts, err = cmd.Flags().GetString("bridge-opts")
			if err != nil {
				return
			}
			networkInterface.BridgeOpts = systemNewBridgeOpts
		case "cnames":
			var systemNewCNames []string
			systemNewCNames, err = cmd.Flags().GetStringSlice("cnames")
			if err != nil {
				return
			}
			networkInterface.CNAMEs = systemNewCNames
		case "connected-mode":
			var systemNewConnectedMode bool
			systemNewConnectedMode, err = cmd.Flags().GetBool("connected-mode")
			if err != nil {
				return
			}
			networkInterface.ConnectedMode = systemNewConnectedMode
		case "dhcp-tag":
			var systemNewDhcpTag string
			systemNewDhcpTag, err = cmd.Flags().GetString("dhcp-tag")
			if err != nil {
				return
			}
			networkInterface.DHCPTag = systemNewDhcpTag
		case "dns-name":
			var systemNewDnsName string
			systemNewDnsName, err = cmd.Flags().GetString("dns-name")
			if err != nil {
				return
			}
			networkInterface.DNSName = systemNewDnsName
		case "if-gateway":
			var systemNewIfGateway string
			systemNewIfGateway, err = cmd.Flags().GetString("if-gateway")
			if err != nil {
				return
			}
			networkInterface.Gateway = systemNewIfGateway
		case "interface-master":
			var systemNewInterfaceMaster string
			systemNewInterfaceMaster, err = cmd.Flags().GetString("interface-master")
			if err != nil {
				return
			}
			networkInterface.InterfaceMaster = systemNewInterfaceMaster
		case "interface-type":
			var systemNewInterfaceType string
			systemNewInterfaceType, err = cmd.Flags().GetString("interface-type")
			if err != nil {
				return
			}
			networkInterface.InterfaceType = systemNewInterfaceType
		case "ip-address":
			var systemNewIpAddress string
			systemNewIpAddress, err = cmd.Flags().GetString("ip-address")
			if err != nil {
				return
			}
			networkInterface.IPAddress = systemNewIpAddress
		case "ipv6-address":
			var systemNewIpv6Address string
			systemNewIpv6Address, err = cmd.Flags().GetString("ipv6-address")
			if err != nil {
				return
			}
			networkInterface.IPv6Address = systemNewIpv6Address
		case "ipv6-default-gateway":
			var systemNewIpv6DefaultGateway string
			systemNewIpv6DefaultGateway, err = cmd.Flags().GetString("ipv6-default-gateway")
			if err != nil {
				return
			}
			networkInterface.IPv6DefaultGateway = systemNewIpv6DefaultGateway
		case "ipv6-mtu":
			var systemNewIpv6Mtu string
			systemNewIpv6Mtu, err = cmd.Flags().GetString("ipv6-mtu")
			if err != nil {
				return
			}
			networkInterface.IPv6MTU = systemNewIpv6Mtu
		case "ipv6-prefix":
			var systemNewIpv6Prefix string
			systemNewIpv6Prefix, err = cmd.Flags().GetString("ipv6-prefix")
			if err != nil {
				return
			}
			networkInterface.IPv6Prefix = systemNewIpv6Prefix
		case "ipv6-secondaries":
			var systemNewIpv6Secondaries []string
			systemNewIpv6Secondaries, err = cmd.Flags().GetStringSlice("ipv6-secondaries")
			if err != nil {
				return
			}
			networkInterface.IPv6Secondaries = systemNewIpv6Secondaries
		case "ipv6-static-routes":
			var systemNewIpv6StaticRoutes []string
			systemNewIpv6StaticRoutes, err = cmd.Flags().GetStringSlice("ipv6-static-routes")
			if err != nil {
				return
			}
			networkInterface.IPv6StaticRoutes = systemNewIpv6StaticRoutes
		case "mac-address":
			var systemNewMacAddress string
			systemNewMacAddress, err = cmd.Flags().GetString("mac-address")
			if err != nil {
				return
			}
			networkInterface.MACAddress = systemNewMacAddress
		case "management":
			var systemNewManagement bool
			systemNewManagement, err = cmd.Flags().GetBool("management")
			if err != nil {
				return
			}
			networkInterface.Management = systemNewManagement
		case "mtu":
			var systemNewMtu string
			systemNewMtu, err = cmd.Flags().GetString("mtu")
			if err != nil {
				return
			}
			networkInterface.MTU = systemNewMtu
		case "netmask":
			var systemNewNetmask string
			systemNewNetmask, err = cmd.Flags().GetString("netmask")
			if err != nil {
				return
			}
			networkInterface.Netmask = systemNewNetmask
		case "static":
			var systemNewStatic bool
			systemNewStatic, err = cmd.Flags().GetBool("static")
			if err != nil {
				return
			}
			networkInterface.Static = systemNewStatic
		case "static-routes":
			var systemNewStaticRoutes []string
			systemNewStaticRoutes, err = cmd.Flags().GetStringSlice("static-routes")
			if err != nil {
				return
			}
			networkInterface.StaticRoutes = systemNewStaticRoutes
		case "virt-bridge":
			var systemNewVirtBridge string
			systemNewVirtBridge, err = cmd.Flags().GetString("virt-bridge")
			if err != nil {
				return
			}
			networkInterface.VirtBridge = systemNewVirtBridge
		}
	})
	// Don't blindly return nil because maybe one of the flags had an issue retrieving an argument.
	return err
}

func NewInterfaceCommand() (*cobra.Command, error) {
	interfaceCmd := &cobra.Command{
		Use:   "interface",
		Short: "Manage interfaces",
		Long:  "Let's you manage network interfaces for systems.",
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
	interfaceAddCmd := &cobra.Command{
		Use: "add",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			systemName, err := cmd.Flags().GetString("system-name")
			if err != nil {
				return err
			}
			networkInterfaceName, err := cmd.Flags().GetString("interface-name")
			if err != nil {
				return err
			}

			systemObject, err := Client.GetSystem(systemName, false, false)
			if err != nil {
				return err
			}

			networkInterface := cobbler.NewInterface()
			err = updateNetworkInterfaceFromFlags(cmd, &networkInterface)
			if err != nil {
				return err
			}

			return systemObject.CreateInterface(networkInterfaceName, networkInterface)
		},
	}
	interfaceAddCmd.Flags().String("interface-name", "", "the interface to operate on")
	interfaceAddCmd.Flags().String("system-name", "", "the system to operate on")
	addStringFlags(interfaceAddCmd, interfaceStringFlagMetadata)
	addBoolFlags(interfaceAddCmd, interfaceBoolFlagMetadata)
	addStringSliceFlags(interfaceAddCmd, interfaceStringSliceFlagMetadata)
	return interfaceAddCmd
}

func NewInterfaceCopyCommand() *cobra.Command {
	interfaceCopyCmd := &cobra.Command{
		Use: "copy",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			systemName, err := cmd.Flags().GetString("system-name")
			if err != nil {
				return err
			}
			networkInterfaceName, err := cmd.Flags().GetString("interface-name")
			if err != nil {
				return err
			}
			newNetworkInterfaceName, err := cmd.Flags().GetString("new-interface-name")
			if err != nil {
				return err
			}

			systemObject, err := Client.GetSystem(systemName, false, false)
			if err != nil {
				return err
			}
			networkInterfaceObject, err := systemObject.GetInterface(networkInterfaceName)
			if err != nil {
				return err
			}
			networkInterfaceObject.MACAddress = ""
			networkInterfaceObject.IPAddress = ""
			networkInterfaceObject.IPv6Address = ""
			err = systemObject.CreateInterface(newNetworkInterfaceName, networkInterfaceObject)
			if err != nil {
				return err
			}

			return nil
		},
	}
	interfaceCopyCmd.Flags().String("interface-name", "", "the interface to operate on")
	interfaceCopyCmd.Flags().String("system-name", "", "the system to operate on")
	interfaceCopyCmd.Flags().String("new-interface-name", "", "the new name for the network interface")
	return interfaceCopyCmd
}

func NewInterfaceEditCommand() *cobra.Command {
	interfaceEditCmd := &cobra.Command{
		Use: "edit",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			systemName, err := cmd.Flags().GetString("system-name")
			if err != nil {
				return err
			}
			networkInterfaceName, err := cmd.Flags().GetString("interface-name")
			if err != nil {
				return err
			}

			systemHandle, err := Client.GetSystemHandle(systemName)
			if err != nil {
				return err
			}

			editedProperties := make(map[string]interface{})
			cmd.Flags().Visit(func(flag *pflag.Flag) {
				propertyName := fmt.Sprintf("%s-%s", networkInterfaceName, strings.Replace(flag.Name, "-", "_", -1))
				editedProperties[propertyName] = flag.Value
			})
			err = Client.ModifyInterface(systemHandle, editedProperties)
			if err != nil {
				return err
			}
			return nil
		},
	}
	interfaceEditCmd.Flags().String("interface-name", "", "the interface to operate on")
	interfaceEditCmd.Flags().String("system-name", "", "the system to operate on")
	addStringFlags(interfaceEditCmd, interfaceStringFlagMetadata)
	addBoolFlags(interfaceEditCmd, interfaceBoolFlagMetadata)
	addStringSliceFlags(interfaceEditCmd, interfaceStringSliceFlagMetadata)
	return interfaceEditCmd
}

func NewInterfaceFindCommand() *cobra.Command {
	interfaceFindCmd := &cobra.Command{
		Use: "find",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			return FindItemNames(cmd, args, "system")
		},
	}
	// Network interface flags
	addStringFlags(interfaceFindCmd, interfaceStringFlagMetadata)
	addBoolFlags(interfaceFindCmd, interfaceBoolFlagMetadata)
	addStringSliceFlags(interfaceFindCmd, interfaceStringSliceFlagMetadata)
	interfaceFindCmd.Flags().String("name", "", "the system to operate on")
	interfaceFindCmd.Flags().String("interface", "", "the interface to operate on")
	return interfaceFindCmd
}

func printNetworkInterfaceNames(cmd *cobra.Command, system *cobbler.System) error {
	networkInterfaces, err := system.GetInterfaces()
	if err != nil {
		return err
	}
	fmt.Fprintf(cmd.OutOrStdout(), "%s:\n", system.Name)
	for interfaceName := range networkInterfaces {
		fmt.Fprintf(cmd.OutOrStdout(), "    %s\n", interfaceName)
	}
	return nil
}

func NewInterfaceListCommand() *cobra.Command {
	interfaceListCmd := &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			systems, err := Client.GetSystems()
			if err != nil {
				return err
			}
			for _, system := range systems {
				err = printNetworkInterfaceNames(cmd, system)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	return interfaceListCmd
}

func NewInterfaceRemoveCommand() *cobra.Command {
	interfaceRemoveCmd := &cobra.Command{
		Use: "remove",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			systemName, err := cmd.Flags().GetString("system-name")
			if err != nil {
				return err
			}
			networkInterfaceName, err := cmd.Flags().GetString("interface-name")
			if err != nil {
				return err
			}

			systemHandle, err := Client.GetSystemHandle(systemName)
			if err != nil {
				return err
			}
			err = Client.DeleteNetworkInterface(systemHandle, networkInterfaceName)
			if err != nil {
				return err
			}

			return nil
		},
	}
	interfaceRemoveCmd.Flags().String("interface-name", "", "the interface to operate on")
	interfaceRemoveCmd.Flags().String("system-name", "", "the system to operate on")
	return interfaceRemoveCmd
}

func NewInterfaceRenameCommand() *cobra.Command {
	interfaceRenameCmd := &cobra.Command{
		Use: "rename",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			systemName, err := cmd.Flags().GetString("system-name")
			if err != nil {
				return err
			}
			networkInterfaceName, err := cmd.Flags().GetString("interface-name")
			if err != nil {
				return err
			}
			newNetworkInterfaceName, err := cmd.Flags().GetString("new-interface-name")
			if err != nil {
				return err
			}

			err = Client.RenameNetworkInterface(systemName, networkInterfaceName, newNetworkInterfaceName)
			if err != nil {
				return err
			}

			return nil
		},
	}
	interfaceRenameCmd.Flags().String("interface-name", "", "the interface to operate on")
	interfaceRenameCmd.Flags().String("system-name", "", "the system to operate on")
	interfaceRenameCmd.Flags().String("new-interface-name", "", "the new name for the network interface")
	return interfaceRenameCmd
}

func NewInterfaceReportCommand() *cobra.Command {
	interfaceReportCmd := &cobra.Command{
		Use: "report",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			systemName, err := cmd.Flags().GetString("system-name")
			if err != nil {
				return err
			}
			interfaceName, err := cmd.Flags().GetString("interface-name")
			if err != nil {
				return err
			}
			itemNames := make([]string, 0)
			if systemName == "" {
				itemNames, err = Client.ListSystemNames()
				if err != nil {
					return err
				}
			} else {
				itemNames = append(itemNames, systemName)
			}
			return reportNetworkInterfaces(cmd, itemNames, interfaceName)
		},
	}
	interfaceReportCmd.Flags().String("interface-name", "", "the interface to operate on")
	interfaceReportCmd.Flags().String("system-name", "", "the system to operate on")
	return interfaceReportCmd
}

func reportNetworkInterfaces(cmd *cobra.Command, systemNames []string, interfaceName string) error {
	for _, itemName := range systemNames {
		system, err := Client.GetSystem(itemName, false, false)
		if err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "%s:\n", itemName)
		if interfaceName == "" {
			networkInterfaces, err := system.GetInterfaces()
			if err != nil {
				return err
			}
			for networkInterfaceName := range networkInterfaces {
				networkInterface := system.Interfaces[networkInterfaceName]
				printStructured(cmd, &networkInterface)
			}
			fmt.Fprintln(cmd.OutOrStdout(), "")
		} else {
			networkInterface, err := system.GetInterface(interfaceName)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), "    Interface %s not found on system", interfaceName)
				continue
			}
			printStructured(cmd, &networkInterface)
		}
	}
	return nil
}

func NewInterfaceExportCmd() *cobra.Command {
	networkInterfaceExportCmd := &cobra.Command{
		Use:   "export",
		Short: "export network interfaces",
		Long:  `Export network interfaces.`,
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

			systemName, err := cmd.Flags().GetString("system-name")
			if err != nil {
				return err
			}
			interfaceName, err := cmd.Flags().GetString("interface-name")
			if err != nil {
				return err
			}
			formatOption, err := cmd.Flags().GetString("format")
			if err != nil {
				return err
			}

			itemNames := make([]string, 0)
			if systemName == "" {
				itemNames, err = Client.ListSystemNames()
				if err != nil {
					return err
				}
			} else {
				itemNames = append(itemNames, systemName)
			}

			for _, itemName := range itemNames {
				system, err := Client.GetSystem(itemName, false, false)
				if err != nil {
					return err
				}

				var systemInterfaces cobbler.Interfaces
				if interfaceName == "" {
					systemInterfaces = system.Interfaces
				} else {
					systemInterfaces = make(map[string]cobbler.Interface)
					intf, interfaceExists := system.Interfaces[interfaceName]
					if interfaceExists {
						systemInterfaces[interfaceName] = intf
					}
				}
				exportData := struct {
					SystemName string `json:"system_name" yaml:"system_name"`
					Interfaces cobbler.Interfaces
				}{
					itemName,
					systemInterfaces,
				}
				if formatOption == "json" {
					jsonDocument, err := json.Marshal(exportData)
					if err != nil {
						return err
					}
					fmt.Fprintln(cmd.OutOrStdout(), string(jsonDocument))
				}
				if formatOption == "yaml" {
					yamlDocument, err := yaml.Marshal(exportData)
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
	networkInterfaceExportCmd.Flags().String("interface-name", "", "the network interface name")
	networkInterfaceExportCmd.Flags().String("system-name", "", "the system name")
	networkInterfaceExportCmd.Flags().String(exportStringMetadata["format"].Name, exportStringMetadata["format"].DefaultValue, exportStringMetadata["format"].Usage)
	return networkInterfaceExportCmd
}
