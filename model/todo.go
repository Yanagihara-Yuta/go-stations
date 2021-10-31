package model

type (
	// A TODO expresses ...
	TODO struct {
		id          int
		subject     string
		description string
		created_at  string
		updated_at  string
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
