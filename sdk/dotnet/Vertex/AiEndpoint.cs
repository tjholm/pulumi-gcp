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
    /// Models are deployed into it, and afterwards Endpoint is called to obtain predictions and explanations.
    /// 
    /// To get more information about Endpoint, see:
    /// 
    /// * [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1beta1/projects.locations.endpoints)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/vertex-ai/docs)
    /// 
    /// ## Example Usage
    /// ### Vertex Ai Endpoint Network
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var vertexNetwork = new Gcp.Compute.Network("vertexNetwork");
    /// 
    ///     var vertexRange = new Gcp.Compute.GlobalAddress("vertexRange", new()
    ///     {
    ///         Purpose = "VPC_PEERING",
    ///         AddressType = "INTERNAL",
    ///         PrefixLength = 24,
    ///         Network = vertexNetwork.Id,
    ///     });
    /// 
    ///     var vertexVpcConnection = new Gcp.ServiceNetworking.Connection("vertexVpcConnection", new()
    ///     {
    ///         Network = vertexNetwork.Id,
    ///         Service = "servicenetworking.googleapis.com",
    ///         ReservedPeeringRanges = new[]
    ///         {
    ///             vertexRange.Name,
    ///         },
    ///     });
    /// 
    ///     var project = Gcp.Organizations.GetProject.Invoke();
    /// 
    ///     var endpoint = new Gcp.Vertex.AiEndpoint("endpoint", new()
    ///     {
    ///         DisplayName = "sample-endpoint",
    ///         Description = "A sample vertex endpoint",
    ///         Location = "us-central1",
    ///         Region = "us-central1",
    ///         Labels = 
    ///         {
    ///             { "label-one", "value-one" },
    ///         },
    ///         Network = Output.Tuple(project, vertexNetwork.Name).Apply(values =&gt;
    ///         {
    ///             var project = values.Item1;
    ///             var name = values.Item2;
    ///             return $"projects/{project.Apply(getProjectResult =&gt; getProjectResult.Number)}/global/networks/{name}";
    ///         }),
    ///         EncryptionSpec = new Gcp.Vertex.Inputs.AiEndpointEncryptionSpecArgs
    ///         {
    ///             KmsKeyName = "kms-name",
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             vertexVpcConnection,
    ///         },
    ///     });
    /// 
    ///     var cryptoKey = new Gcp.Kms.CryptoKeyIAMMember("cryptoKey", new()
    ///     {
    ///         CryptoKeyId = "kms-name",
    ///         Role = "roles/cloudkms.cryptoKeyEncrypterDecrypter",
    ///         Member = $"serviceAccount:service-{project.Apply(getProjectResult =&gt; getProjectResult.Number)}@gcp-sa-aiplatform.iam.gserviceaccount.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Endpoint can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiEndpoint:AiEndpoint default projects/{{project}}/locations/{{location}}/endpoints/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiEndpoint:AiEndpoint default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiEndpoint:AiEndpoint default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:vertex/aiEndpoint:AiEndpoint")]
    public partial class AiEndpoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// (Output)
        /// Output only. Timestamp when the DeployedModel was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and EndpointService.UndeployModel respectively. Models can also be deployed and undeployed using the [Cloud Console](https://console.cloud.google.com/vertex-ai/).
        /// Structure is documented below.
        /// </summary>
        [Output("deployedModels")]
        public Output<ImmutableArray<Outputs.AiEndpointDeployedModel>> DeployedModels { get; private set; } = null!;

        /// <summary>
        /// The description of the Endpoint.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.
        /// Structure is documented below.
        /// </summary>
        [Output("encryptionSpec")]
        public Output<Outputs.AiEndpointEncryptionSpec?> EncryptionSpec { get; private set; } = null!;

        /// <summary>
        /// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The location for the resource
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by CreateModelDeploymentMonitoringJob. Format: `projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}`
        /// </summary>
        [Output("modelDeploymentMonitoringJob")]
        public Output<string> ModelDeploymentMonitoringJob { get; private set; } = null!;

        /// <summary>
        /// The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`. Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
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
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        [Output("pulumiLabels")]
        public Output<ImmutableDictionary<string, string>> PulumiLabels { get; private set; } = null!;

        /// <summary>
        /// The region for the resource
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        /// <summary>
        /// Output only. Timestamp when this Endpoint was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a AiEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AiEndpoint(string name, AiEndpointArgs args, CustomResourceOptions? options = null)
            : base("gcp:vertex/aiEndpoint:AiEndpoint", name, args ?? new AiEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AiEndpoint(string name, Input<string> id, AiEndpointState? state = null, CustomResourceOptions? options = null)
            : base("gcp:vertex/aiEndpoint:AiEndpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AiEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AiEndpoint Get(string name, Input<string> id, AiEndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new AiEndpoint(name, id, state, options);
        }
    }

    public sealed class AiEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Endpoint.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.
        /// Structure is documented below.
        /// </summary>
        [Input("encryptionSpec")]
        public Input<Inputs.AiEndpointEncryptionSpecArgs>? EncryptionSpec { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The location for the resource
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`. Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
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
        /// The region for the resource
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public AiEndpointArgs()
        {
        }
        public static new AiEndpointArgs Empty => new AiEndpointArgs();
    }

    public sealed class AiEndpointState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Output)
        /// Output only. Timestamp when the DeployedModel was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        [Input("deployedModels")]
        private InputList<Inputs.AiEndpointDeployedModelGetArgs>? _deployedModels;

        /// <summary>
        /// Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and EndpointService.UndeployModel respectively. Models can also be deployed and undeployed using the [Cloud Console](https://console.cloud.google.com/vertex-ai/).
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.AiEndpointDeployedModelGetArgs> DeployedModels
        {
            get => _deployedModels ?? (_deployedModels = new InputList<Inputs.AiEndpointDeployedModelGetArgs>());
            set => _deployedModels = value;
        }

        /// <summary>
        /// The description of the Endpoint.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters.
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

        /// <summary>
        /// Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.
        /// Structure is documented below.
        /// </summary>
        [Input("encryptionSpec")]
        public Input<Inputs.AiEndpointEncryptionSpecGetArgs>? EncryptionSpec { get; set; }

        /// <summary>
        /// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The location for the resource
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by CreateModelDeploymentMonitoringJob. Format: `projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}`
        /// </summary>
        [Input("modelDeploymentMonitoringJob")]
        public Input<string>? ModelDeploymentMonitoringJob { get; set; }

        /// <summary>
        /// The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`. Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

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
        /// The region for the resource
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Output only. Timestamp when this Endpoint was last updated.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public AiEndpointState()
        {
        }
        public static new AiEndpointState Empty => new AiEndpointState();
    }
}
