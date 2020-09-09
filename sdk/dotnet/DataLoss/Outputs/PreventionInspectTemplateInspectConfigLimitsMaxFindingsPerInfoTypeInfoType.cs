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
    public sealed class PreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType
    {
        /// <summary>
        /// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
        /// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private PreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(string name)
        {
            Name = name;
        }
    }
}
