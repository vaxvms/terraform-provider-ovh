package ovh

import (
	"fmt"
	"log"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/ovh/helpers"
)

func dataSourceDbaasLogsCluster() *schema.Resource {
	return &schema.Resource{
		Read: func(d *schema.ResourceData, meta interface{}) error {
			return dataSourceDbaasLogsClusterRead(d, meta)
		},
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:        schema.TypeString,
				Description: "The service name",
				Required:    true,
			},
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// Computed
			"urn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cluster_type": {
				Type:        schema.TypeString,
				Description: "Cluster type",
				Computed:    true,
			},
			"dedicated_input_pem": {
				Type:        schema.TypeString,
				Description: "PEM for dedicated inputs",
				Computed:    true,
				Sensitive:   true,
			},
			"archive_allowed_networks": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Allowed networks for ARCHIVE flow type",
				Computed:    true,
			},
			"direct_input_allowed_networks": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Allowed networks for DIRECT_INPUT flow type",
				Computed:    true,
			},
			"direct_input_pem": {
				Type:        schema.TypeString,
				Description: "PEM for direct inputs",
				Computed:    true,
				Sensitive:   true,
			},
			"hostname": {
				Type:        schema.TypeString,
				Description: "hostname",
				Computed:    true,
			},
			"is_default": {
				Type:        schema.TypeBool,
				Description: "All content generated by given service will be placed on this cluster",
				Computed:    true,
			},
			"is_unlocked": {
				Type:        schema.TypeBool,
				Description: "Allow given service to perform advanced operations on cluster",
				Computed:    true,
			},
			"query_allowed_networks": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Allowed networks for QUERY flow type",
				Computed:    true,
			},
			"region": {
				Type:        schema.TypeString,
				Description: "Data center localization",
				Computed:    true,
			},
		},
	}
}

func dataSourceDbaasLogsClusterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	serviceName := d.Get("service_name").(string)
	clusterId := d.Get("cluster_id").(string)

	log.Printf("[DEBUG] Will read dbaas logs cluster %s/%s", serviceName, clusterId)

	d.SetId(clusterId)
	d.Set("urn", helpers.ServiceURN(config.Plate, "ldp", serviceName))

	endpoint := fmt.Sprintf(
		"/dbaas/logs/%s/cluster/%s",
		url.PathEscape(serviceName),
		url.PathEscape(clusterId),
	)

	res := map[string]interface{}{}
	if err := config.OVHClient.Get(endpoint, &res); err != nil {
		return fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
	}

	d.Set("archive_allowed_networks", res["archiveAllowedNetworks"])
	d.Set("cluster_type", res["clusterType"])
	d.Set("dedicated_input_pem", res["dedicatedInputPEM"])
	d.Set("direct_input_allowed_networks", res["directInputAllowedNetworks"])
	d.Set("direct_input_pem", res["directInputPEM"])
	d.Set("hostname", res["hostname"])
	d.Set("is_default", res["isDefault"])
	d.Set("is_unlocked", res["isUnlocked"])
	d.Set("query_allowed_networks", res["queryAllowedNetworks"])
	d.Set("region", res["region"])

	return nil
}
