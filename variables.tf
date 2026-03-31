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
  description = "The private hosted zone ID to associate."
  type        = string
}

variable "vpc_id" {
  description = "The VPC ID to associate with the private hosted zone."
  type        = string
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
