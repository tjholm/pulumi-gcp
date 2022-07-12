// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Sql.Inputs
{

    public sealed class DatabaseInstanceSettingsSqlServerAuditConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the destination bucket (e.g., gs://mybucket).
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// How long to keep generated audit files. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        /// </summary>
        [Input("retentionInterval")]
        public Input<string>? RetentionInterval { get; set; }

        /// <summary>
        /// How often to upload generated audit files. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        /// </summary>
        [Input("uploadInterval")]
        public Input<string>? UploadInterval { get; set; }

        public DatabaseInstanceSettingsSqlServerAuditConfigGetArgs()
        {
        }
    }
}
