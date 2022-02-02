// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A Container Analysis note is a high-level piece of metadata that
 * describes a type of analysis that can be done for a resource.
 *
 * To get more information about Note, see:
 *
 * * [API documentation](https://cloud.google.com/container-analysis/api/reference/rest/)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/container-analysis/)
 *     * [Creating Attestations (Occurrences)](https://cloud.google.com/binary-authorization/docs/making-attestations)
 *
 * ## Example Usage
 * ### Container Analysis Note Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const note = new gcp.containeranalysis.Note("note", {
 *     attestationAuthority: {
 *         hint: {
 *             humanReadableName: "Attestor Note",
 *         },
 *     },
 * });
 * ```
 * ### Container Analysis Note Attestation Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const note = new gcp.containeranalysis.Note("note", {
 *     attestationAuthority: {
 *         hint: {
 *             humanReadableName: "Attestor Note",
 *         },
 *     },
 *     expirationTime: "2120-10-02T15:01:23.045123456Z",
 *     longDescription: "a longer description of test note",
 *     relatedUrls: [
 *         {
 *             label: "foo",
 *             url: "some.url",
 *         },
 *         {
 *             url: "google.com",
 *         },
 *     ],
 *     shortDescription: "test note",
 * });
 * ```
 *
 * ## Import
 *
 * Note can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:containeranalysis/note:Note default projects/{{project}}/notes/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:containeranalysis/note:Note default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:containeranalysis/note:Note default {{name}}
 * ```
 */
export class Note extends pulumi.CustomResource {
    /**
     * Get an existing Note resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NoteState, opts?: pulumi.CustomResourceOptions): Note {
        return new Note(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:containeranalysis/note:Note';

    /**
     * Returns true if the given object is an instance of Note.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Note {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Note.__pulumiType;
    }

    /**
     * Note kind that represents a logical attestation "role" or "authority".
     * For example, an organization might have one AttestationAuthority for
     * "QA" and one for "build". This Note is intended to act strictly as a
     * grouping mechanism for the attached Occurrences (Attestations). This
     * grouping mechanism also provides a security boundary, since IAM ACLs
     * gate the ability for a principle to attach an Occurrence to a given
     * Note. It also provides a single point of lookup to find all attached
     * Attestation Occurrences, even if they don't all live in the same
     * project.
     * Structure is documented below.
     */
    public readonly attestationAuthority!: pulumi.Output<outputs.containeranalysis.NoteAttestationAuthority>;
    /**
     * The time this note was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Time of expiration for this note. Leave empty if note does not expire.
     */
    public readonly expirationTime!: pulumi.Output<string | undefined>;
    /**
     * The type of analysis this note describes
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * A detailed description of the note
     */
    public readonly longDescription!: pulumi.Output<string | undefined>;
    /**
     * The name of the note.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Names of other notes related to this note.
     */
    public readonly relatedNoteNames!: pulumi.Output<string[] | undefined>;
    /**
     * URLs associated with this note and related metadata.
     * Structure is documented below.
     */
    public readonly relatedUrls!: pulumi.Output<outputs.containeranalysis.NoteRelatedUrl[] | undefined>;
    /**
     * A one sentence description of the note.
     */
    public readonly shortDescription!: pulumi.Output<string | undefined>;
    /**
     * The time this note was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Note resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NoteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NoteArgs | NoteState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NoteState | undefined;
            resourceInputs["attestationAuthority"] = state ? state.attestationAuthority : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["expirationTime"] = state ? state.expirationTime : undefined;
            resourceInputs["kind"] = state ? state.kind : undefined;
            resourceInputs["longDescription"] = state ? state.longDescription : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["relatedNoteNames"] = state ? state.relatedNoteNames : undefined;
            resourceInputs["relatedUrls"] = state ? state.relatedUrls : undefined;
            resourceInputs["shortDescription"] = state ? state.shortDescription : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as NoteArgs | undefined;
            if ((!args || args.attestationAuthority === undefined) && !opts.urn) {
                throw new Error("Missing required property 'attestationAuthority'");
            }
            resourceInputs["attestationAuthority"] = args ? args.attestationAuthority : undefined;
            resourceInputs["expirationTime"] = args ? args.expirationTime : undefined;
            resourceInputs["longDescription"] = args ? args.longDescription : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["relatedNoteNames"] = args ? args.relatedNoteNames : undefined;
            resourceInputs["relatedUrls"] = args ? args.relatedUrls : undefined;
            resourceInputs["shortDescription"] = args ? args.shortDescription : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Note.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Note resources.
 */
export interface NoteState {
    /**
     * Note kind that represents a logical attestation "role" or "authority".
     * For example, an organization might have one AttestationAuthority for
     * "QA" and one for "build". This Note is intended to act strictly as a
     * grouping mechanism for the attached Occurrences (Attestations). This
     * grouping mechanism also provides a security boundary, since IAM ACLs
     * gate the ability for a principle to attach an Occurrence to a given
     * Note. It also provides a single point of lookup to find all attached
     * Attestation Occurrences, even if they don't all live in the same
     * project.
     * Structure is documented below.
     */
    attestationAuthority?: pulumi.Input<inputs.containeranalysis.NoteAttestationAuthority>;
    /**
     * The time this note was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Time of expiration for this note. Leave empty if note does not expire.
     */
    expirationTime?: pulumi.Input<string>;
    /**
     * The type of analysis this note describes
     */
    kind?: pulumi.Input<string>;
    /**
     * A detailed description of the note
     */
    longDescription?: pulumi.Input<string>;
    /**
     * The name of the note.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Names of other notes related to this note.
     */
    relatedNoteNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * URLs associated with this note and related metadata.
     * Structure is documented below.
     */
    relatedUrls?: pulumi.Input<pulumi.Input<inputs.containeranalysis.NoteRelatedUrl>[]>;
    /**
     * A one sentence description of the note.
     */
    shortDescription?: pulumi.Input<string>;
    /**
     * The time this note was last updated.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Note resource.
 */
export interface NoteArgs {
    /**
     * Note kind that represents a logical attestation "role" or "authority".
     * For example, an organization might have one AttestationAuthority for
     * "QA" and one for "build". This Note is intended to act strictly as a
     * grouping mechanism for the attached Occurrences (Attestations). This
     * grouping mechanism also provides a security boundary, since IAM ACLs
     * gate the ability for a principle to attach an Occurrence to a given
     * Note. It also provides a single point of lookup to find all attached
     * Attestation Occurrences, even if they don't all live in the same
     * project.
     * Structure is documented below.
     */
    attestationAuthority: pulumi.Input<inputs.containeranalysis.NoteAttestationAuthority>;
    /**
     * Time of expiration for this note. Leave empty if note does not expire.
     */
    expirationTime?: pulumi.Input<string>;
    /**
     * A detailed description of the note
     */
    longDescription?: pulumi.Input<string>;
    /**
     * The name of the note.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Names of other notes related to this note.
     */
    relatedNoteNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * URLs associated with this note and related metadata.
     * Structure is documented below.
     */
    relatedUrls?: pulumi.Input<pulumi.Input<inputs.containeranalysis.NoteRelatedUrl>[]>;
    /**
     * A one sentence description of the note.
     */
    shortDescription?: pulumi.Input<string>;
}
