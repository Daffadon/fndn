package dto

import "time"

type StopwatchModel struct {
	startTime time.Time
	elapsed   time.Duration
	running   bool
}

func (s *StopwatchModel) Start() {
	s.startTime = time.Now()
	s.running = true
}

func (s *StopwatchModel) Stop() {
	if s.running {
		s.elapsed = time.Since(s.startTime)
		s.running = false
	}
}

func (s *StopwatchModel) Elapsed() time.Duration {
	if s.running {
		return time.Since(s.startTime)
	}
	return s.elapsed
}
