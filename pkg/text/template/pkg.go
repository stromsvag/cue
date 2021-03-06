// Code generated by go generate. DO NOT EDIT.

//go:generate rm pkg.go
//go:generate go run ../../gen/gen.go

package template

import (
	"cuelang.org/go/internal/core/adt"
	"cuelang.org/go/pkg/internal"
)

func init() {
	internal.Register("text/template", pkg)
}

var _ = adt.TopKind // in case the adt package isn't used

var pkg = &internal.Package{
	Native: []*internal.Builtin{{
		Name:   "Execute",
		Params: []adt.Kind{adt.StringKind, adt.TopKind},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			templ, data := c.String(0), c.Value(1)
			if c.Do() {
				c.Ret, c.Err = Execute(templ, data)
			}
		},
	}, {
		Name:   "HTMLEscape",
		Params: []adt.Kind{adt.StringKind},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			s := c.String(0)
			if c.Do() {
				c.Ret = HTMLEscape(s)
			}
		},
	}, {
		Name:   "JSEscape",
		Params: []adt.Kind{adt.StringKind},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			s := c.String(0)
			if c.Do() {
				c.Ret = JSEscape(s)
			}
		},
	}},
}
