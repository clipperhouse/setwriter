package setwriter

import (
	"io"

	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(NewSetWriter())
	if err != nil {
		panic(err)
	}
}

type SetWriter struct{}

func NewSetWriter() *SetWriter {
	return &SetWriter{}
}

func (c *SetWriter) Name() string {
	return "set"
}

func (c *SetWriter) Imports(t typewriter.Type) (result []typewriter.ImportSpec) {
	// none
	return result
}

func (c *SetWriter) Write(w io.Writer, t typewriter.Type) error {
	tag, found := t.FindTag(c)

	if !found {
		// nothing to be done
		return nil
	}

	license := `// Set is a modification of https://github.com/deckarep/golang-set
// The MIT License (MIT)
// Copyright (c) 2013 Ralph Caraveo (deckarep@gmail.com)
`

	if _, err := w.Write([]byte(license)); err != nil {
		return err
	}

	if err := set.TypeConstraint.TryType(t); err != nil {
		return err
	}

	tmpl, err := templates.ByTag(t, tag)

	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, t); err != nil {
		return err
	}

	return nil
}
