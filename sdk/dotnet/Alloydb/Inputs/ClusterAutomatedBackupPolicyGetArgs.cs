// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Alloydb.Inputs
{

    public sealed class ClusterAutomatedBackupPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The length of the time window during which a backup can be taken. If a backup does not succeed within this time window, it will be canceled and considered failed.
        /// The backup window must be at least 5 minutes long. There is no upper bound on the window. If not set, it will default to 1 hour.
        /// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        /// </summary>
        [Input("backupWindow")]
        public Input<string>? BackupWindow { get; set; }

        /// <summary>
        /// Whether automated automated backups are enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to backups created using this configuration.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The location where the backup will be stored. Currently, the only supported option is to store the backup in the same region as the cluster.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Quantity-based Backup retention policy to retain recent backups.
        /// Structure is documented below.
        /// </summary>
        [Input("quantityBasedRetention")]
        public Input<Inputs.ClusterAutomatedBackupPolicyQuantityBasedRetentionGetArgs>? QuantityBasedRetention { get; set; }

        /// <summary>
        /// Time-based Backup retention policy.
        /// Structure is documented below.
        /// </summary>
        [Input("timeBasedRetention")]
        public Input<Inputs.ClusterAutomatedBackupPolicyTimeBasedRetentionGetArgs>? TimeBasedRetention { get; set; }

        /// <summary>
        /// Weekly schedule for the Backup.
        /// Structure is documented below.
        /// </summary>
        [Input("weeklySchedule", required: true)]
        public Input<Inputs.ClusterAutomatedBackupPolicyWeeklyScheduleGetArgs> WeeklySchedule { get; set; } = null!;

        public ClusterAutomatedBackupPolicyGetArgs()
        {
        }
        public static new ClusterAutomatedBackupPolicyGetArgs Empty => new ClusterAutomatedBackupPolicyGetArgs();
    }
}
