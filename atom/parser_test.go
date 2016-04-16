package atom

import (
	"io/ioutil"
	"testing"
)

func TestIsDeclared(t *testing.T) {
	var prefix = "../testdata/atom/"

	var namespaces = []struct {
		filename string // input file
		expected bool   // expected result
	}{
		{"unit_01_IsDeclared.xml", true},
		{"unit_02_IsDeclared.xml", true},
		{"unit_03_IsDeclared.xml", false},
	}

	for _, ns := range namespaces {

		f, err := ioutil.ReadFile(prefix + ns.filename)

		if err != nil {
			t.Errorf("[Atom][Unit][IsDeclared] file '%s' : is missing",
				prefix+ns.filename)
		} else {
			res := IsDeclared(f)

			//Test if isDeclared return an expected result
			if res != ns.expected {
				t.Errorf(`[Atom][Unit] file '%s' : expected result '%t',
        actual %t`, ns.filename, ns.expected, res)
			}
		}
	}
}

func TestTextParseContent(t *testing.T) {
	var text Text
	text.XMLContent = "xml content"
	text.TextContent = "text content"

	//empty Type
	text.parseContent()

	// If the "type" attribute is not provided, Atom Processors MUST behave as
	// though it were present with a value of "text"
	// Source : https://tools.ietf.org/html/rfc4287#section-3.1.1
	if text.Type != "text" {
		t.Errorf("[Atom][Unit] Text.Type : expected 'text', actual '%s'", text.Type)
	}

	if text.Content != text.TextContent {
		t.Errorf("[Atom][Unit] Text.Content : expected '%s', actual '%s'",
			text.TextContent, text.Content)
	}

	text.Type = "text"
	text.parseContent()
	if text.Content != text.TextContent {
		t.Errorf("[Atom][Unit] Text.Content : expected '%s', actual '%s'",
			text.TextContent, text.Content)
	}

	text.Type = "xhtml"
	text.parseContent()
	if text.Content != text.XMLContent {
		t.Errorf("[Atom][Unit] Text.Content : expected '%s', actual '%s'",
			text.TextContent, text.Content)
	}
}

func TestContentParseContent(t *testing.T) {
	var c Content
	c.XMLContent = "xml content"
	c.TextContent = "text content"

	c.parseContent()
	if c.Type != "text" {
		t.Errorf("[Atom][Unit] Content.Type : expected 'text', actual '%s'", c.Type)
	}

	if c.Content != c.TextContent {
		t.Errorf("[Atom][Unit] Content.Content : expected '%s', actual '%s'",
			c.TextContent, c.Content)
	}

	c.Type = "xhtml"
	c.parseContent()

	if c.Content != c.XMLContent {
		t.Errorf("[Atom][Unit] Content.Content : expected '%s', actual '%s'",
			c.XMLContent, c.Content)
	}
}

func TestContentParseContent2(t *testing.T) {
	var c Content
	c.TextContent = "text content"
	c.Src = "src"

	c.parseContent()
	if c.Type != "" {
		t.Errorf("[Atom][Unit] Content.Type : expected '', actual '%s'", c.Type)
	}

	if c.Content != c.TextContent {
		t.Errorf("[Atom][Unit] Content.Content : expected '%s', actual '%s'",
			c.TextContent, c.Content)
	}

	c.XMLContent = "xml content"
	c.parseContent()

	if c.Content != c.XMLContent {
		t.Errorf("[Atom][Unit] Content.Content : expected '%s', actual '%s'",
			c.XMLContent, c.Content)
	}
}

func TestLinkParseContent(t *testing.T) {
	var l Link
	l.XMLContent = "xml content"
	l.TextContent = "text content"

	//empty rel
	l.parseContent()

	// If the "rel" attribute is not present, the link element MUST be
	// interpreted as if the link relation type is "alternate".
	// source: https://tools.ietf.org/html/rfc4287#section-4.2.7.2
	if l.Rel != "alternate" {
		t.Errorf("[Atom][Unit] Link.Type : expected 'alternate', actual '%s'",
			l.Rel)
	}

	if l.Content != l.TextContent {
		t.Errorf("[Atom][Unit] Link.Content : expected '%s', actual '%s'",
			l.TextContent, l.Content)
	}

	l.XMLContent = "xml content"
	l.TextContent = ""
	l.parseContent()
	if l.Content != l.XMLContent {
		t.Errorf("[Atom][Unit] Link.Content : expected '%s', actual '%s'",
			l.XMLContent, l.Content)
	}
}

func TestCategoryParseContent(t *testing.T) {
	var c Category
	c.TextContent = "text content"
	c.XMLContent = "xml content"

	c.parseContent()
	if c.Content != c.TextContent {
		t.Errorf("[Atom][Unit] Category.Content : expected '%s', actual '%s'",
			c.TextContent, c.Content)
	}

	c.TextContent = ""
	c.XMLContent = "xml content"
	c.parseContent()
	if c.Content != c.XMLContent {
		t.Errorf("[Atom][Unit] Category.Content : expected '%s', actual '%s'",
			c.XMLContent, c.Content)
	}
}
