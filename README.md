# Overview

This is an example of one way or organising a Github repository with two Go mods: one executable module and the other a library module.

In the case of executable module, this is given a name that is not externally reachable. This module is simply named `myapp`. This external mods can't do a `go install` to pull any artefacts.

In the case of the library module, this is given a name that can be externally accessed by an external Go module. This module is named `github.com/paulwizviz/go-workspace/mylib`.