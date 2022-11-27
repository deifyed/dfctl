# dfctl

## Motivation

Need to handle dotfiles. Why not write some code

## Usage

```shell
# Relink all tracked paths
dfctl relink

# Track a file or folder.
dfctl track <path>

# Untrack a file or folder.
dfctl untrack <path>

# List dfctled paths.
dfctl list

# Taint a certain path. Taint prevents a path from being relinked when running `dfctl relink`.
dfctl taint <path>

# Untaint a certain path. Untaint allows a path to be relinked when running `dfctl relink`.
infect untaint <path>
```

## Install

To install, run `make install`. This will install the `dfctl` binary to `~/.local/bin`. Make sure this directory is in your
`$PATH`.

```shell
# Example
make build && make install
```

To change the installation directory, set the `PREFIX` environment variable.

```shell
# Example
make build && PREFIX=/usr/local make install
```
