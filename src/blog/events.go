package blogs

import (
  "time"
)

type BlogPostCreated struct {
  ID string
  Author string
  Timestamp time.Time
  Title string
  Post string
}
