// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Datastream.Inputs
{

    public sealed class StreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsDatasetTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If supplied, every created dataset will have its name prefixed by the provided value.
        /// The prefix and name will be separated by an underscore. i.e. _.
        /// </summary>
        [Input("datasetIdPrefix")]
        public Input<string>? DatasetIdPrefix { get; set; }

        /// <summary>
        /// Describes the Cloud KMS encryption key that will be used to protect destination BigQuery
        /// table. The BigQuery Service Account associated with your project requires access to this
        /// encryption key. i.e. projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{cryptoKey}.
        /// See https://cloud.google.com/bigquery/docs/customer-managed-encryption for more information.
        /// </summary>
        [Input("kmsKeyName")]
        public Input<string>? KmsKeyName { get; set; }

        /// <summary>
        /// The geographic location where the dataset should reside.
        /// See https://cloud.google.com/bigquery/docs/locations for supported locations.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        public StreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsDatasetTemplateArgs()
        {
        }
        public static new StreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsDatasetTemplateArgs Empty => new StreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsDatasetTemplateArgs();
    }
}
