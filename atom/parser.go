package atom

import (
	"encoding/xml"
	"errors"
	"strings"
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

		f.Entry = append(f.Entry, e)
	}

	f.parseContent()
	err = f.Check(feedIsAbsent)
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
	if c.XmlContent != "" {
		c.Content = c.XmlContent
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

	if l.XmlContent != "" {
		l.Content = l.XmlContent
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
		c.Type = "text"
	}

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
		t.Type = "text"
	}

	if t.Type == "xhtml" {
		t.Content = t.XmlContent
	} else {
		t.Content = t.TextContent
	}
}

func (f *Feed) Check(feedIsAbsent bool) error {
	if !feedIsAbsent {
		if f.ID.Uri == "" {
			return errors.New("atom:feed elements MUST contain exactly one " +
				"atom:id element.")
		}

		if f.Title.Content == "" {
			return errors.New("atom:feed elements MUST contain exactly one " +
				"atom:title element.")
		}

		if f.Updated.DateTime == "" {
			return errors.New("atom:feed elements MUST contain exactly one " +
				"atom:updated element.")
		}
	}

	for _, c := range f.Category {
		err := c.Check()
		if err != nil {
			return errors.New("atom:feed " + err.Error())
		}
	}

	for _, l := range f.Link {
		err := l.Check()
		if err != nil {
			return errors.New("atom:feed " + err.Error())
		}
	}

	for _, a := range f.Author {
		err := a.Check()
		if err != nil {
			return errors.New("atom:feed atom:auhtor " + err.Error())
		}
	}
	// atom:feed elements MUST contain one or more atom:author elements,
	// unless all of the atom:feed element's child atom:entry elements
	// contain at least one atom:author element.
	requiredAuthor := len(f.Author) == 0

	for _, c := range f.Contributor {
		err := c.Check()
		if err != nil {
			return errors.New("atom:feed atom:contributor " + err.Error())
		}
	}

	if err := f.Rights.Check(); err != nil {
		return errors.New("atom:feed atom:rights " + err.Error())
	}

	if err := f.Subtitle.Check(); err != nil {
		return errors.New("atom:feed atom:subtitle " + err.Error())
	}

	if err := f.Title.Check(); err != nil {
		return errors.New("atom:feed atom:title " + err.Error())
	}

	for _, entry := range f.Entry {
		err := entry.Check(requiredAuthor)
		if err != nil {
			return errors.New("atom:feed " + err.Error())
		}
	}

	return nil
}

func (e *Entry) Check(requiredAuthor bool) error {
	if e.ID.Uri == "" {
		return errors.New("atom:entry elements MUST contain exactly one " +
			"atom:id element.")
	}

	if e.Title.Content == "" {
		return errors.New("atom:entry elements MUST contain exactly one " +
			"atom:title element.")
	}

	if e.Updated.DateTime == "" {
		return errors.New("atom:entry elements MUST contain exactly one " +
			"atom:updated element.")
	}

	for _, c := range e.Category {
		err := c.Check()
		if err != nil {
			return errors.New("atom:entry " + err.Error())
		}
	}

	for _, c := range e.Contributor {
		err := c.Check()
		if err != nil {
			return errors.New("atom:entry atom:contributor" + err.Error())
		}
	}

	for _, a := range e.Author {
		err := a.Check()
		if err != nil {
			return errors.New("atom:entry atom:author" + err.Error())
		}
	}

	// atom:entry elements MUST contain one or more atom:author elements,
	// unless the atom:entry contains an atom:source element that
	// contains an atom:author element or, in an Atom Feed Document, the
	// atom:feed element contains an atom:author element itself.
	requiredAuthor = requiredAuthor && len(e.Author) == 0

	if err := e.Source.Check(requiredAuthor); err != nil {
		return errors.New("atom:entry " + err.Error())
	}

	if err := e.Rights.Check(); err != nil {
		return errors.New("atom:entry atom:rights " + err.Error())
	}

	if err := e.Summary.Check(); err != nil {
		return errors.New("atom:entry atom:summary " + err.Error())
	}

	if err := e.Title.Check(); err != nil {
		return errors.New("atom:entry atom:title " + err.Error())
	}

	return nil
}

