package atom

import (
	"errors"
	"strings"
)

func (f *Feed) check(feedIsAbsent bool) error {
	if !feedIsAbsent {
		if f.ID.URI == "" {
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
		err := c.check()
		if err != nil {
			return errors.New("atom:feed " + err.Error())
		}
	}

	for _, l := range f.Link {
		err := l.check()
		if err != nil {
			return errors.New("atom:feed " + err.Error())
		}
	}

	for _, a := range f.Author {
		err := a.check()
		if err != nil {
			return errors.New("atom:feed atom:auhtor " + err.Error())
		}
	}
	// atom:feed elements MUST contain one or more atom:author elements,
	// unless all of the atom:feed element's child atom:entry elements
	// contain at least one atom:author element.
	requiredAuthor := len(f.Author) == 0

	for _, c := range f.Contributor {
		err := c.check()
		if err != nil {
			return errors.New("atom:feed atom:contributor " + err.Error())
		}
	}

	if err := f.Rights.check(); err != nil {
		return errors.New("atom:feed atom:rights " + err.Error())
	}

	if err := f.Subtitle.check(); err != nil {
		return errors.New("atom:feed atom:subtitle " + err.Error())
	}

	if err := f.Title.check(); err != nil {
		return errors.New("atom:feed atom:title " + err.Error())
	}

	for _, entry := range f.Entry {
		err := entry.check(requiredAuthor)
		if err != nil {
			return errors.New("atom:feed " + err.Error())
		}
	}

	return nil
}

func (e *Entry) check(requiredAuthor bool) error {
	if e.ID.URI == "" {
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
		err := c.check()
		if err != nil {
			return errors.New("atom:entry " + err.Error())
		}
	}

	for _, c := range e.Contributor {
		err := c.check()
		if err != nil {
			return errors.New("atom:entry atom:contributor" + err.Error())
		}
	}

	for _, a := range e.Author {
		err := a.check()
		if err != nil {
			return errors.New("atom:entry atom:author" + err.Error())
		}
	}

	// atom:entry elements MUST contain one or more atom:author elements,
	// unless the atom:entry contains an atom:source element that
	// contains an atom:author element or, in an Atom Feed Document, the
	// atom:feed element contains an atom:author element itself.
	requiredAuthor = requiredAuthor && len(e.Author) == 0

	if err := e.Source.check(requiredAuthor); err != nil {
		return errors.New("atom:entry " + err.Error())
	}

	if err := e.Rights.check(); err != nil {
		return errors.New("atom:entry atom:rights " + err.Error())
	}

	if err := e.Summary.check(); err != nil {
		return errors.New("atom:entry atom:summary " + err.Error())
	}

	if err := e.Title.check(); err != nil {
		return errors.New("atom:entry atom:title " + err.Error())
	}

	return nil
}

func (s *Source) check(requiredAuthor bool) error {

	if requiredAuthor && len(s.Author) == 0 {
		return errors.New("atom:source. There is no atom:author. It MUST " +
			"have a least one in atom:feed, atom:entry or atom:source.")
	}

	for _, a := range s.Author {
		err := a.check()
		if err != nil {
			return errors.New("atom:source atom:author" + err.Error())
		}
	}

	for _, c := range s.Category {
		err := c.check()
		if err != nil {
			return errors.New("atom:source " + err.Error())
		}
	}

	for _, c := range s.Contributor {
		err := c.check()
		if err != nil {
			return errors.New("atom:source atom:contributor" + err.Error())
		}
	}

	for _, l := range s.Link {
		err := l.check()
		if err != nil {
			return errors.New("atom:source " + err.Error())
		}
	}

	if err := s.Rights.check(); err != nil {
		return errors.New("atom:source atom:rights" + err.Error())
	}

	if err := s.Subtitle.check(); err != nil {
		return errors.New("atom:source atom:subtitle" + err.Error())
	}

	if err := s.Title.check(); err != nil {
		return errors.New("atom:source atom:title" + err.Error())
	}

	return nil
}

func (c *Category) check() error {

	// https://tools.ietf.org/html/rfc4287#section-4.2.2.1
	if c.Term == "" {
		return errors.New("atom:category elements MUST have a 'term' attribute")
	}

	return nil
}

func (c *Content) check() error {

	// https://tools.ietf.org/html/rfc4287#section-4.1.3.2
	if c.Src != "" && c.Content != "" {
		return errors.New("atom:content, If the 'src' attribute is present, " +
			"atom:content MUST be empty.")
	}

	// https://tools.ietf.org/html/rfc4287#section-4.1.3.2 RULE N°1
	if c.Type == text && c.AnyContent != "" {
		return errors.New("If the value of 'type' is 'text', the content of " +
			"atom:content MUST NOT contain child elements.")
	}

	// https://tools.ietf.org/html/rfc4287#section-4.1.3.2 RULE N°2
	if c.Type == html && c.AnyContent != "" {
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
func (l *Link) check() error {

	// https://tools.ietf.org/html/rfc4287#section-4.2.7.1
	if l.Href == "" {
		return errors.New("atom:link elements MUST have an 'href' attribute")
	}

	return nil
}

// https://tools.ietf.org/html/rfc4287#section-3.1
func (t *Text) check() error {

	// https://tools.ietf.org/html/rfc4287#section-3.1.1
	if t.Type != text && t.Type != html && t.Type != xhtml {
		return errors.New("attr:type MUST be one of 'text', 'html', or 'xhtml'")
	}

	// https://tools.ietf.org/html/rfc4287#section-3.1.1.1
	if t.Type == text && t.AnyContent != "" {
		return errors.New("attr:type set to 'text': the content of the Text " +
			"construct MUST NOT contain child elements.")
	}

	// https://tools.ietf.org/html/rfc4287#section-3.1.1.1
	if t.Type == html && t.AnyContent != "" {
		return errors.New("attr:type set to 'html': the content of the Text " +
			"construct MUST NOT contain child elements. Any markup within " +
			"be escaped.")
	}

	// TODO : check if xhtml is valid
	return nil
}

// https://tools.ietf.org/html/rfc4287#section-3.2
func (p *Person) check() error {

	// https://tools.ietf.org/html/rfc4287#section-3.2.1
	if p.Name == "" {
		return errors.New("person constructs MUST contain exactly one " +
			"'atom:name' element.")
	}

	//TODO no more than one URI and Email Element
	return nil
}
