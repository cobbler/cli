************
cobbler repo
************

Repository mirroring allows Cobbler to mirror not only install trees ("cobbler import" does this for you) but also
optional packages, 3rd party content, and even updates. Mirroring all of this content locally on your network will
result in faster, more up-to-date installations and faster updates. If you are only provisioning a home setup, this will
probably be overkill, though it can be very useful for larger setups (labs, datacenters, etc).

.. code-block:: shell

    $ cobbler repo add --mirror=url --name=string [--rpmlist=list] [--creatrepo-flags=string] [--keep-updated=Y/N] [--priority=number] [--arch=string] [--mirror-locally=Y/N] [--breed=yum|rsync|rhn] [--mirror_type=baseurl|mirrorlist|metalink]

+------------------+---------------------------------------------------------------------------------------------------+
| Name             | Description                                                                                       |
+==================+===================================================================================================+
| apt-components   | Apt Components (apt only) (ex: main restricted universe)                                          |
+------------------+---------------------------------------------------------------------------------------------------+
| apt-dists        | Apt Dist Names (apt only) (ex: precise precise-updates)                                           |
+------------------+---------------------------------------------------------------------------------------------------+
| arch             | Specifies what architecture the repository should use. By default the current system arch (of the |
|                  | server) is used,which may not be desirable. Using this to override the default arch allows        |
|                  | mirroring of source repositories(using ``--arch=src``).                                           |
+------------------+---------------------------------------------------------------------------------------------------+
| breed            | Ordinarily Cobbler's repo system will understand what you mean without supplying this parameter,  |
|                  | though you can set it explicitly if needed.                                                       |
+------------------+---------------------------------------------------------------------------------------------------+
| comment          | Simple attach a description (Free form text) to your distro.                                      |
+------------------+---------------------------------------------------------------------------------------------------+
| createrepo-flags | Specifies optional flags to feed into the createrepo tool, which is called when                   |
|                  | ``cobbler reposync`` is run for the given repository. The defaults are ``-c cache``.              |
+------------------+---------------------------------------------------------------------------------------------------+
| keep-updated     | Specifies that the named repository should not be updated during a normal "cobbler reposync". The |
|                  | repo may still be updated by name. The repo should be synced at least once before disabling this  |
|                  | feature. See "cobbler reposync" below.                                                            |
+------------------+---------------------------------------------------------------------------------------------------+
| **mirror**       | The address of the yum mirror. This can be an ``rsync://``-URL, an ssh location, or a ``http://`` |
|                  | or ``ftp://`` mirror location. Filesystem paths also work.                                        |
|                  |                                                                                                   |
|                  | The mirror address should specify an exact repository to mirror -- just one architecture and just |
|                  | one distribution. If you have a separate repo to mirror for a different arch, add that repo       |
|                  | separately.                                                                                       |
|                  |                                                                                                   |
|                  | Here's an example of what looks like a good URL:                                                  |
|                  |                                                                                                   |
|                  | - ``rsync://yourmirror.example.com/fedora-linux-core/updates/6/i386`` (for rsync protocol)        |
|                  | - ``http://mirrors.kernel.org/fedora/extras/6/i386/`` (for http)                                  |
|                  | - ``user@yourmirror.example.com/fedora-linux-core/updates/6/i386``  (for SSH)                     |
|                  |                                                                                                   |
|                  | Experimental support is also provided for mirroring RHN content when you need a fast local mirror.|
|                  | The mirror syntax for this is ``--mirror=rhn://channel-name`` and you must have entitlements for  |
|                  | this to work. This requires the Cobbler server to be installed on RHEL 5 or later. You will also  |
|                  | need a version of ``yum-utils`` equal or greater to 1.0.4.                                        |
+------------------+---------------------------------------------------------------------------------------------------+
| mirror-locally   | When set to ``N``, specifies that this yum repo is to be referenced directly via automatic        |
|                  | installation files and not mirrored locally on the Cobbler server. Only ``http://`` and ``ftp://``|
|                  | mirror urls are supported when using ``--mirror-locally=N``, you cannot use filesystem URLs.      |
+------------------+---------------------------------------------------------------------------------------------------+
| **name**         | This name is used as the save location for the mirror. If the mirror represented, say, Fedora     |
|                  | Core 6 i386 updates, a good name would be ``fc6i386updates``. Again, be specific.                 |
|                  |                                                                                                   |
|                  | This name corresponds with values given to the ``--repos`` parameter of ``cobbler profile add``.  |
|                  | If a profile has a ``--repos``-value that matches the name given here, that repo can be           |
|                  | automatically set up during provisioning (when supported) and installed systems will also use the |
|                  | boot server as a mirror (unless ``yum_post_install_mirror`` is disabled in the settings file). By |
|                  | default the provisioning server will act as a mirror to systems it installs, which may not be     |
|                  | desirable for laptop configurations, etc.                                                         |
|                  |                                                                                                   |
|                  | Distros that can make use of yum repositories during automatic installation include FC6 and later,|
|                  | RHEL 5 and later, and derivative distributions.                                                   |
|                  |                                                                                                   |
|                  | See the documentation on ``cobbler profile add`` for more information.                            |
+------------------+---------------------------------------------------------------------------------------------------+
| owners           | Users with small sites and a limited number of admins can probably ignore this option. All        |
|                  | objects (distros, profiles, systems, and repos) can take a --owners parameter to specify what     |
|                  | Cobbler users can edit particular objects.This only applies to the Cobbler WebUI and XML-RPC      |
|                  | interface, not the "cobbler" command line tool run from the shell. Furthermore, this is only      |
|                  | respected by the ``authorization.ownership`` module which must be enabled in                      |
|                  | the settings. The value for ``--owners`` is a space separated list of users                       |
|                  | and groups as specified in ``/etc/cobbler/users.conf``.                                           |
|                  | For more information see the users.conf file as well as the Cobbler                               |
|                  | Wiki. In the default Cobbler configuration, this value is completely ignored, as is               |
|                  | ``users.conf``.                                                                                   |
+---------------------+------------------------------------------------------------------------------------------------+
| priority         | Specifies the priority of the repository (the lower the number, the higher the priority), which   |
|                  | applies to installed machines using the repositories that also have the yum priorities plugin     |
|                  | installed. The default priority for the plugins 99, as is that of all Cobbler mirrored            |
|                  | repositories.                                                                                     |
+------------------+---------------------------------------------------------------------------------------------------+
| proxy            | Proxy URL.                                                                                        |
+---------------------+------------------------------------------------------------------------------------------------+
| rpm-list         | By specifying a space-delimited list of package names for ``--rpm-list``, one can decide to mirror|
|                  | only a part of a repo (the list of packages given, plus dependencies). This may be helpful in     |
|                  | conserving time/space/bandwidth. For instance, when mirroring FC6 Extras, it may be desired to    |
|                  | mirror just Cobbler and Koan, and skip all of the game packages. To do this, use                  |
|                  | ``--rpm-list="cobbler koan"``.                                                                    |
|                  |                                                                                                   |
|                  | This option only works for ``http://`` and ``ftp://`` repositories (as it is powered by           |
|                  | yumdownloader). It will be ignored for other mirror types, such as local paths and ``rsync://``   |
|                  | mirrors.                                                                                          |
+------------------+---------------------------------------------------------------------------------------------------+
| yumopts          | Sets values for additional yum options that the repo should use on installed systems. For instance|
|                  | if a yum plugin takes a certain parameter "alpha" and "beta", use something like                  |
|                  | ``--yumopts="alpha=2 beta=3"``.                                                                   |
+------------------+---------------------------------------------------------------------------------------------------+

.. code-block:: shell

    $ cobbler repo autoadd

Add enabled yum repositories from ``dnf repolist --enabled`` list. The repository names are generated using the
<repo id>-<releasever>-<arch> pattern (ex: fedora-32-x86_64). Existing repositories with such names are not overwritten.
