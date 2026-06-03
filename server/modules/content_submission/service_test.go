package content_submission

import (
	"testing"

	contentService "bemunair2026/server/modules/content_submission/service"
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
		if got := contentService.ValidTransition(tt.from, tt.to); got != tt.ok {
			t.Fatalf("%s -> %s = %v", tt.from, tt.to, got)
		}
	}
}
