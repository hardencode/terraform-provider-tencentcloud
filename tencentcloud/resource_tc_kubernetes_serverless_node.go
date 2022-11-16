/*
TKE超级节点
*/

package tencentcloud

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func resourceTkeServerLessNode() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_id": {
			Type:        schema.TypeString,
			Optional:    false,
			ForceNew:    true,
			Description: "超级节点所属的集群id，必填",
		},
		"subnet_id": {
			Type:        schema.TypeString,
			Optional:    false,
			ForceNew:    true,
			Description: "超级节点对应的子网id，必填，且需校验是否可以加入集群",
		},
		"display_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "超级节点展示名称，选填",
		},
		"node_pool_id": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "节点所属的节点池id，云api接口是必填的，tf这里建议选填甚至不对用户暴露？没有的话都归到一个默认的池子里",
		},
		"phase": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "超级节点状态",
		},
		"created_time": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "超级节点创建时间",
		},
	}
}
