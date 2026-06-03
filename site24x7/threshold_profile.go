package site24x7

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
)

// SAMPLE POST JSON
// {
// 	"type": "URL",
//  "profile_name": "Website Threshold Profile",
// 	"profile_type": 1,
// 	"down_location_threshold": 1,
//  "website_content_modified": {
//     "severity": 2,
//     "value": false
//   },
// 	"website_content_changes": [
// 		{
// 		"severity": 2,
// 		"comparison_operator": 1,
// 		"value": 90
// 		}
// 	],
// "response_time_threshold": {
//     "primary": [
//       {
//         "severity": 2,
//         "comparison_operator": 1,
//         "strategy": 1,
//         "value": 10000,
//         "polls_check": 1
//       }
//     ],
//     "secondary": [
//       {
//         "severity": 2,
//         "comparison_operator": 1,
//         "strategy": 1,
//         "polls_check": 5,
//         "value": 10000
//       }
//     ]
//   },
// }

// SSL_CERT Threshold
// {
// 	"type": "SSL_CERT",
// 	"days_until_expiry": [
// 	  {
// 		"severity": 2,
// 		"comparison_operator": 2,
// 		"value": 30
// 	  },
// 	  {
// 		"severity": 3,
// 		"comparison_operator": 2,
// 		"value": 60
// 	  }
// 	],
// 	"profile_type": 1,
// 	"ssl_fingerprint_modified": {
// 	  "value": false
// 	},
// 	"profile_name": "SSL Certificate Threshold"
// }

