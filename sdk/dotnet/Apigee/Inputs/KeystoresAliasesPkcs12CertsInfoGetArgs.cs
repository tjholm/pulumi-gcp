// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Apigee.Inputs
{

    public sealed class KeystoresAliasesPkcs12CertsInfoGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("certInfos")]
        private InputList<Inputs.KeystoresAliasesPkcs12CertsInfoCertInfoGetArgs>? _certInfos;

        /// <summary>
        /// (Output)
        /// List of all properties in the object.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.KeystoresAliasesPkcs12CertsInfoCertInfoGetArgs> CertInfos
        {
            get => _certInfos ?? (_certInfos = new InputList<Inputs.KeystoresAliasesPkcs12CertsInfoCertInfoGetArgs>());
            set => _certInfos = value;
        }

        public KeystoresAliasesPkcs12CertsInfoGetArgs()
        {
        }
        public static new KeystoresAliasesPkcs12CertsInfoGetArgs Empty => new KeystoresAliasesPkcs12CertsInfoGetArgs();
    }
}
