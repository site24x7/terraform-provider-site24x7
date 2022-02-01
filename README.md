<a href="https://terraform.io">
    <img src=".github/terraform_logo.svg" alt="Terraform logo" title="Terraform" align="right" height="50" />
</a>

# Site24x7 Terraform Provider

- Site24x7 Website: <https://www.site24x7.com>
- Terraform Website: <https://www.terraform.io>
- Tutorials: [learn.hashicorp.com](https://learn.hashicorp.com/terraform?track=getting-started#getting-started)
<!-- - Documentation: <https://registry.terraform.io/providers/site24x7/site24x7/latest/docs> -->

- Mailing List: [Google Groups](http://groups.google.com/group/terraform-tool)


The Terraform Site24x7 provider is a plugin for Terraform that allows for the full lifecycle management of Site24x7 resources.
This provider is maintained by Site24x7 team.

Please note: If you believe you have found a security issue in the Terraform Site24x7 Provider, please responsibly disclose by contacting us at support@site24x7.com.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.13+


## Using the provider

To use the latest version of the provider in your Terraform environment, run `terraform init` and Terraform will automatically install the provider.

For Terraform version 0.13.x

```hcl
terraform {
  required_version = "~> 0.15.0"
  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      // Update the latest version from https://registry.terraform.io/providers/site24x7/site24x7/latest 
      version = "1.0.1"
    }
  }
}
```
A terraform provider for managing the following resources in Site24x7:

- Website Monitor - [site24x7_website_monitor](examples/website_monitor_us.tf) ([Site24x7 Website Monitor API doc](https://www.site24x7.com/help/api/#website))
- SSL Certificate Monitor - [site24x7_ssl_monitor](examples/ssl_monitor_us.tf) ([Site24x7 SSL Certificate Monitor API doc](https://www.site24x7.com/help/api/#ssl-certificate))
- Rest API Monitor - [site24x7_rest_api_monitor](examples/rest_api_monitor_us.tf) ([Site24x7 Rest API Monitor API doc](https://www.site24x7.com/help/api/#rest-api))
- Amazon Monitor - [site24x7_amazon_monitor](examples/amazon_monitor_us.tf) ([Site24x7 Amazon Monitor API doc](https://www.site24x7.com/help/api/#amazon-webservice-monitor))
- URL IT Automation - [site24x7_url_action](examples/it_automation_us.tf) ([Site24x7 IT Automation API doc](https://www.site24x7.com/help/api/#it-automation))
- Monitor Group - [site24x7_monitor_group](examples/monitor_group_us.tf) ([Site24x7 Monitor Group API doc](https://www.site24x7.com/help/api/#monitor-groups))
- Threshold Profile - [site24x7_threshold_profile](examples/threshold_profile_us.tf) ([Site24x7 Threshold Profile API doc](https://www.site24x7.com/help/api/#threshold-website))
- Location Profile - [site24x7_location_profile](examples/location_profile_us.tf) ([Site24x7 Location Profile API doc](https://www.site24x7.com/help/api/#location-profiles))
- Notification Profile - [site24x7_notification_profile](examples/notification_profile_us.tf) ([Site24x7 Notification Profile API doc](https://www.site24x7.com/help/api/#notification-profiles))
- User Group - [site24x7_user_group](examples/user_group_us.tf) ([Site24x7 User Group API doc](https://www.site24x7.com/help/api/#user-groups))
- Tag - [site24x7_tag](examples/tag_us.tf) ([Site24x7 Tag API doc](https://www.site24x7.com/help/api/#tags))
- Opsgenie integration - [site24x7_opsgenie_integration](examples/opsgenie_integration_us.tf) ([Site24x7 Opsgenie integration API doc](https://www.site24x7.com/help/api/#create-opsgenie))
- PagerDuty integration - [site24x7_pagerduty_integration](examples/pagerduty_integration_us.tf) ([Site24x7 PagerDuty integration API doc](https://www.site24x7.com/help/api/#create-pagerduty))
- ServiceNow integration - [site24x7_servicenow_integration](examples/servicenow_integration_us.tf) ([Site24x7 ServiceNow integration API doc](https://www.site24x7.com/help/api/#create-servicenow))
- Slack integration - [site24x7_slack_integration](examples/slack_integration_us.tf) ([Site24x7 Slack integration API doc](https://www.site24x7.com/help/api/#create-slack))
- Webhook integration - [site24x7_webhook_integration](examples/webhook_integration_us.tf) ([Site24x7 Webhook integration API doc](https://www.site24x7.com/help/api/#create-webhook))


Usage example
-------------

Refer to the [examples/](examples/) directory for a fully documented usage example.

This is a quick example of the provider configuration:

```terraform
// Authentication API doc - https://www.site24x7.com/help/api/#authentication
provider "site24x7" {
  // The client ID will be looked up in the SITE24X7_OAUTH2_CLIENT_ID
  // environment variable if the attribute is empty or omitted.
  oauth2_client_id = "<SITE24X7_CLIENT_ID>"

  // The client secret will be looked up in the SITE24X7_OAUTH2_CLIENT_SECRET
  // environment variable if the attribute is empty or omitted.
  oauth2_client_secret = "<SITE24X7_CLIENT_SECRET>"

  // The refresh token will be looked up in the SITE24X7_OAUTH2_REFRESH_TOKEN
  // environment variable if the attribute is empty or omitted.
  oauth2_refresh_token = "<SITE24X7_REFRESH_TOKEN>"

  // Specify the data center from which you have obtained your
  // OAuth client credentials and refresh token. It can be (US/EU/IN/AU/CN).
  data_center = "US"
  
  // The minimum time to wait in seconds before retrying failed Site24x7 API requests.
  retry_min_wait = 1

  // The maximum time to wait in seconds before retrying failed Site24x7 API
  // requests. This is the upper limit for the wait duration with exponential
  // backoff.
  retry_max_wait = 30

  // Maximum number of Site24x7 API request retries to perform until giving up.
  max_retries = 4

}

// Website Monitor API doc: https://www.site24x7.com/help/api/#website
resource "site24x7_website_monitor" "website_monitor_us" {
  // (Required) Name of the monitor
  display_name = "Example Monitor"

  // (Required) Website address to monitor.
  website = "https://www.example.com"

  // (Optional) Check interval for monitoring. Default: 1. See
  // https://www.site24x7.com/help/api/#check-interval for all supported
  // values.
  check_frequency = 1
}

```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your
machine (version 1.13+ is _required_). You'll also need to correctly setup a
[GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

<!-- Please see our [CONTRIBUTING][contributing] guide for more detail on the APIs in use by this provider. -->

#### Building

Clone the repository and build the provider:

```sh

git clone git@github.com:Site24x7/terraform-provider-site24x7
cd terraform-provider-site24x7
./build/build_terraform_provider_site24x7.sh

```

This will build the `terraform-provider-site24x7` binary and install it into
the `$HOME/.terraform.d/plugins/registry.zoho.io/zoho/site24x7/1.0.0/linux_amd64` directory.

Place the below content in ~/.terraformrc

```sh

provider_installation {
  filesystem_mirror {
    path = "$HOME/.terraform.d/plugins/"
    include = ["registry.zoho.io/*/*"]
  }
  direct {
    exclude = ["registry.zoho.io/*/*"]
  }
}

```


Please refer the following links for installing custom providers.
- [Installing Terraform Providers](https://www.terraform.io/docs/cloud/run/install-software.html)
- [Overriding Terraform's default installation](https://www.terraform.io/docs/cli/config/config-file.html)
- [Documentation](https://pkg.go.dev/github.com/site24x7/terraform-provider-site24x7)

#### Go Version Support

We'll aim to support the latest supported release of Go, along with the
previous release. This doesn't mean that building with an older version of Go
will not work, but we don't intend to support a Go version in this project that
is not supported by the larger Go community. Please see the [Go
releases][go_releases] page for more details.

[provider_docs]: https://www.terraform.io/docs/providers/site24x7/index.html
[contributing]: https://github.com/site24x7/terraform-provider-site24x7/blob/main/CONTRIBUTING.md
[go_releases]: https://github.com/golang/go/wiki/Go-Release-Cycle


## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Even the tiniest contributions to the script or to the documentation are very welcome and **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b <branch name>`)
3. Commit your Changes (`git commit -m '<commit message>'`)
4. Push to the Branch (`git push origin <branch name>`)
5. Open a Pull Request


## License

Copyright (c) 2022 Zoho Corporation Private Limited

This project is licensed under the MIT License - see [LICENSE](https://github.com/site24x7/terraform-provider-site24x7/blob/main/LICENSE) file for details.


## Acknowledgments

The Site24x7 Terraform Provider uses code from the following library:

 * [Bonial.com](https://github.com/Bonial-International-GmbH), MIT License

























