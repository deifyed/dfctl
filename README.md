# Infect

## Motivation

Need to handle dotfiles. Why not write some code

## Usage

```shell
# Infect filesystem with dotfiles.
# This will link all tracked files and folders to their target locations.
infect

# Track a file or folder.
# This will return the tracked file or folder to this location upon infection.
infect track <path>

# Untrack a file or folder.
# This will unlink the target and return the source file or folder to this location.
infect untrack <path>

# List infected paths.
# This will list all tracked files and folders.
infect list

# Taint a certain path.
# Mark certain paths as to not spread them when infecting filesystem.
infect taint <path>
```
