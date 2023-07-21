package types

type IdExistingStruct struct {
	Id string `json:"id" validate:"regexp=^[1-9]*$"`
}
