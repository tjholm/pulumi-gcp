// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Apigee.Outputs
{

    [OutputType]
    public sealed class KeystoresAliasesPkcs12CertsInfo
    {
        /// <summary>
        /// (Output)
        /// List of all properties in the object.
        /// Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.KeystoresAliasesPkcs12CertsInfoCertInfo> CertInfos;

        [OutputConstructor]
        private KeystoresAliasesPkcs12CertsInfo(ImmutableArray<Outputs.KeystoresAliasesPkcs12CertsInfoCertInfo> certInfos)
        {
            CertInfos = certInfos;
        }
    }
}
