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
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// 	private static string ReadFileBase64(string path) {
    /// 		return Convert.ToBase64String(System.Text.Encoding.UTF8.GetBytes(File.ReadAllText(path)));
    /// 	}
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var apiCfgApi = new Gcp.ApiGateway.Api("apiCfgApi", new()
    ///     {
    ///         ApiId = "my-api",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var apiCfgApiConfig = new Gcp.ApiGateway.ApiConfig("apiCfgApiConfig", new()
    ///     {
    ///         Api = apiCfgApi.ApiId,
    ///         ApiConfigId = "my-config",
    ///         OpenapiDocuments = new[]
    ///         {
    ///             new Gcp.ApiGateway.Inputs.ApiConfigOpenapiDocumentArgs
    ///             {
    ///                 Document = new Gcp.ApiGateway.Inputs.ApiConfigOpenapiDocumentDocumentArgs
    ///                 {
    ///                     Path = "spec.yaml",
    ///                     Contents = ReadFileBase64("test-fixtures/openapi.yaml"),
    ///                 },
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
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
    public partial class ApiConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The API to attach the config to.
        /// 
        /// 
        /// - - -
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
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
        /// clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// Immutable. Gateway specific configuration.
        /// If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
        /// Structure is documented below.
        /// </summary>
        [Output("gatewayConfig")]
        public Output<Outputs.ApiConfigGatewayConfig?> GatewayConfig { get; private set; } = null!;

        /// <summary>
        /// gRPC service definition files. If specified, openapiDocuments must not be included.
        /// Structure is documented below.
        /// </summary>
        [Output("grpcServices")]
        public Output<ImmutableArray<Outputs.ApiConfigGrpcService>> GrpcServices { get; private set; } = null!;

        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents.
        /// If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
        /// Structure is documented below.
        /// </summary>
        [Output("managedServiceConfigs")]
        public Output<ImmutableArray<Outputs.ApiConfigManagedServiceConfig>> ManagedServiceConfigs { get; private set; } = null!;

        /// <summary>
        /// The resource name of the API Config.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
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
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        [Output("pulumiLabels")]
        public Output<ImmutableDictionary<string, string>> PulumiLabels { get; private set; } = null!;

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
                AdditionalSecretOutputs =
                {
                    "effectiveLabels",
                    "pulumiLabels",
                },
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

    public sealed class ApiConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API to attach the config to.
        /// 
        /// 
        /// - - -
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

        [Input("grpcServices")]
        private InputList<Inputs.ApiConfigGrpcServiceArgs>? _grpcServices;

        /// <summary>
        /// gRPC service definition files. If specified, openapiDocuments must not be included.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ApiConfigGrpcServiceArgs> GrpcServices
        {
            get => _grpcServices ?? (_grpcServices = new InputList<Inputs.ApiConfigGrpcServiceArgs>());
            set => _grpcServices = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("managedServiceConfigs")]
        private InputList<Inputs.ApiConfigManagedServiceConfigArgs>? _managedServiceConfigs;

        /// <summary>
        /// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents.
        /// If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ApiConfigManagedServiceConfigArgs> ManagedServiceConfigs
        {
            get => _managedServiceConfigs ?? (_managedServiceConfigs = new InputList<Inputs.ApiConfigManagedServiceConfigArgs>());
            set => _managedServiceConfigs = value;
        }

        [Input("openapiDocuments")]
        private InputList<Inputs.ApiConfigOpenapiDocumentArgs>? _openapiDocuments;

        /// <summary>
        /// OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
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
        public static new ApiConfigArgs Empty => new ApiConfigArgs();
    }

    public sealed class ApiConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API to attach the config to.
        /// 
        /// 
        /// - - -
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

        [Input("effectiveLabels")]
        private InputMap<string>? _effectiveLabels;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
        /// clients and services.
        /// </summary>
        public InputMap<string> EffectiveLabels
        {
            get => _effectiveLabels ?? (_effectiveLabels = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _effectiveLabels = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// Immutable. Gateway specific configuration.
        /// If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
        /// Structure is documented below.
        /// </summary>
        [Input("gatewayConfig")]
        public Input<Inputs.ApiConfigGatewayConfigGetArgs>? GatewayConfig { get; set; }

        [Input("grpcServices")]
        private InputList<Inputs.ApiConfigGrpcServiceGetArgs>? _grpcServices;

        /// <summary>
        /// gRPC service definition files. If specified, openapiDocuments must not be included.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ApiConfigGrpcServiceGetArgs> GrpcServices
        {
            get => _grpcServices ?? (_grpcServices = new InputList<Inputs.ApiConfigGrpcServiceGetArgs>());
            set => _grpcServices = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("managedServiceConfigs")]
        private InputList<Inputs.ApiConfigManagedServiceConfigGetArgs>? _managedServiceConfigs;

        /// <summary>
        /// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents.
        /// If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ApiConfigManagedServiceConfigGetArgs> ManagedServiceConfigs
        {
            get => _managedServiceConfigs ?? (_managedServiceConfigs = new InputList<Inputs.ApiConfigManagedServiceConfigGetArgs>());
            set => _managedServiceConfigs = value;
        }

        /// <summary>
        /// The resource name of the API Config.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("openapiDocuments")]
        private InputList<Inputs.ApiConfigOpenapiDocumentGetArgs>? _openapiDocuments;

        /// <summary>
        /// OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
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

        [Input("pulumiLabels")]
        private InputMap<string>? _pulumiLabels;

        /// <summary>
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        public InputMap<string> PulumiLabels
        {
            get => _pulumiLabels ?? (_pulumiLabels = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _pulumiLabels = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
        /// </summary>
        [Input("serviceConfigId")]
        public Input<string>? ServiceConfigId { get; set; }

        public ApiConfigState()
        {
        }
        public static new ApiConfigState Empty => new ApiConfigState();
    }
}
