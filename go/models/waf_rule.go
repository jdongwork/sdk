package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// WafRule waf rule
// swagger:model WafRule
type WafRule struct {

	// Enable or disable WAF Rule Group. Field introduced in 17.2.1.
	// Required: true
	Enable bool `json:"enable"`

	// Exclude list for the WAF rule. The fields in the exclude list entry are logically and'ed to deduce the exclusion criteria. If there are multiple excludelist entries, it will be 'logical or' of them. Field introduced in 17.2.3.
	ExcludeList []*WafExcludeListEntry `json:"exclude_list,omitempty"`

	//  Field introduced in 17.2.1.
	// Required: true
	Index int32 `json:"index"`

	// User-friendly optional name for a rule. Field introduced in 17.2.1.
	Name string `json:"name,omitempty"`

	// Rule as per Modsec language. Field introduced in 17.2.1.
	// Required: true
	Rule string `json:"rule"`

	// Identifier (id) for a rule per Modsec language. All SecRule and SecAction directives require an id. It is extracted from the id action in a ModSec rule. Rules within a single WAF Policy are required to have unique rule_ids. Field introduced in 17.2.2.
	RuleID string `json:"rule_id,omitempty"`
}
