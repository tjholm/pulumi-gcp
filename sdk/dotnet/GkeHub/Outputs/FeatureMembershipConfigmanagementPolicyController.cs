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
    public sealed class FeatureMembershipConfigmanagementPolicyController
    {
        /// <summary>
        /// Sets the interval for Policy Controller Audit Scans (in seconds). When set to 0, this disables audit functionality altogether.
        /// </summary>
        public readonly string? AuditIntervalSeconds;
        /// <summary>
        /// Enables the installation of Policy Controller. If false, the rest of PolicyController fields take no effect.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The set of namespaces that are excluded from Policy Controller checks. Namespaces do not need to currently exist on the cluster.
        /// </summary>
        public readonly ImmutableArray<string> ExemptableNamespaces;
        /// <summary>
        /// Logs all denies and dry run failures.
        /// </summary>
        public readonly bool? LogDeniesEnabled;
        public readonly Outputs.FeatureMembershipConfigmanagementPolicyControllerMonitoring? Monitoring;
        public readonly bool? MutationEnabled;
        /// <summary>
        /// Enables the ability to use Constraint Templates that reference to objects other than the object currently being evaluated.
        /// </summary>
        public readonly bool? ReferentialRulesEnabled;
        /// <summary>
        /// Installs the default template library along with Policy Controller.
        /// </summary>
        public readonly bool? TemplateLibraryInstalled;

        [OutputConstructor]
        private FeatureMembershipConfigmanagementPolicyController(
            string? auditIntervalSeconds,

            bool? enabled,

            ImmutableArray<string> exemptableNamespaces,

            bool? logDeniesEnabled,

            Outputs.FeatureMembershipConfigmanagementPolicyControllerMonitoring? monitoring,

            bool? mutationEnabled,

            bool? referentialRulesEnabled,

            bool? templateLibraryInstalled)
        {
            AuditIntervalSeconds = auditIntervalSeconds;
            Enabled = enabled;
            ExemptableNamespaces = exemptableNamespaces;
            LogDeniesEnabled = logDeniesEnabled;
            Monitoring = monitoring;
            MutationEnabled = mutationEnabled;
            ReferentialRulesEnabled = referentialRulesEnabled;
            TemplateLibraryInstalled = templateLibraryInstalled;
        }
    }
}
