// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieves the current IAM policy data for tagtemplate
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.datacatalog.getTagTemplateIamPolicy({
 *     tagTemplate: google_data_catalog_tag_template.basic_tag_template.name,
 * });
 * ```
 */
export function getTagTemplateIamPolicy(args: GetTagTemplateIamPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetTagTemplateIamPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:datacatalog/getTagTemplateIamPolicy:getTagTemplateIamPolicy", {
        "project": args.project,
        "region": args.region,
        "tagTemplate": args.tagTemplate,
    }, opts);
}

/**
 * A collection of arguments for invoking getTagTemplateIamPolicy.
 */
export interface GetTagTemplateIamPolicyArgs {
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: string;
    region?: string;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    tagTemplate: string;
}

/**
 * A collection of values returned by getTagTemplateIamPolicy.
 */
export interface GetTagTemplateIamPolicyResult {
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (Required only by `gcp.datacatalog.TagTemplateIamPolicy`) The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: string;
    readonly project: string;
    readonly region: string;
    readonly tagTemplate: string;
}
/**
 * Retrieves the current IAM policy data for tagtemplate
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.datacatalog.getTagTemplateIamPolicy({
 *     tagTemplate: google_data_catalog_tag_template.basic_tag_template.name,
 * });
 * ```
 */
export function getTagTemplateIamPolicyOutput(args: GetTagTemplateIamPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTagTemplateIamPolicyResult> {
    return pulumi.output(args).apply((a: any) => getTagTemplateIamPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getTagTemplateIamPolicy.
 */
export interface GetTagTemplateIamPolicyOutputArgs {
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    tagTemplate: pulumi.Input<string>;
}
