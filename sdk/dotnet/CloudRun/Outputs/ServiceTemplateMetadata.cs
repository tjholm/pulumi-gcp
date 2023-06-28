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
    public sealed class ServiceTemplateMetadata
    {
        /// <summary>
        /// Annotations is a key value map stored with a resource that
        /// may be set by external tools to store and retrieve arbitrary metadata. More
        /// info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations
        /// **Note**: The Cloud Run API may add additional annotations that were not provided in your config.
        /// If the provider plan shows a diff where a server-side annotation is added, you can add it to your config
        /// or apply the lifecycle.ignore_changes rule to the metadata.0.annotations field.
        /// Annotations with `run.googleapis.com/` and `autoscaling.knative.dev` are restricted. Use the following annotation
        /// keys to configure features on a Service:
        /// - `run.googleapis.com/binary-authorization-breakglass` sets the [Binary Authorization breakglass](https://cloud.google.com/sdk/gcloud/reference/run/deploy#--breakglass).
        /// - `run.googleapis.com/binary-authorization` sets the [Binary Authorization](https://cloud.google.com/sdk/gcloud/reference/run/deploy#--binary-authorization).
        /// - `run.googleapis.com/client-name` sets the client name calling the Cloud Run API.
        /// - `run.googleapis.com/custom-audiences` sets the [custom audiences](https://cloud.google.com/sdk/gcloud/reference/alpha/run/deploy#--add-custom-audiences)
        /// that can be used in the audience field of ID token for authenticated requests.
        /// - `run.googleapis.com/description` sets a user defined description for the Service.
        /// - `run.googleapis.com/ingress` sets the [ingress settings](https://cloud.google.com/sdk/gcloud/reference/run/deploy#--ingress)
        /// for the Service. For example, `"run.googleapis.com/ingress" = "all"`.
        /// - `run.googleapis.com/launch-stage` sets the [launch stage](https://cloud.google.com/run/docs/troubleshooting#launch-stage-validation)
        /// when a preview feature is used. For example, `"run.googleapis.com/launch-stage": "BETA"`
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Annotations;
        /// <summary>
        /// (Output)
        /// A sequence number representing a specific generation of the desired state.
        /// </summary>
        public readonly int? Generation;
        /// <summary>
        /// Map of string keys and values that can be used to organize and categorize
        /// (scope and select) objects. May match selectors of replication controllers
        /// and routes.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Labels;
        /// <summary>
        /// Name must be unique within a Google Cloud project and region.
        /// Is required when creating resources. Name is primarily intended
        /// for creation idempotence and configuration definition. Cannot be updated.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// In Cloud Run the namespace must be equal to either the
        /// project ID or project number.
        /// </summary>
        public readonly string? Namespace;
        /// <summary>
        /// (Output)
        /// An opaque value that represents the internal version of this object that
        /// can be used by clients to determine when objects have changed. May be used
        /// for optimistic concurrency, change detection, and the watch operation on a
        /// resource or set of resources. They may only be valid for a
        /// particular resource or set of resources.
        /// </summary>
        public readonly string? ResourceVersion;
        /// <summary>
        /// (Output)
        /// SelfLink is a URL representing this object.
        /// </summary>
        public readonly string? SelfLink;
        /// <summary>
        /// (Output)
        /// UID is a unique id generated by the server on successful creation of a resource and is not
        /// allowed to change on PUT operations.
        /// </summary>
        public readonly string? Uid;

        [OutputConstructor]
        private ServiceTemplateMetadata(
            ImmutableDictionary<string, string>? annotations,

            int? generation,

            ImmutableDictionary<string, string>? labels,

            string? name,

            string? @namespace,

            string? resourceVersion,

            string? selfLink,

            string? uid)
        {
            Annotations = annotations;
            Generation = generation;
            Labels = labels;
            Name = name;
            Namespace = @namespace;
            ResourceVersion = resourceVersion;
            SelfLink = selfLink;
            Uid = uid;
        }
    }
}
