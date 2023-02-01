// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudRun.Outputs
{

    [OutputType]
    public sealed class DomainMappingMetadata
    {
        public readonly ImmutableDictionary<string, string>? Annotations;
        /// <summary>
        /// A sequence number representing a specific generation of the desired state.
        /// </summary>
        public readonly int? Generation;
        /// <summary>
        /// Map of string keys and values that can be used to organize and categorize
        /// (scope and select) objects. May match selectors of replication controllers
        /// and routes.
        /// More info: http://kubernetes.io/docs/user-guide/labels
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Labels;
        /// <summary>
        /// In Cloud Run the namespace must be equal to either the
        /// project ID or project number.
        /// </summary>
        public readonly string Namespace;
        /// <summary>
        /// An opaque value that represents the internal version of this object that
        /// can be used by clients to determine when objects have changed. May be used
        /// for optimistic concurrency, change detection, and the watch operation on a
        /// resource or set of resources. They may only be valid for a
        /// particular resource or set of resources.
        /// More info:
        /// https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
        /// </summary>
        public readonly string? ResourceVersion;
        /// <summary>
        /// SelfLink is a URL representing this object.
        /// </summary>
        public readonly string? SelfLink;
        /// <summary>
        /// UID is a unique id generated by the server on successful creation of a resource and is not
        /// allowed to change on PUT operations.
        /// More info: http://kubernetes.io/docs/user-guide/identifiers#uids
        /// </summary>
        public readonly string? Uid;

        [OutputConstructor]
        private DomainMappingMetadata(
            ImmutableDictionary<string, string>? annotations,

            int? generation,

            ImmutableDictionary<string, string>? labels,

            string @namespace,

            string? resourceVersion,

            string? selfLink,

            string? uid)
        {
            Annotations = annotations;
            Generation = generation;
            Labels = labels;
            Namespace = @namespace;
            ResourceVersion = resourceVersion;
            SelfLink = selfLink;
            Uid = uid;
        }
    }
}
