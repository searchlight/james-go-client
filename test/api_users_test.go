/*
Apache JAMES Web Admin API

Testing UsersAPIService

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

func Test_openapi_UsersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UsersAPIService AddDelegatedUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var baseUser string
		var delegatedUser string

		httpRes, err := apiClient.UsersAPI.AddDelegatedUser(context.Background(), baseUser, delegatedUser).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ChangeUsername", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var oldUser string
		var newUser string

		httpRes, err := apiClient.UsersAPI.ChangeUsername(context.Background(), oldUser, newUser).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService CreateUserIdentity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		httpRes, err := apiClient.UsersAPI.CreateUserIdentity(context.Background(), username).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService DeleteUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		httpRes, err := apiClient.UsersAPI.DeleteUser(context.Background(), username).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ExistsUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		httpRes, err := apiClient.UsersAPI.ExistsUser(context.Background(), username).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ListAllowedFromHeaders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var givenUser string

		resp, httpRes, err := apiClient.UsersAPI.ListAllowedFromHeaders(context.Background(), givenUser).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ListDelegatedUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var baseUser string

		resp, httpRes, err := apiClient.UsersAPI.ListDelegatedUsers(context.Background(), baseUser).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ListUserIdentities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.UsersAPI.ListUserIdentities(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ListUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.ListUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService RemoveAllDelegatedUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var baseUser string

		httpRes, err := apiClient.UsersAPI.RemoveAllDelegatedUsers(context.Background(), baseUser).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService RemoveDelegatedUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var baseUser string
		var delegatedUser string

		httpRes, err := apiClient.UsersAPI.RemoveDelegatedUser(context.Background(), baseUser, delegatedUser).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUserIdentity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string
		var identityId string

		httpRes, err := apiClient.UsersAPI.UpdateUserIdentity(context.Background(), username, identityId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpsertUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		httpRes, err := apiClient.UsersAPI.UpsertUser(context.Background(), username).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}