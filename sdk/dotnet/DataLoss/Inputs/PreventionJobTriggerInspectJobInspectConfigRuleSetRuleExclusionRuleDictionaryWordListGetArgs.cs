// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataLoss.Inputs
{

    public sealed class PreventionJobTriggerInspectJobInspectConfigRuleSetRuleExclusionRuleDictionaryWordListGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("words", required: true)]
        private InputList<string>? _words;

        /// <summary>
        /// Words or phrases defining the dictionary. The dictionary must contain at least one
        /// phrase and every phrase must contain at least 2 characters that are letters or digits.
        /// </summary>
        public InputList<string> Words
        {
            get => _words ?? (_words = new InputList<string>());
            set => _words = value;
        }

        public PreventionJobTriggerInspectJobInspectConfigRuleSetRuleExclusionRuleDictionaryWordListGetArgs()
        {
        }
        public static new PreventionJobTriggerInspectJobInspectConfigRuleSetRuleExclusionRuleDictionaryWordListGetArgs Empty => new PreventionJobTriggerInspectJobInspectConfigRuleSetRuleExclusionRuleDictionaryWordListGetArgs();
    }
}
