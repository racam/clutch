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

package clutch

import (
	"errors"

	"github.com/racam/clutch/atom"
	"github.com/racam/clutch/rss"
)

// Parse parses the ATOM-encoded data into an atom.Feed struct and return it.
// The ATOM parsing do 2 steps :
// * Check that the document is well declared as an ATOM document
// * Parse the structure to fill additionnals fields
func Parse(data []byte) (*Feed, error) {
	res := Feed{}

	if rss, err := rss.Parse(data); err == nil {
		res.RSS = rss
		res.FeedType = FeedTypeRSS
		res.parseRSS()
	} else if atom, err2 := atom.Parse(data); err2 == nil {
		res.Atom = atom
		res.FeedType = FeedTypeAtom
		res.parseAtom()
	} else {
		return nil, errors.New("Data is not recognize as RSS 2.0 or Atom 1.0")
	}

	return &res, nil
}

func (f *Feed) parseRSS() {
	f.Author = &f.RSS.Channel.ManagingEditor
	f.Description = &f.RSS.Channel.Description
	f.Generator = &f.RSS.Channel.Generator
	f.Language = &f.RSS.Channel.Language
	f.Link = &f.RSS.Channel.Link
	f.Logo = &f.RSS.Channel.Image.URL
	f.Rights = &f.RSS.Channel.Copyright
	f.Title = &f.RSS.Channel.Title
	f.Updated = &f.RSS.Channel.PubDate

	f.Category = make([]*string, len(f.RSS.Channel.Category))
	for index := range f.Category {
		f.Category[index] = &f.RSS.Channel.Category[index].Content
	}

	f.Entry = make([]Entry, len(f.RSS.Channel.Item))
	for index := range f.Entry {
		f.Entry[index].Author = &f.RSS.Channel.Item[index].Author
		f.Entry[index].Title = &f.RSS.Channel.Item[index].Title
		f.Entry[index].Description = &f.RSS.Channel.Item[index].Description
		f.Entry[index].ID = &f.RSS.Channel.Item[index].GUID.Content
		f.Entry[index].Link = &f.RSS.Channel.Item[index].Link
		f.Entry[index].Published = &f.RSS.Channel.Item[index].PubDate
		f.Entry[index].Source.Title = &f.RSS.Channel.Item[index].Source.Title
		f.Entry[index].Source.URL = &f.RSS.Channel.Item[index].Source.URL

		nbCat := len(f.RSS.Channel.Item[index].Category)
		f.Entry[index].Category = make([]*string, nbCat)
		for index2 := range f.Entry[index].Category {
			var ptr *string
			ptr = &f.RSS.Channel.Item[index].Category[index2].Content
			f.Entry[index].Category[index2] = ptr
		}
	}
}

func (f *Feed) parseAtom() {
	if len(f.Atom.Author) > 0 {
		f.Author = &f.Atom.Author[0].Name
	} else {
		tmp := "" // Avoid nil pointer
		f.Author = &tmp
	}

	if len(f.Atom.Link) > 0 {
		f.Link = &f.Atom.Link[0].Href
	} else {
		tmp := "" // Avoid nil pointer
		f.Link = &tmp
	}

	f.Description = &f.Atom.Subtitle.Content
	f.Generator = &f.Atom.Generator.Content
	f.Language = &f.Atom.CommonAttributes.Lang
	f.Logo = &f.Atom.Logo.URI
	f.Rights = &f.Atom.Rights.Content
	f.Title = &f.Atom.Title.Content
	f.Updated = &f.Atom.Updated.DateTime

	f.Category = make([]*string, len(f.Atom.Category))
	for index := range f.Category {
		f.Category[index] = &f.Atom.Category[index].Content
	}

	f.Entry = make([]Entry, len(f.Atom.Entry))
	for index := range f.Entry {
		if len(f.Atom.Entry[index].Author) > 0 {
			f.Entry[index].Author = &f.Atom.Entry[index].Author[0].Name
		} else {
			tmp := "" // Avoid nil pointer
			f.Author = &tmp
		}

		if len(f.Atom.Entry[index].Link) > 0 {
			f.Entry[index].Link = &f.Atom.Entry[index].Link[0].Href
		} else {
			tmp := "" // Avoid nil pointer
			f.Link = &tmp
		}

		f.Entry[index].Title = &f.Atom.Entry[index].Title.Content
		f.Entry[index].Description = &f.Atom.Entry[index].Summary.Content
		f.Entry[index].ID = &f.Atom.Entry[index].ID.URI
		f.Entry[index].Published = &f.Atom.Entry[index].Published.DateTime
		f.Entry[index].Source.Title = &f.Atom.Entry[index].Source.Title.Content
		f.Entry[index].Source.URL = &f.Atom.Entry[index].Source.ID.URI

		nbCat := len(f.Atom.Entry[index].Category)
		f.Entry[index].Category = make([]*string, nbCat)
		for index2 := range f.Entry[index].Category {
			var ptr *string
			ptr = &f.Atom.Entry[index].Category[index2].Content
			f.Entry[index].Category[index2] = ptr
		}
	}
}
