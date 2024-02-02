// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.SecurityPosture.Inputs
{

    public sealed class PosturePolicySetPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("complianceStandards")]
        private InputList<Inputs.PosturePolicySetPolicyComplianceStandardGetArgs>? _complianceStandards;

        /// <summary>
        /// Mapping for policy to security standards and controls.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.PosturePolicySetPolicyComplianceStandardGetArgs> ComplianceStandards
        {
            get => _complianceStandards ?? (_complianceStandards = new InputList<Inputs.PosturePolicySetPolicyComplianceStandardGetArgs>());
            set => _complianceStandards = value;
        }

        /// <summary>
        /// Policy constraint definition.It can have the definition of one of following constraints: orgPolicyConstraint orgPolicyConstraintCustom securityHealthAnalyticsModule securityHealthAnalyticsCustomModule
        /// Structure is documented below.
        /// </summary>
        [Input("constraint", required: true)]
        public Input<Inputs.PosturePolicySetPolicyConstraintGetArgs> Constraint { get; set; } = null!;

        /// <summary>
        /// Description of the policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ID of the policy.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        public PosturePolicySetPolicyGetArgs()
        {
        }
        public static new PosturePolicySetPolicyGetArgs Empty => new PosturePolicySetPolicyGetArgs();
    }
}
