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

// Configured, received and negotiated peering capabilities.
type BgpPeeringCapabilities struct {
	// Accept 4-byte AS number.
	As4Capable bool `json:"as4_capable,omitempty"`
	// Sent AS4 capability status.
	As4CapableSent bool `json:"as4_capable_sent,omitempty"`
	// Received AS4 capability status.
	As4CapableReceived bool `json:"as4_capable_received,omitempty"`
	// Resend all prefixes if the peer asks for it.
	RouteRefreshCapable bool `json:"route_refresh_capable,omitempty"`
	// Sent RR capability status.
	RouteRefreshCapableSent bool `json:"route_refresh_capable_sent,omitempty"`
	// Received RR capability status.
	RouteRefreshCapableReceived bool `json:"route_refresh_capable_received,omitempty"`
	// Preserve forwarding state during BGP restart.
	GracefulRestartCapable bool `json:"graceful_restart_capable,omitempty"`
	// Sent graceful restart capability status.
	GracefulRestartCapableSent bool `json:"graceful_restart_capable_sent,omitempty"`
	// Received graceful restart capability status.
	GracefulRestartCapableReceived bool `json:"graceful_restart_capable_received,omitempty"`
	// Use the link-local IPv6 addresses to establish this peering.
	LinkLocalOnly bool `json:"link_local_only,omitempty"`
	// Sent link local only capability status.
	LinkLocalOnlySent bool `json:"link_local_only_sent,omitempty"`
	// Received link local only capability status.
	LinkLocalOnlyReceived bool `json:"link_local_only_received,omitempty"`
	// Whether this peering supports accumulated IGP metrics.
	AigpCapable bool `json:"aigp_capable,omitempty"`
	// Sent AIGP capability status.
	AigpCapableSent bool `json:"aigp_capable_sent,omitempty"`
	// Received AIGP capability status.
	AigpCapableReceived bool `json:"aigp_capable_received,omitempty"`
	// Negotiated AFI/SAFI for which prefixes (NLRIs) are exchanged.
	AfiSafis []string `json:"afi_safis,omitempty"`
	// Sent AFI/SAFIs.
	AfiSafisSent []string `json:"afi_safis_sent,omitempty"`
	// Received AFI/SAFIs.
	AfiSafisReceived []string `json:"afi_safis_received,omitempty"`
	// Negotiated AFI-SAFIs with extended next hop configuration.
	ExtendedNextHops []string `json:"extended_next_hops,omitempty"`
	// Sent extended next hop configuration.
	ExtendedNextHopsSent []string `json:"extended_next_hops_sent,omitempty"`
	// Received extended next hop configuration.
	ExtendedNextHopsReceived []string `json:"extended_next_hops_received,omitempty"`
	// AFI/SAFIs with additional paths enabled.
	AdditionalPaths []float64 `json:"additional_paths,omitempty"`
	// Send additional paths configuration.
	AdditionalPathsSend []float64 `json:"additional_paths_send,omitempty"`
	// Received additional paths configuration.
	AdditionalPathsReceived []float64 `json:"additional_paths_received,omitempty"`
	// Flag indicating whether 6PE is enabled. PE6 allows establishing IPv6 reachability over an IPv4/MPLS core.
	Var6peEnabled bool `json:"6pe_enabled,omitempty"`
	// Accept any remote AS number. If set to true, the remote AS is dynamically configured.
	AnyAs bool `json:"any_as,omitempty"`
}
