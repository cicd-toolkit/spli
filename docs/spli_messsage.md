## spli messsage

messsage

### Synopsis


this command is passing the system messages and print over cli
example:

spli msg
```
[
  {
    "NoAllowedDomainsList": "Security risk warning: Found an empty value for allowedDomainList in the alert_actions.conf configuration file. If you do not configure this setting, then users can send email alerts with search results to any domain. You can add values for allowedDomainList either in the alert_actions.conf file or in Server Settings > Email Settings > Email Domains in Splunk Web.",
    "capabilities": ["admin_all_objects"],
    "eai:acl": null,
    "help": "",
    "message": "Security risk warning: Found an empty value for allowedDomainList in the alert_actions.conf configuration file. If you do not configure this setting, then users can send email alerts with search results to any domain. You can add values for allowedDomainList either in the alert_actions.conf file or in Server Settings > Email Settings > Email Domains in Splunk Web.",
    "message_alternate": "",
    "server": "36c2770ea6dc",
    "severity": "warn",
    "timeCreated_epochSecs": 1730382326,
    "timeCreated_iso": "2024-10-31T13:45:26+00:00"
  }
]
```

use `--jsonpath` to select any output
```
spli msg --jsonpath "entry.#.content.severity|@pretty"
["warn"]
```


```
spli messsage [flags]
```

### Options

```
  -h, --help              help for messsage
      --jsonpath string   jsonpath using gjson (default "entry.#.content|@pretty")
```

### SEE ALSO

* [spli](spli.md)	 - splunk cli

