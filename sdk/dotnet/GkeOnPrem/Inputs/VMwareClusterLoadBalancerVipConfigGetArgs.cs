// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeOnPrem.Inputs
{

    public sealed class VMwareClusterLoadBalancerVipConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The VIP which you previously set aside for the Kubernetes API of this cluster.
        /// </summary>
        [Input("controlPlaneVip")]
        public Input<string>? ControlPlaneVip { get; set; }

        /// <summary>
        /// The VIP which you previously set aside for ingress traffic into this cluster.
        /// </summary>
        [Input("ingressVip")]
        public Input<string>? IngressVip { get; set; }

        public VMwareClusterLoadBalancerVipConfigGetArgs()
        {
        }
        public static new VMwareClusterLoadBalancerVipConfigGetArgs Empty => new VMwareClusterLoadBalancerVipConfigGetArgs();
    }
}
