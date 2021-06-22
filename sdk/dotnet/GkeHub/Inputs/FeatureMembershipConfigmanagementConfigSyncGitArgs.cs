// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeHub.Inputs
{

    public sealed class FeatureMembershipConfigmanagementConfigSyncGitArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// URL for the HTTPS proxy to be used when communicating with the Git repo.
        /// </summary>
        [Input("httpsProxy")]
        public Input<string>? HttpsProxy { get; set; }

        /// <summary>
        /// The path within the Git repository that represents the top level of the repo to sync. Default: the root directory of the repository.
        /// </summary>
        [Input("policyDir")]
        public Input<string>? PolicyDir { get; set; }

        /// <summary>
        /// Type of secret configured for access to the Git repo.
        /// </summary>
        [Input("secretType")]
        public Input<string>? SecretType { get; set; }

        /// <summary>
        /// The branch of the repository to sync from. Default: master.
        /// </summary>
        [Input("syncBranch")]
        public Input<string>? SyncBranch { get; set; }

        /// <summary>
        /// The URL of the Git repository to use as the source of truth.
        /// </summary>
        [Input("syncRepo")]
        public Input<string>? SyncRepo { get; set; }

        /// <summary>
        /// Git revision (tag or hash) to check out. Default HEAD.
        /// </summary>
        [Input("syncRev")]
        public Input<string>? SyncRev { get; set; }

        /// <summary>
        /// Period in seconds between consecutive syncs. Default: 15.
        /// </summary>
        [Input("syncWaitSecs")]
        public Input<string>? SyncWaitSecs { get; set; }

        public FeatureMembershipConfigmanagementConfigSyncGitArgs()
        {
        }
    }
}
