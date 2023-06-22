/*
 * RBFS Operational State API
 *
 * This contract describes the RBFS Operational State API contract defined by RBMS, the RtBrick Management System. This API is a _consumer-driven_ API, which means that all changes to this API **must be approved** by RBMS, the consumer of this API to avoid compatibility issues.  The API is kept backwards-compatible and anyone is allowed to _use_ this API.  The consumer of the API _must_ ignore additional attributes not explained in this specification. Additional attributes are _not_ considered violating backwards compatibility. In contrary, additional attributes allow extending the API while preserving backward compatibility.
 *
 * API version: 1.0.0
 * Contact: martin@rtbrick.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package state

// List of logical interfaces.
type PhysicalInterfaceSummaryRefLogicals struct {
	IflName string `json:"ifl_name,omitempty"`
	// The interface alias.
	IflAlias string `json:"ifl_alias,omitempty"`
	// The interface type.
	IflType string `json:"ifl_type,omitempty"`
	// The administrative interface state.
	AdministrativeState string `json:"administrative_state,omitempty"`
	// The operational interface state.
	OperationalState string `json:"operational_state,omitempty"`
	// The assigned IPv4 addresses.
	Ipv4Addresses []string `json:"ipv4_addresses,omitempty"`
	// The assigned IPv6 addresses.
	Ipv6Addresses []string `json:"ipv6_addresses,omitempty"`
}
