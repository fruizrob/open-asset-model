package network

// RIROrganization represents an organization that is a member
// of a Regional Internet Registry (RIR). An RIR is an organization
// that manages the allocation and registration of IP addresses
// and AS numbers within a specific region of the world.
type RIROrganization struct {
	// Name of the RIR organization, such as "Google LLC".
	Name string `json:"name"`

	// RIRId is the unique identifier of the RIR organization,
	// such as "GOGL". Leave empty if unknown.
	RIRId string `json:"rir_id"`

	// RIR is the name of the RIR that this organization belongs to.
	// It should be one of "AFRINIC", "APNIC", "ARIN", "LACNIC", "RIPE".
	// Leave empty if unknown.
	RIR string `json:"rir"`
}
