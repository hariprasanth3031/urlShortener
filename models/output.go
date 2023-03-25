package models

// All the models related to output are defined
type JSONOutput struct {
	Code int
	Data interface{}
}

type ParsedOutput struct {
	Url string
}
