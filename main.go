// Copyright 2016 The OPA Authors.  All rights reserved.
// Use of this source code is governed by an Apache2
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/cmd"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"

	"github.com/open-policy-agent/opa/util"
)

func main() {
	// refer to https://www.openpolicyagent.org/docs/latest/extensions/#adding-built-in-functions-to-the-opa-runtime
	rego.RegisterBuiltin1(
		&rego.Function{
			Name:    "shell.execute",
			Decl:    types.NewFunction(types.Args(types.S), types.S),
			Memoize: true,
		},
		func(bctx rego.BuiltinContext, a *ast.Term) (*ast.Term, error) {

			var shellCommand string

			if err := ast.As(a.Value, &shellCommand); err != nil {
				return nil, err
			}
			stdout, err := util.ExecuteShell(shellCommand)
			if err != nil {
				// need to show the response if shell command has error return
				// tbd - include the err in resopnse without affecting show stderr
				return ast.StringTerm(string(stdout)), nil
			}

			return ast.StringTerm(string(stdout)), nil
		},
	)

	if err := cmd.RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Capabilities + built-in metadata file generation:
//go:generate build/gen-run-go.sh internal/cmd/genopacapabilities/main.go capabilities.json
//go:generate build/gen-run-go.sh internal/cmd/genbuiltinmetadata/main.go builtin_metadata.json

// WASM base binary generation:
//go:generate build/gen-run-go.sh internal/cmd/genopawasm/main.go -o internal/compiler/wasm/opa/opa.go internal/compiler/wasm/opa/opa.wasm  internal/compiler/wasm/opa/callgraph.csv
