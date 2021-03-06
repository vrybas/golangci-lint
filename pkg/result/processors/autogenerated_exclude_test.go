package processors

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAutogeneratedDetection(t *testing.T) {
	all := `
	// generated by stringer -type Pill pill.go; DO NOT EDIT

// Code generated by "stringer -type Pill pill.go"; DO NOT EDIT

// Code generated by vfsgen; DO NOT EDIT

// Created by cgo -godefs - DO NOT EDIT

/* Created by cgo - DO NOT EDIT. */

// Generated by stringer -i a.out.go -o anames.go -p ppc64
// Do not edit.

// DO NOT EDIT
// generated by: x86map -fmt=decoder ../x86.csv

// DO NOT EDIT.
// Generate with: go run gen.go -full -output md5block.go

// generated by "go run gen.go". DO NOT EDIT.

// DO NOT EDIT. This file is generated by mksyntaxgo from the RE2 distribution.

// GENERATED BY make_perl_groups.pl; DO NOT EDIT.

// generated by mknacl.sh - do not edit

// DO NOT EDIT ** This file was generated with the bake tool ** DO NOT EDIT //

// Generated by running
//  maketables --tables=all --data=http://www.unicode.org/Public/8.0.0/ucd/UnicodeData.txt
// --casefolding=http://www.unicode.org/Public/8.0.0/ucd/CaseFolding.txt
// DO NOT EDIT

/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/gogen/unmarshalmap
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

// AUTOGENERATED FILE: easyjson file.go
`

	generatedCases := strings.Split(all, "\n\n")
	for _, gc := range generatedCases {
		isGenerated := isGeneratedFileByComment(gc)
		assert.True(t, isGenerated)
	}

	notGeneratedCases := []string{
		"code not generated by",
		"test",
	}
	for _, ngc := range notGeneratedCases {
		isGenerated := isGeneratedFileByComment(ngc)
		assert.False(t, isGenerated)
	}
}

func TestGetDoc(t *testing.T) {
	const expectedDoc = `first line
second line
third line`
	doc, err := getDoc(filepath.Join("testdata", "autogen_exclude.go"))
	assert.NoError(t, err)
	assert.Equal(t, expectedDoc, doc)
}
