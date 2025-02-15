package datafs

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gbh007/hgraber-next-agent-core/entities"
	"github.com/gbh007/hgraber-next-agent-core/metrics"
	"github.com/google/uuid"
)

func (s *Storage) Delete(ctx context.Context, fileID uuid.UUID) error {
	startAt := time.Now()
	defer func() {
		metrics.RegisterFSActionTime("delete", time.Since(startAt))
	}()

	filepath := s.filepath(fileID)

	err := os.Remove(filepath)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("local fs: %w", entities.FileNotFoundError)
	}

	if err != nil {
		return fmt.Errorf("local fs: os remove: %w", err)
	}

	return nil
}
