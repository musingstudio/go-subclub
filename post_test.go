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

func TestUpdatePost(t *testing.T) {
	c := NewClient(os.Getenv("API_KEY"))

	p, err := c.Post(&PostParams{
		Content: "This post will be edited.",
	})
	if err != nil {
		t.Error(err)
	}
	if p == nil || !p.Success {
		t.Fatalf("FAILED to post")
	}

	ep, err := c.UpdatePost(p.PostID, &PostParams{
		Content: "UPDATED post!",
	})
	if err != nil {
		t.Error(err)
	}
	if !ep.Success {
		t.Fatalf("Success: %t", ep.Success)
	}

	t.Logf("Updated post: %+v", ep)
}

func TestDeletePost(t *testing.T) {
	c := NewClient(os.Getenv("API_KEY"))

	p, err := c.Post(&PostParams{
		Content: "This post will be deleted.",
	})
	if err != nil {
		t.Error(err)
	}
	if p == nil || !p.Success {
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
