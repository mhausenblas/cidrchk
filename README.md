# cidrchk

![.github/workflows/release.yml](https://github.com/mhausenblas/cidrchk/workflows/.github/workflows/release.yml/badge.svg)

A CLI tool to assist you with CIDR ranges and IPs. 

## Install it

You can download the [latest binary](https://github.com/mhausenblas/cidrchk/releases/latest) 
for Linux (Intel and Arm), macOS, and Windows.

For example, to install `cidrchk` from binary on macOS, do the following:

```sh
curl -L https://github.com/mhausenblas/cidrchk/releases/latest/download/cidrchk_darwin_amd64.tar.gz \
    -o cidrchk.tar.gz && \
    tar xvzf cidrchk.tar.gz cidrchk && \
    mv cidrchk /usr/local/bin && \
    rm cidrchk*
```

## Use it

`cidrchk` can do three things for you: 1. check CIDR ranges for inclusion, 2. check for CIDR range overlaps, and 3. generate all IPs in a CIDR range.

### Check for CIDR range inclusion

To check if a CIDR range contains an IP or another CIDR range:

```
$ cidrchk contains 192.168.0.0/16 192.168.0.42
yes
```

This also works for IPv6 addresses, for example:

```
$ cidrcheck contains \
            0:0:0:0:0:ffff:c0a8:0/30 \
            0:3:ffff:ffff:ffff:ffff:ffff:ffff
yes
``` 

Together with `jq` one can also answer questions like "How many IP addresses
 are there in a CIDR range", for example:

```
$ cidrchk expand 192.168.0.0/16 | jq '.ips[] | length'
65536
```

### Check for CIDR range overlaps

To check if two CIDR ranges overlap you can do:

```
$ cidrchk overlaps 192.168.0.0/16 192.168.1.0/24
yes
```

### Generate IPs from CIDR range

To expand a CIDR range, that is, to generate all IPs in it do the following:

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
to perform CIDR operation, to be precise the 
[built-in Net functions](https://www.openpolicyagent.org/docs/latest/policy-reference/#net).