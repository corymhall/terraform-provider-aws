// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmesh_test

import (
	"context"
	"fmt"
	"testing"

	awstypes "github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfappmesh "github.com/hashicorp/terraform-provider-aws/internal/service/appmesh"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func testAccVirtualService_virtualNode(t *testing.T) {
	ctx := acctest.Context(t)
	var vs awstypes.VirtualServiceData
	resourceName := "aws_appmesh_virtual_service.test"
	meshName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	vnName1 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	vnName2 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	vsName := fmt.Sprintf("tf-acc-test-%d.mesh.local", sdkacctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); acctest.PreCheckPartitionHasService(t, names.AppMeshEndpointID) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AppMeshServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVirtualServiceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVirtualServiceConfig_virtualNode(meshName, vnName1, vnName2, vsName, "aws_appmesh_virtual_node.test1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVirtualServiceExists(ctx, resourceName, &vs),
					resource.TestCheckResourceAttr(resourceName, "name", vsName),
					resource.TestCheckResourceAttr(resourceName, "mesh_name", meshName),
					acctest.CheckResourceAttrAccountID(resourceName, "mesh_owner"),
					resource.TestCheckResourceAttr(resourceName, "spec.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.provider.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.provider.0.virtual_node.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.provider.0.virtual_node.0.virtual_node_name", vnName1),
					resource.TestCheckResourceAttrSet(resourceName, "created_date"),
					resource.TestCheckResourceAttrSet(resourceName, "last_updated_date"),
					acctest.CheckResourceAttrAccountID(resourceName, "resource_owner"),
					acctest.CheckResourceAttrRegionalARN(resourceName, "arn", "appmesh", fmt.Sprintf("mesh/%s/virtualService/%s", meshName, vsName)),
				),
			},
			{
				Config: testAccVirtualServiceConfig_virtualNode(meshName, vnName1, vnName2, vsName, "aws_appmesh_virtual_node.test2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVirtualServiceExists(ctx, resourceName, &vs),
					resource.TestCheckResourceAttr(resourceName, "name", vsName),
					resource.TestCheckResourceAttr(resourceName, "mesh_name", meshName),
					acctest.CheckResourceAttrAccountID(resourceName, "mesh_owner"),
					resource.TestCheckResourceAttr(resourceName, "spec.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.provider.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.provider.0.virtual_node.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.provider.0.virtual_node.0.virtual_node_name", vnName2),
				),
			},
			{
				ResourceName:      resourceName,
				ImportStateId:     fmt.Sprintf("%s/%s", meshName, vsName),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccVirtualService_virtualRouter(t *testing.T) {
	ctx := acctest.Context(t)
	var vs awstypes.VirtualServiceData
	resourceName := "aws_appmesh_virtual_service.test"
	meshName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	vrName1 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	vrName2 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	vsName := fmt.Sprintf("tf-acc-test-%d.mesh.local", sdkacctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); acctest.PreCheckPartitionHasService(t, names.AppMeshEndpointID) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AppMeshServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVirtualServiceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVirtualServiceConfig_virtualRouter(meshName, vrName1, vrName2, vsName, "aws_appmesh_virtual_router.test1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVirtualServiceExists(ctx, resourceName, &vs),
					resource.TestCheckResourceAttr(resourceName, "name", vsName),
					resource.TestCheckResourceAttr(resourceName, "mesh_name", meshName),
					acctest.CheckResourceAttrAccountID(resourceName, "mesh_owner"),
					resource.TestCheckResourceAttr(resourceName, "spec.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.provider.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.provider.0.virtual_router.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.provider.0.virtual_router.0.virtual_router_name", vrName1),
					resource.TestCheckResourceAttrSet(resourceName, "created_date"),
					resource.TestCheckResourceAttrSet(resourceName, "last_updated_date"),
					acctest.CheckResourceAttrAccountID(resourceName, "resource_owner"),
					acctest.CheckResourceAttrRegionalARN(resourceName, "arn", "appmesh", fmt.Sprintf("mesh/%s/virtualService/%s", meshName, vsName))),
			},
			{
				Config: testAccVirtualServiceConfig_virtualRouter(meshName, vrName1, vrName2, vsName, "aws_appmesh_virtual_router.test2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVirtualServiceExists(ctx, resourceName, &vs),
					resource.TestCheckResourceAttr(resourceName, "name", vsName),
					resource.TestCheckResourceAttr(resourceName, "mesh_name", meshName),
					acctest.CheckResourceAttrAccountID(resourceName, "mesh_owner"),
					resource.TestCheckResourceAttr(resourceName, "spec.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.provider.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.provider.0.virtual_router.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.provider.0.virtual_router.0.virtual_router_name", vrName2),
				),
			},
		},
	})
}

