// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

data "aws_region" "current" {}

module "resource_names" {
  source  = "terraform.registry.launch.nttdata.com/module_library/resource_name/launch"
  version = "~> 2.0"

  for_each = var.resource_names_map

  logical_product_family  = var.logical_product_family
  logical_product_service = var.logical_product_service
  class_env               = var.class_env
  instance_env            = var.instance_env
  instance_resource       = var.instance_resource
  cloud_resource_type     = each.value.name
  maximum_length          = each.value.max_length
  region                  = join("", split("-", data.aws_region.current.id))
}

resource "aws_vpc" "primary" {
  cidr_block           = "10.6.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = merge(var.tags, {
    Name = module.resource_names["vpc_primary"].standard
  })
}

resource "aws_vpc" "secondary" {
  cidr_block           = "10.7.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = merge(var.tags, {
    Name = module.resource_names["vpc_secondary"].standard
  })
}

resource "aws_default_security_group" "primary" {
  vpc_id = aws_vpc.primary.id

  tags = merge(var.tags, {
    Name = "${module.resource_names["vpc_primary"].standard}-default-sg"
  })
}

resource "aws_default_security_group" "secondary" {
  vpc_id = aws_vpc.secondary.id

  tags = merge(var.tags, {
    Name = "${module.resource_names["vpc_secondary"].standard}-default-sg"
  })
}

resource "aws_route53_zone" "zone" {
  name = var.domain_name

  vpc {
    vpc_id = aws_vpc.primary.id
  }

  tags = merge(var.tags, {
    Name = module.resource_names["route53_zone"].standard
  })

  lifecycle {
    ignore_changes = [vpc]
  }
}

module "route53_zone_association" {
  source = "../.."

  zone_id    = var.zone_id != null ? var.zone_id : aws_route53_zone.zone.zone_id
  vpc_id     = var.vpc_id != null ? var.vpc_id : aws_vpc.secondary.id
  vpc_region = var.vpc_region
  timeouts   = var.timeouts
}
