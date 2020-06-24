// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudBuild.Outputs
{

    [OutputType]
    public sealed class TriggerGithubPush
    {
        /// <summary>
        /// Regex of branches to match.  Specify only one of branch or tag.
        /// </summary>
        public readonly string? Branch;
        /// <summary>
        /// When true, only trigger a build if the revision regex does NOT match the git_ref regex.
        /// </summary>
        public readonly bool? InvertRegex;
        /// <summary>
        /// Regex of tags to match.  Specify only one of branch or tag.
        /// </summary>
        public readonly string? Tag;

        [OutputConstructor]
        private TriggerGithubPush(
            string? branch,

            bool? invertRegex,

            string? tag)
        {
            Branch = branch;
            InvertRegex = invertRegex;
            Tag = tag;
        }
    }
}
