package eero

import "time"

type GenericResponse struct {
	Meta      Meta        `json:"meta"`
	LoginData interface{} `json:"data"`
}

type Meta struct {
	Code       int        `json:"code"`
	ServerTime *time.Time `json:"server_time"`
	Error      string     `json:"error,omitempty"`
}

type Login struct {
	Login string `json:"login"`
}

type LoginResponse struct {
	Meta Meta              `json:"meta"`
	Data LoginResponseData `json:"data"`
}

type LoginResponseData struct {
	UserToken string `json:"user_token"`
}

type VerifyLogin struct {
	Code string `json:"code"`
}

type AccountResponse struct {
	Meta Meta                `json:"meta"`
	Data AccountResponseData `json:"data"`
}

type AccountResponseData struct {
	Name                      string         `json:"name"`
	Phone                     Phone          `json:"phone"`
	Email                     Email          `json:"email"`
	LogID                     string         `json:"log_id"`
	OrganizationID            interface{}    `json:"organization_id"`
	ImageAssets               interface{}    `json:"image_assets"`
	Networks                  Networks       `json:"networks"`
	Auth                      Auth           `json:"auth"`
	Role                      string         `json:"role"`
	IsBetaBugReporterEligible bool           `json:"is_beta_bug_reporter_eligible"`
	CanTransfer               bool           `json:"can_transfer"`
	IsPremiumCapable          bool           `json:"is_premium_capable"`
	PaymentFailed             bool           `json:"payment_failed"`
	PremiumStatus             string         `json:"premium_status"`
	PremiumDetails            PremiumDetails `json:"premium_details"`
	PushSettings              PushSettings   `json:"push_settings"`
	TrustCertificatesEtag     string         `json:"trust_certificates_etag"`
	CanMigrateToAmazonLogin   bool           `json:"can_migrate_to_amazon_login"`
	EeroForBusiness           bool           `json:"eero_for_business"`
	MduProgram                bool           `json:"mdu_program"`
	BusinessDetails           interface{}    `json:"business_details"`
	Consents                  Consents       `json:"consents"`
}

type Phone struct {
	Value          string `json:"value"`
	CountryCode    string `json:"country_code"`
	NationalNumber string `json:"national_number"`
	Verified       bool   `json:"verified"`
}

type Email struct {
	Value    string `json:"value"`
	Verified bool   `json:"verified"`
}

type Networks struct {
	Count int       `json:"count"`
	Data  []Network `json:"data"`
}

type Network struct {
	URL              string      `json:"url"`
	Name             string      `json:"name"`
	Created          time.Time   `json:"created"`
	NicknameLabel    interface{} `json:"nickname_label"`
	AccessExpiresOn  interface{} `json:"access_expires_on"`
	AmazonDirectedID interface{} `json:"amazon_directed_id"`
}

type Auth struct {
	Type       string      `json:"type"`
	ProviderID interface{} `json:"provider_id"`
	ServiceID  interface{} `json:"service_id"`
}

type PremiumDetails struct {
	TrialEnds            interface{} `json:"trial_ends"`
	HasPaymentInfo       bool        `json:"has_payment_info"`
	Tier                 string      `json:"tier"`
	IsIapCustomer        bool        `json:"is_iap_customer"`
	PaymentMethod        interface{} `json:"payment_method"`
	Interval             string      `json:"interval"`
	NextBillingEventDate interface{} `json:"next_billing_event_date"`
}

type PushSettings struct {
	NetworkOffline bool `json:"networkOffline"`
	NodeOffline    bool `json:"nodeOffline"`
}

type Consents struct {
	MarketingEmails MarketingEmails `json:"marketing_emails"`
}

type MarketingEmails struct {
	Consented bool `json:"consented"`
}

type NetworkResponse struct {
	Meta Meta                   `json:"meta"`
	Data map[string]interface{} `json:"data"`
}

type NetworkResponseData struct {
	Speed NetworkSpeed `json:"speed"`
}

type NetworkSpeed struct {
	Status interface{}         `json:"status"`
	Up     NetworkSpeedResults `json:"up"`
	Down   NetworkSpeedResults `json:"down"`
}

type NetworkSpeedResults struct {
	Units string  `json:"units"`
	Value float64 `json:"value"`
}

type DeviceResponse struct {
	Meta Meta                     `json:"meta"`
	Data []map[string]interface{} `json:"data"`
}

