# min_provider

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.10 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.100.0 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_resource_names"></a> [resource\_names](#module\_resource\_names) | terraform.registry.launch.nttdata.com/module_library/resource_name/launch | ~> 2.0 |
| <a name="module_route53_zone_association"></a> [route53\_zone\_association](#module\_route53\_zone\_association) | ../.. | n/a |

## Resources

| Name | Type |
|------|------|
| [aws_default_security_group.primary](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/default_security_group) | resource |
| [aws_default_security_group.secondary](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/default_security_group) | resource |
| [aws_route53_zone.zone](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_zone) | resource |
| [aws_vpc.primary](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc) | resource |
| [aws_vpc.secondary](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc) | resource |
| [aws_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_class_env"></a> [class\_env](#input\_class\_env) | Environment class for resource naming. | `string` | `"dev"` | no |
| <a name="input_domain_name"></a> [domain\_name](#input\_domain\_name) | Domain name for the private hosted zone. | `string` | `"example.internal"` | no |
| <a name="input_instance_env"></a> [instance\_env](#input\_instance\_env) | Environment instance number for resource naming. | `number` | `0` | no |
| <a name="input_instance_resource"></a> [instance\_resource](#input\_instance\_resource) | Resource instance number for resource naming. | `number` | `0` | no |
| <a name="input_logical_product_family"></a> [logical\_product\_family](#input\_logical\_product\_family) | Logical product family for resource naming. | `string` | `"launch"` | no |
| <a name="input_logical_product_service"></a> [logical\_product\_service](#input\_logical\_product\_service) | Logical product service for resource naming. | `string` | `"r53za"` | no |
| <a name="input_resource_names_map"></a> [resource\_names\_map](#input\_resource\_names\_map) | Map of resource names for the resource naming module. | <pre>map(object({<br/>    name       = string<br/>    max_length = optional(number, 60)<br/>  }))</pre> | <pre>{<br/>  "route53_zone": {<br/>    "max_length": 60,<br/>    "name": "r53zone1"<br/>  },<br/>  "vpc_primary": {<br/>    "max_length": 60,<br/>    "name": "vpcpri1"<br/>  },<br/>  "vpc_secondary": {<br/>    "max_length": 60,<br/>    "name": "vpcsec1"<br/>  }<br/>}</pre> | no |
| <a name="input_tags"></a> [tags](#input\_tags) | Map of tags to assign to resources. | `map(string)` | `{}` | no |
| <a name="input_timeouts"></a> [timeouts](#input\_timeouts) | create = Timeout for creating the zone association. Default 30m.<br/>delete = Timeout for deleting the zone association. Default 30m. | <pre>object({<br/>    create = optional(string)<br/>    delete = optional(string)<br/>  })</pre> | `null` | no |
| <a name="input_vpc_id"></a> [vpc\_id](#input\_vpc\_id) | The VPC ID to associate with the private hosted zone. If not provided, VPCs will be created. | `string` | `null` | no |
| <a name="input_vpc_region"></a> [vpc\_region](#input\_vpc\_region) | The VPC's region. Defaults to the region of the AWS provider. | `string` | `null` | no |
| <a name="input_zone_id"></a> [zone\_id](#input\_zone\_id) | The private hosted zone ID to associate. If not provided, a zone will be created. | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_configured_timeout_create"></a> [configured\_timeout\_create](#output\_configured\_timeout\_create) | Configured create timeout passed to the module in this example. |
| <a name="output_configured_timeout_delete"></a> [configured\_timeout\_delete](#output\_configured\_timeout\_delete) | Configured delete timeout passed to the module in this example. |
| <a name="output_expected_vpc_id"></a> [expected\_vpc\_id](#output\_expected\_vpc\_id) | Expected VPC ID from the test fixture. |
| <a name="output_expected_vpc_region"></a> [expected\_vpc\_region](#output\_expected\_vpc\_region) | Expected VPC region from the provider/test fixture. |
| <a name="output_expected_zone_id"></a> [expected\_zone\_id](#output\_expected\_zone\_id) | Expected zone ID from the test fixture. |
| <a name="output_id"></a> [id](#output\_id) | The calculated unique identifier for the zone association. |
| <a name="output_owning_account"></a> [owning\_account](#output\_owning\_account) | The account ID of the account that created the hosted zone. |
| <a name="output_vpc_id"></a> [vpc\_id](#output\_vpc\_id) | The VPC ID returned by the module. |
| <a name="output_vpc_region"></a> [vpc\_region](#output\_vpc\_region) | The VPC region returned by the module. |
| <a name="output_zone_id"></a> [zone\_id](#output\_zone\_id) | The private hosted zone ID returned by the module. |
<!-- END_TF_DOCS -->
