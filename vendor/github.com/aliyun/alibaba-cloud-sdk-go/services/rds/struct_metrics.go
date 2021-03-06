package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Metrics is a nested struct in rds response
type Metrics struct {
	Unit            string `json:"Unit" xml:"Unit"`
	MetricsKey      string `json:"MetricsKey" xml:"MetricsKey"`
	Method          string `json:"Method" xml:"Method"`
	Dimension       string `json:"Dimension" xml:"Dimension"`
	SortRule        int    `json:"SortRule" xml:"SortRule"`
	GroupKeyType    string `json:"GroupKeyType" xml:"GroupKeyType"`
	MetricsKeyAlias string `json:"MetricsKeyAlias" xml:"MetricsKeyAlias"`
	GroupKey        string `json:"GroupKey" xml:"GroupKey"`
	Description     string `json:"Description" xml:"Description"`
	DbType          string `json:"DbType" xml:"DbType"`
}
