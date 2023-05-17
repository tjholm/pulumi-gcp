// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DatabaseMigrationService.Inputs
{

    public sealed class ConnectionProfileCloudsqlSettingsIpConfigAuthorizedNetworkGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time when this access control entry expires in RFC 3339 format.
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        /// <summary>
        /// A label to identify this entry.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Input only. The time-to-leave of this access control entry.
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        /// <summary>
        /// The allowlisted value for the access control list.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ConnectionProfileCloudsqlSettingsIpConfigAuthorizedNetworkGetArgs()
        {
        }
        public static new ConnectionProfileCloudsqlSettingsIpConfigAuthorizedNetworkGetArgs Empty => new ConnectionProfileCloudsqlSettingsIpConfigAuthorizedNetworkGetArgs();
    }
}
