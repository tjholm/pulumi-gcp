// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Inputs
{

    public sealed class ApplicationIapArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional) Whether the serving infrastructure will authenticate and authorize all incoming requests. 
        /// (default is false)
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// OAuth2 client ID to use for the authentication flow.
        /// </summary>
        [Input("oauth2ClientId", required: true)]
        public Input<string> Oauth2ClientId { get; set; } = null!;

        [Input("oauth2ClientSecret", required: true)]
        private Input<string>? _oauth2ClientSecret;

        /// <summary>
        /// OAuth2 client secret to use for the authentication flow.
        /// The SHA-256 hash of the value is returned in the oauth2ClientSecretSha256 field.
        /// </summary>
        public Input<string>? Oauth2ClientSecret
        {
            get => _oauth2ClientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _oauth2ClientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("oauth2ClientSecretSha256")]
        private Input<string>? _oauth2ClientSecretSha256;

        /// <summary>
        /// Hex-encoded SHA-256 hash of the client secret.
        /// </summary>
        public Input<string>? Oauth2ClientSecretSha256
        {
            get => _oauth2ClientSecretSha256;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _oauth2ClientSecretSha256 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ApplicationIapArgs()
        {
        }
        public static new ApplicationIapArgs Empty => new ApplicationIapArgs();
    }
}
