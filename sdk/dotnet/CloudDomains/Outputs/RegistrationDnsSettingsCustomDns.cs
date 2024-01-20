// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudDomains.Outputs
{

    [OutputType]
    public sealed class RegistrationDnsSettingsCustomDns
    {
        /// <summary>
        /// The list of DS records for this domain, which are used to enable DNSSEC. The domain's DNS provider can provide
        /// the values to set here. If this field is empty, DNSSEC is disabled.
        /// Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.RegistrationDnsSettingsCustomDnsDsRecord> DsRecords;
        /// <summary>
        /// Required. A list of name servers that store the DNS zone for this domain. Each name server is a domain
        /// name, with Unicode domain names expressed in Punycode format.
        /// </summary>
        public readonly ImmutableArray<string> NameServers;

        [OutputConstructor]
        private RegistrationDnsSettingsCustomDns(
            ImmutableArray<Outputs.RegistrationDnsSettingsCustomDnsDsRecord> dsRecords,

            ImmutableArray<string> nameServers)
        {
            DsRecords = dsRecords;
            NameServers = nameServers;
        }
    }
}
