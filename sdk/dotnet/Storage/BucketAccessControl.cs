// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Storage
{
    /// <summary>
    /// Bucket ACLs can be managed authoritatively using the
    /// `storage_bucket_acl` resource. Do not use these two resources in conjunction to manage the same bucket.
    /// 
    /// The BucketAccessControls resource manages the Access Control List
    /// (ACLs) for a single entity/role pairing on a bucket. ACLs let you specify who
    /// has access to your data and to what extent.
    /// 
    /// There are three roles that can be assigned to an entity:
    /// 
    /// READERs can get the bucket, though no acl property will be returned, and
    /// list the bucket's objects.  WRITERs are READERs, and they can insert
    /// objects into the bucket and delete the bucket's objects.  OWNERs are
    /// WRITERs, and they can get the acl property of a bucket, update a bucket,
    /// and call all BucketAccessControls methods on the bucket.  For more
    /// information, see Access Control, with the caveat that this API uses
    /// READER, WRITER, and OWNER instead of READ, WRITE, and FULL_CONTROL.
    /// 
    /// To get more information about BucketAccessControl, see:
    /// 
    /// * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/storage/docs/access-control/lists)
    /// 
    /// ## Example Usage
    /// ### Storage Bucket Access Control Public Bucket
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var bucket = new Gcp.Storage.Bucket("bucket", new()
    ///     {
    ///         Location = "US",
    ///     });
    /// 
    ///     var publicRule = new Gcp.Storage.BucketAccessControl("publicRule", new()
    ///     {
    ///         Bucket = bucket.Name,
    ///         Role = "READER",
    ///         Entity = "allUsers",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// BucketAccessControl can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:storage/bucketAccessControl:BucketAccessControl default {{bucket}}/{{entity}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:storage/bucketAccessControl:BucketAccessControl")]
    public partial class BucketAccessControl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// The domain associated with the entity.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// The email address associated with the entity.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// The entity holding the permission, in one of the following forms:
        /// user-userId
        /// user-email
        /// group-groupId
        /// group-email
        /// domain-domain
        /// project-team-projectId
        /// allUsers
        /// allAuthenticatedUsers
        /// Examples:
        /// The user liz@example.com would be user-liz@example.com.
        /// The group example@googlegroups.com would be
        /// group-example@googlegroups.com.
        /// To refer to all members of the Google Apps for Business domain
        /// example.com, the entity would be domain-example.com.
        /// </summary>
        [Output("entity")]
        public Output<string> Entity { get; private set; } = null!;

        /// <summary>
        /// The access permission for the entity.
        /// Possible values are: `OWNER`, `READER`, `WRITER`.
        /// </summary>
        [Output("role")]
        public Output<string?> Role { get; private set; } = null!;


        /// <summary>
        /// Create a BucketAccessControl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BucketAccessControl(string name, BucketAccessControlArgs args, CustomResourceOptions? options = null)
            : base("gcp:storage/bucketAccessControl:BucketAccessControl", name, args ?? new BucketAccessControlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BucketAccessControl(string name, Input<string> id, BucketAccessControlState? state = null, CustomResourceOptions? options = null)
            : base("gcp:storage/bucketAccessControl:BucketAccessControl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BucketAccessControl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BucketAccessControl Get(string name, Input<string> id, BucketAccessControlState? state = null, CustomResourceOptions? options = null)
        {
            return new BucketAccessControl(name, id, state, options);
        }
    }

    public sealed class BucketAccessControlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// The entity holding the permission, in one of the following forms:
        /// user-userId
        /// user-email
        /// group-groupId
        /// group-email
        /// domain-domain
        /// project-team-projectId
        /// allUsers
        /// allAuthenticatedUsers
        /// Examples:
        /// The user liz@example.com would be user-liz@example.com.
        /// The group example@googlegroups.com would be
        /// group-example@googlegroups.com.
        /// To refer to all members of the Google Apps for Business domain
        /// example.com, the entity would be domain-example.com.
        /// </summary>
        [Input("entity", required: true)]
        public Input<string> Entity { get; set; } = null!;

        /// <summary>
        /// The access permission for the entity.
        /// Possible values are: `OWNER`, `READER`, `WRITER`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public BucketAccessControlArgs()
        {
        }
        public static new BucketAccessControlArgs Empty => new BucketAccessControlArgs();
    }

    public sealed class BucketAccessControlState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// The domain associated with the entity.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The email address associated with the entity.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The entity holding the permission, in one of the following forms:
        /// user-userId
        /// user-email
        /// group-groupId
        /// group-email
        /// domain-domain
        /// project-team-projectId
        /// allUsers
        /// allAuthenticatedUsers
        /// Examples:
        /// The user liz@example.com would be user-liz@example.com.
        /// The group example@googlegroups.com would be
        /// group-example@googlegroups.com.
        /// To refer to all members of the Google Apps for Business domain
        /// example.com, the entity would be domain-example.com.
        /// </summary>
        [Input("entity")]
        public Input<string>? Entity { get; set; }

        /// <summary>
        /// The access permission for the entity.
        /// Possible values are: `OWNER`, `READER`, `WRITER`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public BucketAccessControlState()
        {
        }
        public static new BucketAccessControlState Empty => new BucketAccessControlState();
    }
}
