//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdigitaltwins

import "time"

// AzureDataExplorerConnectionProperties - Properties of a time series database connection to Azure Data Explorer with data
// being sent via an EventHub.
type AzureDataExplorerConnectionProperties struct {
	// REQUIRED; The name of the Azure Data Explorer database.
	AdxDatabaseName *string `json:"adxDatabaseName,omitempty"`

	// REQUIRED; The URI of the Azure Data Explorer endpoint.
	AdxEndpointURI *string `json:"adxEndpointUri,omitempty"`

	// REQUIRED; The resource ID of the Azure Data Explorer cluster.
	AdxResourceID *string `json:"adxResourceId,omitempty"`

	// REQUIRED; The type of time series connection resource.
	ConnectionType *ConnectionType `json:"connectionType,omitempty"`

	// REQUIRED; The URL of the EventHub namespace for identity-based authentication. It must include the protocol sb://
	EventHubEndpointURI *string `json:"eventHubEndpointUri,omitempty"`

	// REQUIRED; The EventHub name in the EventHub namespace for identity-based authentication.
	EventHubEntityPath *string `json:"eventHubEntityPath,omitempty"`

	// REQUIRED; The resource ID of the EventHub namespace.
	EventHubNamespaceResourceID *string `json:"eventHubNamespaceResourceId,omitempty"`

	// The name of the Azure Data Explorer table.
	AdxTableName *string `json:"adxTableName,omitempty"`

	// The EventHub consumer group to use when ADX reads from EventHub. Defaults to $Default.
	EventHubConsumerGroup *string `json:"eventHubConsumerGroup,omitempty"`

	// READ-ONLY; The provisioning state.
	ProvisioningState *TimeSeriesDatabaseConnectionState `json:"provisioningState,omitempty" azure:"ro"`
}

// GetTimeSeriesDatabaseConnectionProperties implements the TimeSeriesDatabaseConnectionPropertiesClassification interface
// for type AzureDataExplorerConnectionProperties.
func (a *AzureDataExplorerConnectionProperties) GetTimeSeriesDatabaseConnectionProperties() *TimeSeriesDatabaseConnectionProperties {
	return &TimeSeriesDatabaseConnectionProperties{
		ConnectionType:    a.ConnectionType,
		ProvisioningState: a.ProvisioningState,
	}
}

// CheckNameRequest - The result returned from a database check name availability request.
type CheckNameRequest struct {
	// REQUIRED; Resource name.
	Name *string `json:"name,omitempty"`

	// REQUIRED; The type of resource, for instance Microsoft.DigitalTwins/digitalTwinsInstances.
	Type *string `json:"type,omitempty"`
}

// CheckNameResult - The result returned from a check name availability request.
type CheckNameResult struct {
	// Message indicating an unavailable name due to a conflict, or a description of the naming rules that are violated.
	Message *string `json:"message,omitempty"`

	// Specifies a Boolean value that indicates if the name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty"`

	// Message providing the reason why the given name is invalid.
	Reason *Reason `json:"reason,omitempty"`
}

// ClientBeginCreateOrUpdateOptions contains the optional parameters for the Client.BeginCreateOrUpdate method.
type ClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClientBeginDeleteOptions contains the optional parameters for the Client.BeginDelete method.
type ClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClientBeginUpdateOptions contains the optional parameters for the Client.BeginUpdate method.
type ClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClientCheckNameAvailabilityOptions contains the optional parameters for the Client.CheckNameAvailability method.
type ClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// ClientGetOptions contains the optional parameters for the Client.Get method.
type ClientGetOptions struct {
	// placeholder for future optional parameters
}

// ClientListByResourceGroupOptions contains the optional parameters for the Client.ListByResourceGroup method.
type ClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ClientListOptions contains the optional parameters for the Client.List method.
type ClientListOptions struct {
	// placeholder for future optional parameters
}

