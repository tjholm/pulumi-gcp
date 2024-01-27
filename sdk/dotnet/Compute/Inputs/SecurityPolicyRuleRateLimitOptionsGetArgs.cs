// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class SecurityPolicyRuleRateLimitOptionsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Can only be specified if the `action` for the rule is `rate_based_ban`.
        /// If specified, determines the time (in seconds) the traffic will continue to be banned by the rate limit after the rate falls below the threshold.
        /// </summary>
        [Input("banDurationSec")]
        public Input<int>? BanDurationSec { get; set; }

        /// <summary>
        /// Can only be specified if the `action` for the rule is `rate_based_ban`.
        /// If specified, the key will be banned for the configured `ban_duration_sec` when the number of requests that exceed the `rate_limit_threshold` also
        /// exceed this `ban_threshold`. Structure is documented below.
        /// </summary>
        [Input("banThreshold")]
        public Input<Inputs.SecurityPolicyRuleRateLimitOptionsBanThresholdGetArgs>? BanThreshold { get; set; }

        /// <summary>
        /// Action to take for requests that are under the configured rate limit threshold. Valid option is `allow` only.
        /// </summary>
        [Input("conformAction", required: true)]
        public Input<string> ConformAction { get; set; } = null!;

        /// <summary>
        /// Determines the key to enforce the rate_limit_threshold on. If not specified, defaults to `ALL`.
        /// </summary>
        [Input("enforceOnKey")]
        public Input<string>? EnforceOnKey { get; set; }

        [Input("enforceOnKeyConfigs")]
        private InputList<Inputs.SecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfigGetArgs>? _enforceOnKeyConfigs;

        /// <summary>
        /// If specified, any combination of values of enforce_on_key_type/enforce_on_key_name is treated as the key on which rate limit threshold/action is enforced. You can specify up to 3 enforce_on_key_configs. If `enforce_on_key_configs` is specified, `enforce_on_key` must be set to an empty string. Structure is documented below.
        /// 
        /// **Note:** To avoid the conflict between `enforce_on_key` and `enforce_on_key_configs`, the field `enforce_on_key` needs to be set to an empty string.
        /// </summary>
        public InputList<Inputs.SecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfigGetArgs> EnforceOnKeyConfigs
        {
            get => _enforceOnKeyConfigs ?? (_enforceOnKeyConfigs = new InputList<Inputs.SecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfigGetArgs>());
            set => _enforceOnKeyConfigs = value;
        }

        /// <summary>
        /// Rate limit key name applicable only for the following key types:
        /// </summary>
        [Input("enforceOnKeyName")]
        public Input<string>? EnforceOnKeyName { get; set; }

        /// <summary>
        /// When a request is denied, returns the HTTP response code specified.
        /// Valid options are `deny()` where valid values for status are 403, 404, 429, and 502.
        /// </summary>
        [Input("exceedAction", required: true)]
        public Input<string> ExceedAction { get; set; } = null!;

        /// <summary>
        /// Parameters defining the redirect action that is used as the exceed action. Cannot be specified if the exceed action is not redirect. Structure is documented below.
        /// 
        /// &lt;a name="nested_threshold"&gt;&lt;/a&gt;The `{ban/rate_limit}_threshold` block supports:
        /// </summary>
        [Input("exceedRedirectOptions")]
        public Input<Inputs.SecurityPolicyRuleRateLimitOptionsExceedRedirectOptionsGetArgs>? ExceedRedirectOptions { get; set; }

        /// <summary>
        /// Threshold at which to begin ratelimiting. Structure is documented below.
        /// </summary>
        [Input("rateLimitThreshold", required: true)]
        public Input<Inputs.SecurityPolicyRuleRateLimitOptionsRateLimitThresholdGetArgs> RateLimitThreshold { get; set; } = null!;

        public SecurityPolicyRuleRateLimitOptionsGetArgs()
        {
        }
        public static new SecurityPolicyRuleRateLimitOptionsGetArgs Empty => new SecurityPolicyRuleRateLimitOptionsGetArgs();
    }
}
