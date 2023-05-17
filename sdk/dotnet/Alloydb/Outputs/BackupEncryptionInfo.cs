// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Alloydb.Outputs
{

    [OutputType]
    public sealed class BackupEncryptionInfo
    {
        /// <summary>
        /// (Output)
        /// Output only. Type of encryption.
        /// </summary>
        public readonly string? EncryptionType;
        /// <summary>
        /// (Output)
        /// Output only. Cloud KMS key versions that are being used to protect the database or the backup.
        /// </summary>
        public readonly ImmutableArray<string> KmsKeyVersions;

        [OutputConstructor]
        private BackupEncryptionInfo(
            string? encryptionType,

            ImmutableArray<string> kmsKeyVersions)
        {
            EncryptionType = encryptionType;
            KmsKeyVersions = kmsKeyVersions;
        }
    }
}
