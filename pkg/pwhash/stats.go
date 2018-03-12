package pwhash

import (
	"sync"
)

type StatsResponse struct {
	Total   int64 `json:"total"`
	Average int64 `json:"average"`
}

type Stats interface {
	Get() (StatsResponse)
	Time(millis int64)
}

type inMemoryStats struct {
	totalTime int64
	count int64
	lck  sync.RWMutex
}

func NewInMemoryStats() Stats {
	return &inMemoryStats{}
}

func (stats *inMemoryStats) Get() (StatsResponse) {
	stats.lck.RLock()
	defer stats.lck.RUnlock()
	Average := (int64) (0)
	if stats.count != 0 {
		Average = stats.totalTime / stats.count
	}
	return StatsResponse{Total: stats.count, Average: Average}
}

func (stats *inMemoryStats) Time(millis int64) {
	stats.lck.Lock()
	defer stats.lck.Unlock()
	stats.count++
	stats.totalTime += millis
}