package azuredevops

import (
	//"fmt"
	//"log"

	//"github.com/microsoft/terraform-provider-azuredevops/azuredevops/utils/converter"

	//"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	//"github.com/microsoft/azure-devops-go-api/azuredevops/taskagent"
)

func resourceVariableGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceVariableGroupCreate,
		Read:   resourceVariableGroupRead,
		Update: resourceVariableGroupUpdate,
		Delete: resourceVariableGroupDelete,

		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"allow_access": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"variables": {
				Type: schema.TypeMap,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"is_secret": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
				Required: true,
			},
		},
	}
}

func resourceVariableGroupCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceVariableGroupRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceVariableGroupUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceVariableGroupDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
