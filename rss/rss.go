// Refer to https://cyber.law.harvard.edu/rss/rss.html

package rss

import (
	"encoding/xml"
)

// https://cyber.law.harvard.edu/rss/rss.html#whatIsRss
type Rss struct {
	Channel Channel  `xml:"channel"`
	Version string   `xml:"version,attr"`
	XMLName xml.Name `rss`
}

// https://cyber.law.harvard.edu/rss/rss.html#requiredChannelElements
type Channel struct {
	Category       []Category `xml:"category"`
	Cloud          string     `xml:"cloud"`
	Copyright      string     `xml:"copyright"`
	Description    string     `xml:"description"`
	Docs           string     `xml:"docs"`
	Generator      string     `xml:"generator"`
	Image          Image      `xml:"image"`
	Item           []Item     `xml:"item"`
	Language       string     `xml:"language"`
	ManagingEditor string     `xml:"managingeditor"`
	Link           string     `xml:"link"`
	LastBuildDate  string     `xml:"lastbuilddate"`
	PubDate        string     `xml:"pubdate"`
	Rating         string     `xml:"rating"`
	SkipDays       string     `xml:"skipdays"`
	SkipHours      string     `xml:"skiphours"`
	TextInput      TextInput  `xml:"textinput"`
	Title          string     `xml:"source"`
	Ttl            string     `xml:"ttl"`
	WebMaster      string     `xml:"webmaster"`
}

// https://cyber.law.harvard.edu/rss/rss.html#hrelementsOfLtitemgt
type Item struct {
	Author      string     `xml:"author"`
	Category    []Category `xml:"category"`
	Comments    string     `xml:"comments"`
	Description string     `xml:"description"`
	Enclosure   string     `xml:"enclosure"`
	Guid        Guid       `xml:"guid"`
	Link        string     `xml:"link"`
	PubDate     string     `xml:"pubdate"`
	Source      Source     `xml:"source"`
	Title       string     `xml:"title"`
}

// https://cyber.law.harvard.edu/rss/rss.html#ltimagegtSubelementOfLtchannelgt
type Image struct {
	Description string `xml:"description"`
	Height      string `xml:"height"`
	Link        string `xml:"link"`
	Title       string `xml:"title"`
	Url         string `xml:"url"`
	Width       string `xml:"width"`
}

// https://cyber.law.harvard.edu/rss/rss.html#lttextinputgtSubelementOfLtchannelgt
type TextInput struct {
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Name        string `xml:"name"`
	Title       string `xml:"title"`
}

// https://cyber.law.harvard.edu/rss/rss.html#ltsourcegtSubelementOfLtitemgt
type Source struct {
	Title string `xml:",chardata"`
	Url   string `xml:"url,attr"`
}

// https://cyber.law.harvard.edu/rss/rss.html#ltcategorygtSubelementOfLtitemgt
type Category struct {
	Content string `xml:",chardata"`
	Domain  string `xml:"domain,attr"`
}

// https://cyber.law.harvard.edu/rss/rss.html#ltguidgtSubelementOfLtitemgt
type Guid struct {
	Content     string `xml:",chardata"`
	IsPermaLink string `xml:"ispermalink,attr"`
}
