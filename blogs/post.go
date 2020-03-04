package blogs

import (
	"github.com/jetbasrawi/go.cqrs"
)

type BlogPost struct {
  *ycq.AggregateBase
  activated bool
  views int
}

func NewBlogPost(id string) (bp *BlogPost) {
  bp = &BlogPost{
    AggregateBase: ycq.NewAggregateBase(id),
  }

  return
}

func (bp *BlogPost) Create(author string, title string, post string) error {

  return nil
}

func (bp *BlogPost) Apply(events ycq.EventMessage, isNew bool) {
  return
}
