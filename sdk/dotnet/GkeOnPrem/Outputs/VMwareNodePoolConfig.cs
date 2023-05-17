// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeOnPrem.Outputs
{

    [OutputType]
    public sealed class VMwareNodePoolConfig
    {
        /// <summary>
        /// VMware disk size to be used during creation.
        /// </summary>
        public readonly int? BootDiskSizeGb;
        /// <summary>
        /// The number of CPUs for each node in the node pool.
        /// </summary>
        public readonly int? Cpus;
        /// <summary>
        /// Allow node pool traffic to be load balanced. Only works for clusters with
        /// MetalLB load balancers.
        /// </summary>
        public readonly bool? EnableLoadBalancer;
        /// <summary>
        /// The OS image name in vCenter, only valid when using Windows.
        /// </summary>
        public readonly string? Image;
        /// <summary>
        /// The OS image to be used for each node in a node pool.
        /// Currently `cos`, `ubuntu`, `ubuntu_containerd` and `windows` are supported.
        /// </summary>
        public readonly string ImageType;
        /// <summary>
        /// The map of Kubernetes labels (key/value pairs) to be applied to each node.
        /// These will added in addition to any default label(s) that
        /// Kubernetes may apply to the node.
        /// In case of conflict in label keys, the applied set may differ depending on
        /// the Kubernetes version -- it's best to assume the behavior is undefined
        /// and conflicts should be avoided.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Labels;
        /// <summary>
        /// The megabytes of memory for each node in the node pool.
        /// </summary>
        public readonly int? MemoryMb;
        /// <summary>
        /// The number of nodes in the node pool.
        /// </summary>
        public readonly int? Replicas;
        /// <summary>
        /// The initial taints assigned to nodes of this node pool.
        /// Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.VMwareNodePoolConfigTaint> Taints;
        /// <summary>
        /// (Output)
        /// Specifies the vSphere config for node pool.
        /// Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.VMwareNodePoolConfigVsphereConfig> VsphereConfigs;

        [OutputConstructor]
        private VMwareNodePoolConfig(
            int? bootDiskSizeGb,

            int? cpus,

            bool? enableLoadBalancer,

            string? image,

            string imageType,

            ImmutableDictionary<string, string>? labels,

            int? memoryMb,

            int? replicas,

            ImmutableArray<Outputs.VMwareNodePoolConfigTaint> taints,

            ImmutableArray<Outputs.VMwareNodePoolConfigVsphereConfig> vsphereConfigs)
        {
            BootDiskSizeGb = bootDiskSizeGb;
            Cpus = cpus;
            EnableLoadBalancer = enableLoadBalancer;
            Image = image;
            ImageType = imageType;
            Labels = labels;
            MemoryMb = memoryMb;
            Replicas = replicas;
            Taints = taints;
            VsphereConfigs = vsphereConfigs;
        }
    }
}
