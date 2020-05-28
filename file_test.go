// Copyright 2020 Marko Mikulicic
// SPDX-License-Identifier: BSD-2-Clause

package filetransformer

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"unicode"

	"golang.org/x/text/runes"
)

func TestTransform(t *testing.T) {
	tmp, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmp.Name())

	fmt.Fprintf(tmp, "abcd")
	tmp.Close()

	if err := Transform(runes.Map(unicode.ToUpper), tmp.Name()); err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadFile(tmp.Name())
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(b), "ABCD"; got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
