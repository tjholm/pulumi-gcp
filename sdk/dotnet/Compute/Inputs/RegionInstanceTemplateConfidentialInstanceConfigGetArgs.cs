// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionInstanceTemplateConfidentialInstanceConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines the confidential computing technology the instance uses. SEV is an AMD feature. One of the following values: `SEV`, `SEV_SNP`. `on_host_maintenance` can be set to MIGRATE if `confidential_instance_type` is set to `SEV` and `min_cpu_platform` is set to `"AMD Milan"`. Otherwise, `on_host_maintenance` has to be set to TERMINATE or this will fail to create the VM. If `SEV_SNP`, currently `min_cpu_platform` has to be set to `"AMD Milan"` or this will fail to create the VM.
        /// </summary>
        [Input("confidentialInstanceType")]
        public Input<string>? ConfidentialInstanceType { get; set; }

        /// <summary>
        /// Defines whether the instance should have confidential compute enabled with AMD SEV. If enabled, `on_host_maintenance` can be set to MIGRATE if `min_cpu_platform` is set to `"AMD Milan"`. Otherwise, `on_host_maintenance` has to be set to TERMINATE or this will fail to create the VM.
        /// </summary>
        [Input("enableConfidentialCompute")]
        public Input<bool>? EnableConfidentialCompute { get; set; }

        public RegionInstanceTemplateConfidentialInstanceConfigGetArgs()
        {
        }
        public static new RegionInstanceTemplateConfidentialInstanceConfigGetArgs Empty => new RegionInstanceTemplateConfidentialInstanceConfigGetArgs();
    }
}
