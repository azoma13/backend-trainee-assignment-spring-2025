package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func sendRequest(t *testing.T, method, url string, body interface{}, token string) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	client := &http.Client{}
	return client.Do(req)
}

func TestIntegrationFlow(t *testing.T) {
	baseURL := "http://localhost:8080"

	moderatorToken := getDummyLoginToken(t, baseURL, "moderator")
	employeeToken := getDummyLoginToken(t, baseURL, "employee")

	pvzID := createPVZ(t, baseURL, moderatorToken)

	receptionResp, err := sendRequest(t, http.MethodPost, baseURL+"/receptions", map[string]string{
		"pvzId": pvzID,
	}, employeeToken)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, receptionResp.StatusCode)

	for i := 1; i <= 50; i++ {
		productType := "электроника"
		if i%2 == 0 {
			productType = "одежда"
		}

		productResp, err := sendRequest(t, http.MethodPost, baseURL+"/products", map[string]string{
			"type":  productType,
			"pvzId": pvzID,
		}, employeeToken)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, productResp.StatusCode)
	}

	closeResp, err := sendRequest(t, http.MethodPost, fmt.Sprintf("%s/pvz/%s/close_last_reception", baseURL, pvzID), nil, employeeToken)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, closeResp.StatusCode)

	var closeResult struct {
		Status string `json:"status"`
	}
	err = json.NewDecoder(closeResp.Body).Decode(&closeResult)
	assert.NoError(t, err)
	assert.Equal(t, "close", closeResult.Status)
}

func getDummyLoginToken(t *testing.T, baseURL, role string) string {
	loginResp, err := sendRequest(t, http.MethodPost, baseURL+"/dummyLogin", map[string]string{
		"role": role,
	}, "")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, loginResp.StatusCode)

	var loginResult struct {
		Token string `json:"token"`
	}
	err = json.NewDecoder(loginResp.Body).Decode(&loginResult)
	assert.NoError(t, err)
	return loginResult.Token
}

func createPVZ(t *testing.T, baseURL, token string) string {
	pvzResp, err := sendRequest(t, http.MethodPost, baseURL+"/pvz", map[string]string{
		"city": "Москва",
	}, token)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, pvzResp.StatusCode)

	var pvzResult struct {
		ID string `json:"id"`
	}
	err = json.NewDecoder(pvzResp.Body).Decode(&pvzResult)
	assert.NoError(t, err)
	return pvzResult.ID
}