// ConnectionProperties - The properties of a private endpoint connection.
type ConnectionProperties struct {
	// The list of group ids for the private endpoint connection.
	GroupIDs []*string `json:"groupIds,omitempty"`

	// The private endpoint.
	PrivateEndpoint *PrivateEndpoint `json:"privateEndpoint,omitempty"`

	// The connection state.
	PrivateLinkServiceConnectionState *ConnectionPropertiesPrivateLinkServiceConnectionState `json:"privateLinkServiceConnectionState,omitempty"`

	// READ-ONLY; The provisioning state.
	ProvisioningState *ConnectionPropertiesProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// ConnectionPropertiesPrivateLinkServiceConnectionState - The connection state.
type ConnectionPropertiesPrivateLinkServiceConnectionState struct {
	// REQUIRED; The description for the current state of a private endpoint connection.
	Description *string `json:"description,omitempty"`

	// REQUIRED; The status of a private endpoint connection.
	Status *PrivateLinkServiceConnectionStatus `json:"status,omitempty"`

	// Actions required for a private endpoint connection.
	ActionsRequired *string `json:"actionsRequired,omitempty"`
}

// ConnectionState - The current state of a private endpoint connection.
type ConnectionState struct {
	// REQUIRED; The description for the current state of a private endpoint connection.
	Description *string `json:"description,omitempty"`

	// REQUIRED; The status of a private endpoint connection.
	Status *PrivateLinkServiceConnectionStatus `json:"status,omitempty"`

	// Actions required for a private endpoint connection.
	ActionsRequired *string `json:"actionsRequired,omitempty"`
}

// Description - The description of the DigitalTwins service.
type Description struct {
	// REQUIRED; The resource location.
	Location *string `json:"location,omitempty"`

	// The managed identity for the DigitalTwinsInstance.
	Identity *Identity `json:"identity,omitempty"`

	// DigitalTwins instance properties.
	Properties *Properties `json:"properties,omitempty"`

	// The resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The resource identifier.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the DigitalTwinsInstance.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DescriptionListResult - A list of DigitalTwins description objects with a next link.
type DescriptionListResult struct {
	// The link used to get the next page of DigitalTwins description objects.
	NextLink *string `json:"nextLink,omitempty"`

	// A list of DigitalTwins description objects.
	Value []*Description `json:"value,omitempty"`
}

// EndpointClientBeginCreateOrUpdateOptions contains the optional parameters for the EndpointClient.BeginCreateOrUpdate method.
type EndpointClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// EndpointClientBeginDeleteOptions contains the optional parameters for the EndpointClient.BeginDelete method.
type EndpointClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// EndpointClientGetOptions contains the optional parameters for the EndpointClient.Get method.
type EndpointClientGetOptions struct {
	// placeholder for future optional parameters
}

// EndpointClientListOptions contains the optional parameters for the EndpointClient.List method.
type EndpointClientListOptions struct {
	// placeholder for future optional parameters
}

// EndpointResource - DigitalTwinsInstance endpoint resource.
type EndpointResource struct {
	// REQUIRED; DigitalTwinsInstance endpoint resource properties.
	Properties EndpointResourcePropertiesClassification `json:"properties,omitempty"`

	// READ-ONLY; The resource identifier.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Extension resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// EndpointResourceListResult - A list of DigitalTwinsInstance Endpoints with a next link.
type EndpointResourceListResult struct {
	// The link used to get the next page of DigitalTwinsInstance Endpoints.
	NextLink *string `json:"nextLink,omitempty"`

	// A list of DigitalTwinsInstance Endpoints.
	Value []*EndpointResource `json:"value,omitempty"`
}

// EndpointResourcePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetEndpointResourceProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *EndpointResourceProperties, *EventGrid, *EventHub, *ServiceBus
type EndpointResourcePropertiesClassification interface {
	// GetEndpointResourceProperties returns the EndpointResourceProperties content of the underlying type.
	GetEndpointResourceProperties() *EndpointResourceProperties
}

// EndpointResourceProperties - Properties related to Digital Twins Endpoint
type EndpointResourceProperties struct {
	// REQUIRED; The type of Digital Twins endpoint
	EndpointType *EndpointType `json:"endpointType,omitempty"`

	// Specifies the authentication type being used for connecting to the endpoint. Defaults to 'KeyBased'. If 'KeyBased' is selected,
	// a connection string must be specified (at least the primary connection
	// string). If 'IdentityBased' is select, the endpointUri and entityPath properties must be specified.
	AuthenticationType *AuthenticationType `json:"authenticationType,omitempty"`

	// Dead letter storage secret for key-based authentication. Will be obfuscated during read.
	DeadLetterSecret *string `json:"deadLetterSecret,omitempty"`

	// Dead letter storage URL for identity-based authentication.
	DeadLetterURI *string `json:"deadLetterUri,omitempty"`

	// READ-ONLY; Time when the Endpoint was added to DigitalTwinsInstance.
	CreatedTime *time.Time `json:"createdTime,omitempty" azure:"ro"`

	// READ-ONLY; The provisioning state.
	ProvisioningState *EndpointProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// GetEndpointResourceProperties implements the EndpointResourcePropertiesClassification interface for type EndpointResourceProperties.
func (e *EndpointResourceProperties) GetEndpointResourceProperties() *EndpointResourceProperties {
	return e
}

// ErrorDefinition - Error definition.
type ErrorDefinition struct {
	// READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Internal error details.
	Details []*ErrorDefinition `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; Description of the error.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// ErrorResponse - Error response.
type ErrorResponse struct {
	// Error description
	Error *ErrorDefinition `json:"error,omitempty"`
}

// EventGrid - Properties related to EventGrid.
type EventGrid struct {
	// REQUIRED; EventGrid secondary accesskey. Will be obfuscated during read.
	AccessKey1 *string `json:"accessKey1,omitempty"`

	// REQUIRED; The type of Digital Twins endpoint
	EndpointType *EndpointType `json:"endpointType,omitempty"`

	// REQUIRED; EventGrid Topic Endpoint.
	TopicEndpoint *string `json:"TopicEndpoint,omitempty"`

	// EventGrid secondary accesskey. Will be obfuscated during read.
	AccessKey2 *string `json:"accessKey2,omitempty"`

	// Specifies the authentication type being used for connecting to the endpoint. Defaults to 'KeyBased'. If 'KeyBased' is selected,
	// a connection string must be specified (at least the primary connection
	// string). If 'IdentityBased' is select, the endpointUri and entityPath properties must be specified.
	AuthenticationType *AuthenticationType `json:"authenticationType,omitempty"`

	// Dead letter storage secret for key-based authentication. Will be obfuscated during read.
	DeadLetterSecret *string `json:"deadLetterSecret,omitempty"`

	// Dead letter storage URL for identity-based authentication.
	DeadLetterURI *string `json:"deadLetterUri,omitempty"`

	// READ-ONLY; Time when the Endpoint was added to DigitalTwinsInstance.
	CreatedTime *time.Time `json:"createdTime,omitempty" azure:"ro"`

	// READ-ONLY; The provisioning state.
	ProvisioningState *EndpointProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// GetEndpointResourceProperties implements the EndpointResourcePropertiesClassification interface for type EventGrid.
func (e *EventGrid) GetEndpointResourceProperties() *EndpointResourceProperties {
	return &EndpointResourceProperties{
		EndpointType:       e.EndpointType,
		ProvisioningState:  e.ProvisioningState,
		CreatedTime:        e.CreatedTime,
		AuthenticationType: e.AuthenticationType,
		DeadLetterSecret:   e.DeadLetterSecret,
		DeadLetterURI:      e.DeadLetterURI,
	}
}

// EventHub - Properties related to EventHub.
type EventHub struct {
	// REQUIRED; The type of Digital Twins endpoint
	EndpointType *EndpointType `json:"endpointType,omitempty"`

	// Specifies the authentication type being used for connecting to the endpoint. Defaults to 'KeyBased'. If 'KeyBased' is selected,
	// a connection string must be specified (at least the primary connection
	// string). If 'IdentityBased' is select, the endpointUri and entityPath properties must be specified.
	AuthenticationType *AuthenticationType `json:"authenticationType,omitempty"`

	// PrimaryConnectionString of the endpoint for key-based authentication. Will be obfuscated during read.
	ConnectionStringPrimaryKey *string `json:"connectionStringPrimaryKey,omitempty"`

	// SecondaryConnectionString of the endpoint for key-based authentication. Will be obfuscated during read.
	ConnectionStringSecondaryKey *string `json:"connectionStringSecondaryKey,omitempty"`

	// Dead letter storage secret for key-based authentication. Will be obfuscated during read.
	DeadLetterSecret *string `json:"deadLetterSecret,omitempty"`

	// Dead letter storage URL for identity-based authentication.
	DeadLetterURI *string `json:"deadLetterUri,omitempty"`

	// The URL of the EventHub namespace for identity-based authentication. It must include the protocol 'sb://'.
	EndpointURI *string `json:"endpointUri,omitempty"`

	// The EventHub name in the EventHub namespace for identity-based authentication.
	EntityPath *string `json:"entityPath,omitempty"`

	// READ-ONLY; Time when the Endpoint was added to DigitalTwinsInstance.
	CreatedTime *time.Time `json:"createdTime,omitempty" azure:"ro"`

	// READ-ONLY; The provisioning state.
	ProvisioningState *EndpointProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// GetEndpointResourceProperties implements the EndpointResourcePropertiesClassification interface for type EventHub.
func (e *EventHub) GetEndpointResourceProperties() *EndpointResourceProperties {
	return &EndpointResourceProperties{
		EndpointType:       e.EndpointType,
		ProvisioningState:  e.ProvisioningState,
		CreatedTime:        e.CreatedTime,
		AuthenticationType: e.AuthenticationType,
		DeadLetterSecret:   e.DeadLetterSecret,
		DeadLetterURI:      e.DeadLetterURI,
	}
}

// ExternalResource - Definition of a resource.
type ExternalResource struct {
	// READ-ONLY; The resource identifier.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Extension resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// GroupIDInformation - The group information for creating a private endpoint on Digital Twin.
type GroupIDInformation struct {
	// REQUIRED; The group information properties.
	Properties *GroupIDInformationProperties `json:"properties,omitempty"`

	// The resource identifier.
	ID *string `json:"id,omitempty"`

	// READ-ONLY; The resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// GroupIDInformationProperties - The properties for a group information object.
type GroupIDInformationProperties struct {
	// The group id.
	GroupID *string `json:"groupId,omitempty"`

	// The required members for a specific group id.
	RequiredMembers []*string `json:"requiredMembers,omitempty"`

	// The required DNS zones for a specific group id.
	RequiredZoneNames []*string `json:"requiredZoneNames,omitempty"`
}

// GroupIDInformationResponse - The available private link resources for a Digital Twin.
type GroupIDInformationResponse struct {
	// The list of available private link resources for a Digital Twin.
	Value []*GroupIDInformation `json:"value,omitempty"`
}

// Identity - The managed identity for the DigitalTwinsInstance.
type Identity struct {
	// The type of Managed Identity used by the DigitalTwinsInstance. Only SystemAssigned is supported.
	Type *DigitalTwinsIdentityType `json:"type,omitempty"`

	// READ-ONLY; The object id of the Managed Identity Resource. This will be sent to the RP from ARM via the x-ms-identity-principal-id
	// header in the PUT request if the resource has a systemAssigned(implicit)
	// identity
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant id of the Managed Identity Resource. This will be sent to the RP from ARM via the x-ms-client-tenant-id
	// header in the PUT request if the resource has a systemAssigned(implicit) identity
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// Operation - DigitalTwins service REST API operation
type Operation struct {
	// Operation properties display
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; If the operation is a data action (for data plane rbac).
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; Operation name: {provider}/{resource}/{read | write | action | delete}
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation.
	Origin *string `json:"origin,omitempty" azure:"ro"`

	// READ-ONLY; Operation properties.
	Properties map[string]interface{} `json:"properties,omitempty" azure:"ro"`
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// READ-ONLY; Friendly description for the operation.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; Name of the operation.
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; Service provider: Microsoft DigitalTwins.
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; Resource Type: DigitalTwinsInstances.
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of DigitalTwins service operations. It contains a list of operations and a URL link to get
// the next set of results.
type OperationListResult struct {
	// The link used to get the next page of DigitalTwins description objects.
	NextLink *string `json:"nextLink,omitempty"`

	// READ-ONLY; A list of DigitalTwins operations supported by the Microsoft.DigitalTwins resource provider.
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PatchDescription - The description of the DigitalTwins service.
type PatchDescription struct {
	// The managed identity for the DigitalTwinsInstance.
	Identity *Identity `json:"identity,omitempty"`

	// Properties for the DigitalTwinsInstance.
	Properties *PatchProperties `json:"properties,omitempty"`

	// Instance patch properties
	Tags map[string]*string `json:"tags,omitempty"`
}

// PatchProperties - The properties of a DigitalTwinsInstance.
type PatchProperties struct {
	// Public network access for the DigitalTwinsInstance.
	PublicNetworkAccess *PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`
}

// PrivateEndpoint - The private endpoint property of a private endpoint connection.
type PrivateEndpoint struct {
	// READ-ONLY; The resource identifier.
	ID *string `json:"id,omitempty" azure:"ro"`
}

// PrivateEndpointConnection - The private endpoint connection of a Digital Twin.
type PrivateEndpointConnection struct {
	// REQUIRED; The connection properties.
	Properties *ConnectionProperties `json:"properties,omitempty"`

	// READ-ONLY; The resource identifier.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the private endpoint connection.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginCreateOrUpdate
// method.
type PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginDelete
// method.
type PrivateEndpointConnectionsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
type PrivateEndpointConnectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientListOptions contains the optional parameters for the PrivateEndpointConnectionsClient.List
// method.
type PrivateEndpointConnectionsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsResponse - The available private link connections for a Digital Twin.
type PrivateEndpointConnectionsResponse struct {
	// The list of available private link connections for a Digital Twin.
	Value []*PrivateEndpointConnection `json:"value,omitempty"`
}

// PrivateLinkResourcesClientGetOptions contains the optional parameters for the PrivateLinkResourcesClient.Get method.
type PrivateLinkResourcesClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientListOptions contains the optional parameters for the PrivateLinkResourcesClient.List method.
type PrivateLinkResourcesClientListOptions struct {
	// placeholder for future optional parameters
}

// Properties - The properties of a DigitalTwinsInstance.
type Properties struct {
	// The private endpoint connections.
	PrivateEndpointConnections []*PrivateEndpointConnection `json:"privateEndpointConnections,omitempty"`

	// Public network access for the DigitalTwinsInstance.
	PublicNetworkAccess *PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// READ-ONLY; Time when DigitalTwinsInstance was created.
	CreatedTime *time.Time `json:"createdTime,omitempty" azure:"ro"`

	// READ-ONLY; Api endpoint to work with DigitalTwinsInstance.
	HostName *string `json:"hostName,omitempty" azure:"ro"`

	// READ-ONLY; Time when DigitalTwinsInstance was updated.
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty" azure:"ro"`

	// READ-ONLY; The provisioning state.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// Resource - The common properties of a DigitalTwinsInstance.
type Resource struct {
	// REQUIRED; The resource location.
	Location *string `json:"location,omitempty"`

	// The managed identity for the DigitalTwinsInstance.
	Identity *Identity `json:"identity,omitempty"`

	// The resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The resource identifier.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the DigitalTwinsInstance.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ServiceBus - Properties related to ServiceBus.
type ServiceBus struct {
	// REQUIRED; The type of Digital Twins endpoint
	EndpointType *EndpointType `json:"endpointType,omitempty"`

	// Specifies the authentication type being used for connecting to the endpoint. Defaults to 'KeyBased'. If 'KeyBased' is selected,
	// a connection string must be specified (at least the primary connection
	// string). If 'IdentityBased' is select, the endpointUri and entityPath properties must be specified.
	AuthenticationType *AuthenticationType `json:"authenticationType,omitempty"`

	// Dead letter storage secret for key-based authentication. Will be obfuscated during read.
	DeadLetterSecret *string `json:"deadLetterSecret,omitempty"`

	// Dead letter storage URL for identity-based authentication.
	DeadLetterURI *string `json:"deadLetterUri,omitempty"`

	// The URL of the ServiceBus namespace for identity-based authentication. It must include the protocol 'sb://'.
	EndpointURI *string `json:"endpointUri,omitempty"`

	// The ServiceBus Topic name for identity-based authentication.
	EntityPath *string `json:"entityPath,omitempty"`

	// PrimaryConnectionString of the endpoint for key-based authentication. Will be obfuscated during read.
	PrimaryConnectionString *string `json:"primaryConnectionString,omitempty"`

	// SecondaryConnectionString of the endpoint for key-based authentication. Will be obfuscated during read.
	SecondaryConnectionString *string `json:"secondaryConnectionString,omitempty"`

	// READ-ONLY; Time when the Endpoint was added to DigitalTwinsInstance.
	CreatedTime *time.Time `json:"createdTime,omitempty" azure:"ro"`

	// READ-ONLY; The provisioning state.
	ProvisioningState *EndpointProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// GetEndpointResourceProperties implements the EndpointResourcePropertiesClassification interface for type ServiceBus.
func (s *ServiceBus) GetEndpointResourceProperties() *EndpointResourceProperties {
	return &EndpointResourceProperties{
		EndpointType:       s.EndpointType,
		ProvisioningState:  s.ProvisioningState,
		CreatedTime:        s.CreatedTime,
		AuthenticationType: s.AuthenticationType,
		DeadLetterSecret:   s.DeadLetterSecret,
		DeadLetterURI:      s.DeadLetterURI,
	}
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// TimeSeriesDatabaseConnection - Describes a time series database connection resource.
type TimeSeriesDatabaseConnection struct {
	// Properties of a specific time series database connection.
	Properties TimeSeriesDatabaseConnectionPropertiesClassification `json:"properties,omitempty"`

	// READ-ONLY; The resource identifier.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Extension resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// TimeSeriesDatabaseConnectionListResult - A pageable list of time series database connection resources.
type TimeSeriesDatabaseConnectionListResult struct {
	// The link used to get the next page of results.
	NextLink *string `json:"nextLink,omitempty"`

	// A list of time series database connection resources.
	Value []*TimeSeriesDatabaseConnection `json:"value,omitempty"`
}

// TimeSeriesDatabaseConnectionPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetTimeSeriesDatabaseConnectionProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureDataExplorerConnectionProperties, *TimeSeriesDatabaseConnectionProperties
type TimeSeriesDatabaseConnectionPropertiesClassification interface {
	// GetTimeSeriesDatabaseConnectionProperties returns the TimeSeriesDatabaseConnectionProperties content of the underlying type.
	GetTimeSeriesDatabaseConnectionProperties() *TimeSeriesDatabaseConnectionProperties
}

// TimeSeriesDatabaseConnectionProperties - Properties of a time series database connection resource.
type TimeSeriesDatabaseConnectionProperties struct {
	// REQUIRED; The type of time series connection resource.
	ConnectionType *ConnectionType `json:"connectionType,omitempty"`

	// READ-ONLY; The provisioning state.
	ProvisioningState *TimeSeriesDatabaseConnectionState `json:"provisioningState,omitempty" azure:"ro"`
}

// GetTimeSeriesDatabaseConnectionProperties implements the TimeSeriesDatabaseConnectionPropertiesClassification interface
// for type TimeSeriesDatabaseConnectionProperties.
func (t *TimeSeriesDatabaseConnectionProperties) GetTimeSeriesDatabaseConnectionProperties() *TimeSeriesDatabaseConnectionProperties {
	return t
}

// TimeSeriesDatabaseConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the TimeSeriesDatabaseConnectionsClient.BeginCreateOrUpdate
// method.
type TimeSeriesDatabaseConnectionsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TimeSeriesDatabaseConnectionsClientBeginDeleteOptions contains the optional parameters for the TimeSeriesDatabaseConnectionsClient.BeginDelete
// method.
type TimeSeriesDatabaseConnectionsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TimeSeriesDatabaseConnectionsClientGetOptions contains the optional parameters for the TimeSeriesDatabaseConnectionsClient.Get
// method.
type TimeSeriesDatabaseConnectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// TimeSeriesDatabaseConnectionsClientListOptions contains the optional parameters for the TimeSeriesDatabaseConnectionsClient.List
// method.
type TimeSeriesDatabaseConnectionsClientListOptions struct {
	// placeholder for future optional parameters
}
