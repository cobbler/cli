**************
cobbler distro
**************

This first step towards configuring what you want to install is to add a distribution record to Cobbler's configuration.

If there is an rsync mirror, DVD, NFS, or filesystem tree available that you would rather ``import`` instead, skip down
to the documentation about the ``import`` command. It's really a lot easier to follow the import workflow -- it only
requires waiting for the mirror content to be copied and/or scanned. Imported mirrors also save time during install
since they don't have to hit external install sources.

If you want to be explicit with distribution definition, however, here's how it works:

.. code-block:: shell

    $ cobbler distro add --name=string --kernel=path --initrd=path [--kernel-options=string] [--kernel-options-post=string] [--autoinstall-meta=string] [--arch=i386|x86_64|ppc|ppc64|ppc64le|arm64] [--breed=redhat|debian|suse] [--template-files=string]

+-----------------+-----------------------------------------------------------------------------------------------------+
| Name            | Description                                                                                         |
+=================+=====================================================================================================+
| arch            | Sets the architecture for the PXE bootloader and also controls how Koan's ``--replace-self`` option |
|                 | will operate.                                                                                       |
|                 |                                                                                                     |
|                 | The default setting (``standard``) will use ``pxelinux``.                                           |
|                 |                                                                                                     |
|                 | ``x86`` and ``x86_64`` effectively do the same thing as standard.                                   |
|                 |                                                                                                     |
|                 | If you perform a ``cobbler import``, the arch field will be auto-assigned.                          |
+-----------------+-----------------------------------------------------------------------------------------------------+
| autoinstall-    | This is an advanced feature that sets automatic installation template variables to substitute, thus |
| meta            | enabling those files to be treated as templates. Templates are powered using Cheetah and are        |
|                 | described further along in this manpage as well as on the Cobbler Wiki.                             |
|                 |                                                                                                     |
|                 | Example: ``--autoinstall-meta="foo=bar baz=3 asdf"``                                                |
|                 |                                                                                                     |
|                 | See the section on "Kickstart Templating" for further information.                                  |
+-----------------+-----------------------------------------------------------------------------------------------------+
| boot-files      | TFTP Boot Files (Files copied into tftpboot beyond the kernel/initrd).                              |
+-----------------+-----------------------------------------------------------------------------------------------------+
| boot-loaders    | Boot loader space delimited list (Network installation boot loaders).                               |
|                 | Valid options for list items are <<inherit>>, `grub`, `pxe`, `ipxe`.                                |
+-----------------+-----------------------------------------------------------------------------------------------------+
| breed           | Controls how various physical and virtual parameters, including kernel arguments for automatic      |
|                 | installation, are to be treated. Defaults to ``redhat``, which is a suitable value for Fedora and   |
|                 | CentOS as well. It means anything Red Hat based.                                                    |
|                 |                                                                                                     |
|                 | There is limited experimental support for specifying "debian", "ubuntu", or "suse", which treats the|
|                 | automatic installation template file as a preseed/autoyast file format and changes the kernel       |
|                 | arguments appropriately. Support for other types of distributions is possible in the future. See the|
|                 | Wiki for the latest information about support for these distributions.                              |
|                 |                                                                                                     |
|                 | The file used for the answer file, regardless of the breed setting, is the value used for           |
|                 | ``--autoinstall`` when creating the profile.                                                        |
+-----------------+-----------------------------------------------------------------------------------------------------+
| comment         | Simple attach a description (Free form text) to your distro.                                        |
+-----------------+-----------------------------------------------------------------------------------------------------+
| fetchable-files | Fetchable Files (Templates for tftp or wget/curl)                                                   |
+-----------------+-----------------------------------------------------------------------------------------------------+
| **initrd**      | An absolute filesystem path to a initrd image.                                                      |
+-----------------+-----------------------------------------------------------------------------------------------------+
| **kernel**      | An absolute filesystem path to a kernel image.                                                      |
+-----------------+-----------------------------------------------------------------------------------------------------+
| kernel-options  | Sets kernel command-line arguments that the distro, and profiles/systems depending on it, will use. |
|                 | To remove a kernel argument that may be added by a higher Cobbler object (or in the global          |
|                 | settings), you can prefix it with a ``!``.                                                          |
|                 |                                                                                                     |
|                 | Example: ``--kernel-options="foo=bar baz=3 asdf !gulp"``                                            |
|                 |                                                                                                     |
|                 | This example passes the arguments ``foo=bar baz=3 asdf`` but will make sure ``gulp`` is not passed  |
|                 | even if it was requested at a level higher up in the Cobbler configuration.                         |
+-----------------+-----------------------------------------------------------------------------------------------------+
| kernel-options- | This is just like ``--kernel-options``, though it governs kernel options on the installed OS, as    |
| post            | opposed to kernel options fed to the installer. The syntax is exactly the same. This requires some  |
|                 | special snippets to be found in your automatic installation template in order for this to work.     |
|                 | Automatic installation templating is described later on in this document.                           |
|                 |                                                                                                     |
|                 | Example: ``noapic``                                                                                 |
+-----------------+-----------------------------------------------------------------------------------------------------+
| mgmt-classes    | Management Classes (Management classes for external config management).                             |
+-----------------+-----------------------------------------------------------------------------------------------------+
| **name**        | A string identifying the distribution, this should be something like ``rhel6``.                     |
+-----------------+-----------------------------------------------------------------------------------------------------+
| os-version      | Generally this field can be ignored. It is intended to alter some hardware setup for virtualized    |
|                 | instances when provisioning guests with Koan. The valid options for ``--os-version`` vary depending |
|                 | on what is specified for ``--breed``. If you specify an invalid option, the error message will      |
|                 | contain a list of valid OS versions that can be used. If you don't know the OS version or it does   |
|                 | not appear in the list, omitting this argument or using ``other`` should be perfectly fine. If you  |
|                 | don't encounter any problems with virtualized instances, this option can be safely ignored.         |
+-----------------+-----------------------------------------------------------------------------------------------------+
| owners          | Users with small sites and a limited number of admins can probably ignore this option. All Cobbler  |
|                 | objects (distros, profiles, systems, and repos) can take a --owners parameter to specify what       |
|                 | Cobbler users can edit particular objects.This only applies to the Cobbler WebUI and XML-RPC        |
|                 | interface, not the "cobbler" command line tool run from the shell. Furthermore, this is only        |
|                 | respected by the ``authorization.ownership`` module which must be enabled in the settings.          |
|                 | The value for ``--owners`` is a space separated list of users and groups as specified in            |
|                 | ``/etc/cobbler/users.conf``. For more information see the users.conf file as well as the Cobbler    |
|                 | Wiki. In the default Cobbler configuration, this value is completely ignored, as is ``users.conf``. |
+-----------------+-----------------------------------------------------------------------------------------------------+
| redhat-         | Management Classes (Management classes for external config management).                             |
| management-key  |                                                                                                     |
+-----------------+-----------------------------------------------------------------------------------------------------+
| remote-boot-    | A URL pointing to the installation initrd of a distribution. If the bootloader has this support,    |
| kernel          | it will directly download the kernel from this URL, instead of the directory of the TFTP client.    |
|                 | Note: The kernel (or initrd below) will still be copied into the image directory of the TFTP server.|
|                 | The above kernel parameter is still needed (e.g. to build iso images, etc.).                        |
|                 | The advantage of letting the boot loader retrieve the kernel/initrd directly is the support of      |
|                 | changing/updated distributions. E.g. openSUSE Tumbleweed is updated on the fly and if Cobbler would |
|                 | copy/cache the kernel/initrd in the TFTP directory, you would get a "kernel does not match          |
|                 | distribution" (or similar) error when trying to install.                                            |
+-----------------+-----------------------------------------------------------------------------------------------------+
| remote-boot-    | See remote-boot-kernel above.                                                                       |
| initrd          |                                                                                                     |
+-----------------+-----------------------------------------------------------------------------------------------------+
| template-files  | This feature allows Cobbler to be used as a configuration management system. The argument is a space|
|                 | delimited string of ``key=value`` pairs. Each key is the path to a template file, each value is the |
|                 | path to install the file on the system. This is described in further detail on the Cobbler Wiki and |
|                 | is implemented using special code in the post install. Koan also can retrieve these files from a    |
|                 | Cobbler server on demand, effectively allowing Cobbler to function as a lightweight templated       |
|                 | configuration management system.                                                                    |
+-----------------+-----------------------------------------------------------------------------------------------------+
