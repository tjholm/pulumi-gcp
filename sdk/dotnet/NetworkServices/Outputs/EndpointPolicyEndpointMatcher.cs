// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.NetworkServices.Outputs
{

    [OutputType]
    public sealed class EndpointPolicyEndpointMatcher
    {
        /// <summary>
        /// The matcher is based on node metadata presented by xDS clients.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.EndpointPolicyEndpointMatcherMetadataLabelMatcher MetadataLabelMatcher;

        [OutputConstructor]
        private EndpointPolicyEndpointMatcher(Outputs.EndpointPolicyEndpointMatcherMetadataLabelMatcher metadataLabelMatcher)
        {
            MetadataLabelMatcher = metadataLabelMatcher;
        }
    }
}
