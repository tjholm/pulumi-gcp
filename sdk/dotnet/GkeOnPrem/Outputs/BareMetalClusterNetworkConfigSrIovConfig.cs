// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeOnPrem.Outputs
{

    [OutputType]
    public sealed class BareMetalClusterNetworkConfigSrIovConfig
    {
        /// <summary>
        /// Whether to install the SR-IOV operator.
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private BareMetalClusterNetworkConfigSrIovConfig(bool? enabled)
        {
            Enabled = enabled;
        }
    }
}
