// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeOnPrem.Inputs
{

    public sealed class BareMetalClusterLoadBalancerManualLbConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether manual load balancing is enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        public BareMetalClusterLoadBalancerManualLbConfigGetArgs()
        {
        }
        public static new BareMetalClusterLoadBalancerManualLbConfigGetArgs Empty => new BareMetalClusterLoadBalancerManualLbConfigGetArgs();
    }
}
