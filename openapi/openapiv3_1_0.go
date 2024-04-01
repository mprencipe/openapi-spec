package openapi

type OpenAPI310ApiResponse struct {
	Openapi           string                           `json:"openapi"`
	Info              OpenAPI310Info                   `json:"info"`
	JsonSchemaDialect *string                          `json:"jsonSchemaDialect"`
	Webhooks          *map[string]OpenAPI310Path       `json:"webhooks"`
	Paths             *map[string]OpenAPI310Path       `json:"paths"`
	Components        OpenAPI310Components             `json:"components"`
	Security          *[]map[string][]string           `json:"security"`
	Servers           *[]OpenAPI310Server              `json:"servers"`
	Tags              *[]OpenAPI310Tag                 `json:"tags"`
	ExternalDocs      *OpenAPI310ExternalDocumentation `json:"externalDocs"`
}

type OpenAPI310Path struct {
	Ref         *string                `json:"$ref"`
	Summary     *string                `json:"summary"`
	Description *string                `json:"description"`
	Get         *OpenAPI310Operation   `json:"get"`
	Put         *OpenAPI310Operation   `json:"put"`
	Post        *OpenAPI310Operation   `json:"post"`
	Delete      *OpenAPI310Operation   `json:"delete"`
	Options     *OpenAPI310Operation   `json:"options"`
	Head        *OpenAPI310Operation   `json:"head"`
	Patch       *OpenAPI310Operation   `json:"patch"`
	Trace       *OpenAPI310Operation   `json:"trace"`
	Servers     *[]OpenAPI310Server    `json:"servers"`
	Parameters  *[]OpenAPI310Parameter `json:"parameters"`
}

type OpenAPI310Info struct {
	Title          string             `json:"title"`
	Summary        *string            `json:"summary"`
	Description    *string            `json:"description"`
	Version        string             `json:"version"`
	TermsOfService *string            `json:"termsOfService"`
	Contact        *OpenAPI310Contact `json:"contact"`
	License        *OpenAPI310License `json:"license"`
}

type OpenAPI310Contact struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
	URL   *string `json:"url"`
}

type OpenAPI310License struct {
	Name       string  `json:"name"`
	Identifier *string `json:"identifier"`
	URL        *string `json:"url"`
}

type OpenAPI310Operation struct {
	Tags         *[]string                        `json:"tags"`
	Summary      *string                          `json:"summary"`
	Description  *string                          `json:"description"`
	ExternalDocs *OpenAPI310ExternalDocumentation `json:"externalDocs"`
	OperationId  *string                          `json:"operationId"`
	Parameters   *[]OpenAPI310Parameter           `json:"parameters"`
	RequestBody  *OpenAPI310RequestBody           `json:"requestBody"`
	Responses    *map[string]OpenAPI310Response   `json:"responses"`
	Callbacks    *map[string]OpenAPI310Callback   `json:"callbacks"`
	Deprecated   *bool                            `json:"deprecated"`
	Security     *[]map[string][]string           `json:"security"`
	Servers      *[]OpenAPI310Server              `json:"servers"`
}

type OpenAPI310Response struct {
	Description string                          `json:"description"`
	Headers     *map[string]OpenAPI310Header    `json:"headers"`
	Content     *map[string]OpenAPI310MediaType `json:"content"`
	Links       *map[string]OpenAPI310Link      `json:"links"`
	Ref         *string                         `json:"$ref"`
	Summary     *string                         `json:"summary"`
}

type OpenAPI310Parameter struct {
	Name            string                          `json:"name"`
	In              string                          `json:"in"`
	Description     *string                         `json:"description"`
	Required        *bool                           `json:"required"`
	Deprecated      *bool                           `json:"deprecated"`
	AllowEmptyValue *bool                           `json:"allowEmptyValue"`
	Style           *string                         `json:"style"`
	Explode         *bool                           `json:"explode"`
	AllowReserved   *bool                           `json:"allowReserved"`
	Schema          *OpenAPI310Schema               `json:"schema"`
	Example         *interface{}                    `json:"example"`
	Examples        *map[string]OpenAPI310Example   `json:"examples"`
	Content         *map[string]OpenAPI310MediaType `json:"content"`
}

type OpenAPI310Schema struct {
	Type          string                           `json:"type"`
	Discriminator *OpenAPI310Discriminator         `json:"discriminator"`
	Xml           *OpenAPI310Xml                   `json:"xml"`
	ExternalDocs  *OpenAPI310ExternalDocumentation `json:"externalDocs"`
	Example       *interface{}                     `json:"example"`
}

type OpenAPI310Discriminator struct {
	PropertyName string             `json:"propertyName"`
	Mapping      *map[string]string `json:"mapping"`
}

type OpenAPI310Xml struct {
	Name      *string `json:"name"`
	Namespace *string `json:"namespace"`
	Prefix    *string `json:"prefix"`
	Attribute *bool   `json:"attribute"`
	Wrapped   *bool   `json:"wrapped"`
}

