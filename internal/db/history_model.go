package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/thefueley/scholar-power-api/internal/history"
)

type HistoryRow struct {
	ID        sql.NullString
	Date      sql.NullString `db:"date"`
	Duration  sql.NullString `db:"duration"`
	Notes     sql.NullString `db:"notes"`
	PlanID    sql.NullString `db:"plan_id"`
	AthleteID sql.NullString `db:"athlete_id"`
}

func (db *Database) CreateHistory(ctx context.Context, hist history.History) error {
	_, err := db.ExecContext(ctx,
		`INSERT INTO history(
			date, 
			duration, 
			notes, 
			plan_id, 
			athlete_id)
		VALUES($1, $2, $3, $4, $5)`,
		hist.Date, hist.Duration, hist.Notes, hist.PlanID, hist.AthleteID,
	)

	if err != nil {
		return fmt.Errorf("could not create exercise: %w", err)
	}
	return nil
}

func (db *Database) GetHistory(ctx context.Context, id string) ([]history.History, error) {
	row, err := db.QueryContext(ctx,
		`SELECT *
		FROM history
		WHERE athlete_id = $1`,
		id,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	foundHistory := make([]history.History, 0)

	for row.Next() {
		var hist history.History
		if err := row.Scan(&hist.ID, &hist.Date, &hist.Duration, &hist.Notes, &hist.PlanID, &hist.AthleteID); err != nil {
			fmt.Printf("model.GetHistoryByUserID: Scan: %v\n", err)
			log.Fatal(err)
		}
		foundHistory = append(foundHistory, hist)
	}

	if err != nil {
		return []history.History{}, fmt.Errorf("could not get workout history: %w", err)
	}

	return foundHistory, nil
}

func (db *Database) UpdateHistory(ctx context.Context, hist history.History) error {
	allCurrentHistory, err := db.GetHistory(ctx, hist.AthleteID)
	var currentHistory history.History
	for _, h := range allCurrentHistory {
		if h.ID == hist.ID {
			currentHistory = h
		}
	}

	finalHistory := reconcileHistory(currentHistory, hist)

	if err != nil {
		return fmt.Errorf("model.UpdateHistory: GetHistoryByUserID: %w", err)
	}

	_, err = db.ExecContext(ctx,
		`UPDATE history 
		SET notes = $1 
		WHERE id = $2`,
		finalHistory.Notes, finalHistory.ID,
	)

	if err != nil {
		return fmt.Errorf("could not update workout history: %w", err)
	}

	return nil
}

func (db *Database) DeleteHistory(ctx context.Context, id string) error {
	_, err := db.ExecContext(ctx,
		`DELETE FROM history
		WHERE id = $1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("could not delete workout history: %w", err)
	}
	return nil
}

func reconcileHistory(current, updated history.History) history.History {
	if updated.Notes == "" {
		updated.Notes = current.Notes
	}
	return updated
}
