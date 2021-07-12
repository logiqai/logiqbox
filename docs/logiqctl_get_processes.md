## logiqctl get processes

List all available processes. This command runs an interactive prompt that lets you choose an application from a list of available applications.

### Synopsis

List all available processes. This command runs an interactive prompt that lets you choose an application from a list of available applications.

```
logiqctl get processes [flags]
```

### Examples

```
logiqctl get processes|proc|p
```

### Options

```
  -h, --help   help for processes
```

### Options inherited from parent commands

```
  -c, --cluster string       Override the default cluster set by `logiqctl set-cluster' command
  -n, --namespace string     Override the default context set by `logiqctl set-context' command
  -o, --output string        Output format. One of: table|json|yaml. 
                             JSON output is not indented, use '| jq' for advanced JSON operations (default "table")
  -t, --time-format string   Time formatting options. One of: relative|epoch|RFC3339. 
                             This is only applicable when the output format is table. JSON and YAML outputs will have time in epoch seconds. (default "relative")
```

### SEE ALSO

* [logiqctl get](logiqctl_get.md)	 - Display one or more of your LOGIQ resources
