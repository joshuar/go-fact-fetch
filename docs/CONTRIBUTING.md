## Adding new commands

You can add commands using the [Cobra
generator](https://github.com/spf13/cobra/blob/master/cobra/README.md).  To add
a new command (after installing the Cobra generator):

```shell
cobra add <some_command>
```

When adding commands:
- The command name should closely follow the name of whatever command-line tool
  it is generating output from.  For example, for the `nmcli` command from
  NetworkManager it is just `go-fact-fetch nmcli` and the output of 
  `nmcli connection` and `nmcli device` is produced from this command. For `tuned-adm`,
  it is just `go-fact-fetch tuned` for simplicity.  
- You can have sub-commands to *filter* the output.  For example, 
  `go-fact-fetch nmcli connection` shows just output from the 
  `nmcli connection show` command.

## Development tips

### Running command-line tools

[os/exec](https://pkg.go.dev/os/exec) is the best option.  In particular,
`exec.Command`.  Output will be a `[]byte`.  

### Parsing command output

[bytes](https://pkg.go.dev/bytes) has functions to handle most things you will
need. For example, the `bytes.TrimSpace` will get rid of extra white-space and
`bytes.FieldsFunc` will allow you to split up a line by some field.

### Generating JSON

[encoding/json](https://pkg.go.dev/encoding/json) has `json.Marshal` which will
be what you need.  




