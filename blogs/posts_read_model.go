package blogs

import (
  "time"

	"github.com/jetbasrawi/go.cqrs"
)

// This part of the code handles the read model for the specified entity.
// Mostly used for projections by the command handlers?

var bullshitDatabase *BullShitDatabase

type BlogPostsFacade interface {
  GetBlogList() []*BlogPostListItem
  GetBlogPostDetails(id string) *BlogPostDetails
}

type BlogPostDetails struct {
  ID string
  Author string
  Title string
  Version int
  Timestamp time.Time
  Post string
}

type BlogPostListItem struct {
  ID string
  Author string
  Title string
  Timestamp time.Time
}

type BlogEngine struct {
}

func NewBlogEngine() *BlogEngine {
  if bullshitDatabase == nil {
		bullshitDatabase = NewBullShitDatabase()
  }

  return &BlogEngine{}
}

func (be *BlogEngine) GetBlogList() []*BlogPostListItem {
  return bullshitDatabase.List
}

func (be *BlogEngine) GetBlogPostDetails(id string) *BlogPostDetails {
  if i, ok := bullshitDatabase.Details[id]; ok {
    return i
  }

  return nil
}

type BlogPostView struct {
}

func NewBlogPostView() *BlogPostView {
  if bullshitDatabase == nil {
    bullshitDatabase = NewBullShitDatabase()
  }

  return &BlogPostView{}
}

func (v *BlogPostView) Handle(message ycq.EventMessage) {

	switch event := message.Event().(type) {
  case *BlogPostCreated:
    bullshitDatabase.Details[message.AggregateID()] = &BlogPostDetails{
			ID:        message.AggregateID(),
			Author:    event.Author,
      Timestamp: time.Now(),
			Version:   0,
      Title:     event.Title,
      Post:      event.Post,
    }
  }
}

type BlogPostListView struct {
}

func NewBlogPostListView() *BlogPostListView {
  if bullshitDatabase == nil {
    bullshitDatabase = NewBullShitDatabase()
  }

  return &BlogPostListView{}
}

func (l *BlogPostListView) Handle(message ycq.EventMessage) {
  switch event := message.Event().(type) {

  case *BlogPostCreated:
    bullshitDatabase.List = append(bullshitDatabase.List, &BlogPostListItem{
      ID:         message.AggregateID(),
      Author:     event.Author,
      Title:      event.Title,
      Timestamp:  time.Now(),
    })
  }
}

//-------------------------------------BULLSHIT-------------------------------
//BullShitDatabase In memory Database
type BullShitDatabase struct {
  Details map[string]*BlogPostDetails
  List []*BlogPostListItem
}

func NewBullShitDatabase() *BullShitDatabase {
  return &BullShitDatabase{
    Details: make(map[string]*BlogPostDetails),
  }
}
