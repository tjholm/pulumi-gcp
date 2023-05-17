// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeOnPrem.Inputs
{

    public sealed class BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The map of Kubernetes labels (key/value pairs) to be applied to
        /// each node. These will added in addition to any default label(s)
        /// that Kubernetes may apply to the node. In case of conflict in
        /// label keys, the applied set may differ depending on the Kubernetes
        /// version -- it's best to assume the behavior is undefined and
        /// conflicts should be avoided. For more information, including usage
        /// and the valid values, see:
        /// http://kubernetes.io/v1.1/docs/user-guide/labels.html
        /// An object containing a list of "key": value pairs.
        /// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("nodeConfigs")]
        private InputList<Inputs.BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigGetArgs>? _nodeConfigs;

        /// <summary>
        /// The list of machine addresses in the Bare Metal Node Pool.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigGetArgs> NodeConfigs
        {
            get => _nodeConfigs ?? (_nodeConfigs = new InputList<Inputs.BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigGetArgs>());
            set => _nodeConfigs = value;
        }

        /// <summary>
        /// Specifies the nodes operating system (default: LINUX).
        /// </summary>
        [Input("operatingSystem")]
        public Input<string>? OperatingSystem { get; set; }

        [Input("taints")]
        private InputList<Inputs.BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintGetArgs>? _taints;

        /// <summary>
        /// The initial taints assigned to nodes of this node pool.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintGetArgs> Taints
        {
            get => _taints ?? (_taints = new InputList<Inputs.BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintGetArgs>());
            set => _taints = value;
        }

        public BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigGetArgs()
        {
        }
        public static new BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigGetArgs Empty => new BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigGetArgs();
    }
}
