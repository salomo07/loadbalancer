package models

type Server struct {
	Address    string `json:"address"`
	Path       string `json:"path"`
	Latency    int64  `json:"latency"`
	LastUpdate string `json:"lastupdate"`
	Online     bool   `json:"online"`
}

type DefaultResponse struct {
	Status   int    `json:"status"`
	Messeges string `json:"messeges"`
}
type UserCDB struct {
	ID       string   `json:"_id"`
	Name     string   `json:"name"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
	Type     string   `json:"type"`
	Database string   `json:"database"`
}
