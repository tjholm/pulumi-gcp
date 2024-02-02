// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Workstations.Inputs
{

    public sealed class WorkstationConfigReadinessCheckGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path to which the request should be sent.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// Port to which the request should be sent.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        public WorkstationConfigReadinessCheckGetArgs()
        {
        }
        public static new WorkstationConfigReadinessCheckGetArgs Empty => new WorkstationConfigReadinessCheckGetArgs();
    }
}
