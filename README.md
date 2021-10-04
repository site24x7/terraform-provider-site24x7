<a href="https://terraform.io">
    <img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-terraform-main.svg" alt="Terraform logo" title="Terraform" align="right" height="50" />
</a>

# Site24x7 Terraform Provider

- Site24x7 Website: <https://www.site24x7.com>
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
A terraform provider for managing the following resources in Site24x7:

- site24x7_website_monitor ([Site24x7 Website Monitor API doc](https://www.site24x7.com/help/api/#website))
- site24x7_ssl_monitor ([Site24x7 SSL Certificate Monitor API doc](https://www.site24x7.com/help/api/#ssl-certificate))
- site24x7_rest_api_monitor ([Site24x7 Rest API Monitor API doc](https://www.site24x7.com/help/api/#rest-api))
- site24x7_amazon_monitor ([Site24x7 Amazon Monitor API doc](https://www.site24x7.com/help/api/#amazon-webservice-monitor))
- site24x7_url_action ([Site24x7 IT Automation API doc](https://www.site24x7.com/help/api/#it-automation))
- site24x7_monitor_group ([Site24x7 Monitor Group API doc](https://www.site24x7.com/help/api/#monitor-groups))
- site24x7_threshold_profile ([Site24x7 Threshold Profile API doc](https://www.site24x7.com/help/api/#threshold-website))
- site24x7_user_group ([Site24x7 User Group API doc](https://www.site24x7.com/help/api/#user-groups))



If you're developing and building the provider, follow the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin). After placing the provider your plugins directory, run `terraform init` to initialize it.

For more information on using the provider and the associated resources, please see the [provider documentation][provider_docs] page.

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

Please see our [CONTRIBUTING][contributing] guide for more detail on the APIs
in use by this provider.

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
[Installing Terraform Providers](https://www.terraform.io/docs/cloud/run/install-software.html)
[Overriding Terraform's default installation](https://www.terraform.io/docs/cli/config/config-file.html)

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

