func testAccVirtualService_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var vs awstypes.VirtualServiceData
	resourceName := "aws_appmesh_virtual_service.test"
	meshName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	vnName1 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	vnName2 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	vsName := fmt.Sprintf("tf-acc-test-%d.mesh.local", sdkacctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); acctest.PreCheckPartitionHasService(t, names.AppMeshEndpointID) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AppMeshServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVirtualServiceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVirtualServiceConfig_tags1(meshName, vnName1, vnName2, vsName, "aws_appmesh_virtual_node.test1", "key1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVirtualServiceExists(ctx, resourceName, &vs),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				Config: testAccVirtualServiceConfig_tags2(meshName, vnName1, vnName2, vsName, "aws_appmesh_virtual_node.test1", "key1", "value1updated", "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVirtualServiceExists(ctx, resourceName, &vs),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				Config: testAccVirtualServiceConfig_tags1(meshName, vnName1, vnName2, vsName, "aws_appmesh_virtual_node.test1", "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVirtualServiceExists(ctx, resourceName, &vs),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportStateId:     fmt.Sprintf("%s/%s", meshName, vsName),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccVirtualService_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var vs awstypes.VirtualServiceData
	resourceName := "aws_appmesh_virtual_service.test"
	meshName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	vnName1 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	vnName2 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	vsName := fmt.Sprintf("tf-acc-test-%d.mesh.local", sdkacctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); acctest.PreCheckPartitionHasService(t, names.AppMeshEndpointID) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AppMeshServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVirtualServiceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVirtualServiceConfig_virtualNode(meshName, vnName1, vnName2, vsName, "aws_appmesh_virtual_node.test1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVirtualServiceExists(ctx, resourceName, &vs),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfappmesh.ResourceVirtualService(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckVirtualServiceDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).AppMeshClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_appmesh_virtual_service" {
				continue
			}

			_, err := tfappmesh.FindVirtualServiceByThreePartKey(ctx, conn, rs.Primary.Attributes["mesh_name"], rs.Primary.Attributes["mesh_owner"], rs.Primary.Attributes["name"])

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("App Mesh Virtual Service %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckVirtualServiceExists(ctx context.Context, n string, v *awstypes.VirtualServiceData) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).AppMeshClient(ctx)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No App Mesh Virtual Service ID is set")
		}

		output, err := tfappmesh.FindVirtualServiceByThreePartKey(ctx, conn, rs.Primary.Attributes["mesh_name"], rs.Primary.Attributes["mesh_owner"], rs.Primary.Attributes["name"])

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccVirtualServiceConfig_virtualNode(meshName, vnName1, vnName2, vsName, rName string) string {
	return fmt.Sprintf(`
resource "aws_appmesh_mesh" "test" {
  name = %[1]q
}

resource "aws_appmesh_virtual_node" "test1" {
  name      = %[2]q
  mesh_name = aws_appmesh_mesh.test.id

  spec {}
}

resource "aws_appmesh_virtual_node" "test2" {
  name      = %[3]q
  mesh_name = aws_appmesh_mesh.test.id

  spec {}
}

resource "aws_appmesh_virtual_service" "test" {
  name      = %[4]q
  mesh_name = aws_appmesh_mesh.test.id

  spec {
    provider {
      virtual_node {
        virtual_node_name = %[5]s.name
      }
    }
  }
}
`, meshName, vnName1, vnName2, vsName, rName)
}

func testAccVirtualServiceConfig_virtualRouter(meshName, vrName1, vrName2, vsName, rName string) string {
	return fmt.Sprintf(`
resource "aws_appmesh_mesh" "test" {
  name = %[1]q
}

resource "aws_appmesh_virtual_router" "test1" {
  name      = %[2]q
  mesh_name = aws_appmesh_mesh.test.id

  spec {
    listener {
      port_mapping {
        port     = 8080
        protocol = "http"
      }
    }
  }
}

resource "aws_appmesh_virtual_router" "test2" {
  name      = %[3]q
  mesh_name = aws_appmesh_mesh.test.id

  spec {
    listener {
      port_mapping {
        port     = 8080
        protocol = "http"
      }
    }
  }
}

resource "aws_appmesh_virtual_service" "test" {
  name      = %[4]q
  mesh_name = aws_appmesh_mesh.test.id

  spec {
    provider {
      virtual_router {
        virtual_router_name = %[5]s.name
      }
    }
  }
}
`, meshName, vrName1, vrName2, vsName, rName)
}

func testAccVirtualServiceConfig_tags1(meshName, vnName1, vnName2, vsName, rName, tagKey1, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_appmesh_mesh" "test" {
  name = %[1]q
}

resource "aws_appmesh_virtual_node" "test1" {
  name      = %[2]q
  mesh_name = aws_appmesh_mesh.test.id

  spec {}
}

resource "aws_appmesh_virtual_node" "test2" {
  name      = %[3]q
  mesh_name = aws_appmesh_mesh.test.id

  spec {}
}

resource "aws_appmesh_virtual_service" "test" {
  name      = %[4]q
  mesh_name = aws_appmesh_mesh.test.id

  spec {
    provider {
      virtual_node {
        virtual_node_name = %[5]s.name
      }
    }
  }

  tags = {
    %[6]s = %[7]q
  }
}
`, meshName, vnName1, vnName2, vsName, rName, tagKey1, tagValue1)
}

func testAccVirtualServiceConfig_tags2(meshName, vnName1, vnName2, vsName, rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_appmesh_mesh" "test" {
  name = %[1]q
}

resource "aws_appmesh_virtual_node" "test1" {
  name      = %[2]q
  mesh_name = aws_appmesh_mesh.test.id

  spec {}
}

resource "aws_appmesh_virtual_node" "test2" {
  name      = %[3]q
  mesh_name = aws_appmesh_mesh.test.id

  spec {}
}

resource "aws_appmesh_virtual_service" "test" {
  name      = %[4]q
  mesh_name = aws_appmesh_mesh.test.id

  spec {
    provider {
      virtual_node {
        virtual_node_name = %[5]s.name
      }
    }
  }

  tags = {
    %[6]s = %[7]q
    %[8]s = %[9]q
  }
}
`, meshName, vnName1, vnName2, vsName, rName, tagKey1, tagValue1, tagKey2, tagValue2)
}
