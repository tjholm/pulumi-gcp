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
    public sealed class EndpointPolicyEndpointMatcherMetadataLabelMatcher
    {
        /// <summary>
        /// Specifies how matching should be done.
        /// Possible values are: `MATCH_ANY`, `MATCH_ALL`.
        /// </summary>
        public readonly string MetadataLabelMatchCriteria;
        /// <summary>
        /// The list of label value pairs that must match labels in the provided metadata based on filterMatchCriteria
        /// Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabel> MetadataLabels;

        [OutputConstructor]
        private EndpointPolicyEndpointMatcherMetadataLabelMatcher(
            string metadataLabelMatchCriteria,

            ImmutableArray<Outputs.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabel> metadataLabels)
        {
            MetadataLabelMatchCriteria = metadataLabelMatchCriteria;
            MetadataLabels = metadataLabels;
        }
    }
}
