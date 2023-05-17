// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeOnPrem.Inputs
{

    public sealed class VMwareClusterFleetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Output)
        /// The name of the managed Hub Membership resource associated to this cluster.
        /// Membership names are formatted as
        /// `projects/&lt;project-number&gt;/locations/&lt;location&gt;/memberships/&lt;cluster-id&gt;`.
        /// </summary>
        [Input("membership")]
        public Input<string>? Membership { get; set; }

        public VMwareClusterFleetArgs()
        {
        }
        public static new VMwareClusterFleetArgs Empty => new VMwareClusterFleetArgs();
    }
}
