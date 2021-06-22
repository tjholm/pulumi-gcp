// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeHub.Outputs
{

    [OutputType]
    public sealed class FeatureMembershipConfigmanagementHierarchyController
    {
        /// <summary>
        /// Whether hierarchical resource quota is enabled in this cluster.
        /// </summary>
        public readonly bool? EnableHierarchicalResourceQuota;
        /// <summary>
        /// Whether pod tree labels are enabled in this cluster.
        /// </summary>
        public readonly bool? EnablePodTreeLabels;
        /// <summary>
        /// Enables the installation of Policy Controller. If false, the rest of PolicyController fields take no effect.
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private FeatureMembershipConfigmanagementHierarchyController(
            bool? enableHierarchicalResourceQuota,

            bool? enablePodTreeLabels,

            bool? enabled)
        {
            EnableHierarchicalResourceQuota = enableHierarchicalResourceQuota;
            EnablePodTreeLabels = enablePodTreeLabels;
            Enabled = enabled;
        }
    }
}
