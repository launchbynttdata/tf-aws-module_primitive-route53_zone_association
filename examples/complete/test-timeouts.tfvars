logical_product_family  = "launch"
logical_product_service = "r53za"
class_env               = "dev"
instance_env            = 0
instance_resource       = 0
domain_name             = "example.internal"
timeouts = {
  create = "45m"
  delete = "45m"
}
tags = {
  "Environment" = "test"
}
