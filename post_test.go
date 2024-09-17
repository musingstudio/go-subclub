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
