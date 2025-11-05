package roleresponse

import "github.com/google/uuid"

type RoleResponse struct {
	ID 			uuid.UUID 	`json:"id"`
	Name		string		`json:"name"`
	CreatedAt 	string		`json:"created_at"`
	UpdatedAt 	string		`json:"updated_at"`
}