---
layout: "aws"
page_title: "AWS: aws_lightsail_instance"
sidebar_current: "docs-aws-datasource-autoscaling-group"
description: |-
  Get information on an Amazon EC2 Autoscaling Group.
---

# Data Source: aws_lightsail_instance

Use this data source to get information on an existing autoscaling group.

## Example Usage

```hcl
data "lightsail_instance" "foo" {
  name = "foo"
}
```

## Argument Reference

* `name` - (Required) The name of the Lightsail Instance. Names be unique within each AWS Region in your Lightsail account.
* `availability_zone` - (Required) The Availability Zone in which to create your
instance (see list below)
* `blueprint_id` - (Required) The ID for a virtual private server image. A list of available blueprint IDs can be obtained using the AWS CLI command: `aws lightsail get-blueprints`
* `bundle_id` - (Required) The bundle of specification information (see list below)
* `key_pair_name` - (Optional) The name of your key pair. Created in the
Lightsail console (cannot use `aws_key_pair` at this time)
* `user_data` - (Optional) launch script to configure server with additional user data
* `tags` - (Optional) A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value.

## Attributes Reference

~> **NOTE:** Some values are not always set and may not be available for
interpolation.

* `id` - The ARN of the Lightsail instance (matches `arn`).
* `arn` - The ARN of the Lightsail instance (matches `id`).
* `created_at` - The timestamp when the instance was created.
* `ipv6_address` - (**Deprecated**) The first IPv6 address of the Lightsail instance. Use `ipv6_addresses` attribute instead.
* `ipv6_addresses` - List of IPv6 addresses for the Lightsail instance.
