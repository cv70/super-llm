package conv

import (
	"context"

	"github.com/JohannesKaufmann/html-to-markdown/v2"
)

func HTML2Markdown(ctx context.Context, html string) (string, error) {
	markdown, err := htmltomarkdown.ConvertString(html)
	if err != nil {
		return "", err
	}
	return markdown, nil
}
