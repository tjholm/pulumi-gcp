// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeBackup
{
    /// <summary>
    /// Represents a Backup Plan instance.
    /// 
    /// To get more information about BackupPlan, see:
    /// 
    /// * [API documentation](https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke/reference/rest/v1/projects.locations.backupPlans)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke)
    /// 
    /// ## Example Usage
    /// ### Gkebackup Backupplan Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var primary = new Gcp.Container.Cluster("primary", new()
    ///     {
    ///         Location = "us-central1",
    ///         InitialNodeCount = 1,
    ///         WorkloadIdentityConfig = new Gcp.Container.Inputs.ClusterWorkloadIdentityConfigArgs
    ///         {
    ///             WorkloadPool = "my-project-name.svc.id.goog",
    ///         },
    ///         AddonsConfig = new Gcp.Container.Inputs.ClusterAddonsConfigArgs
    ///         {
    ///             GkeBackupAgentConfig = new Gcp.Container.Inputs.ClusterAddonsConfigGkeBackupAgentConfigArgs
    ///             {
    ///                 Enabled = true,
    ///             },
    ///         },
    ///         DeletionProtection = true,
    ///     });
    /// 
    ///     var basic = new Gcp.GkeBackup.BackupPlan("basic", new()
    ///     {
    ///         Cluster = primary.Id,
    ///         Location = "us-central1",
    ///         BackupConfig = new Gcp.GkeBackup.Inputs.BackupPlanBackupConfigArgs
    ///         {
    ///             IncludeVolumeData = true,
    ///             IncludeSecrets = true,
    ///             AllNamespaces = true,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Gkebackup Backupplan Autopilot
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var primary = new Gcp.Container.Cluster("primary", new()
    ///     {
    ///         Location = "us-central1",
    ///         EnableAutopilot = true,
    ///         IpAllocationPolicy = null,
    ///         ReleaseChannel = new Gcp.Container.Inputs.ClusterReleaseChannelArgs
    ///         {
    ///             Channel = "RAPID",
    ///         },
    ///         AddonsConfig = new Gcp.Container.Inputs.ClusterAddonsConfigArgs
    ///         {
    ///             GkeBackupAgentConfig = new Gcp.Container.Inputs.ClusterAddonsConfigGkeBackupAgentConfigArgs
    ///             {
    ///                 Enabled = true,
    ///             },
    ///         },
    ///         DeletionProtection = true,
    ///     });
    /// 
    ///     var autopilot = new Gcp.GkeBackup.BackupPlan("autopilot", new()
    ///     {
    ///         Cluster = primary.Id,
    ///         Location = "us-central1",
    ///         BackupConfig = new Gcp.GkeBackup.Inputs.BackupPlanBackupConfigArgs
    ///         {
    ///             IncludeVolumeData = true,
    ///             IncludeSecrets = true,
    ///             AllNamespaces = true,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Gkebackup Backupplan Cmek
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var primary = new Gcp.Container.Cluster("primary", new()
    ///     {
    ///         Location = "us-central1",
    ///         InitialNodeCount = 1,
    ///         WorkloadIdentityConfig = new Gcp.Container.Inputs.ClusterWorkloadIdentityConfigArgs
    ///         {
    ///             WorkloadPool = "my-project-name.svc.id.goog",
    ///         },
    ///         AddonsConfig = new Gcp.Container.Inputs.ClusterAddonsConfigArgs
    ///         {
    ///             GkeBackupAgentConfig = new Gcp.Container.Inputs.ClusterAddonsConfigGkeBackupAgentConfigArgs
    ///             {
    ///                 Enabled = true,
    ///             },
    ///         },
    ///         DeletionProtection = true,
    ///     });
    /// 
    ///     var keyRing = new Gcp.Kms.KeyRing("keyRing", new()
    ///     {
    ///         Location = "us-central1",
    ///     });
    /// 
    ///     var cryptoKey = new Gcp.Kms.CryptoKey("cryptoKey", new()
    ///     {
    ///         KeyRing = keyRing.Id,
    ///     });
    /// 
    ///     var cmek = new Gcp.GkeBackup.BackupPlan("cmek", new()
    ///     {
    ///         Cluster = primary.Id,
    ///         Location = "us-central1",
    ///         BackupConfig = new Gcp.GkeBackup.Inputs.BackupPlanBackupConfigArgs
    ///         {
    ///             IncludeVolumeData = true,
    ///             IncludeSecrets = true,
    ///             SelectedNamespaces = new Gcp.GkeBackup.Inputs.BackupPlanBackupConfigSelectedNamespacesArgs
    ///             {
    ///                 Namespaces = new[]
    ///                 {
    ///                     "default",
    ///                     "test",
    ///                 },
    ///             },
    ///             EncryptionKey = new Gcp.GkeBackup.Inputs.BackupPlanBackupConfigEncryptionKeyArgs
    ///             {
    ///                 GcpKmsEncryptionKey = cryptoKey.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Gkebackup Backupplan Full
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var primary = new Gcp.Container.Cluster("primary", new()
    ///     {
    ///         Location = "us-central1",
    ///         InitialNodeCount = 1,
    ///         WorkloadIdentityConfig = new Gcp.Container.Inputs.ClusterWorkloadIdentityConfigArgs
    ///         {
    ///             WorkloadPool = "my-project-name.svc.id.goog",
    ///         },
    ///         AddonsConfig = new Gcp.Container.Inputs.ClusterAddonsConfigArgs
    ///         {
    ///             GkeBackupAgentConfig = new Gcp.Container.Inputs.ClusterAddonsConfigGkeBackupAgentConfigArgs
    ///             {
    ///                 Enabled = true,
    ///             },
    ///         },
    ///         DeletionProtection = true,
    ///     });
    /// 
    ///     var full = new Gcp.GkeBackup.BackupPlan("full", new()
    ///     {
    ///         Cluster = primary.Id,
    ///         Location = "us-central1",
    ///         RetentionPolicy = new Gcp.GkeBackup.Inputs.BackupPlanRetentionPolicyArgs
    ///         {
    ///             BackupDeleteLockDays = 30,
    ///             BackupRetainDays = 180,
    ///         },
    ///         BackupSchedule = new Gcp.GkeBackup.Inputs.BackupPlanBackupScheduleArgs
    ///         {
    ///             CronSchedule = "0 9 * * 1",
    ///         },
    ///         BackupConfig = new Gcp.GkeBackup.Inputs.BackupPlanBackupConfigArgs
    ///         {
    ///             IncludeVolumeData = true,
    ///             IncludeSecrets = true,
    ///             SelectedApplications = new Gcp.GkeBackup.Inputs.BackupPlanBackupConfigSelectedApplicationsArgs
    ///             {
    ///                 NamespacedNames = new[]
    ///                 {
    ///                     new Gcp.GkeBackup.Inputs.BackupPlanBackupConfigSelectedApplicationsNamespacedNameArgs
    ///                     {
    ///                         Name = "app1",
    ///                         Namespace = "ns1",
    ///                     },
    ///                     new Gcp.GkeBackup.Inputs.BackupPlanBackupConfigSelectedApplicationsNamespacedNameArgs
    ///                     {
    ///                         Name = "app2",
    ///                         Namespace = "ns2",
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
    /// BackupPlan can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:gkebackup/backupPlan:BackupPlan default projects/{{project}}/locations/{{location}}/backupPlans/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:gkebackup/backupPlan:BackupPlan default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:gkebackup/backupPlan:BackupPlan default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:gkebackup/backupPlan:BackupPlan")]
    public partial class BackupPlan : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Defines the configuration of Backups created via this BackupPlan.
        /// Structure is documented below.
        /// </summary>
        [Output("backupConfig")]
        public Output<Outputs.BackupPlanBackupConfig?> BackupConfig { get; private set; } = null!;

        /// <summary>
        /// Defines a schedule for automatic Backup creation via this BackupPlan.
        /// Structure is documented below.
        /// </summary>
        [Output("backupSchedule")]
        public Output<Outputs.BackupPlanBackupSchedule?> BackupSchedule { get; private set; } = null!;

        /// <summary>
        /// The source cluster from which Backups will be created via this BackupPlan.
        /// </summary>
        [Output("cluster")]
        public Output<string> Cluster { get; private set; } = null!;

        /// <summary>
        /// This flag indicates whether this BackupPlan has been deactivated.
        /// Setting this field to True locks the BackupPlan such that no further updates will be allowed
        /// (except deletes), including the deactivated field itself. It also prevents any new Backups
        /// from being created via this BackupPlan (including scheduled Backups).
        /// </summary>
        [Output("deactivated")]
        public Output<bool> Deactivated { get; private set; } = null!;

        /// <summary>
        /// User specified descriptive string for this BackupPlan.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// etag is used for optimistic concurrency control as a way to help prevent simultaneous
        /// updates of a backup plan from overwriting each other. It is strongly suggested that
        /// systems make use of the 'etag' in the read-modify-write cycle to perform BackupPlan updates
        /// in order to avoid race conditions: An etag is returned in the response to backupPlans.get,
        /// and systems are expected to put that etag in the request to backupPlans.patch or
        /// backupPlans.delete to ensure that their change will be applied to the same version of the resource.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Description: A set of custom labels supplied by the user.
        /// A list of key-&gt;value pairs.
        /// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The region of the Backup Plan.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The full name of the BackupPlan Resource.
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
        /// The number of Kubernetes Pods backed up in the last successful Backup created via this BackupPlan.
        /// </summary>
        [Output("protectedPodCount")]
        public Output<int> ProtectedPodCount { get; private set; } = null!;

        /// <summary>
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        [Output("pulumiLabels")]
        public Output<ImmutableDictionary<string, string>> PulumiLabels { get; private set; } = null!;

        /// <summary>
        /// RetentionPolicy governs lifecycle of Backups created under this plan.
        /// Structure is documented below.
        /// </summary>
        [Output("retentionPolicy")]
        public Output<Outputs.BackupPlanRetentionPolicy?> RetentionPolicy { get; private set; } = null!;

        /// <summary>
        /// The State of the BackupPlan.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Detailed description of why BackupPlan is in its current state.
        /// </summary>
        [Output("stateReason")]
        public Output<string> StateReason { get; private set; } = null!;

        /// <summary>
        /// Server generated, unique identifier of UUID format.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;


        /// <summary>
        /// Create a BackupPlan resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackupPlan(string name, BackupPlanArgs args, CustomResourceOptions? options = null)
            : base("gcp:gkebackup/backupPlan:BackupPlan", name, args ?? new BackupPlanArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackupPlan(string name, Input<string> id, BackupPlanState? state = null, CustomResourceOptions? options = null)
            : base("gcp:gkebackup/backupPlan:BackupPlan", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BackupPlan resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackupPlan Get(string name, Input<string> id, BackupPlanState? state = null, CustomResourceOptions? options = null)
        {
            return new BackupPlan(name, id, state, options);
        }
    }

    public sealed class BackupPlanArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines the configuration of Backups created via this BackupPlan.
        /// Structure is documented below.
        /// </summary>
        [Input("backupConfig")]
        public Input<Inputs.BackupPlanBackupConfigArgs>? BackupConfig { get; set; }

        /// <summary>
        /// Defines a schedule for automatic Backup creation via this BackupPlan.
        /// Structure is documented below.
        /// </summary>
        [Input("backupSchedule")]
        public Input<Inputs.BackupPlanBackupScheduleArgs>? BackupSchedule { get; set; }

        /// <summary>
        /// The source cluster from which Backups will be created via this BackupPlan.
        /// </summary>
        [Input("cluster", required: true)]
        public Input<string> Cluster { get; set; } = null!;

        /// <summary>
        /// This flag indicates whether this BackupPlan has been deactivated.
        /// Setting this field to True locks the BackupPlan such that no further updates will be allowed
        /// (except deletes), including the deactivated field itself. It also prevents any new Backups
        /// from being created via this BackupPlan (including scheduled Backups).
        /// </summary>
        [Input("deactivated")]
        public Input<bool>? Deactivated { get; set; }

        /// <summary>
        /// User specified descriptive string for this BackupPlan.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Description: A set of custom labels supplied by the user.
        /// A list of key-&gt;value pairs.
        /// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
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
        /// The region of the Backup Plan.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The full name of the BackupPlan Resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// RetentionPolicy governs lifecycle of Backups created under this plan.
        /// Structure is documented below.
        /// </summary>
        [Input("retentionPolicy")]
        public Input<Inputs.BackupPlanRetentionPolicyArgs>? RetentionPolicy { get; set; }

        public BackupPlanArgs()
        {
        }
        public static new BackupPlanArgs Empty => new BackupPlanArgs();
    }

    public sealed class BackupPlanState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines the configuration of Backups created via this BackupPlan.
        /// Structure is documented below.
        /// </summary>
        [Input("backupConfig")]
        public Input<Inputs.BackupPlanBackupConfigGetArgs>? BackupConfig { get; set; }

        /// <summary>
        /// Defines a schedule for automatic Backup creation via this BackupPlan.
        /// Structure is documented below.
        /// </summary>
        [Input("backupSchedule")]
        public Input<Inputs.BackupPlanBackupScheduleGetArgs>? BackupSchedule { get; set; }

        /// <summary>
        /// The source cluster from which Backups will be created via this BackupPlan.
        /// </summary>
        [Input("cluster")]
        public Input<string>? Cluster { get; set; }

        /// <summary>
        /// This flag indicates whether this BackupPlan has been deactivated.
        /// Setting this field to True locks the BackupPlan such that no further updates will be allowed
        /// (except deletes), including the deactivated field itself. It also prevents any new Backups
        /// from being created via this BackupPlan (including scheduled Backups).
        /// </summary>
        [Input("deactivated")]
        public Input<bool>? Deactivated { get; set; }

        /// <summary>
        /// User specified descriptive string for this BackupPlan.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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
        /// etag is used for optimistic concurrency control as a way to help prevent simultaneous
        /// updates of a backup plan from overwriting each other. It is strongly suggested that
        /// systems make use of the 'etag' in the read-modify-write cycle to perform BackupPlan updates
        /// in order to avoid race conditions: An etag is returned in the response to backupPlans.get,
        /// and systems are expected to put that etag in the request to backupPlans.patch or
        /// backupPlans.delete to ensure that their change will be applied to the same version of the resource.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Description: A set of custom labels supplied by the user.
        /// A list of key-&gt;value pairs.
        /// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
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
        /// The region of the Backup Plan.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The full name of the BackupPlan Resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The number of Kubernetes Pods backed up in the last successful Backup created via this BackupPlan.
        /// </summary>
        [Input("protectedPodCount")]
        public Input<int>? ProtectedPodCount { get; set; }

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
        /// RetentionPolicy governs lifecycle of Backups created under this plan.
        /// Structure is documented below.
        /// </summary>
        [Input("retentionPolicy")]
        public Input<Inputs.BackupPlanRetentionPolicyGetArgs>? RetentionPolicy { get; set; }

        /// <summary>
        /// The State of the BackupPlan.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Detailed description of why BackupPlan is in its current state.
        /// </summary>
        [Input("stateReason")]
        public Input<string>? StateReason { get; set; }

        /// <summary>
        /// Server generated, unique identifier of UUID format.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        public BackupPlanState()
        {
        }
        public static new BackupPlanState Empty => new BackupPlanState();
    }
}
