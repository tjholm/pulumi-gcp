// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudBuild.Inputs
{

    public sealed class TriggerSourceToBuildGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The full resource name of the github enterprise config.
        /// Format: projects/{project}/locations/{location}/githubEnterpriseConfigs/{id}. projects/{project}/githubEnterpriseConfigs/{id}.
        /// </summary>
        [Input("githubEnterpriseConfig")]
        public Input<string>? GithubEnterpriseConfig { get; set; }

        /// <summary>
        /// The branch or tag to use. Must start with "refs/" (required).
        /// </summary>
        [Input("ref", required: true)]
        public Input<string> Ref { get; set; } = null!;

        /// <summary>
        /// The type of the repo, since it may not be explicit from the repo field (e.g from a URL).
        /// Values can be UNKNOWN, CLOUD_SOURCE_REPOSITORIES, GITHUB, BITBUCKET_SERVER
        /// Possible values are `UNKNOWN`, `CLOUD_SOURCE_REPOSITORIES`, `GITHUB`, and `BITBUCKET_SERVER`.
        /// </summary>
        [Input("repoType", required: true)]
        public Input<string> RepoType { get; set; } = null!;

        /// <summary>
        /// The URI of the repo (required).
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public TriggerSourceToBuildGetArgs()
        {
        }
        public static new TriggerSourceToBuildGetArgs Empty => new TriggerSourceToBuildGetArgs();
    }
}
