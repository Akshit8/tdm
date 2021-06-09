package postgresql_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Akshit8/tdm/internal"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	t.Run("Create: OK", func(t *testing.T) {
		t.Parallel()

		task, err := repo.Create(context.Background(), "test", internal.PriorityMedium, internal.Dates{})
		require.NoError(t, err)
		require.NotEmpty(t, task.ID)
	})

	t.Run(("Create: ERR Priority"), func(t *testing.T) {
		t.Parallel()

		task, err := repo.Create(
			context.Background(),
			"invalid priority",
			internal.Priority(-1),
			internal.Dates{},
		)
		require.Error(t, err)
		require.Equal(t, internal.Task{}, task)

		var ierr *internal.Error
		require.ErrorAs(t, err, &ierr)
		require.Equal(t, internal.ErrorCodeUnknown, ierr.Code())
	})
}

func TestFind(t *testing.T) {
	t.Parallel()

	t.Run("Find: OK", func(t *testing.T) {
		t.Parallel()

		newTask, err := repo.Create(context.Background(), "test", internal.PriorityLow, internal.Dates{})
		require.NoError(t, err)

		retrievedTask, err := repo.Find(context.Background(), newTask.ID)
		require.NoError(t, err)

		if !cmp.Equal(newTask, retrievedTask) {
			t.Fatalf("expected result does not match: %s", cmp.Diff(newTask, retrievedTask))
		}
	})

	t.Run("Find: ERR invalid UUID", func(t *testing.T) {
		t.Parallel()

		_, err := repo.Find(context.Background(), "x")
		require.Error(t, err)

		var ierr *internal.Error
		require.ErrorAs(t, err, &ierr)
		require.Equal(t, internal.ErrorCodeInvalidArgument, ierr.Code())
	})

	t.Run("Find: ERR not found", func(t *testing.T) {
		t.Parallel()

		_, err := repo.Find(context.Background(), uuid.NewString())
		require.Error(t, err)

		var ierr *internal.Error
		require.ErrorAs(t, err, &ierr)
		require.Equal(t, internal.ErrorCodeNotFound, ierr.Code())
	})
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	t.Run("Update: OK", func(t *testing.T) {
		t.Parallel()

		newTask, err := repo.Create(context.Background(), "test", internal.PriorityLow, internal.Dates{})
		require.NoError(t, err)

		newTask.Description = "updated test description"
		newTask.Dates.Due = time.Now().UTC()
		newTask.Priority = internal.PriorityHigh

		err = repo.Update(
			context.Background(),
			newTask.ID,
			newTask.Description,
			newTask.Priority,
			newTask.Dates,
			newTask.IsDone,
		)
		require.NoError(t, err)

		retrievedTask, err := repo.Find(context.Background(), newTask.ID)
		require.NoError(t, err)

		if !cmp.Equal(newTask, retrievedTask) {
			t.Fatalf("expected result does not match: %s", cmp.Diff(newTask, retrievedTask))
		}
	})

	t.Run("Update: ERR invalid UUID", func(t *testing.T) {
		t.Parallel()

		err := repo.Update(
			context.Background(),
			"x",
			"",
			internal.PriorityLow,
			internal.Dates{},
			false,
		)
		require.Error(t, err)

		var ierr *internal.Error
		require.ErrorAs(t, err, &ierr)
		require.Equal(t, internal.ErrorCodeInvalidArgument, ierr.Code())
	})

	t.Run("Update: ERR invalid Priority", func(t *testing.T) {
		t.Parallel()

		newTask, err := repo.Create(context.Background(), "test", internal.PriorityLow, internal.Dates{})
		require.NoError(t, err)

		err = repo.Update(
			context.Background(),
			newTask.ID,
			"",
			internal.Priority(-1),
			internal.Dates{},
			false,
		)
		require.Error(t, err)

		var ierr *internal.Error
		require.ErrorAs(t, err, &ierr)
		require.Equal(t, internal.ErrorCodeUnknown, ierr.Code())
	})

	t.Run("Update: ERR not found", func(t *testing.T) {
		t.Parallel()

		err := repo.Update(
			context.Background(),
			uuid.NewString(),
			"",
			internal.PriorityMedium,
			internal.Dates{},
			false,
		)
		require.Error(t, err)

		var ierr *internal.Error
		require.ErrorAs(t, err, &ierr)
		require.Equal(t, internal.ErrorCodeNotFound, ierr.Code())
	})
}

func TestDelete(t *testing.T) {
	t.Parallel()

	t.Run("Delete: OK", func(t *testing.T) {
		t.Parallel()

		newTask, err := repo.Create(context.Background(), "test", internal.PriorityLow, internal.Dates{})
		require.NoError(t, err)

		err = repo.Delete(context.Background(), newTask.ID)
		require.NoError(t, err)

		_, err = repo.Find(context.Background(), newTask.ID)
		require.ErrorIs(t, err, sql.ErrNoRows)
	})

	t.Run("Delete: ERR uuid", func(t *testing.T) {
		t.Parallel()

		err := repo.Delete(context.Background(), "x")

		require.Error(t, err)

		var ierr *internal.Error
		require.ErrorAs(t, err, &ierr)
		require.Equal(t, internal.ErrorCodeInvalidArgument, ierr.Code())
	})

	t.Run("Delete: ERR not found", func(t *testing.T) {
		t.Parallel()

		err := repo.Delete(context.Background(), uuid.NewString())

		require.Error(t, err)

		var ierr *internal.Error
		require.ErrorAs(t, err, &ierr)
		require.Equal(t, internal.ErrorCodeNotFound, ierr.Code())
	})
}
