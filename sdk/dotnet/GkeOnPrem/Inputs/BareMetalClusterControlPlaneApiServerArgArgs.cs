// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeOnPrem.Inputs
{

    public sealed class BareMetalClusterControlPlaneApiServerArgArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The argument name as it appears on the API Server command line please make sure to remove the leading dashes.
        /// </summary>
        [Input("argument", required: true)]
        public Input<string> Argument { get; set; } = null!;

        /// <summary>
        /// The value of the arg as it will be passed to the API Server command line.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public BareMetalClusterControlPlaneApiServerArgArgs()
        {
        }
        public static new BareMetalClusterControlPlaneApiServerArgArgs Empty => new BareMetalClusterControlPlaneApiServerArgArgs();
    }
}
