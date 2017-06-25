package model

import (
	"youtube"
)

type Queue struct {
	Q []QueueItem
}

type QueueItem struct {
	Creator User                 `json:"user"`
	Video   youtube.YoutubeVideo `json:"video"`
}

func (q *Queue) First() QueueItem {
	return q.Q[0]
}

func (q *Queue) PopFirst() {
	q.RemoveIndex(0)
}

func (q *Queue) AddItem(item QueueItem) {
	q.Q = append(q.Q, item)
}

func (q *Queue) RemoveIndex(index int) {
	if q.Count() == 1 {
		q.Q = make([]QueueItem, 0)
	} else {
		q.Q = append(q.Q[:index], q.Q[index+1:]...)
	}
}

func (q *Queue) Count() int {
	return len(q.Q)
}
