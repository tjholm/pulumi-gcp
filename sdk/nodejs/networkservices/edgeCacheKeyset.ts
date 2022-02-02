// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * EdgeCacheKeyset represents a collection of public keys used for validating signed requests.
 *
 * > **Warning:** All arguments including `public_key.public_key.value` will be stored in the raw
 * state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 *
 * ## Example Usage
 * ### Network Services Edge Cache Keyset Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultEdgeCacheKeyset = new gcp.networkservices.EdgeCacheKeyset("default", {
 *     description: "The default keyset",
 *     publicKeys: [
 *         {
 *             id: "my-public-key",
 *             value: "FHsTyFHNmvNpw4o7-rp-M1yqMyBF8vXSBRkZtkQ0RKY",
 *         },
 *         {
 *             id: "my-public-key-2",
 *             value: "hzd03llxB1u5FOLKFkZ6_wCJqC7jtN0bg7xlBqS6WVM",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * EdgeCacheKeyset can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:networkservices/edgeCacheKeyset:EdgeCacheKeyset default projects/{{project}}/locations/global/edgeCacheKeysets/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:networkservices/edgeCacheKeyset:EdgeCacheKeyset default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:networkservices/edgeCacheKeyset:EdgeCacheKeyset default {{name}}
 * ```
 */
export class EdgeCacheKeyset extends pulumi.CustomResource {
    /**
     * Get an existing EdgeCacheKeyset resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EdgeCacheKeysetState, opts?: pulumi.CustomResourceOptions): EdgeCacheKeyset {
        return new EdgeCacheKeyset(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:networkservices/edgeCacheKeyset:EdgeCacheKeyset';

    /**
     * Returns true if the given object is an instance of EdgeCacheKeyset.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EdgeCacheKeyset {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EdgeCacheKeyset.__pulumiType;
    }

    /**
     * A human-readable description of the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Set of label tags associated with the EdgeCache resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Name of the resource; provided by the client when the resource is created.
     * The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
     * and all following characters must be a dash, underscore, letter or digit.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * An ordered list of Ed25519 public keys to use for validating signed requests.
     * You must specify at least one (1) key, and may have up to three (3) keys.
     * Ed25519 public keys are not secret, and only allow Google to validate a request was signed by your corresponding private key.
     * You should ensure that the private key is kept secret, and that only authorized users can add public keys to a keyset.
     * Structure is documented below.
     */
    public readonly publicKeys!: pulumi.Output<outputs.networkservices.EdgeCacheKeysetPublicKey[]>;

    /**
     * Create a EdgeCacheKeyset resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EdgeCacheKeysetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EdgeCacheKeysetArgs | EdgeCacheKeysetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EdgeCacheKeysetState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["publicKeys"] = state ? state.publicKeys : undefined;
        } else {
            const args = argsOrState as EdgeCacheKeysetArgs | undefined;
            if ((!args || args.publicKeys === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publicKeys'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["publicKeys"] = args ? args.publicKeys : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EdgeCacheKeyset.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EdgeCacheKeyset resources.
 */
export interface EdgeCacheKeysetState {
    /**
     * A human-readable description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Set of label tags associated with the EdgeCache resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the resource; provided by the client when the resource is created.
     * The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
     * and all following characters must be a dash, underscore, letter or digit.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * An ordered list of Ed25519 public keys to use for validating signed requests.
     * You must specify at least one (1) key, and may have up to three (3) keys.
     * Ed25519 public keys are not secret, and only allow Google to validate a request was signed by your corresponding private key.
     * You should ensure that the private key is kept secret, and that only authorized users can add public keys to a keyset.
     * Structure is documented below.
     */
    publicKeys?: pulumi.Input<pulumi.Input<inputs.networkservices.EdgeCacheKeysetPublicKey>[]>;
}

/**
 * The set of arguments for constructing a EdgeCacheKeyset resource.
 */
export interface EdgeCacheKeysetArgs {
    /**
     * A human-readable description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Set of label tags associated with the EdgeCache resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the resource; provided by the client when the resource is created.
     * The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
     * and all following characters must be a dash, underscore, letter or digit.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * An ordered list of Ed25519 public keys to use for validating signed requests.
     * You must specify at least one (1) key, and may have up to three (3) keys.
     * Ed25519 public keys are not secret, and only allow Google to validate a request was signed by your corresponding private key.
     * You should ensure that the private key is kept secret, and that only authorized users can add public keys to a keyset.
     * Structure is documented below.
     */
    publicKeys: pulumi.Input<pulumi.Input<inputs.networkservices.EdgeCacheKeysetPublicKey>[]>;
}
