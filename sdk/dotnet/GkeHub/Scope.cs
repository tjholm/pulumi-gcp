// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeHub
{
    /// <summary>
    /// Scope represents a Scope in a Fleet.
    /// 
    /// To get more information about Scope, see:
    /// 
    /// * [API documentation](https://cloud.google.com/anthos/fleet-management/docs/reference/rest/v1/projects.locations.scopes)
    /// * How-to Guides
    ///     * [Registering a Cluster](https://cloud.google.com/anthos/multicluster-management/connect/registering-a-cluster#register_cluster)
    /// 
    /// ## Example Usage
    /// ### Gkehub Scope Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var scope = new Gcp.GkeHub.Scope("scope", new()
    ///     {
    ///         Labels = 
    ///         {
    ///             { "keya", "valuea" },
    ///             { "keyb", "valueb" },
    ///             { "keyc", "valuec" },
    ///         },
    ///         ScopeId = "tf-test-scope%{random_suffix}",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Scope can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:gkehub/scope:Scope default projects/{{project}}/locations/global/scopes/{{scope_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:gkehub/scope:Scope default {{project}}/{{scope_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:gkehub/scope:Scope default {{scope_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:gkehub/scope:Scope")]
    public partial class Scope : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Time the Scope was created in UTC.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Time the Scope was deleted in UTC.
        /// </summary>
        [Output("deleteTime")]
        public Output<string> DeleteTime { get; private set; } = null!;

        /// <summary>
        /// Labels for this Scope.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the scope
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
        /// The client-provided identifier of the scope.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("scopeId")]
        public Output<string> ScopeId { get; private set; } = null!;

        /// <summary>
        /// State of the scope resource.
        /// Structure is documented below.
        /// </summary>
        [Output("states")]
        public Output<ImmutableArray<Outputs.ScopeState>> States { get; private set; } = null!;

        /// <summary>
        /// Google-generated UUID for this resource.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// Time the Scope was updated in UTC.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Scope resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Scope(string name, ScopeArgs args, CustomResourceOptions? options = null)
            : base("gcp:gkehub/scope:Scope", name, args ?? new ScopeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Scope(string name, Input<string> id, ScopeState? state = null, CustomResourceOptions? options = null)
            : base("gcp:gkehub/scope:Scope", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Scope resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Scope Get(string name, Input<string> id, ScopeState? state = null, CustomResourceOptions? options = null)
        {
            return new Scope(name, id, state, options);
        }
    }

    public sealed class ScopeArgs : global::Pulumi.ResourceArgs
    {
        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels for this Scope.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The client-provided identifier of the scope.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("scopeId", required: true)]
        public Input<string> ScopeId { get; set; } = null!;

        public ScopeArgs()
        {
        }
        public static new ScopeArgs Empty => new ScopeArgs();
    }

    public sealed class ScopeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time the Scope was created in UTC.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Time the Scope was deleted in UTC.
        /// </summary>
        [Input("deleteTime")]
        public Input<string>? DeleteTime { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels for this Scope.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The unique identifier of the scope
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
        /// The client-provided identifier of the scope.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("scopeId")]
        public Input<string>? ScopeId { get; set; }

        [Input("states")]
        private InputList<Inputs.ScopeStateGetArgs>? _states;

        /// <summary>
        /// State of the scope resource.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ScopeStateGetArgs> States
        {
            get => _states ?? (_states = new InputList<Inputs.ScopeStateGetArgs>());
            set => _states = value;
        }

        /// <summary>
        /// Google-generated UUID for this resource.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// Time the Scope was updated in UTC.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public ScopeState()
        {
        }
        public static new ScopeState Empty => new ScopeState();
    }
}
