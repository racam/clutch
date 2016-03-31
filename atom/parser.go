package atom

import (
	"encoding/xml"
)

// atomFeed is a simple struct to test if the xml seems to be an atom document
// that begin with a feed element
type atomFeed struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/Atom feed"`
}

// atomEntry is a simple struct to test if the xml seems to be an atom document
// that begin with an entry element
type atomEntry struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/Atom entry"`
}

const text string = "text"
const html string = "html"
const xhtml string = "xhtml"

// IsDeclared tries to find a Feed/Entry element at the root of the xml document and
// looks at the namespace of this Element
func IsDeclared(data []byte) bool {

	//Test Feed at first because it is the most common
	f := atomFeed{}
	err := xml.Unmarshal(data, &f)

	if err == nil {
		return true
	}

	e := atomEntry{}
	err = xml.Unmarshal(data, &e)

	return err == nil
}

// Parse parses the ATOM-encoded data into an Feed struct and return it.
// The ATOM parsing do 3 steps :
// * Check is the document is well declared as an Atom document
// * Parse the structure to fill additionnals fields
// * Check that all requirements are respected
func Parse(data []byte) (*Feed, error) {
	// Test Feed at first because it is the most common
	f := Feed{}

	err := xml.Unmarshal(data, &f)

	// If the Atom document does not start with Feed Element, with test with
	// Entry Element
	feedIsAbsent := err != nil
	if feedIsAbsent {
		e := Entry{}

		err = xml.Unmarshal(data, &e)
		if err != nil {
			return nil, err
		}

		f.Entry = append(f.Entry, e)
	}

	f.parseContent()
	err = f.check(feedIsAbsent) // look at the check.go file
	return &f, err
}

func (f *Feed) parseContent() {
	for _, c := range f.Category {
		c.parseContent()
	}

	for index := range f.Entry {
		f.Entry[index].parseContent()
	}

	for index := range f.Link {
		f.Link[index].parseContent()
	}

	f.Rights.parseContent()
	f.Subtitle.parseContent()
	f.Title.parseContent()
}

func (e *Entry) parseContent() {
	for index := range e.Category {
		e.Category[index].parseContent()
	}

	for index := range e.Link {
		e.Link[index].parseContent()
	}

	e.Content.parseContent()
	e.Rights.parseContent()
	e.Source.parseContent()
	e.Summary.parseContent()
	e.Title.parseContent()
}

func (c *Category) parseContent() {
	if c.XMLContent != "" {
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

	if l.XMLContent != "" {
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

	if c.Type == text || c.Type == xhtml {
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

	s.Rights.parseContent()
	s.Subtitle.parseContent()
	s.Title.parseContent()
}

func (t *Text) parseContent() {

	// If the "type" attribute is not provided, Atom Processors MUST behave as
	// though it were present with a value of "text"
	if t.Type == "" {
		t.Type = text
	}

	if t.Type == xhtml {
		t.Content = t.XMLContent
	} else {
		t.Content = t.TextContent
	}
}
