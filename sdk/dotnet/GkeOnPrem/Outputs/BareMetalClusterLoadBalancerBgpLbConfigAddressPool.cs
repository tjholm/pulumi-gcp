// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeOnPrem.Outputs
{

    [OutputType]
    public sealed class BareMetalClusterLoadBalancerBgpLbConfigAddressPool
    {
        /// <summary>
        /// The addresses that are part of this pool. Each address must be either in the CIDR form (1.2.3.0/24) or range form (1.2.3.1-1.2.3.5).
        /// </summary>
        public readonly ImmutableArray<string> Addresses;
        /// <summary>
        /// If true, avoid using IPs ending in .0 or .255.
        /// This avoids buggy consumer devices mistakenly dropping IPv4 traffic for those special IP addresses.
        /// </summary>
        public readonly bool? AvoidBuggyIps;
        /// <summary>
        /// If true, prevent IP addresses from being automatically assigned.
        /// </summary>
        public readonly string? ManualAssign;
        /// <summary>
        /// The name of the address pool.
        /// </summary>
        public readonly string Pool;

        [OutputConstructor]
        private BareMetalClusterLoadBalancerBgpLbConfigAddressPool(
            ImmutableArray<string> addresses,

            bool? avoidBuggyIps,

            string? manualAssign,

            string pool)
        {
            Addresses = addresses;
            AvoidBuggyIps = avoidBuggyIps;
            ManualAssign = manualAssign;
            Pool = pool;
        }
    }
}
