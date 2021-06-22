// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class ServiceAttachmentConsumerAcceptListGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of consumer forwarding rules the consumer project can
        /// create.
        /// </summary>
        [Input("connectionLimit", required: true)]
        public Input<int> ConnectionLimit { get; set; } = null!;

        /// <summary>
        /// A project that is allowed to connect to this service attachment.
        /// </summary>
        [Input("projectIdOrNum", required: true)]
        public Input<string> ProjectIdOrNum { get; set; } = null!;

        public ServiceAttachmentConsumerAcceptListGetArgs()
        {
        }
    }
}
