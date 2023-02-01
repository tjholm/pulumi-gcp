// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Iam.Outputs
{

    [OutputType]
    public sealed class AccessBoundaryPolicyRuleAccessBoundaryRuleAvailabilityCondition
    {
        /// <summary>
        /// Description of the expression. This is a longer text which describes the expression,
        /// e.g. when hovered over it in a UI.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Textual representation of an expression in Common Expression Language syntax.
        /// </summary>
        public readonly string Expression;
        /// <summary>
        /// String indicating the location of the expression for error reporting,
        /// e.g. a file name and a position in the file.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Title for the expression, i.e. a short string describing its purpose.
        /// This can be used e.g. in UIs which allow to enter the expression.
        /// </summary>
        public readonly string? Title;

        [OutputConstructor]
        private AccessBoundaryPolicyRuleAccessBoundaryRuleAvailabilityCondition(
            string? description,

            string expression,

            string? location,

            string? title)
        {
            Description = description;
            Expression = expression;
            Location = location;
            Title = title;
        }
    }
}
