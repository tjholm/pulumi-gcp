// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Inputs
{

    public sealed class ClusterAddonsConfigGcsFuseCsiDriverConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable Binary Authorization for this cluster. Deprecated in favor of `evaluation_mode`.
        /// for autopilot clusters. Resource limits for `cpu` and `memory` must be defined to enable node auto-provisioning for GKE Standard.
        /// If enabled, pods must be valid under a PodSecurityPolicy to be created.
        /// not.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        public ClusterAddonsConfigGcsFuseCsiDriverConfigArgs()
        {
        }
        public static new ClusterAddonsConfigGcsFuseCsiDriverConfigArgs Empty => new ClusterAddonsConfigGcsFuseCsiDriverConfigArgs();
    }
}
