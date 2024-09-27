# go-subclub

[![Go Reference](https://pkg.go.dev/badge/github.com/musingstudio/go-subclub.svg)](https://pkg.go.dev/github.com/musingstudio/go-subclub)

A Go (golang) library for interacting with the [sub.club](https://sub.club) API.

## Example Usage

```go
package main

import (
	"fmt"
	"github.com/musingstudio/go-subclub"
)

func main() {
	c := subclub.NewClient("YOUR SUB.CLUB KEY")

	p, err := c.Post(&subclub.PostParams{
		Content: "This is my premium post",
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Post: %+v", p)
}

```