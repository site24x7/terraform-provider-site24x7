package api

// AttributeAlertGroup represents an attribute alert group in Site24x7.
type AttributeAlertGroup struct {
	_             struct{} `type:"structure"`
	GroupID       string   `json:"group_id,omitempty"`
	DisplayName   string   `json:"display_name"`
	AttributeList []int    `json:"attribute_list"`
}
