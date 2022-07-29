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
      version = "1.0.23"
    }
  }
}

```
A terraform provider for managing the following resources in Site24x7:

- Website Monitor - [site24x7_website_monitor](examples/website_monitor_us.tf) ([Site24x7 Website Monitor API doc](https://www.site24x7.com/help/api/#website))
- Web Page Speed (Browser) Monitor - [site24x7_web_page_speed_monitor](examples/web_page_speed_monitor_us.tf) ([Site24x7 Web Page Speed Monitor API doc](https://www.site24x7.com/help/api/#websocket))
- SSL Certificate Monitor - [site24x7_ssl_monitor](examples/ssl_monitor_us.tf) ([Site24x7 SSL Certificate Monitor API doc](https://www.site24x7.com/help/api/#ssl-certificate))
- Rest API Monitor - [site24x7_rest_api_monitor](examples/rest_api_monitor_us.tf) ([Site24x7 Rest API Monitor API doc](https://www.site24x7.com/help/api/#rest-api))
- Server Monitor - [site24x7_server_monitor](examples/server_monitor_us.tf) ([Terraform Server Monitor doc](https://registry.terraform.io/providers/site24x7/site24x7/latest/docs/resources/server_monitor))
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
- Schedule Maintenance - [site24x7_schedule_maintenance](examples/schedule_maintenance_us.tf) ([Site24x7 Schedule Maintenance API doc](https://www.site24x7.com/help/api/#schedule-maintenances))


Usage example
-------------

Refer to the [examples/](examples/) directory for a fully documented usage example.

Set your Site24x7 OAuth credentials in the bash environment

```sh

  $ export SITE24X7_OAUTH2_CLIENT_ID="<your_oauth2_client_id>"
  $ export SITE24X7_OAUTH2_CLIENT_SECRET="<your_oauth2_client_secret>"
  $ export SITE24X7_OAUTH2_REFRESH_TOKEN="<your_oauth2_refresh_token>"

```

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

  // ZAAID of the customer under a MSP or BU
  zaaid = "1234"

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
  check_frequency = "1"
}

```

## Steps to generate Site24x7 OAuth credentials

Site24x7 REST APIs uses the OAuth 2.0 protocol to authorize and authenticate calls. To generate Site24x7 OAuth credentials(SITE24X7_CLIENT_ID, SITE24X7_CLIENT_SECRET and SITE24X7_REFRESH_TOKEN) please follow the steps below 

1. Visit https://api-console.zoho.com/
2. Choose the self client option.
3. Copy the Client ID, Client Secret and paste them in the below curl command.
4. Copy and paste `Site24x7.account.All,Site24x7.admin.All,Site24x7.reports.All,Site24x7.operations.All,Site24x7.msp.All,Site24x7.bu.All` in the scope field and click the "Create" button.
5. Copy the generated code, paste it in the below command and execute the same.

```sh

  curl https://accounts.zoho.com/oauth/v2/token
  -X POST
  -d "client_id=<CLIENT_ID>"
  -d "client_secret=<CLIENT_SECRET>"
  -d "code=<GENERATED_CODE>"
  -d "grant_type=authorization_code" --insecure

```

Note: Domain names in the OAuth credentials generation steps vary based on your data center

1. United States (US) - https://accounts.zoho.com and https://api-console.zoho.com
2. Europe (EU) - https://accounts.zoho.eu and https://api-console.zoho.eu
3. China (CN) - https://accounts.zoho.com.cn and https://api-console.zoho.com.cn
4. India (IN) - https://accounts.zoho.in and https://api-console.zoho.in
5. Australia (AU) - https://accounts.zoho.com.au and https://api-console.zoho.com.au


## Steps to import existing monitors and generate terraform resource configuration for the same

#### Clone the repository

Execute the below command to clone Site24x7's terraform provider repository to any desired location in your file system.

```sh

  git clone https://github.com/site24x7/terraform-provider-site24x7.git
  cd terraform-provider-site24x7

```

The current directory denotes your `$SITE24X7_TERRAFORM_PROVIDER_REPOSITORY_HOME`

#### Export your Site24x7 OAuth credentials in the bash environment

```sh

  $ export SITE24X7_OAUTH2_CLIENT_ID="<your_oauth2_client_id>"
  $ export SITE24X7_OAUTH2_CLIENT_SECRET="<your_oauth2_client_secret>"
  $ export SITE24X7_OAUTH2_REFRESH_TOKEN="<your_oauth2_refresh_token>"

```

#### Fetch monitors to import

To fetch all the server monitor IDs using the datasource `site24x7_monitors` paste the below configuration in `$SITE24X7_TERRAFORM_PROVIDER_REPOSITORY_HOME/main.tf`

```terraform

    // Data source to fetch all SERVER monitors
    data "site24x7_monitors" "s247monitors" {
        // (Optional) Type of the monitor. (eg) RESTAPI, SSL_CERT, URL, SERVER etc.
        monitor_type = "SERVER"
    }

    resource "local_file" "key" {
        filename = "${path.module}/utilities/importer/monitors_to_import.json"
        content  = jsonencode(data.site24x7_monitors.s247monitors.ids)
    }

```

Execute the below commands to write all the server monitor IDs in the file `$SITE24X7_TERRAFORM_PROVIDER_REPOSITORY_HOME/utilities/importer/monitors_to_import.json`

```sh

  cd $SITE24X7_TERRAFORM_PROVIDER_REPOSITORY_HOME
  terraform init
  terraform apply

```

#### Generating configuration and import commands

Execute the below commands to generate empty configuration, terraform import commands and the state file configuration.

```sh

  cd utilities/importer
  terraform init
  python site24x7_importer.py --resource site24x7_server_monitor

```

#### Importing monitors to your state

Copy the empty configurations(similar to the one given below) generated in the file `$SITE24X7_TERRAFORM_PROVIDER_REPOSITORY_HOME/empty_configuration.tf` to your terraform configuration file.

```terraform

    resource "site24x7_server_monitor" "SERVER_123456000025786003" {
    }

    resource "site24x7_server_monitor" "SERVER_123456000027570003" {
    }

```

Copy `$SITE24X7_TERRAFORM_PROVIDER_REPOSITORY_HOME/utilities/importer/output/import_commands.sh` to your terraform directory and execute the same to import all the monitors to your terraform state.

```sh

  ./import_commands.sh

```

Copy the resource configurations(similar to the one given below) generated in the file `$SITE24X7_TERRAFORM_PROVIDER_REPOSITORY_HOME/utilities/importer/output/imported_configuration.tf` to your terraform configuration file.

```terraform

  resource "site24x7_server_monitor" "SERVER_123456000025786003" { 
      perform_automation = true 
      log_needed = true 
      notification_profile_id = "123456000000029001" 
      tag_ids = ["123456000024829001", "123456000024829005"] 
      poll_interval = 1 
      monitor_groups = ["123456000000120011"] 
      threshold_profile_id = "123456000000029003" 
      user_group_ids = ["123456000000025005", "123456000000025009"] 
      display_name = "ubuntu-server"
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

  git clone https://github.com/site24x7/terraform-provider-site24x7.git
  cd terraform-provider-site24x7
  ./build/build_terraform_provider_site24x7.sh

```

This will build the `terraform-provider-site24x7` binary and install it into
the `$HOME/.terraform.d/plugins/registry.terraform.io/site24x7/1.0.0/linux_amd64` directory.

Place the below content in ~/.terraformrc

```sh

provider_installation {
	filesystem_mirror {
	  path = "$HOME/.terraform.d/plugins"
	  include = ["registry.terraform.io/site24x7/*"]
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

























