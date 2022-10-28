package integration

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTelegramIntegration(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create telegram integration",
			ExpectedVerb: "POST",
			ExpectedPath: "/integration/telegram",
			ExpectedBody: validation.Fixture(t, "requests/create_telegram_integration.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				telegramIntegration := &api.TelegramIntegration{
					Name:          "Site24x7-Telegram Integration",
					URL:           "https://web.tesdfm.org/z/#-12345",
					BotToken:	   "53asdfg8:ART2345b-u-Ytdfgm-TI8",
					SelectionType: 0,
					TroubleAlert:  true,
					CriticalAlert: false,
					DownAlert:     false,
					Title:         "$MONITOR_NAME is $STATUS",
					AlertTagIDs:   []string{"1234567801"},
				}

				_, err := NewTelegram(c).Create(telegramIntegration)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get telegram integration",
			ExpectedVerb: "GET",
			ExpectedPath: "/integration/telegram/113770000023231022",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_telegram_integration.json"),
			Fn: func(t *testing.T, c rest.Client) {
				telegram_integration, err := NewTelegram(c).Get("113770000023231022")
				require.NoError(t, err)

				expected := &api.TelegramIntegration{
					Name:          "Site24x7-Telegram Integration",
					URL:           "https://web.telegram.org/z/#-1234567",
					BotToken:	   "1234567899",
					ServiceID:     "1234567890",
					ServiceStatus: 0,
					SelectionType: 0,
					Title:         "$MONITOR_NAME is $STATUS",
					AlertTagIDs:   []string{"12345678901"},
				}

				assert.Equal(t, expected, telegram_integration)
			},
		},
		{
			Name:         "update telegram integration",
			ExpectedVerb: "PUT",
			ExpectedPath: "/integration/telegram/123",
			ExpectedBody: validation.Fixture(t, "requests/update_telegram_integration.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				telegram_integration := &api.TelegramIntegration{
					Name:          "Site24x7-Telegram Integration",
					URL:           "https://web.telegram.org/z/#-1234566",
					BotToken:	   "1234567898765",
					ServiceID:     "123",
					SelectionType: 2,
					TroubleAlert:  true,
					CriticalAlert: false,
					DownAlert:     false,
					Monitors:      []string{"123456789", "87654321"},
					Title:         "$MONITOR_NAME is $STATUS",
					AlertTagIDs:   []string{"123456789", "87654321"},
				}

				_, err := NewTelegram(c).Update(telegram_integration)
				require.NoError(t, err)
			},
		},
	})
}
