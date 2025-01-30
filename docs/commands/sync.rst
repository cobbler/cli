************
cobbler sync
************

The sync command is very important, though very often unnecessary for most situations. It's primary purpose is to force
a rewrite of all configuration files, distribution files in the TFTP root, and to restart managed services. So why is it
unnecessary? Because in most common situations (after an object is edited, for example), Cobbler executes what is known
as a "lite sync" which rewrites most critical files.

When is a full sync required? When you are using ``manage_dhcpd`` (Managing DHCP) with systems that use static leases.
In that case, a full sync is required to rewrite the ``dhcpd.conf`` file and to restart the dhcpd service.

Cobbler sync is used to repair or rebuild the contents ``/tftpboot`` or ``/var/www/cobbler`` when something has changed
behind the scenes. It brings the filesystem up to date with the configuration as understood by Cobbler.

Sync should be run whenever files in ``/var/lib/cobbler`` are manually edited (which is not recommended except for the
settings file) or when making changes to automatic installation files. In practice, this should not happen often, though
running sync too many times does not cause any adverse effects.

If using Cobbler to manage a DHCP and/or DNS server (see the advanced section of this manpage), sync does need to be run
after systems are added to regenerate and reload the DHCP/DNS configurations. If you want to trigger the DHCP/DNS
regeneration only and do not want a complete sync, you can use ``cobbler sync --dhcp`` or ``cobbler sync --dns`` or the
combination of both.

``cobbler sync --systems`` is used to only write specific systems (must exists in backend storage) to the TFTP folder.
The expected pattern is a comma separated list of systems e.g. ``sys1.internal,sys2.internal,sys3.internal``.

.. note::
    Please note that at least once a full sync has to be run beforehand.

The sync process can also be kicked off from the web interface.

Example:

.. code-block:: shell

    $ cobbler sync
    $ cobbler sync [--systems=sys1.internal,sys2.internal,sys3.internal]
    $ cobbler sync [--dns]
    $ cobbler sync [--dhcp]
    $ cobbler sync [--dns --dhcp]
