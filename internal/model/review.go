package model

type (
	Review struct {
		Data Editable `json:"data"`
	}

	Editable struct {
		Comment string `json:"comment"`
	}
)
