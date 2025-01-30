***************
cobbler profile
***************

A profile associates a distribution to additional specialized options, such as a installation automation file. Profiles
are the core unit of provisioning and at least one profile must exist for every distribution to be provisioned. A
profile might represent, for instance, a web server or desktop configuration. In this way, profiles define a role to be
performed.

.. code-block:: shell

    $ cobbler profile add --name=string --distro=string [--autoinstall=path] [--kernel-options=string] [--autoinstall-meta=string] [--name-servers=string] [--name-servers-search=string] [--virt-file-size=gigabytes] [--virt-ram=megabytes] [--virt-type=string] [--virt-cpus=integer] [--virt-path=string] [--virt-bridge=string] [--server] [--parent=profile] [--filename=string]

Arguments are the same as listed for distributions, save for the removal of "arch" and "breed", and with the additions
listed below:

+---------------------+------------------------------------------------------------------------------------------------+
| Name                | Description                                                                                    |
+=====================+================================================================================================+
| autoinstall         | Local filesystem path to a automatic installation file, the file must reside under             |
|                     | ``/var/lib/cobbler/templates``                                                                 |
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
| dhcp-tag            | DHCP Tag (see description in system).                                                          |
+---------------------+------------------------------------------------------------------------------------------------+
| **distro**          | The name of a previously defined Cobbler distribution. This value is required.                 |
+---------------------+------------------------------------------------------------------------------------------------+
| enable-ipxe         | Enable iPXE? (Use iPXE instead of PXELINUX for advanced booting options)                       |
+---------------------+------------------------------------------------------------------------------------------------+
| enable-menu         | Enable PXE Menu? (Show this profile in the PXE menu?)                                          |
+---------------------+------------------------------------------------------------------------------------------------+
| fetchable-files     | Fetchable Files (Templates for tftp or wget/curl)                                              |
+---------------------+------------------------------------------------------------------------------------------------+
| filename            | This parameter can be used to select the bootloader for network boot. If specified, this must  |
|                     | be a path relative to the TFTP servers root directory. (e.g. grub/grubx64.efi)                 |
|                     | For most use cases the default bootloader is correct and this can be omitted                   |
+---------------------+------------------------------------------------------------------------------------------------+
| menu                | This is a way of organizing profiles and images in an automatically generated boot menu for    |
|                     | `grub`, `pxe` and `ipxe` boot loaders. Menu created with ``cobbler menu add`` command.         |
+---------------------+------------------------------------------------------------------------------------------------+
| **name**            | A descriptive name. This could be something like ``rhel5webservers`` or ``f9desktops``.        |
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
| parent              | This is an advanced feature.                                                                   |
|                     |                                                                                                |
|                     | Profiles may inherit from other profiles in lieu of specifying ``--distro``. Inherited profiles|
|                     | will override any settings specified in their parent, with the exception of                    |
|                     | ``--autoinstall-meta`` (templating) and ``--kernel-options`` (kernel options), which will be   |
|                     | blended together.                                                                              |
|                     |                                                                                                |
|                     | Example: If profile A has ``--kernel-options="x=7 y=2"``, B inherits from A, and B has         |
|                     | ``--kernel-options="x=9 z=2"``, the actual kernel options that will be used for B are          |
|                     | ``x=9 y=2 z=2``.                                                                               |
|                     |                                                                                                |
|                     | Example: If profile B has ``--virt-ram=256`` and A has ``--virt-ram=512``, profile B will use  |
|                     | the value 256.                                                                                 |
|                     |                                                                                                |
|                     | Example: If profile A has a ``--virt-file-size=5`` and B does not specify a size, B will use   |
|                     | the value from A.                                                                              |
+---------------------+------------------------------------------------------------------------------------------------+
| proxy               | Proxy URL.                                                                                     |
+---------------------+------------------------------------------------------------------------------------------------+
| redhat-             | Management Classes (Management classes for external config management).                        |
| management-key      |                                                                                                |
+---------------------+------------------------------------------------------------------------------------------------+
| repos               | This is a space delimited list of all the repos (created with ``cobbler repo add`` and updated |
|                     | with ``cobbler reposync``)that this profile can make use of during automated installation. For |
|                     | example, an example might be ``--repos="fc6i386updates fc6i386extras"`` if the profile wants to|
|                     | access these two mirrors that are already mirrored on the Cobbler server. Repo management is   |
|                     | described in greater depth later in the manpage.                                               |
+---------------------+------------------------------------------------------------------------------------------------+
| server              | This parameter should be useful only in select circumstances. If machines are on a subnet that |
|                     | cannot access the Cobbler server using the name/IP as configured in the Cobbler settings file, |
|                     | use this parameter to override that servername. See also ``--dhcp-tag`` for configuring the    |
|                     | next server and DHCP information of the system if you are also using Cobbler to help manage    |
|                     | your DHCP configuration.                                                                       |
+---------------------+------------------------------------------------------------------------------------------------+
| template-files      | This feature allows Cobbler to be used as a configuration management system. The argument is a |
|                     | space delimited string of ``key=value`` pairs. Each key is the path to a template file, each   |
|                     | value is the path to install the file on the system. This is described in further detail on    |
|                     | the Cobbler Wiki and is implemented using special code in the post install. Koan also can      |
|                     | retrieve these files from a Cobbler server on demand, effectively allowing Cobbler to function |
|                     | as a lightweight templated configuration management system.                                    |
+---------------------+------------------------------------------------------------------------------------------------+
| virt-auto-boot      | (Virt-only) Virt Auto Boot (Auto boot this VM?).                                               |
+---------------------+------------------------------------------------------------------------------------------------+
| virt-bridge         | (Virt-only) This specifies the default bridge to use for all systems defined under this        |
|                     | profile. If not specified, it will assume the default value in the Cobbler settings file, which|
|                     | as shipped in the RPM is ``virbr0``. If not using NAT, this is most likely not correct. You may|
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
