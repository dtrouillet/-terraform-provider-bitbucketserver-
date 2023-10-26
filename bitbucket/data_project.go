package bitbucket

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
)

func dataSourceProject() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceProjectRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"avatar": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"repos": {
				Type:     schema.TypeList,
				Computed: true,
			},
		},
	}
}

func dataSourceProjectRead(d *schema.ResourceData, m interface{}) error {
	d.SetId(d.Get("key").(string))
	err := resourceProjectRead(d, m)
	if err != nil {
		return err
	}

	project := d.Get("key").(string)

	client := m.(*BitbucketServerProvider).BitbucketClient
	project_repos_req, err := client.Get(fmt.Sprintf("/rest/api/1.0/projects/%s/repos",
		project,
	))

	if err != nil {
		return err
	}

	if project_repos_req.StatusCode == 200 {

		var repos []Repository

		body, readerr := ioutil.ReadAll(project_repos_req.Body)
		if readerr != nil {
			return readerr
		}

		decodeerr := json.Unmarshal(body, &repos)
		if decodeerr != nil {
			return decodeerr
		}

		slugs := make([]string, 0)
		for _, repo := range repos {
			slugs = append(slugs, repo.Slug)
		}

		_ = d.Set("repos", slugs)

	}

	return nil
}
