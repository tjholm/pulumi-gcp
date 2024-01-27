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
    /// ### Network Services Grpc Route Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.NetworkServices.GrpcRoute("default", new()
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
    ///             new Gcp.NetworkServices.Inputs.GrpcRouteRuleArgs
    ///             {
    ///                 Matches = new[]
    ///                 {
    ///                     new Gcp.NetworkServices.Inputs.GrpcRouteRuleMatchArgs
    ///                     {
    ///                         Headers = new[]
    ///                         {
    ///                             new Gcp.NetworkServices.Inputs.GrpcRouteRuleMatchHeaderArgs
    ///                             {
    ///                                 Key = "key",
    ///                                 Value = "value",
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///                 Action = new Gcp.NetworkServices.Inputs.GrpcRouteRuleActionArgs
    ///                 {
    ///                     RetryPolicy = new Gcp.NetworkServices.Inputs.GrpcRouteRuleActionRetryPolicyArgs
    ///                     {
    ///                         RetryConditions = new[]
    ///                         {
    ///                             "cancelled",
    ///                         },
    ///                         NumRetries = 1,
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
    /// ### Network Services Grpc Route Matches And Actions
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.NetworkServices.GrpcRoute("default", new()
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
    ///             new Gcp.NetworkServices.Inputs.GrpcRouteRuleArgs
    ///             {
    ///                 Matches = new[]
    ///                 {
    ///                     new Gcp.NetworkServices.Inputs.GrpcRouteRuleMatchArgs
    ///                     {
    ///                         Headers = new[]
    ///                         {
    ///                             new Gcp.NetworkServices.Inputs.GrpcRouteRuleMatchHeaderArgs
    ///                             {
    ///                                 Key = "key",
    ///                                 Value = "value",
    ///                             },
    ///                         },
    ///                     },
    ///                     new Gcp.NetworkServices.Inputs.GrpcRouteRuleMatchArgs
    ///                     {
    ///                         Headers = new[]
    ///                         {
    ///                             new Gcp.NetworkServices.Inputs.GrpcRouteRuleMatchHeaderArgs
    ///                             {
    ///                                 Key = "key",
    ///                                 Value = "value",
    ///                             },
    ///                         },
    ///                         Method = new Gcp.NetworkServices.Inputs.GrpcRouteRuleMatchMethodArgs
    ///                         {
    ///                             GrpcService = "foo",
    ///                             GrpcMethod = "bar",
    ///                             CaseSensitive = true,
    ///                         },
    ///                     },
    ///                 },
    ///                 Action = new Gcp.NetworkServices.Inputs.GrpcRouteRuleActionArgs
    ///                 {
    ///                     FaultInjectionPolicy = new Gcp.NetworkServices.Inputs.GrpcRouteRuleActionFaultInjectionPolicyArgs
    ///                     {
    ///                         Delay = new Gcp.NetworkServices.Inputs.GrpcRouteRuleActionFaultInjectionPolicyDelayArgs
    ///                         {
    ///                             FixedDelay = "1s",
    ///                             Percentage = 1,
    ///                         },
    ///                         Abort = new Gcp.NetworkServices.Inputs.GrpcRouteRuleActionFaultInjectionPolicyAbortArgs
    ///                         {
    ///                             HttpStatus = 500,
    ///                             Percentage = 1,
    ///                         },
    ///                     },
    ///                     RetryPolicy = new Gcp.NetworkServices.Inputs.GrpcRouteRuleActionRetryPolicyArgs
    ///                     {
    ///                         RetryConditions = new[]
    ///                         {
    ///                             "cancelled",
    ///                         },
    ///                         NumRetries = 1,
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
    /// ### Network Services Grpc Route Actions
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.NetworkServices.GrpcRoute("default", new()
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
    ///             new Gcp.NetworkServices.Inputs.GrpcRouteRuleArgs
    ///             {
    ///                 Action = new Gcp.NetworkServices.Inputs.GrpcRouteRuleActionArgs
    ///                 {
    ///                     FaultInjectionPolicy = new Gcp.NetworkServices.Inputs.GrpcRouteRuleActionFaultInjectionPolicyArgs
    ///                     {
    ///                         Delay = new Gcp.NetworkServices.Inputs.GrpcRouteRuleActionFaultInjectionPolicyDelayArgs
    ///                         {
    ///                             FixedDelay = "1s",
    ///                             Percentage = 1,
    ///                         },
    ///                         Abort = new Gcp.NetworkServices.Inputs.GrpcRouteRuleActionFaultInjectionPolicyAbortArgs
    ///                         {
    ///                             HttpStatus = 500,
    ///                             Percentage = 1,
    ///                         },
    ///                     },
    ///                     RetryPolicy = new Gcp.NetworkServices.Inputs.GrpcRouteRuleActionRetryPolicyArgs
    ///                     {
    ///                         RetryConditions = new[]
    ///                         {
    ///                             "cancelled",
    ///                         },
    ///                         NumRetries = 1,
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
    /// GrpcRoute can be imported using any of these accepted formats* `projects/{{project}}/locations/global/grpcRoutes/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, GrpcRoute can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:networkservices/grpcRoute:GrpcRoute default projects/{{project}}/locations/global/grpcRoutes/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:networkservices/grpcRoute:GrpcRoute default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:networkservices/grpcRoute:GrpcRoute default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:networkservices/grpcRoute:GrpcRoute")]
    public partial class GrpcRoute : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Time the GrpcRoute was created in UTC.
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
        /// List of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway.
        /// </summary>
        [Output("gateways")]
        public Output<ImmutableArray<string>> Gateways { get; private set; } = null!;

        /// <summary>
        /// Required. Service hostnames with an optional port for which this route describes traffic.
        /// </summary>
        [Output("hostnames")]
        public Output<ImmutableArray<string>> Hostnames { get; private set; } = null!;

        /// <summary>
        /// Set of label tags associated with the GrpcRoute resource.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// List of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh.
        /// </summary>
        [Output("meshes")]
        public Output<ImmutableArray<string>> Meshes { get; private set; } = null!;

        /// <summary>
        /// Name of the GrpcRoute resource.
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
        public Output<ImmutableArray<Outputs.GrpcRouteRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL of this resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Time the GrpcRoute was updated in UTC.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a GrpcRoute resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GrpcRoute(string name, GrpcRouteArgs args, CustomResourceOptions? options = null)
            : base("gcp:networkservices/grpcRoute:GrpcRoute", name, args ?? new GrpcRouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GrpcRoute(string name, Input<string> id, GrpcRouteState? state = null, CustomResourceOptions? options = null)
            : base("gcp:networkservices/grpcRoute:GrpcRoute", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GrpcRoute resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GrpcRoute Get(string name, Input<string> id, GrpcRouteState? state = null, CustomResourceOptions? options = null)
        {
            return new GrpcRoute(name, id, state, options);
        }
    }

    public sealed class GrpcRouteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A free-text description of the resource. Max length 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("gateways")]
        private InputList<string>? _gateways;

        /// <summary>
        /// List of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway.
        /// </summary>
        public InputList<string> Gateways
        {
            get => _gateways ?? (_gateways = new InputList<string>());
            set => _gateways = value;
        }

        [Input("hostnames", required: true)]
        private InputList<string>? _hostnames;

        /// <summary>
        /// Required. Service hostnames with an optional port for which this route describes traffic.
        /// </summary>
        public InputList<string> Hostnames
        {
            get => _hostnames ?? (_hostnames = new InputList<string>());
            set => _hostnames = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Set of label tags associated with the GrpcRoute resource.
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
        /// List of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh.
        /// </summary>
        public InputList<string> Meshes
        {
            get => _meshes ?? (_meshes = new InputList<string>());
            set => _meshes = value;
        }

        /// <summary>
        /// Name of the GrpcRoute resource.
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
        private InputList<Inputs.GrpcRouteRuleArgs>? _rules;

        /// <summary>
        /// Rules that define how traffic is routed and handled.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.GrpcRouteRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.GrpcRouteRuleArgs>());
            set => _rules = value;
        }

        public GrpcRouteArgs()
        {
        }
        public static new GrpcRouteArgs Empty => new GrpcRouteArgs();
    }

    public sealed class GrpcRouteState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time the GrpcRoute was created in UTC.
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
        /// List of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway.
        /// </summary>
        public InputList<string> Gateways
        {
            get => _gateways ?? (_gateways = new InputList<string>());
            set => _gateways = value;
        }

        [Input("hostnames")]
        private InputList<string>? _hostnames;

        /// <summary>
        /// Required. Service hostnames with an optional port for which this route describes traffic.
        /// </summary>
        public InputList<string> Hostnames
        {
            get => _hostnames ?? (_hostnames = new InputList<string>());
            set => _hostnames = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Set of label tags associated with the GrpcRoute resource.
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
        /// List of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh.
        /// </summary>
        public InputList<string> Meshes
        {
            get => _meshes ?? (_meshes = new InputList<string>());
            set => _meshes = value;
        }

        /// <summary>
        /// Name of the GrpcRoute resource.
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
        private InputList<Inputs.GrpcRouteRuleGetArgs>? _rules;

        /// <summary>
        /// Rules that define how traffic is routed and handled.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.GrpcRouteRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.GrpcRouteRuleGetArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Server-defined URL of this resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// Time the GrpcRoute was updated in UTC.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public GrpcRouteState()
        {
        }
        public static new GrpcRouteState Empty => new GrpcRouteState();
    }
}
