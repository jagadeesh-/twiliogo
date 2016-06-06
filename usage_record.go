package twiliogo



type UsageRecord struct {
    Category        string `json:"category"`
    Description     string `json:"description"`
    AccountSid      string `json:"account_sid"`
    StartDate       string `json:"start_date"`
    EndDate         string `json:"end_date"`
    Usage           string  `json:"usage"`
    UsageUnit       string  `json:"usage_unit"`
    Count           string  `json:"count"`
    CountUnit       string  `json:"count_unit"`
    Price           string  `json:"price"`
    PriceUnit       string  `json:"price_unit"`
    Uri             string  `json:"uri"`
    SubresourceUris string  `json:"sub_resource_uris"` 
}

