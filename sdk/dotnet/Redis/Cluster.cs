// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Redis
{
    /// <summary>
    /// A Google Cloud Redis Cluster instance.
    /// 
    /// To get more information about Cluster, see:
    /// 
    /// * [API documentation](https://cloud.google.com/memorystore/docs/cluster/reference/rest/v1/projects.locations.clusters)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/memorystore/docs/cluster/)
    /// 
    /// ## Example Usage
    /// ### Redis Cluster Ha
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var producerNet = new Gcp.Compute.Network("producerNet", new()
    ///     {
    ///         AutoCreateSubnetworks = false,
    ///     });
    /// 
    ///     var producerSubnet = new Gcp.Compute.Subnetwork("producerSubnet", new()
    ///     {
    ///         IpCidrRange = "10.0.0.248/29",
    ///         Region = "us-central1",
    ///         Network = producerNet.Id,
    ///     });
    /// 
    ///     var @default = new Gcp.NetworkConnectivity.ServiceConnectionPolicy("default", new()
    ///     {
    ///         Location = "us-central1",
    ///         ServiceClass = "gcp-memorystore-redis",
    ///         Description = "my basic service connection policy",
    ///         Network = producerNet.Id,
    ///         PscConfig = new Gcp.NetworkConnectivity.Inputs.ServiceConnectionPolicyPscConfigArgs
    ///         {
    ///             Subnetworks = new[]
    ///             {
    ///                 producerSubnet.Id,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var cluster_ha = new Gcp.Redis.Cluster("cluster-ha", new()
    ///     {
    ///         ShardCount = 3,
    ///         PscConfigs = new[]
    ///         {
    ///             new Gcp.Redis.Inputs.ClusterPscConfigArgs
    ///             {
    ///                 Network = producerNet.Id,
    ///             },
    ///         },
    ///         Region = "us-central1",
    ///         ReplicaCount = 1,
    ///         TransitEncryptionMode = "TRANSIT_ENCRYPTION_MODE_DISABLED",
    ///         AuthorizationMode = "AUTH_MODE_DISABLED",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             @default,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cluster can be imported using any of these accepted formats* `projects/{{project}}/locations/{{region}}/clusters/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Cluster can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:redis/cluster:Cluster default projects/{{project}}/locations/{{region}}/clusters/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:redis/cluster:Cluster default {{project}}/{{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:redis/cluster:Cluster default {{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:redis/cluster:Cluster default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:redis/cluster:Cluster")]
    public partial class Cluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. The authorization mode of the Redis cluster. If not provided, auth feature is disabled for the cluster.
        /// Default value is `AUTH_MODE_DISABLED`.
        /// Possible values are: `AUTH_MODE_UNSPECIFIED`, `AUTH_MODE_IAM_AUTH`, `AUTH_MODE_DISABLED`.
        /// </summary>
        [Output("authorizationMode")]
        public Output<string?> AuthorizationMode { get; private set; } = null!;

        /// <summary>
        /// The timestamp associated with the cluster creation request. A timestamp in
        /// RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional
        /// digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Output only. Endpoints created on each given network,
        /// for Redis clients to connect to the cluster.
        /// Currently only one endpoint is supported.
        /// Structure is documented below.
        /// </summary>
        [Output("discoveryEndpoints")]
        public Output<ImmutableArray<Outputs.ClusterDiscoveryEndpoint>> DiscoveryEndpoints { get; private set; } = null!;

        /// <summary>
        /// Unique name of the resource in this scope including project and location using the form:
        /// projects/{projectId}/locations/{locationId}/clusters/{clusterId}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Required. Each PscConfig configures the consumer network where two
        /// network addresses will be designated to the cluster for client access.
        /// Currently, only one PscConfig is supported.
        /// Structure is documented below.
        /// </summary>
        [Output("pscConfigs")]
        public Output<ImmutableArray<Outputs.ClusterPscConfig>> PscConfigs { get; private set; } = null!;

        /// <summary>
        /// Output only. PSC connections for discovery of the cluster topology and accessing the cluster.
        /// Structure is documented below.
        /// </summary>
        [Output("pscConnections")]
        public Output<ImmutableArray<Outputs.ClusterPscConnection>> PscConnections { get; private set; } = null!;

        /// <summary>
        /// The name of the region of the Redis cluster.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Optional. The number of replica nodes per shard.
        /// </summary>
        [Output("replicaCount")]
        public Output<int?> ReplicaCount { get; private set; } = null!;

        /// <summary>
        /// Required. Number of shards for the Redis cluster.
        /// </summary>
        [Output("shardCount")]
        public Output<int> ShardCount { get; private set; } = null!;

        /// <summary>
        /// Output only. Redis memory size in GB for the entire cluster.
        /// </summary>
        [Output("sizeGb")]
        public Output<int> SizeGb { get; private set; } = null!;

        /// <summary>
        /// The current state of this cluster. Can be CREATING, READY, UPDATING, DELETING and SUSPENDED
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Output only. Additional information about the current state of the cluster.
        /// Structure is documented below.
        /// </summary>
        [Output("stateInfos")]
        public Output<ImmutableArray<Outputs.ClusterStateInfo>> StateInfos { get; private set; } = null!;

        /// <summary>
        /// Optional. The in-transit encryption for the Redis cluster.
        /// If not provided, encryption is disabled for the cluster.
        /// Default value is `TRANSIT_ENCRYPTION_MODE_DISABLED`.
        /// Possible values are: `TRANSIT_ENCRYPTION_MODE_UNSPECIFIED`, `TRANSIT_ENCRYPTION_MODE_DISABLED`, `TRANSIT_ENCRYPTION_MODE_SERVER_AUTHENTICATION`.
        /// </summary>
        [Output("transitEncryptionMode")]
        public Output<string?> TransitEncryptionMode { get; private set; } = null!;

        /// <summary>
        /// System assigned, unique identifier for the cluster.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("gcp:redis/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("gcp:redis/cluster:Cluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, state, options);
        }
    }

    public sealed class ClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The authorization mode of the Redis cluster. If not provided, auth feature is disabled for the cluster.
        /// Default value is `AUTH_MODE_DISABLED`.
        /// Possible values are: `AUTH_MODE_UNSPECIFIED`, `AUTH_MODE_IAM_AUTH`, `AUTH_MODE_DISABLED`.
        /// </summary>
        [Input("authorizationMode")]
        public Input<string>? AuthorizationMode { get; set; }

        /// <summary>
        /// Unique name of the resource in this scope including project and location using the form:
        /// projects/{projectId}/locations/{locationId}/clusters/{clusterId}
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("pscConfigs", required: true)]
        private InputList<Inputs.ClusterPscConfigArgs>? _pscConfigs;

        /// <summary>
        /// Required. Each PscConfig configures the consumer network where two
        /// network addresses will be designated to the cluster for client access.
        /// Currently, only one PscConfig is supported.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ClusterPscConfigArgs> PscConfigs
        {
            get => _pscConfigs ?? (_pscConfigs = new InputList<Inputs.ClusterPscConfigArgs>());
            set => _pscConfigs = value;
        }

        /// <summary>
        /// The name of the region of the Redis cluster.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Optional. The number of replica nodes per shard.
        /// </summary>
        [Input("replicaCount")]
        public Input<int>? ReplicaCount { get; set; }

        /// <summary>
        /// Required. Number of shards for the Redis cluster.
        /// </summary>
        [Input("shardCount", required: true)]
        public Input<int> ShardCount { get; set; } = null!;

        /// <summary>
        /// Optional. The in-transit encryption for the Redis cluster.
        /// If not provided, encryption is disabled for the cluster.
        /// Default value is `TRANSIT_ENCRYPTION_MODE_DISABLED`.
        /// Possible values are: `TRANSIT_ENCRYPTION_MODE_UNSPECIFIED`, `TRANSIT_ENCRYPTION_MODE_DISABLED`, `TRANSIT_ENCRYPTION_MODE_SERVER_AUTHENTICATION`.
        /// </summary>
        [Input("transitEncryptionMode")]
        public Input<string>? TransitEncryptionMode { get; set; }

        public ClusterArgs()
        {
        }
        public static new ClusterArgs Empty => new ClusterArgs();
    }

    public sealed class ClusterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The authorization mode of the Redis cluster. If not provided, auth feature is disabled for the cluster.
        /// Default value is `AUTH_MODE_DISABLED`.
        /// Possible values are: `AUTH_MODE_UNSPECIFIED`, `AUTH_MODE_IAM_AUTH`, `AUTH_MODE_DISABLED`.
        /// </summary>
        [Input("authorizationMode")]
        public Input<string>? AuthorizationMode { get; set; }

        /// <summary>
        /// The timestamp associated with the cluster creation request. A timestamp in
        /// RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional
        /// digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        [Input("discoveryEndpoints")]
        private InputList<Inputs.ClusterDiscoveryEndpointGetArgs>? _discoveryEndpoints;

        /// <summary>
        /// Output only. Endpoints created on each given network,
        /// for Redis clients to connect to the cluster.
        /// Currently only one endpoint is supported.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ClusterDiscoveryEndpointGetArgs> DiscoveryEndpoints
        {
            get => _discoveryEndpoints ?? (_discoveryEndpoints = new InputList<Inputs.ClusterDiscoveryEndpointGetArgs>());
            set => _discoveryEndpoints = value;
        }

        /// <summary>
        /// Unique name of the resource in this scope including project and location using the form:
        /// projects/{projectId}/locations/{locationId}/clusters/{clusterId}
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("pscConfigs")]
        private InputList<Inputs.ClusterPscConfigGetArgs>? _pscConfigs;

        /// <summary>
        /// Required. Each PscConfig configures the consumer network where two
        /// network addresses will be designated to the cluster for client access.
        /// Currently, only one PscConfig is supported.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ClusterPscConfigGetArgs> PscConfigs
        {
            get => _pscConfigs ?? (_pscConfigs = new InputList<Inputs.ClusterPscConfigGetArgs>());
            set => _pscConfigs = value;
        }

        [Input("pscConnections")]
        private InputList<Inputs.ClusterPscConnectionGetArgs>? _pscConnections;

        /// <summary>
        /// Output only. PSC connections for discovery of the cluster topology and accessing the cluster.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ClusterPscConnectionGetArgs> PscConnections
        {
            get => _pscConnections ?? (_pscConnections = new InputList<Inputs.ClusterPscConnectionGetArgs>());
            set => _pscConnections = value;
        }

        /// <summary>
        /// The name of the region of the Redis cluster.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Optional. The number of replica nodes per shard.
        /// </summary>
        [Input("replicaCount")]
        public Input<int>? ReplicaCount { get; set; }

        /// <summary>
        /// Required. Number of shards for the Redis cluster.
        /// </summary>
        [Input("shardCount")]
        public Input<int>? ShardCount { get; set; }

        /// <summary>
        /// Output only. Redis memory size in GB for the entire cluster.
        /// </summary>
        [Input("sizeGb")]
        public Input<int>? SizeGb { get; set; }

        /// <summary>
        /// The current state of this cluster. Can be CREATING, READY, UPDATING, DELETING and SUSPENDED
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("stateInfos")]
        private InputList<Inputs.ClusterStateInfoGetArgs>? _stateInfos;

        /// <summary>
        /// Output only. Additional information about the current state of the cluster.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ClusterStateInfoGetArgs> StateInfos
        {
            get => _stateInfos ?? (_stateInfos = new InputList<Inputs.ClusterStateInfoGetArgs>());
            set => _stateInfos = value;
        }

        /// <summary>
        /// Optional. The in-transit encryption for the Redis cluster.
        /// If not provided, encryption is disabled for the cluster.
        /// Default value is `TRANSIT_ENCRYPTION_MODE_DISABLED`.
        /// Possible values are: `TRANSIT_ENCRYPTION_MODE_UNSPECIFIED`, `TRANSIT_ENCRYPTION_MODE_DISABLED`, `TRANSIT_ENCRYPTION_MODE_SERVER_AUTHENTICATION`.
        /// </summary>
        [Input("transitEncryptionMode")]
        public Input<string>? TransitEncryptionMode { get; set; }

        /// <summary>
        /// System assigned, unique identifier for the cluster.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        public ClusterState()
        {
        }
        public static new ClusterState Empty => new ClusterState();
    }
}
