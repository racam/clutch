package rss

import (
	"io/ioutil"
	"testing"
)

func TestIsDeclared(t *testing.T) {
	var prefix = "../testdata/rss/"

	var namespaces = []struct {
		filename string // input file
		expected bool   // expected result
	}{
		{"unit_01_IsDeclared.xml", true},
		{"unit_02_IsDeclared.xml", false},
		{"unit_03_IsDeclared.xml", false},
	}

	for _, ns := range namespaces {

		f, err := ioutil.ReadFile(prefix + ns.filename)

		if err != nil {
			t.Errorf("[RSS][Unit][IsDeclared] file '%s' : is missing",
				prefix+ns.filename)
		} else {
			res := IsDeclared(f)

			//Test if isDeclared return an expected result
			if res != ns.expected {
				t.Errorf(`[RSS][Unit] file '%s' : expected result '%t',
        actual %t`, ns.filename, ns.expected, res)
			}
		}
	}
}
