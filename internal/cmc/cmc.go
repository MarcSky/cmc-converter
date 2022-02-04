package cmc

import (
	"cmcconv/internal"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	apiURL = "https://pro-api.coinmarketcap.com/v1/tools/price-conversion?amount=%s&symbol=%s&convert=%s"
)

type api struct {
	cmcToken string
	client   *http.Client
}

func NewCMC(cmcToken string) internal.ConverterClient {
	return &api{
		cmcToken: cmcToken,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *api) reqConvert(amount, from, to string) (float64, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, fmt.Sprintf(apiURL, amount, from, to), http.NoBody)
	if err != nil {
		return 0.0, err
	}
	req.Header.Set("X-CMC_PRO_API_KEY", c.cmcToken)

	res, err := c.client.Do(req)
	if err != nil {
		return 0.0, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0.0, err
	}

	if res.StatusCode != http.StatusOK {
		return 0.0, fmt.Errorf("bad response status %d %s", res.StatusCode, string(body))
	}

	var result Convert
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0.0, err
	}

	return result.Data.Quotes[to].Price, nil
}

func (c *api) Converter(amount, from, to string) (float64, error) {
	return c.reqConvert(amount, from, to)
}
