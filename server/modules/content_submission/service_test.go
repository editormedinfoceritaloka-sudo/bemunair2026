package content_submission

import (
	"testing"

	"bemunair2026/server/database/entities"
)

func TestValidTransition(t *testing.T) {
	tests := []struct {
		from, to string
		ok       bool
	}{
		{entities.StatusPending, entities.StatusInReview, true},
		{entities.StatusInReview, entities.StatusApproved, true},
		{entities.StatusInReview, entities.StatusRejected, true},
		{entities.StatusPending, entities.StatusApproved, false},
	}
	for _, tt := range tests {
		if got := ValidTransition(tt.from, tt.to); got != tt.ok {
			t.Fatalf("%s -> %s = %v", tt.from, tt.to, got)
		}
	}
}
