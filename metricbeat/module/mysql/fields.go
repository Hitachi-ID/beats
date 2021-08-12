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

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package mysql

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "mysql", asset.ModuleFieldsPri, AssetMysql); err != nil {
		panic(err)
	}
}

// AssetMysql returns asset data.
// This is the base64 encoded zlib format compressed contents of module/mysql.
func AssetMysql() string {
	return "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"
}
