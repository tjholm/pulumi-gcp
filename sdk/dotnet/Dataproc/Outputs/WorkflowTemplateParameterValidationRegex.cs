// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Outputs
{

    [OutputType]
    public sealed class WorkflowTemplateParameterValidationRegex
    {
        /// <summary>
        /// Required. RE2 regular expressions used to validate the parameter's value. The value must match the regex in its entirety (substring matches are not sufficient).
        /// </summary>
        public readonly ImmutableArray<string> Regexes;

        [OutputConstructor]
        private WorkflowTemplateParameterValidationRegex(ImmutableArray<string> regexes)
        {
            Regexes = regexes;
        }
    }
}
