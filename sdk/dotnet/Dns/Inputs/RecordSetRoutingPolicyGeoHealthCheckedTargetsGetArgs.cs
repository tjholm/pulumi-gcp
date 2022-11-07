// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dns.Inputs
{

    public sealed class RecordSetRoutingPolicyGeoHealthCheckedTargetsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("internalLoadBalancers", required: true)]
        private InputList<Inputs.RecordSetRoutingPolicyGeoHealthCheckedTargetsInternalLoadBalancerGetArgs>? _internalLoadBalancers;

        /// <summary>
        /// The list of internal load balancers to health check.
        /// Structure is document below.
        /// </summary>
        public InputList<Inputs.RecordSetRoutingPolicyGeoHealthCheckedTargetsInternalLoadBalancerGetArgs> InternalLoadBalancers
        {
            get => _internalLoadBalancers ?? (_internalLoadBalancers = new InputList<Inputs.RecordSetRoutingPolicyGeoHealthCheckedTargetsInternalLoadBalancerGetArgs>());
            set => _internalLoadBalancers = value;
        }

        public RecordSetRoutingPolicyGeoHealthCheckedTargetsGetArgs()
        {
        }
        public static new RecordSetRoutingPolicyGeoHealthCheckedTargetsGetArgs Empty => new RecordSetRoutingPolicyGeoHealthCheckedTargetsGetArgs();
    }
}
