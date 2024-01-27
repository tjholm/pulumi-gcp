// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class SecurityPolicyAdvancedOptionsConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom configuration to apply the JSON parsing. Only applicable when
        /// `json_parsing` is set to `STANDARD`. Structure is documented below.
        /// </summary>
        [Input("jsonCustomConfig")]
        public Input<Inputs.SecurityPolicyAdvancedOptionsConfigJsonCustomConfigArgs>? JsonCustomConfig { get; set; }

        /// <summary>
        /// Whether or not to JSON parse the payload body. Defaults to `DISABLED`.
        /// </summary>
        [Input("jsonParsing")]
        public Input<string>? JsonParsing { get; set; }

        /// <summary>
        /// Log level to use. Defaults to `NORMAL`.
        /// </summary>
        [Input("logLevel")]
        public Input<string>? LogLevel { get; set; }

        [Input("userIpRequestHeaders")]
        private InputList<string>? _userIpRequestHeaders;

        /// <summary>
        /// An optional list of case-insensitive request header names to use for resolving the callers client IP address.
        /// </summary>
        public InputList<string> UserIpRequestHeaders
        {
            get => _userIpRequestHeaders ?? (_userIpRequestHeaders = new InputList<string>());
            set => _userIpRequestHeaders = value;
        }

        public SecurityPolicyAdvancedOptionsConfigArgs()
        {
        }
        public static new SecurityPolicyAdvancedOptionsConfigArgs Empty => new SecurityPolicyAdvancedOptionsConfigArgs();
    }
}
