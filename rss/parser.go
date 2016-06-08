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

import "encoding/xml"

// rssDeclaration is a simple struct to test if the xml seems to be an rss
// document
type rssDeclaration struct {
	Version string   `xml:"version,attr"`
	XMLName xml.Name `xml:"rss"`
}

// IsDeclared tries to find a RSS element at the root of the xml document and
// a version attribute equal to 2.0
// source : https://cyber.law.harvard.edu/rss/rss.html#whatIsRss
func IsDeclared(data []byte) bool {
	r := rssDeclaration{}
	err := xml.Unmarshal(data, &r)

	if err != nil {
		return false
	}

	return r.Version == "2.0"
}

// Parse parses the RSS-encoded data into an RSS struct and return it
func Parse(data []byte) (*RSS, error) {
	r := RSS{}

	err := xml.Unmarshal(data, &r)

	return &r, err
}