var ThresholdProfileSchema = map[string]*schema.Schema{
	"profile_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display Name for the threshold profile",
	},
	"type": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Type of the monitor for which the threshold profile is being created.",
	},
	"profile_type": {
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      1,
		ValidateFunc: validation.IntInSlice([]int{1, 2}),
		Description:  "Static Threshold(1) or AI-based Threshold(2)",
	},
	"profile_type_name": {
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      1,
		ValidateFunc: validation.IntInSlice([]int{1, 2}),
		Description:  "Static Threshold(1) or AI-based Threshold(2)",
	},
	"down_location_threshold": {
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      3,
		ValidateFunc: validation.IntInSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}),
		Description:  "Triggers alert when the monitor is down from configured number of locations. Default value is '3'",
	},
	"website_content_modified": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Triggers alert when the website content is modified.",
	},
	"read_time_out": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{0, 2, 3}), //Down or Trouble or Critical
				},
				"value": {
					Type:     schema.TypeBool,
					Required: true,
				},
			},
		},
		Description: "Triggers alert when not receiving the website entire HTTP response within 30 seconds.",
	},
	"website_content_changes": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{0, 2, 3}), // Trouble or Critical
				},
				"comparison_operator": {
					Type:         schema.TypeInt,
					Optional:     true,
					Default:      1,
					ValidateFunc: validation.IntInSlice([]int{1}),
				},
				"value": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntAtMost(100),
				},
			},
		},
		Description: "Triggers alert when the website content changes by configured percentage.",
	},
	"primary_response_time_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity": {
					Type:     schema.TypeInt,
					Required: true,
					// ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					// 	log.Println("=============================== Validating primary_response_time_trouble_threshold : ", val)
					// 	v := val.(int)
					// 	if v != 2 {
					// 		errs = append(errs, fmt.Errorf("%q must be 2 for trouble threshold, got: %d", key, v))
					// 	}
					// 	return warns, errs
					// },
					ValidateFunc: validation.IntInSlice([]int{2}),
				},
				"comparison_operator": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5}),
				},
				"value": {
					Type:     schema.TypeInt,
					Required: true,
				},
				"strategy": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4}),
				},
				"polls_check": {
					Type:     schema.TypeInt,
					Required: true,
				},
			},
		},
		Description: "Response time trouble threshold for the primary monitoring location. Anomaly Enabled Attribute",
	},
	"primary_response_time_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{3}), // Critical
				},
				"comparison_operator": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5}),
				},
				"value": {
					Type:     schema.TypeInt,
					Required: true,
				},
				"strategy": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4}),
				},
				"polls_check": {
					Type:     schema.TypeInt,
					Required: true,
				},
			},
		},
		Description: "Response time critical threshold for the primary monitoring location. Anomaly Enabled Attribute",
	},
	"secondary_response_time_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{2}), // Trouble
				},
				"comparison_operator": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5}),
				},
				"value": {
					Type:     schema.TypeInt,
					Required: true,
				},
				"strategy": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4}),
				},
				"polls_check": {
					Type:     schema.TypeInt,
					Required: true,
				},
			},
		},
		Description: "Response time trouble threshold for the secondary monitoring location. Anomaly Enabled Attribute",
	},
	"secondary_response_time_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{3}), // Critical
				},
				"comparison_operator": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5}),
				},
				"value": {
					Type:     schema.TypeInt,
					Required: true,
				},
				"strategy": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4}),
				},
				"polls_check": {
					Type:     schema.TypeInt,
					Required: true,
				},
			},
		},
		Description: "Response time critical threshold for the secondary monitoring location. Anomaly Enabled Attribute",
	},
	// SSL_CERT monitor type attributes
	"ssl_cert_fingerprint_modified": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Triggers alert when the ssl certificate is modified.",
	},
	"ssl_cert_days_until_expiry_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{2}), // Trouble
				},
				"value": {
					Type:     schema.TypeInt,
					Required: true,
				},
			},
		},
		Description: "Triggers trouble alert before the SSL certificate expires within the configured number of days.",
	},
	"ssl_cert_days_until_expiry_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{2}), // Trouble
				},
				"value": {
					Type:     schema.TypeInt,
					Required: true,
				},
			},
		},
		Description: "Triggers critical alert before the SSL certificate expires within the configured number of days.",
	},
	// CRON monitor type attributes
	"cron_no_run_alert": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{2}), // Trouble
				},
				"value": {
					Type:     schema.TypeBool,
					Required: true,
				},
			},
		},
		Description: "Triggers Alert, if job does not start on schedule",
	},
	"cron_duration_alert": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"trouble": {
					Type:     schema.TypeInt,
					Required: true,
				},
			},
		},
		Description: "Generate Trouble Alert if not pinged for more than x seconds.",
	},

	// HEARTBEAT monitor type attributes
	"trouble_if_not_pinged_more_than": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Generate Trouble Alert if not pinged for more than x mins.",
	},
	"down_if_not_pinged_more_than": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Generate Down Alert if not pinged for more than x mins.",
	},
	"trouble_if_pinged_within": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Generate Trouble Alert if pinged within x mins",
	},

	// SERVER monitor type attributes
	// Standard threshold attributes - trouble thresholds (severity 2)
	"cpu_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers trouble alert if Server CPU Usage exceeds configured threshold.",
	},
	"cpu_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers critical alert if Server CPU Usage exceeds configured threshold.",
	},
	"memory_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers trouble alert if Server Memory Usage exceeds configured threshold.",
	},
	"memory_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers critical alert if Server Memory Usage exceeds configured threshold.",
	},
	"disk_usage_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers trouble alert if Server Disk utilization exceeds configured threshold.",
	},
	"disk_usage_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers critical alert if Server Disk utilization exceeds configured threshold.",
	},
	"disk_partition_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers trouble alert if disk partition threshold exceeds configured threshold.",
	},
	"disk_partition_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers critical alert if disk partition threshold exceeds configured threshold.",
	},
	"process_cpu_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers trouble alert if any process CPU usage exceeds configured threshold.",
	},
	"process_cpu_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers critical alert if any process CPU usage exceeds configured threshold.",
	},
	"process_memory_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers trouble alert if any process memory usage exceeds configured threshold.",
	},
	"process_memory_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers critical alert if any process memory usage exceeds configured threshold.",
	},
	"process_disk_usage_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers trouble alert if any process disk utilization exceeds configured threshold.",
	},
	"process_disk_usage_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers critical alert if any process disk utilization exceeds configured threshold.",
	},
	"process_handle_count_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers trouble alert if any process handle count exceeds configured threshold.",
	},
	"process_handle_count_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers critical alert if any process handle count exceeds configured threshold.",
	},
	"network_error_packet_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers trouble alert if network error packet exceeds configured threshold.",
	},
	"network_error_packet_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers critical alert if network error packet exceeds configured threshold.",
	},
	"ip_address_change_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers trouble alert if IP address validation fails.",
	},
	"ip_address_change_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers critical alert if IP address validation fails.",
	},
	"process_instance_count_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers trouble alert if any process instance count exceeds configured threshold.",
	},
	"process_instance_count_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers critical alert if any process instance count exceeds configured threshold.",
	},
	// Special attributes
	"process_down_alert": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"value":    {Type: schema.TypeBool, Required: true},
				"severity": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{0, 2, 3})},
			},
		},
		Description: "Triggers alert if process is down.",
	},
	"server_resource_down_alert": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"value":    {Type: schema.TypeBool, Required: true},
				"severity": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{0, 2, 3})},
			},
		},
		Description: "Notify if a resource check fails.",
	},
	"dc_alert": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"value":    {Type: schema.TypeBool, Required: true},
				"severity": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{0, 2, 3})},
			},
		},
		Description: "DC alert configuration.",
	},
	"disk_status_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"value":    {Type: schema.TypeBool, Required: true},
				"severity": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{0, 2, 3})},
			},
		},
		Description: "Disk status threshold alert configuration.",
	},
	"service_status_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"value":       {Type: schema.TypeBool, Required: true},
				"severity":    {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{0, 2, 3})},
				"strategy":    {Type: schema.TypeInt, Optional: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check": {Type: schema.TypeInt, Optional: true},
			},
		},
		Description: "Service status threshold alert configuration.",
	},
	"nw_status_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"value":    {Type: schema.TypeBool, Required: true},
				"severity": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{0, 2, 3})},
			},
		},
		Description: "Network status threshold alert configuration.",
	},
	// Windows specific
	"running_process_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Windows Specific. Triggers trouble alert if running process count exceeds configured threshold.",
	},
	"running_process_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Windows Specific. Triggers critical alert if running process count exceeds configured threshold.",
	},
	"total_service_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Windows Specific. Triggers trouble alert if total service count exceeds configured threshold.",
	},
	"total_service_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Windows Specific. Triggers critical alert if total service count exceeds configured threshold.",
	},
	"process_queue_length_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Windows Specific. Triggers trouble alert if processor queue length exceeds configured threshold.",
	},
	"process_queue_length_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Windows Specific. Triggers critical alert if processor queue length exceeds configured threshold.",
	},
	// Linux specific
	"system_load_1min_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Linux Specific. Triggers trouble alert if system load (1 min avg) exceeds configured threshold.",
	},
	"system_load_1min_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Linux Specific. Triggers critical alert if system load (1 min avg) exceeds configured threshold.",
	},
	"system_load_5min_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Linux Specific. Triggers trouble alert if system load (5 min avg) exceeds configured threshold.",
	},
	"system_load_5min_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Linux Specific. Triggers critical alert if system load (5 min avg) exceeds configured threshold.",
	},
	"system_load_15min_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Linux Specific. Triggers trouble alert if system load (15 min avg) exceeds configured threshold.",
	},
	"system_load_15min_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Linux Specific. Triggers critical alert if system load (15 min avg) exceeds configured threshold.",
	},
	"process_running_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Linux Specific. Triggers trouble alert if running process count exceeds configured threshold.",
	},
	"process_running_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Linux Specific. Triggers critical alert if running process count exceeds configured threshold.",
	},
	"total_process_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Linux Specific. Triggers trouble alert if total process count exceeds configured threshold.",
	},
	"total_process_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Linux Specific. Triggers critical alert if total process count exceeds configured threshold.",
	},
	"blocked_process_trouble_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{2})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Linux Specific. Triggers trouble alert if blocked process count exceeds configured threshold.",
	},
	"blocked_process_critical_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{3})},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"value":               {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Linux Specific. Triggers critical alert if blocked process count exceeds configured threshold.",
	},
	// Special format attributes
	"disk_used_size": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"trouble":             {Type: schema.TypeInt, Required: true},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"unit_id":             {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers alert if disk partition usage threshold exceeds configured threshold.",
	},
	"disk_free_size": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"trouble":             {Type: schema.TypeInt, Required: true},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"unit_id":             {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers alert if disk partition free space exceeds configured threshold.",
	},
	"server_uptime": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"trouble":             {Type: schema.TypeInt, Required: true},
				"comparison_operator": {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5})},
				"polls_check":         {Type: schema.TypeInt, Required: true},
				"strategy":            {Type: schema.TypeInt, Required: true, ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4})},
				"unit_id":             {Type: schema.TypeInt, Required: true},
			},
		},
		Description: "Triggers alert if uptime of the server exceeds configured threshold.",
	},
}

