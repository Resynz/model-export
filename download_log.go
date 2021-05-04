package model_export

type DownloadLog struct {
	Id         int64  `json:"id"`
	Sn         string `json:"sn"`
	Downloader int64  `json:"downloader"`
	CreateTime int64  `json:"create_time"`
	Ip         string `json:"ip"`
	UserAgent  string `json:"user_agent"`
}

func (s *DownloadLog) GetTableName() string {
	return "download_log"
}
