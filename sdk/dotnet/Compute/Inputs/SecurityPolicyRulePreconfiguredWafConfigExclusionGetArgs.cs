// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class SecurityPolicyRulePreconfiguredWafConfigExclusionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("requestCookies")]
        private InputList<Inputs.SecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookyGetArgs>? _requestCookies;

        /// <summary>
        /// Request cookie whose value will be excluded from inspection during preconfigured WAF evaluation. Structure is documented below.
        /// </summary>
        public InputList<Inputs.SecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookyGetArgs> RequestCookies
        {
            get => _requestCookies ?? (_requestCookies = new InputList<Inputs.SecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookyGetArgs>());
            set => _requestCookies = value;
        }

        [Input("requestHeaders")]
        private InputList<Inputs.SecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeaderGetArgs>? _requestHeaders;

        /// <summary>
        /// Request header whose value will be excluded from inspection during preconfigured WAF evaluation. Structure is documented below.
        /// </summary>
        public InputList<Inputs.SecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeaderGetArgs> RequestHeaders
        {
            get => _requestHeaders ?? (_requestHeaders = new InputList<Inputs.SecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeaderGetArgs>());
            set => _requestHeaders = value;
        }

        [Input("requestQueryParams")]
        private InputList<Inputs.SecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamGetArgs>? _requestQueryParams;

        /// <summary>
        /// Request query parameter whose value will be excluded from inspection during preconfigured WAF evaluation. Note that the parameter can be in the query string or in the POST body. Structure is documented below.
        /// </summary>
        public InputList<Inputs.SecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamGetArgs> RequestQueryParams
        {
            get => _requestQueryParams ?? (_requestQueryParams = new InputList<Inputs.SecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamGetArgs>());
            set => _requestQueryParams = value;
        }

        [Input("requestUris")]
        private InputList<Inputs.SecurityPolicyRulePreconfiguredWafConfigExclusionRequestUriGetArgs>? _requestUris;

        /// <summary>
        /// Request URI from the request line to be excluded from inspection during preconfigured WAF evaluation. When specifying this field, the query or fragment part should be excluded. Structure is documented below.
        /// </summary>
        public InputList<Inputs.SecurityPolicyRulePreconfiguredWafConfigExclusionRequestUriGetArgs> RequestUris
        {
            get => _requestUris ?? (_requestUris = new InputList<Inputs.SecurityPolicyRulePreconfiguredWafConfigExclusionRequestUriGetArgs>());
            set => _requestUris = value;
        }

        [Input("targetRuleIds")]
        private InputList<string>? _targetRuleIds;

        /// <summary>
        /// A list of target rule IDs under the WAF rule set to apply the preconfigured WAF exclusion. If omitted, it refers to all the rule IDs under the WAF rule set.
        /// 
        /// &lt;a name="nested_field_params"&gt;&lt;/a&gt;The `request_header`, `request_cookie`, `request_uri` and `request_query_param` blocks support:
        /// </summary>
        public InputList<string> TargetRuleIds
        {
            get => _targetRuleIds ?? (_targetRuleIds = new InputList<string>());
            set => _targetRuleIds = value;
        }

        /// <summary>
        /// Target WAF rule set to apply the preconfigured WAF exclusion.
        /// </summary>
        [Input("targetRuleSet", required: true)]
        public Input<string> TargetRuleSet { get; set; } = null!;

        public SecurityPolicyRulePreconfiguredWafConfigExclusionGetArgs()
        {
        }
        public static new SecurityPolicyRulePreconfiguredWafConfigExclusionGetArgs Empty => new SecurityPolicyRulePreconfiguredWafConfigExclusionGetArgs();
    }
}
