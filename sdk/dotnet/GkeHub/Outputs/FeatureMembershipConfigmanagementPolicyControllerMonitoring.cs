// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeHub.Outputs
{

    [OutputType]
    public sealed class FeatureMembershipConfigmanagementPolicyControllerMonitoring
    {
        /// <summary>
        /// Specifies the list of backends Policy Controller will export to. Must be one of `CLOUD_MONITORING` or `PROMETHEUS`. Defaults to [`CLOUD_MONITORING`, `PROMETHEUS`]. Specifying an empty value `[]` disables metrics export.
        /// </summary>
        public readonly ImmutableArray<string> Backends;

        [OutputConstructor]
        private FeatureMembershipConfigmanagementPolicyControllerMonitoring(ImmutableArray<string> backends)
        {
            Backends = backends;
        }
    }
}
