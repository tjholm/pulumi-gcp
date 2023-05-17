// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataLoss.Inputs
{

    public sealed class PreventionJobTriggerInspectJobInspectConfigRuleSetRuleExclusionRuleDictionaryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Newline-delimited file of words in Cloud Storage. Only a single file is accepted.
        /// Structure is documented below.
        /// </summary>
        [Input("cloudStoragePath")]
        public Input<Inputs.PreventionJobTriggerInspectJobInspectConfigRuleSetRuleExclusionRuleDictionaryCloudStoragePathArgs>? CloudStoragePath { get; set; }

        /// <summary>
        /// List of words or phrases to search for.
        /// Structure is documented below.
        /// </summary>
        [Input("wordList")]
        public Input<Inputs.PreventionJobTriggerInspectJobInspectConfigRuleSetRuleExclusionRuleDictionaryWordListArgs>? WordList { get; set; }

        public PreventionJobTriggerInspectJobInspectConfigRuleSetRuleExclusionRuleDictionaryArgs()
        {
        }
        public static new PreventionJobTriggerInspectJobInspectConfigRuleSetRuleExclusionRuleDictionaryArgs Empty => new PreventionJobTriggerInspectJobInspectConfigRuleSetRuleExclusionRuleDictionaryArgs();
    }
}
