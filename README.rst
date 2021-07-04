:code:`munch`
=====

A means of facilitating Worlds-to-Discord (and back) communication

Synopsis
--------

Currently, Munch is no more than a simple proof-of-concept of how a dummy
account can connect to a WorldServer and intercept information.

Planned Features:

- Worlds communication transmitted to Discord
- Discord communication transmitted to Worlds
- Webhook support

Once stable, the features of Munch will be integrated into the
`Whirlsplash Discord bot <https://github.com/Whirlsplash/bot>`_.

Limitations
^^^^^^^^^^^

Munch never sends any positional information to the RoomServer once connected,
this means that she is never rendered as an avatar, this also means that she is
technically never present within a room (?) so she cannot intercept TEXT
commands, this will be fixed (keep in mind, this is just a proof-of-concept).

Notes
^^^^^

Munch also paves the way for a system of bots enabled by the
Whirlsplash-of-things. "What? ..."; Whirlsplash is planned to have an imperative
"programming" language of sorts which can be used to create bots for Whirl with
ease.

Development
-----------

Docker support coming soon!

Running
^^^^^^^

.. code-blocK:: shell

  $ make run

Building
^^^^^^^^

.. code-blocK:: shell

  $ make build

License
~~~~~~~

`GNU General Public License v3.0 <./LICENSE>`_
