// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class NodeGroupShareSettingsProjectMapArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identifier for this object. Format specified above.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The project id/number should be the same as the key of this project config in the project map.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public NodeGroupShareSettingsProjectMapArgs()
        {
        }
        public static new NodeGroupShareSettingsProjectMapArgs Empty => new NodeGroupShareSettingsProjectMapArgs();
    }
}
