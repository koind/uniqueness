package main

import (
	"testing"
	"bufio"
	"strings"
	"bytes"
)

var testOk = `1
2
3
3
3
3
4
4
5
5`

var testOkResult = `1
2
3
4
5
`

func TestOk(t *testing.T) {

	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := uniq(in, out)

	if err != nil {
		t.Errorf("test for Ok Failed - error")
	}

	result := out.String()
	if result != testOkResult {
		t.Errorf("test for Ok Failed - results not match\n %v %v", result, testOkResult)
	}

}

var testFail = `3
1
4`

func TestForError(t *testing.T) {

	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)
	err := uniq(in, out)

	if err == nil {
		t.Errorf("test for Ok Failed - error")
	}

}