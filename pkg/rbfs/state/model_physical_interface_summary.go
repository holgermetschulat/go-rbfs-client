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

import (
	"time"
)

// A physical interface
type PhysicalInterfaceSummary struct {
	IfpName string `json:"ifp_name,omitempty"`
	// The physical interface alias.
	IfpAlias string `json:"ifp_alias,omitempty"`
	// The MAC address of the physical interface.
	MacAddress string `json:"mac_address,omitempty"`
	// The configured speed limit.  The maximum speed can be read from the bandwidth attribute.
	Speed string `json:"speed,omitempty"`
	// The maximum bandwidth available.
	Bandwidth string `json:"bandwidth,omitempty"`
	// The duplex mode.
	DuplexMode string `json:"duplex_mode,omitempty"`
	// The configured maximum transfer unit (MTU) size in bytes.
	MtuSize int `json:"mtu_size,omitempty"`
	// The operational interface state.
	OperationalState string `json:"operational_state,omitempty"`
	// The administrative interface state.
	AdministrativeState string `json:"administrative_state,omitempty"`
	// Timestamp since when this interface is UP.
	UpTime time.Time `json:"up_time,omitempty"`
	// The logical interfaces defined on this physical interface.
	Ifls        []LogicalInterface         `json:"ifls,omitempty"`
	Lag         *LinkAggregationGroup      `json:"lag,omitempty"`
	IfpCounters *PhysicalInterfaceCounters `json:"ifp_counters,omitempty"`
}
