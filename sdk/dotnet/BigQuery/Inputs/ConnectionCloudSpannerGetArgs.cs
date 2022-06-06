// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery.Inputs
{

    public sealed class ConnectionCloudSpannerGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud Spanner database in the form `project/instance/database'
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// If parallelism should be used when reading from Cloud Spanner
        /// </summary>
        [Input("useParallelism")]
        public Input<bool>? UseParallelism { get; set; }

        public ConnectionCloudSpannerGetArgs()
        {
        }
    }
}
