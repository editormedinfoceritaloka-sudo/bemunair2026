package content_submission

import (
	"testing"

	"bemunair2026/server/pkg/constants"
)

func TestValidTransition(t *testing.T) {
	tests := []struct {
		from, to string
		ok       bool
	}{
		{constants.StatusPending, constants.StatusInReview, true},
		{constants.StatusInReview, constants.StatusApproved, true},
		{constants.StatusInReview, constants.StatusRejected, true},
		{constants.StatusPending, constants.StatusApproved, false},
	}
	for _, tt := range tests {
		if got := ValidTransition(tt.from, tt.to); got != tt.ok {
			t.Fatalf("%s -> %s = %v", tt.from, tt.to, got)
		}
	}
}
