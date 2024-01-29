// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Inputs
{

    public sealed class AwsNodePoolConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Configuration related to CloudWatch metrics collection on the Auto Scaling group of the node pool. When unspecified, metrics collection is disabled.
        /// </summary>
        [Input("autoscalingMetricsCollection")]
        public Input<Inputs.AwsNodePoolConfigAutoscalingMetricsCollectionGetArgs>? AutoscalingMetricsCollection { get; set; }

        /// <summary>
        /// The ARN of the AWS KMS key used to encrypt node pool configuration.
        /// </summary>
        [Input("configEncryption", required: true)]
        public Input<Inputs.AwsNodePoolConfigConfigEncryptionGetArgs> ConfigEncryption { get; set; } = null!;

        /// <summary>
        /// The name of the AWS IAM role assigned to nodes in the pool.
        /// </summary>
        [Input("iamInstanceProfile", required: true)]
        public Input<string> IamInstanceProfile { get; set; } = null!;

        /// <summary>
        /// The OS image type to use on node pool instances.
        /// </summary>
        [Input("imageType")]
        public Input<string>? ImageType { get; set; }

        /// <summary>
        /// Details of placement information for an instance.
        /// </summary>
        [Input("instancePlacement")]
        public Input<Inputs.AwsNodePoolConfigInstancePlacementGetArgs>? InstancePlacement { get; set; }

        /// <summary>
        /// Optional. The AWS instance type. When unspecified, it defaults to `m5.large`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

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
        public Input<Inputs.AwsNodePoolConfigProxyConfigGetArgs>? ProxyConfig { get; set; }

        /// <summary>
        /// Optional. Template for the root volume provisioned for node pool nodes. Volumes will be provisioned in the availability zone assigned to the node pool subnet. When unspecified, it defaults to 32 GiB with the GP2 volume type.
        /// </summary>
        [Input("rootVolume")]
        public Input<Inputs.AwsNodePoolConfigRootVolumeGetArgs>? RootVolume { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// Optional. The IDs of additional security groups to add to nodes in this pool. The manager will automatically create security groups with minimum rules needed for a functioning cluster.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Optional. When specified, the node pool will provision Spot instances from the set of spot_config.instance_types. This field is mutually exclusive with `instance_type`
        /// </summary>
        [Input("spotConfig")]
        public Input<Inputs.AwsNodePoolConfigSpotConfigGetArgs>? SpotConfig { get; set; }

        /// <summary>
        /// Optional. The SSH configuration.
        /// </summary>
        [Input("sshConfig")]
        public Input<Inputs.AwsNodePoolConfigSshConfigGetArgs>? SshConfig { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Optional. Key/value metadata to assign to each underlying AWS resource. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("taints")]
        private InputList<Inputs.AwsNodePoolConfigTaintGetArgs>? _taints;

        /// <summary>
        /// Optional. The initial taints assigned to nodes of this node pool.
        /// </summary>
        public InputList<Inputs.AwsNodePoolConfigTaintGetArgs> Taints
        {
            get => _taints ?? (_taints = new InputList<Inputs.AwsNodePoolConfigTaintGetArgs>());
            set => _taints = value;
        }

        public AwsNodePoolConfigGetArgs()
        {
        }
        public static new AwsNodePoolConfigGetArgs Empty => new AwsNodePoolConfigGetArgs();
    }
}
