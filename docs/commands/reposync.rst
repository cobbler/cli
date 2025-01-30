****************
cobbler reposync
****************

Example:

.. code-block:: shell

    $ cobbler reposync [--only=ONLY] [--tries=TRIES] [--no-fail]

Cobbler reposync is the command to use to update repos as configured with ``cobbler repo add``. Mirroring can
take a long time, and usage of cobbler reposync prior to usage is needed to ensure provisioned systems have the
files they need to actually use the mirrored repositories. If you just add repos and never run ``cobbler reposync``,
the repos will never be mirrored. This is probably a command you would want to put on a crontab, though the
frequency of that crontab and where the output goes is left up to the systems administrator.

For those familiar with dnf’s reposync, cobbler’s reposync is (in most uses) a wrapper around the ``dnf reposync``
command. Please use ``cobbler reposync`` to update cobbler mirrors, as dnf’s reposync does not perform all required steps.
Also cobbler adds support for rsync and SSH locations, where as dnf’s reposync only supports what dnf supports
(http/ftp).

If you ever want to update a certain repository you can run:
``cobbler reposync --only="reponame1" ...``

When updating repos by name, a repo will be updated even if it is set to be not updated during a regular reposync
operation (ex: ``cobbler repo edit –name=reponame1 –keep-updated=0``).
Note that if a cobbler import provides enough information to use the boot server as a yum mirror for core packages,
cobbler can set up automatic installation files to use the cobbler server as a mirror instead of the outside world. If
this feature is desirable, it can be turned on by ``setting yum_post_install_mirror`` to ``True`` in
``/etc/cobbler/settings.yaml`` (and running ``cobbler sync``). You should not use this feature if machines are
provisioned on a different VLAN/network than production, or if you are provisioning laptops that will want to acquire
updates on multiple networks.

The flags ``--tries=N`` (for example, ``--tries=3``) and ``--no-fail`` should likely be used when putting re-posync on a
crontab. They ensure network glitches in one repo can be retried and also that a failure to synchronize one repo does
not stop other repositories from being synchronized.

