// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Inputs
{

    public sealed class AwsClusterControlPlaneGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication configuration for management of AWS resources.
        /// </summary>
        [Input("awsServicesAuthentication", required: true)]
        public Input<Inputs.AwsClusterControlPlaneAwsServicesAuthenticationGetArgs> AwsServicesAuthentication { get; set; } = null!;

        /// <summary>
        /// The ARN of the AWS KMS key used to encrypt cluster configuration.
        /// </summary>
        [Input("configEncryption", required: true)]
        public Input<Inputs.AwsClusterControlPlaneConfigEncryptionGetArgs> ConfigEncryption { get; set; } = null!;

        /// <summary>
        /// The ARN of the AWS KMS key used to encrypt cluster secrets.
        /// </summary>
        [Input("databaseEncryption", required: true)]
        public Input<Inputs.AwsClusterControlPlaneDatabaseEncryptionGetArgs> DatabaseEncryption { get; set; } = null!;

        /// <summary>
        /// The name of the AWS IAM instance pofile to assign to each control plane replica.
        /// </summary>
        [Input("iamInstanceProfile", required: true)]
        public Input<string> IamInstanceProfile { get; set; } = null!;

        /// <summary>
        /// Details of placement information for an instance.
        /// </summary>
        [Input("instancePlacement")]
        public Input<Inputs.AwsClusterControlPlaneInstancePlacementGetArgs>? InstancePlacement { get; set; }

        /// <summary>
        /// Optional. The AWS instance type. When unspecified, it defaults to `m5.large`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// Optional. Configuration related to the main volume provisioned for each control plane replica. The main volume is in charge of storing all of the cluster's etcd state. Volumes will be provisioned in the availability zone associated with the corresponding subnet. When unspecified, it defaults to 8 GiB with the GP2 volume type.
        /// </summary>
        [Input("mainVolume")]
        public Input<Inputs.AwsClusterControlPlaneMainVolumeGetArgs>? MainVolume { get; set; }

        /// <summary>
        /// Proxy configuration for outbound HTTP(S) traffic.
        /// </summary>
        [Input("proxyConfig")]
        public Input<Inputs.AwsClusterControlPlaneProxyConfigGetArgs>? ProxyConfig { get; set; }

        /// <summary>
        /// Optional. Configuration related to the root volume provisioned for each control plane replica. Volumes will be provisioned in the availability zone associated with the corresponding subnet. When unspecified, it defaults to 32 GiB with the GP2 volume type.
        /// </summary>
        [Input("rootVolume")]
        public Input<Inputs.AwsClusterControlPlaneRootVolumeGetArgs>? RootVolume { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// Optional. The IDs of additional security groups to add to control plane replicas. The Anthos Multi-Cloud API will automatically create and manage security groups with the minimum rules needed for a functioning cluster.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Optional. SSH configuration for how to access the underlying control plane machines.
        /// </summary>
        [Input("sshConfig")]
        public Input<Inputs.AwsClusterControlPlaneSshConfigGetArgs>? SshConfig { get; set; }

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The list of subnets where control plane replicas will run. A replica will be provisioned on each subnet and up to three values can be provided. Each subnet must be in a different AWS Availability Zone (AZ).
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Optional. A set of AWS resource tags to propagate to all underlying managed AWS resources. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The Kubernetes version to run on control plane replicas (e.g. `1.19.10-gke.1000`). You can list all supported versions on a given Google Cloud region by calling .
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public AwsClusterControlPlaneGetArgs()
        {
        }
        public static new AwsClusterControlPlaneGetArgs Empty => new AwsClusterControlPlaneGetArgs();
    }
}
