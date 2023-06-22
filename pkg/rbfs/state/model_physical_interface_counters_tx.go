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

// Transmitted data counters.
type PhysicalInterfaceCountersTx struct {
	// Total number of transmitted bytes.
	BytesSent int `json:"bytes_sent,omitempty"`
	// Total number of transmitted packets.
	PacketsSent int `json:"packets_sent,omitempty"`
	// Total number of transmitted unicast packets.
	UnicastPacketsSent int `json:"unicast_packets_sent,omitempty"`
	// Total number of transmitted multicast packets.
	MulticastPacketsSent int `json:"multicast_packets_sent,omitempty"`
	// Total number of transmitted broadcast packets.
	BroadcastPacketsSent int `json:"broadcast_packets_sent,omitempty"`
	// Total number of dropped packets.
	PacketsDropped int `json:"packets_dropped,omitempty"`
}
