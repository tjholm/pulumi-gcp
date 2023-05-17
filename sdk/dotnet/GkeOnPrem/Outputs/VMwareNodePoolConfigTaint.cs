// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeOnPrem.Outputs
{

    [OutputType]
    public sealed class VMwareNodePoolConfigTaint
    {
        /// <summary>
        /// Available taint effects.
        /// Possible values are: `EFFECT_UNSPECIFIED`, `NO_SCHEDULE`, `PREFER_NO_SCHEDULE`, `NO_EXECUTE`.
        /// </summary>
        public readonly string? Effect;
        /// <summary>
        /// Key associated with the effect.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Value associated with the effect.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private VMwareNodePoolConfigTaint(
            string? effect,

            string key,

            string value)
        {
            Effect = effect;
            Key = key;
            Value = value;
        }
    }
}
