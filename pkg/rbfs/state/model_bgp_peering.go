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

// BGP peering details.
type BgpPeering struct {
	InstanceName string `json:"instance_name,omitempty"`
	// The administrative peering state.
	AdministrativeState string    `json:"administrative_state,omitempty"`
	BgpState            *BgpState `json:"bgp_state,omitempty"`
	// Last BGP state transition.
	LastStateTransition string `json:"last_state_transition,omitempty"`
	// The reason for the last BGP reset.
	LastResetReason string `json:"last_reset_reason,omitempty"`
	// The local AS number.
	Asn     int32  `json:"asn,omitempty"`
	IflName string `json:"ifl_name,omitempty"`
	// The local IPv4 address.
	Ipv4Address string `json:"ipv4_address,omitempty"`
	// The local IPv6 address.
	Ipv6Address string `json:"ipv6_address,omitempty"`
	// The peering source port.
	SourcePort int32 `json:"source_port,omitempty"`
	// The local router ID
	RouterId string `json:"router_id,omitempty"`
	// The local BGP host name.
	HostName string `json:"host_name,omitempty"`
	// The local BGP domain name.
	DomainName string `json:"domain_name,omitempty"`
	// The BGP peering type.
	PeeringType string `json:"peering_type,omitempty"`
	// The BGP peer group name.
	PeerGroupName string                  `json:"peer_group_name,omitempty"`
	Peer          *BgpPeer                `json:"peer,omitempty"`
	Capabilities  *BgpPeeringCapabilities `json:"capabilities,omitempty"`
	Statistics    *BgpPeeringStatistics   `json:"statistics,omitempty"`
	Timers        *BgpPeeringTimers       `json:"timers,omitempty"`
}
