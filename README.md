# Infect

## Motivation

Need to handle dotfiles. Why not write some code

## Usage

```shell
# Track a file or folder
infect track <path>

# Infect filesystem with dotfiles
#(in dotfiles repository)
infect

# List infected paths
infect list

# Mark certain paths as to not spread them when infecting filesystem
infect taint <path>
```
