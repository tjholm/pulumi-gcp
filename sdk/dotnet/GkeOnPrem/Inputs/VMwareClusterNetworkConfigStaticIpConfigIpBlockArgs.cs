// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeOnPrem.Inputs
{

    public sealed class VMwareClusterNetworkConfigStaticIpConfigIpBlockArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The network gateway used by the VMware User Cluster.
        /// </summary>
        [Input("gateway", required: true)]
        public Input<string> Gateway { get; set; } = null!;

        [Input("ips", required: true)]
        private InputList<Inputs.VMwareClusterNetworkConfigStaticIpConfigIpBlockIpArgs>? _ips;

        /// <summary>
        /// The node's network configurations used by the VMware User Cluster.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.VMwareClusterNetworkConfigStaticIpConfigIpBlockIpArgs> Ips
        {
            get => _ips ?? (_ips = new InputList<Inputs.VMwareClusterNetworkConfigStaticIpConfigIpBlockIpArgs>());
            set => _ips = value;
        }

        /// <summary>
        /// The netmask used by the VMware User Cluster.
        /// </summary>
        [Input("netmask", required: true)]
        public Input<string> Netmask { get; set; } = null!;

        public VMwareClusterNetworkConfigStaticIpConfigIpBlockArgs()
        {
        }
        public static new VMwareClusterNetworkConfigStaticIpConfigIpBlockArgs Empty => new VMwareClusterNetworkConfigStaticIpConfigIpBlockArgs();
    }
}