func ResourceSite24x7ThresholdProfile() *schema.Resource {
	return &schema.Resource{
		Create: thresholdProfileCreate,
		Read:   thresholdProfileRead,
		Update: thresholdProfileUpdate,
		Delete: thresholdProfileDelete,
		Exists: thresholdProfileExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: ThresholdProfileSchema,
	}
}

func thresholdProfileCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	thresholdProfile := resourceDataToThresholdProfile(d)

	thresholdProfile, err := client.ThresholdProfiles().Create(thresholdProfile)
	if err != nil {
		return err
	}

	d.SetId(thresholdProfile.ProfileID)

	return nil
}

func thresholdProfileRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	thresholdProfile, err := client.ThresholdProfiles().Get(d.Id())
	if err != nil {
		return err
	}

	updateThresholdProfileResourceData(d, thresholdProfile)

	return nil
}

func thresholdProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	thresholdProfile := resourceDataToThresholdProfile(d)
	thresholdProfile, err := client.ThresholdProfiles().Update(thresholdProfile)
	if err != nil {
		return err
	}

	d.SetId(thresholdProfile.ProfileID)

	return nil
}

func thresholdProfileDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	err := client.ThresholdProfiles().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func thresholdProfileExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(Client)

	_, err := client.ThresholdProfiles().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToThresholdProfile(d *schema.ResourceData) *api.ThresholdProfile {

	monitorType := d.Get("type").(string)

	thresholdProfileToReturn := &api.ThresholdProfile{
		ProfileID:   d.Id(),
		ProfileName: d.Get("profile_name").(string),
		Type:        d.Get("type").(string),
		ProfileType: d.Get("profile_type").(int),
	}

	// SSL_CERT attributes
	if monitorType == string(api.SSL_CERT) {
		setSSLCertificateAttributes(d, thresholdProfileToReturn)
	} else if monitorType == string(api.HEARTBEAT) {
		setHeartBeatAttributes(d, thresholdProfileToReturn)
	} else if monitorType == string(api.CRON) {
		setCronAttributes(d, thresholdProfileToReturn)
	} else if monitorType == string(api.SERVER) {
		setServerAttributes(d, thresholdProfileToReturn)
	} else {
		setCommonAttributes(d, thresholdProfileToReturn)
	}

	return thresholdProfileToReturn
}

