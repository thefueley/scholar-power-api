// build +integration
package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thefueley/scholar-power-api/internal/history"
)

func TestCreateHistory(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	hist := history.History{
		Date:      "01-01-2000",
		Duration:  "99:00:00",
		Notes:     "999999",
		PlanID:    "999999",
		AthleteID: "999999",
	}

	err = db.CreateHistory(context.Background(), hist)
	require.NoError(t, err)

	wh, err := db.GetHistory(context.Background(), "999999")
	require.NoError(t, err)
	require.NotEmpty(t, wh)

	err = db.DeleteHistory(context.Background(), wh[0].ID)
	require.NoError(t, err)
}

func TestGetHistory(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	hist := history.History{
		Date:      "01-01-2000",
		Duration:  "99:00:00",
		Notes:     "999999",
		PlanID:    "999999",
		AthleteID: "999999",
	}

	err = db.CreateHistory(context.Background(), hist)
	require.NoError(t, err)

	workouts, err := db.GetHistory(context.Background(), "999999")
	require.NoError(t, err)
	require.NotEmpty(t, workouts)

	require.Equal(t, workouts[0].PlanID, "999999")

	err = db.DeleteHistory(context.Background(), workouts[0].ID)
	require.NoError(t, err)
}

func TestUpdateHistory(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	hist := history.History{
		Date:      "01-01-2000",
		Duration:  "99:00:00",
		Notes:     "999999",
		PlanID:    "999999",
		AthleteID: "999999",
	}

	err = db.CreateHistory(context.Background(), hist)
	require.NoError(t, err)

	wh, err := db.GetHistory(context.Background(), "999999")
	require.NoError(t, err)
	require.NotEmpty(t, wh)

	err = db.UpdateHistory(context.Background(), history.History{ID: wh[0].ID, Notes: "🙃", AthleteID: "999999"})
	require.NoError(t, err)

	err = db.DeleteHistory(context.Background(), wh[0].ID)
	require.NoError(t, err)
}

func TestDeleteHistory(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	hist := history.History{
		Date:      "01-01-2000",
		Duration:  "99:00:00",
		Notes:     "999999",
		PlanID:    "999999",
		AthleteID: "999999",
	}

	err = db.CreateHistory(context.Background(), hist)
	require.NoError(t, err)

	wh, err := db.GetHistory(context.Background(), "999999")
	require.NoError(t, err)
	require.NotEmpty(t, wh)

	err = db.DeleteHistory(context.Background(), wh[0].ID)
	require.NoError(t, err)
}
