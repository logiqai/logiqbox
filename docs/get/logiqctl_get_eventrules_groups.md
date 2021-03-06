# `logiqctl get eventrules groups`

List all event rules

## Synopsis

List all event rules

```
logiqctl get eventrules groups [flags]
```

## Examples

```
logiqctl get eventrules all
```

## Options

```
  -h, --help   help for groups
```

## Options inherited from parent commands

```
  -c, --cluster string       Override the default cluster set by `logiqctl set-cluster' command
  -n, --namespace string     Override the default context set by `logiqctl set-context' command
  -o, --output string        Output format. One of: table|json|yaml. 
                             JSON output is not indented, use '| jq' for advanced JSON operations (default "table")
  -t, --time-format string   Time formatting options. One of: relative|epoch|RFC3339. 
                             This is only applicable when the output format is table. JSON and YAML outputs will have time in epoch seconds. (default "relative")
```

## SEE ALSO

* [logiqctl get eventrules](/get/logiqctl_get_eventrules)	 - List event rules

