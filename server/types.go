package server

const (
	ServerName = "settledb"
)

type Status struct {
	Name string
}

type SetRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//type SetResponse struct {
//	Error error `json:"error"`
//}
