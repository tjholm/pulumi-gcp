// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class FirewallPolicyRuleMatchArgs : global::Pulumi.ResourceArgs
    {
        [Input("destAddressGroups")]
        private InputList<string>? _destAddressGroups;

        /// <summary>
        /// Address groups which should be matched against the traffic destination. Maximum number of destination address groups is 10. Destination address groups is only supported in Egress rules.
        /// </summary>
        public InputList<string> DestAddressGroups
        {
            get => _destAddressGroups ?? (_destAddressGroups = new InputList<string>());
            set => _destAddressGroups = value;
        }

        [Input("destFqdns")]
        private InputList<string>? _destFqdns;

        /// <summary>
        /// Domain names that will be used to match against the resolved domain name of destination of traffic. Can only be specified if DIRECTION is egress.
        /// </summary>
        public InputList<string> DestFqdns
        {
            get => _destFqdns ?? (_destFqdns = new InputList<string>());
            set => _destFqdns = value;
        }

        [Input("destIpRanges")]
        private InputList<string>? _destIpRanges;

        /// <summary>
        /// CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 256.
        /// </summary>
        public InputList<string> DestIpRanges
        {
            get => _destIpRanges ?? (_destIpRanges = new InputList<string>());
            set => _destIpRanges = value;
        }

        [Input("destRegionCodes")]
        private InputList<string>? _destRegionCodes;

        /// <summary>
        /// The Unicode country codes whose IP addresses will be used to match against the source of traffic. Can only be specified if DIRECTION is egress.
        /// </summary>
        public InputList<string> DestRegionCodes
        {
            get => _destRegionCodes ?? (_destRegionCodes = new InputList<string>());
            set => _destRegionCodes = value;
        }

        [Input("destThreatIntelligences")]
        private InputList<string>? _destThreatIntelligences;

        /// <summary>
        /// Name of the Google Cloud Threat Intelligence list.
        /// </summary>
        public InputList<string> DestThreatIntelligences
        {
            get => _destThreatIntelligences ?? (_destThreatIntelligences = new InputList<string>());
            set => _destThreatIntelligences = value;
        }

        [Input("layer4Configs", required: true)]
        private InputList<Inputs.FirewallPolicyRuleMatchLayer4ConfigArgs>? _layer4Configs;

        /// <summary>
        /// Pairs of IP protocols and ports that the rule should match.
        /// </summary>
        public InputList<Inputs.FirewallPolicyRuleMatchLayer4ConfigArgs> Layer4Configs
        {
            get => _layer4Configs ?? (_layer4Configs = new InputList<Inputs.FirewallPolicyRuleMatchLayer4ConfigArgs>());
            set => _layer4Configs = value;
        }

        [Input("srcAddressGroups")]
        private InputList<string>? _srcAddressGroups;

        /// <summary>
        /// Address groups which should be matched against the traffic source. Maximum number of source address groups is 10. Source address groups is only supported in Ingress rules.
        /// </summary>
        public InputList<string> SrcAddressGroups
        {
            get => _srcAddressGroups ?? (_srcAddressGroups = new InputList<string>());
            set => _srcAddressGroups = value;
        }

        [Input("srcFqdns")]
        private InputList<string>? _srcFqdns;

        /// <summary>
        /// Domain names that will be used to match against the resolved domain name of source of traffic. Can only be specified if DIRECTION is ingress.
        /// </summary>
        public InputList<string> SrcFqdns
        {
            get => _srcFqdns ?? (_srcFqdns = new InputList<string>());
            set => _srcFqdns = value;
        }

        [Input("srcIpRanges")]
        private InputList<string>? _srcIpRanges;

        /// <summary>
        /// CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 256.
        /// </summary>
        public InputList<string> SrcIpRanges
        {
            get => _srcIpRanges ?? (_srcIpRanges = new InputList<string>());
            set => _srcIpRanges = value;
        }

        [Input("srcRegionCodes")]
        private InputList<string>? _srcRegionCodes;

        /// <summary>
        /// The Unicode country codes whose IP addresses will be used to match against the source of traffic. Can only be specified if DIRECTION is ingress.
        /// </summary>
        public InputList<string> SrcRegionCodes
        {
            get => _srcRegionCodes ?? (_srcRegionCodes = new InputList<string>());
            set => _srcRegionCodes = value;
        }

        [Input("srcThreatIntelligences")]
        private InputList<string>? _srcThreatIntelligences;

        /// <summary>
        /// Name of the Google Cloud Threat Intelligence list.
        /// 
        /// The `layer4_configs` block supports:
        /// </summary>
        public InputList<string> SrcThreatIntelligences
        {
            get => _srcThreatIntelligences ?? (_srcThreatIntelligences = new InputList<string>());
            set => _srcThreatIntelligences = value;
        }

        public FirewallPolicyRuleMatchArgs()
        {
        }
        public static new FirewallPolicyRuleMatchArgs Empty => new FirewallPolicyRuleMatchArgs();
    }
}
