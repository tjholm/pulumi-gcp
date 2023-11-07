// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Beyondcorp
{
    /// <summary>
    /// A BeyondCorp AppConnection resource represents a BeyondCorp protected AppConnection to a remote application.
    /// It creates all the necessary GCP components needed for creating a BeyondCorp protected AppConnection.
    /// Multiple connectors can be authorised for a single AppConnection.
    /// 
    /// To get more information about AppConnection, see:
    /// 
    /// * [API documentation](https://cloud.google.com/beyondcorp/docs/reference/rest#rest-resource:-v1.projects.locations.appconnections)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/beyondcorp-enterprise/docs/enable-app-connector)
    /// 
    /// ## Example Usage
    /// ### Beyondcorp App Connection Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var serviceAccount = new Gcp.ServiceAccount.Account("serviceAccount", new()
    ///     {
    ///         AccountId = "my-account",
    ///         DisplayName = "Test Service Account",
    ///     });
    /// 
    ///     var appConnector = new Gcp.Beyondcorp.AppConnector("appConnector", new()
    ///     {
    ///         PrincipalInfo = new Gcp.Beyondcorp.Inputs.AppConnectorPrincipalInfoArgs
    ///         {
    ///             ServiceAccount = new Gcp.Beyondcorp.Inputs.AppConnectorPrincipalInfoServiceAccountArgs
    ///             {
    ///                 Email = serviceAccount.Email,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var appConnection = new Gcp.Beyondcorp.AppConnection("appConnection", new()
    ///     {
    ///         Type = "TCP_PROXY",
    ///         ApplicationEndpoint = new Gcp.Beyondcorp.Inputs.AppConnectionApplicationEndpointArgs
    ///         {
    ///             Host = "foo-host",
    ///             Port = 8080,
    ///         },
    ///         Connectors = new[]
    ///         {
    ///             appConnector.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Beyondcorp App Connection Full
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var serviceAccount = new Gcp.ServiceAccount.Account("serviceAccount", new()
    ///     {
    ///         AccountId = "my-account",
    ///         DisplayName = "Test Service Account",
    ///     });
    /// 
    ///     var appGateway = new Gcp.Beyondcorp.AppGateway("appGateway", new()
    ///     {
    ///         Type = "TCP_PROXY",
    ///         HostType = "GCP_REGIONAL_MIG",
    ///     });
    /// 
    ///     var appConnector = new Gcp.Beyondcorp.AppConnector("appConnector", new()
    ///     {
    ///         PrincipalInfo = new Gcp.Beyondcorp.Inputs.AppConnectorPrincipalInfoArgs
    ///         {
    ///             ServiceAccount = new Gcp.Beyondcorp.Inputs.AppConnectorPrincipalInfoServiceAccountArgs
    ///             {
    ///                 Email = serviceAccount.Email,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var appConnection = new Gcp.Beyondcorp.AppConnection("appConnection", new()
    ///     {
    ///         Type = "TCP_PROXY",
    ///         DisplayName = "some display name",
    ///         ApplicationEndpoint = new Gcp.Beyondcorp.Inputs.AppConnectionApplicationEndpointArgs
    ///         {
    ///             Host = "foo-host",
    ///             Port = 8080,
    ///         },
    ///         Connectors = new[]
    ///         {
    ///             appConnector.Id,
    ///         },
    ///         Gateway = new Gcp.Beyondcorp.Inputs.AppConnectionGatewayArgs
    ///         {
    ///             AppGateway = appGateway.Id,
    ///         },
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///             { "bar", "baz" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// AppConnection can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:beyondcorp/appConnection:AppConnection default projects/{{project}}/locations/{{region}}/appConnections/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:beyondcorp/appConnection:AppConnection default {{project}}/{{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:beyondcorp/appConnection:AppConnection default {{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:beyondcorp/appConnection:AppConnection default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:beyondcorp/appConnection:AppConnection")]
    public partial class AppConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Address of the remote application endpoint for the BeyondCorp AppConnection.
        /// Structure is documented below.
        /// </summary>
        [Output("applicationEndpoint")]
        public Output<Outputs.AppConnectionApplicationEndpoint> ApplicationEndpoint { get; private set; } = null!;

        /// <summary>
        /// List of AppConnectors that are authorised to be associated with this AppConnection
        /// </summary>
        [Output("connectors")]
        public Output<ImmutableArray<string>> Connectors { get; private set; } = null!;

        /// <summary>
        /// An arbitrary user-provided name for the AppConnection.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
        /// clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// Gateway used by the AppConnection.
        /// Structure is documented below.
        /// </summary>
        [Output("gateway")]
        public Output<Outputs.AppConnectionGateway> Gateway { get; private set; } = null!;

        /// <summary>
        /// Resource labels to represent user provided metadata.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// ID of the AppConnection.
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
        /// The region of the AppConnection.
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        /// <summary>
        /// The type of network connectivity used by the AppConnection. Refer to
        /// https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#type
        /// for a list of possible values.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AppConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppConnection(string name, AppConnectionArgs args, CustomResourceOptions? options = null)
            : base("gcp:beyondcorp/appConnection:AppConnection", name, args ?? new AppConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppConnection(string name, Input<string> id, AppConnectionState? state = null, CustomResourceOptions? options = null)
            : base("gcp:beyondcorp/appConnection:AppConnection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppConnection Get(string name, Input<string> id, AppConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new AppConnection(name, id, state, options);
        }
    }

    public sealed class AppConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address of the remote application endpoint for the BeyondCorp AppConnection.
        /// Structure is documented below.
        /// </summary>
        [Input("applicationEndpoint", required: true)]
        public Input<Inputs.AppConnectionApplicationEndpointArgs> ApplicationEndpoint { get; set; } = null!;

        [Input("connectors")]
        private InputList<string>? _connectors;

        /// <summary>
        /// List of AppConnectors that are authorised to be associated with this AppConnection
        /// </summary>
        public InputList<string> Connectors
        {
            get => _connectors ?? (_connectors = new InputList<string>());
            set => _connectors = value;
        }

        /// <summary>
        /// An arbitrary user-provided name for the AppConnection.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Gateway used by the AppConnection.
        /// Structure is documented below.
        /// </summary>
        [Input("gateway")]
        public Input<Inputs.AppConnectionGatewayArgs>? Gateway { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels to represent user provided metadata.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// ID of the AppConnection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region of the AppConnection.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The type of network connectivity used by the AppConnection. Refer to
        /// https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#type
        /// for a list of possible values.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AppConnectionArgs()
        {
        }
        public static new AppConnectionArgs Empty => new AppConnectionArgs();
    }

    public sealed class AppConnectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address of the remote application endpoint for the BeyondCorp AppConnection.
        /// Structure is documented below.
        /// </summary>
        [Input("applicationEndpoint")]
        public Input<Inputs.AppConnectionApplicationEndpointGetArgs>? ApplicationEndpoint { get; set; }

        [Input("connectors")]
        private InputList<string>? _connectors;

        /// <summary>
        /// List of AppConnectors that are authorised to be associated with this AppConnection
        /// </summary>
        public InputList<string> Connectors
        {
            get => _connectors ?? (_connectors = new InputList<string>());
            set => _connectors = value;
        }

        /// <summary>
        /// An arbitrary user-provided name for the AppConnection.
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
        /// Gateway used by the AppConnection.
        /// Structure is documented below.
        /// </summary>
        [Input("gateway")]
        public Input<Inputs.AppConnectionGatewayGetArgs>? Gateway { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels to represent user provided metadata.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// ID of the AppConnection.
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

        /// <summary>
        /// The region of the AppConnection.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The type of network connectivity used by the AppConnection. Refer to
        /// https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#type
        /// for a list of possible values.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AppConnectionState()
        {
        }
        public static new AppConnectionState Empty => new AppConnectionState();
    }
}
