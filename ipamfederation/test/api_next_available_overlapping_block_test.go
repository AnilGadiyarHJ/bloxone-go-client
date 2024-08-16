/*
IPAM Federation API

Testing NextAvailableOverlappingBlockAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ipamfederation

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/bloxone-go-client/ipamfederation"
)

func TestNextAvailableOverlappingBlockAPIService(t *testing.T) {

	apiClient := ipamfederation.NewAPIClient()

	t.Run("Test NextAvailableOverlappingBlockAPIService ListNextAvailableOverlappingBlocks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.NextAvailableOverlappingBlockAPI.ListNextAvailableOverlappingBlocks(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}