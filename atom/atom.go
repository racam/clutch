//Refer to the RFC 4287
//link : https://tools.ietf.org/html/rfc4287

package atom

//https://tools.ietf.org/html/rfc4287#section-2
type CommonAttributes struct {
	Base string `xml:"base,attr"`
	Lang string `xml:"lang,attr"`
}

//https://tools.ietf.org/html/rfc4287#section-4.1.1
type Feed struct {
	CommonAttributes
	Author      []*Person   `xml:"author"`
	Category    []*Category `xml:"category"`
	Contributor []*Person   `xml:"contributor"`
	Entry       []*Entry    `xml:"entry"`
	Generator   Generator   `xml:"generator"`
	Icon        Icon        `xml:"icon"`
	ID          Id          `xml:"id"`
	Link        []*Link     `xml:"link"`
	Logo        Logo        `xml:"logo"`
	Rights      Text        `xml:"rights"`
	Subtitle    Text        `xml:"subtitle"`
	Title       Text        `xml:"title"`
	Updated     Date        `xml:"updated"`
}

//https://tools.ietf.org/html/rfc4287#section-4.1.2
type Entry struct {
	CommonAttributes
	Author      []*Person   `xml:"author"`
	Category    []*Category `xml:"category"`
	Content     Content     `xml:"content"`
	Contributor []*Person   `xml:"contributor"`
	ID          Id          `xml:"id"`
	Link        []*Link     `xml:"link"`
	Published   Date        `xml:"published"`
	Rights      Text        `xml:"rights"`
	Source      Source      `xml:"source"`
	Summary     Text        `xml:"summary"`
	Title       Text        `xml:"title"`
	Updated     Date        `xml:"updated"`
}

//https://tools.ietf.org/html/rfc4287#section-4.1.3
type Content struct {
	CommonAttributes
	Text
	Src  string `xml:"src,attr"`
	Type string `xml:"type,attr"`
}

//https://tools.ietf.org/html/rfc4287#section-4.2.7
type Link struct {
	CommonAttributes
	Content     string `-` //Fill with parseContent functions
	Href        string `xml:"href,attr"`
	Hreflang    string `xml:"hreflang,attr"`
	Length      string `xml:"length,attr"`
	Rel         string `xml:"rel,attr"`
	TextContent string `xml:",chardata"`
	Title       string `xml:"title,attr"`
	Type        string `xml:"type,attr"`
	XmlContent  string `xml:",innerxml"`
}

//https://tools.ietf.org/html/rfc4287#section-4.2.11
type Source struct {
	CommonAttributes
	Author      []*Person   `xml:"author"`
	Category    []*Category `xml:"category"`
	Contributor []*Person   `xml:"contributor"`
	Generator   Generator   `xml:"generator"`
	Icon        Icon        `xml:"icon"`
	ID          Id          `xml:"id"`
	Link        []*Link     `xml:"link"`
	Logo        Logo        `xml:"logo"`
	Rights      Text        `xml:"rights"`
	Subtitle    Text        `xml:"subtitle"`
	Title       Text        `xml:"title"`
	Updated     Date        `xml:"updated"`
}

//https://tools.ietf.org/html/rfc4287#section-3.1
type Text struct {
	CommonAttributes
	Content     string `-` //Fill with parseContent functions
	TextContent string `xml:",chardata"`
	Type        string `xml:"type,attr"`
	XmlContent  string `xml:",innerxml"`
}

//https://tools.ietf.org/html/rfc4287#section-3.2
type Person struct {
	CommonAttributes
	Email string `xml:"email"`
	Name  string `xml:"name"`
	Uri   string `xml:"uri"`
}

//https://tools.ietf.org/html/rfc4287#section-3.3
type Date struct {
	CommonAttributes
	DateTime string `xml:",chardata"`
}

//https://tools.ietf.org/html/rfc4287#section-4.2.2
type Category struct {
	CommonAttributes
	Content     string `-` //Fill with parseContent functions
	Label       string `xml:"label,attr"`
	Scheme      string `xml:"scheme,attr"`
	Term        string `xml:"term,attr"`
	TextContent string `xml:",chardata"`
	XmlContent  string `xml:",innerxml"`
}

//https://tools.ietf.org/html/rfc4287#section-4.2.4
type Generator struct {
	CommonAttributes
	Content string `xml:",chardata"`
	Version string `xml:"version,attr"`
	Uri     string `xml:"uri,attr"`
}

type CommonUri struct {
	CommonAttributes
	Uri string `xml:",chardata"`
}

//Icon : https://tools.ietf.org/html/rfc4287#section-4.2.5
type Icon CommonUri

//Icon : https://tools.ietf.org/html/rfc4287#section-4.2.6
type Id CommonUri

//Logo : https://tools.ietf.org/html/rfc4287#section-4.2.8
type Logo CommonUri
