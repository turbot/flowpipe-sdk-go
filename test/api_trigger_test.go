/*
Flowpipe

Testing TriggerApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package flowpipeapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/turbot/flowpipe-sdk-go"
)

func Test_flowpipeapi_TriggerApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TriggerApiService Get", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var triggerName string

		resp, httpRes, err := apiClient.TriggerApi.Get(context.Background(), triggerName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TriggerApiService List", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TriggerApi.List(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
