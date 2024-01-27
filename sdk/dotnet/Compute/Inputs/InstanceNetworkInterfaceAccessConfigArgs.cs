// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class InstanceNetworkInterfaceAccessConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address that will be 1:1 mapped to the instance's
        /// network ip. If not given, one will be generated.
        /// </summary>
        [Input("natIp")]
        public Input<string>? NatIp { get; set; }

        /// <summary>
        /// The service-level to be provided for IPv6 traffic when the
        /// subnet has an external subnet. Only PREMIUM or STANDARD tier is valid for IPv6.
        /// </summary>
        [Input("networkTier")]
        public Input<string>? NetworkTier { get; set; }

        /// <summary>
        /// The domain name to be used when creating DNSv6
        /// records for the external IPv6 ranges..
        /// </summary>
        [Input("publicPtrDomainName")]
        public Input<string>? PublicPtrDomainName { get; set; }

        /// <summary>
        /// Beta A full or partial URL to a security policy to add to this instance. If this field is set to an empty string it will remove the associated security policy.
        /// </summary>
        [Input("securityPolicy")]
        public Input<string>? SecurityPolicy { get; set; }

        public InstanceNetworkInterfaceAccessConfigArgs()
        {
        }
        public static new InstanceNetworkInterfaceAccessConfigArgs Empty => new InstanceNetworkInterfaceAccessConfigArgs();
    }
}
