// Copyright 2016-2018, Pulumi Corporation.
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

package civo

import (
	"unicode"

	"github.com/civo/terraform-provider-civo/civo"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "civo"
	// modules:
	mainMod = "index" // the y module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c *terraform.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := civo.Provider().(*schema.Provider)

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                    p,
		Name:                 "civo",
		Description:          "A Pulumi package for creating and managing Civo cloud resources.",
		Keywords:             []string{"pulumi", "civo"},
		License:              "Apache-2.0",
		Homepage:             "https://pulumi.io",
		Repository:           "https://github.com/pulumi/pulumi-civo",
		GitHubOrg:            "civo",
		Config:               map[string]*tfbridge.SchemaInfo{},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"civo_instance":           {Tok: makeResource(mainMod, "Instance")},
			"civo_network":            {Tok: makeResource(mainMod, "Network")},
			"civo_volume":             {Tok: makeResource(mainMod, "Volume")},
			"civo_volume_attachment":  {Tok: makeResource(mainMod, "VolumeAttachment")},
			"civo_dns_domain_name":    {Tok: makeResource(mainMod, "DnsDomainName")},
			"civo_dns_domain_record":  {Tok: makeResource(mainMod, "DnsDomainRecord")},
			"civo_firewall":           {Tok: makeResource(mainMod, "Firewall")},
			"civo_firewall_rule":      {Tok: makeResource(mainMod, "FirewallRule")},
			"civo_loadbalancer":       {Tok: makeResource(mainMod, "LoadBalancer")},
			"civo_ssh_key":            {Tok: makeResource(mainMod, "SshKey")},
			"civo_template":           {Tok: makeResource(mainMod, "Template")},
			"civo_snapshot":           {Tok: makeResource(mainMod, "Snapshot")},
			"civo_kubernetes_cluster": {Tok: makeResource(mainMod, "KubernetesCluster")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"civo_kubernetes_cluster": {Tok: makeDataSource(mainMod, "getKubernetesCluster")},
			"civo_template":           {Tok: makeDataSource(mainMod, "getTemplate")},
			"civo_kubernetes_version": {Tok: makeDataSource(mainMod, "getKubernetesVersion")},
			"civo_instances_size":     {Tok: makeDataSource(mainMod, "getInstancesSize")},
			"civo_instances":          {Tok: makeDataSource(mainMod, "getInstances")},
			"civo_instance":           {Tok: makeDataSource(mainMod, "getInstance")},
			"civo_dns_domain_name":    {Tok: makeDataSource(mainMod, "getDnsDomainName")},
			"civo_dns_domain_record":  {Tok: makeDataSource(mainMod, "getDnsDomainRecord")},
			"civo_network":            {Tok: makeDataSource(mainMod, "getNetwork")},
			"civo_volume":             {Tok: makeDataSource(mainMod, "getVolume")},
			"civo_loadbalancer":       {Tok: makeDataSource(mainMod, "getLoadBalancer")},
			"civo_ssh_key":            {Tok: makeDataSource(mainMod, "getSshKey")},
			"civo_snapshot":           {Tok: makeDataSource(mainMod, "getSnapshot")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=2.0.0,<3.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: map[string]string{
				"civo": "Civo",
			},
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameProperty = "name"
	for resname, res := range prov.Resources {
		if resSchema := p.ResourcesMap[resname]; resSchema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := resSchema.Schema[nameProperty]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameProperty]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameProperty] = tfbridge.AutoName(nameProperty, 255)
				}
			}
		}
	}

	return prov
}
