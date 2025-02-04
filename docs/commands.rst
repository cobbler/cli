************
CLI-Commands
************

Short Usage: ``cobbler command [subcommand] [--arg1=value1] [--arg2=value2]``

Long Usage:

.. code-block:: shell

    cobbler <distro|profile|system|repo|image|menu> ... [add|edit|copy|get-autoinstall*|list|remove|rename|report] [options|--help]
    cobbler <aclsetup|buildiso|import|list|mkloaders|replicate|report|reposync|sync|validate-autoinstalls|version|signature|hardlink> [options|--help]

.. toctree::
   :maxdepth: 2

   cobbler aclsetup <commands/aclsetup>
   cobbler buildiso <commands/buildiso>
   cobbler distro <commands/distro>
   cobbler hardlink <commands/hardlink>
   cobbler image <commands/image>
   cobbler import <commands/import>
   cobbler list <commands/list>
   cobbler menu <commands/menu>
   cobbler mkloaders <commands/mkloaders>
   cobbler profile <commands/profile>
   cobbler replicate <commands/replicate>
   cobbler repo <commands/repo>
   cobbler report <commands/report>
   cobbler reposync <commands/reposync>
   cobbler signature <commands/signature>
   cobbler sync <commands/sync>
   cobbler system <commands/system>
   cobbler validate-autoinstalls <commands/validate-autoinstalls>
   cobbler version <commands/version>
