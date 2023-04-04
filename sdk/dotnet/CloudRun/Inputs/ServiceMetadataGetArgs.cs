// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudRun.Inputs
{

    public sealed class ServiceMetadataGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// Annotations is a key value map stored with a resource that
        /// may be set by external tools to store and retrieve arbitrary metadata. More
        /// info: http://kubernetes.io/docs/user-guide/annotations
        /// **Note**: The Cloud Run API may add additional annotations that were not provided in your config.
        /// If the provider plan shows a diff where a server-side annotation is added, you can add it to your config
        /// or apply the lifecycle.ignore_changes rule to the metadata.0.annotations field.
        /// Cloud Run (fully managed) uses the following annotation keys to configure features on a Service:
        /// - `run.googleapis.com/ingress` sets the [ingress settings](https://cloud.google.com/sdk/gcloud/reference/run/deploy#--ingress)
        /// for the Service. For example, `"run.googleapis.com/ingress" = "all"`.
        /// </summary>
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// (Output)
        /// A sequence number representing a specific generation of the desired state.
        /// </summary>
        [Input("generation")]
        public Input<int>? Generation { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Map of string keys and values that can be used to organize and categorize
        /// (scope and select) objects. May match selectors of replication controllers
        /// and routes.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// In Cloud Run the namespace must be equal to either the
        /// project ID or project number.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// (Output)
        /// An opaque value that represents the internal version of this object that
        /// can be used by clients to determine when objects have changed. May be used
        /// for optimistic concurrency, change detection, and the watch operation on a
        /// resource or set of resources. They may only be valid for a
        /// particular resource or set of resources.
        /// </summary>
        [Input("resourceVersion")]
        public Input<string>? ResourceVersion { get; set; }

        /// <summary>
        /// (Output)
        /// SelfLink is a URL representing this object.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// (Output)
        /// UID is a unique id generated by the server on successful creation of a resource and is not
        /// allowed to change on PUT operations.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        public ServiceMetadataGetArgs()
        {
        }
        public static new ServiceMetadataGetArgs Empty => new ServiceMetadataGetArgs();
    }
}
