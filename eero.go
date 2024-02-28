package eero

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

const endpoint = "https://api-user.e2ro.com/2.2"

var urlEndpoint, _ = url.Parse("https://api-user.e2ro.com/2.2")

type Eero struct {
	client    *http.Client
	UserToken string
}

func NewEero() *Eero {
	return &Eero{
		client: http.DefaultClient,
	}
}

// Login creates the initial login request to the Eero API.
// The identifier is the phone number or email address associated with the Eero account.
// This must be called before calling VerifyLogin as this sets the token used for the cookie.
func (e *Eero) Login(identifier string) error {
	login := map[string]string{"login": identifier}
	payload, err := json.Marshal(login)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, endpoint+"/login", bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := e.client.Do(req)
	if err != nil {
		return err
	}
	body, err := validateClientResponse(resp)
	if err != nil {
		return err
	}
	loginResp := LoginResponse{}
	json.Unmarshal(body, &loginResp)
	e.UserToken = loginResp.Data.UserToken
	return nil
}

// VerifyLogin verifies the login request to the Eero API.
// It uses the provided code to validate with the session token
// created from calling Login().
func (e *Eero) VerifyLogin(code string) error {
	login := map[string]string{"code": code}
	payload, err := json.Marshal(login)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, endpoint+"/login/verify", bytes.NewReader(payload))
	if err != nil {
		return err
	}
	c := http.Cookie{
		Name:  "s",
		Value: e.UserToken,
	}
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&c)
	resp, err := e.client.Do(req)
	if err != nil {
		return err
	}
	_, err = validateClientResponse(resp)
	if err != nil {
		return err
	}
	if e.client.Jar == nil {
		jar, err := cookiejar.New(nil)
		if err != nil {
			return err
		}
		e.client.Jar = jar
	}
	e.client.Jar.SetCookies(urlEndpoint, []*http.Cookie{&c})
	return nil
}

// RefreshLogin refreshes the UserToken that was used for
// login. This will also set the cookie for the connection.
func (e *Eero) RefreshLogin() error {
	req, err := http.NewRequest(http.MethodGet, endpoint+"/login/refresh", nil)
	if err != nil {
		return err
	}
	resp, err := e.client.Do(req)
	if err != nil {
		return err
	}
	body, err := validateClientResponse(resp)
	if err != nil {
		return err
	}
	loginResp := LoginResponse{}
	json.Unmarshal(body, &loginResp)
	e.UserToken = loginResp.Data.UserToken
	return nil
}

func (e *Eero) GetAccount() (AccountResponseData, error) {
	resp, err := e.authedRequest(http.MethodGet, endpoint+"/account", nil)
	if err != nil {
		return AccountResponseData{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AccountResponseData{}, err
	}
	response := AccountResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return AccountResponseData{}, err
	}
	return response.Data, nil
}

func (e *Eero) GetNetwork(networkID string) (NetworkResponse, error) {
	resp, err := e.authedRequest(http.MethodGet, endpoint+"/networks/"+networkID, nil)
	if err != nil {
		return NetworkResponse{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return NetworkResponse{}, err
	}
	response := NetworkResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return NetworkResponse{}, err
	}
	return response, nil
}

func (e *Eero) GetNetworkDevices(networkID string) (DeviceResponse, error) {
	resp, err := e.authedRequest(http.MethodGet, endpoint+"/networks/"+networkID+"/devices", nil)
	if err != nil {
		return DeviceResponse{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return DeviceResponse{}, err
	}
	response := DeviceResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return DeviceResponse{}, err
	}
	return response, nil
}

func (e *Eero) GetDataBreakdown(networkID string, startTime time.Time, endTime time.Time, tz string) (Data, error) {
	values := url.Values{}
	values.Add("start", startTime.Format("2006-01-02T15:05:05.000Z"))
	values.Add("end", endTime.Format("2006-01-02T15:05:05.000Z"))
	values.Add("timezone", tz)
	encodedParams := values.Encode()

	resp, err := e.authedRequest(http.MethodGet, endpoint+"/networks/"+networkID+"/data_usage/breakdown?"+encodedParams, nil)
	if err != nil {
		return Data{}, err
	}
	var data DataBreakdownResp
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return Data{}, err
	}
	if data.Meta.Code != http.StatusOK {
		return Data{}, fmt.Errorf("(%d) %s", data.Meta.Code, data.Meta.Error)
	}
	return data.Data, nil
}

func (e *Eero) authedRequest(method string, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{
		Name:  "s",
		Value: e.UserToken,
	})
	resp, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func validateClientResponse(resp *http.Response) ([]byte, error) {
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	genericResp := GenericResponse{}
	err = json.Unmarshal(respBody, &genericResp)
	if err != nil {
		return nil, err
	}
	if genericResp.Meta.Code != http.StatusOK && genericResp.Meta.Code != http.StatusCreated {
		return nil, fmt.Errorf("(%d) %s", genericResp.Meta.Code, genericResp.Meta.Error)
	}
	return respBody, nil
}
