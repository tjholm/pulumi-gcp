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
    public sealed class PosturePolicySetPolicyConstraintSecurityHealthAnalyticsModule
    {
        /// <summary>
        /// The state of enablement for the module at its level of the resource hierarchy.
        /// Possible values are: `ENABLEMENT_STATE_UNSPECIFIED`, `ENABLED`, `DISABLED`.
        /// </summary>
        public readonly string? ModuleEnablementState;
        /// <summary>
        /// The name of the module eg: BIGQUERY_TABLE_CMEK_DISABLED.
        /// </summary>
        public readonly string ModuleName;

        [OutputConstructor]
        private PosturePolicySetPolicyConstraintSecurityHealthAnalyticsModule(
            string? moduleEnablementState,

            string moduleName)
        {
            ModuleEnablementState = moduleEnablementState;
            ModuleName = moduleName;
        }
    }
}
