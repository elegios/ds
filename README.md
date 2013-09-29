ds (dotsync)
============

Dotsync uses the fact that there are a lot of programs that can sync folders easily to create a tool that helps syncing multiple dotfiles across computers. It moves all the files you want to sync into a single folder, puts symlinks in their place, and keeps track of which files used to be where.

This information can later be used to clone the setup to another computer. As long as the folder used is synced (via Dropbox or something similar) the files should be kept in sync.


Usage
-----

    ds init <a synced folder>
    ds add <a bunch of dotfiles>

This creates a folder, moves your dotfiles there, and replaces them with symlinks to their new locations. It also creates a profile, with the same name as your computers hostname, which you can use to quickly set up another computer:

    ds clone <same synced folder> <hostname of the first computer>

This replaces any pre-existing dotfiles with symlinks to those in the dotsync folder.
