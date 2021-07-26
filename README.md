# go-gadget-go

[![Go Report Card](https://goreportcard.com/badge/github.com/joshuar/go-gadget-go?style=flat-square)](https://goreportcard.com/report/github.com/joshuar/go-gadget-go)
[![Release](https://img.shields.io/github/release/joshuar/go-gadget-go.svg?style=flat-square)](https://github.com/joshuar/go-gadget-go/releases/latest)

go-gadget-go allows you to add [custom facts](https://docs.ansible.com/ansible/latest/user_guide/playbooks_vars_facts.html)
for use with [Ansible](https://www.ansible.com/).

## Installation

Grab one of the packages appropriate for your system on the [releases
page](https://github.com/joshuar/go-gadget-go/releases).

## Usage

`go-gadget-go` is not really a command-line tool. If you have installed via one
of the packages (rpm/deb), it should already be set up for use with Ansible and
the facts available through the `ansible_local` hierarchy. if you want to see
what facts are available then run:

```shell
go-gadget-go
```

To see the available commands.  Then you can run:

```shell
go-gadget-go <command>
```

To see the facts it will produce (in JSON format).  You can use
[`jq`](https://stedolan.github.io/jq/) to pretty-print the output, for example:

```shell
go-gadget-go nmcli | jq
```

## How it works

go-gadget-go simplifies using executable `.fact` files in `/etc/ansible/fact.d`.
Rather than individual executable or shell scripts, each `.fact` file ends being
something like the following:

```shell
#!/bin/sh

go-gadget-go <some-command>
```

`<some-command>` in go-gadget-go is created using the great
[Cobra](https://github.com/spf13/cobra) library to make self-sufficient that can
then call any executable, parse the output and generate appropriate JSON that
Ansible then absorbs into it's custom fact handling.

Using Cobra, it should be easy to extend go-gadget-go to support any particular
command or fact output you require.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to
discuss what you would like to change. 

Some documentation for adding new commands and parsing command-line output and
generating JSON can be found in [](docs/CONTRIBUTING.md).

## License
[MIT](https://choosealicense.com/licenses/mit/)