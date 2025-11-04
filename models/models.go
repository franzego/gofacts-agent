package models

type FactResponse struct {
	ID          int    `json:"id"`
	Fact        string `json:"fact"`
	Explanation string `json:"explanation"`
}

type HistoryResponse struct {
	Year  int    `json:"year"`
	Event string `json:"event"`
}

type SyntaxRespone struct {
	Response string `json:"response"`
}

type TipsResponse struct {
	Tips string `json:"tips"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type TelexRequest struct {
	Event string `json:"event"`
	Data  struct {
		Content string `json:"content"`
	} `json:"data"`
}

type TelexResponse struct {
	EventName string `json:"event_name"`
	Username  string `json:"username"`
	Message   string `json:"message"`
}
