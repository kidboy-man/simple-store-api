package datatransfers

type BaseQueryParams struct {
	Limit          int
	Offset         int
	Page           int
	IsOnlyCount    bool
	IsWithoutCount bool
	IsPublic       bool
	Search         string
}

type CustomError struct {
	Code    int    `json:"code"`
	Status  int    `json:"status"` // http status code
	Message string `json:"message"`
}

func (ce *CustomError) Error() string {
	return ce.Message
}
