// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.ArtifactRegistry.Outputs
{

    [OutputType]
    public sealed class RepositoryCleanupPolicyMostRecentVersions
    {
        /// <summary>
        /// Minimum number of versions to keep.
        /// </summary>
        public readonly int? KeepCount;
        /// <summary>
        /// Match versions by package prefix. Applied on any prefix match.
        /// </summary>
        public readonly ImmutableArray<string> PackageNamePrefixes;

        [OutputConstructor]
        private RepositoryCleanupPolicyMostRecentVersions(
            int? keepCount,

            ImmutableArray<string> packageNamePrefixes)
        {
            KeepCount = keepCount;
            PackageNamePrefixes = packageNamePrefixes;
        }
    }
}
