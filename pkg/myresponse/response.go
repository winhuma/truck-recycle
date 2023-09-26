package myresponse

type MyResponse struct {
	Message interface{} `json:"message"`
	Data    interface{} `json:"result"`
}

type PaginationResponse struct {
	Message     string      `json:"message"`
	Data        interface{} `json:"result"`
	Offset      interface{} `json:"offset"`
	Page        interface{} `json:"page"`
	NextPage    interface{} `json:"next_page"`
	TotalPage   interface{} `json:"total_page"`
	AllItem     interface{} `json:"total_item"`
	NowShowItem interface{} `json:"now_show_item"`
}

func SetResponse(msg string, data ...interface{}) interface{} {
	var setData interface{}
	if len(data) != 0 {
		setData = data[0]
	}
	return MyResponse{
		Message: msg,
		Data:    setData,
	}
}
