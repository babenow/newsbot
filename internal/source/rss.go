package source

import (
	"context"

	"github.com/SlyMarbo/rss"
	"github.com/babenow/newsbot/internal/model"
)

type RSSSource struct {
	URL        string
	SourceID   int64
	SourceName string
}

func NewRSSSourceFromModel(m *model.Source) *RSSSource {
	return &RSSSource{
		URL:        m.FeedURL,
		SourceID:   m.ID,
		SourceName: m.Name,
	}
}

func (r RSSSource) Fetch(ctx context.Context) ([]model.Item, error) {
	feed, err := r.loadFeed(ctx, r.URL)
	if err != nil {
		return nil, err
	}

	var items []model.Item

	for _, item := range feed.Items {
		m := model.Item{
			Title:      item.Title,
			Link:       item.Link,
			Date:       item.Date,
			Categories: item.Categories,
			Summary:    item.Summary,
			SourceName: r.SourceName,
		}
		items = append(items, m)
	}

	return items, nil
}

func (r RSSSource) loadFeed(ctx context.Context, url string) (*rss.Feed, error) {
	var (
		feedCh = make(chan *rss.Feed)
		errCh  = make(chan error)
	)

	go func() {
		feed, err := rss.Fetch(url)
		if err != nil {
			errCh <- err
			return
		}

		feedCh <- feed
	}()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err := <-errCh:
		return nil, err
	case feed := <-feedCh:
		return feed, nil
	}
}
