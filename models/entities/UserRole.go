package entities

import "mcs_service/models/auxiliary"

type UserRole struct {
	auxiliary.BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
}
