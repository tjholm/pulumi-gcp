// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataLoss.Inputs
{

    public sealed class PreventionJobTriggerInspectJobInspectConfigRuleSetRuleHotwordRuleHotwordRegexGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("groupIndexes")]
        private InputList<int>? _groupIndexes;

        /// <summary>
        /// The index of the submatch to extract as findings. When not specified,
        /// the entire match is returned. No more than 3 may be included.
        /// </summary>
        public InputList<int> GroupIndexes
        {
            get => _groupIndexes ?? (_groupIndexes = new InputList<int>());
            set => _groupIndexes = value;
        }

        /// <summary>
        /// Pattern defining the regular expression. Its syntax
        /// (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        public PreventionJobTriggerInspectJobInspectConfigRuleSetRuleHotwordRuleHotwordRegexGetArgs()
        {
        }
        public static new PreventionJobTriggerInspectJobInspectConfigRuleSetRuleHotwordRuleHotwordRegexGetArgs Empty => new PreventionJobTriggerInspectJobInspectConfigRuleSetRuleHotwordRuleHotwordRegexGetArgs();
    }
}
