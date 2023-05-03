package integration

import (
	"context"
	integration "github.com/yoda/app/internal/integration/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client integration.ClientWithResponsesInterface
)

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
		w.Write([]byte(`{
					"count": 2,
					"items": [
						{
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
        }]}`))
	})

	resp, err := client.GetOrganizationsWithResponse(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode())
	}
	result := resp.JSON200
	if result == nil {
		t.Fatal("expected items, got nil")
	}
	if result.Count != 2 {
		t.Errorf("expected 2 organizations, got %d", result.Count)
	}

	item := result.Items[0]
	if item.Id != "f6a89b8b-964d-11ed-852f-7c10c9211780" {
		t.Errorf("expected id 1, got %s", item.Id)
	}
	if item.Name != "org1" {
		t.Errorf("expected name org1, got %s", item.Name)
	}
}
