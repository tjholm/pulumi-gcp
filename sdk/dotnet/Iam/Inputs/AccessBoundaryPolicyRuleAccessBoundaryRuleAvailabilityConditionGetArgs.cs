// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Iam.Inputs
{

    public sealed class AccessBoundaryPolicyRuleAccessBoundaryRuleAvailabilityConditionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the expression. This is a longer text which describes the expression,
        /// e.g. when hovered over it in a UI.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Textual representation of an expression in Common Expression Language syntax.
        /// </summary>
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        /// <summary>
        /// String indicating the location of the expression for error reporting,
        /// e.g. a file name and a position in the file.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Title for the expression, i.e. a short string describing its purpose.
        /// This can be used e.g. in UIs which allow to enter the expression.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public AccessBoundaryPolicyRuleAccessBoundaryRuleAvailabilityConditionGetArgs()
        {
        }
        public static new AccessBoundaryPolicyRuleAccessBoundaryRuleAvailabilityConditionGetArgs Empty => new AccessBoundaryPolicyRuleAccessBoundaryRuleAvailabilityConditionGetArgs();
    }
}
