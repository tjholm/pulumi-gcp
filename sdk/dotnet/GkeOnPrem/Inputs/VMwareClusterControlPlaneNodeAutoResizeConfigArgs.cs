// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeOnPrem.Inputs
{

    public sealed class VMwareClusterControlPlaneNodeAutoResizeConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable control plane node auto resizing.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        public VMwareClusterControlPlaneNodeAutoResizeConfigArgs()
        {
        }
        public static new VMwareClusterControlPlaneNodeAutoResizeConfigArgs Empty => new VMwareClusterControlPlaneNodeAutoResizeConfigArgs();
    }
}
