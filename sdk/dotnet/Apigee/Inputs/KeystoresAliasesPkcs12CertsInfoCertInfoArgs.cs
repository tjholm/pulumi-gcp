// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Apigee.Inputs
{

    public sealed class KeystoresAliasesPkcs12CertsInfoCertInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Output)
        /// X.509 basic constraints extension.
        /// </summary>
        [Input("basicConstraints")]
        public Input<string>? BasicConstraints { get; set; }

        /// <summary>
        /// (Output)
        /// X.509 notAfter validity period in milliseconds since epoch.
        /// </summary>
        [Input("expiryDate")]
        public Input<string>? ExpiryDate { get; set; }

        /// <summary>
        /// (Output)
        /// Flag that specifies whether the certificate is valid.
        /// Flag is set to Yes if the certificate is valid, No if expired, or Not yet if not yet valid.
        /// </summary>
        [Input("isValid")]
        public Input<string>? IsValid { get; set; }

        /// <summary>
        /// (Output)
        /// X.509 issuer.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// (Output)
        /// Public key component of the X.509 subject public key info.
        /// </summary>
        [Input("publicKey")]
        public Input<string>? PublicKey { get; set; }

        /// <summary>
        /// (Output)
        /// X.509 serial number.
        /// </summary>
        [Input("serialNumber")]
        public Input<string>? SerialNumber { get; set; }

        /// <summary>
        /// (Output)
        /// X.509 signatureAlgorithm.
        /// </summary>
        [Input("sigAlgName")]
        public Input<string>? SigAlgName { get; set; }

        /// <summary>
        /// (Output)
        /// X.509 subject.
        /// </summary>
        [Input("subject")]
        public Input<string>? Subject { get; set; }

        [Input("subjectAlternativeNames")]
        private InputList<string>? _subjectAlternativeNames;

        /// <summary>
        /// (Output)
        /// X.509 subject alternative names (SANs) extension.
        /// </summary>
        public InputList<string> SubjectAlternativeNames
        {
            get => _subjectAlternativeNames ?? (_subjectAlternativeNames = new InputList<string>());
            set => _subjectAlternativeNames = value;
        }

        /// <summary>
        /// (Output)
        /// X.509 notBefore validity period in milliseconds since epoch.
        /// </summary>
        [Input("validFrom")]
        public Input<string>? ValidFrom { get; set; }

        /// <summary>
        /// (Output)
        /// X.509 version.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public KeystoresAliasesPkcs12CertsInfoCertInfoArgs()
        {
        }
        public static new KeystoresAliasesPkcs12CertsInfoCertInfoArgs Empty => new KeystoresAliasesPkcs12CertsInfoCertInfoArgs();
    }
}
