package auth

import (
  "fmt"
  "encoding/json"
  "net/http"
  "io/ioutil"
  "bytes"
)

type AuthTokenResponse struct {
  AccessToken string `json:"access_token"`
  TokenType   string `json:"token_type"`
  ExpiresIn   int    `json:"expires_in"`
  Scope       string `json:"scope"`
}

func GetAuthToken(creds Credentials, token_url string) (resp *AuthTokenResponse, err error) {
  payload, _ := json.Marshal(creds.GetParameters())

  client := &http.Client{}

  requests := 0
  request_body := bytes.NewReader(payload)
  r, e := client.Post(token_url, "application/json", request_body)
  
  // If an error occurs, retry up to five times.
  for e != nil && requests < 5 {
    r, e = client.Post(token_url, "application/json", request_body)
    requests += 1
  }

  if e != nil {
    return nil, e
  }

  if r.StatusCode != 200 {
    error := fmt.Errorf("The call to get auth token returned %d", r.StatusCode)
    return nil, error
  }

  defer r.Body.Close()

  body, _ := ioutil.ReadAll(r.Body)

  resp = &AuthTokenResponse{}

  json.Unmarshal(body, resp)

  return resp, nil
}

