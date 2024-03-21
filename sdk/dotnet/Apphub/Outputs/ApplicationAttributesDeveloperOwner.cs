// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Apphub.Outputs
{

    [OutputType]
    public sealed class ApplicationAttributesDeveloperOwner
    {
        /// <summary>
        /// Optional. Contact's name.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// Required. Email address of the contacts.
        /// </summary>
        public readonly string Email;

        [OutputConstructor]
        private ApplicationAttributesDeveloperOwner(
            string? displayName,

            string email)
        {
            DisplayName = displayName;
            Email = email;
        }
    }
}
