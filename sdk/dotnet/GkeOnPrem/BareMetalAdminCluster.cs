// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeOnPrem
{
    /// <summary>
    /// ## Example Usage
    /// ### Gkeonprem Bare Metal Admin Cluster Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var admin_cluster_basic = new Gcp.GkeOnPrem.BareMetalAdminCluster("admin-cluster-basic", new()
    ///     {
    ///         Location = "us-west1",
    ///         BareMetalVersion = "1.13.4",
    ///         NetworkConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterNetworkConfigArgs
    ///         {
    ///             IslandModeCidr = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterNetworkConfigIslandModeCidrArgs
    ///             {
    ///                 ServiceAddressCidrBlocks = new[]
    ///                 {
    ///                     "172.26.0.0/16",
    ///                 },
    ///                 PodAddressCidrBlocks = new[]
    ///                 {
    ///                     "10.240.0.0/13",
    ///                 },
    ///             },
    ///         },
    ///         NodeConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterNodeConfigArgs
    ///         {
    ///             MaxPodsPerNode = 250,
    ///         },
    ///         ControlPlane = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterControlPlaneArgs
    ///         {
    ///             ControlPlaneNodePoolConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigArgs
    ///             {
    ///                 NodePoolConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigArgs
    ///                 {
    ///                     Labels = null,
    ///                     OperatingSystem = "LINUX",
    ///                     NodeConfigs = new[]
    ///                     {
    ///                         new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigArgs
    ///                         {
    ///                             Labels = null,
    ///                             NodeIp = "10.200.0.2",
    ///                         },
    ///                         new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigArgs
    ///                         {
    ///                             Labels = null,
    ///                             NodeIp = "10.200.0.3",
    ///                         },
    ///                         new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigArgs
    ///                         {
    ///                             Labels = null,
    ///                             NodeIp = "10.200.0.4",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         LoadBalancer = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterLoadBalancerArgs
    ///         {
    ///             PortConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterLoadBalancerPortConfigArgs
    ///             {
    ///                 ControlPlaneLoadBalancerPort = 443,
    ///             },
    ///             VipConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterLoadBalancerVipConfigArgs
    ///             {
    ///                 ControlPlaneVip = "10.200.0.5",
    ///             },
    ///         },
    ///         Storage = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterStorageArgs
    ///         {
    ///             LvpShareConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterStorageLvpShareConfigArgs
    ///             {
    ///                 LvpConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterStorageLvpShareConfigLvpConfigArgs
    ///                 {
    ///                     Path = "/mnt/localpv-share",
    ///                     StorageClass = "local-shared",
    ///                 },
    ///                 SharedPathPvCount = 5,
    ///             },
    ///             LvpNodeMountsConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterStorageLvpNodeMountsConfigArgs
    ///             {
    ///                 Path = "/mnt/localpv-disk",
    ///                 StorageClass = "local-disks",
    ///             },
    ///         },
    ///         NodeAccessConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterNodeAccessConfigArgs
    ///         {
    ///             LoginUser = "root",
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Gkeonprem Bare Metal Admin Cluster Full
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var admin_cluster_basic = new Gcp.GkeOnPrem.BareMetalAdminCluster("admin-cluster-basic", new()
    ///     {
    ///         Location = "us-west1",
    ///         Description = "test description",
    ///         BareMetalVersion = "1.13.4",
    ///         Annotations = null,
    ///         NetworkConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterNetworkConfigArgs
    ///         {
    ///             IslandModeCidr = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterNetworkConfigIslandModeCidrArgs
    ///             {
    ///                 ServiceAddressCidrBlocks = new[]
    ///                 {
    ///                     "172.26.0.0/16",
    ///                 },
    ///                 PodAddressCidrBlocks = new[]
    ///                 {
    ///                     "10.240.0.0/13",
    ///                 },
    ///             },
    ///         },
    ///         NodeConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterNodeConfigArgs
    ///         {
    ///             MaxPodsPerNode = 250,
    ///         },
    ///         ControlPlane = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterControlPlaneArgs
    ///         {
    ///             ControlPlaneNodePoolConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigArgs
    ///             {
    ///                 NodePoolConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigArgs
    ///                 {
    ///                     Labels = null,
    ///                     OperatingSystem = "LINUX",
    ///                     NodeConfigs = new[]
    ///                     {
    ///                         new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigArgs
    ///                         {
    ///                             Labels = null,
    ///                             NodeIp = "10.200.0.2",
    ///                         },
    ///                         new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigArgs
    ///                         {
    ///                             Labels = null,
    ///                             NodeIp = "10.200.0.3",
    ///                         },
    ///                         new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigArgs
    ///                         {
    ///                             Labels = null,
    ///                             NodeIp = "10.200.0.4",
    ///                         },
    ///                     },
    ///                     Taints = new[]
    ///                     {
    ///                         new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintArgs
    ///                         {
    ///                             Key = "test-key",
    ///                             Value = "test-value",
    ///                             Effect = "NO_EXECUTE",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             ApiServerArgs = new[]
    ///             {
    ///                 new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterControlPlaneApiServerArgArgs
    ///                 {
    ///                     Argument = "test argument",
    ///                     Value = "test value",
    ///                 },
    ///             },
    ///         },
    ///         LoadBalancer = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterLoadBalancerArgs
    ///         {
    ///             PortConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterLoadBalancerPortConfigArgs
    ///             {
    ///                 ControlPlaneLoadBalancerPort = 443,
    ///             },
    ///             VipConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterLoadBalancerVipConfigArgs
    ///             {
    ///                 ControlPlaneVip = "10.200.0.5",
    ///             },
    ///             ManualLbConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterLoadBalancerManualLbConfigArgs
    ///             {
    ///                 Enabled = true,
    ///             },
    ///         },
    ///         Storage = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterStorageArgs
    ///         {
    ///             LvpShareConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterStorageLvpShareConfigArgs
    ///             {
    ///                 LvpConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterStorageLvpShareConfigLvpConfigArgs
    ///                 {
    ///                     Path = "/mnt/localpv-share",
    ///                     StorageClass = "local-shared",
    ///                 },
    ///                 SharedPathPvCount = 5,
    ///             },
    ///             LvpNodeMountsConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterStorageLvpNodeMountsConfigArgs
    ///             {
    ///                 Path = "/mnt/localpv-disk",
    ///                 StorageClass = "local-disks",
    ///             },
    ///         },
    ///         NodeAccessConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterNodeAccessConfigArgs
    ///         {
    ///             LoginUser = "root",
    ///         },
    ///         SecurityConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterSecurityConfigArgs
    ///         {
    ///             Authorization = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterSecurityConfigAuthorizationArgs
    ///             {
    ///                 AdminUsers = new[]
    ///                 {
    ///                     new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterSecurityConfigAuthorizationAdminUserArgs
    ///                     {
    ///                         Username = "admin@hashicorptest.com",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         MaintenanceConfig = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterMaintenanceConfigArgs
    ///         {
    ///             MaintenanceAddressCidrBlocks = new[]
    ///             {
    ///                 "10.0.0.1/32",
    ///                 "10.0.0.2/32",
    ///             },
    ///         },
    ///         ClusterOperations = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterClusterOperationsArgs
    ///         {
    ///             EnableApplicationLogs = true,
    ///         },
    ///         Proxy = new Gcp.GkeOnPrem.Inputs.BareMetalAdminClusterProxyArgs
    ///         {
    ///             Uri = "test proxy uri",
    ///             NoProxies = new[]
    ///             {
    ///                 "127.0.0.1",
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// BareMetalAdminCluster can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:gkeonprem/bareMetalAdminCluster:BareMetalAdminCluster default projects/{{project}}/locations/{{location}}/bareMetalAdminClusters/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:gkeonprem/bareMetalAdminCluster:BareMetalAdminCluster default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:gkeonprem/bareMetalAdminCluster:BareMetalAdminCluster default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:gkeonprem/bareMetalAdminCluster:BareMetalAdminCluster")]
    public partial class BareMetalAdminCluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Annotations on the Bare Metal Admin Cluster.
        /// This field has the same restrictions as Kubernetes annotations.
        /// The total size of all keys and values combined is limited to 256k.
        /// Key can have 2 segments: prefix (optional) and name (required),
        /// separated by a slash (/).
        /// Prefix must be a DNS subdomain.
        /// Name must be 63 characters or less, begin and end with alphanumerics,
        /// with dashes (-), underscores (_), dots (.), and alphanumerics between.
        /// </summary>
        [Output("annotations")]
        public Output<ImmutableDictionary<string, string>?> Annotations { get; private set; } = null!;

        /// <summary>
        /// A human readable description of this Bare Metal Admin Cluster.
        /// </summary>
        [Output("bareMetalVersion")]
        public Output<string?> BareMetalVersion { get; private set; } = null!;

        /// <summary>
        /// Specifies the Admin Cluster's observability infrastructure.
        /// Structure is documented below.
        /// </summary>
        [Output("clusterOperations")]
        public Output<Outputs.BareMetalAdminClusterClusterOperations?> ClusterOperations { get; private set; } = null!;

        /// <summary>
        /// Specifies the control plane configuration.
        /// Structure is documented below.
        /// </summary>
        [Output("controlPlane")]
        public Output<Outputs.BareMetalAdminClusterControlPlane?> ControlPlane { get; private set; } = null!;

        /// <summary>
        /// The time the cluster was created, in RFC3339 text format.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The time the cluster was deleted, in RFC3339 text format.
        /// </summary>
        [Output("deleteTime")]
        public Output<string> DeleteTime { get; private set; } = null!;

        /// <summary>
        /// A human readable description of this Bare Metal Admin Cluster.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The IP address name of Bare Metal Admin Cluster's API server.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// This checksum is computed by the server based on the value of other
        /// fields, and may be sent on update and delete requests to ensure the
        /// client has an up-to-date value before proceeding.
        /// Allows clients to perform consistent read-modify-writes
        /// through optimistic concurrency control.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Fleet related configuration.
        /// Fleets are a Google Cloud concept for logically organizing clusters,
        /// letting you use and manage multi-cluster capabilities and apply
        /// consistent policies across your systems.
        /// See [Anthos Fleets](https://cloud.google.com/anthos/multicluster-management/fleets) for
        /// more details on Anthos multi-cluster capabilities using Fleets.
        /// Structure is documented below.
        /// </summary>
        [Output("fleets")]
        public Output<ImmutableArray<Outputs.BareMetalAdminClusterFleet>> Fleets { get; private set; } = null!;

        /// <summary>
        /// Specifies the load balancer configuration.
        /// Structure is documented below.
        /// </summary>
        [Output("loadBalancer")]
        public Output<Outputs.BareMetalAdminClusterLoadBalancer?> LoadBalancer { get; private set; } = null!;

        /// <summary>
        /// The object name of the Bare Metal Admin Cluster custom resource on the
        /// associated admin cluster. This field is used to support conflicting
        /// names when enrolling existing clusters to the API. When used as a part of
        /// cluster enrollment, this field will differ from the ID in the resource
        /// name. For new clusters, this field will match the user provided cluster ID
        /// and be visible in the last component of the resource name. It is not
        /// modifiable.
        /// All users should use this name to access their cluster using gkectl or
        /// kubectl and should expect to see the local name when viewing admin
        /// cluster controller logs.
        /// </summary>
        [Output("localName")]
        public Output<string> LocalName { get; private set; } = null!;

        /// <summary>
        /// The location of the resource.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the workload node configurations.
        /// Structure is documented below.
        /// </summary>
        [Output("maintenanceConfig")]
        public Output<Outputs.BareMetalAdminClusterMaintenanceConfig?> MaintenanceConfig { get; private set; } = null!;

        /// <summary>
        /// The bare metal admin cluster name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Network configuration.
        /// Structure is documented below.
        /// </summary>
        [Output("networkConfig")]
        public Output<Outputs.BareMetalAdminClusterNetworkConfig?> NetworkConfig { get; private set; } = null!;

        /// <summary>
        /// Specifies the node access related settings for the bare metal user cluster.
        /// Structure is documented below.
        /// </summary>
        [Output("nodeAccessConfig")]
        public Output<Outputs.BareMetalAdminClusterNodeAccessConfig?> NodeAccessConfig { get; private set; } = null!;

        /// <summary>
        /// Specifies the workload node configurations.
        /// Structure is documented below.
        /// </summary>
        [Output("nodeConfig")]
        public Output<Outputs.BareMetalAdminClusterNodeConfig?> NodeConfig { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Specifies the cluster proxy configuration.
        /// Structure is documented below.
        /// </summary>
        [Output("proxy")]
        public Output<Outputs.BareMetalAdminClusterProxy?> Proxy { get; private set; } = null!;

        /// <summary>
        /// If set, there are currently changes in flight to the Bare Metal Admin Cluster.
        /// </summary>
        [Output("reconciling")]
        public Output<bool> Reconciling { get; private set; } = null!;

        /// <summary>
        /// Specifies the security related settings for the Bare Metal User Cluster.
        /// Structure is documented below.
        /// </summary>
        [Output("securityConfig")]
        public Output<Outputs.BareMetalAdminClusterSecurityConfig?> SecurityConfig { get; private set; } = null!;

        /// <summary>
        /// (Output)
        /// The lifecycle state of the condition.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// (Output)
        /// Specifies the detailed validation check status
        /// Structure is documented below.
        /// </summary>
        [Output("statuses")]
        public Output<ImmutableArray<Outputs.BareMetalAdminClusterStatus>> Statuses { get; private set; } = null!;

        /// <summary>
        /// Specifies the cluster storage configuration.
        /// Structure is documented below.
        /// </summary>
        [Output("storage")]
        public Output<Outputs.BareMetalAdminClusterStorage?> Storage { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the Bare Metal Admin Cluster.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The time the cluster was last updated, in RFC3339 text format.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Specifies the security related settings for the Bare Metal Admin Cluster.
        /// Structure is documented below.
        /// </summary>
        [Output("validationChecks")]
        public Output<ImmutableArray<Outputs.BareMetalAdminClusterValidationCheck>> ValidationChecks { get; private set; } = null!;


        /// <summary>
        /// Create a BareMetalAdminCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BareMetalAdminCluster(string name, BareMetalAdminClusterArgs args, CustomResourceOptions? options = null)
            : base("gcp:gkeonprem/bareMetalAdminCluster:BareMetalAdminCluster", name, args ?? new BareMetalAdminClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BareMetalAdminCluster(string name, Input<string> id, BareMetalAdminClusterState? state = null, CustomResourceOptions? options = null)
            : base("gcp:gkeonprem/bareMetalAdminCluster:BareMetalAdminCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BareMetalAdminCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BareMetalAdminCluster Get(string name, Input<string> id, BareMetalAdminClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new BareMetalAdminCluster(name, id, state, options);
        }
    }

    public sealed class BareMetalAdminClusterArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// Annotations on the Bare Metal Admin Cluster.
        /// This field has the same restrictions as Kubernetes annotations.
        /// The total size of all keys and values combined is limited to 256k.
        /// Key can have 2 segments: prefix (optional) and name (required),
        /// separated by a slash (/).
        /// Prefix must be a DNS subdomain.
        /// Name must be 63 characters or less, begin and end with alphanumerics,
        /// with dashes (-), underscores (_), dots (.), and alphanumerics between.
        /// </summary>
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// A human readable description of this Bare Metal Admin Cluster.
        /// </summary>
        [Input("bareMetalVersion")]
        public Input<string>? BareMetalVersion { get; set; }

        /// <summary>
        /// Specifies the Admin Cluster's observability infrastructure.
        /// Structure is documented below.
        /// </summary>
        [Input("clusterOperations")]
        public Input<Inputs.BareMetalAdminClusterClusterOperationsArgs>? ClusterOperations { get; set; }

        /// <summary>
        /// Specifies the control plane configuration.
        /// Structure is documented below.
        /// </summary>
        [Input("controlPlane")]
        public Input<Inputs.BareMetalAdminClusterControlPlaneArgs>? ControlPlane { get; set; }

        /// <summary>
        /// A human readable description of this Bare Metal Admin Cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the load balancer configuration.
        /// Structure is documented below.
        /// </summary>
        [Input("loadBalancer")]
        public Input<Inputs.BareMetalAdminClusterLoadBalancerArgs>? LoadBalancer { get; set; }

        /// <summary>
        /// The location of the resource.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Specifies the workload node configurations.
        /// Structure is documented below.
        /// </summary>
        [Input("maintenanceConfig")]
        public Input<Inputs.BareMetalAdminClusterMaintenanceConfigArgs>? MaintenanceConfig { get; set; }

        /// <summary>
        /// The bare metal admin cluster name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network configuration.
        /// Structure is documented below.
        /// </summary>
        [Input("networkConfig")]
        public Input<Inputs.BareMetalAdminClusterNetworkConfigArgs>? NetworkConfig { get; set; }

        /// <summary>
        /// Specifies the node access related settings for the bare metal user cluster.
        /// Structure is documented below.
        /// </summary>
        [Input("nodeAccessConfig")]
        public Input<Inputs.BareMetalAdminClusterNodeAccessConfigArgs>? NodeAccessConfig { get; set; }

        /// <summary>
        /// Specifies the workload node configurations.
        /// Structure is documented below.
        /// </summary>
        [Input("nodeConfig")]
        public Input<Inputs.BareMetalAdminClusterNodeConfigArgs>? NodeConfig { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Specifies the cluster proxy configuration.
        /// Structure is documented below.
        /// </summary>
        [Input("proxy")]
        public Input<Inputs.BareMetalAdminClusterProxyArgs>? Proxy { get; set; }

        /// <summary>
        /// Specifies the security related settings for the Bare Metal User Cluster.
        /// Structure is documented below.
        /// </summary>
        [Input("securityConfig")]
        public Input<Inputs.BareMetalAdminClusterSecurityConfigArgs>? SecurityConfig { get; set; }

        /// <summary>
        /// Specifies the cluster storage configuration.
        /// Structure is documented below.
        /// </summary>
        [Input("storage")]
        public Input<Inputs.BareMetalAdminClusterStorageArgs>? Storage { get; set; }

        public BareMetalAdminClusterArgs()
        {
        }
        public static new BareMetalAdminClusterArgs Empty => new BareMetalAdminClusterArgs();
    }

    public sealed class BareMetalAdminClusterState : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// Annotations on the Bare Metal Admin Cluster.
        /// This field has the same restrictions as Kubernetes annotations.
        /// The total size of all keys and values combined is limited to 256k.
        /// Key can have 2 segments: prefix (optional) and name (required),
        /// separated by a slash (/).
        /// Prefix must be a DNS subdomain.
        /// Name must be 63 characters or less, begin and end with alphanumerics,
        /// with dashes (-), underscores (_), dots (.), and alphanumerics between.
        /// </summary>
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// A human readable description of this Bare Metal Admin Cluster.
        /// </summary>
        [Input("bareMetalVersion")]
        public Input<string>? BareMetalVersion { get; set; }

        /// <summary>
        /// Specifies the Admin Cluster's observability infrastructure.
        /// Structure is documented below.
        /// </summary>
        [Input("clusterOperations")]
        public Input<Inputs.BareMetalAdminClusterClusterOperationsGetArgs>? ClusterOperations { get; set; }

        /// <summary>
        /// Specifies the control plane configuration.
        /// Structure is documented below.
        /// </summary>
        [Input("controlPlane")]
        public Input<Inputs.BareMetalAdminClusterControlPlaneGetArgs>? ControlPlane { get; set; }

        /// <summary>
        /// The time the cluster was created, in RFC3339 text format.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The time the cluster was deleted, in RFC3339 text format.
        /// </summary>
        [Input("deleteTime")]
        public Input<string>? DeleteTime { get; set; }

        /// <summary>
        /// A human readable description of this Bare Metal Admin Cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The IP address name of Bare Metal Admin Cluster's API server.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// This checksum is computed by the server based on the value of other
        /// fields, and may be sent on update and delete requests to ensure the
        /// client has an up-to-date value before proceeding.
        /// Allows clients to perform consistent read-modify-writes
        /// through optimistic concurrency control.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("fleets")]
        private InputList<Inputs.BareMetalAdminClusterFleetGetArgs>? _fleets;

        /// <summary>
        /// Fleet related configuration.
        /// Fleets are a Google Cloud concept for logically organizing clusters,
        /// letting you use and manage multi-cluster capabilities and apply
        /// consistent policies across your systems.
        /// See [Anthos Fleets](https://cloud.google.com/anthos/multicluster-management/fleets) for
        /// more details on Anthos multi-cluster capabilities using Fleets.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.BareMetalAdminClusterFleetGetArgs> Fleets
        {
            get => _fleets ?? (_fleets = new InputList<Inputs.BareMetalAdminClusterFleetGetArgs>());
            set => _fleets = value;
        }

        /// <summary>
        /// Specifies the load balancer configuration.
        /// Structure is documented below.
        /// </summary>
        [Input("loadBalancer")]
        public Input<Inputs.BareMetalAdminClusterLoadBalancerGetArgs>? LoadBalancer { get; set; }

        /// <summary>
        /// The object name of the Bare Metal Admin Cluster custom resource on the
        /// associated admin cluster. This field is used to support conflicting
        /// names when enrolling existing clusters to the API. When used as a part of
        /// cluster enrollment, this field will differ from the ID in the resource
        /// name. For new clusters, this field will match the user provided cluster ID
        /// and be visible in the last component of the resource name. It is not
        /// modifiable.
        /// All users should use this name to access their cluster using gkectl or
        /// kubectl and should expect to see the local name when viewing admin
        /// cluster controller logs.
        /// </summary>
        [Input("localName")]
        public Input<string>? LocalName { get; set; }

        /// <summary>
        /// The location of the resource.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the workload node configurations.
        /// Structure is documented below.
        /// </summary>
        [Input("maintenanceConfig")]
        public Input<Inputs.BareMetalAdminClusterMaintenanceConfigGetArgs>? MaintenanceConfig { get; set; }

        /// <summary>
        /// The bare metal admin cluster name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network configuration.
        /// Structure is documented below.
        /// </summary>
        [Input("networkConfig")]
        public Input<Inputs.BareMetalAdminClusterNetworkConfigGetArgs>? NetworkConfig { get; set; }

        /// <summary>
        /// Specifies the node access related settings for the bare metal user cluster.
        /// Structure is documented below.
        /// </summary>
        [Input("nodeAccessConfig")]
        public Input<Inputs.BareMetalAdminClusterNodeAccessConfigGetArgs>? NodeAccessConfig { get; set; }

        /// <summary>
        /// Specifies the workload node configurations.
        /// Structure is documented below.
        /// </summary>
        [Input("nodeConfig")]
        public Input<Inputs.BareMetalAdminClusterNodeConfigGetArgs>? NodeConfig { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Specifies the cluster proxy configuration.
        /// Structure is documented below.
        /// </summary>
        [Input("proxy")]
        public Input<Inputs.BareMetalAdminClusterProxyGetArgs>? Proxy { get; set; }

        /// <summary>
        /// If set, there are currently changes in flight to the Bare Metal Admin Cluster.
        /// </summary>
        [Input("reconciling")]
        public Input<bool>? Reconciling { get; set; }

        /// <summary>
        /// Specifies the security related settings for the Bare Metal User Cluster.
        /// Structure is documented below.
        /// </summary>
        [Input("securityConfig")]
        public Input<Inputs.BareMetalAdminClusterSecurityConfigGetArgs>? SecurityConfig { get; set; }

        /// <summary>
        /// (Output)
        /// The lifecycle state of the condition.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("statuses")]
        private InputList<Inputs.BareMetalAdminClusterStatusGetArgs>? _statuses;

        /// <summary>
        /// (Output)
        /// Specifies the detailed validation check status
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.BareMetalAdminClusterStatusGetArgs> Statuses
        {
            get => _statuses ?? (_statuses = new InputList<Inputs.BareMetalAdminClusterStatusGetArgs>());
            set => _statuses = value;
        }

        /// <summary>
        /// Specifies the cluster storage configuration.
        /// Structure is documented below.
        /// </summary>
        [Input("storage")]
        public Input<Inputs.BareMetalAdminClusterStorageGetArgs>? Storage { get; set; }

        /// <summary>
        /// The unique identifier of the Bare Metal Admin Cluster.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// The time the cluster was last updated, in RFC3339 text format.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        [Input("validationChecks")]
        private InputList<Inputs.BareMetalAdminClusterValidationCheckGetArgs>? _validationChecks;

        /// <summary>
        /// Specifies the security related settings for the Bare Metal Admin Cluster.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.BareMetalAdminClusterValidationCheckGetArgs> ValidationChecks
        {
            get => _validationChecks ?? (_validationChecks = new InputList<Inputs.BareMetalAdminClusterValidationCheckGetArgs>());
            set => _validationChecks = value;
        }

        public BareMetalAdminClusterState()
        {
        }
        public static new BareMetalAdminClusterState Empty => new BareMetalAdminClusterState();
    }
}
