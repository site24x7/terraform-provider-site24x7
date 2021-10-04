<a href="https://terraform.io">
    <img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" alt="Terraform logo" title="Terraform" align="right" height="50" />
</a>

# Site24x7 Terraform Provider

- Terraform Website: <https://www.terraform.io>
- Tutorials: [learn.hashicorp.com](https://learn.hashicorp.com/terraform?track=getting-started#getting-started)
- Documentation: <https://registry.terraform.io/providers/site24x7/site24x7/latest/docs>
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
  required_version = "~> 0.13.0"
  required_providers {
    site24x7 = {
      source  = "site24x7/site24x7"
      version = "~> 1.0.0"
    }
  }
}
```
A terraform provider for managing Site24x7 monitors which currently supports the following resources:

- site24x7_url_action ([Site24x7 IT Automation API doc](https://www.site24x7.com/help/api/#it-automation))
- site24x7_monitor_group ([Site24x7 Monitor Group API doc](https://www.site24x7.com/help/api/#monitor-groups))
- site24x7_website_monitor ([Site24x7 Monitor API doc](https://www.site24x7.com/help/api/#website))
- site24x7_amazon_monitor


If you're developing and building the provider, follow the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin). After placing the provider your plugins directory, run `terraform init` to initialize it.

For more information on using the provider and the associated resources, please see the [provider documentation][provider_docs] page.

Usage example
-------------

Refer to the [_examples/](_examples/) directory for a fully documented usage example.

This is a quick example of the provider configuration:

```terraform
// Authentication API doc: https://www.site24x7.com/help/api/#authentication
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

  // Specify the following configuration options if you want to use
  // some other Site24x7 data center than the default US one. These
  // must correspond to the data center from which you have obtained your
  // OAuth client credentials and refresh token.

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
  // (Required) Name for the monitor.
  display_name = "zoho"

  // (Required) Website address to monitor.
  website = "https://www.zoho.com"

  // (Optional) Check interval for monitoring. Default: 1. See
  // https://www.site24x7.com/help/api/#check-interval for all supported
  // values.
  check_frequency = 1
}

resource "site24x7_monitor_group" "monitor_group_us" {
  // (Required) Display Name for the Monitor Group.
  display_name = "Website Group"

  // (Optional) Description for the Monitor Group.
  description = "This is the description of the monitor group from terraform"
}

resource "site24x7_url_action" "action_us" {
  // (Required) Display name for the action.
  name = "Vtitan Action"

  // (Required) URL to be invoked for action execution.
  url = "https://www.vtitan.com"

  // (Optional) HTTP Method to access the URL. Default: "P". See
  // https://www.site24x7.com/help/api/#http_methods for allowed values.
  method = "G"

  // (Optional) If send_custom_parameters is set as true. Custom parameters to
  // be passed while accessing the URL.
  custom_parameters = "param=value"

  // (Optional) Configuration to send custom parameters while executing the action.
  send_custom_parameters = true

  // (Optional) Configuration to enable json format for post parameters.
  send_in_json_format = true

  // (Optional) Configuration to send incident parameters while executing the action.
  send_incident_parameters = true

  // (Optional) The amount of time a connection waits to time out. Range 1 - 90. Default: 30.
  timeout = 10
}


resource "site24x7_amazon_monitor" "aws_monitor_site24x7" {
  display_name = "AWS Resource"
  aws_access_key = "<AWS_ACCESS_KEY>"
  aws_secret_key = "<AWS_SECRET_KEY>"
  aws_discovery_frequency = 5
  aws_discover_services = ["1"]
}

```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your
machine (version 1.13+ is _required_). You'll also need to correctly setup a
[GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

Please see our [CONTRIBUTING][contributing] guide for more detail on the APIs
in use by this provider.

#### Building

Clone the repository and build the provider:

```sh
git clone git@github.com:Site24x7/terraform-provider-site24x7
cd terraform-provider-site24x7
make install
```

This will build the `terraform-provider-site24x7` binary and install it into
the `$HOME/.terraform.d/plugins` directory.


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

Copyright (c) 2021 Zoho Corporation Private Limited

This project is licensed under the MIT License - see [LICENSE.md](LICENSE.md) file for details.


## Acknowledgments

The Site24x7 Terraform Provider uses code from the following library:

 * [Bonial.com](https://github.com/Bonial-International-GmbH), MIT License

























