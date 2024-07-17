// Copyright The NRI Plugins Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package watch

import (
	logger "github.com/containers/nri-plugins/pkg/log"
	k8swatch "k8s.io/apimachinery/pkg/watch"
)

type (
	Interface = k8swatch.Interface
	EventType = k8swatch.EventType
	Event     = k8swatch.Event
)

const (
	Added    = k8swatch.Added
	Modified = k8swatch.Modified
	Deleted  = k8swatch.Deleted
	Bookmark = k8swatch.Bookmark
	Error    = k8swatch.Error
)

var (
	log = logger.Get("watch")
)
