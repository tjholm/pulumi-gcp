// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RouterBgpGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// User-specified flag to indicate which mode to use for advertisement.
        /// Default value is `DEFAULT`.
        /// Possible values are `DEFAULT` and `CUSTOM`.
        /// </summary>
        [Input("advertiseMode")]
        public Input<string>? AdvertiseMode { get; set; }

        [Input("advertisedGroups")]
        private InputList<string>? _advertisedGroups;

        /// <summary>
        /// User-specified list of prefix groups to advertise in custom mode.
        /// This field can only be populated if advertiseMode is CUSTOM and
        /// is advertised to all peers of the router. These groups will be
        /// advertised in addition to any specified prefixes. Leave this field
        /// blank to advertise no custom groups.
        /// This enum field has the one valid value: ALL_SUBNETS
        /// </summary>
        public InputList<string> AdvertisedGroups
        {
            get => _advertisedGroups ?? (_advertisedGroups = new InputList<string>());
            set => _advertisedGroups = value;
        }

        [Input("advertisedIpRanges")]
        private InputList<Inputs.RouterBgpAdvertisedIpRangeGetArgs>? _advertisedIpRanges;

        /// <summary>
        /// User-specified list of individual IP ranges to advertise in
        /// custom mode. This field can only be populated if advertiseMode
        /// is CUSTOM and is advertised to all peers of the router. These IP
        /// ranges will be advertised in addition to any specified groups.
        /// Leave this field blank to advertise no custom IP ranges.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.RouterBgpAdvertisedIpRangeGetArgs> AdvertisedIpRanges
        {
            get => _advertisedIpRanges ?? (_advertisedIpRanges = new InputList<Inputs.RouterBgpAdvertisedIpRangeGetArgs>());
            set => _advertisedIpRanges = value;
        }

        /// <summary>
        /// Local BGP Autonomous System Number (ASN). Must be an RFC6996
        /// private ASN, either 16-bit or 32-bit. The value will be fixed for
        /// this router resource. All VPN tunnels that link to this router
        /// will have the same local ASN.
        /// </summary>
        [Input("asn", required: true)]
        public Input<int> Asn { get; set; } = null!;

        /// <summary>
        /// The interval in seconds between BGP keepalive messages that are sent to the peer.
        /// Hold time is three times the interval at which keepalive messages are sent, and the hold time is the
        /// maximum number of seconds allowed to elapse between successive keepalive messages that BGP receives from a peer.
        /// BGP will use the smaller of either the local hold time value or the peer's hold time value as the hold time for
        /// the BGP connection between the two peers. If set, this value must be between 20 and 60. The default is 20.
        /// </summary>
        [Input("keepaliveInterval")]
        public Input<int>? KeepaliveInterval { get; set; }

        public RouterBgpGetArgs()
        {
        }
    }
}
