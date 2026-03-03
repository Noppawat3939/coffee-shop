package dto

import "time"

type GetAuditLogRequest struct {
	ID        *uint      `form:"id"`
	Action    *string    `form:"action"`
	Entity    *string    `form:"entity"`
	StartDate *time.Time `form:"start_date" time_format:"2006-01-02"`
	EndDate   *time.Time `form:"end_date" time_format:"2006-01-02"`
}
