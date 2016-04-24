package engine

import (
	"github.com/Sirupsen/logrus"
	"github.com/hersshel/hersshel/model"
	"github.com/hersshel/hersshel/store"
	"github.com/mmcdole/gofeed"

	"golang.org/x/net/context"
)

// Engine is an interface that handles lifecycle
// of goroutines used to fetch RSS feed.
type Engine interface {
	Schedule(context.Context, *model.Feed)
}

type engine struct {
	parser *gofeed.Parser
}

// New creates a new Engine that can be used to schedule
// fetching of RSS feed.
func New() Engine {
	return &engine{
		parser: gofeed.NewParser(),
	}
}

func (e *engine) Schedule(ctx context.Context, feed *model.Feed) {
	f, err := e.parser.ParseURL(feed.URL)
	if err != nil {
		logrus.WithError(err).Errorf("error parsing the feed")
		return
	}
	logrus.WithField("feed", f).Infof("feed fetched")
	if len(f.Items) > 0 {
		items := make([]*model.Item, len(f.Items))
		for k, v := range f.Items {
			items[k] = &model.Item{
				Link:      v.Link,
				Title:     v.Title,
				Content:   v.Content,
				CreatedAt: v.PublishedParsed,
				UpdatedAt: v.UpdatedParsed,
				FeedID:    feed.ID,
			}
			if v.Author != nil {
				items[k].Author = v.Author.Name
			}
		}
		err = store.CreateItems(ctx, items)
		if err != nil {
			logrus.WithError(err).Errorf("error storing the feed")
			return
		}
	}
}
