// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A `KeyRingImportJob` can be used to create `CryptoKeys` and `CryptoKeyVersions` using pre-existing
 * key material, generated outside of Cloud KMS. A `KeyRingImportJob` expires 3 days after it is created.
 * Once expired, Cloud KMS will no longer be able to import or unwrap any key material that
 * was wrapped with the `KeyRingImportJob`'s public key.
 *
 * > **Note:** KeyRingImportJobs cannot be deleted from Google Cloud Platform.
 * Destroying a provider-managed KeyRingImportJob will remove it from state but
 * *will not delete the resource from the project.*
 *
 * To get more information about KeyRingImportJob, see:
 *
 * * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.importJobs)
 * * How-to Guides
 *     * [Importing a key](https://cloud.google.com/kms/docs/importing-a-key)
 *
 * ## Example Usage
 *
 * ## Import
 *
 * KeyRingImportJob can be imported using any of these accepted formats* `{{name}}` When using the `pulumi import` command, KeyRingImportJob can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:kms/keyRingImportJob:KeyRingImportJob default {{name}}
 * ```
 */
export class KeyRingImportJob extends pulumi.CustomResource {
    /**
     * Get an existing KeyRingImportJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeyRingImportJobState, opts?: pulumi.CustomResourceOptions): KeyRingImportJob {
        return new KeyRingImportJob(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:kms/keyRingImportJob:KeyRingImportJob';

    /**
     * Returns true if the given object is an instance of KeyRingImportJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KeyRingImportJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KeyRingImportJob.__pulumiType;
    }

    /**
     * Statement that was generated and signed by the key creator (for example, an HSM) at key creation time.
     * Use this statement to verify attributes of the key as stored on the HSM, independently of Google.
     * Only present if the chosen ImportMethod is one with a protection level of HSM.
     * Structure is documented below.
     */
    public /*out*/ readonly attestations!: pulumi.Output<outputs.kms.KeyRingImportJobAttestation[]>;
    /**
     * The time at which this resource is scheduled for expiration and can no longer be used.
     * This is in RFC3339 text format.
     */
    public /*out*/ readonly expireTime!: pulumi.Output<string>;
    /**
     * It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}
     *
     *
     * - - -
     */
    public readonly importJobId!: pulumi.Output<string>;
    /**
     * The wrapping method to be used for incoming key material.
     * Possible values are: `RSA_OAEP_3072_SHA1_AES_256`, `RSA_OAEP_4096_SHA1_AES_256`.
     */
    public readonly importMethod!: pulumi.Output<string>;
    /**
     * The KeyRing that this import job belongs to.
     * Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
     */
    public readonly keyRing!: pulumi.Output<string>;
    /**
     * The resource name for this ImportJob in the format projects/*&#47;locations/*&#47;keyRings/*&#47;importJobs/*.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The protection level of the ImportJob. This must match the protectionLevel of the
     * versionTemplate on the CryptoKey you attempt to import into.
     * Possible values are: `SOFTWARE`, `HSM`, `EXTERNAL`.
     */
    public readonly protectionLevel!: pulumi.Output<string>;
    /**
     * The public key with which to wrap key material prior to import. Only returned if state is `ACTIVE`.
     * Structure is documented below.
     */
    public /*out*/ readonly publicKeys!: pulumi.Output<outputs.kms.KeyRingImportJobPublicKey[]>;
    /**
     * The current state of the ImportJob, indicating if it can be used.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a KeyRingImportJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeyRingImportJobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeyRingImportJobArgs | KeyRingImportJobState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KeyRingImportJobState | undefined;
            resourceInputs["attestations"] = state ? state.attestations : undefined;
            resourceInputs["expireTime"] = state ? state.expireTime : undefined;
            resourceInputs["importJobId"] = state ? state.importJobId : undefined;
            resourceInputs["importMethod"] = state ? state.importMethod : undefined;
            resourceInputs["keyRing"] = state ? state.keyRing : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protectionLevel"] = state ? state.protectionLevel : undefined;
            resourceInputs["publicKeys"] = state ? state.publicKeys : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as KeyRingImportJobArgs | undefined;
            if ((!args || args.importJobId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'importJobId'");
            }
            if ((!args || args.importMethod === undefined) && !opts.urn) {
                throw new Error("Missing required property 'importMethod'");
            }
            if ((!args || args.keyRing === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyRing'");
            }
            if ((!args || args.protectionLevel === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protectionLevel'");
            }
            resourceInputs["importJobId"] = args ? args.importJobId : undefined;
            resourceInputs["importMethod"] = args ? args.importMethod : undefined;
            resourceInputs["keyRing"] = args ? args.keyRing : undefined;
            resourceInputs["protectionLevel"] = args ? args.protectionLevel : undefined;
            resourceInputs["attestations"] = undefined /*out*/;
            resourceInputs["expireTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["publicKeys"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KeyRingImportJob.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KeyRingImportJob resources.
 */
export interface KeyRingImportJobState {
    /**
     * Statement that was generated and signed by the key creator (for example, an HSM) at key creation time.
     * Use this statement to verify attributes of the key as stored on the HSM, independently of Google.
     * Only present if the chosen ImportMethod is one with a protection level of HSM.
     * Structure is documented below.
     */
    attestations?: pulumi.Input<pulumi.Input<inputs.kms.KeyRingImportJobAttestation>[]>;
    /**
     * The time at which this resource is scheduled for expiration and can no longer be used.
     * This is in RFC3339 text format.
     */
    expireTime?: pulumi.Input<string>;
    /**
     * It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}
     *
     *
     * - - -
     */
    importJobId?: pulumi.Input<string>;
    /**
     * The wrapping method to be used for incoming key material.
     * Possible values are: `RSA_OAEP_3072_SHA1_AES_256`, `RSA_OAEP_4096_SHA1_AES_256`.
     */
    importMethod?: pulumi.Input<string>;
    /**
     * The KeyRing that this import job belongs to.
     * Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
     */
    keyRing?: pulumi.Input<string>;
    /**
     * The resource name for this ImportJob in the format projects/*&#47;locations/*&#47;keyRings/*&#47;importJobs/*.
     */
    name?: pulumi.Input<string>;
    /**
     * The protection level of the ImportJob. This must match the protectionLevel of the
     * versionTemplate on the CryptoKey you attempt to import into.
     * Possible values are: `SOFTWARE`, `HSM`, `EXTERNAL`.
     */
    protectionLevel?: pulumi.Input<string>;
    /**
     * The public key with which to wrap key material prior to import. Only returned if state is `ACTIVE`.
     * Structure is documented below.
     */
    publicKeys?: pulumi.Input<pulumi.Input<inputs.kms.KeyRingImportJobPublicKey>[]>;
    /**
     * The current state of the ImportJob, indicating if it can be used.
     */
    state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KeyRingImportJob resource.
 */
export interface KeyRingImportJobArgs {
    /**
     * It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}
     *
     *
     * - - -
     */
    importJobId: pulumi.Input<string>;
    /**
     * The wrapping method to be used for incoming key material.
     * Possible values are: `RSA_OAEP_3072_SHA1_AES_256`, `RSA_OAEP_4096_SHA1_AES_256`.
     */
    importMethod: pulumi.Input<string>;
    /**
     * The KeyRing that this import job belongs to.
     * Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
     */
    keyRing: pulumi.Input<string>;
    /**
     * The protection level of the ImportJob. This must match the protectionLevel of the
     * versionTemplate on the CryptoKey you attempt to import into.
     * Possible values are: `SOFTWARE`, `HSM`, `EXTERNAL`.
     */
    protectionLevel: pulumi.Input<string>;
}
