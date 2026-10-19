*************
cobbler check
*************

Check the Cobbler server configuration for common problems.

Example:

.. code-block:: shell

   $ cobbler check

This command asks the Cobbler server to inspect its own configuration (settings, installed dependencies,
boot-loader files, services, etc.) and reports suggestions for items you may want to fix. It does not change
anything on the server; it only lists potential problems for the administrator to review.
