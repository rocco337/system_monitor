package main

import (
	"testing"
	"time"
)

func TestGetTemperatures(t *testing.T) {

	t.Run("GetTemperatures - should pass", func(t *testing.T) {
		index := 0
		ticker := time.NewTicker(100 * time.Millisecond)
		for range ticker.C {
			GetTemperatures()
			index++
			if index > 10 {
				return
			}
		}
	})
}
