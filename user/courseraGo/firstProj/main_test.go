package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOk = `1
2
3
3
4
5`
var testOut = `1
2
3
4
5
`
var testFail = `1
2
1
`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := unique(in, out)

	if err != nil {
		t.Errorf("test for Ok Failed -error")
	}

	if out.String() != testOut {
		t.Errorf("test Ok Failed - results not match\n %v %v", out.String(), testOut)
	}
}

func TestFail(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)
	err := unique(in, out)

	if err == nil {
		t.Errorf("test for Ok Failed -error")
	}

}
