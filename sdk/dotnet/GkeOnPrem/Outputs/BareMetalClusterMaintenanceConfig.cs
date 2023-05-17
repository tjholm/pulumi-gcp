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
    public sealed class BareMetalClusterMaintenanceConfig
    {
        /// <summary>
        /// All IPv4 address from these ranges will be placed into maintenance mode.
        /// Nodes in maintenance mode will be cordoned and drained. When both of these
        /// are true, the "baremetal.cluster.gke.io/maintenance" annotation will be set
        /// on the node resource.
        /// </summary>
        public readonly ImmutableArray<string> MaintenanceAddressCidrBlocks;

        [OutputConstructor]
        private BareMetalClusterMaintenanceConfig(ImmutableArray<string> maintenanceAddressCidrBlocks)
        {
            MaintenanceAddressCidrBlocks = maintenanceAddressCidrBlocks;
        }
    }
}
