// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package deploymentdatasource

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func newEnterpriseSearchResourceInfo() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"elasticsearch_cluster_ref_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"healthy": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"http_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"https_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ref_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"topology": enterpriseSearchTopologySchema(),
		},
	}
}

func enterpriseSearchTopologySchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"instance_configuration_id": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"memory_per_node": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"zone_count": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"node_type_appserver": {
					Type:     schema.TypeBool,
					Computed: true,
				},

				"node_type_connector": {
					Type:     schema.TypeBool,
					Computed: true,
				},

				"node_type_worker": {
					Type:     schema.TypeBool,
					Computed: true,
				},
			},
		},
	}
}
