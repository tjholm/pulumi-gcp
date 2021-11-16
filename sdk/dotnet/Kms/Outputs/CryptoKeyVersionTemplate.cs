// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Kms.Outputs
{

    [OutputType]
    public sealed class CryptoKeyVersionTemplate
    {
        /// <summary>
        /// The algorithm to use when creating a version based on this template.
        /// See the [algorithm reference](https://cloud.google.com/kms/docs/reference/rest/v1/CryptoKeyVersionAlgorithm) for possible inputs.
        /// </summary>
        public readonly string Algorithm;
        /// <summary>
        /// The protection level to use when creating a version based on this template. Possible values include "SOFTWARE", "HSM", "EXTERNAL". Defaults to "SOFTWARE".
        /// </summary>
        public readonly string? ProtectionLevel;

        [OutputConstructor]
        private CryptoKeyVersionTemplate(
            string algorithm,

            string? protectionLevel)
        {
            Algorithm = algorithm;
            ProtectionLevel = protectionLevel;
        }
    }
}