type OpenAPI310ResponseRef struct {
	Ref         string  `json:"$ref"`
	Description *string `json:"description"`
	Summary     *string `json:"summary"`
}

type OpenAPI310Components struct {
	Schemas         *map[string]OpenAPI310Schema         `json:"schemas"`
	Responses       *map[string]OpenAPI310Response       `json:"responses"`
	Parameters      *map[string]OpenAPI310Parameter      `json:"parameters"`
	Examples        *map[string]OpenAPI310Example        `json:"examples"`
	RequestBodies   *map[string]OpenAPI310RequestBody    `json:"requestBodies"`
	Headers         *map[string]OpenAPI310Header         `json:"headers"`
	SecuritySchemes *map[string]OpenAPI310SecurityScheme `json:"securitySchemes"`
	Links           *map[string]OpenAPI310Link           `json:"links"`
	Callbacks       *map[string]OpenAPI310Callback       `json:"callbacks"`
	PathItems       *map[string]OpenAPI310Path           `json:"pathItems"`
}

type OpenAPI310RequestBody struct {
	Description *string                        `json:"description"`
	Content     map[string]OpenAPI310MediaType `json:"content"`
	Required    *bool                          `json:"required"`
}

type OpenAPI310MediaType struct {
	Schema   *OpenAPI310Schema              `json:"schema"`
	Example  *OpenAPI310Example             `json:"example"`
	Examples *map[string]OpenAPI310Example  `json:"examples"`
	Encoding *map[string]OpenAPI310Encoding `json:"encoding"`
}

type OpenAPI310Encoding struct {
	ContentType   *string                      `json:"contentType"`
	Headers       *map[string]OpenAPI310Header `json:"headers"`
	Style         *string                      `json:"style"`
	Explode       *bool                        `json:"explode"`
	AllowReserved *bool                        `json:"allowReserved"`
}

type OpenAPI310Header struct {
	Description     *string                       `json:"description"`
	Required        *bool                         `json:"required"`
	Deprecated      *bool                         `json:"deprecated"`
	AllowEmptyValue *bool                         `json:"allowEmptyValue"`
	Style           *string                       `json:"style"`
	Explode         *bool                         `json:"explode"`
	AllowReserved   *bool                         `json:"allowReserved"`
	Schema          *OpenAPI310Schema             `json:"schema"`
	Example         *interface{}                  `json:"example"`
	Examples        *map[string]OpenAPI310Example `json:"examples"`
}

type OpenAPI310SecurityScheme struct {
	Type             *string               `json:"type"`
	Description      *string               `json:"description"`
	Name             *string               `json:"name"`
	In               *string               `json:"in"`
	Scheme           *string               `json:"scheme"`
	BearerFormat     *string               `json:"bearerFormat"`
	Flows            *OpenAPI310OAuthFlows `json:"flows"`
	OpenIdConnectUrl *string               `json:"openIdConnectUrl"`
}

type OpenAPI310OAuthFlows struct {
	Implicit          *OpenAPI310OAuthFlow `json:"implicit"`
	Password          *OpenAPI310OAuthFlow `json:"password"`
	ClientCredentials *OpenAPI310OAuthFlow `json:"clientCredentials"`
	AuthorizationCode *OpenAPI310OAuthFlow `json:"authorizationCode"`
}

type OpenAPI310OAuthFlow struct {
	AuthorizationUrl *string            `json:"authorizationUrl"`
	TokenUrl         *string            `json:"tokenUrl"`
	RefreshUrl       *string            `json:"refreshUrl"`
	Scopes           *map[string]string `json:"scopes"`
}

type OpenAPI310Link struct {
	OperationRef *string            `json:"operationRef"`
	OperationId  *string            `json:"operationId"`
	Parameters   *map[string]string `json:"parameters"`
	RequestBody  *string            `json:"requestBody"`
	Description  *string            `json:"description"`
	Server       *OpenAPI310Server  `json:"server"`
}

type OpenAPI310Callback struct {
	Expression *string `json:"expression"`
}

type OpenAPI310Example struct {
	Summary       *string `json:"summary"`
	Description   *string `json:"description"`
	Value         *string `json:"value"`
	ExternalValue *string `json:"externalValue"`
}

type OpenAPI310Server struct {
	URL         string                               `json:"url"`
	Description *string                              `json:"description"`
	Variables   *map[string]OpenAPI310ServerVariable `json:"variables"`
}

type OpenAPI310ServerVariable struct {
	Default     string    `json:"default"`
	Description *string   `json:"description"`
	Enum        *[]string `json:"enum"`
}

type OpenAPI310Tag struct {
	Name         string                           `json:"name"`
	Description  *string                          `json:"description"`
	ExternalDocs *OpenAPI310ExternalDocumentation `json:"externalDocs"`
}

type OpenAPI310ExternalDocumentation struct {
	Description *string `json:"description"`
	URL         string  `json:"url"`
}

func (resp OpenAPI310ApiResponse) IsSwaggerResponse() bool {
	return false
}

func (resp OpenAPI310ApiResponse) IsOpenApiResponse() bool {
	return true
}
