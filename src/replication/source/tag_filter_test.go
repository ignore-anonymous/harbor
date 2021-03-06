// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package source

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmware/harbor/src/replication"
	"github.com/vmware/harbor/src/replication/models"
	"github.com/vmware/harbor/src/replication/registry"
)

func TestInitOfTagFilter(t *testing.T) {
	filter := NewTagFilter("", &registry.HarborAdaptor{})
	assert.Nil(t, filter.Init())
}

func TestGetConvertorOfTagFilter(t *testing.T) {
	filter := NewTagFilter("", &registry.HarborAdaptor{})
	assert.NotNil(t, filter.GetConvertor())
}

func TestDoFilterOfTagFilter(t *testing.T) {
	// invalid filter item type
	filter := NewTagFilter("", &registry.HarborAdaptor{})
	items := filter.DoFilter([]models.FilterItem{
		models.FilterItem{
			Kind: "invalid_type",
		},
	})
	assert.Equal(t, 0, len(items))

	// empty pattern
	filter = NewTagFilter("", &registry.HarborAdaptor{})
	items = filter.DoFilter([]models.FilterItem{
		models.FilterItem{
			Kind:  replication.FilterItemKindTag,
			Value: "library/hello-world:latest",
		},
	})
	assert.Equal(t, 1, len(items))

	// non-empty pattern
	filter = NewTagFilter("l*t", &registry.HarborAdaptor{})
	items = filter.DoFilter([]models.FilterItem{
		models.FilterItem{
			Kind:  replication.FilterItemKindTag,
			Value: "library/hello-world:latest",
		},
	})
	assert.Equal(t, 1, len(items))

	// non-empty pattern
	filter = NewTagFilter("lates?", &registry.HarborAdaptor{})
	items = filter.DoFilter([]models.FilterItem{
		models.FilterItem{
			Kind:  replication.FilterItemKindTag,
			Value: "library/hello-world:latest",
		},
	})
	assert.Equal(t, 1, len(items))

	// non-empty pattern
	filter = NewTagFilter("latest?", &registry.HarborAdaptor{})
	items = filter.DoFilter([]models.FilterItem{
		models.FilterItem{
			Kind:  replication.FilterItemKindTag,
			Value: "library/hello-world:latest",
		},
	})
	assert.Equal(t, 0, len(items))
}
