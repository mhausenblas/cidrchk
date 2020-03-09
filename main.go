package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/open-policy-agent/opa/rego"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	// the Rego rules for CIDR operations, see also the standalone example
	// in cidrchck.rego which you can test with the opa CLI like so:
	// opa eval --input input.json --data cidrchk.rego --package cidrchck 'contains'
	module = `package cidrchck

		default contains = "no"
		default overlaps = "no"

		contains = "yes" {
			cr := input.targetcidr
			ior := input.incidr
			net.cidr_contains(cr, ior)
		}

		contains = "yes" {
			cr := input.targetcidr
			ior := input.inip
			net.cidr_contains(cr, ior)
		}

		overlaps = "yes" {
			cr := input.targetcidr
			ior := input.incidr
			net.cidr_intersects(cr, ior)
		}

		expand[ips] {
			cr := input.incidr
			ips := net.cidr_expand(cr)
		}
	`
)

func main() {
	// Define the CLI API:
	flag.Usage = func() {
		fmt.Printf("Usage:\n %s \n", os.Args[0])
		fmt.Println("Example usage:\n cidrchk")
		fmt.Println("Arguments:")
		flag.PrintDefaults()
	}
	showversion := flag.Bool("version", false, "Print the version and exit")
	flag.Parse()

	if *showversion {
		fmt.Printf("%v, commit %v, built at %v\n", version, commit, date)
		os.Exit(0)
	}

	command := os.Args[1]

	switch command {
	case "contains":
		result, err := contains(os.Args[2], os.Args[3])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	case "overlaps":
		result, err := overlaps(os.Args[2], os.Args[3])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	case "expand":
		result, err := expand(os.Args[2])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	default:
		fmt.Println("unknown command")
		flag.Usage()
	}
}

// contains returns "yes" if the  input IP or CIDR is in the target CIDR
// and it returns "no" otherwise
func contains(targetcidr, iporcidr string) (string, error) {
	// log.Printf("Checking if %v is in %v:\n", iporcidr, targetcidr)
	reg := rego.New(
		rego.Query("data.cidrchck.contains"),
		rego.Module("example.rego", module),
		rego.Input(
			map[string]interface{}{
				"targetcidr": targetcidr,
				"inip":       iporcidr,
				"incidr":     iporcidr,
			},
		),
	)
	ctx := context.Background()
	rs, err := reg.Eval(ctx)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf("%v\n", rs[0].Expressions[0])
	return result, nil
}

// overlaps returns "yes" if the input CIDR overlaps with the target CIDR
// and it returns "no" otherwise
func overlaps(targetcidr, inputcidr string) (string, error) {
	// log.Printf("Checking if %v overlaps with %v:\n", inputcidr, targetcidr)
	reg := rego.New(
		rego.Query("data.cidrchck.overlaps"),
		rego.Module("example.rego", module),
		rego.Input(
			map[string]interface{}{
				"targetcidr": targetcidr,
				"incidr":     inputcidr,
			},
		),
	)
	ctx := context.Background()
	rs, err := reg.Eval(ctx)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf("%v\n", rs[0].Expressions[0])
	return result, nil
}

// expand generates a list of all IP addresses in the CIDR
func expand(cidr string) (string, error) {
	// log.Printf("Expanding CIDR %v:\n", cidr)
	reg := rego.New(
		rego.Query("data.cidrchck.expand"),
		rego.Module("example.rego", module),
		rego.Input(
			map[string]interface{}{
				"incidr": cidr,
			},
		),
	)
	ctx := context.Background()
	rs, err := reg.Eval(ctx)
	if err != nil {
		return "", err
	}
	type ipsofcidr struct {
		CIDR string      `json:"cidr"`
		IPs  interface{} `json:"ips"`
	}
	ips := ipsofcidr{
		CIDR: cidr,
		IPs:  rs[0].Expressions[0].Value,
	}
	val, err := json.Marshal(ips)
	if err != nil {
		return "", err
	}
	return string(val), nil
}
