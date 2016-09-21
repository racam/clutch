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

package atom

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// Source: https://github.com/mmcdole/gofeed/blob/master/atom/parser_test.go
func TestParser(t *testing.T) {
	files, _ := filepath.Glob("../testdata/atom/integ_*.xml")
	for _, file := range files {

		base := filepath.Base(file)
		name := strings.TrimSuffix(base, filepath.Ext(base))

		// Get actual source feed
		filename := fmt.Sprintf("../testdata/atom/%s.xml", name)
		atomContent, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Errorf("[Atom][Integ] %s", err)
			continue
		}

		// Get encoded expected feed result
		filename = fmt.Sprintf("../testdata/atom/%s.json", name)
		jsonContent, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Errorf("[Atom][Integ] %s", err)
			continue
		}

		// Parse actual xml feed
		actual, err := Parse(atomContent)
		if err != nil {
			t.Errorf("[Atom][Integ] file %s.xml impossible to parse : %s", name, err)
			continue
		}

		// Parse expected json feed
		var expected Feed
		err = json.Unmarshal(jsonContent, &expected)
		if err != nil {
			t.Errorf("[Atom][Integ] file %s.json impossible to parse : %s", name, err)
			continue
		}

		// Becareful actual is a pointer
		if !reflect.DeepEqual(actual, &expected) {
			// debug
			t.Errorf("[Atom][Integ] file %s, xml and json don't match", name)
			//t.Logf("[DEBUG]%+v", actual)
			//t.Logf("[DEBUG]%+v", &expected)
		}
	}
}
