package common

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBusinessHour(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create business hour",
			ExpectedVerb: "POST",
			ExpectedPath: "/business_hours",
			ExpectedBody: validation.Fixture(t, "requests/create_business_hour.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				businessHourCreate := &api.BusinessHour{
					DisplayName: "General Shift",
					Description: "5-day workweek from 9 AM to 5 PM",
					TimeConfig: []api.TimeSlot{
						{Day: 2, StartTime: "09:00", EndTime: "17:00"},
						{Day: 3, StartTime: "09:00", EndTime: "17:00"},
						{Day: 4, StartTime: "09:00", EndTime: "17:00"},
						{Day: 5, StartTime: "09:00", EndTime: "17:00"},
						{Day: 6, StartTime: "09:00", EndTime: "17:00"},
					},
				}

				_, err := NewBusinessHour(c).Create(businessHourCreate)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get business hour",
			ExpectedVerb: "GET",
			ExpectedPath: "/business_hours/12345",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_business_hour.json"),
			Fn: func(t *testing.T, c rest.Client) {
				group, err := NewBusinessHour(c).Get("12345")
				require.NoError(t, err)

				expected := &api.BusinessHour{
					DisplayName: "General Shift",
					Description: "5-day workweek from 9 AM to 5 PM",
					TimeConfig: []api.TimeSlot{
						{Day: 2, StartTime: "09:00", EndTime: "17:00"},
						{Day: 3, StartTime: "09:00", EndTime: "17:00"},
						{Day: 4, StartTime: "09:00", EndTime: "17:00"},
						{Day: 5, StartTime: "09:00", EndTime: "17:00"},
						{Day: 6, StartTime: "09:00", EndTime: "17:00"},
					},
				}

				assert.Equal(t, expected, group)
			},
		},
		{
			Name:         "list business hours",
			ExpectedVerb: "GET",
			ExpectedPath: "/business_hours",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_business_hours.json"),
			Fn: func(t *testing.T, c rest.Client) {
				groups, err := NewBusinessHour(c).List()
				require.NoError(t, err)

				expected := []*api.BusinessHour{
					{
						ID:          "12345",
						DisplayName: "General Shift",
						Description: "5-day workweek from 9 AM to 5 PM",
						TimeConfig: []api.TimeSlot{
							{Day: 2, StartTime: "09:00", EndTime: "17:00"},
							{Day: 3, StartTime: "09:00", EndTime: "17:00"},
							{Day: 4, StartTime: "09:00", EndTime: "17:00"},
							{Day: 5, StartTime: "09:00", EndTime: "17:00"},
							{Day: 6, StartTime: "09:00", EndTime: "17:00"},
						},
					},
				}

				assert.Equal(t, expected, groups)
			},
		},
		{
			Name:         "update business hour",
			ExpectedVerb: "PUT",
			ExpectedPath: "/business_hours/12345",
			ExpectedBody: validation.Fixture(t, "requests/update_business_hour.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				businessHourUpdate := &api.BusinessHour{
					ID:          "12345",
					DisplayName: "Updated Shift",
					Description: "Updated shift timings",
					TimeConfig: []api.TimeSlot{
						{Day: 2, StartTime: "08:00", EndTime: "16:00"},
						{Day: 3, StartTime: "08:00", EndTime: "16:00"},
						{Day: 4, StartTime: "08:00", EndTime: "16:00"},
					},
				}

				_, err := NewBusinessHour(c).Update(businessHourUpdate)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete business hour",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/businesshours/12345",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewBusinessHour(c).Delete("12345"))
			},
		},
	})
}
