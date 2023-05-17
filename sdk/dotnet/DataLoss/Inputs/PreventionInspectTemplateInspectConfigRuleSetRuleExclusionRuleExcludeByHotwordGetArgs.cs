// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataLoss.Inputs
{

    public sealed class PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleExcludeByHotwordGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Regular expression pattern defining what qualifies as a hotword.
        /// Structure is documented below.
        /// </summary>
        [Input("hotwordRegex", required: true)]
        public Input<Inputs.PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleExcludeByHotwordHotwordRegexGetArgs> HotwordRegex { get; set; } = null!;

        /// <summary>
        /// Proximity of the finding within which the entire hotword must reside. The total length of the window cannot
        /// exceed 1000 characters. Note that the finding itself will be included in the window, so that hotwords may be
        /// used to match substrings of the finding itself. For example, the certainty of a phone number regex
        /// `(\d{3}) \d{3}-\d{4}` could be adjusted upwards if the area code is known to be the local area code of a company
        /// office using the hotword regex `(xxx)`, where `xxx` is the area code in question.
        /// Structure is documented below.
        /// </summary>
        [Input("proximity", required: true)]
        public Input<Inputs.PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleExcludeByHotwordProximityGetArgs> Proximity { get; set; } = null!;

        public PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleExcludeByHotwordGetArgs()
        {
        }
        public static new PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleExcludeByHotwordGetArgs Empty => new PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleExcludeByHotwordGetArgs();
    }
}
