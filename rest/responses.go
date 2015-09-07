package rest

type ListResp struct {
	Meta    ListRespMeta  `json:"meta"`
	Results []interface{} `json:"results"`
}

type ListRespMeta struct {
	Count int `json:"count"` // how many results are being returned
}

type PagedListResp struct {
	Meta    PagedListRespMeta `json:"meta"`
	Results []interface{}     `json:"results"`
}

type PagedListRespMeta struct {
	Count   int `json:"count"`   // how many results are being returned
	PerPage int `json:"perpage"` // how many results per page are being requested
	Pages   int `json:"pages"`   // how many total pages of results
	Results int `json:"results"` // how many results in total to be paged
	Page    int `json:"page"`    // the current page
}

type ErrResp struct {
	Meta ErrRespMeta `json:"meta"`
}

type ErrRespMeta struct {
	ErrCode    int    `json:"error_code"`    // an integer code that described the error that occured
	ErrMessage string `json:"error_message"` // a general error message for the error
}

type FieldErrResp struct {
	Meta FieldErrRespMeta `json:"meta"`
}

type FieldErrRespMeta struct {
	ErrCode    int        `json:"error_code"`    // an integer code that described the error that occured
	ErrMessage string     `json:"error_message"` // a general error message for the error
	ErrFields  []ErrField `json:"error_fields"`  // a list of the fields that generated the error(s)
}

type ErrField struct {
	Field string           `json:"field"`  // the name of the field on which the error occured
	Errs  []ErrFieldObject `json:"errors"` // array of the errors that occurred on the field
}

type ErrFieldObject struct {
	Code    int    `json:"error_code"`    // the error code that occurred
	Message string `json:"error_message"` // the description of the error that occured
}

type CreatedResp struct {
	Meta CreatedRespMeta `json:"meta"`
}

type CreatedRespMeta struct {
	Id int `json:"id"` // the created id for the new object
}
