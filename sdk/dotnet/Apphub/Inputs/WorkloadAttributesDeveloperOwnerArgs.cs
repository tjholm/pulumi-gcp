// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Apphub.Inputs
{

    public sealed class WorkloadAttributesDeveloperOwnerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Contact's name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Email address of the contacts.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        public WorkloadAttributesDeveloperOwnerArgs()
        {
        }
        public static new WorkloadAttributesDeveloperOwnerArgs Empty => new WorkloadAttributesDeveloperOwnerArgs();
    }
}
