# Usage

## Sample input file

Create an input file for the resource to be created.
`assets/gen/resource.yaml` could be a giid location for it.


```yaml
---
  name: Resource
  #apiVer: v1
  #plural: Resources
  propDefs:
    - name: OwnerID
      type: uuid
      length: 36
      isKey: false
      isUnique: false
      ref:
        model: user
        property: ID
    - name: ParentID
      type: uuid
      length: 36
      isKey: false
      isUnique: false
      ref:
        model: account
        property: ID
```

## Run generator

```shell
$ mw help
$ mw help [cmd]
...

$ mw cmd target -using=resource.yaml
$ # i.e:
$ mw generate all assets/gen/resource.yaml
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

