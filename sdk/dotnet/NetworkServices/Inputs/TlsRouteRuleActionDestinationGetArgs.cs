// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.NetworkServices.Inputs
{

    public sealed class TlsRouteRuleActionDestinationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL of a BackendService to route traffic to.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Specifies the proportion of requests forwarded to the backend referenced by the serviceName field.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public TlsRouteRuleActionDestinationGetArgs()
        {
        }
        public static new TlsRouteRuleActionDestinationGetArgs Empty => new TlsRouteRuleActionDestinationGetArgs();
    }
}
