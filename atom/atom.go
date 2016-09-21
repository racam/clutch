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

//Refer to the RFC 4287
//link : https://tools.ietf.org/html/rfc4287

package atom

import "encoding/xml"

// CommonAttributes is a atom structure like describe in
// https://tools.ietf.org/html/rfc4287#section-2
type CommonAttributes struct {
	Base string `xml:"base,attr"`
	Lang string `xml:"lang,attr"`
}

// Feed is a atom structure like describe in
// https://tools.ietf.org/html/rfc4287#section-4.1.1
type Feed struct {
	CommonAttributes
	Author      []Person   `xml:"author"`
	Category    []Category `xml:"category"`
	Contributor []Person   `xml:"contributor"`
	Entry       []Entry    `xml:"entry"`
	Generator   Generator  `xml:"generator"`
	Icon        Icon       `xml:"icon"`
	IsDeclared  bool       `xml:"-"` //internal field for the checking step
	ID          ID         `xml:"id"`
	Link        []Link     `xml:"link"`
	Logo        Logo       `xml:"logo"`
	Rights      Text       `xml:"rights"`
	Subtitle    Text       `xml:"subtitle"`
	Title       Text       `xml:"title"`
	Updated     Date       `xml:"updated"`
	XMLName     xml.Name   `xml:"feed"`
}

// Entry is a atom structure like describe in
// https://tools.ietf.org/html/rfc4287#section-4.1.2
type Entry struct {
	CommonAttributes
	Author      []Person   `xml:"author"`
	Category    []Category `xml:"category"`
	Content     Content    `xml:"content"`
	Contributor []Person   `xml:"contributor"`
	ID          ID         `xml:"id"`
	Link        []Link     `xml:"link"`
	Published   Date       `xml:"published"`
	Rights      Text       `xml:"rights"`
	Source      Source     `xml:"source"`
	Summary     Text       `xml:"summary"`
	Title       Text       `xml:"title"`
	Updated     Date       `xml:"updated"`
}

// Content is a atom structure like describe in
// https://tools.ietf.org/html/rfc4287#section-4.1.3
type Content struct {
	CommonAttributes
	Text
	Src  string `xml:"src,attr"`
	Type string `xml:"type,attr"`
}

// Link is a atom structure like describe in
// https://tools.ietf.org/html/rfc4287#section-4.2.7
type Link struct {
	CommonAttributes
	Content     string `xml:"-"` //Fill with parseContent functions
	Href        string `xml:"href,attr"`
	Hreflang    string `xml:"hreflang,attr"`
	Length      string `xml:"length,attr"`
	Rel         string `xml:"rel,attr"`
	TextContent string `xml:",chardata"`
	Title       string `xml:"title,attr"`
	Type        string `xml:"type,attr"`
	XMLContent  string `xml:",innerxml"`
}

// Source is a atom structure like describe in
// https://tools.ietf.org/html/rfc4287#section-4.2.11
type Source struct {
	CommonAttributes
	Author      []Person   `xml:"author"`
	Category    []Category `xml:"category"`
	Contributor []Person   `xml:"contributor"`
	Generator   Generator  `xml:"generator"`
	Icon        Icon       `xml:"icon"`
	ID          ID         `xml:"id"`
	Link        []Link     `xml:"link"`
	Logo        Logo       `xml:"logo"`
	Rights      Text       `xml:"rights"`
	Subtitle    Text       `xml:"subtitle"`
	Title       Text       `xml:"title"`
	Updated     Date       `xml:"updated"`
}

// Text is a atom structure like describe in
// https://tools.ietf.org/html/rfc4287#section-3.1
type Text struct {
	CommonAttributes
	Content     string `xml:"-"` //Fill with parseContent functions
	TextContent string `xml:",chardata"`
	Type        string `xml:"type,attr"`
	XMLContent  string `xml:",innerxml"`
	AnyContent  string `xml:",any"`
}

// Person is a atom structure like describe in
// https://tools.ietf.org/html/rfc4287#section-3.2
type Person struct {
	CommonAttributes
	Email string `xml:"email"`
	Name  string `xml:"name"`
	URI   string `xml:"uri"`
}

// Date is a atom structure like describe in
// https://tools.ietf.org/html/rfc4287#section-3.3
type Date struct {
	CommonAttributes
	DateTime string `xml:",chardata"`
}

// Category is a atom structure like describe in
// https://tools.ietf.org/html/rfc4287#section-4.2.2
type Category struct {
	CommonAttributes
	Content     string `xml:"-"` //Fill with parseContent functions
	Label       string `xml:"label,attr"`
	Scheme      string `xml:"scheme,attr"`
	Term        string `xml:"term,attr"`
	TextContent string `xml:",chardata"`
	XMLContent  string `xml:",innerxml"`
}

// Generator is a atom structure like describe in
// https://tools.ietf.org/html/rfc4287#section-4.2.4
type Generator struct {
	CommonAttributes
	Content string `xml:",chardata"`
	Version string `xml:"version,attr"`
	URI     string `xml:"uri,attr"`
}

// CommonURI is a atom structure
type CommonURI struct {
	CommonAttributes
	URI string `xml:",chardata"`
}

// Icon : https://tools.ietf.org/html/rfc4287#section-4.2.5
type Icon CommonURI

// ID : https://tools.ietf.org/html/rfc4287#section-4.2.6
type ID CommonURI

// Logo : https://tools.ietf.org/html/rfc4287#section-4.2.8
type Logo CommonURI
