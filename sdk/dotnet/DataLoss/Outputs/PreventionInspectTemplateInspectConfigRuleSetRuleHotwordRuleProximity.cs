// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataLoss.Outputs
{

    [OutputType]
    public sealed class PreventionInspectTemplateInspectConfigRuleSetRuleHotwordRuleProximity
    {
        /// <summary>
        /// Number of characters after the finding to consider. Either this or window_before must be specified
        /// </summary>
        public readonly int? WindowAfter;
        /// <summary>
        /// Number of characters before the finding to consider. Either this or window_after must be specified
        /// </summary>
        public readonly int? WindowBefore;

        [OutputConstructor]
        private PreventionInspectTemplateInspectConfigRuleSetRuleHotwordRuleProximity(
            int? windowAfter,

            int? windowBefore)
        {
            WindowAfter = windowAfter;
            WindowBefore = windowBefore;
        }
    }
}
