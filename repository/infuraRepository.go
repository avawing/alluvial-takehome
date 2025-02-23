package repository

import (
	"alluvial/models/models"

	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Repository struct {
	Client *http.Client
	APIKey string
	Node   string
}

func NewInfuraRepository(c *http.Client, a string) *Repository {
	return &Repository{
		Client: c,
		APIKey: a,
	}
}

func (r *Repository) GetBalanceByID(c context.Context, id string) (string, error) {
	url := fmt.Sprintf("https://mainnet.infura.io/v3/%s", r.APIKey)
	reqBody, jsonErr := json.Marshal(models.InfuraRequest{
		Jsonrpc: "2.0",
		Method:  "eth_getBalance",
		Params: []string{
			id,
			"latest",
		},
		Id: 1,
	})
	if jsonErr != nil {
		return "", fmt.Errorf("GetBalanceByID: %v+", jsonErr)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("GetBalanceByID: %v+", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, clientErr := r.Client.Do(req)
	if clientErr != nil {
		return "", fmt.Errorf("GetBalanceByID: %v+", clientErr)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", resp.StatusCode)
		return "", fmt.Errorf("GetBalanceByID: Request failed with status code: %v+", resp.StatusCode)
	}

	body, ioerr := io.ReadAll(resp.Body)
	if ioerr != nil {
		return "", fmt.Errorf("GetBalanceByID: %v+", ioerr)
	}

	var response models.InfuraResponse
	unmErr := json.Unmarshal(body, &response)
	if unmErr != nil {
		return "", fmt.Errorf("GetBalanceByID: %v+", unmErr)
	}
	return response.Result, nil
}
