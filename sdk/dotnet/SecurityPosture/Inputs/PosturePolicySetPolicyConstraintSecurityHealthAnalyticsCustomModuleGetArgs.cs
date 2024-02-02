// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.SecurityPosture.Inputs
{

    public sealed class PosturePolicySetPolicyConstraintSecurityHealthAnalyticsCustomModuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom module details.
        /// Structure is documented below.
        /// </summary>
        [Input("config", required: true)]
        public Input<Inputs.PosturePolicySetPolicyConstraintSecurityHealthAnalyticsCustomModuleConfigGetArgs> Config { get; set; } = null!;

        /// <summary>
        /// The display name of the Security Health Analytics custom module. This
        /// display name becomes the finding category for all findings that are
        /// returned by this custom module.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// (Output)
        /// A server generated id of custom module.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The state of enablement for the module at its level of the resource hierarchy.
        /// Possible values are: `ENABLEMENT_STATE_UNSPECIFIED`, `ENABLED`, `DISABLED`.
        /// </summary>
        [Input("moduleEnablementState")]
        public Input<string>? ModuleEnablementState { get; set; }

        public PosturePolicySetPolicyConstraintSecurityHealthAnalyticsCustomModuleGetArgs()
        {
        }
        public static new PosturePolicySetPolicyConstraintSecurityHealthAnalyticsCustomModuleGetArgs Empty => new PosturePolicySetPolicyConstraintSecurityHealthAnalyticsCustomModuleGetArgs();
    }
}
