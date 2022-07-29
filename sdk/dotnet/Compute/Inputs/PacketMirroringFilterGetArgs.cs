// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class PacketMirroringFilterGetArgs : Pulumi.ResourceArgs
    {
        [Input("cidrRanges")]
        private InputList<string>? _cidrRanges;

        /// <summary>
        /// IP CIDR ranges that apply as a filter on the source (ingress) or
        /// destination (egress) IP in the IP header. Only IPv4 is supported.
        /// </summary>
        public InputList<string> CidrRanges
        {
            get => _cidrRanges ?? (_cidrRanges = new InputList<string>());
            set => _cidrRanges = value;
        }

        /// <summary>
        /// Direction of traffic to mirror.
        /// Default value is `BOTH`.
        /// Possible values are `INGRESS`, `EGRESS`, and `BOTH`.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        [Input("ipProtocols")]
        private InputList<string>? _ipProtocols;

        /// <summary>
        /// Possible IP protocols including tcp, udp, icmp and esp
        /// </summary>
        public InputList<string> IpProtocols
        {
            get => _ipProtocols ?? (_ipProtocols = new InputList<string>());
            set => _ipProtocols = value;
        }

        public PacketMirroringFilterGetArgs()
        {
        }
    }
}
