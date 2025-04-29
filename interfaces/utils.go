package interfaces

// Update Cond ...
type UpdateCond struct {
	Ids  string
	Vals string
}

// Multy Update Cond ...
type MultyUpdateCond struct {
	Ids  []string
	Vals []string
}

type ArrVals struct {
	Vals []string
}

type StandardResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type StandardResponsePagination struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Info    interface{} `json:"pagination_info"`
	Data    interface{} `json:"data,omitempty"`
}

type ResInfoPagination struct {
	TotalData int64 `json:"total_data"`
	Page      int   `json:"curent_page"`
	TotalPage int   `json:"total_page"`
	PerPage   int   `json:"limit_per_page"`
}
