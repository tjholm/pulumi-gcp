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
    /// ### Gkeonprem Vmware Node Pool Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var default_basic = new Gcp.GkeOnPrem.VMwareCluster("default-basic", new()
    ///     {
    ///         Location = "us-west1",
    ///         AdminClusterMembership = "projects/870316890899/locations/global/memberships/gkeonprem-terraform-test",
    ///         Description = "test cluster",
    ///         OnPremVersion = "1.13.1-gke.35",
    ///         NetworkConfig = new Gcp.GkeOnPrem.Inputs.VMwareClusterNetworkConfigArgs
    ///         {
    ///             ServiceAddressCidrBlocks = new[]
    ///             {
    ///                 "10.96.0.0/12",
    ///             },
    ///             PodAddressCidrBlocks = new[]
    ///             {
    ///                 "192.168.0.0/16",
    ///             },
    ///             DhcpIpConfig = new Gcp.GkeOnPrem.Inputs.VMwareClusterNetworkConfigDhcpIpConfigArgs
    ///             {
    ///                 Enabled = true,
    ///             },
    ///         },
    ///         ControlPlaneNode = new Gcp.GkeOnPrem.Inputs.VMwareClusterControlPlaneNodeArgs
    ///         {
    ///             Cpus = 4,
    ///             Memory = 8192,
    ///             Replicas = 1,
    ///         },
    ///         LoadBalancer = new Gcp.GkeOnPrem.Inputs.VMwareClusterLoadBalancerArgs
    ///         {
    ///             VipConfig = new Gcp.GkeOnPrem.Inputs.VMwareClusterLoadBalancerVipConfigArgs
    ///             {
    ///                 ControlPlaneVip = "10.251.133.5",
    ///                 IngressVip = "10.251.135.19",
    ///             },
    ///             MetalLbConfig = new Gcp.GkeOnPrem.Inputs.VMwareClusterLoadBalancerMetalLbConfigArgs
    ///             {
    ///                 AddressPools = new[]
    ///                 {
    ///                     new Gcp.GkeOnPrem.Inputs.VMwareClusterLoadBalancerMetalLbConfigAddressPoolArgs
    ///                     {
    ///                         Pool = "ingress-ip",
    ///                         ManualAssign = true,
    ///                         Addresses = new[]
    ///                         {
    ///                             "10.251.135.19",
    ///                         },
    ///                     },
    ///                     new Gcp.GkeOnPrem.Inputs.VMwareClusterLoadBalancerMetalLbConfigAddressPoolArgs
    ///                     {
    ///                         Pool = "lb-test-ip",
    ///                         ManualAssign = true,
    ///                         Addresses = new[]
    ///                         {
    ///                             "10.251.135.19",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var nodepool_basic = new Gcp.GkeOnPrem.VMwareNodePool("nodepool-basic", new()
    ///     {
    ///         Location = "us-west1",
    ///         VmwareCluster = default_basic.Name,
    ///         Config = new Gcp.GkeOnPrem.Inputs.VMwareNodePoolConfigArgs
    ///         {
    ///             Replicas = 3,
    ///             ImageType = "ubuntu_containerd",
    ///             EnableLoadBalancer = true,
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Gkeonprem Vmware Node Pool Full
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var default_full = new Gcp.GkeOnPrem.VMwareCluster("default-full", new()
    ///     {
    ///         Location = "us-west1",
    ///         AdminClusterMembership = "projects/870316890899/locations/global/memberships/gkeonprem-terraform-test",
    ///         Description = "test cluster",
    ///         OnPremVersion = "1.13.1-gke.35",
    ///         NetworkConfig = new Gcp.GkeOnPrem.Inputs.VMwareClusterNetworkConfigArgs
    ///         {
    ///             ServiceAddressCidrBlocks = new[]
    ///             {
    ///                 "10.96.0.0/12",
    ///             },
    ///             PodAddressCidrBlocks = new[]
    ///             {
    ///                 "192.168.0.0/16",
    ///             },
    ///             DhcpIpConfig = new Gcp.GkeOnPrem.Inputs.VMwareClusterNetworkConfigDhcpIpConfigArgs
    ///             {
    ///                 Enabled = true,
    ///             },
    ///         },
    ///         ControlPlaneNode = new Gcp.GkeOnPrem.Inputs.VMwareClusterControlPlaneNodeArgs
    ///         {
    ///             Cpus = 4,
    ///             Memory = 8192,
    ///             Replicas = 1,
    ///         },
    ///         LoadBalancer = new Gcp.GkeOnPrem.Inputs.VMwareClusterLoadBalancerArgs
    ///         {
    ///             VipConfig = new Gcp.GkeOnPrem.Inputs.VMwareClusterLoadBalancerVipConfigArgs
    ///             {
    ///                 ControlPlaneVip = "10.251.133.5",
    ///                 IngressVip = "10.251.135.19",
    ///             },
    ///             MetalLbConfig = new Gcp.GkeOnPrem.Inputs.VMwareClusterLoadBalancerMetalLbConfigArgs
    ///             {
    ///                 AddressPools = new[]
    ///                 {
    ///                     new Gcp.GkeOnPrem.Inputs.VMwareClusterLoadBalancerMetalLbConfigAddressPoolArgs
    ///                     {
    ///                         Pool = "ingress-ip",
    ///                         ManualAssign = true,
    ///                         Addresses = new[]
    ///                         {
    ///                             "10.251.135.19",
    ///                         },
    ///                     },
    ///                     new Gcp.GkeOnPrem.Inputs.VMwareClusterLoadBalancerMetalLbConfigAddressPoolArgs
    ///                     {
    ///                         Pool = "lb-test-ip",
    ///                         ManualAssign = true,
    ///                         Addresses = new[]
    ///                         {
    ///                             "10.251.135.19",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var nodepool_full = new Gcp.GkeOnPrem.VMwareNodePool("nodepool-full", new()
    ///     {
    ///         Location = "us-west1",
    ///         VmwareCluster = default_full.Name,
    ///         Annotations = null,
    ///         Config = new Gcp.GkeOnPrem.Inputs.VMwareNodePoolConfigArgs
    ///         {
    ///             Cpus = 4,
    ///             MemoryMb = 8196,
    ///             Replicas = 3,
    ///             ImageType = "ubuntu_containerd",
    ///             Image = "image",
    ///             BootDiskSizeGb = 10,
    ///             Taints = new[]
    ///             {
    ///                 new Gcp.GkeOnPrem.Inputs.VMwareNodePoolConfigTaintArgs
    ///                 {
    ///                     Key = "key",
    ///                     Value = "value",
    ///                 },
    ///                 new Gcp.GkeOnPrem.Inputs.VMwareNodePoolConfigTaintArgs
    ///                 {
    ///                     Key = "key",
    ///                     Value = "value",
    ///                     Effect = "NO_SCHEDULE",
    ///                 },
    ///             },
    ///             Labels = null,
    ///             EnableLoadBalancer = true,
    ///         },
    ///         NodePoolAutoscaling = new Gcp.GkeOnPrem.Inputs.VMwareNodePoolNodePoolAutoscalingArgs
    ///         {
    ///             MinReplicas = 1,
    ///             MaxReplicas = 5,
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
    /// VmwareNodePool can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:gkeonprem/vMwareNodePool:VMwareNodePool default projects/{{project}}/locations/{{location}}/vmwareClusters/{{vmware_cluster}}/vmwareNodePools/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:gkeonprem/vMwareNodePool:VMwareNodePool default {{project}}/{{location}}/{{vmware_cluster}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:gkeonprem/vMwareNodePool:VMwareNodePool default {{location}}/{{vmware_cluster}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:gkeonprem/vMwareNodePool:VMwareNodePool")]
    public partial class VMwareNodePool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Annotations on the node Pool.
        /// This field has the same restrictions as Kubernetes annotations.
        /// The total size of all keys and values combined is limited to 256k.
        /// Key can have 2 segments: prefix (optional) and name (required),
        /// separated by a slash (/).
        /// Prefix must be a DNS subdomain.
        /// Name must be 63 characters or less, begin and end with alphanumerics,
        /// with dashes (-), underscores (_), dots (.), and alphanumerics between.
        /// </summary>
        [Output("annotations")]
        public Output<ImmutableDictionary<string, string>> Annotations { get; private set; } = null!;

        /// <summary>
        /// The node configuration of the node pool.
        /// Structure is documented below.
        /// </summary>
        [Output("config")]
        public Output<Outputs.VMwareNodePoolConfig> Config { get; private set; } = null!;

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
        /// The display name for the node pool.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

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
        /// The location of the resource.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The vmware node pool name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Node Pool autoscaling config for the node pool.
        /// Structure is documented below.
        /// </summary>
        [Output("nodePoolAutoscaling")]
        public Output<Outputs.VMwareNodePoolNodePoolAutoscaling?> NodePoolAutoscaling { get; private set; } = null!;

        /// <summary>
        /// Anthos version for the node pool. Defaults to the user cluster version.
        /// </summary>
        [Output("onPremVersion")]
        public Output<string> OnPremVersion { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// If set, there are currently changes in flight to the node pool.
        /// </summary>
        [Output("reconciling")]
        public Output<bool> Reconciling { get; private set; } = null!;

        /// <summary>
        /// (Output)
        /// The lifecycle state of the condition.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// ResourceStatus representing detailed cluster state.
        /// Structure is documented below.
        /// </summary>
        [Output("statuses")]
        public Output<ImmutableArray<Outputs.VMwareNodePoolStatus>> Statuses { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the node pool.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The time the cluster was last updated, in RFC3339 text format.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// The cluster this node pool belongs to.
        /// </summary>
        [Output("vmwareCluster")]
        public Output<string> VmwareCluster { get; private set; } = null!;


        /// <summary>
        /// Create a VMwareNodePool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VMwareNodePool(string name, VMwareNodePoolArgs args, CustomResourceOptions? options = null)
            : base("gcp:gkeonprem/vMwareNodePool:VMwareNodePool", name, args ?? new VMwareNodePoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VMwareNodePool(string name, Input<string> id, VMwareNodePoolState? state = null, CustomResourceOptions? options = null)
            : base("gcp:gkeonprem/vMwareNodePool:VMwareNodePool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VMwareNodePool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VMwareNodePool Get(string name, Input<string> id, VMwareNodePoolState? state = null, CustomResourceOptions? options = null)
        {
            return new VMwareNodePool(name, id, state, options);
        }
    }

    public sealed class VMwareNodePoolArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// Annotations on the node Pool.
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
        /// The node configuration of the node pool.
        /// Structure is documented below.
        /// </summary>
        [Input("config", required: true)]
        public Input<Inputs.VMwareNodePoolConfigArgs> Config { get; set; } = null!;

        /// <summary>
        /// The display name for the node pool.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The vmware node pool name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Node Pool autoscaling config for the node pool.
        /// Structure is documented below.
        /// </summary>
        [Input("nodePoolAutoscaling")]
        public Input<Inputs.VMwareNodePoolNodePoolAutoscalingArgs>? NodePoolAutoscaling { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The cluster this node pool belongs to.
        /// </summary>
        [Input("vmwareCluster", required: true)]
        public Input<string> VmwareCluster { get; set; } = null!;

        public VMwareNodePoolArgs()
        {
        }
        public static new VMwareNodePoolArgs Empty => new VMwareNodePoolArgs();
    }

    public sealed class VMwareNodePoolState : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// Annotations on the node Pool.
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
        /// The node configuration of the node pool.
        /// Structure is documented below.
        /// </summary>
        [Input("config")]
        public Input<Inputs.VMwareNodePoolConfigGetArgs>? Config { get; set; }

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
        /// The display name for the node pool.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// This checksum is computed by the server based on the value of other
        /// fields, and may be sent on update and delete requests to ensure the
        /// client has an up-to-date value before proceeding.
        /// Allows clients to perform consistent read-modify-writes
        /// through optimistic concurrency control.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The vmware node pool name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Node Pool autoscaling config for the node pool.
        /// Structure is documented below.
        /// </summary>
        [Input("nodePoolAutoscaling")]
        public Input<Inputs.VMwareNodePoolNodePoolAutoscalingGetArgs>? NodePoolAutoscaling { get; set; }

        /// <summary>
        /// Anthos version for the node pool. Defaults to the user cluster version.
        /// </summary>
        [Input("onPremVersion")]
        public Input<string>? OnPremVersion { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// If set, there are currently changes in flight to the node pool.
        /// </summary>
        [Input("reconciling")]
        public Input<bool>? Reconciling { get; set; }

        /// <summary>
        /// (Output)
        /// The lifecycle state of the condition.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("statuses")]
        private InputList<Inputs.VMwareNodePoolStatusGetArgs>? _statuses;

        /// <summary>
        /// ResourceStatus representing detailed cluster state.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.VMwareNodePoolStatusGetArgs> Statuses
        {
            get => _statuses ?? (_statuses = new InputList<Inputs.VMwareNodePoolStatusGetArgs>());
            set => _statuses = value;
        }

        /// <summary>
        /// The unique identifier of the node pool.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// The time the cluster was last updated, in RFC3339 text format.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// The cluster this node pool belongs to.
        /// </summary>
        [Input("vmwareCluster")]
        public Input<string>? VmwareCluster { get; set; }

        public VMwareNodePoolState()
        {
        }
        public static new VMwareNodePoolState Empty => new VMwareNodePoolState();
    }
}
