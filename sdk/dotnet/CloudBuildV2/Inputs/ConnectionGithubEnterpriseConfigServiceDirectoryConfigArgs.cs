// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudBuildV2.Inputs
{

    public sealed class ConnectionGithubEnterpriseConfigServiceDirectoryConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The Service Directory service name. Format: projects/{project}/locations/{location}/namespaces/{namespace}/services/{service}.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public ConnectionGithubEnterpriseConfigServiceDirectoryConfigArgs()
        {
        }
        public static new ConnectionGithubEnterpriseConfigServiceDirectoryConfigArgs Empty => new ConnectionGithubEnterpriseConfigServiceDirectoryConfigArgs();
    }
}
