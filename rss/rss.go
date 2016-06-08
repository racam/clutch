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

// Package rss please Refer to https://cyber.law.harvard.edu/rss/rss.html
package rss

import (
	"encoding/xml"
)

// RSS is a RSS structure like describe in
// https://cyber.law.harvard.edu/rss/rss.html#whatIsRss
type RSS struct {
	Channel Channel  `xml:"channel"`
	Version string   `xml:"version,attr"`
	XMLName xml.Name `xml:"rss"`
}

// Channel is a RSS structure like describe in
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
	TTL            string     `xml:"ttl"`
	WebMaster      string     `xml:"webmaster"`
}

// Item is a RSS structure like describe in
// https://cyber.law.harvard.edu/rss/rss.html#hrelementsOfLtitemgt
type Item struct {
	Author      string     `xml:"author"`
	Category    []Category `xml:"category"`
	Comments    string     `xml:"comments"`
	Description string     `xml:"description"`
	Enclosure   string     `xml:"enclosure"`
	GUID        GUID       `xml:"guid"`
	Link        string     `xml:"link"`
	PubDate     string     `xml:"pubdate"`
	Source      Source     `xml:"source"`
	Title       string     `xml:"title"`
}

// Image is a RSS structure like describe in
// https://cyber.law.harvard.edu/rss/rss.html#ltimagegtSubelementOfLtchannelgt
type Image struct {
	Description string `xml:"description"`
	Height      string `xml:"height"`
	Link        string `xml:"link"`
	Title       string `xml:"title"`
	URL         string `xml:"url"`
	Width       string `xml:"width"`
}

// TextInput is a RSS structure like describe in
// https://cyber.law.harvard.edu/rss/rss.html#lttextinputgtSubelementOfLtchannelgt
type TextInput struct {
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Name        string `xml:"name"`
	Title       string `xml:"title"`
}

// Source is a RSS structure like describe in
// https://cyber.law.harvard.edu/rss/rss.html#ltsourcegtSubelementOfLtitemgt
type Source struct {
	Title string `xml:",chardata"`
	URL   string `xml:"url,attr"`
}

// Category is a RSS structure like describe in
// https://cyber.law.harvard.edu/rss/rss.html#ltcategorygtSubelementOfLtitemgt
type Category struct {
	Content string `xml:",chardata"`
	Domain  string `xml:"domain,attr"`
}

// GUID is a RSS structure like describe in
// https://cyber.law.harvard.edu/rss/rss.html#ltguidgtSubelementOfLtitemgt
type GUID struct {
	Content     string `xml:",chardata"`
	IsPermaLink string `xml:"ispermalink,attr"`
}
