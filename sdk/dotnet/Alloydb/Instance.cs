// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Alloydb
{
    /// <summary>
    /// A managed alloydb cluster instance.
    /// 
    /// To get more information about Instance, see:
    /// 
    /// * [API documentation](https://cloud.google.com/alloydb/docs/reference/rest/v1/projects.locations.clusters.instances/create)
    /// * How-to Guides
    ///     * [AlloyDB](https://cloud.google.com/alloydb/docs/)
    /// 
    /// ## Example Usage
    /// ### Alloydb Instance Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var defaultNetwork = new Gcp.Compute.Network("defaultNetwork");
    /// 
    ///     var defaultCluster = new Gcp.Alloydb.Cluster("defaultCluster", new()
    ///     {
    ///         ClusterId = "alloydb-cluster",
    ///         Location = "us-central1",
    ///         Network = defaultNetwork.Id,
    ///         InitialUser = new Gcp.Alloydb.Inputs.ClusterInitialUserArgs
    ///         {
    ///             Password = "alloydb-cluster",
    ///         },
    ///     });
    /// 
    ///     var privateIpAlloc = new Gcp.Compute.GlobalAddress("privateIpAlloc", new()
    ///     {
    ///         AddressType = "INTERNAL",
    ///         Purpose = "VPC_PEERING",
    ///         PrefixLength = 16,
    ///         Network = defaultNetwork.Id,
    ///     });
    /// 
    ///     var vpcConnection = new Gcp.ServiceNetworking.Connection("vpcConnection", new()
    ///     {
    ///         Network = defaultNetwork.Id,
    ///         Service = "servicenetworking.googleapis.com",
    ///         ReservedPeeringRanges = new[]
    ///         {
    ///             privateIpAlloc.Name,
    ///         },
    ///     });
    /// 
    ///     var defaultInstance = new Gcp.Alloydb.Instance("defaultInstance", new()
    ///     {
    ///         Cluster = defaultCluster.Name,
    ///         InstanceId = "alloydb-instance",
    ///         InstanceType = "PRIMARY",
    ///         MachineConfig = new Gcp.Alloydb.Inputs.InstanceMachineConfigArgs
    ///         {
    ///             CpuCount = 2,
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             vpcConnection,
    ///         },
    ///     });
    /// 
    ///     var project = Gcp.Organizations.GetProject.Invoke();
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Instance can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:alloydb/instance:Instance default projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/instances/{{instance_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:alloydb/instance:Instance default {{project}}/{{location}}/{{cluster}}/{{instance_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:alloydb/instance:Instance default {{location}}/{{cluster}}/{{instance_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:alloydb/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.
        /// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
        /// Please refer to the field `effective_annotations` for all of the annotations present on the resource.
        /// </summary>
        [Output("annotations")]
        public Output<ImmutableDictionary<string, string>?> Annotations { get; private set; } = null!;

        /// <summary>
        /// 'Availability type of an Instance. Defaults to REGIONAL for both primary and read instances.
        /// Note that primary and read instances can have different availability types.
        /// Only READ_POOL instance supports ZONAL type. Users can't specify the zone for READ_POOL instance.
        /// Zone is automatically chosen from the list of zones in the region specified.
        /// Read pool of size 1 can only have zonal availability. Read pools with node count of 2 or more
        /// can have regional availability (nodes are present in 2 or more zones in a region).'
        /// Possible values are: `AVAILABILITY_TYPE_UNSPECIFIED`, `ZONAL`, `REGIONAL`.
        /// </summary>
        [Output("availabilityType")]
        public Output<string> AvailabilityType { get; private set; } = null!;

        /// <summary>
        /// Identifies the alloydb cluster. Must be in the format
        /// 'projects/{project}/locations/{location}/clusters/{cluster_id}'
        /// </summary>
        [Output("cluster")]
        public Output<string> Cluster { get; private set; } = null!;

        /// <summary>
        /// Time the Instance was created in UTC.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Database flags. Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary.
        /// </summary>
        [Output("databaseFlags")]
        public Output<ImmutableDictionary<string, string>?> DatabaseFlags { get; private set; } = null!;

        /// <summary>
        /// User-settable and human-readable display name for the Instance.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through
        /// Terraform, other clients and services.
        /// </summary>
        [Output("effectiveAnnotations")]
        public Output<ImmutableDictionary<string, string>> EffectiveAnnotations { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.
        /// </summary>
        [Output("gceZone")]
        public Output<string?> GceZone { get; private set; } = null!;

        /// <summary>
        /// The ID of the alloydb instance.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The type of the instance. If the instance type is READ_POOL, provide the associated PRIMARY instance in the `depends_on` meta-data attribute.
        /// Possible values are: `PRIMARY`, `READ_POOL`.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The IP address for the Instance. This is the connection endpoint for an end-user application.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// User-defined labels for the alloydb instance.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Configurations for the machines that host the underlying database engine.
        /// Structure is documented below.
        /// </summary>
        [Output("machineConfig")]
        public Output<Outputs.InstanceMachineConfig> MachineConfig { get; private set; } = null!;

        /// <summary>
        /// The name of the instance resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        [Output("pulumiLabels")]
        public Output<ImmutableDictionary<string, string>> PulumiLabels { get; private set; } = null!;

        /// <summary>
        /// Configuration for query insights.
        /// Structure is documented below.
        /// </summary>
        [Output("queryInsightsConfig")]
        public Output<Outputs.InstanceQueryInsightsConfig> QueryInsightsConfig { get; private set; } = null!;

        /// <summary>
        /// Read pool specific config. If the instance type is READ_POOL, this configuration must be provided.
        /// Structure is documented below.
        /// </summary>
        [Output("readPoolConfig")]
        public Output<Outputs.InstanceReadPoolConfig?> ReadPoolConfig { get; private set; } = null!;

        /// <summary>
        /// Set to true if the current state of Instance does not match the user's intended state, and the service is actively updating the resource to reconcile them. This can happen due to user-triggered updates or system actions like failover or maintenance.
        /// </summary>
        [Output("reconciling")]
        public Output<bool> Reconciling { get; private set; } = null!;

        /// <summary>
        /// The current state of the alloydb instance.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The system-generated UID of the resource.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// Time the Instance was updated in UTC.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("gcp:alloydb/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("gcp:alloydb/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.
        /// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
        /// Please refer to the field `effective_annotations` for all of the annotations present on the resource.
        /// </summary>
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// 'Availability type of an Instance. Defaults to REGIONAL for both primary and read instances.
        /// Note that primary and read instances can have different availability types.
        /// Only READ_POOL instance supports ZONAL type. Users can't specify the zone for READ_POOL instance.
        /// Zone is automatically chosen from the list of zones in the region specified.
        /// Read pool of size 1 can only have zonal availability. Read pools with node count of 2 or more
        /// can have regional availability (nodes are present in 2 or more zones in a region).'
        /// Possible values are: `AVAILABILITY_TYPE_UNSPECIFIED`, `ZONAL`, `REGIONAL`.
        /// </summary>
        [Input("availabilityType")]
        public Input<string>? AvailabilityType { get; set; }

        /// <summary>
        /// Identifies the alloydb cluster. Must be in the format
        /// 'projects/{project}/locations/{location}/clusters/{cluster_id}'
        /// </summary>
        [Input("cluster", required: true)]
        public Input<string> Cluster { get; set; } = null!;

        [Input("databaseFlags")]
        private InputMap<string>? _databaseFlags;

        /// <summary>
        /// Database flags. Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary.
        /// </summary>
        public InputMap<string> DatabaseFlags
        {
            get => _databaseFlags ?? (_databaseFlags = new InputMap<string>());
            set => _databaseFlags = value;
        }

        /// <summary>
        /// User-settable and human-readable display name for the Instance.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.
        /// </summary>
        [Input("gceZone")]
        public Input<string>? GceZone { get; set; }

        /// <summary>
        /// The ID of the alloydb instance.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The type of the instance. If the instance type is READ_POOL, provide the associated PRIMARY instance in the `depends_on` meta-data attribute.
        /// Possible values are: `PRIMARY`, `READ_POOL`.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-defined labels for the alloydb instance.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Configurations for the machines that host the underlying database engine.
        /// Structure is documented below.
        /// </summary>
        [Input("machineConfig")]
        public Input<Inputs.InstanceMachineConfigArgs>? MachineConfig { get; set; }

        /// <summary>
        /// Configuration for query insights.
        /// Structure is documented below.
        /// </summary>
        [Input("queryInsightsConfig")]
        public Input<Inputs.InstanceQueryInsightsConfigArgs>? QueryInsightsConfig { get; set; }

        /// <summary>
        /// Read pool specific config. If the instance type is READ_POOL, this configuration must be provided.
        /// Structure is documented below.
        /// </summary>
        [Input("readPoolConfig")]
        public Input<Inputs.InstanceReadPoolConfigArgs>? ReadPoolConfig { get; set; }

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.
        /// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
        /// Please refer to the field `effective_annotations` for all of the annotations present on the resource.
        /// </summary>
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// 'Availability type of an Instance. Defaults to REGIONAL for both primary and read instances.
        /// Note that primary and read instances can have different availability types.
        /// Only READ_POOL instance supports ZONAL type. Users can't specify the zone for READ_POOL instance.
        /// Zone is automatically chosen from the list of zones in the region specified.
        /// Read pool of size 1 can only have zonal availability. Read pools with node count of 2 or more
        /// can have regional availability (nodes are present in 2 or more zones in a region).'
        /// Possible values are: `AVAILABILITY_TYPE_UNSPECIFIED`, `ZONAL`, `REGIONAL`.
        /// </summary>
        [Input("availabilityType")]
        public Input<string>? AvailabilityType { get; set; }

        /// <summary>
        /// Identifies the alloydb cluster. Must be in the format
        /// 'projects/{project}/locations/{location}/clusters/{cluster_id}'
        /// </summary>
        [Input("cluster")]
        public Input<string>? Cluster { get; set; }

        /// <summary>
        /// Time the Instance was created in UTC.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        [Input("databaseFlags")]
        private InputMap<string>? _databaseFlags;

        /// <summary>
        /// Database flags. Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary.
        /// </summary>
        public InputMap<string> DatabaseFlags
        {
            get => _databaseFlags ?? (_databaseFlags = new InputMap<string>());
            set => _databaseFlags = value;
        }

        /// <summary>
        /// User-settable and human-readable display name for the Instance.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("effectiveAnnotations")]
        private InputMap<string>? _effectiveAnnotations;

        /// <summary>
        /// All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through
        /// Terraform, other clients and services.
        /// </summary>
        public InputMap<string> EffectiveAnnotations
        {
            get => _effectiveAnnotations ?? (_effectiveAnnotations = new InputMap<string>());
            set => _effectiveAnnotations = value;
        }

        [Input("effectiveLabels")]
        private InputMap<string>? _effectiveLabels;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
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

        /// <summary>
        /// The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.
        /// </summary>
        [Input("gceZone")]
        public Input<string>? GceZone { get; set; }

        /// <summary>
        /// The ID of the alloydb instance.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The type of the instance. If the instance type is READ_POOL, provide the associated PRIMARY instance in the `depends_on` meta-data attribute.
        /// Possible values are: `PRIMARY`, `READ_POOL`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The IP address for the Instance. This is the connection endpoint for an end-user application.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-defined labels for the alloydb instance.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Configurations for the machines that host the underlying database engine.
        /// Structure is documented below.
        /// </summary>
        [Input("machineConfig")]
        public Input<Inputs.InstanceMachineConfigGetArgs>? MachineConfig { get; set; }

        /// <summary>
        /// The name of the instance resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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
        /// Configuration for query insights.
        /// Structure is documented below.
        /// </summary>
        [Input("queryInsightsConfig")]
        public Input<Inputs.InstanceQueryInsightsConfigGetArgs>? QueryInsightsConfig { get; set; }

        /// <summary>
        /// Read pool specific config. If the instance type is READ_POOL, this configuration must be provided.
        /// Structure is documented below.
        /// </summary>
        [Input("readPoolConfig")]
        public Input<Inputs.InstanceReadPoolConfigGetArgs>? ReadPoolConfig { get; set; }

        /// <summary>
        /// Set to true if the current state of Instance does not match the user's intended state, and the service is actively updating the resource to reconcile them. This can happen due to user-triggered updates or system actions like failover or maintenance.
        /// </summary>
        [Input("reconciling")]
        public Input<bool>? Reconciling { get; set; }

        /// <summary>
        /// The current state of the alloydb instance.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The system-generated UID of the resource.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// Time the Instance was updated in UTC.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}
