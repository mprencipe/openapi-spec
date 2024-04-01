package parser

import (
	json "encoding/json"
	"errors"
	"fmt"

	"github.com/mprencipe/openapi-spec/openapi"
	"github.com/mprencipe/openapi-spec/swagger"
)

// Implemented by all response types
type ApiResponse interface {
	IsSwaggerResponse() bool
	IsOpenApiResponse() bool
}

func isSwaggerResponse(resp *map[string]interface{}) bool {
	return (*resp)["swagger"] != nil
}

func isOpenApiResponse(resp *map[string]interface{}) bool {
	return (*resp)["openapi"] != nil
}

func swaggerResponseOrError(jsonBytes []byte, resp *map[string]interface{}) (ApiResponse, error) {
	swaggerResponse := swagger.SwaggerV2Response{}
	err := json.Unmarshal(jsonBytes, &swaggerResponse)
	if err != nil {
		return nil, errors.New("Failed to parse Swagger response")
	}
	return &swaggerResponse, nil
}

func openApiResponseOrError(bytes []byte, resp *map[string]interface{}) (ApiResponse, error) {
	openApiResponse := openapi.OpenAPI310ApiResponse{}
	err := json.Unmarshal(bytes, &openApiResponse)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse OpenAPI response: %w", err)
	}
	return &openApiResponse, nil
}

func ParseApiResponse(bytes []byte) (ApiResponse, error) {
	resp := make(map[string]interface{})
	err := json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, errors.New("Failed to parse API response")
	}
	if isSwaggerResponse(&resp) {
		return swaggerResponseOrError(bytes, &resp)
	} else if isOpenApiResponse(&resp) {
		return openApiResponseOrError(bytes, &resp)
	}
	return nil, errors.New("Unsupported API type, not Swagger or OpenAPI")
}