// Called during read and sets thresholdProfile in API response to ResourceData
func updateThresholdProfileResourceData(d *schema.ResourceData, thresholdProfile *api.ThresholdProfile) {

	monitorType := thresholdProfile.Type
	d.Set("profile_name", thresholdProfile.ProfileName)
	d.Set("type", thresholdProfile.Type)
	d.Set("profile_type", thresholdProfile.ProfileType)

	if monitorType == string(api.SSL_CERT) {
		setSSLCertificateResourceData(d, thresholdProfile)
	} else if monitorType == string(api.HEARTBEAT) {
		setHeartBeatResourceData(d, thresholdProfile)
	} else if monitorType == string(api.CRON) {
		setCronResourceData(d, thresholdProfile)
	} else if monitorType == string(api.SERVER) {
		setServerResourceData(d, thresholdProfile)
	} else {
		setCommonResourceData(d, thresholdProfile)
	}

}

func setSSLCertificateAttributes(d *schema.ResourceData, thresholdProfile *api.ThresholdProfile) {
	// SSL Certificate days until expiry
	var sslCertDaysUntilExpiry []map[string]interface{}
	if sslCertDaysUntilExp, ok := d.GetOk("ssl_cert_days_until_expiry_trouble_threshold"); ok {
		sslCertDaysUntilExpiryMap := sslCertDaysUntilExp.(map[string]interface{})
		sslCertDaysUntilExpiryMap["severity"] = "2"
		sslCertDaysUntilExpiry = append(sslCertDaysUntilExpiry, sslCertDaysUntilExpiryMap)
	}
	if sslCertDaysUntilExpCritical, ok := d.GetOk("ssl_cert_days_until_expiry_critical_threshold"); ok {
		sslCertDaysUntilExpiryCriticalMap := sslCertDaysUntilExpCritical.(map[string]interface{})
		sslCertDaysUntilExpiryCriticalMap["severity"] = "3"
		sslCertDaysUntilExpiry = append(sslCertDaysUntilExpiry, sslCertDaysUntilExpiryCriticalMap)
	}
	thresholdProfile.SSLCertificateDaysUntilExpiry = sslCertDaysUntilExpiry

	// SSL certificate fingerprint modified
	if sslCertFingerprintModified, ok := d.GetOk("ssl_cert_fingerprint_modified"); ok {
		sslCertFingerprintModifiedMap := make(map[string]interface{})
		sslCertFingerprintModifiedMap["value"] = sslCertFingerprintModified.(bool)
		thresholdProfile.SSLCertificateFingerprintModified = sslCertFingerprintModifiedMap
	}
}

func setSSLCertificateResourceData(d *schema.ResourceData, thresholdProfile *api.ThresholdProfile) {
	if len(thresholdProfile.SSLCertificateDaysUntilExpiry) > 0 {
		for _, sslCertDaysUntilExpiry := range thresholdProfile.SSLCertificateDaysUntilExpiry {
			sslCertDaysUntilExpiryMap := sslCertDaysUntilExpiry
			if secondarySeverity, ok := sslCertDaysUntilExpiryMap["severity"]; ok {
				if secondarySeverity == 2 {
					d.Set("ssl_cert_days_until_expiry_trouble_threshold", sslCertDaysUntilExpiryMap)
				}
				if secondarySeverity == 3 {
					d.Set("ssl_cert_days_until_expiry_critical_threshold", sslCertDaysUntilExpiryMap)
				}
			}
		}
	}

	if thresholdProfile.SSLCertificateFingerprintModified != nil {
		d.Set("ssl_cert_fingerprint_modified", thresholdProfile.SSLCertificateFingerprintModified["value"].(bool))
	}
}

