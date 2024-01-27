// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Outputs
{

    [OutputType]
    public sealed class ClusterProtectConfigWorkloadConfig
    {
        /// <summary>
        /// Sets which mode of auditing should be used for the cluster's workloads. Accepted values are DISABLED, BASIC.
        /// </summary>
        public readonly string AuditMode;

        [OutputConstructor]
        private ClusterProtectConfigWorkloadConfig(string auditMode)
        {
            AuditMode = auditMode;
        }
    }
}
