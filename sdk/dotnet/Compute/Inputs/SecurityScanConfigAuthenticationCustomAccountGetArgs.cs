// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class SecurityScanConfigAuthenticationCustomAccountGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The login form URL of the website.
        /// </summary>
        [Input("loginUrl", required: true)]
        public Input<string> LoginUrl { get; set; } = null!;

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// The password of the custom account. The credential is stored encrypted
        /// in GCP.
        /// **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The user name of the custom account.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public SecurityScanConfigAuthenticationCustomAccountGetArgs()
        {
        }
        public static new SecurityScanConfigAuthenticationCustomAccountGetArgs Empty => new SecurityScanConfigAuthenticationCustomAccountGetArgs();
    }
}
