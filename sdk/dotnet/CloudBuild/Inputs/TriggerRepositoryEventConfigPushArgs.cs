// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudBuild.Inputs
{

    public sealed class TriggerRepositoryEventConfigPushArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Regex of branches to match.  Specify only one of branch or tag.
        /// </summary>
        [Input("branch")]
        public Input<string>? Branch { get; set; }

        /// <summary>
        /// When true, only trigger a build if the revision regex does NOT match the git_ref regex.
        /// </summary>
        [Input("invertRegex")]
        public Input<bool>? InvertRegex { get; set; }

        /// <summary>
        /// Regex of tags to match.  Specify only one of branch or tag.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        public TriggerRepositoryEventConfigPushArgs()
        {
        }
        public static new TriggerRepositoryEventConfigPushArgs Empty => new TriggerRepositoryEventConfigPushArgs();
    }
}
