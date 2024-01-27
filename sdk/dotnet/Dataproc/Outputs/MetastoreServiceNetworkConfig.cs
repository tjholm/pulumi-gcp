// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Outputs
{

    [OutputType]
    public sealed class MetastoreServiceNetworkConfig
    {
        /// <summary>
        /// The consumer-side network configuration for the Dataproc Metastore instance.
        /// Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.MetastoreServiceNetworkConfigConsumer> Consumers;
        /// <summary>
        /// (Optional, Beta)
        /// Enables custom routes to be imported and exported for the Dataproc Metastore service's peered VPC network.
        /// </summary>
        public readonly bool? CustomRoutesEnabled;

        [OutputConstructor]
        private MetastoreServiceNetworkConfig(
            ImmutableArray<Outputs.MetastoreServiceNetworkConfigConsumer> consumers,

            bool? customRoutesEnabled)
        {
            Consumers = consumers;
            CustomRoutesEnabled = customRoutesEnabled;
        }
    }
}
