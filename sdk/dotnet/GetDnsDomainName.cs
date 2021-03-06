// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Civo
{
    public static class GetDnsDomainName
    {
        public static Task<GetDnsDomainNameResult> InvokeAsync(GetDnsDomainNameArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDnsDomainNameResult>("civo:index/getDnsDomainName:getDnsDomainName", args ?? new GetDnsDomainNameArgs(), options.WithVersion());
    }


    public sealed class GetDnsDomainNameArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the domain.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the domain.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetDnsDomainNameArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDnsDomainNameResult
    {
        /// <summary>
        /// A unique ID that can be used to identify and reference a domain.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the domain.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private GetDnsDomainNameResult(
            string? id,

            string? name)
        {
            Id = id;
            Name = name;
        }
    }
}
