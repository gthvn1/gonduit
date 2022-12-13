package gonduit

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/gthvn1/gonduit/core"
	"github.com/gthvn1/gonduit/requests"
	"github.com/gthvn1/gonduit/responses"
	"github.com/gthvn1/gonduit/test/server"
)

const buildableSearchResponseJSON = `{
  "result": {
    "data": [
      {
        "id": 54057,
        "type": "HMBB",
        "phid": "PHID-HMBB-6tceawkrkt55btokp7es",
        "fields": {
          "objectPHID": "PHID-DIFF-7grzvaqb24vaorwgj6f6",
          "containerPHID": "PHID-DREV-ea4xglpfvktonm7cyzmq",
          "buildableStatus": {
            "value": "failed"
          },
          "isManual": false,
          "uri": "https://www.example.com/B54057",
          "dateCreated": 1419993553,
          "dateModified": 1419994281,
          "policy": {
            "view": "users",
            "edit": "users"
          }
        },
        "attachments": {}
      }
    ],
    "maps": {},
    "query": {
      "queryKey": null
    },
    "cursor": {
      "limit": 100,
      "after": null,
      "before": null,
      "order": null
    }
  }
}`

func TestHarbormasterBuildableSearch(t *testing.T) {
	s := server.New()
	defer s.Close()
	s.RegisterCapabilities()
	response := server.ResponseFromJSON(buildableSearchResponseJSON)
	s.RegisterMethod(HarbormasterBuildableSearchMethod, http.StatusOK, response)

	c, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})
	assert.Nil(t, err)
	req := requests.HarbormasterBuildableSearchRequest{
		Constraints: &requests.HarbormasterBuildableSearchConstraints{
			IDs: []int{54057},
		},
	}
	resp, err := c.HarbormasterBuildableSearch(req)
	assert.NoError(t, err)
	want := responses.HarbormasterBuildableSearchResponse{
		Data: []*responses.HarbormasterBuildableSearchResponseItem{
			{
				ResponseObject: responses.ResponseObject{
					ID:   54057,
					Type: "HMBB",
					PHID: "PHID-HMBB-6tceawkrkt55btokp7es",
				},
				Fields: responses.HarbormasterBuildableSearchResponseItemFields{
					ObjectPHID:    "PHID-DIFF-7grzvaqb24vaorwgj6f6",
					ContainerPHID: "PHID-DREV-ea4xglpfvktonm7cyzmq",
					BuildableStatus: responses.BuildableStatus{
						Value: "failed",
					},
					IsManual:     false,
					URI:          "https://www.example.com/B54057",
					DateCreated:  timestamp(1419993553),
					DateModified: timestamp(1419994281),
				},
			},
		},
		Cursor: responses.SearchCursor{
			Limit: 100,
		},
	}
	assert.Equal(t, &want, resp)
}

const buildSearchResponseJSON = `{
  "result": {
	"data": [
	  {
	    "id": 133550439,
	    "type": "HMBD",
	    "phid": "PHID-HMBD-zkyrlqywr4ahneo3pgih",
	    "fields": {
	  	"buildablePHID": "PHID-HMBB-75lynfoojfsgoezq4dg6",
	  	"buildPlanPHID": "PHID-HMCP-ieis77mkeighxvxu3zvi",
	  	"buildStatus": {
	  	  "value": "passed",
	  	  "name": "Passed",
	  	  "color.ansi": "green"
	  	},
	  	"initiatorPHID": "PHID-HRUL-57w2uekbfssb5py6uf57",
	  	"name": "metadata-check-diff for go-code",
	  	"dateCreated": 1623311845,
	  	"dateModified": 1623312245,
	  	"policy": {
	  	  "view": "users",
	  	  "edit": "users"
	  	}
	    },
	    "attachments": {}
	  }
	],
    "maps": {},
    "query": {
      "queryKey": null
    },
    "cursor": {
      "limit": 100,
      "after": null,
      "before": null,
      "order": null
    }
  }
}`

func TestHarbormasterBuildSearch(t *testing.T) {
	s := server.New()
	defer s.Close()
	s.RegisterCapabilities()
	response := server.ResponseFromJSON(buildSearchResponseJSON)
	s.RegisterMethod(HarbormasterBuildSearchMethod, http.StatusOK, response)

	c, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})
	assert.Nil(t, err)
	req := requests.HarbormasterBuildSearchRequest{
		Constraints: &requests.HarbormasterBuildSearchConstraints{
			IDs: []int{133550439},
		},
	}
	resp, err := c.HarbormasterBuildSearch(req)
	assert.NoError(t, err)
	want := responses.HarbormasterBuildSearchResponse{
		Data: []*responses.HarbormasterBuildSearchResponseItem{
			{
				ResponseObject: responses.ResponseObject{
					ID:   133550439,
					Type: "HMBD",
					PHID: "PHID-HMBD-zkyrlqywr4ahneo3pgih",
				},
				Fields: responses.HarbormasterBuildSearchResponseItemFields{
					BuildablePHID: "PHID-HMBB-75lynfoojfsgoezq4dg6",
					BuildPlanPHID: "PHID-HMCP-ieis77mkeighxvxu3zvi",
					BuildStatus: responses.BuildStatus{
						Value: "passed",
					},
					InitiatorPHID: "PHID-HRUL-57w2uekbfssb5py6uf57",
					Name:          "metadata-check-diff for go-code",
					DateCreated:   timestamp(1623311845),
					DateModified:  timestamp(1623312245),
				},
			},
		},
		Cursor: responses.SearchCursor{
			Limit: 100,
		},
	}
	assert.Equal(t, &want, resp)
}
