// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataform.Outputs
{

    [OutputType]
    public sealed class RepositoryGitRemoteSettings
    {
        /// <summary>
        /// The name of the Secret Manager secret version to use as an authentication token for Git operations. Must be in the format projects/*/secrets/*/versions/*.
        /// </summary>
        public readonly string AuthenticationTokenSecretVersion;
        /// <summary>
        /// The Git remote's default branch name.
        /// </summary>
        public readonly string DefaultBranch;
        /// <summary>
        /// -
        /// Indicates the status of the Git access token. https://cloud.google.com/dataform/reference/rest/v1beta1/projects.locations.repositories#TokenStatus
        /// </summary>
        public readonly string? TokenStatus;
        /// <summary>
        /// The Git remote's URL.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private RepositoryGitRemoteSettings(
            string authenticationTokenSecretVersion,

            string defaultBranch,

            string? tokenStatus,

            string url)
        {
            AuthenticationTokenSecretVersion = authenticationTokenSecretVersion;
            DefaultBranch = defaultBranch;
            TokenStatus = tokenStatus;
            Url = url;
        }
    }
}
