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

// BgpState : BGP FSM state.
type BgpState string

// List of BgpState
const (
	CONNECT_BgpState      BgpState = "CONNECT"
	IDLE_BgpState         BgpState = "IDLE"
	OPEN_CONFIRM_BgpState BgpState = "OPEN_CONFIRM"
	OPEN_SENT_BgpState    BgpState = "OPEN_SENT"
	ACTIVE_BgpState       BgpState = "ACTIVE"
	ESTABLISHED_BgpState  BgpState = "ESTABLISHED"
)