func setCronAttributes(d *schema.ResourceData, thresholdProfile *api.ThresholdProfile) {

	if cronNoRunAlert, ok := d.GetOk("cron_no_run_alert"); ok {
		thresholdProfile.CronNoRunAlert = cronNoRunAlert.(map[string]interface{})
	}

	if cronDurationAlert, ok := d.GetOk("cron_duration_alert"); ok {
		thresholdProfile.CronDurationAlert = cronDurationAlert.(map[string]interface{})
	}
}

func setCronResourceData(d *schema.ResourceData, thresholdProfile *api.ThresholdProfile) {
	cronNoRunAlertMap := make(map[string]interface{})
	cronNoRunAlertMap["severity"] = int(thresholdProfile.CronNoRunAlert["severity"].(float64))
	cronNoRunAlertMap["value"] = thresholdProfile.CronNoRunAlert["value"].(bool)
	d.Set("cron_no_run_alert", cronNoRunAlertMap)

	cronDurationAlertMap := make(map[string]interface{})
	if thresholdProfile.CronDurationAlert["trouble"] == nil {
		d.Set("cron_duration_alert", cronDurationAlertMap)
	} else {
		cronDurationAlertMap["trouble"] = int(thresholdProfile.CronDurationAlert["trouble"].(float64))
		d.Set("cron_duration_alert", cronDurationAlertMap)
	}
}

func setHeartBeatAttributes(d *schema.ResourceData, thresholdProfile *api.ThresholdProfile) {
	troubleIfNotPingedMoreThanMap := make(map[string]interface{})
	if troubleIfNotPingedMoreThan, ok := d.GetOk("trouble_if_not_pinged_more_than"); ok {
		troubleIfNotPingedMoreThanMap["comparison_operator"] = 1
		troubleIfNotPingedMoreThanMap["trouble"] = troubleIfNotPingedMoreThan
		troubleIfNotPingedMoreThanMap["strategy"] = 1
		troubleIfNotPingedMoreThanMap["polls_check"] = 5
		thresholdProfile.TroubleIfNotPingedMoreThan = troubleIfNotPingedMoreThanMap
	}
	downIfNotPingedMoreThanMap := make(map[string]interface{})
	if downIfNotPingedMoreThan, ok := d.GetOk("down_if_not_pinged_more_than"); ok {
		downIfNotPingedMoreThanMap["comparison_operator"] = 1
		downIfNotPingedMoreThanMap["trouble"] = downIfNotPingedMoreThan
		downIfNotPingedMoreThanMap["strategy"] = 1
		downIfNotPingedMoreThanMap["polls_check"] = 5
		thresholdProfile.DownIfNotPingedMoreThan = downIfNotPingedMoreThanMap
	}
	troubleIfPingedWithinMap := make(map[string]interface{})
	if troubleIfPingedWithin, ok := d.GetOk("trouble_if_pinged_within"); ok {
		troubleIfPingedWithinMap["comparison_operator"] = 1
		troubleIfPingedWithinMap["trouble"] = troubleIfPingedWithin
		troubleIfPingedWithinMap["strategy"] = 1
		troubleIfPingedWithinMap["polls_check"] = 5
		thresholdProfile.TroubleIfPingedWithin = troubleIfPingedWithinMap
	}
}

func setHeartBeatResourceData(d *schema.ResourceData, thresholdProfile *api.ThresholdProfile) {
	if thresholdProfile.TroubleIfNotPingedMoreThan != nil {
		if troubleIfNotPingedMoreThan, ok := thresholdProfile.TroubleIfNotPingedMoreThan["trouble"]; ok {
			d.Set("trouble_if_not_pinged_more_than", troubleIfNotPingedMoreThan)

		}
	}

	if thresholdProfile.DownIfNotPingedMoreThan != nil {
		if downIfNotPingedMoreThan, ok := thresholdProfile.DownIfNotPingedMoreThan["trouble"]; ok {
			d.Set("down_if_not_pinged_more_than", downIfNotPingedMoreThan)

		}
	}

	if thresholdProfile.TroubleIfPingedWithin != nil {
		if troubleIfPingedWithin, ok := thresholdProfile.TroubleIfPingedWithin["trouble"]; ok {
			d.Set("trouble_if_pinged_within", troubleIfPingedWithin)

		}
	}

}

