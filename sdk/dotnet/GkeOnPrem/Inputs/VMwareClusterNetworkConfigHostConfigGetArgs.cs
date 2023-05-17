// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeOnPrem.Inputs
{

    public sealed class VMwareClusterNetworkConfigHostConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("dnsSearchDomains")]
        private InputList<string>? _dnsSearchDomains;

        /// <summary>
        /// DNS search domains.
        /// </summary>
        public InputList<string> DnsSearchDomains
        {
            get => _dnsSearchDomains ?? (_dnsSearchDomains = new InputList<string>());
            set => _dnsSearchDomains = value;
        }

        [Input("dnsServers")]
        private InputList<string>? _dnsServers;

        /// <summary>
        /// DNS servers.
        /// </summary>
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        [Input("ntpServers")]
        private InputList<string>? _ntpServers;

        /// <summary>
        /// NTP servers.
        /// </summary>
        public InputList<string> NtpServers
        {
            get => _ntpServers ?? (_ntpServers = new InputList<string>());
            set => _ntpServers = value;
        }

        public VMwareClusterNetworkConfigHostConfigGetArgs()
        {
        }
        public static new VMwareClusterNetworkConfigHostConfigGetArgs Empty => new VMwareClusterNetworkConfigHostConfigGetArgs();
    }
}
