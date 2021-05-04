package model_export

type ExportTaskStatus uint8

const (
	ExportTaskStatusUnknown ExportTaskStatus = iota
	ExportTaskStatusPending
	ExportTaskStatusProcessing
	ExportTaskStatusFailed
	ExportTaskStatusSuccess
	ExportTaskStatusCanceled
)

type ExportTaskLogType uint8

const (
	ExportTaskLogTypeUnknown ExportTaskLogType = iota
	ExportTaskLogTypeCreate
	ExportTaskLogTypeExecuteStart
	ExportTaskLogTypeExecuteFailed
	ExportTaskLogTypeExecuteSuccess
	ExportTaskLogTypeExecuteCanceled
)

const System int64 = -1
