// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Memcache
{
    /// <summary>
    /// A Google Cloud Memcache instance.
    /// 
    /// To get more information about Instance, see:
    /// 
    /// * [API documentation](https://cloud.google.com/memorystore/docs/memcached/reference/rest/v1/projects.locations.instances)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/memcache/docs/creating-instances)
    /// 
    /// ## Example Usage
    /// ### Memcache Instance Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // This example assumes this network already exists.
    ///     // The API creates a tenant network per network authorized for a
    ///     // Memcache instance and that network is not deleted when the user-created
    ///     // network (authorized_network) is deleted, so this prevents issues
    ///     // with tenant network quota.
    ///     // If this network hasn't been created and you are using this example in your
    ///     // config, add an additional network resource or change
    ///     // this from "data"to "resource"
    ///     var memcacheNetwork = new Gcp.Compute.Network("memcacheNetwork");
    /// 
    ///     var serviceRange = new Gcp.Compute.GlobalAddress("serviceRange", new()
    ///     {
    ///         Purpose = "VPC_PEERING",
    ///         AddressType = "INTERNAL",
    ///         PrefixLength = 16,
    ///         Network = memcacheNetwork.Id,
    ///     });
    /// 
    ///     var privateServiceConnection = new Gcp.ServiceNetworking.Connection("privateServiceConnection", new()
    ///     {
    ///         Network = memcacheNetwork.Id,
    ///         Service = "servicenetworking.googleapis.com",
    ///         ReservedPeeringRanges = new[]
    ///         {
    ///             serviceRange.Name,
    ///         },
    ///     });
    /// 
    ///     var instance = new Gcp.Memcache.Instance("instance", new()
    ///     {
    ///         AuthorizedNetwork = privateServiceConnection.Network,
    ///         Labels = 
    ///         {
    ///             { "env", "test" },
    ///         },
    ///         NodeConfig = new Gcp.Memcache.Inputs.InstanceNodeConfigArgs
    ///         {
    ///             CpuCount = 1,
    ///             MemorySizeMb = 1024,
    ///         },
    ///         NodeCount = 1,
    ///         MemcacheVersion = "MEMCACHE_1_5",
    ///         MaintenancePolicy = new Gcp.Memcache.Inputs.InstanceMaintenancePolicyArgs
    ///         {
    ///             WeeklyMaintenanceWindows = new[]
    ///             {
    ///                 new Gcp.Memcache.Inputs.InstanceMaintenancePolicyWeeklyMaintenanceWindowArgs
    ///                 {
    ///                     Day = "SATURDAY",
    ///                     Duration = "14400s",
    ///                     StartTime = new Gcp.Memcache.Inputs.InstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeArgs
    ///                     {
    ///                         Hours = 0,
    ///                         Minutes = 30,
    ///                         Seconds = 0,
    ///                         Nanos = 0,
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Instance can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:memcache/instance:Instance default projects/{{project}}/locations/{{region}}/instances/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:memcache/instance:Instance default {{project}}/{{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:memcache/instance:Instance default {{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:memcache/instance:Instance default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:memcache/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The full name of the GCE network to connect the instance to.  If not provided,
        /// 'default' will be used.
        /// </summary>
        [Output("authorizedNetwork")]
        public Output<string> AuthorizedNetwork { get; private set; } = null!;

        /// <summary>
        /// (Output)
        /// Output only. The time when the policy was created.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
        /// resolution and up to nine fractional digits
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Endpoint for Discovery API
        /// </summary>
        [Output("discoveryEndpoint")]
        public Output<string> DiscoveryEndpoint { get; private set; } = null!;

        /// <summary>
        /// A user-visible name for the instance.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
        /// clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Maintenance policy for an instance.
        /// Structure is documented below.
        /// </summary>
        [Output("maintenancePolicy")]
        public Output<Outputs.InstanceMaintenancePolicy?> MaintenancePolicy { get; private set; } = null!;

        /// <summary>
        /// Output only. Published maintenance schedule.
        /// Structure is documented below.
        /// </summary>
        [Output("maintenanceSchedules")]
        public Output<ImmutableArray<Outputs.InstanceMaintenanceSchedule>> MaintenanceSchedules { get; private set; } = null!;

        /// <summary>
        /// The full version of memcached server running on this instance.
        /// </summary>
        [Output("memcacheFullVersion")]
        public Output<string> MemcacheFullVersion { get; private set; } = null!;

        /// <summary>
        /// Additional information about the instance state, if available.
        /// Structure is documented below.
        /// </summary>
        [Output("memcacheNodes")]
        public Output<ImmutableArray<Outputs.InstanceMemcacheNode>> MemcacheNodes { get; private set; } = null!;

        /// <summary>
        /// User-specified parameters for this memcache instance.
        /// Structure is documented below.
        /// </summary>
        [Output("memcacheParameters")]
        public Output<Outputs.InstanceMemcacheParameters?> MemcacheParameters { get; private set; } = null!;

        /// <summary>
        /// The major version of Memcached software. If not provided, latest supported version will be used.
        /// Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
        /// determined by our system based on the latest supported minor version.
        /// Default value is `MEMCACHE_1_5`.
        /// Possible values are: `MEMCACHE_1_5`.
        /// </summary>
        [Output("memcacheVersion")]
        public Output<string?> MemcacheVersion { get; private set; } = null!;

        /// <summary>
        /// The resource name of the instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Configuration for memcache nodes.
        /// Structure is documented below.
        /// </summary>
        [Output("nodeConfig")]
        public Output<Outputs.InstanceNodeConfig> NodeConfig { get; private set; } = null!;

        /// <summary>
        /// Number of nodes in the memcache instance.
        /// </summary>
        [Output("nodeCount")]
        public Output<int> NodeCount { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        [Output("pulumiLabels")]
        public Output<ImmutableDictionary<string, string>> PulumiLabels { get; private set; } = null!;

        /// <summary>
        /// The region of the Memcache instance. If it is not provided, the provider region is used.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Zones where memcache nodes should be provisioned.  If not
        /// provided, all zones will be used.
        /// </summary>
        [Output("zones")]
        public Output<ImmutableArray<string>> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("gcp:memcache/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("gcp:memcache/instance:Instance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "effectiveLabels",
                    "pulumiLabels",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The full name of the GCE network to connect the instance to.  If not provided,
        /// 'default' will be used.
        /// </summary>
        [Input("authorizedNetwork")]
        public Input<string>? AuthorizedNetwork { get; set; }

        /// <summary>
        /// A user-visible name for the instance.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Maintenance policy for an instance.
        /// Structure is documented below.
        /// </summary>
        [Input("maintenancePolicy")]
        public Input<Inputs.InstanceMaintenancePolicyArgs>? MaintenancePolicy { get; set; }

        /// <summary>
        /// User-specified parameters for this memcache instance.
        /// Structure is documented below.
        /// </summary>
        [Input("memcacheParameters")]
        public Input<Inputs.InstanceMemcacheParametersArgs>? MemcacheParameters { get; set; }

        /// <summary>
        /// The major version of Memcached software. If not provided, latest supported version will be used.
        /// Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
        /// determined by our system based on the latest supported minor version.
        /// Default value is `MEMCACHE_1_5`.
        /// Possible values are: `MEMCACHE_1_5`.
        /// </summary>
        [Input("memcacheVersion")]
        public Input<string>? MemcacheVersion { get; set; }

        /// <summary>
        /// The resource name of the instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configuration for memcache nodes.
        /// Structure is documented below.
        /// </summary>
        [Input("nodeConfig", required: true)]
        public Input<Inputs.InstanceNodeConfigArgs> NodeConfig { get; set; } = null!;

        /// <summary>
        /// Number of nodes in the memcache instance.
        /// </summary>
        [Input("nodeCount", required: true)]
        public Input<int> NodeCount { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region of the Memcache instance. If it is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// Zones where memcache nodes should be provisioned.  If not
        /// provided, all zones will be used.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The full name of the GCE network to connect the instance to.  If not provided,
        /// 'default' will be used.
        /// </summary>
        [Input("authorizedNetwork")]
        public Input<string>? AuthorizedNetwork { get; set; }

        /// <summary>
        /// (Output)
        /// Output only. The time when the policy was created.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
        /// resolution and up to nine fractional digits
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Endpoint for Discovery API
        /// </summary>
        [Input("discoveryEndpoint")]
        public Input<string>? DiscoveryEndpoint { get; set; }

        /// <summary>
        /// A user-visible name for the instance.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("effectiveLabels")]
        private InputMap<string>? _effectiveLabels;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
        /// clients and services.
        /// </summary>
        public InputMap<string> EffectiveLabels
        {
            get => _effectiveLabels ?? (_effectiveLabels = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _effectiveLabels = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Maintenance policy for an instance.
        /// Structure is documented below.
        /// </summary>
        [Input("maintenancePolicy")]
        public Input<Inputs.InstanceMaintenancePolicyGetArgs>? MaintenancePolicy { get; set; }

        [Input("maintenanceSchedules")]
        private InputList<Inputs.InstanceMaintenanceScheduleGetArgs>? _maintenanceSchedules;

        /// <summary>
        /// Output only. Published maintenance schedule.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.InstanceMaintenanceScheduleGetArgs> MaintenanceSchedules
        {
            get => _maintenanceSchedules ?? (_maintenanceSchedules = new InputList<Inputs.InstanceMaintenanceScheduleGetArgs>());
            set => _maintenanceSchedules = value;
        }

        /// <summary>
        /// The full version of memcached server running on this instance.
        /// </summary>
        [Input("memcacheFullVersion")]
        public Input<string>? MemcacheFullVersion { get; set; }

        [Input("memcacheNodes")]
        private InputList<Inputs.InstanceMemcacheNodeGetArgs>? _memcacheNodes;

        /// <summary>
        /// Additional information about the instance state, if available.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.InstanceMemcacheNodeGetArgs> MemcacheNodes
        {
            get => _memcacheNodes ?? (_memcacheNodes = new InputList<Inputs.InstanceMemcacheNodeGetArgs>());
            set => _memcacheNodes = value;
        }

        /// <summary>
        /// User-specified parameters for this memcache instance.
        /// Structure is documented below.
        /// </summary>
        [Input("memcacheParameters")]
        public Input<Inputs.InstanceMemcacheParametersGetArgs>? MemcacheParameters { get; set; }

        /// <summary>
        /// The major version of Memcached software. If not provided, latest supported version will be used.
        /// Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
        /// determined by our system based on the latest supported minor version.
        /// Default value is `MEMCACHE_1_5`.
        /// Possible values are: `MEMCACHE_1_5`.
        /// </summary>
        [Input("memcacheVersion")]
        public Input<string>? MemcacheVersion { get; set; }

        /// <summary>
        /// The resource name of the instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configuration for memcache nodes.
        /// Structure is documented below.
        /// </summary>
        [Input("nodeConfig")]
        public Input<Inputs.InstanceNodeConfigGetArgs>? NodeConfig { get; set; }

        /// <summary>
        /// Number of nodes in the memcache instance.
        /// </summary>
        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("pulumiLabels")]
        private InputMap<string>? _pulumiLabels;

        /// <summary>
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        public InputMap<string> PulumiLabels
        {
            get => _pulumiLabels ?? (_pulumiLabels = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _pulumiLabels = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// The region of the Memcache instance. If it is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// Zones where memcache nodes should be provisioned.  If not
        /// provided, all zones will be used.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}
