package datatransfers

type ListQueryParams struct {
	Limit          int
	Offset         int
	Page           int
	IsOnlyCount    bool
	IsWithoutCount bool
}
