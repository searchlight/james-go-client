/*
Apache JAMES Web Admin API

Testing GlobalQuotaAPIService

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

func Test_openapi_GlobalQuotaAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GlobalQuotaAPIService DeleteGlobalQuotaCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.GlobalQuotaAPI.DeleteGlobalQuotaCount(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalQuotaAPIService DeleteGlobalQuotaSize", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.GlobalQuotaAPI.DeleteGlobalQuotaSize(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalQuotaAPIService GetGlobalQuota", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GlobalQuotaAPI.GetGlobalQuota(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalQuotaAPIService GetGlobalQuotaCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GlobalQuotaAPI.GetGlobalQuotaCount(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalQuotaAPIService GetGlobalQuotaSize", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GlobalQuotaAPI.GetGlobalQuotaSize(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalQuotaAPIService UpdateGlobalQuota", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.GlobalQuotaAPI.UpdateGlobalQuota(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalQuotaAPIService UpdateGlobalQuotaCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.GlobalQuotaAPI.UpdateGlobalQuotaCount(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalQuotaAPIService UpdateGlobalQuotaSize", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.GlobalQuotaAPI.UpdateGlobalQuotaSize(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
