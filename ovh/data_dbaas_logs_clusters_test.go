package ovh

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const testAccDataSourceDbaasLogsClusters = `
data "ovh_dbaas_logs_cluster" "ldp" {
  service_name = "%s"
}
`

func TestAccDataSourceDbaasLogsClusters(t *testing.T) {
	serviceName := os.Getenv("OVH_DBAAS_LOGS_SERVICE_TEST")
	clusterId := os.Getenv("OVH_DBAAS_LOGS_CLUSTER_ID")

	config := fmt.Sprintf(
		testAccDataSourceDbaasLogsClusters,
		serviceName,
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheckDbaasLogs(t) },

		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ovh_dbaas_logs_cluster.ldp",
						"service_name",
						serviceName,
					),
				),
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckTypeSetElemAttr(
						"data.ovh_dbaas_logs_cluster.ldp",
						"uuids",
						clusterId,
					),
				),
			},
		},
	})
}
