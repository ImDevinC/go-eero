package eero_test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/imdevinc/go-eero"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
)

var networkID string = os.Getenv("EERO_NETWORK_ID")

func getAuthedClient(t *testing.T) *eero.Eero {
	t.Helper()
	client := eero.NewEero()
	client.UserToken = os.Getenv("EERO_USERTOKEN")
	return client
}

func TestLogin(t *testing.T) {
	client := eero.NewEero()
	err := client.Login(os.Getenv("EERO_USERTOKEN"))
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	fmt.Println(client.UserToken)
}

func TestVerifyLogin(t *testing.T) {
	client := getAuthedClient(t)
	err := client.VerifyLogin(os.Getenv("EERO_USERTOKEN"))
	if !assert.NoError(t, err) {
		t.FailNow()
	}
}

func TestGetAccount(t *testing.T) {
	client := getAuthedClient(t)
	account, err := client.GetAccount()
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	expected := os.Getenv("EERO_EXPECTED_ACCOUNT_NAME")
	if account.Name != expected {
		t.Errorf("expected %s, but got %s", expected, account.Name)
		t.FailNow()
	}
}

func TestGetNetwork(t *testing.T) {
	// TODO: Network has a lot of attributes, need to validate some during this test
	client := getAuthedClient(t)
	resp, err := client.GetNetwork(networkID)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	t.Logf("%+v", resp.Data["health"])
}

func TestGetNetworkDevices(t *testing.T) {
	// TODO: Network has a lot of attributes, need to validate some during this test
	client := getAuthedClient(t)
	resp, err := client.GetNetworkDevices(networkID)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	if !assert.GreaterOrEqual(t, len(resp.Data), 1, "expected at least one device") {
		t.FailNow()
	}
}

func TestGetDataBreakdown(t *testing.T) {
	client := getAuthedClient(t)
	loc, err := time.LoadLocation("America/Los_Angeles")
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	resp, err := client.GetDataBreakdown(networkID, time.Date(2024, 2, 25, 8, 0, 0, 0, loc), time.Date(2024, 2, 25, 9, 0, 0, 0, loc), "America/Los_Angeles")
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	json.NewEncoder(log.Writer()).Encode(resp)
}
