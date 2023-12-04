package pkg

import (
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
)

func TestEssTagResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `resource "aliyun_ess_tag" "test" {
						  ess_id = "asg-id"
						  tags = {
							"a" = "b"
							"c" = "d"
						  }
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aliyun_ess_tag.test", "ess_id", "asg-id"),
					resource.TestCheckResourceAttr("aliyun_ess_tag.test", "tags.a", "b"),
					resource.TestCheckResourceAttr("aliyun_ess_tag.test", "tags.b", "d"),
				),
			},
			// Update and Read testing
			{
				Config: providerConfig + `resource "aliyun_ess_tag" "test" {
						  ess_id = "asg-id"
						  tags = {
							"a" = "bb"
							"c" = "dd"
						  }
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aliyun_ess_tag.test", "tags.a", "bb"),
					resource.TestCheckResourceAttr("aliyun_ess_tag.test", "tags.b", "dd"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
