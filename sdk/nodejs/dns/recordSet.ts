// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * DNS record sets can be imported using either of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:dns/recordSet:RecordSet frontend projects/{{project}}/managedZones/{{zone}}/rrsets/{{name}}/{{type}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dns/recordSet:RecordSet frontend {{project}}/{{zone}}/{{name}}/{{type}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dns/recordSet:RecordSet frontend {{zone}}/{{name}}/{{type}}
 * ```
 *
 *  NoteThe record name must include the trailing dot at the end.
 */
export class RecordSet extends pulumi.CustomResource {
    /**
     * Get an existing RecordSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RecordSetState, opts?: pulumi.CustomResourceOptions): RecordSet {
        return new RecordSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dns/recordSet:RecordSet';

    /**
     * Returns true if the given object is an instance of RecordSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RecordSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RecordSet.__pulumiType;
    }

    /**
     * The name of the zone in which this record set will
     * reside.
     */
    public readonly managedZone!: pulumi.Output<string>;
    /**
     * The DNS name this record set will apply to.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The string data for the records in this record set whose meaning depends on the DNS type. For TXT record, if the string
     * data contains spaces, add surrounding \" if you don't want your string to get split on spaces. To specify a single
     * record value longer than 255 characters such as a TXT record for DKIM, add \"\" inside the Terraform configuration
     * string (e.g. "first255characters\"\"morecharacters").
     */
    public readonly rrdatas!: pulumi.Output<string[]>;
    /**
     * The time-to-live of this record set (seconds).
     */
    public readonly ttl!: pulumi.Output<number | undefined>;
    /**
     * The DNS record set type.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a RecordSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RecordSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RecordSetArgs | RecordSetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RecordSetState | undefined;
            resourceInputs["managedZone"] = state ? state.managedZone : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["rrdatas"] = state ? state.rrdatas : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as RecordSetArgs | undefined;
            if ((!args || args.managedZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'managedZone'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.rrdatas === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rrdatas'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["managedZone"] = args ? args.managedZone : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["rrdatas"] = args ? args.rrdatas : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RecordSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RecordSet resources.
 */
export interface RecordSetState {
    /**
     * The name of the zone in which this record set will
     * reside.
     */
    managedZone?: pulumi.Input<string>;
    /**
     * The DNS name this record set will apply to.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The string data for the records in this record set whose meaning depends on the DNS type. For TXT record, if the string
     * data contains spaces, add surrounding \" if you don't want your string to get split on spaces. To specify a single
     * record value longer than 255 characters such as a TXT record for DKIM, add \"\" inside the Terraform configuration
     * string (e.g. "first255characters\"\"morecharacters").
     */
    rrdatas?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The time-to-live of this record set (seconds).
     */
    ttl?: pulumi.Input<number>;
    /**
     * The DNS record set type.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RecordSet resource.
 */
export interface RecordSetArgs {
    /**
     * The name of the zone in which this record set will
     * reside.
     */
    managedZone: pulumi.Input<string>;
    /**
     * The DNS name this record set will apply to.
     */
    name: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The string data for the records in this record set whose meaning depends on the DNS type. For TXT record, if the string
     * data contains spaces, add surrounding \" if you don't want your string to get split on spaces. To specify a single
     * record value longer than 255 characters such as a TXT record for DKIM, add \"\" inside the Terraform configuration
     * string (e.g. "first255characters\"\"morecharacters").
     */
    rrdatas: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The time-to-live of this record set (seconds).
     */
    ttl?: pulumi.Input<number>;
    /**
     * The DNS record set type.
     */
    type: pulumi.Input<string>;
}
