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
	"github.com/racam/clutch/atom"
	"github.com/racam/clutch/rss"
)

// FeedType represents one of the possible feed
// types that we can detect.
// Source for FeedType : https://github.com/mmcdole/gofeed
type FeedType int

const (
	// FeedTypeUnknown represents a feed that could not have its
	// type determiend.
	FeedTypeUnknown FeedType = iota
	// FeedTypeAtom repesents an Atom feed
	FeedTypeAtom
	// FeedTypeRSS represents an RSS feed
	FeedTypeRSS
)

// Feed is the root element of this common structure for RSS/Atom.
// It has few metadata and contains all the 'entry'
type Feed struct {
	Author      *string
	Category    []*string
	Description *string
	Entry       []Entry
	Generator   *string
	Language    *string
	Logo        *string
	Link        *string
	Rights      *string
	Title       *string
	Updated     *string
	RSS         *rss.RSS
	Atom        *atom.Feed
	FeedType    FeedType
}

// Entry is the principal element of this common structure for RSS/Atom.
// An entry is acting as a container for metadata and data associated
// with the entry. An entry may represent one news or an article for example.
type Entry struct {
	Author      *string
	Category    []*string
	Description *string
	ID          *string
	Link        *string
	Published   *string
	Source      Source
	Title       *string
}

// Source is useful if an entry is forwarded from an existing RSS/Atom feed
type Source struct {
	Title *string
	URL   *string
}
