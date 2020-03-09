# cidrchk

A CLI tool to assist you with CIDR ranges and IPs. 

`cidrchk` can do three things for you:

1. Check if a CIDR range contains an IP or another CIDR range:

```
$ cidrchk contains 192.168.0.0/16 192.168.0.42
yes
```

2. Check if two CIDR ranges overlap:

```
$ cidrchk overlaps 192.168.0.0/16 192.168.1.0/24
yes
```

3. Expand a CIDR range, that is, list all IPs in it:

```
$ cidrchk expand 192.168.0.0/30 | jq .
{
  "cidr": "192.168.0.0/30",
  "ips": [
    [
      "192.168.0.0",
      "192.168.0.1",
      "192.168.0.2",
      "192.168.0.3"
    ]
  ]
}
```

Note that `cidrchk`  uses [OPA](https://www.openpolicyagent.org/) Rego 
to perform CIDR operation, to be precise the [built-in Net functions](https://www.openpolicyagent.org/docs/latest/policy-reference/#net).