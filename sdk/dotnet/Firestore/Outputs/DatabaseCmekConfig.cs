// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Firestore.Outputs
{

    [OutputType]
    public sealed class DatabaseCmekConfig
    {
        /// <summary>
        /// (Output)
        /// Currently in-use KMS key versions (https://cloud.google.com/kms/docs/resource-hierarchy#key_versions).
        /// During key rotation (https://cloud.google.com/kms/docs/key-rotation), there can be
        /// multiple in-use key versions.
        /// The expected format is
        /// `projects/{project_id}/locations/{kms_location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}/cryptoKeyVersions/{key_version}`.
        /// </summary>
        public readonly ImmutableArray<string> ActiveKeyVersions;
        /// <summary>
        /// The resource ID of a Cloud KMS key. If set, the database created will
        /// be a Customer-managed Encryption Key (CMEK) database encrypted with
        /// this key. This feature is allowlist only in initial launch.
        /// Only keys in the same location as this database are allowed to be used
        /// for encryption. For Firestore's nam5 multi-region, this corresponds to Cloud KMS
        /// multi-region us. For Firestore's eur3 multi-region, this corresponds to
        /// Cloud KMS multi-region europe. See https://cloud.google.com/kms/docs/locations.
        /// This value should be the KMS key resource ID in the format of
        /// `projects/{project_id}/locations/{kms_location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
        /// How to retrive this resource ID is listed at
        /// https://cloud.google.com/kms/docs/getting-resource-ids#getting_the_id_for_a_key_and_version.
        /// </summary>
        public readonly string KmsKeyName;

        [OutputConstructor]
        private DatabaseCmekConfig(
            ImmutableArray<string> activeKeyVersions,

            string kmsKeyName)
        {
            ActiveKeyVersions = activeKeyVersions;
            KmsKeyName = kmsKeyName;
        }
    }
}
