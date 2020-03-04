package blogs

import (
  "testing"
)

// TestNewBlogPost 
func TestNewBlogPostCanBeCreated(t *testing.T) {
  //
  blogPost := NewBlogPost("111111111")

  if blogPost == nil {
    t.Errorf("Expected blogPost to exist")
  }
}
