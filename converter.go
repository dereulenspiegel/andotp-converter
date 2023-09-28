package andotpconverter

import (
	"time"

	"github.com/dereulenspiegel/andotp-converter/andotp"
	"github.com/dereulenspiegel/andotp-converter/twofas"
	"github.com/google/uuid"
)

func FromAndOtpTo2Fas(andOTPServices []*andotp.Service) (*twofas.RemoteBackup, error) {
	updatedAt := uint64(time.Now().UnixMilli())
	twofasBackup := &twofas.RemoteBackup{
		UpdatedAt:      updatedAt,
		SchemaVersion:  twofas.DEFAULT_SCHEMA_VERSION,
		AppVersionCode: twofas.DEFAULT_APP_VERSION_CODE,
		AppOrigin:      twofas.DEAFULT_APP_ORIGIN,
	}

	groupMapping := map[string]twofas.RemoteGroup{}
	defaultSource := twofas.DEFAULT_SOURCE

	for i, andOTPService := range andOTPServices {
		var groupId *string
		if len(andOTPService.Tags) > 0 {
			// Use the first tag as group
			tag := andOTPService.Tags[0]
			if group, exists := groupMapping[tag]; !exists {
				group = twofas.RemoteGroup{
					Name:       tag,
					UpdatedAt:  updatedAt,
					IsExpanded: true,
					ID:         uuid.Must(uuid.NewRandom()).String(),
				}
				groupMapping[tag] = group
				groupId = &group.ID
			} else {
				groupId = &group.ID
			}
		}

		name := andOTPService.Issuer
		if name == "" {
			name = andOTPService.Label
		}
		remoteService := twofas.RemoteService{
			Name:      name,
			Secret:    andOTPService.Secret,
			UpdatedAt: andOTPService.LastUsed,
			Order:     twofas.Order{Position: i},
			Type:      thumbnailToServiceType(andOTPService.Thumbnail),
			GroupId:   groupId,
			Otp: twofas.Otp{
				Label:     &andOTPService.Label,
				Account:   &andOTPService.Label,
				Issuer:    &andOTPService.Issuer,
				Digits:    &andOTPService.Digits,
				Period:    &andOTPService.Period,
				Algorithm: &andOTPService.Algorithm,
				Counter:   &andOTPService.UsedFrequency,
				TokenType: &andOTPService.Type,
				Source:    &defaultSource,
			},
		}

		twofasBackup.Services = append(twofasBackup.Services, &remoteService)
	}
	for _, group := range groupMapping {
		twofasBackup.Groups = append(twofasBackup.Groups, &group)
	}
	return twofasBackup, nil
}

func thumbnailToServiceType(thumbnail string) *twofas.ServiceType {
	for _, st := range twofas.ServiceTypes {
		if string(st) == thumbnail {
			return &st
		}
	}
	unknown := twofas.ServiceType("Unknwown")
	return &unknown
}
