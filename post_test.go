package subclub

import (
	"os"
	"testing"
)

func TestPost(t *testing.T) {
	c := NewClient(os.Getenv("API_KEY"))

	p, err := c.Post(&PostParams{
		Content: "Test post",
	})
	if err != nil {
		t.Error(err)
	}

	t.Logf("Post: %+v", p)
}

func TestEditPost(t *testing.T) {
	c := NewClient(os.Getenv("API_KEY"))

	p, err := c.EditPost(&PostUpdateParams{
		PostID:  os.Getenv("POST_ID"),
		Content: "Test post (updated)",
	})
	if err != nil {
		t.Error(err)
	}
	if !p.Success {
		t.Fatalf("Success: %t", p.Success)
	}

	t.Logf("Updated post: %+v", p)
}
