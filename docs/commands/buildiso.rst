****************
cobbler buildiso
****************

This command may not behave like you expect it without installing additional dependencies and configuration. The in
depth explanation can be found at
`Building ISOs <https://cobbler.readthedocs.io/en/latest/user-guide/building-isos.html>`_.

.. note:: Systems refers to systems that are profile based. Systems with a parent image based systems will be skipped.

+--------------+-------------------------------------------------------------------------------------------------------+
| Name         | Description                                                                                           |
+--------------+-------------------------------------------------------------------------------------------------------+
| iso          | Output ISO to this file. If the file exists it will be truncated to zero before.                      |
+--------------+-------------------------------------------------------------------------------------------------------+
| profiles     | Use these profiles only for information collection.                                                   |
+--------------+-------------------------------------------------------------------------------------------------------+
| systems      | (net-only) Use these systems only for information collection.                                         |
+--------------+-------------------------------------------------------------------------------------------------------+
| tempdir      | Working directory for building the ISO. The default value is set in the settings file.                |
+--------------+-------------------------------------------------------------------------------------------------------+
| distro       | Used to detect the architecture of the ISO you are building. Specifies also the used Kernel and       |
|              | Initrd.                                                                                               |
+--------------+-------------------------------------------------------------------------------------------------------+
| standalone   | (offline-only) Creates a standalone ISO with all required distribution files but without any added    |
|              | repositories.                                                                                         |
+--------------+-------------------------------------------------------------------------------------------------------+
| airgapped    | (offline-only) Implies --standalone but additionally includes repo files for disconnected system      |
|              | installations.                                                                                        |
+--------------+-------------------------------------------------------------------------------------------------------+
| source       | (offline-only) Used with --standalone or --airgapped to specify a source for the distribution files.  |
+--------------+-------------------------------------------------------------------------------------------------------+
| exclude-dns  | (net-only) Prevents addition of name server addresses to the kernel boot options.                     |
+--------------+-------------------------------------------------------------------------------------------------------+
| xorriso-opts | Extra options for xorriso.                                                                            |
+--------------+-------------------------------------------------------------------------------------------------------+

Example: The following command builds a single ISO file for all profiles and systems present under the distro `test`.

.. code-block:: shell

    $ cobbler buildiso --distro=test

