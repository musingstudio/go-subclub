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

func TestDeletePost(t *testing.T) {
	c := NewClient(os.Getenv("API_KEY"))

	p, err := c.Post(&PostParams{
		Content: "This post will be deleted.",
	})
	if err != nil {
		t.Error(err)
	}
	if !p.Success {
		t.Fatalf("FAILED to post")
	}

	t.Logf("Posted: %+v", p)

	r, err := c.DeletePost(p.PostID)
	if err != nil {
		t.Error(err)
	}
	if !r.Deleted {
		t.Fatalf("NOT deleted")
	}
}
