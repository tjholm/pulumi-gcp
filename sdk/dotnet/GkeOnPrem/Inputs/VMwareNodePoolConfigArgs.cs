// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeOnPrem.Inputs
{

    public sealed class VMwareNodePoolConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// VMware disk size to be used during creation.
        /// </summary>
        [Input("bootDiskSizeGb")]
        public Input<int>? BootDiskSizeGb { get; set; }

        /// <summary>
        /// The number of CPUs for each node in the node pool.
        /// </summary>
        [Input("cpus")]
        public Input<int>? Cpus { get; set; }

        /// <summary>
        /// Allow node pool traffic to be load balanced. Only works for clusters with
        /// MetalLB load balancers.
        /// </summary>
        [Input("enableLoadBalancer")]
        public Input<bool>? EnableLoadBalancer { get; set; }

        /// <summary>
        /// The OS image name in vCenter, only valid when using Windows.
        /// </summary>
        [Input("image")]
        public Input<string>? Image { get; set; }

        /// <summary>
        /// The OS image to be used for each node in a node pool.
        /// Currently `cos`, `ubuntu`, `ubuntu_containerd` and `windows` are supported.
        /// </summary>
        [Input("imageType", required: true)]
        public Input<string> ImageType { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The map of Kubernetes labels (key/value pairs) to be applied to each node.
        /// These will added in addition to any default label(s) that
        /// Kubernetes may apply to the node.
        /// In case of conflict in label keys, the applied set may differ depending on
        /// the Kubernetes version -- it's best to assume the behavior is undefined
        /// and conflicts should be avoided.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The megabytes of memory for each node in the node pool.
        /// </summary>
        [Input("memoryMb")]
        public Input<int>? MemoryMb { get; set; }

        /// <summary>
        /// The number of nodes in the node pool.
        /// </summary>
        [Input("replicas")]
        public Input<int>? Replicas { get; set; }

        [Input("taints")]
        private InputList<Inputs.VMwareNodePoolConfigTaintArgs>? _taints;

        /// <summary>
        /// The initial taints assigned to nodes of this node pool.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.VMwareNodePoolConfigTaintArgs> Taints
        {
            get => _taints ?? (_taints = new InputList<Inputs.VMwareNodePoolConfigTaintArgs>());
            set => _taints = value;
        }

        [Input("vsphereConfigs")]
        private InputList<Inputs.VMwareNodePoolConfigVsphereConfigArgs>? _vsphereConfigs;

        /// <summary>
        /// Specifies the vSphere config for node pool.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.VMwareNodePoolConfigVsphereConfigArgs> VsphereConfigs
        {
            get => _vsphereConfigs ?? (_vsphereConfigs = new InputList<Inputs.VMwareNodePoolConfigVsphereConfigArgs>());
            set => _vsphereConfigs = value;
        }

        public VMwareNodePoolConfigArgs()
        {
        }
        public static new VMwareNodePoolConfigArgs Empty => new VMwareNodePoolConfigArgs();
    }
}
