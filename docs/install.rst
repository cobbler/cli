*************************************
Installing the Cobbler standalone CLI
*************************************

.. note::
   Before version 1.0 is released, expect information on this page to be a WIP.

.. warning::
   The standalone Cobbler CLI will conflict with the bundled CLI of the Cobbler daemon with version 3.3.x and older.

RPM
###

#. Add the OBS Repository for your operating system
#. Enable and refresh the previously added repository
#. Install the binary package ``cobbler-cli``

DEB
###

#. Add the OBS Repository for your operating system
#. Enable and refresh the previously added repository
#. Install the binary package ``cobbler-cli``

Docker
######

.. code-block:: shell-session

   docker run --rm ghcr.io/cobbler/cli:latest


.. code-block:: shell-session

   podman run --rm ghcr.io/cobbler/cli:latest

Source
######

.. code-block:: shell-session

   make build
