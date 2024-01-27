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
    /// An AlloyDB Backup.
    /// 
    /// To get more information about Backup, see:
    /// 
    /// * [API documentation](https://cloud.google.com/alloydb/docs/reference/rest/v1/projects.locations.backups/create)
    /// * How-to Guides
    ///     * [AlloyDB](https://cloud.google.com/alloydb/docs/)
    /// 
    /// ## Example Usage
    /// ### Alloydb Backup Basic
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
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             vpcConnection,
    ///         },
    ///     });
    /// 
    ///     var defaultBackup = new Gcp.Alloydb.Backup("defaultBackup", new()
    ///     {
    ///         Location = "us-central1",
    ///         BackupId = "alloydb-backup",
    ///         ClusterName = defaultCluster.Name,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             defaultInstance,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Alloydb Backup Full
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
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             vpcConnection,
    ///         },
    ///     });
    /// 
    ///     var defaultBackup = new Gcp.Alloydb.Backup("defaultBackup", new()
    ///     {
    ///         Location = "us-central1",
    ///         BackupId = "alloydb-backup",
    ///         ClusterName = defaultCluster.Name,
    ///         Description = "example description",
    ///         Type = "ON_DEMAND",
    ///         Labels = 
    ///         {
    ///             { "label", "key" },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             defaultInstance,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Backup can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/backups/{{backup_id}}` * `{{project}}/{{location}}/{{backup_id}}` * `{{location}}/{{backup_id}}` When using the `pulumi import` command, Backup can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:alloydb/backup:Backup default projects/{{project}}/locations/{{location}}/backups/{{backup_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:alloydb/backup:Backup default {{project}}/{{location}}/{{backup_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:alloydb/backup:Backup default {{location}}/{{backup_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:alloydb/backup:Backup")]
    public partial class Backup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128
        /// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
        /// Please refer to the field `effective_annotations` for all of the annotations present on the resource.
        /// </summary>
        [Output("annotations")]
        public Output<ImmutableDictionary<string, string>?> Annotations { get; private set; } = null!;

        /// <summary>
        /// The ID of the alloydb backup.
        /// </summary>
        [Output("backupId")]
        public Output<string> BackupId { get; private set; } = null!;

        /// <summary>
        /// The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// Output only. The system-generated UID of the cluster which was used to create this resource.
        /// </summary>
        [Output("clusterUid")]
        public Output<string> ClusterUid { get; private set; } = null!;

        /// <summary>
        /// Output only. Create time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Output only. Delete time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("deleteTime")]
        public Output<string> DeleteTime { get; private set; } = null!;

        /// <summary>
        /// User-provided description of the backup.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// User-settable and human-readable display name for the Backup.
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
        /// EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
        /// Structure is documented below.
        /// </summary>
        [Output("encryptionConfig")]
        public Output<Outputs.BackupEncryptionConfig?> EncryptionConfig { get; private set; } = null!;

        /// <summary>
        /// EncryptionInfo describes the encryption information of a cluster or a backup.
        /// Structure is documented below.
        /// </summary>
        [Output("encryptionInfos")]
        public Output<ImmutableArray<Outputs.BackupEncryptionInfo>> EncryptionInfos { get; private set; } = null!;

        /// <summary>
        /// For Resource freshness validation (https://google.aip.dev/154)
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Output only. The QuantityBasedExpiry of the backup, specified by the backup's retention policy.
        /// Once the expiry quantity is over retention, the backup is eligible to be garbage collected.
        /// Structure is documented below.
        /// </summary>
        [Output("expiryQuantities")]
        public Output<ImmutableArray<Outputs.BackupExpiryQuantity>> ExpiryQuantities { get; private set; } = null!;

        /// <summary>
        /// Output only. The time at which after the backup is eligible to be garbage collected.
        /// It is the duration specified by the backup's retention policy, added to the backup's createTime.
        /// </summary>
        [Output("expiryTime")]
        public Output<string> ExpiryTime { get; private set; } = null!;

        /// <summary>
        /// User-defined labels for the alloydb backup. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The location where the alloydb backup should reside.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Output only. The name of the backup resource with the format: * projects/{project}/locations/{region}/backups/{backupId}
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
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        [Output("pulumiLabels")]
        public Output<ImmutableDictionary<string, string>> PulumiLabels { get; private set; } = null!;

        /// <summary>
        /// Output only. Reconciling (https://google.aip.dev/128#reconciliation), if true, indicates that the service is actively updating the resource.
        /// This can happen due to user-triggered updates or system actions like failover or maintenance.
        /// </summary>
        [Output("reconciling")]
        public Output<bool> Reconciling { get; private set; } = null!;

        /// <summary>
        /// Output only. The size of the backup in bytes.
        /// </summary>
        [Output("sizeBytes")]
        public Output<string> SizeBytes { get; private set; } = null!;

        /// <summary>
        /// Output only. The current state of the backup.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The backup type, which suggests the trigger for the backup.
        /// Possible values are: `TYPE_UNSPECIFIED`, `ON_DEMAND`, `AUTOMATED`, `CONTINUOUS`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Output only. The system-generated UID of the resource. The UID is assigned when the resource is created, and it is retained until it is deleted.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// Output only. Update time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Backup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Backup(string name, BackupArgs args, CustomResourceOptions? options = null)
            : base("gcp:alloydb/backup:Backup", name, args ?? new BackupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Backup(string name, Input<string> id, BackupState? state = null, CustomResourceOptions? options = null)
            : base("gcp:alloydb/backup:Backup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Backup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Backup Get(string name, Input<string> id, BackupState? state = null, CustomResourceOptions? options = null)
        {
            return new Backup(name, id, state, options);
        }
    }

    public sealed class BackupArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128
        /// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
        /// Please refer to the field `effective_annotations` for all of the annotations present on the resource.
        /// </summary>
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// The ID of the alloydb backup.
        /// </summary>
        [Input("backupId", required: true)]
        public Input<string> BackupId { get; set; } = null!;

        /// <summary>
        /// The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// User-provided description of the backup.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User-settable and human-readable display name for the Backup.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
        /// Structure is documented below.
        /// </summary>
        [Input("encryptionConfig")]
        public Input<Inputs.BackupEncryptionConfigArgs>? EncryptionConfig { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-defined labels for the alloydb backup. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
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
        /// The location where the alloydb backup should reside.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The backup type, which suggests the trigger for the backup.
        /// Possible values are: `TYPE_UNSPECIFIED`, `ON_DEMAND`, `AUTOMATED`, `CONTINUOUS`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public BackupArgs()
        {
        }
        public static new BackupArgs Empty => new BackupArgs();
    }

    public sealed class BackupState : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128
        /// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
        /// Please refer to the field `effective_annotations` for all of the annotations present on the resource.
        /// </summary>
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// The ID of the alloydb backup.
        /// </summary>
        [Input("backupId")]
        public Input<string>? BackupId { get; set; }

        /// <summary>
        /// The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// Output only. The system-generated UID of the cluster which was used to create this resource.
        /// </summary>
        [Input("clusterUid")]
        public Input<string>? ClusterUid { get; set; }

        /// <summary>
        /// Output only. Create time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Output only. Delete time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("deleteTime")]
        public Input<string>? DeleteTime { get; set; }

        /// <summary>
        /// User-provided description of the backup.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User-settable and human-readable display name for the Backup.
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
        /// EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
        /// Structure is documented below.
        /// </summary>
        [Input("encryptionConfig")]
        public Input<Inputs.BackupEncryptionConfigGetArgs>? EncryptionConfig { get; set; }

        [Input("encryptionInfos")]
        private InputList<Inputs.BackupEncryptionInfoGetArgs>? _encryptionInfos;

        /// <summary>
        /// EncryptionInfo describes the encryption information of a cluster or a backup.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.BackupEncryptionInfoGetArgs> EncryptionInfos
        {
            get => _encryptionInfos ?? (_encryptionInfos = new InputList<Inputs.BackupEncryptionInfoGetArgs>());
            set => _encryptionInfos = value;
        }

        /// <summary>
        /// For Resource freshness validation (https://google.aip.dev/154)
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("expiryQuantities")]
        private InputList<Inputs.BackupExpiryQuantityGetArgs>? _expiryQuantities;

        /// <summary>
        /// Output only. The QuantityBasedExpiry of the backup, specified by the backup's retention policy.
        /// Once the expiry quantity is over retention, the backup is eligible to be garbage collected.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.BackupExpiryQuantityGetArgs> ExpiryQuantities
        {
            get => _expiryQuantities ?? (_expiryQuantities = new InputList<Inputs.BackupExpiryQuantityGetArgs>());
            set => _expiryQuantities = value;
        }

        /// <summary>
        /// Output only. The time at which after the backup is eligible to be garbage collected.
        /// It is the duration specified by the backup's retention policy, added to the backup's createTime.
        /// </summary>
        [Input("expiryTime")]
        public Input<string>? ExpiryTime { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-defined labels for the alloydb backup. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
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
        /// The location where the alloydb backup should reside.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Output only. The name of the backup resource with the format: * projects/{project}/locations/{region}/backups/{backupId}
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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
        /// Output only. Reconciling (https://google.aip.dev/128#reconciliation), if true, indicates that the service is actively updating the resource.
        /// This can happen due to user-triggered updates or system actions like failover or maintenance.
        /// </summary>
        [Input("reconciling")]
        public Input<bool>? Reconciling { get; set; }

        /// <summary>
        /// Output only. The size of the backup in bytes.
        /// </summary>
        [Input("sizeBytes")]
        public Input<string>? SizeBytes { get; set; }

        /// <summary>
        /// Output only. The current state of the backup.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The backup type, which suggests the trigger for the backup.
        /// Possible values are: `TYPE_UNSPECIFIED`, `ON_DEMAND`, `AUTOMATED`, `CONTINUOUS`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Output only. The system-generated UID of the resource. The UID is assigned when the resource is created, and it is retained until it is deleted.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// Output only. Update time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public BackupState()
        {
        }
        public static new BackupState Empty => new BackupState();
    }
}
