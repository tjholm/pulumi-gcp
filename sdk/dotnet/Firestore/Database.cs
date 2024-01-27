// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Firestore
{
    /// <summary>
    /// A Cloud Firestore Database.
    /// 
    /// If you wish to use Firestore with App Engine, use the
    /// `gcp.appengine.Application`
    /// resource instead. If you were previously using the `gcp.appengine.Application` resource exclusively for managing a Firestore database
    /// and would like to use the `gcp.firestore.Database` resource instead, please follow the instructions
    /// [here](https://cloud.google.com/firestore/docs/app-engine-requirement).
    /// 
    /// To get more information about Database, see:
    /// 
    /// * [API documentation](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/firestore/docs/)
    /// 
    /// ## Example Usage
    /// ### Firestore Default Database
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var database = new Gcp.Firestore.Database("database", new()
    ///     {
    ///         DeleteProtectionState = "DELETE_PROTECTION_ENABLED",
    ///         DeletionPolicy = "DELETE",
    ///         LocationId = "nam5",
    ///         Project = "my-project-name",
    ///         Type = "FIRESTORE_NATIVE",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Firestore Database
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var database = new Gcp.Firestore.Database("database", new()
    ///     {
    ///         AppEngineIntegrationMode = "DISABLED",
    ///         ConcurrencyMode = "OPTIMISTIC",
    ///         DeleteProtectionState = "DELETE_PROTECTION_ENABLED",
    ///         DeletionPolicy = "DELETE",
    ///         LocationId = "nam5",
    ///         PointInTimeRecoveryEnablement = "POINT_IN_TIME_RECOVERY_ENABLED",
    ///         Project = "my-project-name",
    ///         Type = "FIRESTORE_NATIVE",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Firestore Default Database In Datastore Mode
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var datastoreModeDatabase = new Gcp.Firestore.Database("datastoreModeDatabase", new()
    ///     {
    ///         DeleteProtectionState = "DELETE_PROTECTION_ENABLED",
    ///         DeletionPolicy = "DELETE",
    ///         LocationId = "nam5",
    ///         Project = "my-project-name",
    ///         Type = "DATASTORE_MODE",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Firestore Database In Datastore Mode
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var datastoreModeDatabase = new Gcp.Firestore.Database("datastoreModeDatabase", new()
    ///     {
    ///         AppEngineIntegrationMode = "DISABLED",
    ///         ConcurrencyMode = "OPTIMISTIC",
    ///         DeleteProtectionState = "DELETE_PROTECTION_ENABLED",
    ///         DeletionPolicy = "DELETE",
    ///         LocationId = "nam5",
    ///         PointInTimeRecoveryEnablement = "POINT_IN_TIME_RECOVERY_ENABLED",
    ///         Project = "my-project-name",
    ///         Type = "DATASTORE_MODE",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Database can be imported using any of these accepted formats* `projects/{{project}}/databases/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Database can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:firestore/database:Database default projects/{{project}}/databases/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:firestore/database:Database default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:firestore/database:Database default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:firestore/database:Database")]
    public partial class Database : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The App Engine integration mode to use for this database.
        /// Possible values are: `ENABLED`, `DISABLED`.
        /// </summary>
        [Output("appEngineIntegrationMode")]
        public Output<string> AppEngineIntegrationMode { get; private set; } = null!;

        /// <summary>
        /// The concurrency control mode to use for this database.
        /// Possible values are: `OPTIMISTIC`, `PESSIMISTIC`, `OPTIMISTIC_WITH_ENTITY_GROUPS`.
        /// </summary>
        [Output("concurrencyMode")]
        public Output<string> ConcurrencyMode { get; private set; } = null!;

        /// <summary>
        /// Output only. The timestamp at which this database was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// State of delete protection for the database. When delete protection is enabled, this database cannot be deleted. The
        /// default value is 'DELETE_PROTECTION_STATE_UNSPECIFIED', which is currently equivalent to 'DELETE_PROTECTION_DISABLED'.
        /// **Note:** Additionally, to delete this database using 'terraform destroy', 'deletion_policy' must be set to 'DELETE'.
        /// Possible values: ["DELETE_PROTECTION_STATE_UNSPECIFIED", "DELETE_PROTECTION_ENABLED", "DELETE_PROTECTION_DISABLED"]
        /// </summary>
        [Output("deleteProtectionState")]
        public Output<string> DeleteProtectionState { get; private set; } = null!;

        /// <summary>
        /// Deletion behavior for this database. If the deletion policy is 'ABANDON', the database will be removed from Terraform
        /// state but not deleted from Google Cloud upon destruction. If the deletion policy is 'DELETE', the database will both be
        /// removed from Terraform state and deleted from Google Cloud upon destruction. The default value is 'ABANDON'. See also
        /// 'delete_protection'.
        /// </summary>
        [Output("deletionPolicy")]
        public Output<string?> DeletionPolicy { get; private set; } = null!;

        /// <summary>
        /// Output only. The earliest timestamp at which older versions of the data can be read from the database. See versionRetentionPeriod above; this field is populated with now - versionRetentionPeriod.
        /// This value is continuously updated, and becomes stale the moment it is queried. If you are using this value to recover data, make sure to account for the time from the moment when the value is queried to the moment when you initiate the recovery.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("earliestVersionTime")]
        public Output<string> EarliestVersionTime { get; private set; } = null!;

        /// <summary>
        /// Output only. This checksum is computed by the server based on the value of other fields,
        /// and may be sent on update and delete requests to ensure the client has an
        /// up-to-date value before proceeding.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Output only. The keyPrefix for this database.
        /// This keyPrefix is used, in combination with the project id ("~") to construct the application id
        /// that is returned from the Cloud Datastore APIs in Google App Engine first generation runtimes.
        /// This value may be empty in which case the appid to use for URL-encoded keys is the project_id (eg: foo instead of v~foo).
        /// </summary>
        [Output("keyPrefix")]
        public Output<string> KeyPrefix { get; private set; } = null!;

        /// <summary>
        /// The location of the database. Available locations are listed at
        /// https://cloud.google.com/firestore/docs/locations.
        /// </summary>
        [Output("locationId")]
        public Output<string> LocationId { get; private set; } = null!;

        /// <summary>
        /// The ID to use for the database, which will become the final
        /// component of the database's resource name. This value should be 4-63
        /// characters. Valid characters are /[a-z][0-9]-/ with first character
        /// a letter and the last a letter or a number. Must not be
        /// UUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/.
        /// "(default)" database id is also valid.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Whether to enable the PITR feature on this database.
        /// If `POINT_IN_TIME_RECOVERY_ENABLED` is selected, reads are supported on selected versions of the data from within the past 7 days.
        /// versionRetentionPeriod and earliestVersionTime can be used to determine the supported versions. These include reads against any timestamp within the past hour
        /// and reads against 1-minute snapshots beyond 1 hour and within 7 days.
        /// If `POINT_IN_TIME_RECOVERY_DISABLED` is selected, reads are supported on any version of the data from within the past 1 hour.
        /// Default value is `POINT_IN_TIME_RECOVERY_DISABLED`.
        /// Possible values are: `POINT_IN_TIME_RECOVERY_ENABLED`, `POINT_IN_TIME_RECOVERY_DISABLED`.
        /// </summary>
        [Output("pointInTimeRecoveryEnablement")]
        public Output<string?> PointInTimeRecoveryEnablement { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The type of the database.
        /// See https://cloud.google.com/datastore/docs/firestore-or-datastore
        /// for information about how to choose.
        /// Possible values are: `FIRESTORE_NATIVE`, `DATASTORE_MODE`.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Output only. The system-generated UUID4 for this Database.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// Output only. The timestamp at which this database was most recently updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Output only. The period during which past versions of data are retained in the database.
        /// Any read or query can specify a readTime within this window, and will read the state of the database at that time.
        /// If the PITR feature is enabled, the retention period is 7 days. Otherwise, the retention period is 1 hour.
        /// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
        /// </summary>
        [Output("versionRetentionPeriod")]
        public Output<string> VersionRetentionPeriod { get; private set; } = null!;


        /// <summary>
        /// Create a Database resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Database(string name, DatabaseArgs args, CustomResourceOptions? options = null)
            : base("gcp:firestore/database:Database", name, args ?? new DatabaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Database(string name, Input<string> id, DatabaseState? state = null, CustomResourceOptions? options = null)
            : base("gcp:firestore/database:Database", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Database resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Database Get(string name, Input<string> id, DatabaseState? state = null, CustomResourceOptions? options = null)
        {
            return new Database(name, id, state, options);
        }
    }

    public sealed class DatabaseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The App Engine integration mode to use for this database.
        /// Possible values are: `ENABLED`, `DISABLED`.
        /// </summary>
        [Input("appEngineIntegrationMode")]
        public Input<string>? AppEngineIntegrationMode { get; set; }

        /// <summary>
        /// The concurrency control mode to use for this database.
        /// Possible values are: `OPTIMISTIC`, `PESSIMISTIC`, `OPTIMISTIC_WITH_ENTITY_GROUPS`.
        /// </summary>
        [Input("concurrencyMode")]
        public Input<string>? ConcurrencyMode { get; set; }

        /// <summary>
        /// State of delete protection for the database. When delete protection is enabled, this database cannot be deleted. The
        /// default value is 'DELETE_PROTECTION_STATE_UNSPECIFIED', which is currently equivalent to 'DELETE_PROTECTION_DISABLED'.
        /// **Note:** Additionally, to delete this database using 'terraform destroy', 'deletion_policy' must be set to 'DELETE'.
        /// Possible values: ["DELETE_PROTECTION_STATE_UNSPECIFIED", "DELETE_PROTECTION_ENABLED", "DELETE_PROTECTION_DISABLED"]
        /// </summary>
        [Input("deleteProtectionState")]
        public Input<string>? DeleteProtectionState { get; set; }

        /// <summary>
        /// Deletion behavior for this database. If the deletion policy is 'ABANDON', the database will be removed from Terraform
        /// state but not deleted from Google Cloud upon destruction. If the deletion policy is 'DELETE', the database will both be
        /// removed from Terraform state and deleted from Google Cloud upon destruction. The default value is 'ABANDON'. See also
        /// 'delete_protection'.
        /// </summary>
        [Input("deletionPolicy")]
        public Input<string>? DeletionPolicy { get; set; }

        /// <summary>
        /// The location of the database. Available locations are listed at
        /// https://cloud.google.com/firestore/docs/locations.
        /// </summary>
        [Input("locationId", required: true)]
        public Input<string> LocationId { get; set; } = null!;

        /// <summary>
        /// The ID to use for the database, which will become the final
        /// component of the database's resource name. This value should be 4-63
        /// characters. Valid characters are /[a-z][0-9]-/ with first character
        /// a letter and the last a letter or a number. Must not be
        /// UUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/.
        /// "(default)" database id is also valid.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether to enable the PITR feature on this database.
        /// If `POINT_IN_TIME_RECOVERY_ENABLED` is selected, reads are supported on selected versions of the data from within the past 7 days.
        /// versionRetentionPeriod and earliestVersionTime can be used to determine the supported versions. These include reads against any timestamp within the past hour
        /// and reads against 1-minute snapshots beyond 1 hour and within 7 days.
        /// If `POINT_IN_TIME_RECOVERY_DISABLED` is selected, reads are supported on any version of the data from within the past 1 hour.
        /// Default value is `POINT_IN_TIME_RECOVERY_DISABLED`.
        /// Possible values are: `POINT_IN_TIME_RECOVERY_ENABLED`, `POINT_IN_TIME_RECOVERY_DISABLED`.
        /// </summary>
        [Input("pointInTimeRecoveryEnablement")]
        public Input<string>? PointInTimeRecoveryEnablement { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The type of the database.
        /// See https://cloud.google.com/datastore/docs/firestore-or-datastore
        /// for information about how to choose.
        /// Possible values are: `FIRESTORE_NATIVE`, `DATASTORE_MODE`.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DatabaseArgs()
        {
        }
        public static new DatabaseArgs Empty => new DatabaseArgs();
    }

    public sealed class DatabaseState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The App Engine integration mode to use for this database.
        /// Possible values are: `ENABLED`, `DISABLED`.
        /// </summary>
        [Input("appEngineIntegrationMode")]
        public Input<string>? AppEngineIntegrationMode { get; set; }

        /// <summary>
        /// The concurrency control mode to use for this database.
        /// Possible values are: `OPTIMISTIC`, `PESSIMISTIC`, `OPTIMISTIC_WITH_ENTITY_GROUPS`.
        /// </summary>
        [Input("concurrencyMode")]
        public Input<string>? ConcurrencyMode { get; set; }

        /// <summary>
        /// Output only. The timestamp at which this database was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// State of delete protection for the database. When delete protection is enabled, this database cannot be deleted. The
        /// default value is 'DELETE_PROTECTION_STATE_UNSPECIFIED', which is currently equivalent to 'DELETE_PROTECTION_DISABLED'.
        /// **Note:** Additionally, to delete this database using 'terraform destroy', 'deletion_policy' must be set to 'DELETE'.
        /// Possible values: ["DELETE_PROTECTION_STATE_UNSPECIFIED", "DELETE_PROTECTION_ENABLED", "DELETE_PROTECTION_DISABLED"]
        /// </summary>
        [Input("deleteProtectionState")]
        public Input<string>? DeleteProtectionState { get; set; }

        /// <summary>
        /// Deletion behavior for this database. If the deletion policy is 'ABANDON', the database will be removed from Terraform
        /// state but not deleted from Google Cloud upon destruction. If the deletion policy is 'DELETE', the database will both be
        /// removed from Terraform state and deleted from Google Cloud upon destruction. The default value is 'ABANDON'. See also
        /// 'delete_protection'.
        /// </summary>
        [Input("deletionPolicy")]
        public Input<string>? DeletionPolicy { get; set; }

        /// <summary>
        /// Output only. The earliest timestamp at which older versions of the data can be read from the database. See versionRetentionPeriod above; this field is populated with now - versionRetentionPeriod.
        /// This value is continuously updated, and becomes stale the moment it is queried. If you are using this value to recover data, make sure to account for the time from the moment when the value is queried to the moment when you initiate the recovery.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("earliestVersionTime")]
        public Input<string>? EarliestVersionTime { get; set; }

        /// <summary>
        /// Output only. This checksum is computed by the server based on the value of other fields,
        /// and may be sent on update and delete requests to ensure the client has an
        /// up-to-date value before proceeding.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Output only. The keyPrefix for this database.
        /// This keyPrefix is used, in combination with the project id ("~") to construct the application id
        /// that is returned from the Cloud Datastore APIs in Google App Engine first generation runtimes.
        /// This value may be empty in which case the appid to use for URL-encoded keys is the project_id (eg: foo instead of v~foo).
        /// </summary>
        [Input("keyPrefix")]
        public Input<string>? KeyPrefix { get; set; }

        /// <summary>
        /// The location of the database. Available locations are listed at
        /// https://cloud.google.com/firestore/docs/locations.
        /// </summary>
        [Input("locationId")]
        public Input<string>? LocationId { get; set; }

        /// <summary>
        /// The ID to use for the database, which will become the final
        /// component of the database's resource name. This value should be 4-63
        /// characters. Valid characters are /[a-z][0-9]-/ with first character
        /// a letter and the last a letter or a number. Must not be
        /// UUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/.
        /// "(default)" database id is also valid.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether to enable the PITR feature on this database.
        /// If `POINT_IN_TIME_RECOVERY_ENABLED` is selected, reads are supported on selected versions of the data from within the past 7 days.
        /// versionRetentionPeriod and earliestVersionTime can be used to determine the supported versions. These include reads against any timestamp within the past hour
        /// and reads against 1-minute snapshots beyond 1 hour and within 7 days.
        /// If `POINT_IN_TIME_RECOVERY_DISABLED` is selected, reads are supported on any version of the data from within the past 1 hour.
        /// Default value is `POINT_IN_TIME_RECOVERY_DISABLED`.
        /// Possible values are: `POINT_IN_TIME_RECOVERY_ENABLED`, `POINT_IN_TIME_RECOVERY_DISABLED`.
        /// </summary>
        [Input("pointInTimeRecoveryEnablement")]
        public Input<string>? PointInTimeRecoveryEnablement { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The type of the database.
        /// See https://cloud.google.com/datastore/docs/firestore-or-datastore
        /// for information about how to choose.
        /// Possible values are: `FIRESTORE_NATIVE`, `DATASTORE_MODE`.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Output only. The system-generated UUID4 for this Database.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// Output only. The timestamp at which this database was most recently updated.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// Output only. The period during which past versions of data are retained in the database.
        /// Any read or query can specify a readTime within this window, and will read the state of the database at that time.
        /// If the PITR feature is enabled, the retention period is 7 days. Otherwise, the retention period is 1 hour.
        /// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
        /// </summary>
        [Input("versionRetentionPeriod")]
        public Input<string>? VersionRetentionPeriod { get; set; }

        public DatabaseState()
        {
        }
        public static new DatabaseState Empty => new DatabaseState();
    }
}
