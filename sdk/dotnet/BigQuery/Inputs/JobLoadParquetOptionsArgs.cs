// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery.Inputs
{

    public sealed class JobLoadParquetOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If sourceFormat is set to PARQUET, indicates whether to use schema inference specifically for Parquet LIST logical type.
        /// </summary>
        [Input("enableListInference")]
        public Input<bool>? EnableListInference { get; set; }

        /// <summary>
        /// If sourceFormat is set to PARQUET, indicates whether to infer Parquet ENUM logical type as STRING instead of BYTES by default.
        /// </summary>
        [Input("enumAsString")]
        public Input<bool>? EnumAsString { get; set; }

        public JobLoadParquetOptionsArgs()
        {
        }
        public static new JobLoadParquetOptionsArgs Empty => new JobLoadParquetOptionsArgs();
    }
}
