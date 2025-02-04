*************
Configuration
*************

General Notes
#############

You can configure the Cobbler CLI via various mechanisms. The order in which the values are overwritten is the
following:

* Defaults
* YAML configuration file
* Environment variables
* CLI flags

Keeping this order in mind, the defaults have the lowest priority and the CLI flags the highest.

YAML configuration file notes
#############################

The following paths are searched for the YAML configuration file:

* The explicitly given one in the ``-c``/``--config`` flag.
* The users home directory ``$HOME/.cobbler.yaml``

Settings
########

Cobbler Server URL
==================

The full URL that the CLI will use to connect to the API.

* Default: ``http://127.0.0.1/cobbler_api``
* YAML key: ``server_url``
* Env variable: ``SERVER_URL``
* CLI flag: None

Cobbler Server Username
=======================

The username that should be used for logging into the Cobbler server.

* Default: ``cobbler``
* YAML key: ``server_username``
* Env variable: ``SERVER_USERNAME``
* CLI flag: None

Cobbler Server Password
=======================

The password that should be used for logging into the Cobbler server.

* Default: ``cobbler``
* YAML key: ``server_password``
* Env variable: ``SERVER_PASSWORD``
* CLI flag: None
