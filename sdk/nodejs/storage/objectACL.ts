// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Authoritatively manages the access control list (ACL) for an object in a Google
 * Cloud Storage (GCS) bucket. Removing a `gcp.storage.ObjectACL` sets the
 * acl to the `private` [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl).
 *
 * For more information see
 * [the official documentation](https://cloud.google.com/storage/docs/access-control/lists)
 * and
 * [API](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls).
 *
 * > Want fine-grained control over object ACLs? Use `gcp.storage.ObjectAccessControl` to control individual
 * role entity pairs.
 *
 * ## Example Usage
 *
 * Create an object ACL with one owner and one reader.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const image_store = new gcp.storage.Bucket("image-store", {location: "EU"});
 * const image = new gcp.storage.BucketObject("image", {
 *     bucket: image_store.name,
 *     source: new pulumi.asset.FileAsset("image1.jpg"),
 * });
 * const image_store_acl = new gcp.storage.ObjectACL("image-store-acl", {
 *     bucket: image_store.name,
 *     object: image.outputName,
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
export class ObjectACL extends pulumi.CustomResource {
    /**
     * Get an existing ObjectACL resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ObjectACLState, opts?: pulumi.CustomResourceOptions): ObjectACL {
        return new ObjectACL(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:storage/objectACL:ObjectACL';

    /**
     * Returns true if the given object is an instance of ObjectACL.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ObjectACL {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ObjectACL.__pulumiType;
    }

    /**
     * The name of the bucket the object is stored in.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * The name of the object to apply the acl to.
     */
    public readonly object!: pulumi.Output<string>;
    /**
     * The "canned" [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl) to apply. Must be set if `roleEntity` is not.
     */
    public readonly predefinedAcl!: pulumi.Output<string | undefined>;
    /**
     * List of role/entity pairs in the form `ROLE:entity`. See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
     * Must be set if `predefinedAcl` is not.
     */
    public readonly roleEntities!: pulumi.Output<string[]>;

    /**
     * Create a ObjectACL resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ObjectACLArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ObjectACLArgs | ObjectACLState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ObjectACLState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["object"] = state ? state.object : undefined;
            resourceInputs["predefinedAcl"] = state ? state.predefinedAcl : undefined;
            resourceInputs["roleEntities"] = state ? state.roleEntities : undefined;
        } else {
            const args = argsOrState as ObjectACLArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.object === undefined) && !opts.urn) {
                throw new Error("Missing required property 'object'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["object"] = args ? args.object : undefined;
            resourceInputs["predefinedAcl"] = args ? args.predefinedAcl : undefined;
            resourceInputs["roleEntities"] = args ? args.roleEntities : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ObjectACL.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ObjectACL resources.
 */
export interface ObjectACLState {
    /**
     * The name of the bucket the object is stored in.
     */
    bucket?: pulumi.Input<string>;
    /**
     * The name of the object to apply the acl to.
     */
    object?: pulumi.Input<string>;
    /**
     * The "canned" [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl) to apply. Must be set if `roleEntity` is not.
     */
    predefinedAcl?: pulumi.Input<string>;
    /**
     * List of role/entity pairs in the form `ROLE:entity`. See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
     * Must be set if `predefinedAcl` is not.
     */
    roleEntities?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ObjectACL resource.
 */
export interface ObjectACLArgs {
    /**
     * The name of the bucket the object is stored in.
     */
    bucket: pulumi.Input<string>;
    /**
     * The name of the object to apply the acl to.
     */
    object: pulumi.Input<string>;
    /**
     * The "canned" [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl) to apply. Must be set if `roleEntity` is not.
     */
    predefinedAcl?: pulumi.Input<string>;
    /**
     * List of role/entity pairs in the form `ROLE:entity`. See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
     * Must be set if `predefinedAcl` is not.
     */
    roleEntities?: pulumi.Input<pulumi.Input<string>[]>;
}
