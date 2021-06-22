// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class SecurityPolicyAdaptiveProtectionConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ) Configuration for [Google Cloud Armor Adaptive Protection Layer 7 DDoS Defense](https://cloud.google.com/armor/docs/adaptive-protection-overview?hl=en). Structure is documented below.
        /// </summary>
        [Input("layer7DdosDefenseConfig")]
        public Input<Inputs.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigGetArgs>? Layer7DdosDefenseConfig { get; set; }

        public SecurityPolicyAdaptiveProtectionConfigGetArgs()
        {
        }
    }
}
