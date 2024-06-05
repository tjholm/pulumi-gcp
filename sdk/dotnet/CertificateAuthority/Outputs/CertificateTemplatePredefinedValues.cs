// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CertificateAuthority.Outputs
{

    [OutputType]
    public sealed class CertificateTemplatePredefinedValues
    {
        /// <summary>
        /// Optional. Describes custom X.509 extensions.
        /// Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.CertificateTemplatePredefinedValuesAdditionalExtension> AdditionalExtensions;
        /// <summary>
        /// Optional. Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate.
        /// </summary>
        public readonly ImmutableArray<string> AiaOcspServers;
        /// <summary>
        /// Optional. Describes options in this X509Parameters that are relevant in a CA certificate.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.CertificateTemplatePredefinedValuesCaOptions? CaOptions;
        /// <summary>
        /// Optional. Indicates the intended use for keys that correspond to a certificate.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.CertificateTemplatePredefinedValuesKeyUsage? KeyUsage;
        /// <summary>
        /// Optional. Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4.
        /// Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.CertificateTemplatePredefinedValuesPolicyId> PolicyIds;

        [OutputConstructor]
        private CertificateTemplatePredefinedValues(
            ImmutableArray<Outputs.CertificateTemplatePredefinedValuesAdditionalExtension> additionalExtensions,

            ImmutableArray<string> aiaOcspServers,

            Outputs.CertificateTemplatePredefinedValuesCaOptions? caOptions,

            Outputs.CertificateTemplatePredefinedValuesKeyUsage? keyUsage,

            ImmutableArray<Outputs.CertificateTemplatePredefinedValuesPolicyId> policyIds)
        {
            AdditionalExtensions = additionalExtensions;
            AiaOcspServers = aiaOcspServers;
            CaOptions = caOptions;
            KeyUsage = keyUsage;
            PolicyIds = policyIds;
        }
    }
}
