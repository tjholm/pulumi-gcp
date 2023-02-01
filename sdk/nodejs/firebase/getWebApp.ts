// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getWebApp(args: GetWebAppArgs, opts?: pulumi.InvokeOptions): Promise<GetWebAppResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:firebase/getWebApp:getWebApp", {
        "appId": args.appId,
    }, opts);
}

/**
 * A collection of arguments for invoking getWebApp.
 */
export interface GetWebAppArgs {
    /**
     * The appIp of name of the Firebase webApp.
     */
    appId: string;
}

/**
 * A collection of values returned by getWebApp.
 */
export interface GetWebAppResult {
    /**
     * Immutable. The globally unique, Firebase-assigned identifier of the App.
     * This identifier should be treated as an opaque token, as the data format is not specified.
     */
    readonly appId: string;
    readonly appUrls: string[];
    readonly deletionPolicy: string;
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The fully qualified resource name of the App, for example:
     * projects/projectId/webApps/appId
     */
    readonly name: string;
    readonly project: string;
}
export function getWebAppOutput(args: GetWebAppOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWebAppResult> {
    return pulumi.output(args).apply((a: any) => getWebApp(a, opts))
}

/**
 * A collection of arguments for invoking getWebApp.
 */
export interface GetWebAppOutputArgs {
    /**
     * The appIp of name of the Firebase webApp.
     */
    appId: pulumi.Input<string>;
}
