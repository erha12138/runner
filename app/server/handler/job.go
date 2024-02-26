package handler

import (
	"context"
	"time"

	"github.com/cox96de/runner/log"
	"github.com/samber/lo"

	"github.com/cox96de/runner/app/server/dispatch"
	"github.com/cox96de/runner/db"
	"github.com/pkg/errors"

	"github.com/cox96de/runner/api"
	"github.com/cox96de/runner/lib"
)

func (h *Handler) UpdateJobExecution(ctx context.Context, request *api.UpdateJobExecutionRequest) (*api.UpdateJobExecutionResponse, error) {
	logger := log.ExtractLogger(ctx).WithFields(log.Fields{
		"job_id":           request.JobID,
		"job_execution_id": request.ID,
	})
	lock, err := h.locker.Lock(ctx, lib.BuildJobExecutionLockKey(request.ID), "update_job_execution", time.Second)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to lock job execution '%d'", request.ID)
	}
	// May be delay and retry ?
	if !lock {
		return nil, errors.Errorf("job execution '%d' is locked", request.ID)
	}
	defer func() {
		_, _ = h.locker.Unlock(ctx, lib.BuildJobExecutionLockKey(request.ID))
	}()
	jobExecution, err := h.db.GetJobExecution(ctx, request.ID)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to get job execution '%d'", request.ID)
	}
	if request.Status != nil {
		logger.Infof("update job execution status from %s to '%s'", jobExecution.Status, *request.Status)
		if !dispatch.CheckStatus(jobExecution.Status, *request.Status) {
			return nil, errors.Errorf("invalid status transition from '%s' to '%s'", jobExecution.Status, *request.Status)
		}
		jobExecution.Status = *request.Status
		updateJobExecutionOption := &db.UpdateJobExecutionOption{
			ID:     jobExecution.JobID,
			Status: request.Status,
		}
		switch {
		case *request.Status == api.StatusPreparing:
			// TODO: add preparing at.
		case *request.Status == api.StatusRunning:
			updateJobExecutionOption.StartedAt = lo.ToPtr(time.Now())
		case (*request).Status.IsCompleted():
			updateJobExecutionOption.CompletedAt = lo.ToPtr(time.Now())
		}
		err := h.db.UpdateJobExecution(ctx, updateJobExecutionOption)
		if err != nil {
			return nil, errors.WithMessagef(err, "failed to update job execution '%d'", request.ID)
		}
	}
	return &api.UpdateJobExecutionResponse{
		Job: db.PackJobExecution(jobExecution),
	}, nil
}