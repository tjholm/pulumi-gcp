// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.SecurityCenter.Inputs
{

    public sealed class SourceIamMemberConditionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the source (max of 1024 characters).
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public SourceIamMemberConditionGetArgs()
        {
        }
        public static new SourceIamMemberConditionGetArgs Empty => new SourceIamMemberConditionGetArgs();
    }
}
