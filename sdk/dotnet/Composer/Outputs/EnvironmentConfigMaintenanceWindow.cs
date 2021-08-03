// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Composer.Outputs
{

    [OutputType]
    public sealed class EnvironmentConfigMaintenanceWindow
    {
        /// <summary>
        /// Maintenance window end time. It is used only to calculate the duration of the maintenance window.
        /// The value for end-time must be in the future, relative to 'start_time'.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// Maintenance window recurrence. Format is a subset of RFC-5545 (https://tools.ietf.org/html/rfc5545) 'RRULE'.
        /// The only allowed values for 'FREQ' field are 'FREQ=DAILY' and 'FREQ=WEEKLY;BYDAY=...'.
        /// Example values: 'FREQ=WEEKLY;BYDAY=TU,WE', 'FREQ=DAILY'.
        /// </summary>
        public readonly string Recurrence;
        /// <summary>
        /// Start time of the first recurrence of the maintenance window.
        /// </summary>
        public readonly string StartTime;

        [OutputConstructor]
        private EnvironmentConfigMaintenanceWindow(
            string endTime,

            string recurrence,

            string startTime)
        {
            EndTime = endTime;
            Recurrence = recurrence;
            StartTime = startTime;
        }
    }
}