func (s *Source) Check(requiredAuthor bool) error {

	if requiredAuthor && len(s.Author) == 0 {
		return errors.New("atom:source. There is no atom:author. It MUST " +
			"have a least one in atom:feed, atom:entry or atom:source.")
	}

	for _, a := range s.Author {
		err := a.Check()
		if err != nil {
			return errors.New("atom:source atom:author" + err.Error())
		}
	}

	for _, c := range s.Category {
		err := c.Check()
		if err != nil {
			return errors.New("atom:source " + err.Error())
		}
	}

	for _, c := range s.Contributor {
		err := c.Check()
		if err != nil {
			return errors.New("atom:source atom:contributor" + err.Error())
		}
	}

	for _, l := range s.Link {
		err := l.Check()
		if err != nil {
			return errors.New("atom:source " + err.Error())
		}
	}

	if err := s.Rights.Check(); err != nil {
		return errors.New("atom:source atom:rights" + err.Error())
	}

	if err := s.Subtitle.Check(); err != nil {
		return errors.New("atom:source atom:subtitle" + err.Error())
	}

	if err := s.Title.Check(); err != nil {
		return errors.New("atom:source atom:title" + err.Error())
	}

	return nil
}

func (c *Category) Check() error {

	// https://tools.ietf.org/html/rfc4287#section-4.2.2.1
	if c.Term == "" {
		return errors.New("atom:category elements MUST have a 'term' attribute.")
	}

	return nil
}

func (c *Content) Check() error {

	// https://tools.ietf.org/html/rfc4287#section-4.1.3.2
	if c.Src != "" && c.Content != "" {
		return errors.New("atom:content, If the 'src' attribute is present, " +
			"atom:content MUST be empty.")
	}

	// https://tools.ietf.org/html/rfc4287#section-4.1.3.2 RULE N°1
	if c.Type == "text" && c.AnyContent != "" {
		return errors.New("If the value of 'type' is 'text', the content of " +
			"atom:content MUST NOT contain child elements.")
	}

	// https://tools.ietf.org/html/rfc4287#section-4.1.3.2 RULE N°2
	if c.Type == "html" && c.AnyContent != "" {
		return errors.New("If the value of 'type' is 'html', the content of " +
			"atom:content MUST NOT contain child elements. Any markup within " +
			"be escaped.")
	}

	// https://tools.ietf.org/html/rfc4287#section-4.1.3.2 RULE N°5
	lowerType := strings.ToLower(c.Type)
	if strings.HasPrefix(lowerType, "text/") && c.AnyContent != "" {
		return errors.New("If the value of 'type' begins with 'text/'' " +
			"(case insensitive), the content of atom:content MUST NOT " +
			"contain child elements.")
	}

	// TODO : check if xhtml is valid
	return nil
}

// https://tools.ietf.org/html/rfc4287#section-4.2.7
func (l *Link) Check() error {

	// https://tools.ietf.org/html/rfc4287#section-4.2.7.1
	if l.Href == "" {
		return errors.New("atom:link elements MUST have an 'href' attribute")
	}

	return nil
}

// https://tools.ietf.org/html/rfc4287#section-3.1
func (t *Text) Check() error {

	// https://tools.ietf.org/html/rfc4287#section-3.1.1
	if t.Type != "text" && t.Type != "html" && t.Type != "xhtml" {
		return errors.New("attr:type MUST be one of 'text', 'html', or 'xhtml'")
	}

	// https://tools.ietf.org/html/rfc4287#section-3.1.1.1
	if t.Type == "text" && t.AnyContent != "" {
		return errors.New("attr:type set to 'text': the content of the Text " +
			"construct MUST NOT contain child elements.")
	}

	// https://tools.ietf.org/html/rfc4287#section-3.1.1.1
	if t.Type == "html" && t.AnyContent != "" {
		return errors.New("attr:type set to 'html': the content of the Text " +
			"construct MUST NOT contain child elements. Any markup within " +
			"be escaped.")
	}

	// TODO : check if xhtml is valid
	return nil
}

// https://tools.ietf.org/html/rfc4287#section-3.2
func (p *Person) Check() error {

	// https://tools.ietf.org/html/rfc4287#section-3.2.1
	if p.Name == "" {
		return errors.New("person constructs MUST contain exactly one " +
			"'atom:name' element.")
	}

	//TODO no more than one URI and Email Element
	return nil
}
