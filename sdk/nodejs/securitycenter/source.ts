// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A Cloud Security Command Center's (Cloud SCC) finding source. A finding
 * source is an entity or a mechanism that can produce a finding. A source is
 * like a container of findings that come from the same scanner, logger,
 * monitor, etc.
 *
 * To get more information about Source, see:
 *
 * * [API documentation](https://cloud.google.com/security-command-center/docs/reference/rest/v1/organizations.sources)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/security-command-center/docs)
 *
 * ## Example Usage
 * ### Scc Source Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const customSource = new gcp.securitycenter.Source("custom_source", {
 *     description: "My custom Cloud Security Command Center Finding Source",
 *     displayName: "My Source",
 *     organization: "123456789",
 * });
 * ```
 *
 * ## Import
 *
 * Source can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:securitycenter/source:Source default organizations/{{organization}}/sources/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:securitycenter/source:Source default {{organization}}/{{name}}
 * ```
 */
export class Source extends pulumi.CustomResource {
    /**
     * Get an existing Source resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SourceState, opts?: pulumi.CustomResourceOptions): Source {
        return new Source(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:securitycenter/source:Source';

    /**
     * Returns true if the given object is an instance of Source.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Source {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Source.__pulumiType;
    }

    /**
     * The description of the source (max of 1024 characters).
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The source’s display name. A source’s display name must be unique
     * amongst its siblings, for example, two sources with the same parent
     * can't share the same display name. The display name must start and end
     * with a letter or digit, may contain letters, digits, spaces, hyphens,
     * and underscores, and can be no longer than 32 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The resource name of this source, in the format 'organizations/{{organization}}/sources/{{source}}'.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The organization whose Cloud Security Command Center the Source
     * lives in.
     */
    public readonly organization!: pulumi.Output<string>;

    /**
     * Create a Source resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SourceArgs | SourceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SourceState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organization"] = state ? state.organization : undefined;
        } else {
            const args = argsOrState as SourceArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.organization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organization'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Source.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Source resources.
 */
export interface SourceState {
    /**
     * The description of the source (max of 1024 characters).
     */
    description?: pulumi.Input<string>;
    /**
     * The source’s display name. A source’s display name must be unique
     * amongst its siblings, for example, two sources with the same parent
     * can't share the same display name. The display name must start and end
     * with a letter or digit, may contain letters, digits, spaces, hyphens,
     * and underscores, and can be no longer than 32 characters.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The resource name of this source, in the format 'organizations/{{organization}}/sources/{{source}}'.
     */
    name?: pulumi.Input<string>;
    /**
     * The organization whose Cloud Security Command Center the Source
     * lives in.
     */
    organization?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Source resource.
 */
export interface SourceArgs {
    /**
     * The description of the source (max of 1024 characters).
     */
    description?: pulumi.Input<string>;
    /**
     * The source’s display name. A source’s display name must be unique
     * amongst its siblings, for example, two sources with the same parent
     * can't share the same display name. The display name must start and end
     * with a letter or digit, may contain letters, digits, spaces, hyphens,
     * and underscores, and can be no longer than 32 characters.
     */
    displayName: pulumi.Input<string>;
    /**
     * The organization whose Cloud Security Command Center the Source
     * lives in.
     */
    organization: pulumi.Input<string>;
}
