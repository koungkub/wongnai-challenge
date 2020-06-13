package model

type (
	// Review request body for editable
	Review struct {
		Data Editable `json:"data"`
	}

	// Editable raw data for editable
	Editable struct {
		Comment string `json:"comment"`
	}
)
