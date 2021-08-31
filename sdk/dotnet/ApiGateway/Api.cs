// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.ApiGateway
{
    /// <summary>
    /// A consumable API that can be used by multiple Gateways.
    /// 
    /// To get more information about Api, see:
    /// 
    /// * [API documentation](https://cloud.google.com/api-gateway/docs/reference/rest/v1beta/projects.locations.apis)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/api-gateway/docs/quickstart)
    /// 
    /// ## Example Usage
    /// ### Apigateway Api Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var api = new Gcp.ApiGateway.Api("api", new Gcp.ApiGateway.ApiArgs
    ///         {
    ///             ApiId = "api",
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Api can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:apigateway/api:Api default projects/{{project}}/locations/global/apis/{{api_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:apigateway/api:Api default {{project}}/{{api_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:apigateway/api:Api default {{api_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:apigateway/api:Api")]
    public partial class Api : Pulumi.CustomResource
    {
        /// <summary>
        /// Identifier to assign to the API. Must be unique within scope of the parent resource(project)
        /// </summary>
        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// A user-visible name for the API.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
        /// If not specified, a new Service will automatically be created in the same project as this API.
        /// </summary>
        [Output("managedService")]
        public Output<string> ManagedService { get; private set; } = null!;

        /// <summary>
        /// The resource name of the API. Format 'projects/{{project}}/locations/global/apis/{{apiId}}'
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a Api resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Api(string name, ApiArgs args, CustomResourceOptions? options = null)
            : base("gcp:apigateway/api:Api", name, args ?? new ApiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Api(string name, Input<string> id, ApiState? state = null, CustomResourceOptions? options = null)
            : base("gcp:apigateway/api:Api", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Api resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Api Get(string name, Input<string> id, ApiState? state = null, CustomResourceOptions? options = null)
        {
            return new Api(name, id, state, options);
        }
    }

    public sealed class ApiArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier to assign to the API. Must be unique within scope of the parent resource(project)
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// A user-visible name for the API.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
        /// If not specified, a new Service will automatically be created in the same project as this API.
        /// </summary>
        [Input("managedService")]
        public Input<string>? ManagedService { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ApiArgs()
        {
        }
    }

    public sealed class ApiState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier to assign to the API. Must be unique within scope of the parent resource(project)
        /// </summary>
        [Input("apiId")]
        public Input<string>? ApiId { get; set; }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// A user-visible name for the API.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
        /// If not specified, a new Service will automatically be created in the same project as this API.
        /// </summary>
        [Input("managedService")]
        public Input<string>? ManagedService { get; set; }

        /// <summary>
        /// The resource name of the API. Format 'projects/{{project}}/locations/global/apis/{{apiId}}'
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ApiState()
        {
        }
    }
}
