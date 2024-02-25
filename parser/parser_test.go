// generate test for parser

package parser

import (
	"testing"
)

func TestParsingFailure(t *testing.T) {
	bytes := []byte(`}`)
	_, err := ParseApiResponse(bytes)
	if err == nil {
		t.Error("Should fail with parsing failure")
	}
}

func TestUnsupportedSpec(t *testing.T) {
	bytes := []byte(`{"asdf":"0.0.1"}`)
	_, err := ParseApiResponse(bytes)
	if err == nil {
		t.Error("Should fail with unknown API response")
	}
}

// Swagger

func TestSwagger2_0(t *testing.T) {
	bytes := []byte(`{"swagger":"2.0"}`)
	_, err := ParseApiResponse(bytes)
	if err != nil {
		t.Errorf("Error parsing Swagger response: %s", err)
	}
}

func TestMalformedSwaggerResponse(t *testing.T) {
	bytes := []byte(`{"swagger":"2.0", "info":[]}`)
	_, err := ParseApiResponse(bytes)
	if err == nil {
		t.Error("Should fail with malformed Swagger response")
	}
}

// OpenAPI

func TestOpenAPI3_0_x(t *testing.T) {
	bytes := []byte(`{"openapi":"3.0.1"}`)
	_, err := ParseApiResponse(bytes)
	if err != nil {
		t.Errorf("Error parsing OpenAPI response: %s", err)
	}
}

func TestOpenAPI3_1_0(t *testing.T) {
	bytes := []byte(`{"openapi":"3.1.0"}`)
	_, err := ParseApiResponse(bytes)
	if err != nil {
		t.Errorf("Error parsing OpenAPI response: %s", err)
	}
}

func TestMalformedOpenAPIResponse(t *testing.T) {
	bytes := []byte(`{"openapi":"3.1.0", "info":true}`)
	_, err := ParseApiResponse(bytes)
	if err == nil {
		t.Error("Should fail with malformed OpenAPI response")
	}
}
