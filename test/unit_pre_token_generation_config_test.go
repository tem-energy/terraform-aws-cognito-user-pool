package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitPreTokenGenerationConfig(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "unit-pre-token-generation-config",
		EnvVars: map[string]string{
			"AWS_ACCESS_KEY_ID":     "test",
			"AWS_SECRET_ACCESS_KEY": "test",
		},
		Upgrade: true,
	}

	plan := terraform.InitAndPlanAndShowWithStructNoLogTempPlanFile(t, terraformOptions)
	resource, exists := plan.ResourcePlannedValuesMap["module.test.aws_cognito_user_pool.user_pool[0]"]
	require.True(t, exists, "expected Cognito user pool in plan")

	lambdaConfig, exists := resource.AttributeValues["lambda_config"]
	require.True(t, exists, "expected Lambda configuration in plan")
	lambdaConfigs, ok := lambdaConfig.([]interface{})
	require.True(t, ok)
	require.Len(t, lambdaConfigs, 1)

	lambdaConfigValues, ok := lambdaConfigs[0].(map[string]interface{})
	require.True(t, ok)
	preTokenConfigs, ok := lambdaConfigValues["pre_token_generation_config"].([]interface{})
	require.True(t, ok)
	require.Len(t, preTokenConfigs, 1)

	preTokenConfig, ok := preTokenConfigs[0].(map[string]interface{})
	require.True(t, ok)
	assert.Equal(t, "arn:aws:lambda:us-east-1:123456789012:function:pre-token-generation", preTokenConfig["lambda_arn"])
	assert.Equal(t, "V2_0", preTokenConfig["lambda_version"])
}
