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

output "id" {
  description = "The calculated unique identifier for the zone association."
  value       = module.route53_zone_association.id
}

output "owning_account" {
  description = "The account ID of the account that created the hosted zone."
  value       = module.route53_zone_association.owning_account
}

output "zone_id" {
  description = "The private hosted zone ID returned by the module."
  value       = module.route53_zone_association.zone_id
}

output "vpc_id" {
  description = "The VPC ID returned by the module."
  value       = module.route53_zone_association.vpc_id
}

output "vpc_region" {
  description = "The VPC region returned by the module."
  value       = module.route53_zone_association.vpc_region
}

output "expected_zone_id" {
  description = "Expected zone ID from the test fixture."
  value       = aws_route53_zone.zone.zone_id
}

output "expected_vpc_id" {
  description = "Expected VPC ID from the test fixture."
  value       = aws_vpc.secondary.id
}

output "expected_vpc_region" {
  description = "Expected VPC region from the provider/test fixture."
  value       = data.aws_region.current.name
}

output "configured_timeout_create" {
  description = "Configured create timeout passed to the module in this example."
  value       = try(var.timeouts.create, null)
}

output "configured_timeout_delete" {
  description = "Configured delete timeout passed to the module in this example."
  value       = try(var.timeouts.delete, null)
}
