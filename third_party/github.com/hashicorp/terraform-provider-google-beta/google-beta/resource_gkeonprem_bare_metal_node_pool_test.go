package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGkeonpremBareMetalNodePool_bareMetalNodePoolUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckGkeonpremBareMetalNodePoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGkeonpremBareMetalNodePool_bareMetalNodePoolUpdateStart(context),
			},
			{
				ResourceName:      "google_gkeonprem_bare_metal_node_pool.nodepool",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGkeonpremBareMetalNodePool_bareMetalNodePoolUpdate(context),
			},
			{
				ResourceName:      "google_gkeonprem_bare_metal_node_pool.nodepool",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccGkeonpremBareMetalNodePool_bareMetalNodePoolUpdateStart(context map[string]interface{}) string {
	return Nprintf(`

  resource "google_gkeonprem_bare_metal_cluster" "cluster" {
    provider = google-beta
    name = "cluster"
    location = "us-west1"
    admin_cluster_membership = "projects/870316890899/locations/global/memberships/gkeonprem-terraform-test"
    bare_metal_version = "1.12.3"
    network_config {
      island_mode_cidr {
        service_address_cidr_blocks = ["172.26.0.0/16"]
        pod_address_cidr_blocks = ["10.240.0.0/13"]
      }
    }
    control_plane {
      control_plane_node_pool_config {
        node_pool_config {
          labels = {}
          operating_system = "LINUX"
          node_configs {
            labels = {}
            node_ip = "10.200.0.9"
          }
        }
      }
    }
    load_balancer {
      port_config {
        control_plane_load_balancer_port = 443
      }
      vip_config {
        control_plane_vip = "10.200.0.13"
        ingress_vip = "10.200.0.14"
      }
      metal_lb_config {
        address_pools {
          pool = "pool1"
          addresses = [
            "10.200.0.14/32",
            "10.200.0.15/32",
            "10.200.0.16/32",
            "10.200.0.17/32",
            "10.200.0.18/32",
            "fd00:1::f/128",
            "fd00:1::10/128",
            "fd00:1::11/128",
            "fd00:1::12/128"
          ]
        }
      }
    }
    storage {
      lvp_share_config {
        lvp_config {
          path = "/mnt/localpv-share"
          storage_class = "local-shared"
        }
        shared_path_pv_count = 5
      }
      lvp_node_mounts_config {
        path = "/mnt/localpv-disk"
        storage_class = "local-disks"
      }
    }
    security_config {
      authorization {
        admin_users {
          username = "admin@hashicorptest.com"
        }
      }
    }
  }

  resource "google_gkeonprem_bare_metal_node_pool" "nodepool" {
    provider = google-beta
    location = "us-west1"
    name = "nodepool"
    bare_metal_cluster = google_gkeonprem_bare_metal_cluster.cluster.name
    annotations = {}
    node_pool_config {
      operating_system = "LINUX"
      labels = {}
      node_configs {
        node_ip = "10.200.0.11"
        labels = {}
      }
    }
  }
`, context)
}

func testAccGkeonpremBareMetalNodePool_bareMetalNodePoolUpdate(context map[string]interface{}) string {
	return Nprintf(`

  resource "google_gkeonprem_bare_metal_cluster" "cluster" {
    provider = google-beta
    name = "cluster"
    location = "us-west1"
    admin_cluster_membership = "projects/870316890899/locations/global/memberships/gkeonprem-terraform-test"
    bare_metal_version = "1.12.3"
    network_config {
      island_mode_cidr {
        service_address_cidr_blocks = ["172.26.0.0/16"]
        pod_address_cidr_blocks = ["10.240.0.0/13"]
      }
    }
    control_plane {
      control_plane_node_pool_config {
        node_pool_config {
          labels = {}
          operating_system = "LINUX"
          node_configs {
            labels = {}
            node_ip = "10.200.0.9"
          }
        }
      }
    }
    load_balancer {
      port_config {
        control_plane_load_balancer_port = 443
      }
      vip_config {
        control_plane_vip = "10.200.0.13"
        ingress_vip = "10.200.0.14"
      }
      metal_lb_config {
        address_pools {
          pool = "pool1"
          addresses = [
            "10.200.0.14/32",
            "10.200.0.15/32",
            "10.200.0.16/32",
            "10.200.0.17/32",
            "10.200.0.18/32",
            "fd00:1::f/128",
            "fd00:1::10/128",
            "fd00:1::11/128",
            "fd00:1::12/128"
          ]
        }
      }
    }
    storage {
      lvp_share_config {
        lvp_config {
          path = "/mnt/localpv-share"
          storage_class = "local-shared"
        }
        shared_path_pv_count = 5
      }
      lvp_node_mounts_config {
        path = "/mnt/localpv-disk"
        storage_class = "local-disks"
      }
    }
    security_config {
      authorization {
        admin_users {
          username = "admin@hashicorptest.com"
        }
      }
    }
  }

  resource "google_gkeonprem_bare_metal_node_pool" "nodepool" {
    provider = google-beta
    location = "us-west1"
    name = "nodepool"
    bare_metal_cluster = google_gkeonprem_bare_metal_cluster.cluster.name
    annotations = {}
    node_pool_config {
      operating_system = "LINUX"
      labels = {}
      node_configs {
        node_ip = "10.200.0.12"
        labels = {}
      }
    }
  }
`, context)
}
