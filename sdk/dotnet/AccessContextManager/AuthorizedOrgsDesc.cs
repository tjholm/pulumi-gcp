// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AccessContextManager
{
    /// <summary>
    /// An authorized organizations description describes a list of organizations
    /// (1) that have been authorized to use certain asset (for example, device) data
    /// owned by different organizations at the enforcement points, or (2) with certain
    /// asset (for example, device) have been authorized to access the resources in
    /// another organization at the enforcement points.
    /// 
    /// To get more information about AuthorizedOrgsDesc, see:
    /// 
    /// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.authorizedOrgsDescs)
    /// * How-to Guides
    ///     * [gcloud docs](https://cloud.google.com/beyondcorp-enterprise/docs/cross-org-authorization)
    /// 
    /// &gt; **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
    /// you must specify a `billing_project` and set `user_project_override` to true
    /// in the provider configuration. Otherwise the ACM API will return a 403 error.
    /// Your account must have the `serviceusage.services.use` permission on the
    /// `billing_project` you defined.
    /// 
    /// ## Example Usage
    /// ### Access Context Manager Authorized Orgs Desc Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test_access = new Gcp.AccessContextManager.AccessPolicy("test-access", new()
    ///     {
    ///         Parent = "organizations/",
    ///         Title = "my policy",
    ///     });
    /// 
    ///     var authorized_orgs_desc = new Gcp.AccessContextManager.AuthorizedOrgsDesc("authorized-orgs-desc", new()
    ///     {
    ///         AssetType = "ASSET_TYPE_CREDENTIAL_STRENGTH",
    ///         AuthorizationDirection = "AUTHORIZATION_DIRECTION_TO",
    ///         AuthorizationType = "AUTHORIZATION_TYPE_TRUST",
    ///         Orgs = new[]
    ///         {
    ///             "organizations/12345",
    ///             "organizations/98765",
    ///         },
    ///         Parent = test_access.Name.Apply(name =&gt; $"accessPolicies/{name}"),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// AuthorizedOrgsDesc can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:accesscontextmanager/authorizedOrgsDesc:AuthorizedOrgsDesc default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:accesscontextmanager/authorizedOrgsDesc:AuthorizedOrgsDesc")]
    public partial class AuthorizedOrgsDesc : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The type of entities that need to use the authorization relationship during
        /// evaluation, such as a device. Valid values are "ASSET_TYPE_DEVICE" and
        /// "ASSET_TYPE_CREDENTIAL_STRENGTH".
        /// Possible values are: `ASSET_TYPE_DEVICE`, `ASSET_TYPE_CREDENTIAL_STRENGTH`.
        /// </summary>
        [Output("assetType")]
        public Output<string?> AssetType { get; private set; } = null!;

        /// <summary>
        /// The direction of the authorization relationship between this organization
        /// and the organizations listed in the "orgs" field. The valid values for this
        /// field include the following:
        /// AUTHORIZATION_DIRECTION_FROM: Allows this organization to evaluate traffic
        /// in the organizations listed in the `orgs` field.
        /// AUTHORIZATION_DIRECTION_TO: Allows the organizations listed in the `orgs`
        /// field to evaluate the traffic in this organization.
        /// For the authorization relationship to take effect, all of the organizations
        /// must authorize and specify the appropriate relationship direction. For
        /// example, if organization A authorized organization B and C to evaluate its
        /// traffic, by specifying "AUTHORIZATION_DIRECTION_TO" as the authorization
        /// direction, organizations B and C must specify
        /// "AUTHORIZATION_DIRECTION_FROM" as the authorization direction in their
        /// "AuthorizedOrgsDesc" resource.
        /// Possible values are: `AUTHORIZATION_DIRECTION_TO`, `AUTHORIZATION_DIRECTION_FROM`.
        /// </summary>
        [Output("authorizationDirection")]
        public Output<string?> AuthorizationDirection { get; private set; } = null!;

        /// <summary>
        /// A granular control type for authorization levels. Valid value is "AUTHORIZATION_TYPE_TRUST".
        /// Possible values are: `AUTHORIZATION_TYPE_TRUST`.
        /// </summary>
        [Output("authorizationType")]
        public Output<string?> AuthorizationType { get; private set; } = null!;

        /// <summary>
        /// Time the AuthorizedOrgsDesc was created in UTC.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Resource name for the `AuthorizedOrgsDesc`. Format:
        /// `accessPolicies/{access_policy}/authorizedOrgsDescs/{authorized_orgs_desc}`.
        /// The `authorized_orgs_desc` component must begin with a letter, followed by
        /// alphanumeric characters or `_`.
        /// After you create an `AuthorizedOrgsDesc`, you cannot change its `name`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The list of organization ids in this AuthorizedOrgsDesc.
        /// Format: `organizations/&lt;org_number&gt;`
        /// Example: `organizations/123456`
        /// </summary>
        [Output("orgs")]
        public Output<ImmutableArray<string>> Orgs { get; private set; } = null!;

        /// <summary>
        /// Required. Resource name for the access policy which owns this `AuthorizedOrgsDesc`.
        /// </summary>
        [Output("parent")]
        public Output<string> Parent { get; private set; } = null!;

        /// <summary>
        /// Time the AuthorizedOrgsDesc was updated in UTC.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a AuthorizedOrgsDesc resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthorizedOrgsDesc(string name, AuthorizedOrgsDescArgs args, CustomResourceOptions? options = null)
            : base("gcp:accesscontextmanager/authorizedOrgsDesc:AuthorizedOrgsDesc", name, args ?? new AuthorizedOrgsDescArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthorizedOrgsDesc(string name, Input<string> id, AuthorizedOrgsDescState? state = null, CustomResourceOptions? options = null)
            : base("gcp:accesscontextmanager/authorizedOrgsDesc:AuthorizedOrgsDesc", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthorizedOrgsDesc resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthorizedOrgsDesc Get(string name, Input<string> id, AuthorizedOrgsDescState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthorizedOrgsDesc(name, id, state, options);
        }
    }

    public sealed class AuthorizedOrgsDescArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of entities that need to use the authorization relationship during
        /// evaluation, such as a device. Valid values are "ASSET_TYPE_DEVICE" and
        /// "ASSET_TYPE_CREDENTIAL_STRENGTH".
        /// Possible values are: `ASSET_TYPE_DEVICE`, `ASSET_TYPE_CREDENTIAL_STRENGTH`.
        /// </summary>
        [Input("assetType")]
        public Input<string>? AssetType { get; set; }

        /// <summary>
        /// The direction of the authorization relationship between this organization
        /// and the organizations listed in the "orgs" field. The valid values for this
        /// field include the following:
        /// AUTHORIZATION_DIRECTION_FROM: Allows this organization to evaluate traffic
        /// in the organizations listed in the `orgs` field.
        /// AUTHORIZATION_DIRECTION_TO: Allows the organizations listed in the `orgs`
        /// field to evaluate the traffic in this organization.
        /// For the authorization relationship to take effect, all of the organizations
        /// must authorize and specify the appropriate relationship direction. For
        /// example, if organization A authorized organization B and C to evaluate its
        /// traffic, by specifying "AUTHORIZATION_DIRECTION_TO" as the authorization
        /// direction, organizations B and C must specify
        /// "AUTHORIZATION_DIRECTION_FROM" as the authorization direction in their
        /// "AuthorizedOrgsDesc" resource.
        /// Possible values are: `AUTHORIZATION_DIRECTION_TO`, `AUTHORIZATION_DIRECTION_FROM`.
        /// </summary>
        [Input("authorizationDirection")]
        public Input<string>? AuthorizationDirection { get; set; }

        /// <summary>
        /// A granular control type for authorization levels. Valid value is "AUTHORIZATION_TYPE_TRUST".
        /// Possible values are: `AUTHORIZATION_TYPE_TRUST`.
        /// </summary>
        [Input("authorizationType")]
        public Input<string>? AuthorizationType { get; set; }

        /// <summary>
        /// Resource name for the `AuthorizedOrgsDesc`. Format:
        /// `accessPolicies/{access_policy}/authorizedOrgsDescs/{authorized_orgs_desc}`.
        /// The `authorized_orgs_desc` component must begin with a letter, followed by
        /// alphanumeric characters or `_`.
        /// After you create an `AuthorizedOrgsDesc`, you cannot change its `name`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("orgs")]
        private InputList<string>? _orgs;

        /// <summary>
        /// The list of organization ids in this AuthorizedOrgsDesc.
        /// Format: `organizations/&lt;org_number&gt;`
        /// Example: `organizations/123456`
        /// </summary>
        public InputList<string> Orgs
        {
            get => _orgs ?? (_orgs = new InputList<string>());
            set => _orgs = value;
        }

        /// <summary>
        /// Required. Resource name for the access policy which owns this `AuthorizedOrgsDesc`.
        /// </summary>
        [Input("parent", required: true)]
        public Input<string> Parent { get; set; } = null!;

        public AuthorizedOrgsDescArgs()
        {
        }
        public static new AuthorizedOrgsDescArgs Empty => new AuthorizedOrgsDescArgs();
    }

    public sealed class AuthorizedOrgsDescState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of entities that need to use the authorization relationship during
        /// evaluation, such as a device. Valid values are "ASSET_TYPE_DEVICE" and
        /// "ASSET_TYPE_CREDENTIAL_STRENGTH".
        /// Possible values are: `ASSET_TYPE_DEVICE`, `ASSET_TYPE_CREDENTIAL_STRENGTH`.
        /// </summary>
        [Input("assetType")]
        public Input<string>? AssetType { get; set; }

        /// <summary>
        /// The direction of the authorization relationship between this organization
        /// and the organizations listed in the "orgs" field. The valid values for this
        /// field include the following:
        /// AUTHORIZATION_DIRECTION_FROM: Allows this organization to evaluate traffic
        /// in the organizations listed in the `orgs` field.
        /// AUTHORIZATION_DIRECTION_TO: Allows the organizations listed in the `orgs`
        /// field to evaluate the traffic in this organization.
        /// For the authorization relationship to take effect, all of the organizations
        /// must authorize and specify the appropriate relationship direction. For
        /// example, if organization A authorized organization B and C to evaluate its
        /// traffic, by specifying "AUTHORIZATION_DIRECTION_TO" as the authorization
        /// direction, organizations B and C must specify
        /// "AUTHORIZATION_DIRECTION_FROM" as the authorization direction in their
        /// "AuthorizedOrgsDesc" resource.
        /// Possible values are: `AUTHORIZATION_DIRECTION_TO`, `AUTHORIZATION_DIRECTION_FROM`.
        /// </summary>
        [Input("authorizationDirection")]
        public Input<string>? AuthorizationDirection { get; set; }

        /// <summary>
        /// A granular control type for authorization levels. Valid value is "AUTHORIZATION_TYPE_TRUST".
        /// Possible values are: `AUTHORIZATION_TYPE_TRUST`.
        /// </summary>
        [Input("authorizationType")]
        public Input<string>? AuthorizationType { get; set; }

        /// <summary>
        /// Time the AuthorizedOrgsDesc was created in UTC.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Resource name for the `AuthorizedOrgsDesc`. Format:
        /// `accessPolicies/{access_policy}/authorizedOrgsDescs/{authorized_orgs_desc}`.
        /// The `authorized_orgs_desc` component must begin with a letter, followed by
        /// alphanumeric characters or `_`.
        /// After you create an `AuthorizedOrgsDesc`, you cannot change its `name`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("orgs")]
        private InputList<string>? _orgs;

        /// <summary>
        /// The list of organization ids in this AuthorizedOrgsDesc.
        /// Format: `organizations/&lt;org_number&gt;`
        /// Example: `organizations/123456`
        /// </summary>
        public InputList<string> Orgs
        {
            get => _orgs ?? (_orgs = new InputList<string>());
            set => _orgs = value;
        }

        /// <summary>
        /// Required. Resource name for the access policy which owns this `AuthorizedOrgsDesc`.
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        /// <summary>
        /// Time the AuthorizedOrgsDesc was updated in UTC.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public AuthorizedOrgsDescState()
        {
        }
        public static new AuthorizedOrgsDescState Empty => new AuthorizedOrgsDescState();
    }
}
