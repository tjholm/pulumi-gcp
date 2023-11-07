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
    /// A managed alloydb cluster.
    /// 
    /// To get more information about Cluster, see:
    /// 
    /// * [API documentation](https://cloud.google.com/alloydb/docs/reference/rest/v1/projects.locations.clusters/create)
    /// * How-to Guides
    ///     * [AlloyDB](https://cloud.google.com/alloydb/docs/)
    /// 
    /// &gt; **Warning:** All arguments including the following potentially sensitive
    /// values will be stored in the raw state as plain text: `initial_user.password`.
    /// Read more about sensitive data in state.
    /// 
    /// ## Example Usage
    /// ### Alloydb Cluster Basic
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
    ///     var project = Gcp.Organizations.GetProject.Invoke();
    /// 
    /// });
    /// ```
    /// ### Alloydb Cluster Full
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.Compute.Network("default");
    /// 
    ///     var full = new Gcp.Alloydb.Cluster("full", new()
    ///     {
    ///         ClusterId = "alloydb-cluster-full",
    ///         Location = "us-central1",
    ///         Network = @default.Id,
    ///         InitialUser = new Gcp.Alloydb.Inputs.ClusterInitialUserArgs
    ///         {
    ///             User = "alloydb-cluster-full",
    ///             Password = "alloydb-cluster-full",
    ///         },
    ///         ContinuousBackupConfig = new Gcp.Alloydb.Inputs.ClusterContinuousBackupConfigArgs
    ///         {
    ///             Enabled = true,
    ///             RecoveryWindowDays = 14,
    ///         },
    ///         AutomatedBackupPolicy = new Gcp.Alloydb.Inputs.ClusterAutomatedBackupPolicyArgs
    ///         {
    ///             Location = "us-central1",
    ///             BackupWindow = "1800s",
    ///             Enabled = true,
    ///             WeeklySchedule = new Gcp.Alloydb.Inputs.ClusterAutomatedBackupPolicyWeeklyScheduleArgs
    ///             {
    ///                 DaysOfWeeks = new[]
    ///                 {
    ///                     "MONDAY",
    ///                 },
    ///                 StartTimes = new[]
    ///                 {
    ///                     new Gcp.Alloydb.Inputs.ClusterAutomatedBackupPolicyWeeklyScheduleStartTimeArgs
    ///                     {
    ///                         Hours = 23,
    ///                         Minutes = 0,
    ///                         Seconds = 0,
    ///                         Nanos = 0,
    ///                     },
    ///                 },
    ///             },
    ///             QuantityBasedRetention = new Gcp.Alloydb.Inputs.ClusterAutomatedBackupPolicyQuantityBasedRetentionArgs
    ///             {
    ///                 Count = 1,
    ///             },
    ///             Labels = 
    ///             {
    ///                 { "test", "alloydb-cluster-full" },
    ///             },
    ///         },
    ///         Labels = 
    ///         {
    ///             { "test", "alloydb-cluster-full" },
    ///         },
    ///     });
    /// 
    ///     var project = Gcp.Organizations.GetProject.Invoke();
    /// 
    /// });
    /// ```
    /// ### Alloydb Cluster Restore
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = Gcp.Compute.GetNetwork.Invoke(new()
    ///     {
    ///         Name = "alloydb-network",
    ///     });
    /// 
    ///     var sourceCluster = new Gcp.Alloydb.Cluster("sourceCluster", new()
    ///     {
    ///         ClusterId = "alloydb-source-cluster",
    ///         Location = "us-central1",
    ///         Network = @default.Apply(@default =&gt; @default.Apply(getNetworkResult =&gt; getNetworkResult.Id)),
    ///         InitialUser = new Gcp.Alloydb.Inputs.ClusterInitialUserArgs
    ///         {
    ///             Password = "alloydb-source-cluster",
    ///         },
    ///     });
    /// 
    ///     var privateIpAlloc = new Gcp.Compute.GlobalAddress("privateIpAlloc", new()
    ///     {
    ///         AddressType = "INTERNAL",
    ///         Purpose = "VPC_PEERING",
    ///         PrefixLength = 16,
    ///         Network = @default.Apply(@default =&gt; @default.Apply(getNetworkResult =&gt; getNetworkResult.Id)),
    ///     });
    /// 
    ///     var vpcConnection = new Gcp.ServiceNetworking.Connection("vpcConnection", new()
    ///     {
    ///         Network = @default.Apply(@default =&gt; @default.Apply(getNetworkResult =&gt; getNetworkResult.Id)),
    ///         Service = "servicenetworking.googleapis.com",
    ///         ReservedPeeringRanges = new[]
    ///         {
    ///             privateIpAlloc.Name,
    ///         },
    ///     });
    /// 
    ///     var sourceInstance = new Gcp.Alloydb.Instance("sourceInstance", new()
    ///     {
    ///         Cluster = sourceCluster.Name,
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
    ///     var sourceBackup = new Gcp.Alloydb.Backup("sourceBackup", new()
    ///     {
    ///         BackupId = "alloydb-backup",
    ///         Location = "us-central1",
    ///         ClusterName = sourceCluster.Name,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             sourceInstance,
    ///         },
    ///     });
    /// 
    ///     var restoredFromBackup = new Gcp.Alloydb.Cluster("restoredFromBackup", new()
    ///     {
    ///         ClusterId = "alloydb-backup-restored",
    ///         Location = "us-central1",
    ///         Network = @default.Apply(@default =&gt; @default.Apply(getNetworkResult =&gt; getNetworkResult.Id)),
    ///         RestoreBackupSource = new Gcp.Alloydb.Inputs.ClusterRestoreBackupSourceArgs
    ///         {
    ///             BackupName = sourceBackup.Name,
    ///         },
    ///     });
    /// 
    ///     var restoredViaPitr = new Gcp.Alloydb.Cluster("restoredViaPitr", new()
    ///     {
    ///         ClusterId = "alloydb-pitr-restored",
    ///         Location = "us-central1",
    ///         Network = @default.Apply(@default =&gt; @default.Apply(getNetworkResult =&gt; getNetworkResult.Id)),
    ///         RestoreContinuousBackupSource = new Gcp.Alloydb.Inputs.ClusterRestoreContinuousBackupSourceArgs
    ///         {
    ///             Cluster = sourceCluster.Name,
    ///             PointInTime = "2023-08-03T19:19:00.094Z",
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
    /// Cluster can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:alloydb/cluster:Cluster default projects/{{project}}/locations/{{location}}/clusters/{{cluster_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:alloydb/cluster:Cluster default {{project}}/{{location}}/{{cluster_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:alloydb/cluster:Cluster default {{location}}/{{cluster_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:alloydb/cluster:Cluster default {{cluster_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:alloydb/cluster:Cluster")]
    public partial class Cluster : global::Pulumi.CustomResource
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
        /// The automated backup policy for this cluster. AutomatedBackupPolicy is disabled by default.
        /// Structure is documented below.
        /// </summary>
        [Output("automatedBackupPolicy")]
        public Output<Outputs.ClusterAutomatedBackupPolicy> AutomatedBackupPolicy { get; private set; } = null!;

        /// <summary>
        /// Cluster created from backup.
        /// Structure is documented below.
        /// </summary>
        [Output("backupSources")]
        public Output<ImmutableArray<Outputs.ClusterBackupSource>> BackupSources { get; private set; } = null!;

        /// <summary>
        /// The ID of the alloydb cluster.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The continuous backup config for this cluster.
        /// If no policy is provided then the default policy will be used. The default policy takes one backup a day and retains backups for 14 days.
        /// Structure is documented below.
        /// </summary>
        [Output("continuousBackupConfig")]
        public Output<Outputs.ClusterContinuousBackupConfig> ContinuousBackupConfig { get; private set; } = null!;

        /// <summary>
        /// ContinuousBackupInfo describes the continuous backup properties of a cluster.
        /// Structure is documented below.
        /// </summary>
        [Output("continuousBackupInfos")]
        public Output<ImmutableArray<Outputs.ClusterContinuousBackupInfo>> ContinuousBackupInfos { get; private set; } = null!;

        /// <summary>
        /// The database engine major version. This is an output-only field and it's populated at the Cluster creation time. This field cannot be changed after cluster creation.
        /// </summary>
        [Output("databaseVersion")]
        public Output<string> DatabaseVersion { get; private set; } = null!;

        /// <summary>
        /// User-settable and human-readable display name for the Cluster.
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
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
        /// clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
        /// Structure is documented below.
        /// </summary>
        [Output("encryptionConfig")]
        public Output<Outputs.ClusterEncryptionConfig?> EncryptionConfig { get; private set; } = null!;

        /// <summary>
        /// (Output)
        /// Output only. The encryption information for the WALs and backups required for ContinuousBackup.
        /// Structure is documented below.
        /// </summary>
        [Output("encryptionInfos")]
        public Output<ImmutableArray<Outputs.ClusterEncryptionInfo>> EncryptionInfos { get; private set; } = null!;

        /// <summary>
        /// For Resource freshness validation (https://google.aip.dev/154)
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Initial user to setup during cluster creation.
        /// Structure is documented below.
        /// </summary>
        [Output("initialUser")]
        public Output<Outputs.ClusterInitialUser?> InitialUser { get; private set; } = null!;

        /// <summary>
        /// User-defined labels for the alloydb cluster.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The location where the alloydb cluster should reside.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Cluster created via DMS migration.
        /// Structure is documented below.
        /// </summary>
        [Output("migrationSources")]
        public Output<ImmutableArray<Outputs.ClusterMigrationSource>> MigrationSources { get; private set; } = null!;

        /// <summary>
        /// The name of the cluster resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// (Optional, Deprecated)
        /// The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
        /// "projects/{projectNumber}/global/networks/{network_id}".
        /// 
        /// &gt; **Warning:** `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// Metadata related to network configuration.
        /// Structure is documented below.
        /// </summary>
        [Output("networkConfig")]
        public Output<Outputs.ClusterNetworkConfig> NetworkConfig { get; private set; } = null!;

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
        /// Output only. Reconciling (https://google.aip.dev/128#reconciliation).
        /// Set to true if the current state of Cluster does not match the user's intended state, and the service is actively updating the resource to reconcile them.
        /// This can happen due to user-triggered updates or system actions like failover or maintenance.
        /// </summary>
        [Output("reconciling")]
        public Output<bool> Reconciling { get; private set; } = null!;

        /// <summary>
        /// The source when restoring from a backup. Conflicts with 'restore_continuous_backup_source', both can't be set together.
        /// Structure is documented below.
        /// </summary>
        [Output("restoreBackupSource")]
        public Output<Outputs.ClusterRestoreBackupSource?> RestoreBackupSource { get; private set; } = null!;

        /// <summary>
        /// The source when restoring via point in time recovery (PITR). Conflicts with 'restore_backup_source', both can't be set together.
        /// Structure is documented below.
        /// </summary>
        [Output("restoreContinuousBackupSource")]
        public Output<Outputs.ClusterRestoreContinuousBackupSource?> RestoreContinuousBackupSource { get; private set; } = null!;

        /// <summary>
        /// Output only. The current serving state of the cluster.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The system-generated UID of the resource.
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
            : base("gcp:alloydb/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("gcp:alloydb/cluster:Cluster", name, state, MakeResourceOptions(options, id))
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
        /// The automated backup policy for this cluster. AutomatedBackupPolicy is disabled by default.
        /// Structure is documented below.
        /// </summary>
        [Input("automatedBackupPolicy")]
        public Input<Inputs.ClusterAutomatedBackupPolicyArgs>? AutomatedBackupPolicy { get; set; }

        /// <summary>
        /// The ID of the alloydb cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The continuous backup config for this cluster.
        /// If no policy is provided then the default policy will be used. The default policy takes one backup a day and retains backups for 14 days.
        /// Structure is documented below.
        /// </summary>
        [Input("continuousBackupConfig")]
        public Input<Inputs.ClusterContinuousBackupConfigArgs>? ContinuousBackupConfig { get; set; }

        /// <summary>
        /// User-settable and human-readable display name for the Cluster.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
        /// Structure is documented below.
        /// </summary>
        [Input("encryptionConfig")]
        public Input<Inputs.ClusterEncryptionConfigArgs>? EncryptionConfig { get; set; }

        /// <summary>
        /// For Resource freshness validation (https://google.aip.dev/154)
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Initial user to setup during cluster creation.
        /// Structure is documented below.
        /// </summary>
        [Input("initialUser")]
        public Input<Inputs.ClusterInitialUserArgs>? InitialUser { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-defined labels for the alloydb cluster.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The location where the alloydb cluster should reside.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// (Optional, Deprecated)
        /// The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
        /// "projects/{projectNumber}/global/networks/{network_id}".
        /// 
        /// &gt; **Warning:** `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// Metadata related to network configuration.
        /// Structure is documented below.
        /// </summary>
        [Input("networkConfig")]
        public Input<Inputs.ClusterNetworkConfigArgs>? NetworkConfig { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The source when restoring from a backup. Conflicts with 'restore_continuous_backup_source', both can't be set together.
        /// Structure is documented below.
        /// </summary>
        [Input("restoreBackupSource")]
        public Input<Inputs.ClusterRestoreBackupSourceArgs>? RestoreBackupSource { get; set; }

        /// <summary>
        /// The source when restoring via point in time recovery (PITR). Conflicts with 'restore_backup_source', both can't be set together.
        /// Structure is documented below.
        /// </summary>
        [Input("restoreContinuousBackupSource")]
        public Input<Inputs.ClusterRestoreContinuousBackupSourceArgs>? RestoreContinuousBackupSource { get; set; }

        public ClusterArgs()
        {
        }
        public static new ClusterArgs Empty => new ClusterArgs();
    }

    public sealed class ClusterState : global::Pulumi.ResourceArgs
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
        /// The automated backup policy for this cluster. AutomatedBackupPolicy is disabled by default.
        /// Structure is documented below.
        /// </summary>
        [Input("automatedBackupPolicy")]
        public Input<Inputs.ClusterAutomatedBackupPolicyGetArgs>? AutomatedBackupPolicy { get; set; }

        [Input("backupSources")]
        private InputList<Inputs.ClusterBackupSourceGetArgs>? _backupSources;

        /// <summary>
        /// Cluster created from backup.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ClusterBackupSourceGetArgs> BackupSources
        {
            get => _backupSources ?? (_backupSources = new InputList<Inputs.ClusterBackupSourceGetArgs>());
            set => _backupSources = value;
        }

        /// <summary>
        /// The ID of the alloydb cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The continuous backup config for this cluster.
        /// If no policy is provided then the default policy will be used. The default policy takes one backup a day and retains backups for 14 days.
        /// Structure is documented below.
        /// </summary>
        [Input("continuousBackupConfig")]
        public Input<Inputs.ClusterContinuousBackupConfigGetArgs>? ContinuousBackupConfig { get; set; }

        [Input("continuousBackupInfos")]
        private InputList<Inputs.ClusterContinuousBackupInfoGetArgs>? _continuousBackupInfos;

        /// <summary>
        /// ContinuousBackupInfo describes the continuous backup properties of a cluster.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ClusterContinuousBackupInfoGetArgs> ContinuousBackupInfos
        {
            get => _continuousBackupInfos ?? (_continuousBackupInfos = new InputList<Inputs.ClusterContinuousBackupInfoGetArgs>());
            set => _continuousBackupInfos = value;
        }

        /// <summary>
        /// The database engine major version. This is an output-only field and it's populated at the Cluster creation time. This field cannot be changed after cluster creation.
        /// </summary>
        [Input("databaseVersion")]
        public Input<string>? DatabaseVersion { get; set; }

        /// <summary>
        /// User-settable and human-readable display name for the Cluster.
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

        /// <summary>
        /// EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
        /// Structure is documented below.
        /// </summary>
        [Input("encryptionConfig")]
        public Input<Inputs.ClusterEncryptionConfigGetArgs>? EncryptionConfig { get; set; }

        [Input("encryptionInfos")]
        private InputList<Inputs.ClusterEncryptionInfoGetArgs>? _encryptionInfos;

        /// <summary>
        /// (Output)
        /// Output only. The encryption information for the WALs and backups required for ContinuousBackup.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ClusterEncryptionInfoGetArgs> EncryptionInfos
        {
            get => _encryptionInfos ?? (_encryptionInfos = new InputList<Inputs.ClusterEncryptionInfoGetArgs>());
            set => _encryptionInfos = value;
        }

        /// <summary>
        /// For Resource freshness validation (https://google.aip.dev/154)
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Initial user to setup during cluster creation.
        /// Structure is documented below.
        /// </summary>
        [Input("initialUser")]
        public Input<Inputs.ClusterInitialUserGetArgs>? InitialUser { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-defined labels for the alloydb cluster.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The location where the alloydb cluster should reside.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("migrationSources")]
        private InputList<Inputs.ClusterMigrationSourceGetArgs>? _migrationSources;

        /// <summary>
        /// Cluster created via DMS migration.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ClusterMigrationSourceGetArgs> MigrationSources
        {
            get => _migrationSources ?? (_migrationSources = new InputList<Inputs.ClusterMigrationSourceGetArgs>());
            set => _migrationSources = value;
        }

        /// <summary>
        /// The name of the cluster resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (Optional, Deprecated)
        /// The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
        /// "projects/{projectNumber}/global/networks/{network_id}".
        /// 
        /// &gt; **Warning:** `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// Metadata related to network configuration.
        /// Structure is documented below.
        /// </summary>
        [Input("networkConfig")]
        public Input<Inputs.ClusterNetworkConfigGetArgs>? NetworkConfig { get; set; }

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
        /// Output only. Reconciling (https://google.aip.dev/128#reconciliation).
        /// Set to true if the current state of Cluster does not match the user's intended state, and the service is actively updating the resource to reconcile them.
        /// This can happen due to user-triggered updates or system actions like failover or maintenance.
        /// </summary>
        [Input("reconciling")]
        public Input<bool>? Reconciling { get; set; }

        /// <summary>
        /// The source when restoring from a backup. Conflicts with 'restore_continuous_backup_source', both can't be set together.
        /// Structure is documented below.
        /// </summary>
        [Input("restoreBackupSource")]
        public Input<Inputs.ClusterRestoreBackupSourceGetArgs>? RestoreBackupSource { get; set; }

        /// <summary>
        /// The source when restoring via point in time recovery (PITR). Conflicts with 'restore_backup_source', both can't be set together.
        /// Structure is documented below.
        /// </summary>
        [Input("restoreContinuousBackupSource")]
        public Input<Inputs.ClusterRestoreContinuousBackupSourceGetArgs>? RestoreContinuousBackupSource { get; set; }

        /// <summary>
        /// Output only. The current serving state of the cluster.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The system-generated UID of the resource.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        public ClusterState()
        {
        }
        public static new ClusterState Empty => new ClusterState();
    }
}
