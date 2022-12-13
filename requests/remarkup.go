package requests

import "github.com/gthvn1/gonduit/constants"

// RemarkupProcessQuery represents a request to project.query.
type RemarkupProcessRequest struct {
	Context  constants.RemarkupProcessContextType `json:"context"`
	Contents []string                             `json:"contents"`
	Request
}
