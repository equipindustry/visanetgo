package client

import (
  "encoding/base64"
  "fmt"
  "github.com/spf13/viper"
  "io"
  "io/ioutil"
  "net/http"
)

func ClientAuthorization() ([]byte, error) {
  url := viper.Get("VISANET_DEV_AUTHORIZATION_API")
  req, err := http.NewRequest("POST", url, nil)
  if err != nil {
    return nil, err
  }
  user := viper.Get("VISANET_DEV_USER")
  pwd := viper.Get("VISANET_DEV_PASSWD")
  encodedCredential := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, pwd)))
  req.Header.Set("Authorization", fmt.Sprintf("Basic %s", encodedCredential))

  c := &http.Client{}
  res, err := c.Do(req)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()

  obj, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return nil, err
  }

  if err != nil {
    err = fmt.Errorf("%v: %s", err, string(obj))
    return nil, err
  }

  if res.StatusCode == 201 {
    return obj, nil
  }

  return nil, err
}

func Client(method string, url string, body io.Reader, token string) ([]byte, error) {
  req, err := http.NewRequest(method, url, body)
  if err != nil {
    return nil, err
  }
  req.Header.Set("Authorization", token)

  c := &http.Client{}
  res, err := c.Do(req)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()

  obj, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return nil, err
  }

  if err != nil {
    err = fmt.Errorf("%v: %s", err, string(obj))
    return nil, err
  }

  if res.StatusCode == 200 {
    return obj, nil
  }

  return nil, err
}
