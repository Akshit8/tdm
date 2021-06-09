// Package postgresql implements queries to interact with postgres db
package postgresql

import (
	"database/sql"
	"time"

	"github.com/Akshit8/tdm/internal"
)

//go:generate sqlc generate

func convertPriority(p Priority) (internal.Priority, error) {
	switch p {
	case PriorityNone:
		return internal.PriorityNone, nil
	case PriorityLow:
		return internal.PriorityLow, nil
	case PriorityMedium:
		return internal.PriorityMedium, nil
	case PriorityHigh:
		return internal.PriorityHigh, nil
	}

	return internal.Priority(-1), internal.NewErrorf(internal.ErrorCodeInvalidArgument, "unknow value")
}

func newNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: !t.IsZero(),
	}
}

func newPriority(p internal.Priority) Priority {
	switch p {
	case internal.PriorityNone:
		return PriorityNone
	case internal.PriorityLow:
		return PriorityLow
	case internal.PriorityMedium:
		return PriorityMedium
	case internal.PriorityHigh:
		return PriorityHigh
	}

	// since for priority we are using enum, postgres will throw error for below string.
	return "invalid"
}
