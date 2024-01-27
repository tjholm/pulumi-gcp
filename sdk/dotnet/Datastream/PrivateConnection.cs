// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Datastream
{
    /// <summary>
    /// The PrivateConnection resource is used to establish private connectivity between Datastream and a customer's network.
    /// 
    /// To get more information about PrivateConnection, see:
    /// 
    /// * [API documentation](https://cloud.google.com/datastream/docs/reference/rest/v1/projects.locations.privateConnections)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/datastream/docs/create-a-private-connectivity-configuration)
    /// 
    /// ## Example Usage
    /// ### Datastream Private Connection Full
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var defaultNetwork = new Gcp.Compute.Network("defaultNetwork");
    /// 
    ///     var defaultPrivateConnection = new Gcp.Datastream.PrivateConnection("defaultPrivateConnection", new()
    ///     {
    ///         DisplayName = "Connection profile",
    ///         Location = "us-central1",
    ///         PrivateConnectionId = "my-connection",
    ///         Labels = 
    ///         {
    ///             { "key", "value" },
    ///         },
    ///         VpcPeeringConfig = new Gcp.Datastream.Inputs.PrivateConnectionVpcPeeringConfigArgs
    ///         {
    ///             Vpc = defaultNetwork.Id,
    ///             Subnet = "10.0.0.0/29",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// PrivateConnection can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/privateConnections/{{private_connection_id}}` * `{{project}}/{{location}}/{{private_connection_id}}` * `{{location}}/{{private_connection_id}}` When using the `pulumi import` command, PrivateConnection can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:datastream/privateConnection:PrivateConnection default projects/{{project}}/locations/{{location}}/privateConnections/{{private_connection_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:datastream/privateConnection:PrivateConnection default {{project}}/{{location}}/{{private_connection_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:datastream/privateConnection:PrivateConnection default {{location}}/{{private_connection_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:datastream/privateConnection:PrivateConnection")]
    public partial class PrivateConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Display name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// The PrivateConnection error in case of failure.
        /// Structure is documented below.
        /// </summary>
        [Output("errors")]
        public Output<ImmutableArray<Outputs.PrivateConnectionError>> Errors { get; private set; } = null!;

        /// <summary>
        /// Labels.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The name of the location this private connection is located in.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource's name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The private connectivity identifier.
        /// </summary>
        [Output("privateConnectionId")]
        public Output<string> PrivateConnectionId { get; private set; } = null!;

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
        /// State of the PrivateConnection.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The VPC Peering configuration is used to create VPC peering
        /// between Datastream and the consumer's VPC.
        /// Structure is documented below.
        /// </summary>
        [Output("vpcPeeringConfig")]
        public Output<Outputs.PrivateConnectionVpcPeeringConfig> VpcPeeringConfig { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateConnection(string name, PrivateConnectionArgs args, CustomResourceOptions? options = null)
            : base("gcp:datastream/privateConnection:PrivateConnection", name, args ?? new PrivateConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateConnection(string name, Input<string> id, PrivateConnectionState? state = null, CustomResourceOptions? options = null)
            : base("gcp:datastream/privateConnection:PrivateConnection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PrivateConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateConnection Get(string name, Input<string> id, PrivateConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new PrivateConnection(name, id, state, options);
        }
    }

    public sealed class PrivateConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Display name.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The name of the location this private connection is located in.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The private connectivity identifier.
        /// </summary>
        [Input("privateConnectionId", required: true)]
        public Input<string> PrivateConnectionId { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The VPC Peering configuration is used to create VPC peering
        /// between Datastream and the consumer's VPC.
        /// Structure is documented below.
        /// </summary>
        [Input("vpcPeeringConfig", required: true)]
        public Input<Inputs.PrivateConnectionVpcPeeringConfigArgs> VpcPeeringConfig { get; set; } = null!;

        public PrivateConnectionArgs()
        {
        }
        public static new PrivateConnectionArgs Empty => new PrivateConnectionArgs();
    }

    public sealed class PrivateConnectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

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

        [Input("errors")]
        private InputList<Inputs.PrivateConnectionErrorGetArgs>? _errors;

        /// <summary>
        /// The PrivateConnection error in case of failure.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.PrivateConnectionErrorGetArgs> Errors
        {
            get => _errors ?? (_errors = new InputList<Inputs.PrivateConnectionErrorGetArgs>());
            set => _errors = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The name of the location this private connection is located in.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource's name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The private connectivity identifier.
        /// </summary>
        [Input("privateConnectionId")]
        public Input<string>? PrivateConnectionId { get; set; }

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
        /// State of the PrivateConnection.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The VPC Peering configuration is used to create VPC peering
        /// between Datastream and the consumer's VPC.
        /// Structure is documented below.
        /// </summary>
        [Input("vpcPeeringConfig")]
        public Input<Inputs.PrivateConnectionVpcPeeringConfigGetArgs>? VpcPeeringConfig { get; set; }

        public PrivateConnectionState()
        {
        }
        public static new PrivateConnectionState Empty => new PrivateConnectionState();
    }
}
