// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Alloydb.Inputs
{

    public sealed class ClusterEncryptionInfoGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Output)
        /// Output only. Type of encryption.
        /// </summary>
        [Input("encryptionType")]
        public Input<string>? EncryptionType { get; set; }

        [Input("kmsKeyVersions")]
        private InputList<string>? _kmsKeyVersions;

        /// <summary>
        /// (Output)
        /// Output only. Cloud KMS key versions that are being used to protect the database or the backup.
        /// </summary>
        public InputList<string> KmsKeyVersions
        {
            get => _kmsKeyVersions ?? (_kmsKeyVersions = new InputList<string>());
            set => _kmsKeyVersions = value;
        }

        public ClusterEncryptionInfoGetArgs()
        {
        }
        public static new ClusterEncryptionInfoGetArgs Empty => new ClusterEncryptionInfoGetArgs();
    }
}
