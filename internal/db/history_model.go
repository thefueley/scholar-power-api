package db

import (
	"context"
	"database/sql"
	"errors"
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
	result, err := db.ExecContext(ctx,
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
		fmt.Printf("model.CreateHistory: ExecContext: %v\n", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("model.CreateHistory: RowsAffected: %v\n", err)
	}

	if rows != 1 {
		return errors.New("could not create workout history")
	}

	return nil
}

func (db *Database) GetHistory(ctx context.Context, uid string) ([]history.History, error) {
	row, err := db.QueryContext(ctx,
		`SELECT *
		FROM history
		WHERE athlete_id = $1`,
		uid,
	)

	if err != nil {
		fmt.Printf("model.GetHistory: QueryContext: %v\n", err)
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

	if len(foundHistory) == 0 {
		return []history.History{}, history.ErrHistoryNotFound
	}

	return foundHistory, nil
}

func (db *Database) UpdateHistory(ctx context.Context, hist history.History) error {
	fmt.Printf("got athlete id: %s\n", hist.AthleteID)
	allCurrentHistory, err := db.GetHistory(ctx, hist.AthleteID)

	if err != nil {
		return fmt.Errorf("model.UpdateHistory: could not get current workout history")
	}

	var currentHistory history.History
	for _, h := range allCurrentHistory {
		if h.ID == hist.ID {
			currentHistory = h
		}
	}

	finalHistory := reconcileHistory(currentHistory, hist)

	result, err := db.ExecContext(ctx,
		`UPDATE history 
		SET notes = $1 
		WHERE id = $2`,
		finalHistory.Notes, finalHistory.ID,
	)

	if err != nil {
		fmt.Printf("model.UpdateHistory: ExecContext: %v\n", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("model.UpdateHistory: RowsAffected: %v\n", err)
	}

	if rows != 1 {
		return errors.New("could not update workout history")
	}

	return nil
}

func (db *Database) DeleteHistory(ctx context.Context, id string) error {

	result, err := db.ExecContext(ctx,
		`DELETE FROM history
		WHERE id = $1`,
		id,
	)

	if err != nil {
		fmt.Printf("model.DeleteHistory: ExecContext: %v\n", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("model.DeleteHistory: RowsAffected: %v\n", err)
	}

	if rows != 1 {
		return errors.New("could not delete workout history")
	}

	return nil
}

func reconcileHistory(current, updated history.History) history.History {
	if updated.Notes == "" {
		updated.Notes = current.Notes
	}
	return updated
}
