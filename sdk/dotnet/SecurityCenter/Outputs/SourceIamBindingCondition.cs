// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.SecurityCenter.Outputs
{

    [OutputType]
    public sealed class SourceIamBindingCondition
    {
        /// <summary>
        /// The description of the source (max of 1024 characters).
        /// </summary>
        public readonly string? Description;
        public readonly string Expression;
        public readonly string Title;

        [OutputConstructor]
        private SourceIamBindingCondition(
            string? description,

            string expression,

            string title)
        {
            Description = description;
            Expression = expression;
            Title = title;
        }
    }
}
