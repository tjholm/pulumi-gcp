// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataLoss.Inputs
{

    public sealed class PreventionStoredInfoTypeDictionaryCloudStoragePathArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A url representing a file or path (no wildcards) in Cloud Storage. Example: `gs://[BUCKET_NAME]/dictionary.txt`
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public PreventionStoredInfoTypeDictionaryCloudStoragePathArgs()
        {
        }
    }
}
