package services

// ExternalAPIService defines the contract for external API operations
type ExternalAPIService interface {
	ValidateData(data interface{}) error
	FetchExternalData(id string) (map[string]interface{}, error)
}

// IDGenerator defines the contract for ID generation
type IDGenerator interface {
	GenerateID() string
}
