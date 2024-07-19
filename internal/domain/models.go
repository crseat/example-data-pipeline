package domain

type PostData struct {
	IPAddress    string                 `json:"ip_address" validate:"required,ipv4"`
	UserAgent    string                 `json:"user_agent" validate:"required"`
	ReferringURL string                 `json:"referring_url" validate:"required,url"`
	AdvertiserID string                 `json:"advertiser_id" validate:"required"`
	Metadata     map[string]interface{} `json:"metadata" validate:"required"`
}
