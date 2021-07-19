# go-gadget-go

go-gadget-go allows you to add [custom facts](https://docs.ansible.com/ansible/latest/user_guide/playbooks_vars_facts.html)
for use with [Ansible](https://www.ansible.com/).

## Installation

Grab one of the packages appropriate for your system on the [releases page](https://github.com/joshuar/go-gadget-go/releases).

## Usage

```shell
go-gadget-go <command>
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
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)