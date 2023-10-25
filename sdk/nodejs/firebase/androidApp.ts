// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### Firebase Android App Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basic = new gcp.firebase.AndroidApp("basic", {
 *     project: "my-project-name",
 *     displayName: "Display Name Basic",
 *     packageName: "android.package.app",
 *     sha1Hashes: ["2145bdf698b8715039bd0e83f2069bed435ac21c"],
 *     sha256Hashes: ["2145bdf698b8715039bd0e83f2069bed435ac21ca1b2c3d4e5f6123456789abc"],
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Firebase Android App Custom Api Key
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const android = new gcp.projects.ApiKey("android", {
 *     displayName: "Display Name",
 *     project: "my-project-name",
 *     restrictions: {
 *         androidKeyRestrictions: {
 *             allowedApplications: [{
 *                 packageName: "android.package.app",
 *                 sha1Fingerprint: "2145bdf698b8715039bd0e83f2069bed435ac21c",
 *             }],
 *         },
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * const _default = new gcp.firebase.AndroidApp("default", {
 *     project: "my-project-name",
 *     displayName: "Display Name",
 *     packageName: "android.package.app",
 *     sha1Hashes: ["2145bdf698b8715039bd0e83f2069bed435ac21c"],
 *     sha256Hashes: ["2145bdf698b8715039bd0e83f2069bed435ac21ca1b2c3d4e5f6123456789abc"],
 *     apiKeyId: android.uid,
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * AndroidApp can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:firebase/androidApp:AndroidApp default {{project}} projects/{{project}}/androidApps/{{app_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:firebase/androidApp:AndroidApp default projects/{{project}}/androidApps/{{app_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:firebase/androidApp:AndroidApp default {{project}}/{{project}}/{{app_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:firebase/androidApp:AndroidApp default androidApps/{{app_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:firebase/androidApp:AndroidApp default {{app_id}}
 * ```
 */
export class AndroidApp extends pulumi.CustomResource {
    /**
     * Get an existing AndroidApp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AndroidAppState, opts?: pulumi.CustomResourceOptions): AndroidApp {
        return new AndroidApp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:firebase/androidApp:AndroidApp';

    /**
     * Returns true if the given object is an instance of AndroidApp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AndroidApp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AndroidApp.__pulumiType;
    }

    /**
     * The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the AndroidApp.
     * If apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the AndroidApp.
     * This auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned.
     */
    public readonly apiKeyId!: pulumi.Output<string>;
    /**
     * The globally unique, Firebase-assigned identifier of the AndroidApp.
     * This identifier should be treated as an opaque token, as the data format is not specified.
     */
    public /*out*/ readonly appId!: pulumi.Output<string>;
    /**
     * (Optional) Set to 'ABANDON' to allow the AndroidApp to be untracked from terraform state rather than deleted upon
     * 'terraform destroy'. This is useful because the AndroidApp may be serving traffic. Set to 'DELETE' to delete the
     * AndroidApp. Defaults to 'DELETE'.
     */
    public readonly deletionPolicy!: pulumi.Output<string | undefined>;
    /**
     * The user-assigned display name of the AndroidApp.
     *
     *
     * - - -
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * This checksum is computed by the server based on the value of other fields, and it may be sent
     * with update requests to ensure the client has an up-to-date value before proceeding.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The fully qualified resource name of the AndroidApp, for example:
     * projects/projectId/androidApps/appId
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Immutable. The canonical package name of the Android app as would appear in the Google Play
     * Developer Console.
     */
    public readonly packageName!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The SHA1 certificate hashes for the AndroidApp.
     */
    public readonly sha1Hashes!: pulumi.Output<string[] | undefined>;
    /**
     * The SHA256 certificate hashes for the AndroidApp.
     */
    public readonly sha256Hashes!: pulumi.Output<string[] | undefined>;

    /**
     * Create a AndroidApp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AndroidAppArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AndroidAppArgs | AndroidAppState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AndroidAppState | undefined;
            resourceInputs["apiKeyId"] = state ? state.apiKeyId : undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["deletionPolicy"] = state ? state.deletionPolicy : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["packageName"] = state ? state.packageName : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["sha1Hashes"] = state ? state.sha1Hashes : undefined;
            resourceInputs["sha256Hashes"] = state ? state.sha256Hashes : undefined;
        } else {
            const args = argsOrState as AndroidAppArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["apiKeyId"] = args ? args.apiKeyId : undefined;
            resourceInputs["deletionPolicy"] = args ? args.deletionPolicy : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["packageName"] = args ? args.packageName : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["sha1Hashes"] = args ? args.sha1Hashes : undefined;
            resourceInputs["sha256Hashes"] = args ? args.sha256Hashes : undefined;
            resourceInputs["appId"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AndroidApp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AndroidApp resources.
 */
export interface AndroidAppState {
    /**
     * The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the AndroidApp.
     * If apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the AndroidApp.
     * This auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned.
     */
    apiKeyId?: pulumi.Input<string>;
    /**
     * The globally unique, Firebase-assigned identifier of the AndroidApp.
     * This identifier should be treated as an opaque token, as the data format is not specified.
     */
    appId?: pulumi.Input<string>;
    /**
     * (Optional) Set to 'ABANDON' to allow the AndroidApp to be untracked from terraform state rather than deleted upon
     * 'terraform destroy'. This is useful because the AndroidApp may be serving traffic. Set to 'DELETE' to delete the
     * AndroidApp. Defaults to 'DELETE'.
     */
    deletionPolicy?: pulumi.Input<string>;
    /**
     * The user-assigned display name of the AndroidApp.
     *
     *
     * - - -
     */
    displayName?: pulumi.Input<string>;
    /**
     * This checksum is computed by the server based on the value of other fields, and it may be sent
     * with update requests to ensure the client has an up-to-date value before proceeding.
     */
    etag?: pulumi.Input<string>;
    /**
     * The fully qualified resource name of the AndroidApp, for example:
     * projects/projectId/androidApps/appId
     */
    name?: pulumi.Input<string>;
    /**
     * Immutable. The canonical package name of the Android app as would appear in the Google Play
     * Developer Console.
     */
    packageName?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The SHA1 certificate hashes for the AndroidApp.
     */
    sha1Hashes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The SHA256 certificate hashes for the AndroidApp.
     */
    sha256Hashes?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a AndroidApp resource.
 */
export interface AndroidAppArgs {
    /**
     * The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the AndroidApp.
     * If apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the AndroidApp.
     * This auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned.
     */
    apiKeyId?: pulumi.Input<string>;
    /**
     * (Optional) Set to 'ABANDON' to allow the AndroidApp to be untracked from terraform state rather than deleted upon
     * 'terraform destroy'. This is useful because the AndroidApp may be serving traffic. Set to 'DELETE' to delete the
     * AndroidApp. Defaults to 'DELETE'.
     */
    deletionPolicy?: pulumi.Input<string>;
    /**
     * The user-assigned display name of the AndroidApp.
     *
     *
     * - - -
     */
    displayName: pulumi.Input<string>;
    /**
     * Immutable. The canonical package name of the Android app as would appear in the Google Play
     * Developer Console.
     */
    packageName?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The SHA1 certificate hashes for the AndroidApp.
     */
    sha1Hashes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The SHA256 certificate hashes for the AndroidApp.
     */
    sha256Hashes?: pulumi.Input<pulumi.Input<string>[]>;
}