func setCommonAttributes(d *schema.ResourceData, thresholdProfile *api.ThresholdProfile) {
	thresholdProfile.DownLocationThreshold = d.Get("down_location_threshold").(int)
	thresholdProfile.WebsiteContentModified = d.Get("website_content_modified").(bool)

	// Website content changes
	var websiteContentChanges []map[string]interface{}
	if contentChangesList, ok := d.GetOk("website_content_changes"); ok {
		for _, urlContentChanges := range contentChangesList.([]interface{}) {
			urlContentChangesMap, ok := urlContentChanges.(map[string]interface{})
			if ok {
				websiteContentChanges = append(websiteContentChanges, urlContentChangesMap)
			}
		}
	}
	thresholdProfile.WebsiteContentChanges = websiteContentChanges

	if readTimeOut, ok := d.GetOk("read_time_out"); ok {
		thresholdProfile.ReadTimeOut = readTimeOut.(map[string]interface{})
	}

	// Response Time Threshold
	var setResponseTimeThresholdMap bool
	responseTimeThresholdMap := make(map[string]interface{})
	var primaryThresholdList []map[string]interface{}
	var secondaryThresholdList []map[string]interface{}
	// Primary Threshold
	if primaryResponseTimeTroubleThreshold, ok := d.GetOk("primary_response_time_trouble_threshold"); ok {
		primaryResponseTimeTroubleThresholdMap := primaryResponseTimeTroubleThreshold.(map[string]interface{})
		primaryResponseTimeTroubleThresholdMap["severity"] = "2"
		primaryThresholdList = append(primaryThresholdList, primaryResponseTimeTroubleThresholdMap)
	}
	if primaryResponseTimeCriticalThreshold, ok := d.GetOk("primary_response_time_critical_threshold"); ok {
		primaryResponseTimeCriticalThresholdMap := primaryResponseTimeCriticalThreshold.(map[string]interface{})
		primaryResponseTimeCriticalThresholdMap["severity"] = "3"
		primaryThresholdList = append(primaryThresholdList, primaryResponseTimeCriticalThresholdMap)
	}
	// Secondary Threshold
	if secondaryResponseTimeTroubleThreshold, ok := d.GetOk("secondary_response_time_trouble_threshold"); ok {
		secondaryResponseTimeTroubleThresholdMap := secondaryResponseTimeTroubleThreshold.(map[string]interface{})
		secondaryResponseTimeTroubleThresholdMap["severity"] = "2"
		secondaryThresholdList = append(secondaryThresholdList, secondaryResponseTimeTroubleThresholdMap)
	}
	if secondaryResponseTimeCriticalThreshold, ok := d.GetOk("secondary_response_time_critical_threshold"); ok {
		secondaryResponseTimeCriticalThresholdMap := secondaryResponseTimeCriticalThreshold.(map[string]interface{})
		secondaryResponseTimeCriticalThresholdMap["severity"] = "3"
		secondaryThresholdList = append(secondaryThresholdList, secondaryResponseTimeCriticalThresholdMap)
	}
	if len(primaryThresholdList) > 0 {
		responseTimeThresholdMap["primary"] = primaryThresholdList
		setResponseTimeThresholdMap = true
	}
	if len(secondaryThresholdList) > 0 {
		responseTimeThresholdMap["secondary"] = secondaryThresholdList
		setResponseTimeThresholdMap = true
	}
	if setResponseTimeThresholdMap {
		thresholdProfile.ResponseTimeThreshold = responseTimeThresholdMap
	}
}

func setCommonResourceData(d *schema.ResourceData, thresholdProfile *api.ThresholdProfile) {
	d.Set("down_location_threshold", thresholdProfile.DownLocationThreshold)
	d.Set("website_content_modified", thresholdProfile.WebsiteContentModified)
	d.Set("website_content_changes", thresholdProfile.WebsiteContentChanges)
	d.Set("read_time_out", thresholdProfile.ReadTimeOut)
	// Response Time Primary Threshold
	if primaryThreshold, ok := thresholdProfile.ResponseTimeThreshold["primary"]; ok {
		primaryThresholdList := primaryThreshold.([]interface{})
		if len(primaryThresholdList) > 0 {
			for _, primaryThresh := range primaryThresholdList {
				primaryThresholdMap := primaryThresh.(map[string]interface{})
				if primarySeverity, ok := primaryThresholdMap["severity"]; ok {
					if primarySeverity == 2 {
						d.Set("primary_response_time_trouble_threshold", primaryThresholdMap)
					}
					if primarySeverity == 3 {
						d.Set("primary_response_time_critical_threshold", primaryThresholdMap)
					}
				}
			}
		}

	}
	// Response Time Secondary Threshold
	if secondaryThreshold, ok := thresholdProfile.ResponseTimeThreshold["secondary"]; ok {
		secondaryThresholdList := secondaryThreshold.([]interface{})
		if len(secondaryThresholdList) > 0 {
			for _, secondaryThresh := range secondaryThresholdList {
				secondaryThresholdMap := secondaryThresh.(map[string]interface{})
				if secondarySeverity, ok := secondaryThresholdMap["severity"]; ok {
					if secondarySeverity == 2 {
						d.Set("secondary_response_time_trouble_threshold", secondaryThresholdMap)
					}
					if secondarySeverity == 3 {
						d.Set("secondary_response_time_critical_threshold", secondaryThresholdMap)
					}
				}
			}
		}
	}
}

