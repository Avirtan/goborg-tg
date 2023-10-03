package method_dto

// https://core.telegram.org/bots/api#setpassportdataerrors
type PassportElementError struct {
	Source      string   `json:"source"`
	Type        string   `json:"type"`
	FieldName   string   `json:"field_name,omitempty"`
	DataHash    string   `json:"data_hash,omitempty"`
	Message     string   `json:"message"`
	FileHash    string   `json:"file_hash,omitempty"`
	FileHashes  []string `json:"file_hashes,omitempty"`
	ElementHash string   `json:"element_hash,omitempty"`
}

func NewPassportElementError(source, typeStr, message string) PassportElementError {
	return PassportElementError{
		Source:  source,
		Type:    typeStr,
		Message: message,
	}
}

func (p *PassportElementError) SetFieldName(fieldName string) {
	p.FieldName = fieldName
}

func (p *PassportElementError) SetDataHash(dataHash string) {
	p.DataHash = dataHash
}

func (p *PassportElementError) SetFileHash(fileHash string) {
	p.FileHash = fileHash
}

func (p *PassportElementError) SetFileHashes(fileHashes []string) {
	p.FileHashes = fileHashes
}

func (p *PassportElementError) SetElementHash(elementHash string) {
	p.ElementHash = elementHash
}
