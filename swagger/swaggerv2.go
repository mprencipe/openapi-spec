package swagger

type SwaggerV2Response struct {
	Swagger      string                         `json:"swagger"`
	Info         Info                           `json:"info"`
	Host         *string                        `json:"host"`
	BasePath     *string                        `json:"basePath"`
	Schemes      *[]string                      `json:"schemes"`
	Consumes     *[]string                      `json:"consumes"`
	Produces     *[]string                      `json:"produces"`
	Paths        *map[string]Path               `json:"paths"`
	Definitions  *map[string]Schema             `json:"definitions"`
	Parameters   *map[string]Parameter          `json:"parameters"`
	Responses    *map[string]Response           `json:"responses"`
	Security     *[]map[string][]string         `json:"security"`
	SecurityDefs *map[string]SecurityDefinition `json:"securityDefinitions"`
	Tags         *[]Tag                         `json:"tags"`
	ExternalDocs *ExternalDocumentation         `json:"externalDocs"`
}

type Info struct {
	Title          string   `json:"title"`
	Description    *string  `json:"description"`
	TermsOfService *string  `json:"termsOfService"`
	Contact        *Contact `json:"contact"`
	License        *License `json:"license"`
	Version        string   `json:"version"`
}

type Contact struct {
	Name  *string `json:"name"`
	Url   *string `json:"url"`
	Email *string `json:"email"`
}

type License struct {
	Name *string `json:"name"`
	Url  *string `json:"url"`
}

type Path struct {
	Ref        *string               `json:"$ref"`
	Get        *Operation            `json:"get"`
	Put        *Operation            `json:"put"`
	Post       *Operation            `json:"post"`
	Delete     *Operation            `json:"delete"`
	Options    *Operation            `json:"options"`
	Head       *Operation            `json:"head"`
	Patch      *Operation            `json:"patch"`
	Parameters *map[string]Parameter `json:"parameters"`
}

type Operation struct {
	Tags         *[]string                      `json:"tags"`
	Summary      *string                        `json:"summary"`
	Description  *string                        `json:"description"`
	ExternalDocs *ExternalDocumentation         `json:"externalDocs"`
	OperationId  *string                        `json:"operationId"`
	Consumes     *[]string                      `json:"consumes"`
	Produces     *[]string                      `json:"produces"`
	Parameters   *[]Parameter                   `json:"parameters"`
	Responses    *map[string]Response           `json:"responses"`
	Schemes      *[]string                      `json:"schemes"`
	Deprecated   *bool                          `json:"deprecated"`
	Security     *[]map[string][]string         `json:"security"`
	SecurityDefs *map[string]SecurityDefinition `json:"securityDefinitions"`
}

type Parameter struct {
	Name             string  `json:"name"`
	In               string  `json:"in"`
	Description      *string `json:"description"`
	Required         bool    `json:"required"`
	Schema           *Schema `json:"schema"`
	Type             string  `json:"type"`
	Format           string  `json:"format"`
	Items            *Items  `json:"items"`
	CollectionFormat string  `json:"collectionFormat"`
	Default          *string `json:"default"`
}

type Schema struct {
	Ref                  *string                `json:"$ref"`
	Format               *string                `json:"format"`
	Title                *string                `json:"title"`
	Description          *string                `json:"description"`
	Default              *string                `json:"default"`
	MultipleOf           *float64               `json:"multipleOf"`
	Maximum              *float64               `json:"maximum"`
	ExclusiveMaximum     *bool                  `json:"exclusiveMaximum"`
	Minimum              *float64               `json:"minimum"`
	ExclusiveMinimum     *bool                  `json:"exclusiveMinimum"`
	MaxLength            *int                   `json:"maxLength"`
	MinLength            *int                   `json:"minLength"`
	Pattern              *string                `json:"pattern"`
	MaxItems             *int                   `json:"maxItems"`
	MinItems             *int                   `json:"minItems"`
	UniqueItems          *bool                  `json:"uniqueItems"`
	MaxProperties        *int                   `json:"maxProperties"`
	MinProperties        *int                   `json:"minProperties"`
	Required             *[]string              `json:"required"`
	Enum                 *[]string              `json:"enum"`
	Type                 *string                `json:"type"`
	Items                *Items                 `json:"items"`
	AllOf                *[]Schema              `json:"allOf"`
	Properties           *map[string]Schema     `json:"properties"`
	AdditionalProperties *Schema                `json:"additionalProperties"`
	ReadOnly             *bool                  `json:"readOnly"`
	Xml                  *XML                   `json:"xml"`
	ExternalDocs         *ExternalDocumentation `json:"externalDocs"`
	Example              *string                `json:"example"`
}

