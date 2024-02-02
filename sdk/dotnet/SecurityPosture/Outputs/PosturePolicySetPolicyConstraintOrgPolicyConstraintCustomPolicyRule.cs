// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.SecurityPosture.Outputs
{

    [OutputType]
    public sealed class PosturePolicySetPolicyConstraintOrgPolicyConstraintCustomPolicyRule
    {
        /// <summary>
        /// Setting this to true means that all values are allowed. This field can be set only in policies for list constraints.
        /// </summary>
        public readonly bool? AllowAll;
        /// <summary>
        /// Setting this to true means that all values are denied. This field can be set only in policies for list constraints.
        /// </summary>
        public readonly bool? DenyAll;
        /// <summary>
        /// If `true`, then the policy is enforced. If `false`, then any configuration is acceptable.
        /// This field can be set only in policies for boolean constraints.
        /// </summary>
        public readonly bool? Enforce;
        /// <summary>
        /// Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language.
        /// This page details the objects and attributes that are used to the build the CEL expressions for
        /// custom access levels - https://cloud.google.com/access-context-manager/docs/custom-access-level-spec.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.PosturePolicySetPolicyConstraintOrgPolicyConstraintCustomPolicyRuleExpr? Expr;
        /// <summary>
        /// List of values to be used for this policy rule. This field can be set only in policies for list constraints.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.PosturePolicySetPolicyConstraintOrgPolicyConstraintCustomPolicyRuleValues? Values;

        [OutputConstructor]
        private PosturePolicySetPolicyConstraintOrgPolicyConstraintCustomPolicyRule(
            bool? allowAll,

            bool? denyAll,

            bool? enforce,

            Outputs.PosturePolicySetPolicyConstraintOrgPolicyConstraintCustomPolicyRuleExpr? expr,

            Outputs.PosturePolicySetPolicyConstraintOrgPolicyConstraintCustomPolicyRuleValues? values)
        {
            AllowAll = allowAll;
            DenyAll = denyAll;
            Enforce = enforce;
            Expr = expr;
            Values = values;
        }
    }
}
