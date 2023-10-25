// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The Eventarc GoogleChannelConfig resource
 *
 * ## Example Usage
 * ### Basic
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const testProject = gcp.organizations.getProject({
 *     projectId: "my-project-name",
 * });
 * const testKeyRing = gcp.kms.getKMSKeyRing({
 *     name: "keyring",
 *     location: "us-west1",
 * });
 * const key = testKeyRing.then(testKeyRing => gcp.kms.getKMSCryptoKey({
 *     name: "key",
 *     keyRing: testKeyRing.id,
 * }));
 * const key1Member = new gcp.kms.CryptoKeyIAMMember("key1Member", {
 *     cryptoKeyId: data.google_kms_crypto_key.key1.id,
 *     role: "roles/cloudkms.cryptoKeyEncrypterDecrypter",
 *     member: testProject.then(testProject => `serviceAccount:service-${testProject.number}@gcp-sa-eventarc.iam.gserviceaccount.com`),
 * });
 * const primary = new gcp.eventarc.GoogleChannelConfig("primary", {
 *     location: "us-west1",
 *     project: testProject.then(testProject => testProject.projectId),
 *     cryptoKeyName: data.google_kms_crypto_key.key1.id,
 * }, {
 *     dependsOn: [key1Member],
 * });
 * ```
 *
 * ## Import
 *
 * GoogleChannelConfig can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:eventarc/googleChannelConfig:GoogleChannelConfig default projects/{{project}}/locations/{{location}}/googleChannelConfig
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:eventarc/googleChannelConfig:GoogleChannelConfig default {{project}}/{{location}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:eventarc/googleChannelConfig:GoogleChannelConfig default {{location}}
 * ```
 */
export class GoogleChannelConfig extends pulumi.CustomResource {
    /**
     * Get an existing GoogleChannelConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GoogleChannelConfigState, opts?: pulumi.CustomResourceOptions): GoogleChannelConfig {
        return new GoogleChannelConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:eventarc/googleChannelConfig:GoogleChannelConfig';

    /**
     * Returns true if the given object is an instance of GoogleChannelConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GoogleChannelConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GoogleChannelConfig.__pulumiType;
    }

    /**
     * Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern `projects/*&#47;locations/*&#47;keyRings/*&#47;cryptoKeys/*`.
     */
    public readonly cryptoKeyName!: pulumi.Output<string | undefined>;
    /**
     * The location for the resource
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Required. The resource name of the config. Must be in the format of, `projects/{project}/locations/{location}/googleChannelConfig`.
     *
     *
     *
     * - - -
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project for the resource
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Output only. The last-modified time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a GoogleChannelConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GoogleChannelConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GoogleChannelConfigArgs | GoogleChannelConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GoogleChannelConfigState | undefined;
            resourceInputs["cryptoKeyName"] = state ? state.cryptoKeyName : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as GoogleChannelConfigArgs | undefined;
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            resourceInputs["cryptoKeyName"] = args ? args.cryptoKeyName : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GoogleChannelConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GoogleChannelConfig resources.
 */
export interface GoogleChannelConfigState {
    /**
     * Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern `projects/*&#47;locations/*&#47;keyRings/*&#47;cryptoKeys/*`.
     */
    cryptoKeyName?: pulumi.Input<string>;
    /**
     * The location for the resource
     */
    location?: pulumi.Input<string>;
    /**
     * Required. The resource name of the config. Must be in the format of, `projects/{project}/locations/{location}/googleChannelConfig`.
     *
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
    /**
     * Output only. The last-modified time.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GoogleChannelConfig resource.
 */
export interface GoogleChannelConfigArgs {
    /**
     * Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern `projects/*&#47;locations/*&#47;keyRings/*&#47;cryptoKeys/*`.
     */
    cryptoKeyName?: pulumi.Input<string>;
    /**
     * The location for the resource
     */
    location: pulumi.Input<string>;
    /**
     * Required. The resource name of the config. Must be in the format of, `projects/{project}/locations/{location}/googleChannelConfig`.
     *
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
}
