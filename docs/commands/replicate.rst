*****************
cobbler replicate
*****************

Cobbler can replicate configurations from a master Cobbler server. Each Cobbler server is still expected to have a
locally relevant ``/etc/cobbler/settings.yaml``, as this file is not synced.

This feature is intended for load-balancing, disaster-recovery, backup, or multiple geography support.

Cobbler can replicate data from a central server.

Objects that need to be replicated should be specified with a pattern, such as ``--profiles="webservers* dbservers*"``
or ``--systems="*.example.org"``. All objects matched by the pattern, and all dependencies of those objects matched by
the pattern (recursively) will be transferred from the remote server to the central server. This is to say if you intend
to transfer ``*.example.org`` and the definition of the systems have not changed, but a profile above them has changed,
the changes to that profile will also be transferred.

In the case where objects are more recent on the local server, those changes will not be overridden locally.

Common data locations will be rsync'ed from the master server unless ``--omit-data`` is specified.

To delete objects that are no longer present on the master server, use ``--prune``.

**Warning**: This will delete all object types not present on the remote server from the local server, and is recursive.
If you use prune, it is best to manage Cobbler centrally and not expect changes made on the slave servers to be
preserved. It is not currently possible to just prune objects of a specific type.

Example:

.. code-block:: shell

    $ cobbler replicate --master=cobbler.example.org [--distros=pattern] [--profiles=pattern] [--systems=pattern] [--repos-pattern] [--images=pattern] [--prune] [--omit-data]
