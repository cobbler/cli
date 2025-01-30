**************
cobbler system
**************

System records map a piece of hardware (or a virtual machine) with the Cobbler profile to be assigned to run on it. This
may be thought of as choosing a role for a specific system.

Note that if provisioning via Koan and PXE menus alone, it is not required to create system records in Cobbler, though
they are useful when system specific customizations are required. One such customization would be defining the MAC
address. If there is a specific role intended for a given machine, system records should be created for it.

System commands have a wider variety of control offered over network details. In order to use these to the fullest
possible extent, the automatic installation template used by Cobbler must contain certain automatic installation
snippets (sections of code specifically written for Cobbler to make these values become reality). Compare your automatic
installation templates with the stock ones in ``/var/lib/cobbler/templates`` if you have upgraded, to make sure
you can take advantage of all options to their fullest potential. If you are a new Cobbler user, base your automatic
installation templates off of these templates.

Read more about networking setup at: `Advanced Networking`_

Example:

.. code-block:: shell

    $ cobbler system add --name=string --profile=string [--mac=macaddress] [--ip-address=ipaddress] [--hostname=hostname] [--kernel-options=string] [--autoinstall-meta=string] [--autoinstall=path] [--netboot-enabled=Y/N] [--server=string] [--gateway=string] [--dns-name=string] [--static-routes=string] [--power-address=string] [--power-type=string] [--power-user=string] [--power-pass=string] [--power-id=string]

Adds a Cobbler System to the configuration. Arguments are specified as per "profile add" with the following changes:

