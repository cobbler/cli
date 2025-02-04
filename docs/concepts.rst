******************
General Principles
******************

This should just be a brief overview. For the detailed explanations please refer to
`Readthedocs <https://cobbler.readthedocs.io/>`_.

Distros, Profiles and Systems
#############################

Cobbler has a system of inheritance when it comes to managing the information you want to apply to a certain system.

Images
######

Repositories
############

Management Classes
##################

Deleting configuration entries
##############################

If you want to remove a specific object, use the remove command with the name that was used to add it.

.. code-block:: shell

    cobbler distro|profile|system|repo|image|menu remove --name=string

Editing
#######

If you want to change a particular setting without doing an ``add`` again, use the ``edit`` command, using the same name
you gave when you added the item. Anything supplied in the parameter list will overwrite the settings in the existing
object, preserving settings not mentioned.

.. code-block:: shell

    cobbler distro|profile|system|repo|image|menu edit --name=string [parameterlist]

Copying
#######

Objects can also be copied:

.. code-block:: shell

    cobbler distro|profile|system|repo|image|menu copy --name=oldname --newname=newname

Renaming
########

Objects can also be renamed, as long as other objects don't reference them.

.. code-block:: shell

    cobbler distro|profile|system|repo|image|menu rename --name=oldname --newname=newname
