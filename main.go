package main

import (
  "github.com/blackhatbrigade/blackhattery-go/blogs"

	"github.com/jetbasrawi/go.cqrs"
)

var (
	dispatcher ycq.Dispatcher
  blogEngine blogs.BlogPostsFacade
)

func main() {
  blogEngine = blogs.NewBlogEngine()
	eventBus := ycq.NewInternalEventBus()


}
