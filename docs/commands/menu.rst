************
cobbler menu
************

By default, Cobbler builds a single-level boot menu for profiles and images. To simplify navigation through a large number
of OS boot items, you can create `menu` objects and place any number of submenus, profiles, and images there. The menu is
hierarchical, to indicate the nesting of one submenu in another, you can use the `parent` property. If the `parent` property
for a submenu, or the `menu` property for a profile or images are not set or have an empty value, then the corresponding
element will be displayed in the top-level menu. If a submenu does not have descendants in the form of profiles or images,
then such a submenu will not be displayed in the boot menu.

.. code-block:: shell

    $ cobbler menu add --name=string [--display-name=string] [--parent=string]

+------------------+---------------------------------------------------------------------------------------------------+
| Name             | Description                                                                                       |
+==================+===================================================================================================+
| display-name     | This is a human-readable name to display in the boot menu.                                        |
+------------------+---------------------------------------------------------------------------------------------------+
| **name**         | This name can be used as a `--parent` for a submenu, or as a `--menu` for a profile or image.     |
+------------------+---------------------------------------------------------------------------------------------------+
| parent           | This value can be set to indicate the nesting of this submenu in another.                         |
+------------------+---------------------------------------------------------------------------------------------------+
