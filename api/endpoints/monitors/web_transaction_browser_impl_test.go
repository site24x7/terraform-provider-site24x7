package monitors

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebTransactionBrowserMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create web transaction browser monitor",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitors",
			ExpectedBody: validation.Fixture(t, "requests/create_web_transaction_browser_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				webTransactionBrowserMonitor := &api.WebTransactionBrowserMonitor{
					DisplayName:        "RBM-Terraform",
					Type:               string(api.REALBROWSER),
					BaseURL:            "https://www.example.com/",
					AsyncDCEnabled:     false,
					CheckFrequency:     "15",
					IgnoreCertError:    false,
					IPType:             0,
					SeleniumScript:     "Script for the monitor",
					ScriptType:         "txt",
					ThresholdProfileId: "789",
					PageLoadTime:       0,
					PerformAutomation:  false,
					Resolution:         "1600,900",
					LocationProfileID:  "456",
					MonitorGroups:      []string{"234", "567"},
					UserGroupIDs:       []string{"123", "456"},
					TagIDs:             []string{"123"},
				}

				_, err := NewWebTransactionBrowserMonitors(c).Create(webTransactionBrowserMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get web transaction browser monitor",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors/897654345678",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_web_transaction_browser_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				webTransactionBrowserMonitor, err := NewWebTransactionBrowserMonitors(c).Get("897654345678")
				require.NoError(t, err)

				expected := &api.WebTransactionBrowserMonitor{
					MonitorID:          "897654345678",
					DisplayName:        "RBM-Terraform",
					Type:               string(api.REALBROWSER),
					BaseURL:            "https://www.example.com/",
					SeleniumScript:     "Script for the monitor",
					ScriptType:         "txt",
					ThresholdProfileId: "789",
					Resolution:         "1600,900",
					LocationProfileID:  "456",
					MonitorGroups:      []string{"234", "567"},
					UserGroupIDs:       []string{"123", "456"},
					TagIDs:             []string{"123"},
				}

				assert.Equal(t, expected, webTransactionBrowserMonitor)
			},
		},
		{
			Name:         "list eb transaction browser monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_web_transaction_browser_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				webTransactionBrowserMonitor, err := NewWebTransactionBrowserMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.WebTransactionBrowserMonitor{
					{
						MonitorID:          "897654345678",
						DisplayName:        "RBM-Terraform",
						Type:               string(api.REALBROWSER),
						BaseURL:            "https://www.example.com/",
						SeleniumScript:     "Script for the monitor",
						ScriptType:         "txt",
						ThresholdProfileId: "789",
						Resolution:         "1600,900",
						LocationProfileID:  "456",
						MonitorGroups:      []string{"234", "567"},
						UserGroupIDs:       []string{"123", "456"},
						TagIDs:             []string{"123"},
					},
					{
						MonitorID:          "987554574575",
						DisplayName:        "RBM-Terraform",
						Type:               string(api.REALBROWSER),
						BaseURL:            "https://www.example.com/",
						SeleniumScript:     "Script for the monitor",
						ScriptType:         "txt",
						ThresholdProfileId: "789",
						Resolution:         "1600,900",
						LocationProfileID:  "456",
						MonitorGroups:      []string{"234", "567"},
						UserGroupIDs:       []string{"123", "456"},
						TagIDs:             []string{"123"},
					},
				}

				assert.Equal(t, expected, webTransactionBrowserMonitor)
			},
		},
		{
			Name:         "update web transaction browser monitor",
			ExpectedVerb: "PUT",
			ExpectedPath: "/monitors/897654345678",
			ExpectedBody: validation.Fixture(t, "requests/update_web_transaction_browser_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				webTransactionBrowserMonitor := &api.WebTransactionBrowserMonitor{
					MonitorID:          "897654345678",
					DisplayName:        "RBM-Terraform",
					Type:               string(api.REALBROWSER),
					BaseURL:            "https://www.example.com/",
					AsyncDCEnabled:     false,
					BrowserVersion:     0,
					CheckFrequency:     "15",
					IgnoreCertError:    false,
					IPType:             0,
					ThresholdProfileId: "789",
					PageLoadTime:       0,
					PerformAutomation:  false,
					Resolution:         "1600,900",
					LocationProfileID:  "456",
					MonitorGroups:      []string{"234", "567"},
					UserGroupIDs:       []string{"123", "456"},
					TagIDs:             []string{"123"},
				}

				_, err := NewWebTransactionBrowserMonitors(c).Update(webTransactionBrowserMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete web transaction browser monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/897654345678",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewWebTransactionBrowserMonitors(c).Delete("897654345678"))
			},
		},
	})
}
