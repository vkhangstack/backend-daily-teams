package model

type User struct {
	Username   string `json:"username" db:"username"`
	Password   string `json:"password" db:"password"`
	Email      string `json:"email" db:"email,omitempty"`
	FirstName  string `json:"firstName,omitempty" db:"first_name"`
	LastName   string `json:"lastName,omitempty" db:"last_ame"`
	Phone      string `json:"phone,omitempty" db:"phone"`
	ProviderId string `json:"providerId,omitempty" db:"provider_id"`
	AvatarURL  string `json:"avatarUrl,omitempty" db:"avatar_url"`
	Address    string `json:"address,omitempty" db:"address"`
	Status     int8   `json:"status" db:"status"`
	//UID        string `json:"uid" db:"uid,omitempty"`
	//ID         uint64 `json:"id" db:"id"`
	*SqlModel
}

func (User) TableName() string {
	return "users"
}

type UserUpdate struct {
	Password   string `json:"password,omitempty" db:"password"`
	FirstName  string `json:"firstName,omitempty" db:"first_name"`
	LastName   string `json:"lastName,omitempty" db:"last_ame"`
	Phone      string `json:"phone,omitempty" db:"phone"`
	ProviderId string `json:"providerId,omitempty" db:"provider_id"`
	AvatarURL  string `json:"avatarUrl,omitempty" db:"avatar_url"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

type UserLoginSocial struct {
	Email     string `json:"email" db:"email"`
	UID       string `json:"uid" db:"uid"`
	Username  string `json:"username,omitempty" db:"username"`
	FirstName string `json:"firstName,omitempty" db:"first_name"`
	LastName  string `json:"lastName,omitempty" db:"last_name"`
	Phone     string `json:"phone,omitempty" db:"phone"`
	AvatarURL string `json:"avatarUrl,omitempty" db:"avatar_url"`
	Address   string `json:"address,omitempty" db:"address"`
}

func (UserLoginSocial) TableName() string {
	return User{}.TableName()
}
