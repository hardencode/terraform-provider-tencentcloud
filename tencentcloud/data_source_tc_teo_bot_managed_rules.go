/*
Use this data source to query detailed information of teo botManagedRules

Example Usage

```hcl
data "tencentcloud_teo_bot_managed_rules" "botManagedRules" {
  zone_id = ""
  entity = ""
  rule_type = ""
  }
```
*/

package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	teo "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/teo/v20220106"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func dataSourceTencentCloudTeoBotManagedRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudTeoBotManagedRulesRead,
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Site ID.",
			},

			"entity": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Subdomain or application name.",
			},

			"rule_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "rule type, idcid/sipbot/uabot.",
			},

			"rules": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Managed rules list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Rule ID.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the rule.",
						},
						"rule_type_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the rule.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of the rule.",
						},
					},
				},
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudTeoBotManagedRulesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloud_teo_bot_managed_rules.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	param := make(map[string]interface{})
	if v, ok := d.GetOk("zone_id"); ok {
		param["zone_id"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("entity"); ok {
		param["entity"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("rule_type"); ok {
		param["rule_type"] = helper.String(v.(string))
	}

	teoService := TeoService{client: meta.(*TencentCloudClient).apiV3Conn}

	var rules []*teo.BotManagedRuleDetail
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		results, e := teoService.DescribeTeoBotManagedRulesByFilter(ctx, param)
		if e != nil {
			return retryError(e)
		}
		rules = results
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read Teo rules failed, reason:%+v", logId, err)
		return err
	}

	if rules != nil {
		ruleList := []interface{}{}
		for _, rule := range rules {
			ruleMap := map[string]interface{}{}
			if rule.RuleId != nil {
				ruleMap["rule_id"] = rule.RuleId
			}
			if rule.Description != nil {
				ruleMap["description"] = rule.Description
			}
			if rule.RuleTypeName != nil {
				ruleMap["rule_type_name"] = rule.RuleTypeName
			}
			if rule.Status != nil {
				ruleMap["status"] = rule.Status
			}

			ruleList = append(ruleList, ruleMap)
		}
		_ = d.Set("rules", ruleList)
	}

	return nil
}
