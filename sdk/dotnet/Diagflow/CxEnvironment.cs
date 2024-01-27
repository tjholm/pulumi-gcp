// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Diagflow
{
    /// <summary>
    /// Represents an environment for an agent. You can create multiple versions of your agent and publish them to separate environments.
    /// When you edit an agent, you are editing the draft agent. At any point, you can save the draft agent as an agent version, which is an immutable snapshot of your agent.
    /// When you save the draft agent, it is published to the default environment. When you create agent versions, you can publish them to custom environments. You can create a variety of custom environments for testing, development, production, etc.
    /// 
    /// To get more information about Environment, see:
    /// 
    /// * [API documentation](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.environments)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/dialogflow/cx/docs)
    /// 
    /// ## Example Usage
    /// ### Dialogflowcx Environment Full
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var agent = new Gcp.Diagflow.CxAgent("agent", new()
    ///     {
    ///         DisplayName = "dialogflowcx-agent",
    ///         Location = "global",
    ///         DefaultLanguageCode = "en",
    ///         SupportedLanguageCodes = new[]
    ///         {
    ///             "fr",
    ///             "de",
    ///             "es",
    ///         },
    ///         TimeZone = "America/New_York",
    ///         Description = "Example description.",
    ///         AvatarUri = "https://cloud.google.com/_static/images/cloud/icons/favicons/onecloud/super_cloud.png",
    ///         EnableStackdriverLogging = true,
    ///         EnableSpellCorrection = true,
    ///         SpeechToTextSettings = new Gcp.Diagflow.Inputs.CxAgentSpeechToTextSettingsArgs
    ///         {
    ///             EnableSpeechAdaptation = true,
    ///         },
    ///     });
    /// 
    ///     var version1 = new Gcp.Diagflow.CxVersion("version1", new()
    ///     {
    ///         Parent = agent.StartFlow,
    ///         DisplayName = "1.0.0",
    ///         Description = "version 1.0.0",
    ///     });
    /// 
    ///     var development = new Gcp.Diagflow.CxEnvironment("development", new()
    ///     {
    ///         Parent = agent.Id,
    ///         DisplayName = "Development",
    ///         Description = "Development Environment",
    ///         VersionConfigs = new[]
    ///         {
    ///             new Gcp.Diagflow.Inputs.CxEnvironmentVersionConfigArgs
    ///             {
    ///                 Version = version1.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Environment can be imported using any of these accepted formats* `{{parent}}/environments/{{name}}` * `{{parent}}/{{name}}` When using the `pulumi import` command, Environment can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:diagflow/cxEnvironment:CxEnvironment default {{parent}}/environments/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:diagflow/cxEnvironment:CxEnvironment default {{parent}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:diagflow/cxEnvironment:CxEnvironment")]
    public partial class CxEnvironment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The name of the environment.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Agent to create an Environment for.
        /// Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;.
        /// </summary>
        [Output("parent")]
        public Output<string?> Parent { get; private set; } = null!;

        /// <summary>
        /// Update time of this environment. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// A list of configurations for flow versions. You should include version configs for all flows that are reachable from [Start Flow][Agent.start_flow] in the agent. Otherwise, an error will be returned.
        /// Structure is documented below.
        /// </summary>
        [Output("versionConfigs")]
        public Output<ImmutableArray<Outputs.CxEnvironmentVersionConfig>> VersionConfigs { get; private set; } = null!;


        /// <summary>
        /// Create a CxEnvironment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CxEnvironment(string name, CxEnvironmentArgs args, CustomResourceOptions? options = null)
            : base("gcp:diagflow/cxEnvironment:CxEnvironment", name, args ?? new CxEnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CxEnvironment(string name, Input<string> id, CxEnvironmentState? state = null, CustomResourceOptions? options = null)
            : base("gcp:diagflow/cxEnvironment:CxEnvironment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CxEnvironment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CxEnvironment Get(string name, Input<string> id, CxEnvironmentState? state = null, CustomResourceOptions? options = null)
        {
            return new CxEnvironment(name, id, state, options);
        }
    }

    public sealed class CxEnvironmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The Agent to create an Environment for.
        /// Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;.
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        [Input("versionConfigs", required: true)]
        private InputList<Inputs.CxEnvironmentVersionConfigArgs>? _versionConfigs;

        /// <summary>
        /// A list of configurations for flow versions. You should include version configs for all flows that are reachable from [Start Flow][Agent.start_flow] in the agent. Otherwise, an error will be returned.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.CxEnvironmentVersionConfigArgs> VersionConfigs
        {
            get => _versionConfigs ?? (_versionConfigs = new InputList<Inputs.CxEnvironmentVersionConfigArgs>());
            set => _versionConfigs = value;
        }

        public CxEnvironmentArgs()
        {
        }
        public static new CxEnvironmentArgs Empty => new CxEnvironmentArgs();
    }

    public sealed class CxEnvironmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The name of the environment.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Agent to create an Environment for.
        /// Format: projects/&lt;Project ID&gt;/locations/&lt;Location ID&gt;/agents/&lt;Agent ID&gt;.
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        /// <summary>
        /// Update time of this environment. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        [Input("versionConfigs")]
        private InputList<Inputs.CxEnvironmentVersionConfigGetArgs>? _versionConfigs;

        /// <summary>
        /// A list of configurations for flow versions. You should include version configs for all flows that are reachable from [Start Flow][Agent.start_flow] in the agent. Otherwise, an error will be returned.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.CxEnvironmentVersionConfigGetArgs> VersionConfigs
        {
            get => _versionConfigs ?? (_versionConfigs = new InputList<Inputs.CxEnvironmentVersionConfigGetArgs>());
            set => _versionConfigs = value;
        }

        public CxEnvironmentState()
        {
        }
        public static new CxEnvironmentState Empty => new CxEnvironmentState();
    }
}
