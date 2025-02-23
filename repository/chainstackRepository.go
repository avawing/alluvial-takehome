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

func NewChainstackRepository(c *http.Client, a string) *Repository {
	return &Repository{
		Client: c,
		APIKey: a,
	}
}

func (r *Repository) GetBalanceByIDCS(c context.Context, id string) (string, error) {
	url := fmt.Sprintf("https://base-mainnet.core.chainstack.com/%s", r.APIKey)
	reqBody, jsonErr := json.Marshal(models.ChainstackRequest{
		Jsonrpc: "2.0",
		Method:  "eth_getBalance",
		Params: []string{
			id,
			"latest",
		},
		Id: 1,
	})
	if jsonErr != nil {
		return "", fmt.Errorf("GetBalanceByIDCS: %v+", jsonErr)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("GetBalanceByIDCS: %v+", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, clientErr := r.Client.Do(req)
	if clientErr != nil {
		return "", fmt.Errorf("GetBalanceByIDCS: %v+", clientErr)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", resp.StatusCode)
		return "", fmt.Errorf("GetBalanceByIDCS: Request failed with status code: %v+", resp.StatusCode)
	}

	body, ioerr := io.ReadAll(resp.Body)
	if ioerr != nil {
		return "", fmt.Errorf("GetBalanceByIDCS: %v+", ioerr)
	}

	var response models.ChainstackResponse
	unmErr := json.Unmarshal(body, &response)
	if unmErr != nil {
		return "", fmt.Errorf("GetBalanceByIDCS: %v+", unmErr)
	}
	return response.Result, nil
}
