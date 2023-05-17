// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeOnPrem.Inputs
{

    public sealed class BareMetalClusterNetworkConfigIslandModeCidrArgs : global::Pulumi.ResourceArgs
    {
        [Input("podAddressCidrBlocks", required: true)]
        private InputList<string>? _podAddressCidrBlocks;

        /// <summary>
        /// All pods in the cluster are assigned an RFC1918 IPv4 address from these ranges. This field cannot be changed after creation.
        /// </summary>
        public InputList<string> PodAddressCidrBlocks
        {
            get => _podAddressCidrBlocks ?? (_podAddressCidrBlocks = new InputList<string>());
            set => _podAddressCidrBlocks = value;
        }

        [Input("serviceAddressCidrBlocks", required: true)]
        private InputList<string>? _serviceAddressCidrBlocks;

        /// <summary>
        /// All services in the cluster are assigned an RFC1918 IPv4 address from these ranges. This field cannot be changed after creation.
        /// </summary>
        public InputList<string> ServiceAddressCidrBlocks
        {
            get => _serviceAddressCidrBlocks ?? (_serviceAddressCidrBlocks = new InputList<string>());
            set => _serviceAddressCidrBlocks = value;
        }

        public BareMetalClusterNetworkConfigIslandModeCidrArgs()
        {
        }
        public static new BareMetalClusterNetworkConfigIslandModeCidrArgs Empty => new BareMetalClusterNetworkConfigIslandModeCidrArgs();
    }
}
