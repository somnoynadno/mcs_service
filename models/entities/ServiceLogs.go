package entities

import (
	"mcs_service/models/auxiliary"
)

// persistent logs
type ServiceLogs struct {
	auxiliary.BaseModel
	RealIP    string
	URL       string
	Method    string
	UserAgent string
	Headers   string
}
