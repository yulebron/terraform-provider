package alicloud

import (
	"fmt"
	"strings"
	"github.com/hashicorp/terraform/helper/schema")

func resourceAliyunCen() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliyunCenCreate,
		Read:   resourceAliyunCenRead,
		Update: resourceAliyunCenUpdate,
		Delete: resourceAliyunCenDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					if len(value) < 2 || len(value) > 128 {
						errors = append(errors, fmt.Errorf("%s cannot be shorter than 2 characters or longer than 128 characters", k))
					}

					if strings.HasPrefix(value, "http://") || strings.HasPrefix(value, "https://") {
						errors = append(errors, fmt.Errorf("%s cannot starts with http:// or https://", k))
					}

					return
				},
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					if len(value) < 2 || len(value) > 256 {
						errors = append(errors, fmt.Errorf("%s cannot be shorter than 2 characters or longer than 256 characters", k))
					}

					if strings.HasPrefix(value, "http://") || strings.HasPrefix(value, "https://") {
						errors = append(errors, fmt.Errorf("%s cannot starts with http:// or https://", k))
					}

					return
				},
			},
		},
	}
}

func resourceAliyunCenCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAliyunCenRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*AliyunClient)

	resp, err := client.DescribeCen(d.Id())

	if err != nil {
		if NotFoundError(err) {
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", resp.Name)
	d.Set("description", resp.Description)

	return nil
}

func resourceAliyunCenUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAliyunCenDelete(d *schema.ResourceData, meta interface{}) error  {
	return nil
}
