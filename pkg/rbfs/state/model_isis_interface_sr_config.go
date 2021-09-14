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

type IsisInterfaceSrConfig struct {
	IflName string `json:"ifl_name,omitempty"`
	// IPv4 SR index
	Ipv4Sid int32 `json:"ipv4_sid,omitempty"`
	// IPv6 SR index
	Ipv6Sid int32 `json:"ipv6_sid,omitempty"`
	// IPv4 SR anycast index
	Ipv4AnycastSid int32 `json:"ipv4_anycast_sid,omitempty"`
	// IPv6 SR anycast index
	Ipv6AnycastSid int32 `json:"ipv6_anycast_sid,omitempty"`
}
