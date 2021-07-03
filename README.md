# Overview

This repository demonstrates a Go workspace with multiple Go modules. In this repository, there is an module named [myapp][1] and a module named formally [github.com/paulwizviz/go-workspace/mylib][2] and informally `mylib`.

## The `myapp` module

[myapp][1] has a file structure to support the creation of on one or more executable applications. The structure that follows two aspects of the [Standard Go Project Layout](https://github.com/golang-standards/project-layout):

```
myapp/
  cmd/
    app1/
      main.go
  internal/
    mathutil/
```

The `cmd` folder host packages to create lanuch points (i.e. main) for one or more executable applications packages. In this example, an application name `app1` is packaged under a folder with the same name.

The `internal` folder host packages of non-exportable utilies or common code based for intended for consumption by packages found in `cmd` folder.

Since [myapp][1] module is never intended to share codes with external modules (including [mylib][2] there is no need to assign a full qualified domain name so this module in `go.mod` is simply named [myapp[][1].

> NOTE: Beware of potential namespace clashes particularly with Go standard packages.

## The `mylib` module

[mylib][2] is an externally accessible Go module. This means its content can be accessed by modules outside [mylib][2]. You could view this as a public library. By convention, any Go packages or files are hosted under the root folder of the module -- i.e. `mylib`.

Since this is externally accessible this module follow the Go convention of using a fully qualified url name -- i.e. [github.com/paulwizviz/go-workspace/mylib][2].

## Modules lifecycle management

Whilst [myapp][1] and [mylib][2] share the same workspace, you can manage their lifecycle independently. Use the following convention to Git tag the repo:

1) `myapp` tag it with this naming convention `myapp/v<semver>` e.g. `myapp/v0.0.1`

2) `mylib` tag it with this naming convention `mylib/<semver>` e.g. `mylib/v0.0.1`.

## Further Reading

For futher information about organising and managing multi-module repositories, please refer to the [oficial documentation](https://github.com/golang/go/wiki/Modules#faqs--multi-module-repositories).

[1]: ./myapp/go.mod
[2]: ./mylib/go.mod