package postgresql_test

import (
	"context"
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

		task, err := repo.Create(context.Background(), "invalid priority", internal.Priority(-1), internal.Dates{})
		require.Error(t, err)
		require.Equal(t, task, internal.Task{})
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
	})

	t.Run("Find: ERR not found", func(t *testing.T) {
		t.Parallel()

		_, err := repo.Find(context.Background(), uuid.NewString())
		require.Error(t, err)
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

		err = repo.Update(context.Background(),
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

		err := repo.Update(context.Background(),
			"x",
			"",
			internal.PriorityLow,
			internal.Dates{},
			false,
		)
		require.Error(t, err)
	})

	t.Run("Update: ERR invalid Priority", func(t *testing.T) {
		t.Parallel()

		newTask, err := repo.Create(context.Background(), "test", internal.PriorityLow, internal.Dates{})
		require.NoError(t, err)

		err = repo.Update(context.Background(),
			newTask.ID,
			"",
			internal.Priority(-1),
			internal.Dates{},
			false,
		)
		require.Error(t, err)
	})
}
