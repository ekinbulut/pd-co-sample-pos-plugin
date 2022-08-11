package requests

type CatalogImportRequest struct {
	CatalogImportID string `json:"catalogImportId"`
	Status          string `json:"status"`
	Message         string `json:"message"`
	Details         []struct {
		Status           string `json:"status"`
		PosVendorID      string `json:"posVendorId"`
		PlatformVendorID string `json:"platformVendorId"`
		GlobalEntityID   string `json:"globalEntityId"`
	} `json:"details"`
}
