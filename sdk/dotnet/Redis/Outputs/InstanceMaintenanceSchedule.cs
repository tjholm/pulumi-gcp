// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Redis.Outputs
{

    [OutputType]
    public sealed class InstanceMaintenanceSchedule
    {
        /// <summary>
        /// -
        /// Output only. The end time of any upcoming scheduled maintenance for this instance.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
        /// resolution and up to nine fractional digits.
        /// </summary>
        public readonly string? EndTime;
        /// <summary>
        /// -
        /// Output only. The deadline that the maintenance schedule start time
        /// can not go beyond, including reschedule.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
        /// resolution and up to nine fractional digits.
        /// </summary>
        public readonly string? ScheduleDeadlineTime;
        /// <summary>
        /// -
        /// Output only. The start time of any upcoming scheduled maintenance for this instance.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
        /// resolution and up to nine fractional digits.
        /// </summary>
        public readonly string? StartTime;

        [OutputConstructor]
        private InstanceMaintenanceSchedule(
            string? endTime,

            string? scheduleDeadlineTime,

            string? startTime)
        {
            EndTime = endTime;
            ScheduleDeadlineTime = scheduleDeadlineTime;
            StartTime = startTime;
        }
    }
}
