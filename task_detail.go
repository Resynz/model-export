package model_export

type TaskDetail struct {
	Id         int64  `json:"id"`
	Sn         string `json:"sn"`
	Action     string `json:"action"`
	Condition  string `json:"condition"`
	ResultPath string `json:"result_path"`
	NotifyUrl  string `json:"notify_url"`
	FileName   string `json:"file_name"`
}

func (s *TaskDetail) GetTableName() string {
	return "task_detail"
}
