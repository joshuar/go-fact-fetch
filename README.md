# go-fact-fetch

[![Go Report Card](https://goreportcard.com/badge/github.com/joshuar/go-fact-fetch?style=flat-square)](https://goreportcard.com/report/github.com/joshuar/go-fact-fetch)
[![Release](https://img.shields.io/github/release/joshuar/go-fact-fetch.svg?style=flat-square)](https://github.com/joshuar/go-fact-fetch/releases/latest)

`go-fact-fetch` allows you to add [custom facts](https://docs.ansible.com/ansible/latest/user_guide/playbooks_vars_facts.html)
for use with [Ansible](https://www.ansible.com/).

## Installation

Grab one of the packages appropriate for your system on the [releases
page](https://github.com/joshuar/go-fact-fetch/releases).

## Usage

### In Ansible

After installing `go-fact-fetch`, custom facts it produces should be available/visible with:

```ansible
- name: grab new facts
  ansible.builtin.setup:
```

If using Fedora, the following would download the latest version of `go-fact-fetch` and retrieve the facts:

```ansible
- name: install go-fact-fetch
  ansible.builtin.dnf:
    name: "https://github.com/joshuar/go-fact-fetch/releases/download/v{{ go_fact_fetch_version }}/go-fact-fetch_{{ go_fact_fetch_version }}_linux_x86_64.rpm"
    state: installed
    disable_gpg_check: true
  become: true
  when: ansible_distribution == "Fedora"

- name: grab new facts
  ansible.builtin.setup:
```

Set the `go_fact_fetch_version` to the latest available version.

### On the command-line

`go-fact-fetch` is not really a command-line tool. If you have installed via one
of the packages (rpm/deb), it should already be set up for use with Ansible and
the facts available through the `ansible_local` hierarchy. If you want to see
what facts are available then run:

```shell
go-fact-fetch
```

To see the available commands.  Then you can run:

```shell
go-fact-fetch <command>
```

To see the facts it will produce (in JSON format).  You can use
[`jq`](https://stedolan.github.io/jq/) to pretty-print the output, for example:

```shell
go-fact-fetch nmcli | jq
```

## How it works

`go-fact-fetch` simplifies using executable `.fact` files in `/etc/ansible/fact.d`.
Rather than individual executable or shell scripts, each `.fact` file ends being
something like the following:

```shell
#!/bin/sh

go-fact-fetch <some-command>
```

`<some-command>` in go-fact-fetch is created using the great
[Cobra](https://github.com/spf13/cobra) library to make self-sufficient that can
then call any executable, parse the output and generate appropriate JSON that
Ansible then absorbs into its custom fact handling.

Using Cobra, it should be easy to extend go-fact-fetch to support any particular
command or fact output you require.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to
discuss what you would like to change. 

Some documentation for adding new commands and parsing command-line output and
generating JSON can be found in [CONTRIBUTING.md](docs/CONTRIBUTING.md).

## License
[MIT](https://choosealicense.com/licenses/mit/)