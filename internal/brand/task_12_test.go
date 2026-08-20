package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask12(t *testing.T) {
	now := time.Now()
	s := NewService(NewRegistry(), func() time.Time { return now })
	in := Inspection{ID: "i", StartedAt: now.Add(-time.Hour), CompletedAt: now, Sections: []InspectionSection{{Code: "food", Required: true, Score: 90, Evidence: []string{"photo"}}}}
	out, err := s.CompleteInspection(context.Background(), in, 80)
	require.NoError(t, err)
	require.True(t, out.Passed)
}
