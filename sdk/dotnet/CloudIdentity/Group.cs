// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudIdentity
{
    /// <summary>
    /// A Cloud Identity resource representing a Group.
    /// 
    /// To get more information about Group, see:
    /// 
    /// * [API documentation](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/identity/docs/how-to/setup)
    /// 
    /// &gt; **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
    /// you must specify a `billing_project` and set `user_project_override` to true
    /// in the provider configuration. Otherwise the Cloud Identity API will return a 403 error.
    /// Your account must have the `serviceusage.services.use` permission on the
    /// `billing_project` you defined.
    /// 
    /// ## Example Usage
    /// ### Cloud Identity Groups Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var cloudIdentityGroupBasic = new Gcp.CloudIdentity.Group("cloudIdentityGroupBasic", new()
    ///     {
    ///         DisplayName = "my-identity-group",
    ///         GroupKey = new Gcp.CloudIdentity.Inputs.GroupGroupKeyArgs
    ///         {
    ///             Id = "my-identity-group@example.com",
    ///         },
    ///         InitialGroupConfig = "WITH_INITIAL_OWNER",
    ///         Labels = 
    ///         {
    ///             { "cloudidentity.googleapis.com/groups.discussion_forum", "" },
    ///         },
    ///         Parent = "customers/A01b123xz",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Group can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:cloudidentity/group:Group default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:cloudidentity/group:Group")]
    public partial class Group : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The time when the Group was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// An extended description to help users determine the purpose of a Group.
        /// Must not be longer than 4,096 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the Group.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// EntityKey of the Group.
        /// Structure is documented below.
        /// </summary>
        [Output("groupKey")]
        public Output<Outputs.GroupGroupKey> GroupKey { get; private set; } = null!;

        /// <summary>
        /// The initial configuration options for creating a Group.
        /// See the
        /// [API reference](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups/create#initialgroupconfig)
        /// for possible values.
        /// Default value is `EMPTY`.
        /// Possible values are: `INITIAL_GROUP_CONFIG_UNSPECIFIED`, `WITH_INITIAL_OWNER`, `EMPTY`.
        /// </summary>
        [Output("initialGroupConfig")]
        public Output<string?> InitialGroupConfig { get; private set; } = null!;

        /// <summary>
        /// One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value.
        /// Google Groups are the default type of group and have a label with a key of cloudidentity.googleapis.com/groups.discussion_forum and an empty value.
        /// Existing Google Groups can have an additional label with a key of cloudidentity.googleapis.com/groups.security and an empty value added to them. This is an immutable change and the security label cannot be removed once added.
        /// Dynamic groups have a label with a key of cloudidentity.googleapis.com/groups.dynamic.
        /// Identity-mapped groups for Cloud Search have a label with a key of system/groups/external and an empty value.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Resource name of the Group in the format: groups/{group_id}, where group_id
        /// is the unique ID assigned to the Group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The resource name of the entity under which this Group resides in the
        /// Cloud Identity resource hierarchy.
        /// Must be of the form identitysources/{identity_source_id} for external-identity-mapped
        /// groups or customers/{customer_id} for Google Groups.
        /// </summary>
        [Output("parent")]
        public Output<string> Parent { get; private set; } = null!;

        /// <summary>
        /// The time when the Group was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs args, CustomResourceOptions? options = null)
            : base("gcp:cloudidentity/group:Group", name, args ?? new GroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
            : base("gcp:cloudidentity/group:Group", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Group resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Group Get(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Group(name, id, state, options);
        }
    }

    public sealed class GroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An extended description to help users determine the purpose of a Group.
        /// Must not be longer than 4,096 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the Group.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// EntityKey of the Group.
        /// Structure is documented below.
        /// </summary>
        [Input("groupKey", required: true)]
        public Input<Inputs.GroupGroupKeyArgs> GroupKey { get; set; } = null!;

        /// <summary>
        /// The initial configuration options for creating a Group.
        /// See the
        /// [API reference](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups/create#initialgroupconfig)
        /// for possible values.
        /// Default value is `EMPTY`.
        /// Possible values are: `INITIAL_GROUP_CONFIG_UNSPECIFIED`, `WITH_INITIAL_OWNER`, `EMPTY`.
        /// </summary>
        [Input("initialGroupConfig")]
        public Input<string>? InitialGroupConfig { get; set; }

        [Input("labels", required: true)]
        private InputMap<string>? _labels;

        /// <summary>
        /// One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value.
        /// Google Groups are the default type of group and have a label with a key of cloudidentity.googleapis.com/groups.discussion_forum and an empty value.
        /// Existing Google Groups can have an additional label with a key of cloudidentity.googleapis.com/groups.security and an empty value added to them. This is an immutable change and the security label cannot be removed once added.
        /// Dynamic groups have a label with a key of cloudidentity.googleapis.com/groups.dynamic.
        /// Identity-mapped groups for Cloud Search have a label with a key of system/groups/external and an empty value.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The resource name of the entity under which this Group resides in the
        /// Cloud Identity resource hierarchy.
        /// Must be of the form identitysources/{identity_source_id} for external-identity-mapped
        /// groups or customers/{customer_id} for Google Groups.
        /// </summary>
        [Input("parent", required: true)]
        public Input<string> Parent { get; set; } = null!;

        public GroupArgs()
        {
        }
        public static new GroupArgs Empty => new GroupArgs();
    }

    public sealed class GroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time when the Group was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// An extended description to help users determine the purpose of a Group.
        /// Must not be longer than 4,096 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the Group.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// EntityKey of the Group.
        /// Structure is documented below.
        /// </summary>
        [Input("groupKey")]
        public Input<Inputs.GroupGroupKeyGetArgs>? GroupKey { get; set; }

        /// <summary>
        /// The initial configuration options for creating a Group.
        /// See the
        /// [API reference](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups/create#initialgroupconfig)
        /// for possible values.
        /// Default value is `EMPTY`.
        /// Possible values are: `INITIAL_GROUP_CONFIG_UNSPECIFIED`, `WITH_INITIAL_OWNER`, `EMPTY`.
        /// </summary>
        [Input("initialGroupConfig")]
        public Input<string>? InitialGroupConfig { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value.
        /// Google Groups are the default type of group and have a label with a key of cloudidentity.googleapis.com/groups.discussion_forum and an empty value.
        /// Existing Google Groups can have an additional label with a key of cloudidentity.googleapis.com/groups.security and an empty value added to them. This is an immutable change and the security label cannot be removed once added.
        /// Dynamic groups have a label with a key of cloudidentity.googleapis.com/groups.dynamic.
        /// Identity-mapped groups for Cloud Search have a label with a key of system/groups/external and an empty value.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Resource name of the Group in the format: groups/{group_id}, where group_id
        /// is the unique ID assigned to the Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The resource name of the entity under which this Group resides in the
        /// Cloud Identity resource hierarchy.
        /// Must be of the form identitysources/{identity_source_id} for external-identity-mapped
        /// groups or customers/{customer_id} for Google Groups.
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        /// <summary>
        /// The time when the Group was last updated.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public GroupState()
        {
        }
        public static new GroupState Empty => new GroupState();
    }
}
