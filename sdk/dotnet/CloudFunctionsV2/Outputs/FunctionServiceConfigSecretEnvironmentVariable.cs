// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudFunctionsV2.Outputs
{

    [OutputType]
    public sealed class FunctionServiceConfigSecretEnvironmentVariable
    {
        /// <summary>
        /// Name of the environment variable.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Project identifier (preferrably project number but can also be the project ID) of the project that contains the secret. If not set, it will be populated with the function's project assuming that the secret exists in the same project as of the function.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// Name of the secret in secret manager (not the full resource name).
        /// </summary>
        public readonly string Secret;
        /// <summary>
        /// Version of the secret (version number or the string 'latest'). It is preferable to use latest version with secret volumes as secret value changes are reflected immediately.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private FunctionServiceConfigSecretEnvironmentVariable(
            string key,

            string projectId,

            string secret,

            string version)
        {
            Key = key;
            ProjectId = projectId;
            Secret = secret;
            Version = version;
        }
    }
}
