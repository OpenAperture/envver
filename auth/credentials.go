package auth

type Credentials interface {
  GetParameters() map[string]string
}

type ClientCredentials struct {
  ClientId     string
  ClientSecret string
}

func (c ClientCredentials) GetParameters() map[string]string {
  return map[string]string {
    "grant_type": "client_credentials",
    "client_id": c.ClientId,
    "client_secret": c.ClientSecret,
  }
}

type PasswordCredentials struct {
  Username string
  Password string
}

func (p PasswordCredentials) GetParameters() map[string]string {
  return map[string]string {
    "grant_type": "password",
    "username": p.Username,
    "password": p.Password,
  }
}