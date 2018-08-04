// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Get an active folder within GCP by `display_name` and `parent`.
 */
export function getActiveFolder(args: GetActiveFolderArgs): Promise<GetActiveFolderResult> {
    return pulumi.runtime.invoke("gcp:organizations/getActiveFolder:getActiveFolder", {
        "displayName": args.displayName,
        "parent": args.parent,
    });
}

/**
 * A collection of arguments for invoking getActiveFolder.
 */
export interface GetActiveFolderArgs {
    /**
     * The folder's display name.
     */
    readonly displayName: string;
    /**
     * The resource name of the parent Folder or Organization.
     */
    readonly parent: string;
}

/**
 * A collection of values returned by getActiveFolder.
 */
export interface GetActiveFolderResult {
    /**
     * The resource name of the Folder. This uniquely identifies the folder.
     */
    readonly name: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
