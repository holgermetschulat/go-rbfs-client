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

// An active subscriber session.
type SubscriberSession struct {
	// The current subscriber FSM state.
	SubscriberState string `json:"subscriber_state,omitempty"`
	// The timestamp of the last FSM state transition.
	LastStateTransition string `json:"last_state_transition,omitempty"`
	// Reason why the subscriber session got terminated.
	SubscriberTerminationReason string `json:"subscriber_termination_reason,omitempty"`
	SubscriberId                int32  `json:"subscriber_id,omitempty"`
	// String subscriber_id representation for I-JSON compliance.
	SubscriberIdStr string `json:"subscriber_id_str,omitempty"`
	IfpName         string `json:"ifp_name,omitempty"`
	// The MAC address of the physical interface  the subscriber is connected to.
	IfpMac string `json:"ifp_mac,omitempty"`
	// The inner VLAN (C-tag) associated with the subscriber.
	InnerVlan int32 `json:"inner_vlan,omitempty"`
	// The outer VLAN (S-tag) associated with the subscriber.
	OuterVlan int32 `json:"outer_vlan,omitempty"`
	// The client MAC address.
	ClientMac string `json:"client_mac,omitempty"`
	// The subscriber's username.
	SubscriberUserName string `json:"subscriber_user_name,omitempty"`
	// The subscriber's Agent-Remote-ID.
	AgentRemoteId string `json:"agent_remote_id,omitempty"`
	// The subscriber's Agent-Circuit-ID.
	AgentCircuitId string `json:"agent_circuit_id,omitempty"`
	// The name of the subscriber's logical interface.
	IflName    string                 `json:"ifl_name,omitempty"`
	AccessType *SubscriberAccessType  `json:"access_type,omitempty"`
	Ipv4       *SubscriberSessionIpv4 `json:"ipv4,omitempty"`
	Ipv6       *SubscriberSessionIpv6 `json:"ipv6,omitempty"`
	// The MTU level assigned to the subscriber.
	L3Mtu int32 `json:"l3_mtu,omitempty"`
	// The AAA profile name.
	AaaProfileName string `json:"aaa_profile_name,omitempty"`
	// The access profile name.
	AccessProfileName string `json:"access_profile_name,omitempty"`
	// The service profile name.
	ServiceProfileName string                    `json:"service_profile_name,omitempty"`
	Accounting         *SubscriberAccountingInfo `json:"accounting,omitempty"`
	Qos                *SubscriberQosProfile     `json:"qos,omitempty"`
	L2tp               *SubscriberSessionL2tp    `json:"l2tp,omitempty"`
	Pppoe              *SubscriberSessionPppoe   `json:"pppoe,omitempty"`
}
