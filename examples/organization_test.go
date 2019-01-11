package examples

import (
	"testing"

	"github.com/PennState/golang_scimclient/scim"
	"github.com/stretchr/testify/assert"
)

const organizationJSON = `
{
	"id": "430beb5c-a361-4c04-b308-2845789a496e",
	"schemas": ["urn:com:example:2.0:Organization"],
	"name": "Tour Promotion",
	"type": "Department",
	"parent": "../Organizations/4a7741a3-a436-4a52-a6d5-149e6c1b9578",
	"children": [
		"../Organizations/7eb59c46-35a4-4443-b8c1-5de8be88f973",
		"../Organizations/66506f29-8c44-414e-b52d-a993b94f370c",
		"../Organizations/0a365d4f-10e5-45c5-ae05-ee5184b59627"
	],
	"meta": {
		"resourceType": "Organization",
		"created": "2010-01-23T04:56:22Z",
		"lastModified": "2011-05-13T04:42:34Z",
		"version": "W/3694e05e9dff590",
		"location": "https://example.com/v2/Organizations/430beb5c-a361-4c04-b308-2845789a496e"
	}
}`

func TestOrganizationUnmarshaling(t *testing.T) {
	var organization Organization
	err := scim.Unmarshal([]byte(organizationJSON), &organization)
	assert.Nil(t, err)
}
