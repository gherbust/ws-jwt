package domain

type User struct {
	Name      string `json:"name"`
	Password  string `json:"password,omitempty"`
	Role      string `json:"role"`
	ClientKey string `json:"client_key,omitempty"`
}
