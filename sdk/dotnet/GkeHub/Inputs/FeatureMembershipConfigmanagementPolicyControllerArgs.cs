// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeHub.Inputs
{

    public sealed class FeatureMembershipConfigmanagementPolicyControllerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sets the interval for Policy Controller Audit Scans (in seconds). When set to 0, this disables audit functionality altogether.
        /// </summary>
        [Input("auditIntervalSeconds")]
        public Input<string>? AuditIntervalSeconds { get; set; }

        /// <summary>
        /// Enables the installation of Policy Controller. If false, the rest of PolicyController fields take no effect.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("exemptableNamespaces")]
        private InputList<string>? _exemptableNamespaces;

        /// <summary>
        /// The set of namespaces that are excluded from Policy Controller checks. Namespaces do not need to currently exist on the cluster.
        /// </summary>
        public InputList<string> ExemptableNamespaces
        {
            get => _exemptableNamespaces ?? (_exemptableNamespaces = new InputList<string>());
            set => _exemptableNamespaces = value;
        }

        /// <summary>
        /// Logs all denies and dry run failures.
        /// </summary>
        [Input("logDeniesEnabled")]
        public Input<bool>? LogDeniesEnabled { get; set; }

        /// <summary>
        /// Enables the ability to use Constraint Templates that reference to objects other than the object currently being evaluated.
        /// </summary>
        [Input("referentialRulesEnabled")]
        public Input<bool>? ReferentialRulesEnabled { get; set; }

        /// <summary>
        /// Installs the default template library along with Policy Controller.
        /// </summary>
        [Input("templateLibraryInstalled")]
        public Input<bool>? TemplateLibraryInstalled { get; set; }

        public FeatureMembershipConfigmanagementPolicyControllerArgs()
        {
        }
    }
}
