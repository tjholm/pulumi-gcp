// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataLoss.Inputs
{

    public sealed class PreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of information the findings limit applies to. Only one limit per infoType should be provided. If InfoTypeLimit does
        /// not have an infoType, the DLP API applies the limit against all infoTypes that are found but not
        /// specified in another InfoTypeLimit.
        /// Structure is documented below.
        /// </summary>
        [Input("infoType")]
        public Input<Inputs.PreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeGetArgs>? InfoType { get; set; }

        /// <summary>
        /// Max findings limit for the given infoType.
        /// </summary>
        [Input("maxFindings")]
        public Input<int>? MaxFindings { get; set; }

        public PreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeGetArgs()
        {
        }
        public static new PreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeGetArgs Empty => new PreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeGetArgs();
    }
}
