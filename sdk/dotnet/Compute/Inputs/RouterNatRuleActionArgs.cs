// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RouterNatRuleActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceNatActiveIps")]
        private InputList<string>? _sourceNatActiveIps;

        /// <summary>
        /// A list of URLs of the IP resources used for this NAT rule.
        /// These IP addresses must be valid static external IP addresses assigned to the project.
        /// This field is used for public NAT.
        /// </summary>
        public InputList<string> SourceNatActiveIps
        {
            get => _sourceNatActiveIps ?? (_sourceNatActiveIps = new InputList<string>());
            set => _sourceNatActiveIps = value;
        }

        [Input("sourceNatActiveRanges")]
        private InputList<string>? _sourceNatActiveRanges;

        /// <summary>
        /// A list of URLs of the subnetworks used as source ranges for this NAT Rule.
        /// These subnetworks must have purpose set to PRIVATE_NAT.
        /// This field is used for private NAT.
        /// </summary>
        public InputList<string> SourceNatActiveRanges
        {
            get => _sourceNatActiveRanges ?? (_sourceNatActiveRanges = new InputList<string>());
            set => _sourceNatActiveRanges = value;
        }

        [Input("sourceNatDrainIps")]
        private InputList<string>? _sourceNatDrainIps;

        /// <summary>
        /// A list of URLs of the IP resources to be drained.
        /// These IPs must be valid static external IPs that have been assigned to the NAT.
        /// These IPs should be used for updating/patching a NAT rule only.
        /// This field is used for public NAT.
        /// </summary>
        public InputList<string> SourceNatDrainIps
        {
            get => _sourceNatDrainIps ?? (_sourceNatDrainIps = new InputList<string>());
            set => _sourceNatDrainIps = value;
        }

        [Input("sourceNatDrainRanges")]
        private InputList<string>? _sourceNatDrainRanges;

        /// <summary>
        /// A list of URLs of subnetworks representing source ranges to be drained.
        /// This is only supported on patch/update, and these subnetworks must have previously been used as active ranges in this NAT Rule.
        /// This field is used for private NAT.
        /// </summary>
        public InputList<string> SourceNatDrainRanges
        {
            get => _sourceNatDrainRanges ?? (_sourceNatDrainRanges = new InputList<string>());
            set => _sourceNatDrainRanges = value;
        }

        public RouterNatRuleActionArgs()
        {
        }
        public static new RouterNatRuleActionArgs Empty => new RouterNatRuleActionArgs();
    }
}
