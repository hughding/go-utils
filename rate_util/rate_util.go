package rate_util

import (
	"sync"
)

type RateUtil struct {
	limit int32
	count *int32
	mu    sync.Mutex
}

func NewRateUtil(limit int) *RateUtil {
	count := int32(0)
	return &RateUtil{
		limit: int32(limit),
		count: &count,
		mu:    sync.Mutex{},
	}
}

func (ru *RateUtil) Allow() bool {
	ru.mu.Lock()
	defer ru.mu.Unlock()

	if *ru.count >= ru.limit {
		return false
	}
	*ru.count++
	return true
}

func (ru *RateUtil) Done() bool {
	ru.mu.Lock()
	defer ru.mu.Unlock()

	if *ru.count > 0 {
		*ru.count--
		return true
	}
	return false
}
