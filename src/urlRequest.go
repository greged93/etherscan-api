package urlRequest

import (
  "net/url"
  "fmt"
)

type UrlRequest struct {
	BaseUrl string
	Queries map[string]string
}

type ErrEthAddressFormat string

func (e ErrEthAddressFormat) Error() string {
  return fmt.Sprintf("Invalid Ethereum address format, got %s", string(e))
}

func checkEthAddressFormat(eth string) bool {
  return eth[0:2] == "0x" && len(eth) == 42
}

func (ureq UrlRequest) MakeUrl() (*url.URL, error) {
	u, err := url.Parse(ureq.BaseUrl)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	for key, value := range ureq.Queries {
    if key == "address" && !checkEthAddressFormat(value){
      return nil, ErrEthAddressFormat(value)
    }
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()
	return u, nil
}
