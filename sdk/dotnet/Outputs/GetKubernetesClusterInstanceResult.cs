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
    public sealed class GetKubernetesClusterInstanceResult
    {
        /// <summary>
        /// Total cpu of the inatance.
        /// </summary>
        public readonly int CpuCores;
        /// <summary>
        /// The date where the Kubernetes cluster was create.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The size of the disk.
        /// </summary>
        public readonly int DiskGb;
        /// <summary>
        /// The firewall id assigned to the instance
        /// </summary>
        public readonly string FirewallId;
        /// <summary>
        /// The hostname of the instance.
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// The public ip of the instances, only available if the instances is the master
        /// </summary>
        public readonly string PublicIp;
        /// <summary>
        /// Total ram of the instance
        /// </summary>
        public readonly int RamMb;
        /// <summary>
        /// The region where instance are.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The size of the instance.
        /// </summary>
        public readonly string Size;
        /// <summary>
        /// The status of Kubernetes cluster.
        /// * `ready` -If the Kubernetes cluster is ready.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The tag of the instances
        /// </summary>
        public readonly ImmutableArray<string> Tags;

        [OutputConstructor]
        private GetKubernetesClusterInstanceResult(
            int cpuCores,

            string createdAt,

            int diskGb,

            string firewallId,

            string hostname,

            string publicIp,

            int ramMb,

            string region,

            string size,

            string status,

            ImmutableArray<string> tags)
        {
            CpuCores = cpuCores;
            CreatedAt = createdAt;
            DiskGb = diskGb;
            FirewallId = firewallId;
            Hostname = hostname;
            PublicIp = publicIp;
            RamMb = ramMb;
            Region = region;
            Size = size;
            Status = status;
            Tags = tags;
        }
    }
}
