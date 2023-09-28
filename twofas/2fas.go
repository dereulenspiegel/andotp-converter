package twofas

const (
	DEFAULT_SCHEMA_VERSION   = 3
	DEFAULT_APP_VERSION_CODE = 4070000
	DEAFULT_APP_VERSION_NAME = "4.7.0"
	DEAFULT_APP_ORIGIN       = "android"

	TOKEN_TYPE_TOTP = "TOTP"

	DEFAULT_SOURCE          = "Link"
	DEFAULT_ICON_COLLECTION = "a5b3fb65-4ec5-43e6-8ec1-49e24ca9e7ad"
)

type Otp struct {
	Link      *string `json:"link,omitempty"`
	Label     *string `json:"label,,omitempty"`
	Account   *string `json:"account,omitempty"`
	Issuer    *string `json:"issuer,omitempty"`
	Digits    *int    `json:"digits,omitempty"`
	Period    *int    `json:"period,omitempty"`
	Algorithm *string `json:"algorith,omitempty"`
	Counter   *int    `json:"counter,omitempty"`
	TokenType *string `json:"tokenType,omitempty"`
	Source    *string `json:"source,omitempty"`
}

type Order struct {
	Position int `json:"position"`
}

type Badge struct {
	Color Tint `json:"color"`
}

type Icon struct {
	Selected       string          `json:"selected"` // One of Brand, Label, IconCollection
	Brand          *Brand          `json:"brand,omitempty"`
	Label          *Label          `json:"label,omitempty"`
	IconCollection *IconCollection `json:"iconCollection,omitempty"`
}

type Brand struct {
	ID *string // Is a ServiceType
}

type IconCollection struct {
	ID string `json:"id"`
}

type Label struct {
	Text            string `json:"text"`
	Backgroundcolor Tint   `json:"backgroundColor"`
}

type RemoteService struct {
	Name      string       `json:"name"`
	Secret    string       `json:"secret"`
	UpdatedAt uint64       `json:"updatedAt"`
	Type      *ServiceType `json:"type,omitempty"` //TODO import enum ServiceType somehow
	Otp       Otp          `json:"otp"`
	Order     Order        `json:"order"`
	Badge     *Badge       `json:"badge,omitempty"`
	Icon      *Icon        `json:"icon,omitempty"`
	GroupId   *string      `json:"groupId,omitempty"`
}

type RemoteGroup struct {
	ID         string `json:"id"` // Should be a UUIDv4
	Name       string `json:"name"`
	IsExpanded bool   `json:"isExpanded"`
	UpdatedAt  uint64 `json:"updatedAt"`
}

type RemoteBackup struct {
	Services          []*RemoteService `json:"services"`
	UpdatedAt         uint64           `json:"updatedAt"`
	SchemaVersion     int              `json:"schemaVersion"`  //Currently at 3
	AppVersionCode    int              `json:"appVersionCode"` //TODO find out a valid value
	AppOrigin         string           `json:"appOrigin"`      //defaults to android
	Groups            []*RemoteGroup   `json:"groups"`
	Account           *string          `json:"account,omitempty"`
	ServicesEncrypted *string          `json:"servicesEncrypted,omitempty"`
	Reference         *string          `json:"reference,omitempty"`
}
