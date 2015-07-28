package environment

import (
  "fmt"
  "encoding/json"
  "net/http"
  "io/ioutil"
  "net/url"
)

type EnvironmentVariable struct {
  Id                   int `json:"id"`
  ProductId            int `json:"product_id"`
  ProductEnvironmentId int `json:"product_environment_id"`
  InsertedAt           string `json:"inserted_at"`
  UpdatedAt            string `json:"updated_at"`
  Name                 string `json:"name"`
  Value                string `json:"value"`
}

func GetEnvironmentVariables(productName string, environmentName string, authToken string, baseUrl string) ([]EnvironmentVariable, error) {
  path := fmt.Sprintf("products/%s/environments/%s/variables", productName, environmentName)
  url := buildUrl(baseUrl, path, "coalesced=true")

  return makeRequest(url, authToken)
}

func GetProductEnvironmentVariables(productName string, authToken string, baseUrl string) ([]EnvironmentVariable, error) {
  path := fmt.Sprintf("products/%s/environmental_variables", productName)
  url := buildUrl(baseUrl, path, "coalesced=true")

  return makeRequest(url, authToken)
}

func buildUrl(baseUrl string, path string, query string) (string) {
  urrl, err := url.Parse(baseUrl)

  if err != nil {
    fmt.Printf("Error: Can't parse URL: %s\n", baseUrl)
    return ""
  }

  urrl.Path = path
  urrl.RawQuery = query

  return urrl.String()
}

func makeRequest(url string, authToken string) ([]EnvironmentVariable, error) {
  req, err := http.NewRequest("GET", url, nil)
  req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)

  requests := 0

  // If an error occurs, retry up to five times.
  for err != nil && requests < 5 {
    resp, err = client.Do(req)
    requests += 1
  }

  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)

  var variables []EnvironmentVariable
  json.Unmarshal(body, &variables)

  return variables, nil
}