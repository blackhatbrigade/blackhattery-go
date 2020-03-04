package main

import (
  "fmt"
  "net/http"
  "time"

  "github.com/blackhatbrigade/blackhattery-go/blogs"

	"github.com/jetbasrawi/go.cqrs"
)

var (
	dispatcher ycq.Dispatcher
  blogEngine blogs.BlogPostsFacade
)

func main() {
  log, err := NewLogger(1)
  if err != nil {
    fmt.Println(err)
  }

  blogEngine = blogs.NewBlogEngine()
  //eventBus := ycq.NewInternalEventBus()

  mux := setupHandlers()
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}

func setupHandlers() *http.ServeMux {
	mux := http.NewServeMux()
  log, err := NewLogger(1)
  if err != nil {
    fmt.Println(err)
  }

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		posts := blogEngine.GetBlogList()

    for _, v := range posts {
      fmt.Fprintf(w, "%s : by %s", v.Title, v.Author)
    }
	})

	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				log.Fatal(err)
			}

			id := ycq.NewUUID()
      em := ycq.NewCommandMessage(id, &blogs.BlogPostCreated{
        ID: ycq.NewUUID(),
				Author: r.Form.Get("Author"),
        Timestamp: time.Now(),
        Title: r.Form.Get("Title"),
        Post: r.Form.Get("Post"),
			})

			err = dispatcher.Dispatch(em)
			if err != nil {
				log.Info(err)
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	})

  return mux
}
