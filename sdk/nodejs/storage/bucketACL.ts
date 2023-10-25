// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Authoritatively manages a bucket's ACLs in Google cloud storage service (GCS). For more information see
 * [the official documentation](https://cloud.google.com/storage/docs/access-control/lists)
 * and
 * [API](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls).
 *
 * Bucket ACLs can be managed non authoritatively using the `storageBucketAccessControl` resource. Do not use these two resources in conjunction to manage the same bucket.
 *
 * Permissions can be granted either by ACLs or Cloud IAM policies. In general, permissions granted by Cloud IAM policies do not appear in ACLs, and permissions granted by ACLs do not appear in Cloud IAM policies. The only exception is for ACLs applied directly on a bucket and certain bucket-level Cloud IAM policies, as described in [Cloud IAM relation to ACLs](https://cloud.google.com/storage/docs/access-control/iam#acls).
 *
 * **NOTE** This resource will not remove the `project-owners-<project_id>` entity from the `OWNER` role.
 *
 * ## Example Usage
 *
 * Example creating an ACL on a bucket with one owner, and one reader.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const image_store = new gcp.storage.Bucket("image-store", {location: "EU"});
 * const image_store_acl = new gcp.storage.BucketACL("image-store-acl", {
 *     bucket: image_store.name,
 *     roleEntities: [
 *         "OWNER:user-my.email@gmail.com",
 *         "READER:group-mygroup",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * This resource does not support import.
 */
export class BucketACL extends pulumi.CustomResource {
    /**
     * Get an existing BucketACL resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketACLState, opts?: pulumi.CustomResourceOptions): BucketACL {
        return new BucketACL(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:storage/bucketACL:BucketACL';

    /**
     * Returns true if the given object is an instance of BucketACL.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BucketACL {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BucketACL.__pulumiType;
    }

    /**
     * The name of the bucket it applies to.
     *
     * - - -
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * Configure this ACL to be the default ACL.
     */
    public readonly defaultAcl!: pulumi.Output<string | undefined>;
    /**
     * The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if `roleEntity` is not.
     */
    public readonly predefinedAcl!: pulumi.Output<string | undefined>;
    /**
     * List of role/entity pairs in the form `ROLE:entity`. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)  for more details. Must be set if `predefinedAcl` is not.
     */
    public readonly roleEntities!: pulumi.Output<string[]>;

    /**
     * Create a BucketACL resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketACLArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketACLArgs | BucketACLState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BucketACLState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["defaultAcl"] = state ? state.defaultAcl : undefined;
            resourceInputs["predefinedAcl"] = state ? state.predefinedAcl : undefined;
            resourceInputs["roleEntities"] = state ? state.roleEntities : undefined;
        } else {
            const args = argsOrState as BucketACLArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["defaultAcl"] = args ? args.defaultAcl : undefined;
            resourceInputs["predefinedAcl"] = args ? args.predefinedAcl : undefined;
            resourceInputs["roleEntities"] = args ? args.roleEntities : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BucketACL.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BucketACL resources.
 */
export interface BucketACLState {
    /**
     * The name of the bucket it applies to.
     *
     * - - -
     */
    bucket?: pulumi.Input<string>;
    /**
     * Configure this ACL to be the default ACL.
     */
    defaultAcl?: pulumi.Input<string>;
    /**
     * The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if `roleEntity` is not.
     */
    predefinedAcl?: pulumi.Input<string>;
    /**
     * List of role/entity pairs in the form `ROLE:entity`. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)  for more details. Must be set if `predefinedAcl` is not.
     */
    roleEntities?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a BucketACL resource.
 */
export interface BucketACLArgs {
    /**
     * The name of the bucket it applies to.
     *
     * - - -
     */
    bucket: pulumi.Input<string>;
    /**
     * Configure this ACL to be the default ACL.
     */
    defaultAcl?: pulumi.Input<string>;
    /**
     * The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if `roleEntity` is not.
     */
    predefinedAcl?: pulumi.Input<string>;
    /**
     * List of role/entity pairs in the form `ROLE:entity`. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)  for more details. Must be set if `predefinedAcl` is not.
     */
    roleEntities?: pulumi.Input<pulumi.Input<string>[]>;
}
