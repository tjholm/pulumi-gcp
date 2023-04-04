// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionInstanceGroupManagerStatefulExternalIpGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// , A value that prescribes what should happen to the external ip when the VM instance is deleted. The available options are `NEVER` and `ON_PERMANENT_INSTANCE_DELETION`. `NEVER` - detach the ip when the VM is deleted, but do not delete the ip. `ON_PERMANENT_INSTANCE_DELETION` will delete the external ip when the VM is permanently deleted from the instance group.
        /// </summary>
        [Input("deleteRule")]
        public Input<string>? DeleteRule { get; set; }

        /// <summary>
        /// , The network interface name of the external Ip. Possible value: `nic0`.
        /// </summary>
        [Input("interfaceName")]
        public Input<string>? InterfaceName { get; set; }

        public RegionInstanceGroupManagerStatefulExternalIpGetArgs()
        {
        }
        public static new RegionInstanceGroupManagerStatefulExternalIpGetArgs Empty => new RegionInstanceGroupManagerStatefulExternalIpGetArgs();
    }
}
