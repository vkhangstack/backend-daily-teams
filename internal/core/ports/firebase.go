package ports

import "firebase.google.com/go/auth"

type FirebaseAdapter interface {
	GetUser(uid string) (*auth.UserRecord, error)
}
