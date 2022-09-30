package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNotificationProfiles(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create notification profile",
			ExpectedVerb: "POST",
			ExpectedPath: "/notification_profiles",
			ExpectedBody: validation.Fixture(t, "requests/create_notification_profile.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				notificationProfile := &api.NotificationProfile{
					RcaNeeded:                   true,
					NotifyAfterExecutingActions: true,
					ProfileName:                 "Notifi Profile",
				}

				_, err := NewNotificationProfiles(c).Create(notificationProfile)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get notification profile",
			ExpectedVerb: "GET",
			ExpectedPath: "/notification_profiles/123",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_notification_profile.json"),
			Fn: func(t *testing.T, c rest.Client) {
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
			Name:         "list notification profiles",
			ExpectedVerb: "GET",
			ExpectedPath: "/notification_profiles",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_notification_profiles.json"),
			Fn: func(t *testing.T, c rest.Client) {
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
			Name:         "update notification profile",
			ExpectedVerb: "PUT",
			ExpectedPath: "/notification_profiles/123",
			ExpectedBody: validation.Fixture(t, "requests/update_notification_profile.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				notificationProfile := &api.NotificationProfile{
					ProfileID:   "123",
					ProfileName: "Notifi Profile",
					RcaNeeded:   true,

					NotifyAfterExecutingActions: true,
					SuppressAutomation:          false,
				}

				_, err := NewNotificationProfiles(c).Update(notificationProfile)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete notification profile",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/notification_profiles/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewNotificationProfiles(c).Delete("123"))
			},
		},
	})
}
