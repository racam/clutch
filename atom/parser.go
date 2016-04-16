package atom

import (
	"encoding/xml"
)

// atomFeed is a simple struct to test if the xml seems to be an atom document
// that begin with an atom:feed element
type atomFeed struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/Atom feed"`
}

// atomEntry is a simple struct to test if the xml seems to be an atom document
// that begin with an atom:entry element
type atomEntry struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/Atom entry"`
}

const text string = "text"
const html string = "html"
const xhtml string = "xhtml"

// IsDeclared tries to find a atom:feed or atom:Entry element at the root of the
// XML document and looks at the namespace of this element
func IsDeclared(data []byte) bool {

	//Test atom:feed at first because it is the most common
	f := atomFeed{}
	err := xml.Unmarshal(data, &f)

	if err == nil {
		return true
	}

	e := atomEntry{}
	err = xml.Unmarshal(data, &e)

	return err == nil
}

// Parse parses the ATOM-encoded data into an atom.Feed struct and return it.
// The ATOM parsing do 2 steps :
// * Check that the document is well declared as an ATOM document
// * Parse the structure to fill additionnals fields
func Parse(data []byte) (*Feed, error) {
	// Test atom:feed at first because it is the most common
	f := Feed{}
	err := xml.Unmarshal(data, &f)

	// If the ATOM document does not start with atom:feed element, we test with
	// atom:entry element
	feedIsAbsent := err != nil
	if feedIsAbsent {
		e := Entry{}
		err = xml.Unmarshal(data, &e)

		if err != nil {
			return nil, err
		}

		f.Entry = append(f.Entry, e)
		f.IsDeclared = false

	} else {
		f.IsDeclared = true
	}

	//Parse the structure in order to fill additionnals fields
	f.parseContent()
	return &f, nil
}

func (f *Feed) parseContent() {
	for index := range f.Category {
		f.Category[index].parseContent()
	}

	for index := range f.Entry {
		f.Entry[index].parseContent()
	}

	for index := range f.Link {
		f.Link[index].parseContent()
	}

	if f.Rights != nil {
		f.Rights.parseContent()
	}
	if f.Subtitle != nil {
		f.Subtitle.parseContent()
	}
	if f.Title != nil {
		f.Title.parseContent()
	}
}

func (e *Entry) parseContent() {
	for index := range e.Category {
		e.Category[index].parseContent()
	}

	for index := range e.Link {
		e.Link[index].parseContent()
	}

	if e.Content != nil {
		e.Content.parseContent()
	}
	if e.Rights != nil {
		e.Rights.parseContent()
	}
	if e.Source != nil {
		e.Source.parseContent()
	}
	if e.Summary != nil {
		e.Summary.parseContent()
	}
	if e.Title != nil {
		e.Title.parseContent()
	}
}

func (c *Category) parseContent() {
	if c.XMLContent != "" && c.TextContent == "" {
		c.Content = c.XMLContent
	} else {
		c.Content = c.TextContent
	}
}

func (l *Link) parseContent() {

	// If the "rel" attribute is not present, the link element MUST be
	// interpreted as if the link relation type is "alternate".
	// source: https://tools.ietf.org/html/rfc4287#section-4.2.7.2
	if l.Rel == "" {
		l.Rel = "alternate"
	}

	if l.XMLContent != "" && l.TextContent == "" {
		l.Content = l.XMLContent
	} else {
		l.Content = l.TextContent
	}
}

func (c *Content) parseContent() {

	// If neither the "type" attribute nor the "src" attribute is provided, Atom
	// Processors MUST behave as though the "type" attribute were present with a
	// value of "text".
	// source: https://tools.ietf.org/html/rfc4287#section-4.1.3.1

	if c.Type == "" && c.Src == "" {
		c.Type = text
	}

	if c.Type == xhtml {
		c.Content = c.XMLContent
		return
	}

	if c.Type == text || c.Type == html {
		c.Content = c.TextContent
		return
	}

	if c.XMLContent != "" {
		c.Content = c.XMLContent
	} else {
		c.Content = c.TextContent
	}

}

func (s *Source) parseContent() {
	for index := range s.Category {
		s.Category[index].parseContent()
	}

	for index := range s.Link {
		s.Link[index].parseContent()
	}

	if s.Rights != nil {
		s.Rights.parseContent()
	}
	if s.Subtitle != nil {
		s.Subtitle.parseContent()
	}
	if s.Title != nil {
		s.Title.parseContent()
	}
}

func (t *Text) parseContent() {

	// If the "type" attribute is not provided, Atom Processors MUST behave as
	// though it were present with a value of "text"
	// Source : https://tools.ietf.org/html/rfc4287#section-3.1.1
	if t.Type == "" {
		t.Type = text
	}

	if t.Type == xhtml {
		t.Content = t.XMLContent
	} else {
		t.Content = t.TextContent
	}
}
