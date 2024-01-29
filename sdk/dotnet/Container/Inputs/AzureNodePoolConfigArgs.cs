// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Inputs
{

    public sealed class AzureNodePoolConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The OS image type to use on node pool instances.
        /// </summary>
        [Input("imageType")]
        public Input<string>? ImageType { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. The initial labels assigned to nodes of this node pool. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Proxy configuration for outbound HTTP(S) traffic.
        /// </summary>
        [Input("proxyConfig")]
        public Input<Inputs.AzureNodePoolConfigProxyConfigArgs>? ProxyConfig { get; set; }

        /// <summary>
        /// Optional. Configuration related to the root volume provisioned for each node pool machine. When unspecified, it defaults to a 32-GiB Azure Disk.
        /// </summary>
        [Input("rootVolume")]
        public Input<Inputs.AzureNodePoolConfigRootVolumeArgs>? RootVolume { get; set; }

        /// <summary>
        /// SSH configuration for how to access the node pool machines.
        /// </summary>
        [Input("sshConfig", required: true)]
        public Input<Inputs.AzureNodePoolConfigSshConfigArgs> SshConfig { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Optional. A set of tags to apply to all underlying Azure resources for this node pool. This currently only includes Virtual Machine Scale Sets. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Optional. The Azure VM size name. Example: `Standard_DS2_v2`. See (/anthos/clusters/docs/azure/reference/supported-vms) for options. When unspecified, it defaults to `Standard_DS2_v2`.
        /// </summary>
        [Input("vmSize")]
        public Input<string>? VmSize { get; set; }

        public AzureNodePoolConfigArgs()
        {
        }
        public static new AzureNodePoolConfigArgs Empty => new AzureNodePoolConfigArgs();
    }
}
