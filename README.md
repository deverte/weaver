# Weaver
The Weaver program is a package manager for symlink programs (variation of
portable applications).

---
- [Information](#information)
- [Weaver - Scoop similarity](#weaver-scoop-similarity)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [settings.yaml](#settingsyaml)
- [Environment variables](#environment-variables)
- [app.yaml](#appyaml)
---

## Information
When some portable applications creators are modifying these apps or make them
run inside virtual environments, symlink way is something in the middle.
To make symlink app, you can trace for file system and registry (in Windows)
changes while installation (sometimes after some usage) and import/export these
changes when you want without modifying this app and without license violation.
This method also gives you an ability to install/uninstall any symlink app
within fractions of a second.

## Weaver - Scoop similarity
Weaver is inspired by [Scoop](https://scoop.sh/) package manager, but pursues a
different goal.
When some applications does not have portable versions it's hard sometimes to
effectively manipulate them (e.g. if you want to keep your system clean).
Scoop provides *main* repo for portable applications, but not all *extra*
applications are portable (and this is not Scoop's fault).

Note that ***Weaver*** is not fork of the ***Scoop***, but adopts some
conceptions from this package manager.
There are some similarities with ***Scoop***'s folder structure:
```
weaver                      ← scoop
weaver/apps                 ← scoop/apps
weaver/cache                ← scoop/cache
weaver/shims                ← scoop/shims
weaver/tangle               ← scoop/buckets
weaver/tangle/$repo         ← scoop/buckets/$repo
weaver/tangle/$repo/fiber   ← scoop/buckets/$repo/bucket
```

## Installation
**Note: in future this installation method will be replaced with single
scipt.**

1. Add `WEAVER_HOME` environment variable where you prefer to store symlink
apps.
2. Download last version of the Weaver executable from
[Releases](https://github.com/deverte/weaver/releases) page.
3. Create Weaver folder structure inside `$WEAVER_HOME` as mentioned in
[Weaver - Scoop similarity](#weaver-scoop-similarity) and put `weaver.exe` into
`$WEAVER_HOME/apps/weaver` folder.
4. Add `weaver.exe` startup script into `$WEAVER_HOME/shims`. E.g. `weaver.ps1`
that contains `& "$env:WEAVER_HOME\apps\weaver\weaver.exe" $args` line.
5. Add `$WEAVER_HOME/shims` into `$Path` environment variable.

## Usage
Let us consider work with `hello_world` application.
All examples below are applicable to any application controlled by `weaver`.

**FUTURE FUNCTIONALITY!** Download and append `hello_world` application from
`main` repository.
```sh
weaver add hello_world
```

Install `hello_world` application.
```sh
weaver install hello_world
```

**FUTURE FUNCTIONALITY!** Or you can use single command to download and install
`hello_world` application.
```sh
weaver add hello_world -i
```

To list all installed applications and it's statuses, type:
```sh
weaver list
```

You can test `hello_world` app by running:
```sh
hello_world
```

Output should be:
```
Hello, World!
```

You can uninstall `hello_world` application with `uninstall` command.
But this command will only delete symlinks, registry keys, temporary files and
execute uninstall scripts.
```sh
weaver uninstall hello_world
```

Now, if you call `list` command, you will see that this program is present
inside list but with `uninstalled` status.

**FUTURE FUNCTIONALITY!** If you want to completely delete `hello_world`
application with user data and it's files, execute:
```sh
weaver remove hello_world
```

## Commands
```
add [application name]              Search application in added repositories and
                                    download it if present.
    -i, --install                   Also executes `install` command.
    -s, --source [source folder]    You can specify custom application source
                                    folder that does not presented in
                                    repositories. If not presented, application
                                    will be searched inside `weaver/apps`
                                    folder. Conflicts with `--target` flag.
    -t, --target [target folder]    You can specify custom application folder if
                                    you don't want to use `weaver/apps` folder.
                                    Conflicts with `--source` flag.
info [application name]             Information about specified application.
install [application name]          Install application.
list                                List of all added applications with
                                    installation status.
remove [application name]           Completely remove application, it's files
                                    and user files connected with this
                                    application.
search [application name]           Search application in added repositories.
tangle
        add [tangle path]           Add tangle from git repository.
        list                        List of all added tangles.
        remove [tangle name]        Remove tangle.
uninstall [application name]        Uninstall application, with ability to
                                    restore it with all user data.
version                             Print the version number of the Weaver.

Flags:
    -h, --help                      Outputs documentation about `weaver` or it's
                                    command.
```

Some commands don't work yet: `add`, `remove`, `search` and will be implemented
in future versions.

## settings.yaml
**FUTURE FUNCTIONALITY**
```yaml
external:
  # Applications added with `--target` option
  targets:
    - name: app_name_1
      path: path/to/target/1
    - name: app_name_2
      path: path/to/target/2
  # Applications added with `--source` option
  sources:
    - name: app_name_3
      path: path/to/source/3
    - name: app_name_4
      path: path/to/source/4
```

## Environment variables
`WEAVER_HOME` — Weaver home directory.

## app.yaml
```yaml
name: Application name
version: 1.0.0
install:
  symlinks:
    - source: path/to/source/1
      target: path/to/target/1
    - source: path/to/source/2
      target: path/to/target/2
  copy:
    - source: path/to/source/file/1.ext
      target: path/to/target/directory/1
    - source: path/to/source/file/2.ext
      target: path/to/target/directory/2
  reg:
    - path: path/to/install.reg
  scripts:
    - path: path/to/install.ps1
    - path: path/to/install.cmd
    - path: path/to/install.bat
uninstall:
  symlinks:
    - target: path/to/target/1
    - target: path/to/target/2
  delete:
    - path: path/to/target/directory/1/1.ext
    - path: path/to/target/directory/2/2.ext
  reg:
    - path: path/to/uninstall.reg
  scripts:
    - path: path/to/uninstall.ps1
    - path: path/to/uninstall.cmd
    - path: path/to/uninstall.bat
```