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
    /// A backup schedule for a Cloud Firestore Database.
    /// This resource is owned by the database it is backing up, and is deleted along with the database.
    /// The actual backups are not though.
    /// 
    /// To get more information about BackupSchedule, see:
    /// 
    /// * [API documentation](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.backupSchedules)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/firestore/docs/backups)
    /// 
    /// &gt; **Warning:** This resource creates a Firestore Backup Schedule on a project that already has
    /// a Firestore database.
    /// This resource is owned by the database it is backing up, and is deleted along
    /// with the database. The actual backups are not though.
    /// 
    /// ## Example Usage
    /// ### Firestore Backup Schedule Daily
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
    ///         Project = "my-project-name",
    ///         LocationId = "nam5",
    ///         Type = "FIRESTORE_NATIVE",
    ///         DeleteProtectionState = "DELETE_PROTECTION_ENABLED",
    ///         DeletionPolicy = "DELETE",
    ///     });
    /// 
    ///     var daily_backup = new Gcp.Firestore.BackupSchedule("daily-backup", new()
    ///     {
    ///         Project = "my-project-name",
    ///         Database = database.Name,
    ///         Retention = "604800s",
    ///         DailyRecurrence = null,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Firestore Backup Schedule Weekly
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
    ///         Project = "my-project-name",
    ///         LocationId = "nam5",
    ///         Type = "FIRESTORE_NATIVE",
    ///         DeleteProtectionState = "DELETE_PROTECTION_ENABLED",
    ///         DeletionPolicy = "DELETE",
    ///     });
    /// 
    ///     var weekly_backup = new Gcp.Firestore.BackupSchedule("weekly-backup", new()
    ///     {
    ///         Project = "my-project-name",
    ///         Database = database.Name,
    ///         Retention = "8467200s",
    ///         WeeklyRecurrence = new Gcp.Firestore.Inputs.BackupScheduleWeeklyRecurrenceArgs
    ///         {
    ///             Day = "SUNDAY",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// BackupSchedule can be imported using any of these accepted formats* `projects/{{project}}/databases/{{database}}/backupSchedules/{{name}}` * `{{project}}/{{database}}/{{name}}` * `{{database}}/{{name}}` When using the `pulumi import` command, BackupSchedule can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:firestore/backupSchedule:BackupSchedule default projects/{{project}}/databases/{{database}}/backupSchedules/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:firestore/backupSchedule:BackupSchedule default {{project}}/{{database}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:firestore/backupSchedule:BackupSchedule default {{database}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:firestore/backupSchedule:BackupSchedule")]
    public partial class BackupSchedule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// For a schedule that runs daily at a specified time.
        /// </summary>
        [Output("dailyRecurrence")]
        public Output<Outputs.BackupScheduleDailyRecurrence?> DailyRecurrence { get; private set; } = null!;

        /// <summary>
        /// The Firestore database id. Defaults to `"(default)"`.
        /// </summary>
        [Output("database")]
        public Output<string?> Database { get; private set; } = null!;

        /// <summary>
        /// The unique backup schedule identifier across all locations and databases for the given project. Format:
        /// `projects/{{project}}/databases/{{database}}/backupSchedules/{{backupSchedule}}`
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
        /// At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.
        /// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
        /// For a daily backup recurrence, set this to a value up to 7 days. If you set a weekly backup recurrence, set this to a value up to 14 weeks.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("retention")]
        public Output<string> Retention { get; private set; } = null!;

        /// <summary>
        /// For a schedule that runs weekly on a specific day and time.
        /// Structure is documented below.
        /// </summary>
        [Output("weeklyRecurrence")]
        public Output<Outputs.BackupScheduleWeeklyRecurrence?> WeeklyRecurrence { get; private set; } = null!;


        /// <summary>
        /// Create a BackupSchedule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackupSchedule(string name, BackupScheduleArgs args, CustomResourceOptions? options = null)
            : base("gcp:firestore/backupSchedule:BackupSchedule", name, args ?? new BackupScheduleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackupSchedule(string name, Input<string> id, BackupScheduleState? state = null, CustomResourceOptions? options = null)
            : base("gcp:firestore/backupSchedule:BackupSchedule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BackupSchedule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackupSchedule Get(string name, Input<string> id, BackupScheduleState? state = null, CustomResourceOptions? options = null)
        {
            return new BackupSchedule(name, id, state, options);
        }
    }

    public sealed class BackupScheduleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// For a schedule that runs daily at a specified time.
        /// </summary>
        [Input("dailyRecurrence")]
        public Input<Inputs.BackupScheduleDailyRecurrenceArgs>? DailyRecurrence { get; set; }

        /// <summary>
        /// The Firestore database id. Defaults to `"(default)"`.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.
        /// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
        /// For a daily backup recurrence, set this to a value up to 7 days. If you set a weekly backup recurrence, set this to a value up to 14 weeks.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("retention", required: true)]
        public Input<string> Retention { get; set; } = null!;

        /// <summary>
        /// For a schedule that runs weekly on a specific day and time.
        /// Structure is documented below.
        /// </summary>
        [Input("weeklyRecurrence")]
        public Input<Inputs.BackupScheduleWeeklyRecurrenceArgs>? WeeklyRecurrence { get; set; }

        public BackupScheduleArgs()
        {
        }
        public static new BackupScheduleArgs Empty => new BackupScheduleArgs();
    }

    public sealed class BackupScheduleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// For a schedule that runs daily at a specified time.
        /// </summary>
        [Input("dailyRecurrence")]
        public Input<Inputs.BackupScheduleDailyRecurrenceGetArgs>? DailyRecurrence { get; set; }

        /// <summary>
        /// The Firestore database id. Defaults to `"(default)"`.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// The unique backup schedule identifier across all locations and databases for the given project. Format:
        /// `projects/{{project}}/databases/{{database}}/backupSchedules/{{backupSchedule}}`
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
        /// At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.
        /// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
        /// For a daily backup recurrence, set this to a value up to 7 days. If you set a weekly backup recurrence, set this to a value up to 14 weeks.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("retention")]
        public Input<string>? Retention { get; set; }

        /// <summary>
        /// For a schedule that runs weekly on a specific day and time.
        /// Structure is documented below.
        /// </summary>
        [Input("weeklyRecurrence")]
        public Input<Inputs.BackupScheduleWeeklyRecurrenceGetArgs>? WeeklyRecurrence { get; set; }

        public BackupScheduleState()
        {
        }
        public static new BackupScheduleState Empty => new BackupScheduleState();
    }
}
