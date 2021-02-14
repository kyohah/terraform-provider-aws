package aws

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAwsLightsailInstanceDataSource_Arn(t *testing.T) {
	rName := acctest.RandomWithPrefix("tf-acc-test")
	dataSourceName := "data.aws_lightsail_instance.test"
	resourceName := "aws_lightsail_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckAwsLightsailInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAwsLightsailInstanceDataSourceConfigArn(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "arn", resourceName, "arn"),
					resource.TestCheckResourceAttrPair(dataSourceName, "blueprint_id", resourceName, "blueprint_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "bundle_id", resourceName, "bundle_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ipv6_address", resourceName, "ipv6_address"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ipv6_addresses.#", resourceName, "ipv6_addresses.#"),
					resource.TestCheckResourceAttrPair(dataSourceName, "key_pair_name", "key_pair_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "tags.%", resourceName, "tags.%"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ram_size", resourceName, "ram_size"),
				),
			},
		},
	})
}

func testAccAwsLightsailInstanceDataSourceConfigArn(rName string) string {
	return fmt.Sprintf(`
data "aws_region" "current" {}
data "aws_partition" "current" {}
data "aws_availability_zones" "available" {
  state = "available"

  filter {
    name   = "opt-in-status"
    values = ["opt-in-not-required"]
  }
}

resource "aws_lightsail_instance" "test" {
  name              = "%s"
  availability_zone = data.aws_availability_zones.available.names[0]
  blueprint_id      = "amazon_linux"
  bundle_id         = "nano_1_0"
}

data "aws_lightsail_instance" "test" {
  arn = aws_lightsail_instance.test.arn
}
`, rName)
}