type Device struct {
	Mac                      string             `json:"mac"`
	Manufacturer             string             `json:"manufacturer"`
	IP                       string             `json:"ip"`
	IPs                      []string           `json:"ips"`
	IPv6Addresses            []string           `json:"ipv6_addresses"`
	Nickname                 string             `json:"nickname"`
	Hostname                 string             `json:"hostname"`
	Connected                bool               `json:"connected"`
	Wireless                 bool               `json:"wireless"`
	ConnectionType           string             `json:"connection_type"`
	Source                   DeviceSource       `json:"source"`
	LastActive               time.Time          `json:"last_active"`
	FirstActive              time.Time          `json:"first_active"`
	Connectivity             DeviceConnectivity `json:"connectivity"`
	Interface                DeviceInterface    `json:"interface"`
	Usage                    interface{}        `json:"usage"`
	Profile                  interface{}        `json:"profile"`
	DeviceType               string             `json:"device_type"`
	Blacklisted              bool               `json:"blacklisted"`
	IsGuest                  bool               `json:"is_guest"`
	Paused                   bool               `json:"paused"`
	Channel                  int                `json:"channel"`
	Auth                     string             `json:"auth"`
	IsPrivate                bool               `json:"is_private"`
	SecondaryWanDenyAccess   bool               `json:"secondary_wan_deny_access"`
	Ipv4                     interface{}        `json:"ipv4"`
	IsProxiedNode            bool               `json:"is_proxied_node"`
	ManufacturerDeviceTypeID interface{}        `json:"manufacturer_device_type_id"`
	Ssid                     string             `json:"ssid"`
	DisplayName              string             `json:"display_name"`
}

type DeviceSource struct {
	Location      string `json:"location"`
	IsGateway     bool   `json:"is_gateway"`
	Model         string `json:"model"`
	DisplayName   string `json:"display_name"`
	SerialNumber  string `json:"serial_number"`
	IsProxiedNode bool   `json:"is_proxied_node"`
}

type DeviceConnectivity struct {
	RxBitrate      string         `json:"rx_bitrate"`
	Signal         string         `json:"signal"`
	Score          float64        `json:"score"`
	ScoreBars      int            `json:"score_bars"`
	Frequency      int            `json:"frequency"`
	RxRateInfo     DeviceRateInfo `json:"rx_rate_info"`
	TxRateInfo     DeviceRateInfo `json:"tx_rate_info"`
	EthernetStatus interface{}    `json:"ethernet_status"`
}

type DeviceRateInfo struct {
	RateBps       int    `json:"rate_bps"`
	Mcs           int    `json:"mcs"`
	Nss           int    `json:"nss"`
	GuardInterval string `json:"guard_interval"`
	ChannelWidth  string `json:"channel_width"`
	PhyType       string `json:"phy_type"`
}

type DeviceInterface struct {
	Frequency     string `json:"frequency"`
	FrequencyUnit string `json:"frequency_unit"`
}

type DataBreakdownResp struct {
	Meta Meta `json:"meta"`
	Data Data `json:"data"`
}

type Data struct {
	Start      time.Time    `json:"start"`
	End        time.Time    `json:"end"`
	Upload     int64        `json:"upload"`
	Download   int64        `json:"download"`
	Eeros      []EeroDevice `json:"eeros"`
	Devices    []DeviceData `json:"devices"`
	Unprofiled []DeviceData `json:"unprofiled"`
}

type EeroDevice struct {
	ID          int    `json:"id"`
	ModelNumber string `json:"model_number"`
	Serial      string `json:"serial"`
	Location    string `json:"location"`
	Upload      int64  `json:"upload"`
	Download    int64  `json:"download"`
}

type DeviceData struct {
	URL                string      `json:"url"`
	Nickname           string      `json:"nickname"`
	Hostname           string      `json:"hostname"`
	Manufacturer       string      `json:"manufacturer"`
	MAC                string      `json:"mac"`
	Type               string      `json:"device_type"`
	IsPrivate          bool        `json:"is_private"`
	ConnectionType     string      `json:"connection_type"`
	Upload             int64       `json:"upload"`
	Download           int64       `json:"download"`
	DisplayName        string      `json:"display_name"`
	ModelName          string      `json:"model_name"`
	IsProxiedNode      bool        `json:"is_proxied_node"`
	AmazonDeviceTypeID interface{} `json:"amazon_device_type_id"`
}
