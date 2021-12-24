package opsgenie

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/opsgenie/opsgenie-go-sdk-v2/team"
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
			"rights": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"right": {
							Type:     schema.TypeString,
							Required: true,
						},
						"granted": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
					},
				},
			},
		},
	}
}

// TODO: Read and set values before returning
func resourceOpsGenieTeamRoleCreate(d *schema.ResourceData, meta interface{}) error {
	client, err := team.NewClient(meta.(*OpsgenieClient).client.Config)
	if err != nil {
		return err
	}

	name := d.Get("name").(string)
	teamId := d.Get("team_id").(string)

	createRequest := team.CreateTeamRoleRequest{
		TeamIdentifierType:  team.Id,
		TeamIdentifierValue: teamId,
		Name:                name,
	}
	createRequest.Rights = expandOpsGenieTeamRoleRights(d)

	_, err = client.CreateRole(context.Background(), &createRequest)
	if err != nil {
		return err
	}

	getRequest := &team.GetTeamRoleRequest{
		TeamID:   teamId,
		RoleName: name,
	}

	getResponse, err := client.GetRole(context.Background(), getRequest)
	if err != nil {
		return err
	}

	d.SetId(getResponse.Id)
	return nil
}

func expandOpsGenieTeamRoleRights(d *schema.ResourceData) []team.Right {
	input := d.Get("rights").([]interface{})
	rights := make([]team.Right, 0, len(input))

	for _, v := range input {
		config := v.(map[string]interface{})

		rightName := config["right"].(string)
		granted := config["granted"].(bool)

		right := team.Right{
			Right:   rightName,
			Granted: &granted,
		}

		rights = append(rights, right)
	}

	return rights
}

func resourceOpsGenieReadContext(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics {
	panic("Not implemented")
}

func resourceOpsGenieUpdateContext(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics {
	panic("Not implemented")
}

func resourceOpsGenieDeleteContext(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics {
	panic("Not implemented")
}

func validateOpsGenieteamRoleName(interface{}, string) ([]string, []error) {
	panic("Not implemented")
}

func validateOpsGenieTeamRoleTeamId(i interface{}, s string) ([]string, []error) {
	panic("Not implemented")
}
