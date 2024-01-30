package fancy

import (
	"sync"
)

type Item struct {
	title       string
	description string
	by          string
	id          int
	score       int
	time        int
	jobType     string
	url         string
}

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.description }
func (i Item) FilterValue() string { return i.title }

type ItemGenerator struct {
  items []Item
	itemIdx int
	mtx        *sync.Mutex
	get        *sync.Once
}

func (r *ItemGenerator) reset() {
	r.mtx = &sync.Mutex{}
	r.get = &sync.Once{}
  r.items = []Item{}
	r.get.Do(func() {
		ids := GetJobIds()
		for _, id := range ids {
			job, err := FetchJob(id)
			if err != nil {
				continue
			}
      r.items = append(r.items, Item{
        title:       job.JsonTitle,
        description: job.Text,
        by:          job.By,
        id:          job.Id,
        score:       job.Score,
        time:        job.Time,
        jobType:     job.Type,
        url:         job.Url,
      })
		}
	})
}

func (r *ItemGenerator) next() Item {
	if r.mtx == nil {
		r.reset()
	}
	r.mtx.Lock()
	defer r.mtx.Unlock()
	i := Item{
		title:       r.items[r.itemIdx].title,
		description: r.items[r.itemIdx].description,
	}
	r.itemIdx++
	if r.itemIdx >= len(r.items) {
		r.itemIdx = 0
	}
	return i
}
