# cidrchk

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

`cidrchk` can do three things for you: check CIDR ranges for inclusion and overlap,
and generate all IPs in a CIDR range.

To check if a CIDR range contains an IP or another CIDR range:

```
$ cidrchk contains 192.168.0.0/16 192.168.0.42
yes
```

To Check if two CIDR ranges overlap:

```
$ cidrchk overlaps 192.168.0.0/16 192.168.1.0/24
yes
```

To expand a CIDR range, that is, to generate all IPs in it:

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