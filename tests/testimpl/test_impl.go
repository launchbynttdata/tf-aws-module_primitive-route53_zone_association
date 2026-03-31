package testimpl

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	t.Run("TestRoute53ZoneAssociationOutputs", func(t *testing.T) {
		id := terraform.Output(t, ctx.TerratestTerraformOptions(), "id")
		require.NotEmpty(t, id, "Association ID must not be empty")

		owningAccount := terraform.Output(t, ctx.TerratestTerraformOptions(), "owning_account")
		require.NotEmpty(t, owningAccount, "Owning account must not be empty")

		zoneID := terraform.Output(t, ctx.TerratestTerraformOptions(), "zone_id")
		require.NotEmpty(t, zoneID, "Zone ID must not be empty")

		vpcID := terraform.Output(t, ctx.TerratestTerraformOptions(), "vpc_id")
		require.NotEmpty(t, vpcID, "VPC ID must not be empty")
	})

	t.Run("TestRoute53ZoneAssociationViaAPI", func(t *testing.T) {
		zoneID := terraform.Output(t, ctx.TerratestTerraformOptions(), "zone_id")
		vpcID := terraform.Output(t, ctx.TerratestTerraformOptions(), "vpc_id")

		cfg, err := config.LoadDefaultConfig(context.Background())
		require.NoError(t, err, "Failed to load AWS config")

		client := route53.NewFromConfig(cfg)
		result, err := client.GetHostedZone(context.Background(), &route53.GetHostedZoneInput{
			Id: &zoneID,
		})
		require.NoError(t, err, "Failed to get hosted zone")
		require.NotNil(t, result.VPCs, "VPCs list must not be nil")

		found := false
		for _, vpc := range result.VPCs {
			if vpc.VPCId != nil && *vpc.VPCId == vpcID {
				found = true
				break
			}
		}
		assert.True(t, found, "Secondary VPC %s should be associated with zone %s", vpcID, zoneID)
	})
}

func TestComposableCompleteReadonly(t *testing.T, ctx types.TestContext) {
	t.Run("TestRoute53ZoneAssociationOutputsReadonly", func(t *testing.T) {
		id := terraform.Output(t, ctx.TerratestTerraformOptions(), "id")
		require.NotEmpty(t, id, "Association ID must not be empty")

		owningAccount := terraform.Output(t, ctx.TerratestTerraformOptions(), "owning_account")
		require.NotEmpty(t, owningAccount, "Owning account must not be empty")

		zoneID := terraform.Output(t, ctx.TerratestTerraformOptions(), "zone_id")
		require.NotEmpty(t, zoneID, "Zone ID must not be empty")

		vpcID := terraform.Output(t, ctx.TerratestTerraformOptions(), "vpc_id")
		require.NotEmpty(t, vpcID, "VPC ID must not be empty")
	})

	t.Run("TestRoute53ZoneAssociationExistsViaAPI", func(t *testing.T) {
		zoneID := terraform.Output(t, ctx.TerratestTerraformOptions(), "zone_id")
		vpcID := terraform.Output(t, ctx.TerratestTerraformOptions(), "vpc_id")

		cfg, err := config.LoadDefaultConfig(context.Background())
		require.NoError(t, err, "Failed to load AWS config")

		client := route53.NewFromConfig(cfg)
		result, err := client.GetHostedZone(context.Background(), &route53.GetHostedZoneInput{
			Id: &zoneID,
		})
		require.NoError(t, err, "Failed to get hosted zone")
		require.NotNil(t, result.VPCs, "VPCs list must not be nil")

		found := false
		for _, vpc := range result.VPCs {
			if vpc.VPCId != nil && *vpc.VPCId == vpcID {
				found = true
				break
			}
		}
		assert.True(t, found, "Secondary VPC %s should be associated with zone %s", vpcID, zoneID)
	})
}
