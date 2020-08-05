## logiqctl license set

Configure license for LOGIQ deployment

### Synopsis

Configure license for LOGIQ deployment

```
logiqctl license set [flags]
```

### Examples

```
logiqctl license set -l=<license-file-path>
```

### Options

```
  -h, --help             help for set
  -l, --license string   license file path
```

### Options inherited from parent commands

```
  -c, --cluster string       Override the default cluster set by `logiqctl set-cluster' command
  -n, --namespace string     Override the default context set by `logiqctl set-context' command
  -o, --output string        Output format. One of: table|json|yaml. (default "table")
  -t, --time-format string   Time formatting options. One of: relative|epoch|RFC3339. 
                             This is only applicable when the output format is table. json and yaml outputs will have time in epoch seconds.
                             json output is not indented, use '| jq' for advanced json operations (default "relative")
```

### SEE ALSO

* [logiqctl license](logiqctl_license.md)	 - set and get license
