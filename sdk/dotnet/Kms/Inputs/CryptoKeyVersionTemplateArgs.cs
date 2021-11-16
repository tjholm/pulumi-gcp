// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Kms.Inputs
{

    public sealed class CryptoKeyVersionTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The algorithm to use when creating a version based on this template.
        /// See the [algorithm reference](https://cloud.google.com/kms/docs/reference/rest/v1/CryptoKeyVersionAlgorithm) for possible inputs.
        /// </summary>
        [Input("algorithm", required: true)]
        public Input<string> Algorithm { get; set; } = null!;

        /// <summary>
        /// The protection level to use when creating a version based on this template. Possible values include "SOFTWARE", "HSM", "EXTERNAL". Defaults to "SOFTWARE".
        /// </summary>
        [Input("protectionLevel")]
        public Input<string>? ProtectionLevel { get; set; }

        public CryptoKeyVersionTemplateArgs()
        {
        }
    }
}