// serverThresholdAttributeMapping maps API field names to their trouble/critical terraform attribute names
var serverThresholdAttributeMapping = []struct {
	apiField     string
	troubleAttr  string
	criticalAttr string
	getField     func(*api.ThresholdProfile) []map[string]interface{}
	setField     func(*api.ThresholdProfile, []map[string]interface{})
}{
	{"cpu_threshold", "cpu_trouble_threshold", "cpu_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.CpuThreshold },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.CpuThreshold = v }},
	{"memory_threshold", "memory_trouble_threshold", "memory_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.MemoryThreshold },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.MemoryThreshold = v }},
	{"disk_usage_threshold", "disk_usage_trouble_threshold", "disk_usage_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.DiskUsageThreshold },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.DiskUsageThreshold = v }},
	{"disk_partition_threshold", "disk_partition_trouble_threshold", "disk_partition_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.DiskPartitionThreshold },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.DiskPartitionThreshold = v }},
	{"process_cpu_threshold", "process_cpu_trouble_threshold", "process_cpu_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.ProcessCpuThreshold },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.ProcessCpuThreshold = v }},
	{"process_memory_threshold", "process_memory_trouble_threshold", "process_memory_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.ProcessMemoryThreshold },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.ProcessMemoryThreshold = v }},
	{"process_disk_usage_threshold", "process_disk_usage_trouble_threshold", "process_disk_usage_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.ProcessDiskUsageThreshold },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.ProcessDiskUsageThreshold = v }},
	{"process_handle_count_threshold", "process_handle_count_trouble_threshold", "process_handle_count_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.ProcessHandleCountThreshold },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.ProcessHandleCountThreshold = v }},
	{"network_error_packet_threshold", "network_error_packet_trouble_threshold", "network_error_packet_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.NetworkErrorPacketThreshold },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.NetworkErrorPacketThreshold = v }},
	{"ip_address_change", "ip_address_change_trouble_threshold", "ip_address_change_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.IpAddressChange },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.IpAddressChange = v }},
	{"process_instance_count_threshold", "process_instance_count_trouble_threshold", "process_instance_count_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.ProcessInstanceCountThreshold },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.ProcessInstanceCountThreshold = v }},
	// Windows specific
	{"running_process", "running_process_trouble_threshold", "running_process_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.RunningProcess },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.RunningProcess = v }},
	{"total_service", "total_service_trouble_threshold", "total_service_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.TotalService },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.TotalService = v }},
	{"process_queue_length", "process_queue_length_trouble_threshold", "process_queue_length_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.ProcessQueueLength },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.ProcessQueueLength = v }},
	// Linux specific
	{"system_load_1min", "system_load_1min_trouble_threshold", "system_load_1min_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.SystemLoad1Min },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.SystemLoad1Min = v }},
	{"system_load_5min", "system_load_5min_trouble_threshold", "system_load_5min_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.SystemLoad5Min },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.SystemLoad5Min = v }},
	{"system_load_15min", "system_load_15min_trouble_threshold", "system_load_15min_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.SystemLoad15Min },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.SystemLoad15Min = v }},
	{"process_running", "process_running_trouble_threshold", "process_running_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.ProcessRunning },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.ProcessRunning = v }},
	{"total_process", "total_process_trouble_threshold", "total_process_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.TotalProcess },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.TotalProcess = v }},
	{"blocked_process", "blocked_process_trouble_threshold", "blocked_process_critical_threshold",
		func(t *api.ThresholdProfile) []map[string]interface{} { return t.BlockedProcess },
		func(t *api.ThresholdProfile, v []map[string]interface{}) { t.BlockedProcess = v }},
}

func setServerAttributes(d *schema.ResourceData, thresholdProfile *api.ThresholdProfile) {
	// Standard threshold attributes with trouble/critical split - sent as flat arrays to the API
	for _, mapping := range serverThresholdAttributeMapping {
		troubleMap, hasTrouble := d.GetOk(mapping.troubleAttr)
		criticalMap, hasCritical := d.GetOk(mapping.criticalAttr)

		if hasTrouble || hasCritical {
			var thresholdList []map[string]interface{}
			if hasTrouble {
				thresholdList = append(thresholdList, troubleMap.(map[string]interface{}))
			}
			if hasCritical {
				thresholdList = append(thresholdList, criticalMap.(map[string]interface{}))
			}
			mapping.setField(thresholdProfile, thresholdList)
		}
	}

	// Process down alert
	if processDownAlert, ok := d.GetOk("process_down_alert"); ok {
		thresholdProfile.ProcessDownAlert = processDownAlert.(map[string]interface{})
	}

	// Server resource down alert
	if serverResourceDownAlert, ok := d.GetOk("server_resource_down_alert"); ok {
		thresholdProfile.ServerResourceDownAlert = serverResourceDownAlert.(map[string]interface{})
	}

	// DC alert
	if dcAlert, ok := d.GetOk("dc_alert"); ok {
		thresholdProfile.DcAlert = dcAlert.(map[string]interface{})
	}

	// Disk status threshold
	if diskStatusThreshold, ok := d.GetOk("disk_status_threshold"); ok {
		thresholdProfile.DiskStatusThreshold = diskStatusThreshold.(map[string]interface{})
	}

	// Service status threshold
	if serviceStatusThreshold, ok := d.GetOk("service_status_threshold"); ok {
		thresholdProfile.ServiceStatusThreshold = serviceStatusThreshold.(map[string]interface{})
	}

	// Network status threshold
	if nwStatusThreshold, ok := d.GetOk("nw_status_threshold"); ok {
		thresholdProfile.NwStatusThreshold = nwStatusThreshold.(map[string]interface{})
	}

	// Disk used size
	if diskUsedSize, ok := d.GetOk("disk_used_size"); ok {
		thresholdProfile.DiskUsedSize = diskUsedSize.(map[string]interface{})
	}

	// Disk free size
	if diskFreeSize, ok := d.GetOk("disk_free_size"); ok {
		thresholdProfile.DiskFreeSize = diskFreeSize.(map[string]interface{})
	}

	// Server uptime
	if serverUptime, ok := d.GetOk("server_uptime"); ok {
		thresholdProfile.ServerUptime = serverUptime.(map[string]interface{})
	}
}

func setServerResourceData(d *schema.ResourceData, thresholdProfile *api.ThresholdProfile) {
	// Standard threshold attributes with trouble/critical split
	// API returns flat arrays: [{severity:2,...}, {severity:0,...}]
	for _, mapping := range serverThresholdAttributeMapping {
		fieldValue := mapping.getField(thresholdProfile)
		if len(fieldValue) == 0 {
			continue
		}

		for _, thresholdMap := range fieldValue {
			if severity, ok := thresholdMap["severity"]; ok {
				switch severity {
				case float64(2), 2, "2":
					d.Set(mapping.troubleAttr, thresholdMap)
				case float64(3), 3, "3", float64(0), 0, "0":
					d.Set(mapping.criticalAttr, thresholdMap)
				}
			}
		}
	}

	// Process down alert
	if thresholdProfile.ProcessDownAlert != nil {
		d.Set("process_down_alert", thresholdProfile.ProcessDownAlert)
	}

	// Server resource down alert
	if thresholdProfile.ServerResourceDownAlert != nil {
		d.Set("server_resource_down_alert", thresholdProfile.ServerResourceDownAlert)
	}

	// DC alert
	if thresholdProfile.DcAlert != nil {
		d.Set("dc_alert", thresholdProfile.DcAlert)
	}

	// Disk status threshold
	if thresholdProfile.DiskStatusThreshold != nil {
		d.Set("disk_status_threshold", thresholdProfile.DiskStatusThreshold)
	}

	// Service status threshold
	if thresholdProfile.ServiceStatusThreshold != nil {
		d.Set("service_status_threshold", thresholdProfile.ServiceStatusThreshold)
	}

	// Network status threshold
	if thresholdProfile.NwStatusThreshold != nil {
		d.Set("nw_status_threshold", thresholdProfile.NwStatusThreshold)
	}

	// Disk used size
	if thresholdProfile.DiskUsedSize != nil {
		d.Set("disk_used_size", thresholdProfile.DiskUsedSize)
	}

	// Disk free size
	if thresholdProfile.DiskFreeSize != nil {
		d.Set("disk_free_size", thresholdProfile.DiskFreeSize)
	}

	// Server uptime
	if thresholdProfile.ServerUptime != nil {
		d.Set("server_uptime", thresholdProfile.ServerUptime)
	}
}
