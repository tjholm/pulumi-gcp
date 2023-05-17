// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeOnPrem.Inputs
{

    public sealed class VMwareClusterDataplaneV2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable advanced networking which requires dataplane_v2_enabled to be set true.
        /// </summary>
        [Input("advancedNetworking")]
        public Input<bool>? AdvancedNetworking { get; set; }

        /// <summary>
        /// Enables Dataplane V2.
        /// </summary>
        [Input("dataplaneV2Enabled")]
        public Input<bool>? DataplaneV2Enabled { get; set; }

        /// <summary>
        /// Enable Dataplane V2 for clusters with Windows nodes.
        /// </summary>
        [Input("windowsDataplaneV2Enabled")]
        public Input<bool>? WindowsDataplaneV2Enabled { get; set; }

        public VMwareClusterDataplaneV2Args()
        {
        }
        public static new VMwareClusterDataplaneV2Args Empty => new VMwareClusterDataplaneV2Args();
    }
}
