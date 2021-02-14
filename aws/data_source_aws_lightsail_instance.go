package aws

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lightsail"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/keyvaluetags"
)

func dataSourceAwsLightsailInstance() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsLightsailInstanceRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"availability_zone": {
				Type:     schema.TypeString,
				Required: true,
			},
			"blueprint_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"bundle_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional attributes
			"key_pair_name": {
				// Not compatible with aws_key_pair (yet)
				// We'll need a new aws_lightsail_key_pair resource
				Type:     schema.TypeString,
				Optional: true,
			},

			// cannot be retrieved from the API
			"user_data": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// additional info returned from the API
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ram_size": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"ipv6_address": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: "use `ipv6_addresses` attribute instead",
			},
			"ipv6_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"is_static_ip": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"private_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsSchema(),
		},
	}
}

func dataSourceAwsLightsailInstanceRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).lightsailconn
	ignoreTagsConfig := meta.(*AWSClient).IgnoreTagsConfig

	output, err := conn.GetInstance(&lightsail.GetInstanceInput{
		InstanceName: aws.String(d.Id()),
	})

	if err != nil {
		log.Printf("Lightsail Instance (%s) not found, removing from state", d.Id())
		return nil
	}

	if output == nil {
		return fmt.Errorf("Lightsail Instance (%s) not found, nil response from server, removing from state: %w", err)
	}

	if output == nil || output.Image == nil {
		return fmt.Errorf("error getting Image Builder Image: empty response")
	}

	instance := output.Instance

	d.Set("availability_zone", instance.Location.AvailabilityZone)
	d.Set("blueprint_id", instance.BlueprintId)
	d.Set("bundle_id", instance.BundleId)
	d.Set("key_pair_name", instance.SshKeyName)
	d.Set("name", instance.Name)

	// additional attributes
	d.Set("arn", instance.Arn)
	d.Set("username", instance.Username)
	d.Set("created_at", instance.CreatedAt.Format(time.RFC3339))
	d.Set("cpu_count", instance.Hardware.CpuCount)
	d.Set("ram_size", instance.Hardware.RamSizeInGb)

	// Deprecated: AWS Go SDK v1.36.25 removed Ipv6Address field
	if len(i.Ipv6Addresses) > 0 {
		d.Set("ipv6_address", aws.StringValue(instance.Ipv6Addresses[0]))
	}

	d.Set("ipv6_addresses", aws.StringValueSlice(instance.Ipv6Addresses))
	d.Set("is_static_ip", instance.IsStaticIp)
	d.Set("private_ip_address", instance.PrivateIpAddress)
	d.Set("public_ip_address", instance.PublicIpAddress)

	d.Set("tags", keyvaluetags.ImagebuilderKeyValueTags(image.Tags).IgnoreAws().IgnoreConfig(meta.(*AWSClient).IgnoreTagsConfig).Map())

	return nil
}
