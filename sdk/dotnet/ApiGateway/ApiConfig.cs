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
    /// An API Configuration is an association of an API Controller Config and a Gateway Config
    /// 
    /// To get more information about ApiConfig, see:
    /// 
    /// * [API documentation](https://cloud.google.com/api-gateway/docs/reference/rest/v1beta/projects.locations.apis.configs)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/api-gateway/docs/creating-api-config)
    /// 
    /// ## Example Usage
    /// ### Apigateway Api Config Basic
    /// 
    /// ```csharp
    /// using System;
    /// using System.IO;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    /// 	private static string ReadFileBase64(string path) {
    /// 		return Convert.ToBase64String(Encoding.UTF8.GetBytes(File.ReadAllText(path)))
    /// 	}
    /// 
    ///     public MyStack()
    ///     {
    ///         var apiCfgApi = new Gcp.ApiGateway.Api("apiCfgApi", new Gcp.ApiGateway.ApiArgs
    ///         {
    ///             ApiId = "api-cfg",
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///         var apiCfgApiConfig = new Gcp.ApiGateway.ApiConfig("apiCfgApiConfig", new Gcp.ApiGateway.ApiConfigArgs
    ///         {
    ///             Api = apiCfgApi.ApiId,
    ///             ApiConfigId = "cfg",
    ///             OpenapiDocuments = 
    ///             {
    ///                 new Gcp.ApiGateway.Inputs.ApiConfigOpenapiDocumentArgs
    ///                 {
    ///                     Document = new Gcp.ApiGateway.Inputs.ApiConfigOpenapiDocumentDocumentArgs
    ///                     {
    ///                         Path = "spec.yaml",
    ///                         Contents = ReadFileBase64("test-fixtures/apigateway/openapi.yaml"),
    ///                     },
    ///                 },
    ///             },
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
    /// ApiConfig can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:apigateway/apiConfig:ApiConfig default projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:apigateway/apiConfig:ApiConfig default {{project}}/{{api}}/{{api_config_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:apigateway/apiConfig:ApiConfig default {{api}}/{{api_config_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:apigateway/apiConfig:ApiConfig")]
    public partial class ApiConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// The API to attach the config to.
        /// </summary>
        [Output("api")]
        public Output<string> Api { get; private set; } = null!;

        /// <summary>
        /// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
        /// </summary>
        [Output("apiConfigId")]
        public Output<string> ApiConfigId { get; private set; } = null!;

        /// <summary>
        /// Creates a unique name beginning with the
        /// specified prefix. If this and api_config_id are unspecified, a random value is chosen for the name.
        /// </summary>
        [Output("apiConfigIdPrefix")]
        public Output<string> ApiConfigIdPrefix { get; private set; } = null!;

        /// <summary>
        /// A user-visible name for the API.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Immutable. Gateway specific configuration.
        /// If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
        /// Structure is documented below.
        /// </summary>
        [Output("gatewayConfig")]
        public Output<Outputs.ApiConfigGatewayConfig?> GatewayConfig { get; private set; } = null!;

        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The resource name of the API Config.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An OpenAPI Specification Document describing an API.
        /// Structure is documented below.
        /// </summary>
        [Output("openapiDocuments")]
        public Output<ImmutableArray<Outputs.ApiConfigOpenapiDocument>> OpenapiDocuments { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
        /// </summary>
        [Output("serviceConfigId")]
        public Output<string> ServiceConfigId { get; private set; } = null!;


        /// <summary>
        /// Create a ApiConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiConfig(string name, ApiConfigArgs args, CustomResourceOptions? options = null)
            : base("gcp:apigateway/apiConfig:ApiConfig", name, args ?? new ApiConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApiConfig(string name, Input<string> id, ApiConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:apigateway/apiConfig:ApiConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApiConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiConfig Get(string name, Input<string> id, ApiConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new ApiConfig(name, id, state, options);
        }
    }

    public sealed class ApiConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API to attach the config to.
        /// </summary>
        [Input("api", required: true)]
        public Input<string> Api { get; set; } = null!;

        /// <summary>
        /// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
        /// </summary>
        [Input("apiConfigId")]
        public Input<string>? ApiConfigId { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the
        /// specified prefix. If this and api_config_id are unspecified, a random value is chosen for the name.
        /// </summary>
        [Input("apiConfigIdPrefix")]
        public Input<string>? ApiConfigIdPrefix { get; set; }

        /// <summary>
        /// A user-visible name for the API.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Immutable. Gateway specific configuration.
        /// If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
        /// Structure is documented below.
        /// </summary>
        [Input("gatewayConfig")]
        public Input<Inputs.ApiConfigGatewayConfigArgs>? GatewayConfig { get; set; }

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

        [Input("openapiDocuments", required: true)]
        private InputList<Inputs.ApiConfigOpenapiDocumentArgs>? _openapiDocuments;

        /// <summary>
        /// An OpenAPI Specification Document describing an API.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ApiConfigOpenapiDocumentArgs> OpenapiDocuments
        {
            get => _openapiDocuments ?? (_openapiDocuments = new InputList<Inputs.ApiConfigOpenapiDocumentArgs>());
            set => _openapiDocuments = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ApiConfigArgs()
        {
        }
    }

    public sealed class ApiConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API to attach the config to.
        /// </summary>
        [Input("api")]
        public Input<string>? Api { get; set; }

        /// <summary>
        /// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
        /// </summary>
        [Input("apiConfigId")]
        public Input<string>? ApiConfigId { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the
        /// specified prefix. If this and api_config_id are unspecified, a random value is chosen for the name.
        /// </summary>
        [Input("apiConfigIdPrefix")]
        public Input<string>? ApiConfigIdPrefix { get; set; }

        /// <summary>
        /// A user-visible name for the API.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Immutable. Gateway specific configuration.
        /// If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
        /// Structure is documented below.
        /// </summary>
        [Input("gatewayConfig")]
        public Input<Inputs.ApiConfigGatewayConfigGetArgs>? GatewayConfig { get; set; }

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
        /// The resource name of the API Config.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("openapiDocuments")]
        private InputList<Inputs.ApiConfigOpenapiDocumentGetArgs>? _openapiDocuments;

        /// <summary>
        /// An OpenAPI Specification Document describing an API.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ApiConfigOpenapiDocumentGetArgs> OpenapiDocuments
        {
            get => _openapiDocuments ?? (_openapiDocuments = new InputList<Inputs.ApiConfigOpenapiDocumentGetArgs>());
            set => _openapiDocuments = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
        /// </summary>
        [Input("serviceConfigId")]
        public Input<string>? ServiceConfigId { get; set; }

        public ApiConfigState()
        {
        }
    }
}
