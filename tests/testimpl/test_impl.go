package testimpl

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	route53types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type zoneAssociationOutputs struct {
	id                string
	owningAccount     string
	zoneID            string
	vpcID             string
	vpcRegion         string
	expectedZoneID    string
	expectedVPCID     string
	expectedVPCRegion string
}

func getOutputs(t *testing.T, ctx types.TestContext) zoneAssociationOutputs {
	t.Helper()
	opts := ctx.TerratestTerraformOptions()
	return zoneAssociationOutputs{
		id:                terraform.Output(t, opts, "id"),
		owningAccount:     terraform.Output(t, opts, "owning_account"),
		zoneID:            terraform.Output(t, opts, "zone_id"),
		vpcID:             terraform.Output(t, opts, "vpc_id"),
		vpcRegion:         terraform.Output(t, opts, "vpc_region"),
		expectedZoneID:    terraform.Output(t, opts, "expected_zone_id"),
		expectedVPCID:     terraform.Output(t, opts, "expected_vpc_id"),
		expectedVPCRegion: terraform.Output(t, opts, "expected_vpc_region"),
	}
}

func assertOutputContract(t *testing.T, outputs zoneAssociationOutputs) {
	t.Helper()

	assert.Regexp(t, "^Z[A-Z0-9]+$", outputs.zoneID, "zone_id must be a Route53 hosted zone ID")
	assert.Regexp(t, "^vpc-[0-9a-f]+$", outputs.vpcID, "vpc_id must be a VPC ID")
	assert.Regexp(t, "^[0-9]{12}$", outputs.owningAccount, "owning_account must be a 12-digit AWS account ID")
	assert.Equal(t, outputs.expectedZoneID, outputs.zoneID, "zone_id output must match the created hosted zone")
	assert.Equal(t, outputs.expectedVPCID, outputs.vpcID, "vpc_id output must match the associated secondary VPC")
	assert.Equal(t, outputs.expectedVPCRegion, outputs.vpcRegion, "vpc_region output must match the provider region")

	idParts := strings.Split(outputs.id, ":")
	require.GreaterOrEqual(t, len(idParts), 2, "id must include at least zone_id and vpc_id")
	assert.Equal(t, outputs.zoneID, idParts[0], "id must embed zone_id as the first segment")
	assert.Equal(t, outputs.vpcID, idParts[1], "id must embed vpc_id as the second segment")
}

func verifyAssociationViaAPI(t *testing.T, client *route53.Client, outputs zoneAssociationOutputs) string {
	t.Helper()
	result, err := client.GetHostedZone(context.Background(), &route53.GetHostedZoneInput{
		Id: aws.String(outputs.zoneID),
	})
	require.NoError(t, err, "failed to get hosted zone")
	require.NotNil(t, result.HostedZone, "hosted zone response must include HostedZone")
	require.NotNil(t, result.VPCs, "VPCs list must not be nil")

	zoneIDFromAPI := strings.TrimPrefix(aws.ToString(result.HostedZone.Id), "/hostedzone/")
	assert.Equal(t, outputs.zoneID, zoneIDFromAPI, "zone ID from API must match Terraform output")

	found := false
	for _, vpc := range result.VPCs {
		if aws.ToString(vpc.VPCId) == outputs.vpcID {
			found = true
			break
		}
	}
	assert.True(t, found, "secondary VPC %s should be associated with zone %s", outputs.vpcID, outputs.zoneID)

	return strings.TrimSuffix(aws.ToString(result.HostedZone.Name), ".")
}

func assertOwningAccountMatchesCaller(t *testing.T, cfg aws.Config, owningAccount string) {
	t.Helper()
	stsClient := sts.NewFromConfig(cfg)
	identity, err := stsClient.GetCallerIdentity(context.Background(), &sts.GetCallerIdentityInput{})
	require.NoError(t, err, "failed to get caller identity")
	assert.Equal(t, owningAccount, aws.ToString(identity.Account), "owning_account should match the caller account in this example")
}

func createAndDeleteRoute53Record(t *testing.T, client *route53.Client, zoneID, zoneName string) {
	t.Helper()
	recordName := fmt.Sprintf("terratest-zone-association-%d.%s", time.Now().UnixNano(), zoneName)
	createInput := &route53.ChangeResourceRecordSetsInput{
		HostedZoneId: aws.String(zoneID),
		ChangeBatch: &route53types.ChangeBatch{
			Changes: []route53types.Change{
				{
					Action: route53types.ChangeActionUpsert,
					ResourceRecordSet: &route53types.ResourceRecordSet{
						Name: aws.String(recordName),
						Type: route53types.RRTypeA,
						TTL:  aws.Int64(60),
						ResourceRecords: []route53types.ResourceRecord{
							{Value: aws.String("192.0.2.1")},
						},
					},
				},
			},
		},
	}
	_, err := client.ChangeResourceRecordSets(context.Background(), createInput)
	require.NoError(t, err, "failed to create/write Route53 test record")

	t.Cleanup(func() {
		deleteInput := &route53.ChangeResourceRecordSetsInput{
			HostedZoneId: aws.String(zoneID),
			ChangeBatch: &route53types.ChangeBatch{
				Changes: []route53types.Change{
					{
						Action: route53types.ChangeActionDelete,
						ResourceRecordSet: &route53types.ResourceRecordSet{
							Name: aws.String(recordName),
							Type: route53types.RRTypeA,
							TTL:  aws.Int64(60),
							ResourceRecords: []route53types.ResourceRecord{
								{Value: aws.String("192.0.2.1")},
							},
						},
					},
				},
			},
		}
		_, deleteErr := client.ChangeResourceRecordSets(context.Background(), deleteInput)
		if deleteErr != nil {
			t.Logf("warning: failed to delete Route53 test record %s: %v", recordName, deleteErr)
		}
	})
}

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	outputs := getOutputs(t, ctx)
	assertOutputContract(t, outputs)

	cfg, err := config.LoadDefaultConfig(context.Background())
	require.NoError(t, err, "failed to load AWS config")

	assertOwningAccountMatchesCaller(t, cfg, outputs.owningAccount)
	route53Client := route53.NewFromConfig(cfg)
	zoneName := verifyAssociationViaAPI(t, route53Client, outputs)

	t.Run("TestRoute53WriteOperation", func(t *testing.T) {
		createAndDeleteRoute53Record(t, route53Client, outputs.zoneID, zoneName)
	})
}

func TestComposableCompleteReadonly(t *testing.T, ctx types.TestContext) {
	outputs := getOutputs(t, ctx)
	assertOutputContract(t, outputs)

	cfg, err := config.LoadDefaultConfig(context.Background())
	require.NoError(t, err, "failed to load AWS config")

	assertOwningAccountMatchesCaller(t, cfg, outputs.owningAccount)
	route53Client := route53.NewFromConfig(cfg)
	verifyAssociationViaAPI(t, route53Client, outputs)
}
