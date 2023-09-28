package andotp

import (
	"encoding/json"
	"fmt"
	"os"
)

type Service struct {
	Secret        string   `json:"secret"`
	Issuer        string   `json:"issuer"`
	Label         string   `json:"label"`
	Digits        int      `json:"digits"`
	Type          string   `json:"type"`
	Algorithm     string   `json:"algorithm"`
	Thumbnail     string   `json:"thumbnail"`
	LastUsed      uint64   `json:"last_used"`
	UsedFrequency int      `json:"used_frequency"`
	Period        int      `json:"period"`
	Tags          []string `json:"tags"`
}

func Import(filepath string) (services []*Service, err error) {
	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filepath, err)
	}
	services = make([]*Service, 0)
	if err = json.Unmarshal(fileBytes, &services); err != nil {
		return nil, fmt.Errorf("failed to unmarshal file %s: %w", filepath, err)
	}
	return
}
