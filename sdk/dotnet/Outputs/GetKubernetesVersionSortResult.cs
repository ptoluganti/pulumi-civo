// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Civo.Outputs
{

    [OutputType]
    public sealed class GetKubernetesVersionSortResult
    {
        /// <summary>
        /// The sort direction. This may be either `asc` or `desc`.
        /// </summary>
        public readonly string? Direction;
        /// <summary>
        /// Sort the sizes by this key. This may be one of `version`.
        /// </summary>
        public readonly string Key;

        [OutputConstructor]
        private GetKubernetesVersionSortResult(
            string? direction,

            string key)
        {
            Direction = direction;
            Key = key;
        }
    }
}
