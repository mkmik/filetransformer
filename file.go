// Copyright 2020 Marko Mikulicic
// SPDX-License-Identifier: BSD-2-Clause

/*
Package filetransformer offers a function to atomically replace the content of a file
by passing the previous content through a golang.org/x/text/transform.Transformer .
*/
package filetransformer

import (
	"io"
	"os"

	"github.com/google/renameio"
	"golang.org/x/text/transform"
)

// Transform applies a transformer on the contents of a file and writes
// the output back into the same file atomically.
// It calls Reset on t.
func Transform(t transform.Transformer, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w, err := renameio.TempFile("", filename)
	if err != nil {
		return err
	}
	defer w.Cleanup()

	defer t.Reset()
	if _, err := io.Copy(w, transform.NewReader(f, t)); err != nil {
		return err
	}
	return w.CloseAtomicallyReplace()
}
