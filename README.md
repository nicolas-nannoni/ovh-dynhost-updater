# OVH DynHost Updater #

This tool is here to help with updating the IP address of DynHost ([OVH DynDNS records](https://docs.ovh.com/fr/fr/web/domains/utilisation-dynhost/)) entries.

## I need help! ##
Just ask the program!

    ovh-dynhost-updater help

    NAME:
       ovh-dynhost-upgrader - Tool to perform DynHost DNS records for OVH domains.

    USAGE:
       ovh-dynhost-updater [global options] command [command options] [arguments...]

    VERSION:
       0.0.1 (build ae840b2bfcdfe7016d1ad0c7b4478350175c5f6b)

    AUTHOR(S):
       Nicolas Nannoni

    COMMANDS:
         update-record  Update a DynHost record
         help, h        Shows a list of commands or help for one command

    GLOBAL OPTIONS:
       --username value  The OVH DynHost username
       --password value  The OVH DynHost password
       --debug, -d       Enable debug mode
       --help, -h        show help
       --version, -v     print the version


## How to build and run? ##

    git clone git@github.com:nicolas-nannoni/ovh-dynhost-updater.git
    go get
    make

Binaries will be in the `bin` folder

### Cross compile

    make linux
    make osx
    make windows

On Linux:

    ./bin/ovh-dynhost-updater-linux

On Mac:

    ./bin/ovh-dynhost-updater-osx

On Windows:

    ./bin/ovh-dynhost-updater.exe

Or to just run it without building:

    make run