// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeOnPrem.Inputs
{

    public sealed class VMwareClusterVcenterGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The load balancer's IP address.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// (Output)
        /// Contains the vCenter CA certificate public key for SSL verification.
        /// </summary>
        [Input("caCertData")]
        public Input<string>? CaCertData { get; set; }

        /// <summary>
        /// (Output)
        /// The name of the vCenter cluster for the user cluster.
        /// </summary>
        [Input("cluster")]
        public Input<string>? Cluster { get; set; }

        /// <summary>
        /// (Output)
        /// The name of the vCenter datacenter for the user cluster.
        /// </summary>
        [Input("datacenter")]
        public Input<string>? Datacenter { get; set; }

        /// <summary>
        /// (Output)
        /// The Vsphere datastore used by the Control Plane Node.
        /// </summary>
        [Input("datastore")]
        public Input<string>? Datastore { get; set; }

        /// <summary>
        /// (Output)
        /// The name of the vCenter folder for the user cluster.
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

        /// <summary>
        /// (Output)
        /// The name of the vCenter resource pool for the user cluster.
        /// </summary>
        [Input("resourcePool")]
        public Input<string>? ResourcePool { get; set; }

        public VMwareClusterVcenterGetArgs()
        {
        }
        public static new VMwareClusterVcenterGetArgs Empty => new VMwareClusterVcenterGetArgs();
    }
}
