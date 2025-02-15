package item

type CreateRequest struct {
	Date        string
	Formula     string
	CategoryID  string
	Description string
}

type UpdateRequest struct {
	ID          string
	Date        string
	Formula     string
	CategoryID  string
	Description string
}

type IndexRequest struct {
	Year  string
	Month string
}