+---------------------+------------------------------------------------------------------------------------------------+
| Name                | Description                                                                                    |
+=====================+================================================================================================+
| autoinstall         | While it is recommended that the ``--autoinstall`` parameter is only used within for the       |
|                     | "profile add" command, there are limited scenarios when an install base switching to Cobbler   |
|                     | may have legacy automatic installation files created on aper-system basis (one automatic       |
|                     | installation file for each system, nothing shared) and may not want to immediately make use of |
|                     | the Cobbler templating system. This allows specifying a automatic installation file for use on |
|                     | a per-system basis. Creation of a parent profile is still required. If the automatic           |
|                     | installation file is a filesystem location, it will still be treated as a Cobbler template.    |
+---------------------+------------------------------------------------------------------------------------------------+
| autoinstall-meta    | Automatic Installation Metadata (Ex: `dog=fang agent=86`).                                     |
+---------------------+------------------------------------------------------------------------------------------------+
| boot-files          | TFTP Boot Files (Files copied into tftpboot beyond the kernel/initrd).                         |
+---------------------+------------------------------------------------------------------------------------------------+
| boot-loaders        | Boot loader space delimited list (Network installation boot loaders).                          |
|                     | Valid options for list items are <<inherit>>, `grub`, `pxe`, `ipxe`.                           |
+---------------------+------------------------------------------------------------------------------------------------+
| comment             | Simple attach a description (Free form text) to your distro.                                   |
+---------------------+------------------------------------------------------------------------------------------------+
| dhcp-tag            | If you are setting up a PXE environment with multiple subnets/gateways, and are using Cobbler  |
|                     | to manage a DHCP configuration, you will probably want to use this option. If not, it can be   |
|                     | ignored.                                                                                       |
|                     |                                                                                                |
|                     | By default, the dhcp tag for all systems is "default" and means that in the DHCP template      |
|                     | files the systems will expand out where $insert_cobbler_systems_definitions is found in the    |
|                     | DHCP template. However, you may want certain systems to expand out in other places in the DHCP |
|                     | config file. Setting ``--dhcp-tag=subnet2`` for instance, will cause that system to expand out |
|                     | where $insert_cobbler_system_definitions_subnet2 is found, allowing you to insert directives   |
|                     | to specify different subnets (or other parameters) before the DHCP configuration entries for   |
|                     | those particular systems.                                                                      |
|                     |                                                                                                |
|                     | This is described further on the Cobbler Wiki.                                                 |
+---------------------+------------------------------------------------------------------------------------------------+
| dns-name            | If using the DNS management feature (see advanced section -- Cobbler supports auto-setup of    |
|                     | BIND and dnsmasq), use this to define a hostname for the system to receive from DNS.           |
|                     |                                                                                                |
|                     | Example: ``--dns-name=mycomputer.example.com``                                                 |
|                     |                                                                                                |
|                     | This is a per-interface parameter. If you have multiple interfaces, it may be different for    |
|                     | each interface, for example, assume a DMZ / dual-homed setup.                                  |
+---------------------+------------------------------------------------------------------------------------------------+
| enable-ipxe         | Enable iPXE? (Use iPXE instead of PXELINUX for advanced booting options)                       |
+---------------------+------------------------------------------------------------------------------------------------+
| fetchable-files     | Fetchable Files (Templates for tftp or wget/curl)                                              |
+---------------------+------------------------------------------------------------------------------------------------+
| filename            | This parameter can be used to select the bootloader for network boot. If specified, this must  |
|                     | be a path relative to the TFTP servers root directory. (e.g. grub/grubx64.efi)                 |
|                     | For most use cases the default bootloader is correct and this can be omitted                   |
+---------------------+------------------------------------------------------------------------------------------------+
| gateway and netmask | If you are using static IP configurations and the interface is flagged ``--static=1``, these   |
|                     | will be applied.                                                                               |
|                     |                                                                                                |
|                     | Netmask is a per-interface parameter. Because of the way gateway is stored on the installed OS,|
|                     | gateway is a global parameter. You may use ``--static-routes`` for per-interface customizations|
|                     | if required.                                                                                   |
+---------------------+------------------------------------------------------------------------------------------------+
| hostname            | This field corresponds to the hostname set in a systems ``/etc/sysconfig/network`` file. This  |
|                     | has no bearing on DNS, even when manage_dns is enabled. Use ``--dns-name`` instead for that    |
|                     | feature.                                                                                       |
|                     |                                                                                                |
|                     | This parameter is assigned once per system, it is not a per-interface setting.                 |
+---------------------+------------------------------------------------------------------------------------------------+
| interface           | By default flags like ``--ip``, ``--mac``, ``--dhcp-tag``, ``--dns-name``, ``--netmask``,      |
|                     | ``--virt-bridge``, and ``--static-routes`` operate on the first network interface defined for  |
|                     | a system (eth0).                                                                               |
|                     | However, Cobbler supports an arbitrary number of interfaces. Using ``--interface=eth1`` for    |
|                     | instance, will allow creating and editing of a second interface.                               |
|                     |                                                                                                |
|                     | Interface naming notes:                                                                        |
|                     |                                                                                                |
|                     | Additional interfaces can be specified (for example: eth1, or any name you like, as long as it |
|                     | does not conflict with any reserved names such as kernel module names) for use with the edit   |
|                     | command. Defining VLANs this way is also supported, of you want to add VLAN 5 on interface     |
|                     | eth0, simply name your interface eth0.5.                                                       |
|                     |                                                                                                |
|                     | Example:                                                                                       |
|                     |                                                                                                |
|                     | cobbler system edit --name=foo --ip-address=192.168.1.50 --mac=AA:BB:CC:DD:EE:A0               |
|                     |                                                                                                |
|                     | cobbler system edit --name=foo --interface=eth0 --ip-address=10.1.1.51 --mac=AA:BB:CC:DD:EE:A1 |
|                     |                                                                                                |
|                     | cobbler system report foo                                                                      |
|                     |                                                                                                |
|                     | Interfaces can be deleted using the --delete-interface option.                                 |
|                     |                                                                                                |
|                     | Example:                                                                                       |
|                     |                                                                                                |
|                     | cobbler system edit --name=foo --interface=eth2 --delete-interface                             |
+---------------------+------------------------------------------------------------------------------------------------+
| interface-type,     | One of the other advanced networking features supported by Cobbler is NIC bonding, bridging    |
| interface-master,   | and BMC. You can use this to bond multiple physical network interfaces to one single logical   |
| bonding-opts,       | interface to reduce single points of failure in your network, to create bridged interfaces for |
| bridge-opts         | things like tunnels and virtual machine networks, or to manage BMC interface by DHCP.          |
|                     | Supported values for the ``--interface-type`` parameter are "bond", "bond_slave", "bridge",    |
|                     | "bridge_slave","bonded_bridge_slave" and "bmc". If one of the "_slave" options is specified,   |
|                     | you also need to define the master-interface for this bond using                               |
|                     | ``--interface-master=INTERFACE``. Bonding and bridge options for the master-interface may be   |
|                     | specified using ``--bonding-opts="foo=1 bar=2"`` or ``--bridge-opts="foo=1 bar=2"``.           |
|                     |                                                                                                |
|                     | Example:                                                                                       |
|                     |                                                                                                |
|                     | .. code-block:: shell-session                                                                  |
|                     |                                                                                                |
|                     |     cobbler system edit --name=foo \                                                           |
|                     |                         --interface=eth0 \                                                     |
|                     |                         --mac=AA:BB:CC:DD:EE:00 \                                              |
|                     |                         --interface-type=bond_slave \                                          |
|                     |                         --interface-master=bond0                                               |
|                     |     cobbler system edit --name=foo \                                                           |
|                     |                         --interface=eth1 \                                                     |
|                     |                         --mac=AA:BB:CC:DD:EE:01 \                                              |
|                     |                         --interface-type=bond_slave \                                          |
|                     |                         --interface-master=bond0                                               |
|                     |     cobbler system edit --name=foo \                                                           |
|                     |                         --interface=bond0 \                                                    |
|                     |                         --interface-type=bond \                                                |
|                     |                         --bonding-opts="mode=active-backup miimon=100" \                       |
|                     |                         --ip-address=192.168.0.63 \                                            |
|                     |                         --netmask=255.255.255.0 \                                              |
|                     |                         --gateway=192.168.0.1 \                                                |
|                     |                         --static=1                                                             |
|                     |                                                                                                |
|                     | More information about networking setup is available at `Advanced Networking`_                 |
|                     |                                                                                                |
|                     | To review what networking configuration you have for any object, run "cobbler system report"   |
|                     | at any time:                                                                                   |
|                     |                                                                                                |
|                     | Example:                                                                                       |
|                     |                                                                                                |
|                     | .. code-block:: shell-session                                                                  |
|                     |                                                                                                |
|                     |     cobbler system report --name=foo                                                           |
|                     |                                                                                                |
+---------------------+------------------------------------------------------------------------------------------------+
| if-gateway          | If you are using static IP configurations and have multiple interfaces, use this to define     |
|                     | different gateway for each interface.                                                          |
|                     |                                                                                                |
|                     | This is a per-interface setting.                                                               |
+---------------------+------------------------------------------------------------------------------------------------+
| ip-address,         | If Cobbler is configured to generate a DHCP configuration (see advanced section), use this     |
| ipv6-address        | setting to define a specific IP for this system in DHCP. Leaving off this parameter will       |
|                     | result in no DHCP management for this particular system.                                       |
|                     |                                                                                                |
|                     | Example: ``--ip-address=192.168.1.50``                                                         |
|                     |                                                                                                |
|                     | If DHCP management is disabled and the interface is labelled ``--static=1``, this setting will |
|                     | be used for static IP configuration.                                                           |
|                     |                                                                                                |
|                     | Special feature: To control the default PXE behavior for an entire subnet, this field can also |
|                     | be passed in using CIDR notation. If ``--ip`` is CIDR, do not specify any other arguments      |
|                     | other than ``--name`` and ``--profile``.                                                       |
|                     |                                                                                                |
|                     | When using the CIDR notation trick, don't specify any arguments other than ``--name`` and      |
|                     | ``--profile``, as they won't be used.                                                          |
+---------------------+------------------------------------------------------------------------------------------------+
| kernel-options      | Sets kernel command-line arguments that the distro, and profiles/systems depending on it, will |
|                     | use. To remove a kernel argument that may be added by a higher Cobbler object (or in the global|
|                     | settings), you can prefix it with a ``!``.                                                     |
|                     |                                                                                                |
|                     | Example: ``--kernel-options="foo=bar baz=3 asdf !gulp"``                                       |
|                     |                                                                                                |
|                     | This example passes the arguments ``foo=bar baz=3 asdf`` but will make sure ``gulp`` is not    |
|                     | passed even if it was requested at a level higher up in the Cobbler configuration.             |
+---------------------+------------------------------------------------------------------------------------------------+
| kernel-options-post | This is just like ``--kernel-options``, though it governs kernel options on the installed OS,  |
|                     | as opposed to kernel options fed to the installer. The syntax is exactly the same. This        |
|                     | requires some special snippets to be found in your automatic installation template in order    |
|                     | for this to work. Automatic installation templating is described later on in this document.    |
|                     |                                                                                                |
|                     | Example: ``noapic``                                                                            |
+---------------------+------------------------------------------------------------------------------------------------+
| mac,                | Specifying a mac address via ``--mac`` allows the system object to boot directly to a specific |
| mac-address         | profile via PXE, bypassing Cobbler's PXE menu. If the name of the Cobbler system already looks |
|                     | like a mac address, this is inferred from the system name and does not need to be specified.   |
|                     |                                                                                                |
|                     | MAC addresses have the format AA:BB:CC:DD:EE:FF. It's highly recommended to register your MAC  |
|                     | addresses in Cobbler if you're using static addressing with multiple interfaces, or if you are |
|                     | using any of the advanced networking features like bonding, bridges or VLANs.                  |
|                     |                                                                                                |
|                     | Cobbler does contain a feature (enabled in /etc/cobbler/settings.yaml) that can automatically  |
|                     | add new system records when it finds profiles being provisioned on hardware it has seen before.|
|                     | This may help if you do not have a report of all the MAC addresses in your datacenter/lab      |
|                     | configuration.                                                                                 |
+---------------------+------------------------------------------------------------------------------------------------+
| mgmt-classes        | Management Classes (Management classes for external config management).                        |
+---------------------+------------------------------------------------------------------------------------------------+
| mgmt-parameters     | Management Parameters which will be handed to your management application.                     |
|                     | (Must be valid YAML dictionary)                                                                |
+---------------------+------------------------------------------------------------------------------------------------+
| **name**            | The system name works like the name option for other commands.                                 |
|                     |                                                                                                |
|                     | If the name looks like a MAC address or an IP, the name will implicitly be used for either     |
|                     | ``--mac`` or ``--ip`` of the first interface, respectively. However, it's usually better to    |
|                     | give a descriptive name -- don't rely on this behavior.                                        |
|                     |                                                                                                |
|                     | A system created with name "default" has special semantics. If a default system object exists, |
|                     | it sets all undefined systems to PXE to a specific profile. Without a "default" system name    |
|                     | created, PXE will fall through to local boot for unconfigured systems.                         |
|                     |                                                                                                |
|                     | When using "default" name, don't specify any other arguments than ``--profile``, as they won't |
|                     | be used.                                                                                       |
+---------------------+------------------------------------------------------------------------------------------------+
| name-servers        | If your nameservers are not provided by DHCP, you can specify a space separated list of        |
|                     | addresses here to configure each of the installed nodes to use them (provided the automatic    |
|                     | installation files used are installed on a per-system basis). Users with DHCP setups should not|
|                     | need to use this option. This is available to set in profiles to avoid having to set it        |
|                     | repeatedly for each system record.                                                             |
+---------------------+------------------------------------------------------------------------------------------------+
| name-servers-search | You can specify a space separated list of domain names to configure each of the installed nodes|
|                     | to use them as domain search path. This is available to set in profiles to avoid having to set |
|                     | it repeatedly for each system record.                                                          |
+---------------------+------------------------------------------------------------------------------------------------+
| netboot-enabled     | If set false, the system will be provisionable through Koan but not through standard PXE.      |
|                     | This will allow the system to fall back to default PXE boot behavior without deleting the      |
|                     | Cobbler system object. The default value allows PXE. Cobbler contains a PXE boot loop          |
|                     | prevention feature (pxe_just_once, can be enabled in /etc/cobbler/settings.yaml) that can      |
|                     | automatically trip off this value after a system gets done installing. This can prevent        |
|                     | installs from appearing in an endless loop when the system is set to PXE first in the BIOS     |
|                     | order.                                                                                         |
+---------------------+------------------------------------------------------------------------------------------------+
| next-server         | To override the Next server.                                                                   |
+---------------------+------------------------------------------------------------------------------------------------+
| owners              | Users with small sites and a limited number of admins can probably ignore this option. All     |
|                     | objects (distros, profiles, systems, and repos) can take a --owners parameter to specify what  |
|                     | Cobbler users can edit particular objects.This only applies to the Cobbler WebUI and XML-RPC   |
|                     | interface, not the "cobbler" command line tool run from the shell. Furthermore, this is only   |
|                     | respected by the ``authorization.ownership`` module which must be enabled in                   |
|                     | the settings. The value for ``--owners`` is a space separated list of users                    |
|                     | and groups as specified in ``/etc/cobbler/users.conf``.                                        |
|                     | For more information see the users.conf file as well as the Cobbler                            |
|                     | Wiki. In the default Cobbler configuration, this value is completely ignored, as is            |
|                     | ``users.conf``.                                                                                |
+---------------------+------------------------------------------------------------------------------------------------+
| power-address,      | Cobbler contains features that enable integration with power management for easier             |
| power-type,         | installation, reinstallation, and management of machines in a datacenter environment. These    |
| power-user,         | parameters are described online at `power-management`. If you have a power-managed             |
| power-pass,         | datacenter/lab setup, usage of these features may be something you are interested in.          |
| power-id,           |                                                                                                |
| power-options,      |                                                                                                |
| power-identity-file |                                                                                                |
+---------------------+------------------------------------------------------------------------------------------------+
| **profile**         | The name of Cobbler profile the system will inherit its properties.                            |
+---------------------+------------------------------------------------------------------------------------------------+
| proxy               | Proxy URL.                                                                                     |
+---------------------+------------------------------------------------------------------------------------------------+
| redhat-             | Management Classes (Management classes for external config management).                        |
| management-key      |                                                                                                |
+---------------------+------------------------------------------------------------------------------------------------+
| repos-enabled       | If set true, Koan can reconfigure repositories after installation. This is described further   |
|                     | on the Cobbler Wiki,https://github.com/cobbler/cobbler/wiki/Manage-yum-repos.                  |
+---------------------+------------------------------------------------------------------------------------------------+
| static              | Indicates that this interface is statically configured. Many fields (such as gateway/netmask)  |
|                     | will not be used unless this field is enabled.                                                 |
|                     |                                                                                                |
|                     | This is a per-interface setting.                                                               |
+---------------------+------------------------------------------------------------------------------------------------+
| static-routes       | This is a space delimited list of ip/mask:gateway routing information in that format.          |
|                     | Most systems will not need this information.                                                   |
|                     |                                                                                                |
|                     | This is a per-interface setting.                                                               |
+---------------------+------------------------------------------------------------------------------------------------+
| virt-auto-boot      | (Virt-only) Virt Auto Boot (Auto boot this VM?).                                               |
+---------------------+------------------------------------------------------------------------------------------------+
| virt-bridge         | (Virt-only) This specifies the default bridge to use for all systems defined under this        |
|                     | profile. If not specified, it will assume the default value in the Cobbler settings file, which|
|                     | as shipped in the RPM is ``virbr0``. If no using NAT, this is most likely not correct. You may |
|                     | want to override this setting in the system object. Bridge settings are important as they      |
|                     | define how outside networking will reach the guest. For more information on bridge setup, see  |
|                     | the Cobbler Wiki, where there is a section describing Koan usage.                              |
+---------------------+------------------------------------------------------------------------------------------------+
| virt-cpus           | (Virt-only) How many virtual CPUs should Koan give the virtual machine? The default is 1. This |
|                     | is an integer.                                                                                 |
+---------------------+------------------------------------------------------------------------------------------------+
| virt-disk-driver    | (Virt-only) Virt Disk Driver Type (The on-disk format for the virtualization disk).            |
|                     | Valid options are <<inherit>>, `raw`, `qcow2`, `qed`, `vdi`, `vmdk`                            |
+---------------------+------------------------------------------------------------------------------------------------+
| virt-file-size      | (Virt-only) How large the disk image should be in Gigabytes. The default is 5. This can be a   |
|                     | comma separated list (ex: ``5,6,7``) to allow for multiple disks of different sizes depending  |
|                     | on what is given to ``--virt-path``. This should be input as a integer or decimal value without|
|                     | units.                                                                                         |
+---------------------+------------------------------------------------------------------------------------------------+
| virt-path           | (Virt-only) Where to store the virtual image on the host system. Except for advanced cases,    |
|                     | this parameter can usually be omitted. For disk images, the value is usually an absolute path  |
|                     | to an existing directory with an optional filename component. There is support for specifying  |
|                     | partitions ``/dev/sda4`` or volume groups ``VolGroup00``, etc.                                 |
|                     |                                                                                                |
|                     | For multiple disks, separate the values with commas such as ``VolGroup00,VolGroup00`` or       |
|                     | ``/dev/sda4,/dev/sda5``. Both those examples would create two disks for the VM.                |
+---------------------+------------------------------------------------------------------------------------------------+
| virt-ram            | (Virt-only) How many megabytes of RAM to consume. The default is 512 MB. This should be input  |
|                     | as an integer without units.                                                                   |
+---------------------+------------------------------------------------------------------------------------------------+
| virt-type           | (Virt-only) Koan can install images using either Xen paravirt (``xenpv``) or QEMU/KVM          |
|                     | (``qemu``/``kvm``). Choose one or the other strings to specify, or values will default to      |
|                     | attempting to find a compatible installation type on the client system("auto"). See the "Koan" |
|                     | manpage for more documentation. The default ``--virt-type`` can be configured in the Cobbler   |
|                     | settings file such that this parameter does not have to be provided. Other virtualization types|
|                     | are supported, for information on those options (such as VMware), see the Cobbler Wiki.        |
+---------------------+------------------------------------------------------------------------------------------------+

.. _Advanced Networking: https://cobbler.readthedocs.io/en/latest/user-guide/advanced-networking.html
