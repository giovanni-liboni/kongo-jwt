package kongojwt

import (
	"errors"
	"net/http"
)

type KongoJWT struct {
	Server string
}

// GetUserFromToken returns JWT token from an username and a custom ID.
func (kongojwt *KongoJWT) GetToken(username string, customID string) (string, error) {
	data := KongData{Username: username, CustomID: customID, Server: kongojwt.Server}

	status, err := data.getJWTCredentials()
	if err != nil {
		return "", err
	} else {
		if status == http.StatusNotFound {
			// Trying to create a new customer
			err = data.createCustomer()
			if err != nil {
				return "", err
			}
			err = data.createJWTCredentials()
			if err != nil {
				return "", err
			}
		} else if status == http.StatusFound || status == http.StatusOK {
			// Select default token
			err = data.setDefaultJWTResult()
			if err != nil {
				return "", err
			}
		}
	}
	// Returns token
	return data.Token, nil
}

func New(server string) (*KongoJWT, error) {
	if server == "" {
		return nil, errors.New("Setup a valid server string")
	} else {
		return &KongoJWT{Server: server}, nil
	}
}
