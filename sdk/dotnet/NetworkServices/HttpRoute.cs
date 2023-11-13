// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.NetworkServices
{
    /// <summary>
    /// ## Example Usage
    /// ### Network Services Http Route Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.NetworkServices.HttpRoute("default", new()
    ///     {
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Description = "my description",
    ///         Hostnames = new[]
    ///         {
    ///             "example",
    ///         },
    ///         Rules = new[]
    ///         {
    ///             new Gcp.NetworkServices.Inputs.HttpRouteRuleArgs
    ///             {
    ///                 Matches = new[]
    ///                 {
    ///                     new Gcp.NetworkServices.Inputs.HttpRouteRuleMatchArgs
    ///                     {
    ///                         QueryParameters = new[]
    ///                         {
    ///                             new Gcp.NetworkServices.Inputs.HttpRouteRuleMatchQueryParameterArgs
    ///                             {
    ///                                 QueryParameter = "key",
    ///                                 ExactMatch = "value",
    ///                             },
    ///                         },
    ///                         FullPathMatch = "example",
    ///                     },
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
    /// ### Network Services Http Route Matches And Actions
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.NetworkServices.HttpRoute("default", new()
    ///     {
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Description = "my description",
    ///         Hostnames = new[]
    ///         {
    ///             "example",
    ///         },
    ///         Rules = new[]
    ///         {
    ///             new Gcp.NetworkServices.Inputs.HttpRouteRuleArgs
    ///             {
    ///                 Matches = new[]
    ///                 {
    ///                     new Gcp.NetworkServices.Inputs.HttpRouteRuleMatchArgs
    ///                     {
    ///                         Headers = new[]
    ///                         {
    ///                             new Gcp.NetworkServices.Inputs.HttpRouteRuleMatchHeaderArgs
    ///                             {
    ///                                 Header = "header",
    ///                                 InvertMatch = false,
    ///                                 RegexMatch = "header-value",
    ///                             },
    ///                         },
    ///                         QueryParameters = new[]
    ///                         {
    ///                             new Gcp.NetworkServices.Inputs.HttpRouteRuleMatchQueryParameterArgs
    ///                             {
    ///                                 QueryParameter = "key",
    ///                                 ExactMatch = "value",
    ///                             },
    ///                         },
    ///                         PrefixMatch = "example",
    ///                         IgnoreCase = false,
    ///                     },
    ///                     new Gcp.NetworkServices.Inputs.HttpRouteRuleMatchArgs
    ///                     {
    ///                         Headers = new[]
    ///                         {
    ///                             new Gcp.NetworkServices.Inputs.HttpRouteRuleMatchHeaderArgs
    ///                             {
    ///                                 Header = "header",
    ///                                 InvertMatch = false,
    ///                                 PresentMatch = true,
    ///                             },
    ///                         },
    ///                         QueryParameters = new[]
    ///                         {
    ///                             new Gcp.NetworkServices.Inputs.HttpRouteRuleMatchQueryParameterArgs
    ///                             {
    ///                                 QueryParameter = "key",
    ///                                 RegexMatch = "value",
    ///                             },
    ///                         },
    ///                         RegexMatch = "example",
    ///                         IgnoreCase = false,
    ///                     },
    ///                     new Gcp.NetworkServices.Inputs.HttpRouteRuleMatchArgs
    ///                     {
    ///                         Headers = new[]
    ///                         {
    ///                             new Gcp.NetworkServices.Inputs.HttpRouteRuleMatchHeaderArgs
    ///                             {
    ///                                 Header = "header",
    ///                                 InvertMatch = false,
    ///                                 PresentMatch = true,
    ///                             },
    ///                         },
    ///                         QueryParameters = new[]
    ///                         {
    ///                             new Gcp.NetworkServices.Inputs.HttpRouteRuleMatchQueryParameterArgs
    ///                             {
    ///                                 QueryParameter = "key",
    ///                                 PresentMatch = true,
    ///                             },
    ///                         },
    ///                         FullPathMatch = "example",
    ///                         IgnoreCase = false,
    ///                     },
    ///                 },
    ///                 Action = new Gcp.NetworkServices.Inputs.HttpRouteRuleActionArgs
    ///                 {
    ///                     Redirect = new Gcp.NetworkServices.Inputs.HttpRouteRuleActionRedirectArgs
    ///                     {
    ///                         HostRedirect = "new-host",
    ///                         PathRedirect = "new-path",
    ///                         PrefixRewrite = "new-prefix",
    ///                         HttpsRedirect = true,
    ///                         StripQuery = true,
    ///                         PortRedirect = 8081,
    ///                     },
    ///                     UrlRewrite = new Gcp.NetworkServices.Inputs.HttpRouteRuleActionUrlRewriteArgs
    ///                     {
    ///                         PathPrefixRewrite = "new-prefix",
    ///                         HostRewrite = "new-host",
    ///                     },
    ///                     RetryPolicy = new Gcp.NetworkServices.Inputs.HttpRouteRuleActionRetryPolicyArgs
    ///                     {
    ///                         RetryConditions = new[]
    ///                         {
    ///                             "server_error",
    ///                         },
    ///                         NumRetries = 1,
    ///                         PerTryTimeout = "1s",
    ///                     },
    ///                     RequestMirrorPolicy = new Gcp.NetworkServices.Inputs.HttpRouteRuleActionRequestMirrorPolicyArgs
    ///                     {
    ///                         Destination = new Gcp.NetworkServices.Inputs.HttpRouteRuleActionRequestMirrorPolicyDestinationArgs
    ///                         {
    ///                             ServiceName = "new",
    ///                             Weight = 1,
    ///                         },
    ///                     },
    ///                     CorsPolicy = new Gcp.NetworkServices.Inputs.HttpRouteRuleActionCorsPolicyArgs
    ///                     {
    ///                         AllowOrigins = new[]
    ///                         {
    ///                             "example",
    ///                         },
    ///                         AllowMethods = new[]
    ///                         {
    ///                             "GET",
    ///                             "PUT",
    ///                         },
    ///                         AllowHeaders = new[]
    ///                         {
    ///                             "version",
    ///                             "type",
    ///                         },
    ///                         ExposeHeaders = new[]
    ///                         {
    ///                             "version",
    ///                             "type",
    ///                         },
    ///                         MaxAge = "1s",
    ///                         AllowCredentials = true,
    ///                         Disabled = false,
    ///                     },
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
    /// ### Network Services Http Route Actions
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.NetworkServices.HttpRoute("default", new()
    ///     {
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Description = "my description",
    ///         Hostnames = new[]
    ///         {
    ///             "example",
    ///         },
    ///         Rules = new[]
    ///         {
    ///             new Gcp.NetworkServices.Inputs.HttpRouteRuleArgs
    ///             {
    ///                 Action = new Gcp.NetworkServices.Inputs.HttpRouteRuleActionArgs
    ///                 {
    ///                     FaultInjectionPolicy = new Gcp.NetworkServices.Inputs.HttpRouteRuleActionFaultInjectionPolicyArgs
    ///                     {
    ///                         Delay = new Gcp.NetworkServices.Inputs.HttpRouteRuleActionFaultInjectionPolicyDelayArgs
    ///                         {
    ///                             FixedDelay = "1s",
    ///                             Percentage = 1,
    ///                         },
    ///                         Abort = new Gcp.NetworkServices.Inputs.HttpRouteRuleActionFaultInjectionPolicyAbortArgs
    ///                         {
    ///                             HttpStatus = 500,
    ///                             Percentage = 1,
    ///                         },
    ///                     },
    ///                     UrlRewrite = new Gcp.NetworkServices.Inputs.HttpRouteRuleActionUrlRewriteArgs
    ///                     {
    ///                         PathPrefixRewrite = "new-prefix",
    ///                         HostRewrite = "new-host",
    ///                     },
    ///                     RetryPolicy = new Gcp.NetworkServices.Inputs.HttpRouteRuleActionRetryPolicyArgs
    ///                     {
    ///                         RetryConditions = new[]
    ///                         {
    ///                             "server_error",
    ///                         },
    ///                         NumRetries = 1,
    ///                         PerTryTimeout = "1s",
    ///                     },
    ///                     RequestMirrorPolicy = new Gcp.NetworkServices.Inputs.HttpRouteRuleActionRequestMirrorPolicyArgs
    ///                     {
    ///                         Destination = new Gcp.NetworkServices.Inputs.HttpRouteRuleActionRequestMirrorPolicyDestinationArgs
    ///                         {
    ///                             ServiceName = "new",
    ///                             Weight = 1,
    ///                         },
    ///                     },
    ///                     CorsPolicy = new Gcp.NetworkServices.Inputs.HttpRouteRuleActionCorsPolicyArgs
    ///                     {
    ///                         AllowOrigins = new[]
    ///                         {
    ///                             "example",
    ///                         },
    ///                         AllowMethods = new[]
    ///                         {
    ///                             "GET",
    ///                             "PUT",
    ///                         },
    ///                         AllowHeaders = new[]
    ///                         {
    ///                             "version",
    ///                             "type",
    ///                         },
    ///                         ExposeHeaders = new[]
    ///                         {
    ///                             "version",
    ///                             "type",
    ///                         },
    ///                         MaxAge = "1s",
    ///                         AllowCredentials = true,
    ///                         Disabled = false,
    ///                     },
    ///                     RequestHeaderModifier = new Gcp.NetworkServices.Inputs.HttpRouteRuleActionRequestHeaderModifierArgs
    ///                     {
    ///                         Set = 
    ///                         {
    ///                             { "version", "1" },
    ///                             { "type", "json" },
    ///                         },
    ///                         Add = 
    ///                         {
    ///                             { "minor-version", "1" },
    ///                         },
    ///                         Removes = new[]
    ///                         {
    ///                             "arg",
    ///                         },
    ///                     },
    ///                     ResponseHeaderModifier = new Gcp.NetworkServices.Inputs.HttpRouteRuleActionResponseHeaderModifierArgs
    ///                     {
    ///                         Set = 
    ///                         {
    ///                             { "version", "1" },
    ///                             { "type", "json" },
    ///                         },
    ///                         Add = 
    ///                         {
    ///                             { "minor-version", "1" },
    ///                         },
    ///                         Removes = new[]
    ///                         {
    ///                             "removearg",
    ///                         },
    ///                     },
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
    /// ### Network Services Http Route Mesh Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var defaultMesh = new Gcp.NetworkServices.Mesh("defaultMesh", new()
    ///     {
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Description = "my description",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var defaultHttpRoute = new Gcp.NetworkServices.HttpRoute("defaultHttpRoute", new()
    ///     {
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Description = "my description",
    ///         Hostnames = new[]
    ///         {
    ///             "example",
    ///         },
    ///         Meshes = new[]
    ///         {
    ///             defaultMesh.Id,
    ///         },
    ///         Rules = new[]
    ///         {
    ///             new Gcp.NetworkServices.Inputs.HttpRouteRuleArgs
    ///             {
    ///                 Matches = new[]
    ///                 {
    ///                     new Gcp.NetworkServices.Inputs.HttpRouteRuleMatchArgs
    ///                     {
    ///                         QueryParameters = new[]
    ///                         {
    ///                             new Gcp.NetworkServices.Inputs.HttpRouteRuleMatchQueryParameterArgs
    ///                             {
    ///                                 QueryParameter = "key",
    ///                                 ExactMatch = "value",
    ///                             },
    ///                         },
    ///                         FullPathMatch = "example",
    ///                     },
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
    /// HttpRoute can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:networkservices/httpRoute:HttpRoute default projects/{{project}}/locations/global/httpRoutes/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:networkservices/httpRoute:HttpRoute default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:networkservices/httpRoute:HttpRoute default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:networkservices/httpRoute:HttpRoute")]
    public partial class HttpRoute : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Time the HttpRoute was created in UTC.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// A free-text description of the resource. Max length 1024 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// Gateways defines a list of gateways this HttpRoute is attached to, as one of the routing rules to route the requests served by the gateway.
        /// Each gateway reference should match the pattern: projects/*/locations/global/gateways/&lt;gateway_name&gt;
        /// </summary>
        [Output("gateways")]
        public Output<ImmutableArray<string>> Gateways { get; private set; } = null!;

        /// <summary>
        /// Set of hosts that should match against the HTTP host header to select a HttpRoute to process the request.
        /// </summary>
        [Output("hostnames")]
        public Output<ImmutableArray<string>> Hostnames { get; private set; } = null!;

        /// <summary>
        /// Set of label tags associated with the HttpRoute resource.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Meshes defines a list of meshes this HttpRoute is attached to, as one of the routing rules to route the requests served by the mesh.
        /// Each mesh reference should match the pattern: projects/*/locations/global/meshes/&lt;mesh_name&gt;.
        /// The attached Mesh should be of a type SIDECAR.
        /// </summary>
        [Output("meshes")]
        public Output<ImmutableArray<string>> Meshes { get; private set; } = null!;

        /// <summary>
        /// Name of the HttpRoute resource.
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
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        [Output("pulumiLabels")]
        public Output<ImmutableDictionary<string, string>> PulumiLabels { get; private set; } = null!;

        /// <summary>
        /// Rules that define how traffic is routed and handled.
        /// Structure is documented below.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.HttpRouteRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL of this resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Time the HttpRoute was updated in UTC.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a HttpRoute resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HttpRoute(string name, HttpRouteArgs args, CustomResourceOptions? options = null)
            : base("gcp:networkservices/httpRoute:HttpRoute", name, args ?? new HttpRouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HttpRoute(string name, Input<string> id, HttpRouteState? state = null, CustomResourceOptions? options = null)
            : base("gcp:networkservices/httpRoute:HttpRoute", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HttpRoute resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HttpRoute Get(string name, Input<string> id, HttpRouteState? state = null, CustomResourceOptions? options = null)
        {
            return new HttpRoute(name, id, state, options);
        }
    }

    public sealed class HttpRouteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A free-text description of the resource. Max length 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("gateways")]
        private InputList<string>? _gateways;

        /// <summary>
        /// Gateways defines a list of gateways this HttpRoute is attached to, as one of the routing rules to route the requests served by the gateway.
        /// Each gateway reference should match the pattern: projects/*/locations/global/gateways/&lt;gateway_name&gt;
        /// </summary>
        public InputList<string> Gateways
        {
            get => _gateways ?? (_gateways = new InputList<string>());
            set => _gateways = value;
        }

        [Input("hostnames", required: true)]
        private InputList<string>? _hostnames;

        /// <summary>
        /// Set of hosts that should match against the HTTP host header to select a HttpRoute to process the request.
        /// </summary>
        public InputList<string> Hostnames
        {
            get => _hostnames ?? (_hostnames = new InputList<string>());
            set => _hostnames = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Set of label tags associated with the HttpRoute resource.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("meshes")]
        private InputList<string>? _meshes;

        /// <summary>
        /// Meshes defines a list of meshes this HttpRoute is attached to, as one of the routing rules to route the requests served by the mesh.
        /// Each mesh reference should match the pattern: projects/*/locations/global/meshes/&lt;mesh_name&gt;.
        /// The attached Mesh should be of a type SIDECAR.
        /// </summary>
        public InputList<string> Meshes
        {
            get => _meshes ?? (_meshes = new InputList<string>());
            set => _meshes = value;
        }

        /// <summary>
        /// Name of the HttpRoute resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("rules", required: true)]
        private InputList<Inputs.HttpRouteRuleArgs>? _rules;

        /// <summary>
        /// Rules that define how traffic is routed and handled.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.HttpRouteRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.HttpRouteRuleArgs>());
            set => _rules = value;
        }

        public HttpRouteArgs()
        {
        }
        public static new HttpRouteArgs Empty => new HttpRouteArgs();
    }

    public sealed class HttpRouteState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time the HttpRoute was created in UTC.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// A free-text description of the resource. Max length 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("effectiveLabels")]
        private InputMap<string>? _effectiveLabels;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
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

        [Input("gateways")]
        private InputList<string>? _gateways;

        /// <summary>
        /// Gateways defines a list of gateways this HttpRoute is attached to, as one of the routing rules to route the requests served by the gateway.
        /// Each gateway reference should match the pattern: projects/*/locations/global/gateways/&lt;gateway_name&gt;
        /// </summary>
        public InputList<string> Gateways
        {
            get => _gateways ?? (_gateways = new InputList<string>());
            set => _gateways = value;
        }

        [Input("hostnames")]
        private InputList<string>? _hostnames;

        /// <summary>
        /// Set of hosts that should match against the HTTP host header to select a HttpRoute to process the request.
        /// </summary>
        public InputList<string> Hostnames
        {
            get => _hostnames ?? (_hostnames = new InputList<string>());
            set => _hostnames = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Set of label tags associated with the HttpRoute resource.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("meshes")]
        private InputList<string>? _meshes;

        /// <summary>
        /// Meshes defines a list of meshes this HttpRoute is attached to, as one of the routing rules to route the requests served by the mesh.
        /// Each mesh reference should match the pattern: projects/*/locations/global/meshes/&lt;mesh_name&gt;.
        /// The attached Mesh should be of a type SIDECAR.
        /// </summary>
        public InputList<string> Meshes
        {
            get => _meshes ?? (_meshes = new InputList<string>());
            set => _meshes = value;
        }

        /// <summary>
        /// Name of the HttpRoute resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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

        [Input("rules")]
        private InputList<Inputs.HttpRouteRuleGetArgs>? _rules;

        /// <summary>
        /// Rules that define how traffic is routed and handled.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.HttpRouteRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.HttpRouteRuleGetArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Server-defined URL of this resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// Time the HttpRoute was updated in UTC.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public HttpRouteState()
        {
        }
        public static new HttpRouteState Empty => new HttpRouteState();
    }
}
