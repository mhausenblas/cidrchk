package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/open-policy-agent/opa/rego"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
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
		contains(os.Args[2], os.Args[3])
	case "overlaps":
		fmt.Println("overlaps not yet implemented")
	case "expand":
		fmt.Println("expand not yet implemented")
	default:
		fmt.Println("unknown command")
		flag.Usage()
	}
}

func contains(cidrange, iporcidrange string) {
	fmt.Printf("Checking if %v is in %v:\n", iporcidrange, cidrange)
	query, err := rego.New(
		rego.Query("x = data.example.authz.allow"),
		rego.Module("example.rego", module),
	).PrepareForEval(ctx)

	if err != nil {
		// Handle error.
	}
	input := map[string]interface{}{
		"method": "GET",
		"path":   []interface{}{"salary", "bob"},
		"subject": map[string]interface{}{
			"user":   "bob",
			"groups": []interface{}{"sales", "marketing"},
		},
	}

	results, err := query.Eval(context.Context, rego.EvalInput(input))
	if err != nil {
		// Handle evaluation error.
	} else if len(results) == 0 {
		// Handle undefined result.
	} else if result, ok := results[0].Bindings["x"].(bool); !ok {
		// Handle unexpected result type.
	} else {
		// Handle result/decision.
	}
}
