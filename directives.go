// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gocyclo

import (
	"go/ast"
	"slices"
	"strings"
)

type directives []string

func (ds directives) HasIgnore() bool {
	return ds.isPresent("ignore")
}

func (ds directives) isPresent(name string) bool {
	return slices.Contains(ds, name)
}

func parseDirectives(doc *ast.CommentGroup) directives {
	if doc == nil {
		return directives{}
	}
	ds := make(directives, 0, len(doc.List))
	for _, comment := range doc.List {
		if after, ok := strings.CutPrefix(comment.Text, "//gocyclo:"); ok {
			ds = append(ds, strings.TrimSpace(after))
		}
	}
	return ds
}
