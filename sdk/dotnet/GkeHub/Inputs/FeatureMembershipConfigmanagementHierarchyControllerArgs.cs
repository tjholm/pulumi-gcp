// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeHub.Inputs
{

    public sealed class FeatureMembershipConfigmanagementHierarchyControllerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether hierarchical resource quota is enabled in this cluster.
        /// </summary>
        [Input("enableHierarchicalResourceQuota")]
        public Input<bool>? EnableHierarchicalResourceQuota { get; set; }

        /// <summary>
        /// Whether pod tree labels are enabled in this cluster.
        /// </summary>
        [Input("enablePodTreeLabels")]
        public Input<bool>? EnablePodTreeLabels { get; set; }

        /// <summary>
        /// Enables the installation of Policy Controller. If false, the rest of PolicyController fields take no effect.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public FeatureMembershipConfigmanagementHierarchyControllerArgs()
        {
        }
    }
}
