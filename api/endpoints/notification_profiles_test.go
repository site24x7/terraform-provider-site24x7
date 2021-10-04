package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNotificationProfiles(t *testing.T) {
	runTests(t, []*endpointTest{
		{
			name:         "create notification profile",
			expectedVerb: "POST",
			expectedPath: "/notification_profiles",
			expectedBody: fixture(t, "requests/create_notification_profile.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
				notificationProfile := &api.NotificationProfile{
					RcaNeeded:                   true,
					NotifyAfterExecutingActions: true,
					ProfileName:                 "Notifi Profile",
					EscalationWaitTime:          60,
				}

				_, err := NewNotificationProfiles(c).Create(notificationProfile)
				require.NoError(t, err)
			},
		},
		{
			name:         "get notification profile",
			expectedVerb: "GET",
			expectedPath: "/notification_profiles/123",
			statusCode:   200,
			responseBody: fixture(t, "responses/get_notification_profile.json"),
			fn: func(t *testing.T, c rest.Client) {
				notificationProfile, err := NewNotificationProfiles(c).Get("123")
				require.NoError(t, err)

				expected := &api.NotificationProfile{
					ProfileID:   "123",
					ProfileName: "Notifi Profile",
					RcaNeeded:   true,
				}

				assert.Equal(t, expected, notificationProfile)
			},
		},
		{
			name:         "list notification profiles",
			expectedVerb: "GET",
			expectedPath: "/notification_profiles",
			statusCode:   200,
			responseBody: fixture(t, "responses/list_notification_profiles.json"),
			fn: func(t *testing.T, c rest.Client) {
				notificationProfiles, err := NewNotificationProfiles(c).List()
				require.NoError(t, err)

				expected := []*api.NotificationProfile{
					{
						ProfileID:   "123",
						ProfileName: "Notifi Profile",
						RcaNeeded:   true,
					},
					{
						ProfileID:   "456",
						ProfileName: "TEST",
						RcaNeeded:   false,
					},
				}

				assert.Equal(t, expected, notificationProfiles)
			},
		},
		{
			name:         "update notification profile",
			expectedVerb: "PUT",
			expectedPath: "/notification_profiles/123",
			expectedBody: fixture(t, "requests/update_notification_profile.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
				notificationProfile := &api.NotificationProfile{
					ProfileID:                   "123",
					ProfileName:                 "Notifi Profile",
					RcaNeeded:                   true,
					EscalationWaitTime:          60,
					NotifyAfterExecutingActions: true,
				}

				_, err := NewNotificationProfiles(c).Update(notificationProfile)
				require.NoError(t, err)
			},
		},
		{
			name:         "delete notification profile",
			expectedVerb: "DELETE",
			expectedPath: "/notification_profiles/123",
			statusCode:   200,
			fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewNotificationProfiles(c).Delete("123"))
			},
		},
	})
}
