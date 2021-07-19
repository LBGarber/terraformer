// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package linode

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/linode/linodego"
)

type ObjectStorageGenerator struct {
	LinodeService
}

func (g ObjectStorageGenerator) createResources(bucketList []linodego.ObjectStorageBucket) []terraformutils.Resource {
	var resources []terraformutils.Resource

	for _, bucket := range bucketList {
		id := fmt.Sprintf("%s:%s", bucket.Cluster, bucket.Label)

		resources = append(resources, terraformutils.NewResource(
			id,
			id,
			"linode_object_storage_bucket",
			"linode",
			map[string]string{
				"cluster": bucket.Cluster,
				"label": bucket.Label,
			},
			[]string{},
			map[string]interface{}{}))
	}
	return resources
}

func (g *ObjectStorageGenerator) InitResources() error {
	client := g.generateClient()
	buckets, err := client.ListObjectStorageBuckets(context.Background(), nil)
	if err != nil {
		return err
	}
	
	g.Resources = g.createResources(buckets)
	
	return nil
}
