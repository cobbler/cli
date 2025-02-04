*****************
cobbler mkloaders
*****************

This command is used for generating UEFI bootable GRUB 2 bootloaders. This command has no options and is configured via
the settings file of Cobbler. If available on the operating system Cobbler is running on, then this also generates
bootloaders for different architectures then the one of the system.

.. note:: This command should be executed every time the bootloader modules are being updated, running it more
          frequently does not help, running it less frequently will cause the bootloader to be possibly vulnerable.

Example:

.. code-block:: shell

    $ cobbler mkloaders
