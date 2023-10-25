// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Google Cloud Filestore instance.
 *
 * To get more information about Instance, see:
 *
 * * [API documentation](https://cloud.google.com/filestore/docs/reference/rest/v1beta1/projects.locations.instances/create)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/filestore/docs/creating-instances)
 *     * [Use with Kubernetes](https://cloud.google.com/filestore/docs/accessing-fileshares)
 *     * [Copying Data In/Out](https://cloud.google.com/filestore/docs/copying-data)
 *
 * ## Example Usage
 * ### Filestore Instance Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const instance = new gcp.filestore.Instance("instance", {
 *     fileShares: {
 *         capacityGb: 1024,
 *         name: "share1",
 *     },
 *     location: "us-central1-b",
 *     networks: [{
 *         modes: ["MODE_IPV4"],
 *         network: "default",
 *     }],
 *     tier: "BASIC_HDD",
 * });
 * ```
 * ### Filestore Instance Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const instance = new gcp.filestore.Instance("instance", {
 *     fileShares: {
 *         capacityGb: 2560,
 *         name: "share1",
 *         nfsExportOptions: [
 *             {
 *                 accessMode: "READ_WRITE",
 *                 ipRanges: ["10.0.0.0/24"],
 *                 squashMode: "NO_ROOT_SQUASH",
 *             },
 *             {
 *                 accessMode: "READ_ONLY",
 *                 anonGid: 456,
 *                 anonUid: 123,
 *                 ipRanges: ["10.10.0.0/24"],
 *                 squashMode: "ROOT_SQUASH",
 *             },
 *         ],
 *     },
 *     location: "us-central1-b",
 *     networks: [{
 *         connectMode: "DIRECT_PEERING",
 *         modes: ["MODE_IPV4"],
 *         network: "default",
 *     }],
 *     tier: "BASIC_SSD",
 * });
 * ```
 * ### Filestore Instance Enterprise
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const filestoreKeyring = new gcp.kms.KeyRing("filestoreKeyring", {location: "us-central1"});
 * const filestoreKey = new gcp.kms.CryptoKey("filestoreKey", {keyRing: filestoreKeyring.id});
 * const instance = new gcp.filestore.Instance("instance", {
 *     location: "us-central1",
 *     tier: "ENTERPRISE",
 *     fileShares: {
 *         capacityGb: 1024,
 *         name: "share1",
 *     },
 *     networks: [{
 *         network: "default",
 *         modes: ["MODE_IPV4"],
 *     }],
 *     kmsKeyName: filestoreKey.id,
 * });
 * ```
 *
 * ## Import
 *
 * Instance can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:filestore/instance:Instance default projects/{{project}}/locations/{{location}}/instances/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:filestore/instance:Instance default {{project}}/{{location}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:filestore/instance:Instance default {{location}}/{{name}}
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:filestore/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * A description of the instance.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Server-specified ETag for the instance resource to prevent
     * simultaneous updates from overwriting each other.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * File system shares on the instance. For this version, only a
     * single file share is supported.
     * Structure is documented below.
     */
    public readonly fileShares!: pulumi.Output<outputs.filestore.InstanceFileShares>;
    /**
     * KMS key name used for data encryption.
     */
    public readonly kmsKeyName!: pulumi.Output<string | undefined>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The resource name of the instance.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * VPC networks to which the instance is connected. For this version,
     * only a single network is supported.
     * Structure is documented below.
     */
    public readonly networks!: pulumi.Output<outputs.filestore.InstanceNetwork[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The service tier of the instance.
     * Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD, ZONAL and ENTERPRISE
     */
    public readonly tier!: pulumi.Output<string>;
    /**
     * (Optional, Deprecated)
     * The name of the Filestore zone of the instance.
     *
     * > **Warning:** `zone` is deprecated and will be removed in a future major release. Use `location` instead.
     *
     * @deprecated `zone` is deprecated and will be removed in a future major release. Use `location` instead.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["fileShares"] = state ? state.fileShares : undefined;
            resourceInputs["kmsKeyName"] = state ? state.kmsKeyName : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networks"] = state ? state.networks : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["tier"] = state ? state.tier : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.fileShares === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileShares'");
            }
            if ((!args || args.networks === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networks'");
            }
            if ((!args || args.tier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tier'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["fileShares"] = args ? args.fileShares : undefined;
            resourceInputs["kmsKeyName"] = args ? args.kmsKeyName : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networks"] = args ? args.networks : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["tier"] = args ? args.tier : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    createTime?: pulumi.Input<string>;
    /**
     * A description of the instance.
     */
    description?: pulumi.Input<string>;
    /**
     * Server-specified ETag for the instance resource to prevent
     * simultaneous updates from overwriting each other.
     */
    etag?: pulumi.Input<string>;
    /**
     * File system shares on the instance. For this version, only a
     * single file share is supported.
     * Structure is documented below.
     */
    fileShares?: pulumi.Input<inputs.filestore.InstanceFileShares>;
    /**
     * KMS key name used for data encryption.
     */
    kmsKeyName?: pulumi.Input<string>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
     */
    location?: pulumi.Input<string>;
    /**
     * The resource name of the instance.
     */
    name?: pulumi.Input<string>;
    /**
     * VPC networks to which the instance is connected. For this version,
     * only a single network is supported.
     * Structure is documented below.
     */
    networks?: pulumi.Input<pulumi.Input<inputs.filestore.InstanceNetwork>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The service tier of the instance.
     * Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD, ZONAL and ENTERPRISE
     */
    tier?: pulumi.Input<string>;
    /**
     * (Optional, Deprecated)
     * The name of the Filestore zone of the instance.
     *
     * > **Warning:** `zone` is deprecated and will be removed in a future major release. Use `location` instead.
     *
     * @deprecated `zone` is deprecated and will be removed in a future major release. Use `location` instead.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * A description of the instance.
     */
    description?: pulumi.Input<string>;
    /**
     * File system shares on the instance. For this version, only a
     * single file share is supported.
     * Structure is documented below.
     */
    fileShares: pulumi.Input<inputs.filestore.InstanceFileShares>;
    /**
     * KMS key name used for data encryption.
     */
    kmsKeyName?: pulumi.Input<string>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
     */
    location?: pulumi.Input<string>;
    /**
     * The resource name of the instance.
     */
    name?: pulumi.Input<string>;
    /**
     * VPC networks to which the instance is connected. For this version,
     * only a single network is supported.
     * Structure is documented below.
     */
    networks: pulumi.Input<pulumi.Input<inputs.filestore.InstanceNetwork>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The service tier of the instance.
     * Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD, ZONAL and ENTERPRISE
     */
    tier: pulumi.Input<string>;
    /**
     * (Optional, Deprecated)
     * The name of the Filestore zone of the instance.
     *
     * > **Warning:** `zone` is deprecated and will be removed in a future major release. Use `location` instead.
     *
     * @deprecated `zone` is deprecated and will be removed in a future major release. Use `location` instead.
     */
    zone?: pulumi.Input<string>;
}
