*************
cobbler image
*************

The primary and recommended use of Cobbler is to deploy systems by building them like from the OS manufacturer's
distribution, e.g Redhat kickstart. This method is generally easier to work with and provides an infrastructure which is
not only more sustainable but also much more flexible across varieties of hardware.

But Cobbler can also help with image-based booting, physically and virtually. Some manual use of other commands beyond what
is typically required of Cobbler may be needed to prepare images for use with this feature and the usage of these
commands varies substantially depending on the type of image.

For now we just have 1 example of using the "memdisk" image type:

Example:

.. code-block:: shell

    $ cobbler image

memdisk - Oracle / Sun Maintenance CD
-------------------------------------

The 'memdisk' image type can be used to PXE boot Oracle / Sun maintenance CDs.
`Their manual <https://docs.oracle.com/cd/E19121-01/sf.x2250/820-4593-12/AppB.html#50540564_72480>`_ gives details on
how to copy the image from a CD to a PXE server. The procedure is even easier with Cobbler since the system takes care
of most of it for you.

Take your ISO for the boot CD and mount it as a loopback mount somewhere on your Cobbler server then copy the
``boot.img`` file into your tftpboot directory. Then add an image of type ``memdisk`` which uses it. Right now the
following shell command will fail due to a known bug but the web interface can be used instead to add the image.

.. code-block:: shell

   > cobbler image add --name=MyName --image-type=memdisk --file=/tftpboot/oracle/SF2250/boot.img
   > usage: cobbler [options]
   >
   > cobbler: error: option --image-type: invalid choice: 'memdisk' (choose from 'iso', 'direct', 'virt-image')


Now just boot your machine from the network and select the image "MyName".

Memtest
-------

If installed Cobbler will put an entry into all of your PXE menus allowing you to run memtest on physical systems
without making changes in Cobbler. This can be handy for some simple diagnostics.

Steps to get memtest to show up in your PXE menus:

.. code-block:: shell

   $ zypper/dnf install memtest86+
   $ cobbler image add --name=memtest86+ --file=/path/to/memtest86+ --image-type=direct
   $ cobbler sync

Targeted Memtesting
-------------------

However if you already have a Cobbler system record for the system you won't get the menu. To solve this:

.. code-block:: shell

   cobbler image add --name=foo --file=/path/to/memtest86 --image-type=direct
   cobbler system edit --name=bar --mac=AA:BB:CC:DD:EE:FF --image=foo --netboot-enabled=1

The system will boot to memtest until you put it back to its original profile.

.. warning:: When restoring the system back from memtest, make sure you turn its netboot flag **off** if you have it
             set to PXE first in the BIOS order unless you want to reinstall the system!

.. code-block:: shell

   $ cobbler system edit --name=bar --profile=old_profile_name --netboot-enabled=0

If you do want to reinstall it after running memtest, use ``--netboot-enabled=true``.
