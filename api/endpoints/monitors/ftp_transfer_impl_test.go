package monitors

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFTPTransferMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create ftp transfer monitor",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitors",
			ExpectedBody: validation.Fixture(t, "requests/create_ftp_transfer_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				ftpTransferMonitor := &api.FTPTransferMonitor{
					DisplayName:           "FTP Transfer Monitor",
					HostName:              "www.example.com",
					Protocol:              "FTP",
					Type:                  "FTP",
					Port:                  443,
					CheckFrequency:        "5",
					Timeout:               30,
					CheckUpload:           true,
					CheckDownload:         true,
					Username:              "sas",
					Password:              "sas",
					Destination:           "/Home/sas/",
					PerformAutomation:     true,
					CredentialProfileID:   "2345536536",
					OnCallScheduleID:      "8687567555",
					LocationProfileID:     "456",
					NotificationProfileID: "789",
					ThresholdProfileID:    "012",
					MonitorGroups:         []string{"234", "567"},
					DependencyResourceIDs: []string{"234", "567"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
				}

				_, err := NewFTPTransferMonitors(c).Create(ftpTransferMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get ftp transfer monitor",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors/897654345678",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_ftp_transfer_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				ftpTransferMonitor, err := NewFTPTransferMonitors(c).Get("897654345678")
				require.NoError(t, err)

				expected := &api.FTPTransferMonitor{
					DisplayName:           "FTP Transfer Monitor",
					HostName:              "www.example.com",
					Protocol:              "FTP",
					Type:                  "FTP",
					Port:                  443,
					CheckFrequency:        "5",
					Timeout:               30,
					CheckUpload:           true,
					CheckDownload:         true,
					Username:              "sas",
					Password:              "sas",
					Destination:           "/Home/sas/",
					PerformAutomation:     true,
					CredentialProfileID:   "2345536536",
					OnCallScheduleID:      "8687567555",
					LocationProfileID:     "456",
					NotificationProfileID: "789",
					ThresholdProfileID:    "012",
					MonitorGroups:         []string{"234", "567"},
					DependencyResourceIDs: []string{"234", "567"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
				}

				assert.Equal(t, expected, ftpTransferMonitor)
			},
		},
		{
			Name:         "list ftp transfer monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_ftp_transfer_monitors.json"),
			Fn: func(t *testing.T, c rest.Client) {
				ftpTransferMonitors, err := NewFTPTransferMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.FTPTransferMonitor{
					{
						MonitorID:             "897654345678",
						DisplayName:           "FTP Transfer Monitor",
						HostName:              "www.example.com",
						Protocol:              "FTP",
						Type:                  "FTP",
						Port:                  443,
						CheckFrequency:        "5",
						Timeout:               30,
						CheckUpload:           true,
						CheckDownload:         true,
						Username:              "sas",
						Password:              "sas",
						Destination:           "/Home/sas/",
						PerformAutomation:     true,
						CredentialProfileID:   "2345536536",
						OnCallScheduleID:      "8687567555",
						LocationProfileID:     "456",
						NotificationProfileID: "789",
						ThresholdProfileID:    "012",
						MonitorGroups:         []string{"234", "567"},
						DependencyResourceIDs: []string{"234", "567"},
						UserGroupIDs:          []string{"123", "456"},
						TagIDs:                []string{"123"},
					},
					{
						MonitorID:             "933654345678",
						DisplayName:           "FTP Transfer Monitor",
						HostName:              "www.example.com",
						Protocol:              "FTP",
						Type:                  "FTP",
						Port:                  443,
						CheckFrequency:        "5",
						Timeout:               30,
						CheckUpload:           true,
						CheckDownload:         true,
						Username:              "sas",
						Password:              "sas",
						Destination:           "/Home/sas/",
						PerformAutomation:     true,
						CredentialProfileID:   "2345536536",
						OnCallScheduleID:      "8687567555",
						LocationProfileID:     "456",
						NotificationProfileID: "789",
						ThresholdProfileID:    "012",
						MonitorGroups:         []string{"234", "567"},
						DependencyResourceIDs: []string{"234", "567"},
						UserGroupIDs:          []string{"123", "456"},
						TagIDs:                []string{"123"},
					},
				}

				assert.Equal(t, expected, ftpTransferMonitors)
			},
		},
		{
			Name:         "update ftp transfer monitor",
			ExpectedVerb: "PUT",
			ExpectedPath: "/monitors/123",
			ExpectedBody: validation.Fixture(t, "requests/update_ftp_transfer_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				ftpTransferMonitor := &api.FTPTransferMonitor{
					MonitorID:             "123",
					DisplayName:           "FTP Transfer Monitor",
					HostName:              "www.example.com",
					Protocol:              "FTP",
					Type:                  "FTP",
					Port:                  443,
					CheckFrequency:        "5",
					Timeout:               30,
					CheckUpload:           true,
					CheckDownload:         true,
					Username:              "sas",
					Password:              "sas",
					Destination:           "/Home/sas/",
					PerformAutomation:     true,
					CredentialProfileID:   "2345536536",
					OnCallScheduleID:      "8687567555",
					LocationProfileID:     "456",
					NotificationProfileID: "789",
					ThresholdProfileID:    "012",
					MonitorGroups:         []string{"234", "567"},
					DependencyResourceIDs: []string{"234", "567"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
				}

				_, err := NewFTPTransferMonitors(c).Update(ftpTransferMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete ftp transfer monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewFTPTransferMonitors(c).Delete("123"))
			},
		},
	})
}
