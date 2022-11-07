// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Vertex.Outputs
{

    [OutputType]
    public sealed class AiEndpointDeployedModelDedicatedResourceMachineSpec
    {
        public readonly int? AcceleratorCount;
        public readonly string? AcceleratorType;
        public readonly string? MachineType;

        [OutputConstructor]
        private AiEndpointDeployedModelDedicatedResourceMachineSpec(
            int? acceleratorCount,

            string? acceleratorType,

            string? machineType)
        {
            AcceleratorCount = acceleratorCount;
            AcceleratorType = acceleratorType;
            MachineType = machineType;
        }
    }
}
