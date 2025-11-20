package payload

import "github.com/google/uuid"

type ImageUploadPayload struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Type      string    `json:"type"`
	FileBytes []byte    `json:"file_bytes,omitempty"`
	Folder    string    `json:"folder"`
	Filename  string    `json:"filename"`
}

// Implementasi interface uploadable
func (p *ImageUploadPayload) GetFileBytes() []byte { return p.FileBytes }
func (p *ImageUploadPayload) GetFolder() string    { return p.Folder }
func (p *ImageUploadPayload) GetFilename() string  { return p.Filename }
func (p *ImageUploadPayload) GetType() string      { return p.Type }