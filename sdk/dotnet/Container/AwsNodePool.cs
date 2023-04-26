// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container
{
    /// <summary>
    /// An Anthos node pool running on AWS.
    /// 
    /// For more information, see:
    /// * [Multicloud overview](https://cloud.google.com/anthos/clusters/docs/multi-cloud)
    /// ## Example Usage
    /// ### Basic_aws_cluster
    /// A basic example of a containeraws node pool
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var versions = Gcp.Container.GetAwsVersions.Invoke(new()
    ///     {
    ///         Project = "my-project-name",
    ///         Location = "us-west1",
    ///     });
    /// 
    ///     var primaryAwsCluster = new Gcp.Container.AwsCluster("primaryAwsCluster", new()
    ///     {
    ///         Authorization = new Gcp.Container.Inputs.AwsClusterAuthorizationArgs
    ///         {
    ///             AdminUsers = new[]
    ///             {
    ///                 new Gcp.Container.Inputs.AwsClusterAuthorizationAdminUserArgs
    ///                 {
    ///                     Username = "my@service-account.com",
    ///                 },
    ///             },
    ///         },
    ///         AwsRegion = "my-aws-region",
    ///         ControlPlane = new Gcp.Container.Inputs.AwsClusterControlPlaneArgs
    ///         {
    ///             AwsServicesAuthentication = new Gcp.Container.Inputs.AwsClusterControlPlaneAwsServicesAuthenticationArgs
    ///             {
    ///                 RoleArn = "arn:aws:iam::012345678910:role/my--1p-dev-oneplatform",
    ///                 RoleSessionName = "my--1p-dev-session",
    ///             },
    ///             ConfigEncryption = new Gcp.Container.Inputs.AwsClusterControlPlaneConfigEncryptionArgs
    ///             {
    ///                 KmsKeyArn = "arn:aws:kms:my-aws-region:012345678910:key/12345678-1234-1234-1234-123456789111",
    ///             },
    ///             DatabaseEncryption = new Gcp.Container.Inputs.AwsClusterControlPlaneDatabaseEncryptionArgs
    ///             {
    ///                 KmsKeyArn = "arn:aws:kms:my-aws-region:012345678910:key/12345678-1234-1234-1234-123456789111",
    ///             },
    ///             IamInstanceProfile = "my--1p-dev-controlplane",
    ///             SubnetIds = new[]
    ///             {
    ///                 "subnet-00000000000000000",
    ///             },
    ///             Version = versions.Apply(getAwsVersionsResult =&gt; getAwsVersionsResult.ValidVersions[0]),
    ///             InstanceType = "t3.medium",
    ///             MainVolume = new Gcp.Container.Inputs.AwsClusterControlPlaneMainVolumeArgs
    ///             {
    ///                 Iops = 3000,
    ///                 KmsKeyArn = "arn:aws:kms:my-aws-region:012345678910:key/12345678-1234-1234-1234-123456789111",
    ///                 SizeGib = 10,
    ///                 VolumeType = "GP3",
    ///             },
    ///             ProxyConfig = new Gcp.Container.Inputs.AwsClusterControlPlaneProxyConfigArgs
    ///             {
    ///                 SecretArn = "arn:aws:secretsmanager:us-west-2:126285863215:secret:proxy_config20210824150329476300000001-ABCDEF",
    ///                 SecretVersion = "12345678-ABCD-EFGH-IJKL-987654321098",
    ///             },
    ///             RootVolume = new Gcp.Container.Inputs.AwsClusterControlPlaneRootVolumeArgs
    ///             {
    ///                 Iops = 3000,
    ///                 KmsKeyArn = "arn:aws:kms:my-aws-region:012345678910:key/12345678-1234-1234-1234-123456789111",
    ///                 SizeGib = 10,
    ///                 VolumeType = "GP3",
    ///             },
    ///             SecurityGroupIds = new[]
    ///             {
    ///                 "sg-00000000000000000",
    ///             },
    ///             SshConfig = new Gcp.Container.Inputs.AwsClusterControlPlaneSshConfigArgs
    ///             {
    ///                 Ec2KeyPair = "my--1p-dev-ssh",
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "owner", "my@service-account.com" },
    ///             },
    ///         },
    ///         Fleet = new Gcp.Container.Inputs.AwsClusterFleetArgs
    ///         {
    ///             Project = "my-project-number",
    ///         },
    ///         Location = "us-west1",
    ///         Networking = new Gcp.Container.Inputs.AwsClusterNetworkingArgs
    ///         {
    ///             PodAddressCidrBlocks = new[]
    ///             {
    ///                 "10.2.0.0/16",
    ///             },
    ///             ServiceAddressCidrBlocks = new[]
    ///             {
    ///                 "10.1.0.0/16",
    ///             },
    ///             VpcId = "vpc-00000000000000000",
    ///         },
    ///         Annotations = 
    ///         {
    ///             { "label-one", "value-one" },
    ///         },
    ///         Description = "A sample aws cluster",
    ///         Project = "my-project-name",
    ///     });
    /// 
    ///     var primaryAwsNodePool = new Gcp.Container.AwsNodePool("primaryAwsNodePool", new()
    ///     {
    ///         Autoscaling = new Gcp.Container.Inputs.AwsNodePoolAutoscalingArgs
    ///         {
    ///             MaxNodeCount = 5,
    ///             MinNodeCount = 1,
    ///         },
    ///         Cluster = primaryAwsCluster.Name,
    ///         Config = new Gcp.Container.Inputs.AwsNodePoolConfigArgs
    ///         {
    ///             ConfigEncryption = new Gcp.Container.Inputs.AwsNodePoolConfigConfigEncryptionArgs
    ///             {
    ///                 KmsKeyArn = "arn:aws:kms:my-aws-region:012345678910:key/12345678-1234-1234-1234-123456789111",
    ///             },
    ///             IamInstanceProfile = "my--1p-dev-nodepool",
    ///             InstanceType = "t3.medium",
    ///             Labels = 
    ///             {
    ///                 { "label-one", "value-one" },
    ///             },
    ///             RootVolume = new Gcp.Container.Inputs.AwsNodePoolConfigRootVolumeArgs
    ///             {
    ///                 Iops = 3000,
    ///                 KmsKeyArn = "arn:aws:kms:my-aws-region:012345678910:key/12345678-1234-1234-1234-123456789111",
    ///                 SizeGib = 10,
    ///                 VolumeType = "GP3",
    ///             },
    ///             SecurityGroupIds = new[]
    ///             {
    ///                 "sg-00000000000000000",
    ///             },
    ///             ProxyConfig = new Gcp.Container.Inputs.AwsNodePoolConfigProxyConfigArgs
    ///             {
    ///                 SecretArn = "arn:aws:secretsmanager:us-west-2:126285863215:secret:proxy_config20210824150329476300000001-ABCDEF",
    ///                 SecretVersion = "12345678-ABCD-EFGH-IJKL-987654321098",
    ///             },
    ///             SshConfig = new Gcp.Container.Inputs.AwsNodePoolConfigSshConfigArgs
    ///             {
    ///                 Ec2KeyPair = "my--1p-dev-ssh",
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "tag-one", "value-one" },
    ///             },
    ///             Taints = new[]
    ///             {
    ///                 new Gcp.Container.Inputs.AwsNodePoolConfigTaintArgs
    ///                 {
    ///                     Effect = "PREFER_NO_SCHEDULE",
    ///                     Key = "taint-key",
    ///                     Value = "taint-value",
    ///                 },
    ///             },
    ///         },
    ///         Location = "us-west1",
    ///         MaxPodsConstraint = new Gcp.Container.Inputs.AwsNodePoolMaxPodsConstraintArgs
    ///         {
    ///             MaxPodsPerNode = 110,
    ///         },
    ///         SubnetId = "subnet-00000000000000000",
    ///         Version = versions.Apply(getAwsVersionsResult =&gt; getAwsVersionsResult.ValidVersions[0]),
    ///         Annotations = 
    ///         {
    ///             { "label-one", "value-one" },
    ///         },
    ///         Project = "my-project-name",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// NodePool can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:container/awsNodePool:AwsNodePool default projects/{{project}}/locations/{{location}}/awsClusters/{{cluster}}/awsNodePools/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:container/awsNodePool:AwsNodePool default {{project}}/{{location}}/{{cluster}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:container/awsNodePool:AwsNodePool default {{location}}/{{cluster}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:container/awsNodePool:AwsNodePool")]
    public partial class AwsNodePool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. Annotations on the node pool. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
        /// </summary>
        [Output("annotations")]
        public Output<ImmutableDictionary<string, string>?> Annotations { get; private set; } = null!;

        /// <summary>
        /// Autoscaler configuration for this node pool.
        /// </summary>
        [Output("autoscaling")]
        public Output<Outputs.AwsNodePoolAutoscaling> Autoscaling { get; private set; } = null!;

        /// <summary>
        /// The awsCluster for the resource
        /// </summary>
        [Output("cluster")]
        public Output<string> Cluster { get; private set; } = null!;

        /// <summary>
        /// The configuration of the node pool.
        /// </summary>
        [Output("config")]
        public Output<Outputs.AwsNodePoolConfig> Config { get; private set; } = null!;

        /// <summary>
        /// Output only. The time at which this node pool was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The location for the resource
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
        /// </summary>
        [Output("maxPodsConstraint")]
        public Output<Outputs.AwsNodePoolMaxPodsConstraint> MaxPodsConstraint { get; private set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project for the resource
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Output only. If set, there are currently changes in flight to the node pool.
        /// </summary>
        [Output("reconciling")]
        public Output<bool> Reconciling { get; private set; } = null!;

        /// <summary>
        /// Output only. The lifecycle state of the node pool. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR, DEGRADED
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The subnet where the node pool node run.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Output only. A globally unique identifier for the node pool.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// Output only. The time at which this node pool was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// The Kubernetes version to run on this node pool (e.g. `1.19.10-gke.1000`). You can list all supported versions on a given Google Cloud region by calling GetAwsServerConfig.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a AwsNodePool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AwsNodePool(string name, AwsNodePoolArgs args, CustomResourceOptions? options = null)
            : base("gcp:container/awsNodePool:AwsNodePool", name, args ?? new AwsNodePoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AwsNodePool(string name, Input<string> id, AwsNodePoolState? state = null, CustomResourceOptions? options = null)
            : base("gcp:container/awsNodePool:AwsNodePool", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AwsNodePool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AwsNodePool Get(string name, Input<string> id, AwsNodePoolState? state = null, CustomResourceOptions? options = null)
        {
            return new AwsNodePool(name, id, state, options);
        }
    }

    public sealed class AwsNodePoolArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// Optional. Annotations on the node pool. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
        /// </summary>
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// Autoscaler configuration for this node pool.
        /// </summary>
        [Input("autoscaling", required: true)]
        public Input<Inputs.AwsNodePoolAutoscalingArgs> Autoscaling { get; set; } = null!;

        /// <summary>
        /// The awsCluster for the resource
        /// </summary>
        [Input("cluster", required: true)]
        public Input<string> Cluster { get; set; } = null!;

        /// <summary>
        /// The configuration of the node pool.
        /// </summary>
        [Input("config", required: true)]
        public Input<Inputs.AwsNodePoolConfigArgs> Config { get; set; } = null!;

        /// <summary>
        /// The location for the resource
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
        /// </summary>
        [Input("maxPodsConstraint", required: true)]
        public Input<Inputs.AwsNodePoolMaxPodsConstraintArgs> MaxPodsConstraint { get; set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project for the resource
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The subnet where the node pool node run.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        /// <summary>
        /// The Kubernetes version to run on this node pool (e.g. `1.19.10-gke.1000`). You can list all supported versions on a given Google Cloud region by calling GetAwsServerConfig.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public AwsNodePoolArgs()
        {
        }
        public static new AwsNodePoolArgs Empty => new AwsNodePoolArgs();
    }

    public sealed class AwsNodePoolState : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// Optional. Annotations on the node pool. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
        /// </summary>
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// Autoscaler configuration for this node pool.
        /// </summary>
        [Input("autoscaling")]
        public Input<Inputs.AwsNodePoolAutoscalingGetArgs>? Autoscaling { get; set; }

        /// <summary>
        /// The awsCluster for the resource
        /// </summary>
        [Input("cluster")]
        public Input<string>? Cluster { get; set; }

        /// <summary>
        /// The configuration of the node pool.
        /// </summary>
        [Input("config")]
        public Input<Inputs.AwsNodePoolConfigGetArgs>? Config { get; set; }

        /// <summary>
        /// Output only. The time at which this node pool was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The location for the resource
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
        /// </summary>
        [Input("maxPodsConstraint")]
        public Input<Inputs.AwsNodePoolMaxPodsConstraintGetArgs>? MaxPodsConstraint { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project for the resource
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Output only. If set, there are currently changes in flight to the node pool.
        /// </summary>
        [Input("reconciling")]
        public Input<bool>? Reconciling { get; set; }

        /// <summary>
        /// Output only. The lifecycle state of the node pool. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR, DEGRADED
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The subnet where the node pool node run.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// Output only. A globally unique identifier for the node pool.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// Output only. The time at which this node pool was last updated.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// The Kubernetes version to run on this node pool (e.g. `1.19.10-gke.1000`). You can list all supported versions on a given Google Cloud region by calling GetAwsServerConfig.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public AwsNodePoolState()
        {
        }
        public static new AwsNodePoolState Empty => new AwsNodePoolState();
    }
}
