package model_export

type Task struct {
	Id          int64            `json:"id"`
	Sn          string           `json:"sn"`
	Status      ExportTaskStatus `json:"status"`
	Creator     int64            `json:"creator"`
	CreateTime  int64            `json:"create_time"`
	ExecuteTime int64            `json:"execute_time"`
	FinishTime  int64            `json:"finish_time"`
	CostTime    int64            `json:"cost_time"`
	UpdateTime  int64            `json:"update_time"`
}

func (s *Task) GetTableName() string {
	return "task"
}
