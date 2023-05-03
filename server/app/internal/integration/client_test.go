package integration

import (
	"context"
	integration "github.com/yoda/app/internal/integration/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

const apiURL = "http://api.github.com"

var (
	mux    *http.ServeMux
	server *httptest.Server
	client integration.ClientWithResponsesInterface
)

func BaseURL(baseURL string) integration.ClientOption {
	return func(c *integration.Client) error {
		c.Server = baseURL
		return nil
	}
}

func setup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client, _ = integration.NewClientWithResponses(server.URL)

	return func() {
		server.Close()
	}
}

func TestGetOrganization(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/organizations", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{
            "id": "f6a89b8b-964d-11ed-852f-7c10c9211780",
            "name": "org1",
            "inn": "",
            "kpp": "",
            "updateAt": "2023-04-27T09:50:33"
        },
        {
            "id": "ecf9c10e-9bdf-11ed-852f-7c10c9211780",
            "name": "org2",
            "inn": "1111111111",
            "kpp": "",
            "updateAt": "2023-04-27T16:14:36"
        }]`))
	})

	resp, err := client.GetOrganizationsWithResponse(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode())
	}
	items := resp.JSON200
	if items == nil {
		t.Fatal("expected items, got nil")
	}
	if len(*items) != 2 {
		t.Errorf("expected 2 organizations, got %d", len(*items))
	}

	item := (*items)[0]
	if item.Id != "f6a89b8b-964d-11ed-852f-7c10c9211780" {
		t.Errorf("expected id 1, got %s", item.Id)
	}
	if item.Name != "org1" {
		t.Errorf("expected name org1, got %s", item.Name)
	}
}
