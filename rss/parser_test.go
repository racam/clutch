// This file is part of clutch.
//
// clutch is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// clutch is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.

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
