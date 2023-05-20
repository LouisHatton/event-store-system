package users

import "firebase.google.com/go/v4/auth"

type User struct {
	Id            string `json:"id"`
	DisplayName   string `json:"displayName"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"emailVerified"`
	TenantID      string `json:"tenantId"`
	PhotoURL      string `json:"photoUrl"`
	ProviderID    string `json:"providerId"`
}

func AuthUserRecordToUser(user *auth.UserRecord) User {
	return User{
		Id:            user.UID,
		DisplayName:   user.DisplayName,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		TenantID:      user.TenantID,
		PhotoURL:      user.PhotoURL,
		ProviderID:    user.ProviderID,
	}
}
