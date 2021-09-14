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

// The ephemeral configuration of a L2BSA service. This configuration is lost if the switch reboots.
type L2bsaServiceConfig struct {
	IfpName string `json:"ifp_name,omitempty"`
	// The ANP VLAN ID.
	AnpVlan int32 `json:"anp_vlan,omitempty"`
	// The AAA profile name.
	AaaProfileName string                 `json:"aaa_profile_name,omitempty"`
	L2x            *L2xConfig             `json:"l2x,omitempty"`
	Qos            *L2bsaServiceConfigQos `json:"qos,omitempty"`
	AccessLineInfo *AccessLineInfo        `json:"access_line_info,omitempty"`
}
