package main

import (
	"flag"
	"testing"

	"github.com/google/go-cmdtest"
)

var update = flag.Bool("update", false, "update test files with results")

func TestCLI(t *testing.T) {
	ts, err := cmdtest.Read("testdata")
	if err != nil {
		t.Fatal(err)
	}

	ts.Commands["cmd"] = cmdtest.InProcessProgram("cmd", Execute)
	ts.Run(t, *update)
}
