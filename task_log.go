package model_export

type TaskLog struct {
	Id         int64             `json:"id"`
	Sn         string            `json:"sn"`
	Type       ExportTaskLogType `json:"type"`
	Operator   int64             `json:"operator"`
	Remark     string            `json:"remark"`
	CreateTime int64             `json:"create_time"`
}

func (s *TaskLog) GetTableName() string {
	return "task_log"
}
