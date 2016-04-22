package kongojwt

import "net/http"

// GetUserFromToken returns JWT token from an username and a custom ID.
func GetToken(username string, customID string) (string, error) {
	data := KongData{Username: username, CustomID: customID}

	status, err := data.GetJWTCredentials()
	if err != nil {
		return "", err
	} else {
		if status == http.StatusNotFound {
			// Trying to create a new customer
			err = data.CreateCustomer()
			if err != nil {
				return "", err
			}
			err = data.CreateJWTCredentials()
			if err != nil {
				return "", err
			}
		} else if status == http.StatusFound || status == http.StatusOK {
			// Select default token
			err = data.SetDefaultJWTResult()
			if err != nil {
				return "", err
			}
		}
	}
	// Returns token
	return data.Token, nil
}
