// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeOnPrem.Inputs
{

    public sealed class BareMetalClusterNetworkConfigSrIovConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to install the SR-IOV operator.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public BareMetalClusterNetworkConfigSrIovConfigArgs()
        {
        }
        public static new BareMetalClusterNetworkConfigSrIovConfigArgs Empty => new BareMetalClusterNetworkConfigSrIovConfigArgs();
    }
}
