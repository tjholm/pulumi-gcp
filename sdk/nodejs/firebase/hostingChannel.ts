// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### Firebasehosting Channel Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultHostingSite = new gcp.firebase.HostingSite("defaultHostingSite", {
 *     project: "my-project-name",
 *     siteId: "site-with-channel",
 * }, {
 *     provider: google_beta,
 * });
 * const defaultHostingChannel = new gcp.firebase.HostingChannel("defaultHostingChannel", {
 *     siteId: defaultHostingSite.siteId,
 *     channelId: "channel-basic",
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Firebasehosting Channel Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.firebase.HostingSite("default", {
 *     project: "my-project-name",
 *     siteId: "site-with-channel",
 * }, {
 *     provider: google_beta,
 * });
 * const full = new gcp.firebase.HostingChannel("full", {
 *     siteId: _default.siteId,
 *     channelId: "channel-full",
 *     ttl: "86400s",
 *     retainedReleaseCount: 20,
 *     labels: {
 *         "some-key": "some-value",
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * Channel can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:firebase/hostingChannel:HostingChannel default sites/{{site_id}}/channels/{{channel_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:firebase/hostingChannel:HostingChannel default {{site_id}}/{{channel_id}}
 * ```
 */
export class HostingChannel extends pulumi.CustomResource {
    /**
     * Get an existing HostingChannel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostingChannelState, opts?: pulumi.CustomResourceOptions): HostingChannel {
        return new HostingChannel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:firebase/hostingChannel:HostingChannel';

    /**
     * Returns true if the given object is an instance of HostingChannel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HostingChannel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HostingChannel.__pulumiType;
    }

    /**
     * Required. Immutable. A unique ID within the site that identifies the channel.
     *
     *
     * - - -
     */
    public readonly channelId!: pulumi.Output<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The time at which the channel will be automatically deleted. If null, the channel
     * will not be automatically deleted. This field is present in the output whether it's
     * set directly or via the `ttl` field.
     */
    public readonly expireTime!: pulumi.Output<string>;
    /**
     * Text labels used for extra metadata and/or filtering
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The fully-qualified resource name for the channel, in the format:
     * sites/SITE_ID/channels/CHANNEL_ID
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    public /*out*/ readonly pulumiLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The number of previous releases to retain on the channel for rollback or other
     * purposes. Must be a number between 1-100. Defaults to 10 for new channels.
     */
    public readonly retainedReleaseCount!: pulumi.Output<number>;
    /**
     * Required. The ID of the site in which to create this channel.
     */
    public readonly siteId!: pulumi.Output<string>;
    /**
     * Input only. A time-to-live for this channel. Sets `expireTime` to the provided
     * duration past the time of the request. A duration in seconds with up to nine fractional
     * digits, terminated by 's'. Example: "86400s" (one day).
     */
    public readonly ttl!: pulumi.Output<string | undefined>;

    /**
     * Create a HostingChannel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostingChannelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostingChannelArgs | HostingChannelState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HostingChannelState | undefined;
            resourceInputs["channelId"] = state ? state.channelId : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["expireTime"] = state ? state.expireTime : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["retainedReleaseCount"] = state ? state.retainedReleaseCount : undefined;
            resourceInputs["siteId"] = state ? state.siteId : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
        } else {
            const args = argsOrState as HostingChannelArgs | undefined;
            if ((!args || args.channelId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'channelId'");
            }
            if ((!args || args.siteId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'siteId'");
            }
            resourceInputs["channelId"] = args ? args.channelId : undefined;
            resourceInputs["expireTime"] = args ? args.expireTime : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["retainedReleaseCount"] = args ? args.retainedReleaseCount : undefined;
            resourceInputs["siteId"] = args ? args.siteId : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(HostingChannel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HostingChannel resources.
 */
export interface HostingChannelState {
    /**
     * Required. Immutable. A unique ID within the site that identifies the channel.
     *
     *
     * - - -
     */
    channelId?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The time at which the channel will be automatically deleted. If null, the channel
     * will not be automatically deleted. This field is present in the output whether it's
     * set directly or via the `ttl` field.
     */
    expireTime?: pulumi.Input<string>;
    /**
     * Text labels used for extra metadata and/or filtering
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The fully-qualified resource name for the channel, in the format:
     * sites/SITE_ID/channels/CHANNEL_ID
     */
    name?: pulumi.Input<string>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    pulumiLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The number of previous releases to retain on the channel for rollback or other
     * purposes. Must be a number between 1-100. Defaults to 10 for new channels.
     */
    retainedReleaseCount?: pulumi.Input<number>;
    /**
     * Required. The ID of the site in which to create this channel.
     */
    siteId?: pulumi.Input<string>;
    /**
     * Input only. A time-to-live for this channel. Sets `expireTime` to the provided
     * duration past the time of the request. A duration in seconds with up to nine fractional
     * digits, terminated by 's'. Example: "86400s" (one day).
     */
    ttl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HostingChannel resource.
 */
export interface HostingChannelArgs {
    /**
     * Required. Immutable. A unique ID within the site that identifies the channel.
     *
     *
     * - - -
     */
    channelId: pulumi.Input<string>;
    /**
     * The time at which the channel will be automatically deleted. If null, the channel
     * will not be automatically deleted. This field is present in the output whether it's
     * set directly or via the `ttl` field.
     */
    expireTime?: pulumi.Input<string>;
    /**
     * Text labels used for extra metadata and/or filtering
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The number of previous releases to retain on the channel for rollback or other
     * purposes. Must be a number between 1-100. Defaults to 10 for new channels.
     */
    retainedReleaseCount?: pulumi.Input<number>;
    /**
     * Required. The ID of the site in which to create this channel.
     */
    siteId: pulumi.Input<string>;
    /**
     * Input only. A time-to-live for this channel. Sets `expireTime` to the provided
     * duration past the time of the request. A duration in seconds with up to nine fractional
     * digits, terminated by 's'. Example: "86400s" (one day).
     */
    ttl?: pulumi.Input<string>;
}
