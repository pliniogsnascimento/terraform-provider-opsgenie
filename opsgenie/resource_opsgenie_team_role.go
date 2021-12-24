package opsgenie

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceOpsGenieTeamRole() *schema.Resource {
	return &schema.Resource{
		Create:        resourceOpsGenieTeamRoleCreate,
		ReadContext:   resourceOpsGenieReadContext,
		UpdateContext: resourceOpsGenieUpdateContext,
		DeleteContext: resourceOpsGenieDeleteContext,
		// Importer: ,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateOpsGenieteamRoleName,
			},
			"team_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"granted_rights": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice(validCustomRolesRights, false),
				},
				Set: schema.HashString,
			},
			"disallowed_rights": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice(validCustomRolesRights, false),
				},
				Set: schema.HashString,
			},
		},
	}
}

func resourceOpsGenieTeamRoleCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceOpsGenieReadContext(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics {
	return nil
}

func resourceOpsGenieUpdateContext(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics {
	return nil
}

func resourceOpsGenieDeleteContext(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics {
	return nil
}

func validateOpsGenieteamRoleName(interface{}, string) ([]string, []error) {
	return nil, nil
}

func validateOpsGenieTeamRoleTeamId(i interface{}, s string) ([]string, []error) {
	return nil, nil
}
