package subclub

import (
	"fmt"
)

type (
	Post struct {
		Success bool `json:"success"`

		PostID string `json:"postId"`
		Post   string `json:"post"`
		URI    string `json:"uri"`
		URL    string `json:"url"`
	}

	PostParams struct {
		Content string `json:"content"`
	}

	// PostUpdateParams has parameters for editing a sub.club post.
	PostUpdateParams struct {
		PostID  string `json:"postId"`
		Content string `json:"content"`
	}

	// PostDeleteParams has parameters for deleting a sub.club post.
	PostDeleteParams struct {
		PostID string `json:"postId"`
	}

	// PostDeleteResult represents an API response after deleting a sub.club post.
	PostDeleteResult struct {
		Deleted bool `json:"deleted"`
	}
)

func (c *Client) Post(pp *PostParams) (*Post, error) {
	p := &Post{}

	env, err := c.post("/post", pp, p)
	if err != nil {
		return nil, err
	}

	var ok bool
	if p, ok = env.(*Post); !ok {
		return nil, fmt.Errorf("wrong data returned from API")
	}

	return p, nil
}

// EditPost edits the given post with the supplied PostUpdateParams.
func (c *Client) EditPost(pup *PostUpdateParams) (*Post, error) {
	p := &Post{}

	env, err := c.post("/post/edit", pup, p)
	if err != nil {
		return nil, err
	}

	var ok bool
	if p, ok = env.(*Post); !ok {
		return nil, fmt.Errorf("wrong data returned from API")
	}

	return p, nil
}

func (c *Client) DeletePost(postID string) (*PostDeleteResult, error) {
	res := &PostDeleteResult{}

	env, err := c.post("/post/delete", &PostDeleteParams{PostID: postID}, res)
	if err != nil {
		return nil, err
	}

	var ok bool
	if res, ok = env.(*PostDeleteResult); !ok {
		return nil, fmt.Errorf("wrong data returned from API")
	}

	return res, nil
}
