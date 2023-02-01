// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// A config defined for a single managed instance that belongs to an instance group manager. It preserves the instance name
    /// across instance group manager operations and can define stateful disks or metadata that are unique to the instance.
    /// This resource works with regional instance group managers.
    /// 
    /// To get more information about RegionPerInstanceConfig, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/compute/docs/instance-groups/stateful-migs#per-instance_configs)
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// RegionPerInstanceConfig can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/regionPerInstanceConfig:RegionPerInstanceConfig default projects/{{project}}/regions/{{region}}/instanceGroupManagers/{{region_instance_group_manager}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/regionPerInstanceConfig:RegionPerInstanceConfig default {{project}}/{{region}}/{{region_instance_group_manager}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/regionPerInstanceConfig:RegionPerInstanceConfig default {{region}}/{{region_instance_group_manager}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/regionPerInstanceConfig:RegionPerInstanceConfig default {{region_instance_group_manager}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:compute/regionPerInstanceConfig:RegionPerInstanceConfig")]
    public partial class RegionPerInstanceConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The minimal action to perform on the instance during an update.
        /// Default is `NONE`. Possible values are:
        /// * REPLACE
        /// * RESTART
        /// * REFRESH
        /// * NONE
        /// </summary>
        [Output("minimalAction")]
        public Output<string?> MinimalAction { get; private set; } = null!;

        /// <summary>
        /// The most disruptive action to perform on the instance during an update.
        /// Default is `REPLACE`. Possible values are:
        /// * REPLACE
        /// * RESTART
        /// * REFRESH
        /// * NONE
        /// </summary>
        [Output("mostDisruptiveAllowedAction")]
        public Output<string?> MostDisruptiveAllowedAction { get; private set; } = null!;

        /// <summary>
        /// The name for this per-instance config and its corresponding instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The preserved state for this instance.
        /// Structure is documented below.
        /// </summary>
        [Output("preservedState")]
        public Output<Outputs.RegionPerInstanceConfigPreservedState?> PreservedState { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Region where the containing instance group manager is located
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The region instance group manager this instance config is part of.
        /// </summary>
        [Output("regionInstanceGroupManager")]
        public Output<string> RegionInstanceGroupManager { get; private set; } = null!;

        /// <summary>
        /// When true, deleting this config will immediately remove any specified state from the underlying instance.
        /// When false, deleting this config will *not* immediately remove any state from the underlying instance.
        /// State will be removed on the next instance recreation or update.
        /// </summary>
        [Output("removeInstanceStateOnDestroy")]
        public Output<bool?> RemoveInstanceStateOnDestroy { get; private set; } = null!;


        /// <summary>
        /// Create a RegionPerInstanceConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegionPerInstanceConfig(string name, RegionPerInstanceConfigArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/regionPerInstanceConfig:RegionPerInstanceConfig", name, args ?? new RegionPerInstanceConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegionPerInstanceConfig(string name, Input<string> id, RegionPerInstanceConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/regionPerInstanceConfig:RegionPerInstanceConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RegionPerInstanceConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegionPerInstanceConfig Get(string name, Input<string> id, RegionPerInstanceConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new RegionPerInstanceConfig(name, id, state, options);
        }
    }

    public sealed class RegionPerInstanceConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The minimal action to perform on the instance during an update.
        /// Default is `NONE`. Possible values are:
        /// * REPLACE
        /// * RESTART
        /// * REFRESH
        /// * NONE
        /// </summary>
        [Input("minimalAction")]
        public Input<string>? MinimalAction { get; set; }

        /// <summary>
        /// The most disruptive action to perform on the instance during an update.
        /// Default is `REPLACE`. Possible values are:
        /// * REPLACE
        /// * RESTART
        /// * REFRESH
        /// * NONE
        /// </summary>
        [Input("mostDisruptiveAllowedAction")]
        public Input<string>? MostDisruptiveAllowedAction { get; set; }

        /// <summary>
        /// The name for this per-instance config and its corresponding instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The preserved state for this instance.
        /// Structure is documented below.
        /// </summary>
        [Input("preservedState")]
        public Input<Inputs.RegionPerInstanceConfigPreservedStateArgs>? PreservedState { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Region where the containing instance group manager is located
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The region instance group manager this instance config is part of.
        /// </summary>
        [Input("regionInstanceGroupManager", required: true)]
        public Input<string> RegionInstanceGroupManager { get; set; } = null!;

        /// <summary>
        /// When true, deleting this config will immediately remove any specified state from the underlying instance.
        /// When false, deleting this config will *not* immediately remove any state from the underlying instance.
        /// State will be removed on the next instance recreation or update.
        /// </summary>
        [Input("removeInstanceStateOnDestroy")]
        public Input<bool>? RemoveInstanceStateOnDestroy { get; set; }

        public RegionPerInstanceConfigArgs()
        {
        }
        public static new RegionPerInstanceConfigArgs Empty => new RegionPerInstanceConfigArgs();
    }

    public sealed class RegionPerInstanceConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The minimal action to perform on the instance during an update.
        /// Default is `NONE`. Possible values are:
        /// * REPLACE
        /// * RESTART
        /// * REFRESH
        /// * NONE
        /// </summary>
        [Input("minimalAction")]
        public Input<string>? MinimalAction { get; set; }

        /// <summary>
        /// The most disruptive action to perform on the instance during an update.
        /// Default is `REPLACE`. Possible values are:
        /// * REPLACE
        /// * RESTART
        /// * REFRESH
        /// * NONE
        /// </summary>
        [Input("mostDisruptiveAllowedAction")]
        public Input<string>? MostDisruptiveAllowedAction { get; set; }

        /// <summary>
        /// The name for this per-instance config and its corresponding instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The preserved state for this instance.
        /// Structure is documented below.
        /// </summary>
        [Input("preservedState")]
        public Input<Inputs.RegionPerInstanceConfigPreservedStateGetArgs>? PreservedState { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Region where the containing instance group manager is located
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The region instance group manager this instance config is part of.
        /// </summary>
        [Input("regionInstanceGroupManager")]
        public Input<string>? RegionInstanceGroupManager { get; set; }

        /// <summary>
        /// When true, deleting this config will immediately remove any specified state from the underlying instance.
        /// When false, deleting this config will *not* immediately remove any state from the underlying instance.
        /// State will be removed on the next instance recreation or update.
        /// </summary>
        [Input("removeInstanceStateOnDestroy")]
        public Input<bool>? RemoveInstanceStateOnDestroy { get; set; }

        public RegionPerInstanceConfigState()
        {
        }
        public static new RegionPerInstanceConfigState Empty => new RegionPerInstanceConfigState();
    }
}
