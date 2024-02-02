// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudRun.Outputs
{

    [OutputType]
    public sealed class GetServiceTemplateSpecVolumeCsiResult
    {
        /// <summary>
        /// Unique name representing the type of file system to be created. Cloud Run supports the following values:
        ///   * gcsfuse.run.googleapis.com: Mount a Google Cloud Storage bucket using GCSFuse. This driver requires the
        ///     run.googleapis.com/execution-environment annotation to be set to "gen2" and
        ///     run.googleapis.com/launch-stage set to "BETA" or "ALPHA".
        /// </summary>
        public readonly string Driver;
        /// <summary>
        /// If true, all mounts created from this volume will be read-only.
        /// </summary>
        public readonly bool ReadOnly;
        /// <summary>
        /// Driver-specific attributes. The following options are supported for available drivers:
        ///   * gcsfuse.run.googleapis.com
        ///     * bucketName: The name of the Cloud Storage Bucket that backs this volume. The Cloud Run Service identity must have access to this bucket.
        /// </summary>
        public readonly ImmutableDictionary<string, string> VolumeAttributes;

        [OutputConstructor]
        private GetServiceTemplateSpecVolumeCsiResult(
            string driver,

            bool readOnly,

            ImmutableDictionary<string, string> volumeAttributes)
        {
            Driver = driver;
            ReadOnly = readOnly;
            VolumeAttributes = volumeAttributes;
        }
    }
}
