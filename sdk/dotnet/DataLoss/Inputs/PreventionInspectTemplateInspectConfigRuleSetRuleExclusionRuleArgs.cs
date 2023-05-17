// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataLoss.Inputs
{

    public sealed class PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dictionary which defines the rule.
        /// Structure is documented below.
        /// </summary>
        [Input("dictionary")]
        public Input<Inputs.PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleDictionaryArgs>? Dictionary { get; set; }

        /// <summary>
        /// Drop if the hotword rule is contained in the proximate context.
        /// For tabular data, the context includes the column name.
        /// Structure is documented below.
        /// </summary>
        [Input("excludeByHotword")]
        public Input<Inputs.PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleExcludeByHotwordArgs>? ExcludeByHotword { get; set; }

        /// <summary>
        /// Set of infoTypes for which findings would affect this rule.
        /// Structure is documented below.
        /// </summary>
        [Input("excludeInfoTypes")]
        public Input<Inputs.PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleExcludeInfoTypesArgs>? ExcludeInfoTypes { get; set; }

        /// <summary>
        /// How the rule is applied. See the documentation for more information: https://cloud.google.com/dlp/docs/reference/rest/v2/InspectConfig#MatchingType
        /// Possible values are: `MATCHING_TYPE_FULL_MATCH`, `MATCHING_TYPE_PARTIAL_MATCH`, `MATCHING_TYPE_INVERSE_MATCH`.
        /// </summary>
        [Input("matchingType", required: true)]
        public Input<string> MatchingType { get; set; } = null!;

        /// <summary>
        /// Regular expression which defines the rule.
        /// Structure is documented below.
        /// </summary>
        [Input("regex")]
        public Input<Inputs.PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleRegexArgs>? Regex { get; set; }

        public PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleArgs()
        {
        }
        public static new PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleArgs Empty => new PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleArgs();
    }
}
