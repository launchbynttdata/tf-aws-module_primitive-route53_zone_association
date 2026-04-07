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

variable "zone_id" {
  description = "The private hosted zone ID to associate. If not provided, a zone will be created."
  type        = string
  default     = null
}

variable "vpc_id" {
  description = "The VPC ID to associate with the private hosted zone. If not provided, VPCs will be created."
  type        = string
  default     = null
}

variable "vpc_region" {
  description = "The VPC's region. Defaults to the region of the AWS provider."
  type        = string
  default     = null
}

variable "timeouts" {
  description = <<-EOT
    create = Timeout for creating the zone association. Default 30m.
    delete = Timeout for deleting the zone association. Default 30m.
  EOT
  type = object({
    create = optional(string)
    delete = optional(string)
  })
  default = null
}

variable "logical_product_family" {
  description = "Logical product family for resource naming."
  type        = string
  default     = "launch"
}

variable "logical_product_service" {
  description = "Logical product service for resource naming."
  type        = string
  default     = "r53za"
}

variable "class_env" {
  description = "Environment class for resource naming."
  type        = string
  default     = "dev"
}

variable "instance_env" {
  description = "Environment instance number for resource naming."
  type        = number
  default     = 0
}

variable "instance_resource" {
  description = "Resource instance number for resource naming."
  type        = number
  default     = 0
}

variable "resource_names_map" {
  description = "Map of resource names for the resource naming module."
  type = map(object({
    name       = string
    max_length = optional(number, 60)
  }))
  default = {
    vpc_primary = {
      name       = "vpcpri1"
      max_length = 60
    }
    vpc_secondary = {
      name       = "vpcsec1"
      max_length = 60
    }
    route53_zone = {
      name       = "r53zone1"
      max_length = 60
    }
  }
}

variable "domain_name" {
  description = "Domain name for the private hosted zone."
  type        = string
  default     = "example.internal"
}

variable "tags" {
  description = "Map of tags to assign to resources."
  type        = map(string)
  default     = {}
}
