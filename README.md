# cidrchk

A CLI tool that to assist with CIDR ranges and IPs. Uses [OPA](https://www.openpolicyagent.org/) Rego to check for CIDR ranges.

```
$ cidrchk contains 192.168.0.0/16 192.168.0.42
yes

$ cidrchk overlaps 192.168.0.0/16 192.168.1.0/24
yes

$ cidrchk expand 192.168.0.0/30
{“192.168.0.0", "192.168.0.1", "192.168.0.2", "192.168.0.3"}
```