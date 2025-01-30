**************
cobbler import
**************

.. note::
   When running Cobbler via systemd, you cannot mount the ISO to ``/tmp`` or a sub-folder of it because we are using the
   option `Private Temporary Directory`, to enhance the security of our application.

Example:

.. code-block:: shell

    $ cobbler import
