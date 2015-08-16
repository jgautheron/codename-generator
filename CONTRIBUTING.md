# Contributing to `codename-generator`

This tool assumes you are working in a standard Go workspace, as described in http://golang.org/doc/code.html.

## Workflow

1. Fork the repository.
2. Create a dedicated branch for the feature/fix you'd like to add.
3. Keep in mind that commit messages should follow the [AngularJS Git Commit Message Conventions](https://docs.google.com/document/d/1QrDFcIiPjSLDn3EL15IJygNPiHORgU1_OOAqWjiDU5Y/edit).
4. Submit a PR! It'll get reviewed shortly.

## Dependency management

This project uses `godep` for managing dependencies. If you added some 3rd party imports, you will have to save the new dependencies:

    $ cd gocha
    $ go get -u github.com/tools/godep
    $ $GOPATH/bin/godep save