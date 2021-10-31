package model

type (
	// A TODO expresses ...
	TODO struct {
		Id          int
		Subject     string
		Description string
		Created_at  string
		Updated_at  string
	}

	// A CreateTODORequest expresses ...
	CreateTODORequest struct{}
	// A CreateTODOResponse expresses ...
	CreateTODOResponse struct{}

	// A ReadTODORequest expresses ...
	ReadTODORequest struct{}
	// A ReadTODOResponse expresses ...
	ReadTODOResponse struct{}

	// A UpdateTODORequest expresses ...
	UpdateTODORequest struct{}
	// A UpdateTODOResponse expresses ...
	UpdateTODOResponse struct{}

	// A DeleteTODORequest expresses ...
	DeleteTODORequest struct{}
	// A DeleteTODOResponse expresses ...
	DeleteTODOResponse struct{}
)
