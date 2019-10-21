# Generator
Resource files generator.

## Notes
Work in progress.
Not ready to be used.

## Install

```shell
$ go get -u https://gitlab.com/mikrowezel/backend/generator
$ alias mw=generator
```

Another option to avoid alias

```shell
$ git clone https://gitlab.com/mikrowezel/backend/generator.git
$ cd generator
$ make install
```

## Usage
```shell
$ mw help
$ mw help [cmd]
...

$ mw cmd target -using=resource.yaml
...
```

**Where**

  * **cmd** is the name of the command.
    * **help:** Shows main help
    * **generate:** Shows help for generate command

  * **target** is one of this list:
    * **handler:** generates a handler
    * **migration:** generates a migration
    * **model:** generates a model
    * **repo:** generates a repo handler
    * **test:** generates a test for handler
    * **all:** generates all resource files

