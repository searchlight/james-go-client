/*
Apache JAMES Web Admin API

Testing AddressGroupAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/searchlight/james-go-client"
)

func Test_openapi_AddressGroupAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AddressGroupAPIService AddMember", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupAddress string

		httpRes, err := apiClient.AddressGroupAPI.AddMember(context.Background(), groupAddress).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AddressGroupAPIService ListGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AddressGroupAPI.ListGroups(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AddressGroupAPIService ListMembers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupAddress string

		resp, httpRes, err := apiClient.AddressGroupAPI.ListMembers(context.Background(), groupAddress).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AddressGroupAPIService RemoveMember", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupAddress string

		httpRes, err := apiClient.AddressGroupAPI.RemoveMember(context.Background(), groupAddress).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}