type Items struct {
	Ref                  *string                `json:"$ref"`
	Format               *string                `json:"format"`
	Title                *string                `json:"title"`
	Description          *string                `json:"description"`
	Default              *string                `json:"default"`
	MultipleOf           *float64               `json:"multipleOf"`
	Maximum              *float64               `json:"maximum"`
	ExclusiveMaximum     *bool                  `json:"exclusiveMaximum"`
	Minimum              *float64               `json:"minimum"`
	ExclusiveMinimum     *bool                  `json:"exclusiveMinimum"`
	MaxLength            *int                   `json:"maxLength"`
	MinLength            *int                   `json:"minLength"`
	Pattern              *string                `json:"pattern"`
	MaxItems             *int                   `json:"maxItems"`
	MinItems             *int                   `json:"minItems"`
	UniqueItems          *bool                  `json:"uniqueItems"`
	MaxProperties        *int                   `json:"maxProperties"`
	MinProperties        *int                   `json:"minProperties"`
	Required             *[]string              `json:"required"`
	Enum                 *[]string              `json:"enum"`
	Type                 *string                `json:"type"`
	Items                *Items                 `json:"items"`
	AllOf                *[]Schema              `json:"allOf"`
	Properties           *map[string]Schema     `json:"properties"`
	AdditionalProperties *Schema                `json:"additionalProperties"`
	ReadOnly             *bool                  `json:"readOnly"`
	Xml                  *XML                   `json:"xml"`
	ExternalDocs         *ExternalDocumentation `json:"externalDocs"`
	Example              *string                `json:"example"`
}

type XML struct {
	Name      *string `json:"name"`
	Namespace *string `json:"namespace"`
	Prefix    *string `json:"prefix"`
	Attribute *bool   `json:"attribute"`
	Wrapped   *bool   `json:"wrapped"`
}

type Response struct {
	Description *string             `json:"description"`
	Schema      *Schema             `json:"schema"`
	Headers     *map[string]Header  `json:"headers"`
	Examples    *map[string]Example `json:"examples"`
}

type Header struct {
	Description      *string   `json:"description"`
	Type             *string   `json:"type"`
	Format           *string   `json:"format"`
	Items            *Items    `json:"items"`
	CollectionFormat *string   `json:"collectionFormat"`
	Default          *string   `json:"default"`
	Maximum          *float64  `json:"maximum"`
	ExclusiveMaximum *bool     `json:"exclusiveMaximum"`
	Minimum          *float64  `json:"minimum"`
	ExclusiveMinimum *bool     `json:"exclusiveMinimum"`
	MaxLength        *int      `json:"maxLength"`
	MinLength        *int      `json:"minLength"`
	Pattern          *string   `json:"pattern"`
	MaxItems         *int      `json:"maxItems"`
	MinItems         *int      `json:"minItems"`
	UniqueItems      *bool     `json:"uniqueItems"`
	Enum             *[]string `json:"enum"`
	MultipleOf       *float64  `json:"multipleOf"`
}

type Example struct {
	Ref           *string `json:"$ref"`
	Summary       *string `json:"summary"`
	Description   *string `json:"description"`
	Value         *string `json:"value"`
	ExternalValue *string `json:"externalValue"`
}

type SecurityDefinition struct {
	Type             string             `json:"type"`
	Description      *string            `json:"description"`
	Name             string             `json:"name"`
	In               string             `json:"in"`
	Flow             string             `json:"flow"`
	AuthorizationUrl *string            `json:"authorizationUrl"`
	TokenUrl         *string            `json:"tokenUrl"`
	Scopes           *map[string]string `json:"scopes"`
}

type Tag struct {
	Name         string                 `json:"name"`
	Description  *string                `json:"description"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs"`
}

type ExternalDocumentation struct {
	Description *string `json:"description"`
	Url         string  `json:"url"`
}

func (resp SwaggerV2Response) IsSwaggerResponse() bool {
	return true
}

func (resp SwaggerV2Response) IsOpenApiResponse() bool {
	return false
}
