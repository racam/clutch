package atom

import (
	"encoding/xml"
)

// A simple struct to test if the xml seems to be an atom document
type atomFeed struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/Atom feed"`
}

// A simple struct to test if the xml seems to be an atom document
type atomEntry struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/Atom entry"`
}

// IsAtom tries to find a Feed/Entry element at the root of the xml document and
// looks at the namespace of this Element
func IsAtom(data []byte) bool {

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

		f.Entry = append(f.Entry, &e)
	}

	f.parseContent()
	return &f, err
}

func (f *Feed) parseContent() {
	for _, c := range f.Category {
		c.parseContent()
	}

	for _, e := range f.Entry {
		e.parseContent()
	}

	for _, l := range f.Link {
		l.parseContent()
	}

	f.Rights.parseContent()
	f.Subtitle.parseContent()
	f.Title.parseContent()
}

func (e *Entry) parseContent() {
	for _, c := range e.Category {
		c.parseContent()
	}

	for _, l := range e.Link {
		l.parseContent()
	}

	e.Content.parseContent()
	e.Rights.parseContent()
	e.Source.parseContent()
	e.Summary.parseContent()
	e.Title.parseContent()
}

func (c *Category) parseContent() {
	if c.XmlContent != "" {
		c.Content = c.XmlContent
	} else {
		c.Content = c.TextContent
	}
}

func (l *Link) parseContent() {
	if l.XmlContent != "" {
		l.Content = l.XmlContent
	} else {
		l.Content = l.TextContent
	}
}

func (c *Content) parseContent() {
	if c.Type == "xhtml" {
		c.Content = c.XmlContent
		return
	}

	if c.Type == "text" || c.Type == "html" {
		c.Content = c.TextContent
		return
	}

	if c.XmlContent != "" {
		c.Content = c.XmlContent
	} else {
		c.Content = c.TextContent
	}

}

func (s *Source) parseContent() {
	for _, c := range s.Category {
		c.parseContent()
	}

	for _, l := range s.Link {
		l.parseContent()
	}

	s.Rights.parseContent()
	s.Subtitle.parseContent()
	s.Title.parseContent()
}

func (t *Text) parseContent() {
	if t.Type == "xhtml" {
		t.Content = t.XmlContent
	} else {
		t.Content = t.TextContent
	}
}
