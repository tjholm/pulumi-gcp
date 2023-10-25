// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Vertex
{
    /// <summary>
    /// An endpoint indexes are deployed into. An index endpoint can have multiple deployed indexes.
    /// 
    /// To get more information about IndexEndpoint, see:
    /// 
    /// * [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.indexEndpoints/)
    /// 
    /// ## Example Usage
    /// ### Vertex Ai Index Endpoint
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var vertexNetwork = Gcp.Compute.GetNetwork.Invoke(new()
    ///     {
    ///         Name = "network-name",
    ///     });
    /// 
    ///     var vertexRange = new Gcp.Compute.GlobalAddress("vertexRange", new()
    ///     {
    ///         Purpose = "VPC_PEERING",
    ///         AddressType = "INTERNAL",
    ///         PrefixLength = 24,
    ///         Network = vertexNetwork.Apply(getNetworkResult =&gt; getNetworkResult.Id),
    ///     });
    /// 
    ///     var vertexVpcConnection = new Gcp.ServiceNetworking.Connection("vertexVpcConnection", new()
    ///     {
    ///         Network = vertexNetwork.Apply(getNetworkResult =&gt; getNetworkResult.Id),
    ///         Service = "servicenetworking.googleapis.com",
    ///         ReservedPeeringRanges = new[]
    ///         {
    ///             vertexRange.Name,
    ///         },
    ///     });
    /// 
    ///     var project = Gcp.Organizations.GetProject.Invoke();
    /// 
    ///     var indexEndpoint = new Gcp.Vertex.AiIndexEndpoint("indexEndpoint", new()
    ///     {
    ///         DisplayName = "sample-endpoint",
    ///         Description = "A sample vertex endpoint",
    ///         Region = "us-central1",
    ///         Labels = 
    ///         {
    ///             { "label-one", "value-one" },
    ///         },
    ///         Network = Output.Tuple(project, vertexNetwork).Apply(values =&gt;
    ///         {
    ///             var project = values.Item1;
    ///             var vertexNetwork = values.Item2;
    ///             return $"projects/{project.Apply(getProjectResult =&gt; getProjectResult.Number)}/global/networks/{vertexNetwork.Apply(getNetworkResult =&gt; getNetworkResult.Name)}";
    ///         }),
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             vertexVpcConnection,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Vertex Ai Index Endpoint With Public Endpoint
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var indexEndpoint = new Gcp.Vertex.AiIndexEndpoint("indexEndpoint", new()
    ///     {
    ///         Description = "A sample vertex endpoint with an public endpoint",
    ///         DisplayName = "sample-endpoint",
    ///         Labels = 
    ///         {
    ///             { "label-one", "value-one" },
    ///         },
    ///         PublicEndpointEnabled = true,
    ///         Region = "us-central1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// IndexEndpoint can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiIndexEndpoint:AiIndexEndpoint default projects/{{project}}/locations/{{region}}/indexEndpoints/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiIndexEndpoint:AiIndexEndpoint default {{project}}/{{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiIndexEndpoint:AiIndexEndpoint default {{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiIndexEndpoint:AiIndexEndpoint default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:vertex/aiIndexEndpoint:AiIndexEndpoint")]
    public partial class AiIndexEndpoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The timestamp of when the Index was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The description of the Index.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the Index. The name can be up to 128 characters long and can consist of any UTF-8 characters.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Used to perform consistent read-modify-write updates.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The labels with user-defined metadata to organize your Indexes.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The resource name of the Index.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the index endpoint should be peered.
        /// Private services access must already be configured for the network. If left unspecified, the index endpoint is not peered with any network.
        /// [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`.
        /// Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
        /// </summary>
        [Output("network")]
        public Output<string?> Network { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// If publicEndpointEnabled is true, this field will be populated with the domain name to use for this index endpoint.
        /// </summary>
        [Output("publicEndpointDomainName")]
        public Output<string> PublicEndpointDomainName { get; private set; } = null!;

        /// <summary>
        /// If true, the deployed index will be accessible through public endpoint.
        /// </summary>
        [Output("publicEndpointEnabled")]
        public Output<bool?> PublicEndpointEnabled { get; private set; } = null!;

        /// <summary>
        /// The region of the index endpoint. eg us-central1
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        /// <summary>
        /// The timestamp of when the Index was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a AiIndexEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AiIndexEndpoint(string name, AiIndexEndpointArgs args, CustomResourceOptions? options = null)
            : base("gcp:vertex/aiIndexEndpoint:AiIndexEndpoint", name, args ?? new AiIndexEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AiIndexEndpoint(string name, Input<string> id, AiIndexEndpointState? state = null, CustomResourceOptions? options = null)
            : base("gcp:vertex/aiIndexEndpoint:AiIndexEndpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AiIndexEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AiIndexEndpoint Get(string name, Input<string> id, AiIndexEndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new AiIndexEndpoint(name, id, state, options);
        }
    }

    public sealed class AiIndexEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Index.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the Index. The name can be up to 128 characters long and can consist of any UTF-8 characters.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels with user-defined metadata to organize your Indexes.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the index endpoint should be peered.
        /// Private services access must already be configured for the network. If left unspecified, the index endpoint is not peered with any network.
        /// [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`.
        /// Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// If true, the deployed index will be accessible through public endpoint.
        /// </summary>
        [Input("publicEndpointEnabled")]
        public Input<bool>? PublicEndpointEnabled { get; set; }

        /// <summary>
        /// The region of the index endpoint. eg us-central1
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public AiIndexEndpointArgs()
        {
        }
        public static new AiIndexEndpointArgs Empty => new AiIndexEndpointArgs();
    }

    public sealed class AiIndexEndpointState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The timestamp of when the Index was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The description of the Index.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the Index. The name can be up to 128 characters long and can consist of any UTF-8 characters.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Used to perform consistent read-modify-write updates.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels with user-defined metadata to organize your Indexes.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The resource name of the Index.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the index endpoint should be peered.
        /// Private services access must already be configured for the network. If left unspecified, the index endpoint is not peered with any network.
        /// [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`.
        /// Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// If publicEndpointEnabled is true, this field will be populated with the domain name to use for this index endpoint.
        /// </summary>
        [Input("publicEndpointDomainName")]
        public Input<string>? PublicEndpointDomainName { get; set; }

        /// <summary>
        /// If true, the deployed index will be accessible through public endpoint.
        /// </summary>
        [Input("publicEndpointEnabled")]
        public Input<bool>? PublicEndpointEnabled { get; set; }

        /// <summary>
        /// The region of the index endpoint. eg us-central1
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The timestamp of when the Index was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public AiIndexEndpointState()
        {
        }
        public static new AiIndexEndpointState Empty => new AiIndexEndpointState();
    }
}
