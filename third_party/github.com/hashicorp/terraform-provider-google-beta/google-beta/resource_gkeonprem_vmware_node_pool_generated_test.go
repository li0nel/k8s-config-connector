// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccGkeonpremVmwareNodePool_gkeonpremVmwareNodePoolBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckGkeonpremVmwareNodePoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGkeonpremVmwareNodePool_gkeonpremVmwareNodePoolBasicExample(context),
			},
			{
				ResourceName:            "google_gkeonprem_vmware_node_pool.nodepool-basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "vmware_cluster", "location"},
			},
		},
	})
}

func testAccGkeonpremVmwareNodePool_gkeonpremVmwareNodePoolBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_gkeonprem_vmware_cluster" "default-basic" {
  provider = google-beta
  location = "us-west1"
  name = "default-basic"
  admin_cluster_membership = "projects/870316890899/locations/global/memberships/gkeonprem-terraform-test"
  description = "test cluster"
  on_prem_version = "1.13.1-gke.35"
  network_config {
    service_address_cidr_blocks = ["10.96.0.0/12"]
    pod_address_cidr_blocks = ["192.168.0.0/16"]
    dhcp_ip_config {
      enabled = true
    }
  }
  control_plane_node {
     cpus = 4
     memory = 8192
     replicas = 1
  }
  load_balancer {
    vip_config {
      control_plane_vip = "10.251.133.5"
      ingress_vip = "10.251.135.19"
    }
    metal_lb_config {
      address_pools {
        pool = "ingress-ip"
        manual_assign = "true"
        addresses = ["10.251.135.19"]
      }
      address_pools {
        pool = "lb-test-ip"
        manual_assign = "true"
        addresses = ["10.251.135.19"]
      }
    }
  }
}

resource "google_gkeonprem_vmware_node_pool" "nodepool-basic" {
  provider = google-beta
  location = "us-west1"
  name = "np-nodepool%{random_suffix}"
  vmware_cluster = google_gkeonprem_vmware_cluster.default-basic.name
  config {
    replicas = 3
    image_type = "ubuntu_containerd"
    enable_load_balancer = true
  }
}
`, context)
}

func TestAccGkeonpremVmwareNodePool_gkeonpremVmwareNodePoolFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckGkeonpremVmwareNodePoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGkeonpremVmwareNodePool_gkeonpremVmwareNodePoolFullExample(context),
			},
			{
				ResourceName:            "google_gkeonprem_vmware_node_pool.nodepool-full",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "vmware_cluster", "location"},
			},
		},
	})
}

func testAccGkeonpremVmwareNodePool_gkeonpremVmwareNodePoolFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_gkeonprem_vmware_cluster" "default-full" {
  provider = google-beta
  location = "us-west1"
  name = "default-full"
  admin_cluster_membership = "projects/870316890899/locations/global/memberships/gkeonprem-terraform-test"
  description = "test cluster"
  on_prem_version = "1.13.1-gke.35"
  network_config {
    service_address_cidr_blocks = ["10.96.0.0/12"]
    pod_address_cidr_blocks = ["192.168.0.0/16"]
    dhcp_ip_config {
      enabled = true
    }
  }
  control_plane_node {
     cpus = 4
     memory = 8192
     replicas = 1
  }
  load_balancer {
    vip_config {
      control_plane_vip = "10.251.133.5"
      ingress_vip = "10.251.135.19"
    }
    metal_lb_config {
      address_pools {
        pool = "ingress-ip"
        manual_assign = "true"
        addresses = ["10.251.135.19"]
      }
      address_pools {
        pool = "lb-test-ip"
        manual_assign = "true"
        addresses = ["10.251.135.19"]
      }
    }
  }
}

resource "google_gkeonprem_vmware_node_pool" "nodepool-full" {
  provider = google-beta
  location = "us-west1"
  name = "np-nodepool%{random_suffix}"
  vmware_cluster = google_gkeonprem_vmware_cluster.default-full.name
  annotations = {}
  config {
    cpus = 4
    memory_mb = 8196
    replicas = 3
    image_type = "ubuntu_containerd"
    image = "image"
    boot_disk_size_gb = 10
    taints {
        key = "key"
        value = "value"
    }
    taints {
        key = "key"
        value = "value"
        effect = "NO_SCHEDULE"
    }
    labels = {}
    enable_load_balancer = true
  }
  node_pool_autoscaling {
    min_replicas = 1
    max_replicas = 5
  }
}
`, context)
}

func testAccCheckGkeonpremVmwareNodePoolDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_gkeonprem_vmware_node_pool" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{GkeonpremBasePath}}projects/{{project}}/locations/{{location}}/vmwareClusters/{{vmware_cluster}}/vmwareNodePools/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("GkeonpremVmwareNodePool still exists at %s", url)
			}
		}

		return nil
	}
}
