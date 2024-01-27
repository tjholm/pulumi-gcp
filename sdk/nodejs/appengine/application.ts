// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Allows creation and management of an App Engine application.
 *
 * > App Engine applications cannot be deleted once they're created; you have to delete the
 *    entire project to delete the application. This provider will report the application has been
 *    successfully deleted; this is a limitation of the provider, and will go away in the future.
 *    This provider is not able to delete App Engine applications.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const myProject = new gcp.organizations.Project("myProject", {orgId: "1234567"});
 * const app = new gcp.appengine.Application("app", {
 *     project: myProject.projectId,
 *     locationId: "us-central",
 * });
 * ```
 *
 * ## Import
 *
 * Applications can be imported using the ID of the project the application belongs to, e.g. * `{{project-id}}` When using the `pulumi import` command, Applications can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:appengine/application:Application default {{project-id}}
 * ```
 */
export class Application extends pulumi.CustomResource {
    /**
     * Get an existing Application resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationState, opts?: pulumi.CustomResourceOptions): Application {
        return new Application(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:appengine/application:Application';

    /**
     * Returns true if the given object is an instance of Application.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Application {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Application.__pulumiType;
    }

    /**
     * Identifier of the app, usually `{PROJECT_ID}`
     */
    public /*out*/ readonly appId!: pulumi.Output<string>;
    /**
     * The domain to authenticate users with when using App Engine's User API.
     */
    public readonly authDomain!: pulumi.Output<string>;
    /**
     * The GCS bucket code is being stored in for this app.
     */
    public /*out*/ readonly codeBucket!: pulumi.Output<string>;
    /**
     * The type of the Cloud Firestore or Cloud Datastore database associated with this application.
     * Can be `CLOUD_FIRESTORE` or `CLOUD_DATASTORE_COMPATIBILITY` for new
     * instances.  To support old instances, the value `CLOUD_DATASTORE` is accepted by the provider, but will be rejected by the API.
     * To create a Cloud Firestore database without creating an App Engine application, use the
     * `gcp.firestore.Database`
     * resource instead.
     */
    public readonly databaseType!: pulumi.Output<string>;
    /**
     * The GCS bucket content is being stored in for this app.
     */
    public /*out*/ readonly defaultBucket!: pulumi.Output<string>;
    /**
     * The default hostname for this app.
     */
    public /*out*/ readonly defaultHostname!: pulumi.Output<string>;
    /**
     * A block of optional settings to configure specific App Engine features:
     */
    public readonly featureSettings!: pulumi.Output<outputs.appengine.ApplicationFeatureSettings>;
    /**
     * The GCR domain used for storing managed Docker images for this app.
     */
    public /*out*/ readonly gcrDomain!: pulumi.Output<string>;
    /**
     * Settings for enabling Cloud Identity Aware Proxy
     */
    public readonly iap!: pulumi.Output<outputs.appengine.ApplicationIap>;
    /**
     * The [location](https://cloud.google.com/appengine/docs/locations)
     * to serve the app from.
     */
    public readonly locationId!: pulumi.Output<string>;
    /**
     * Unique name of the app, usually `apps/{PROJECT_ID}`
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The project ID to create the application under.
     * ~>**NOTE:** GCP only accepts project ID, not project number. If you are using number,
     * you may get a "Permission denied" error.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The serving status of the app.
     */
    public readonly servingStatus!: pulumi.Output<string>;
    /**
     * A list of dispatch rule blocks. Each block has a `domain`, `path`, and `service` field.
     */
    public /*out*/ readonly urlDispatchRules!: pulumi.Output<outputs.appengine.ApplicationUrlDispatchRule[]>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationArgs | ApplicationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationState | undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["authDomain"] = state ? state.authDomain : undefined;
            resourceInputs["codeBucket"] = state ? state.codeBucket : undefined;
            resourceInputs["databaseType"] = state ? state.databaseType : undefined;
            resourceInputs["defaultBucket"] = state ? state.defaultBucket : undefined;
            resourceInputs["defaultHostname"] = state ? state.defaultHostname : undefined;
            resourceInputs["featureSettings"] = state ? state.featureSettings : undefined;
            resourceInputs["gcrDomain"] = state ? state.gcrDomain : undefined;
            resourceInputs["iap"] = state ? state.iap : undefined;
            resourceInputs["locationId"] = state ? state.locationId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["servingStatus"] = state ? state.servingStatus : undefined;
            resourceInputs["urlDispatchRules"] = state ? state.urlDispatchRules : undefined;
        } else {
            const args = argsOrState as ApplicationArgs | undefined;
            if ((!args || args.locationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'locationId'");
            }
            resourceInputs["authDomain"] = args ? args.authDomain : undefined;
            resourceInputs["databaseType"] = args ? args.databaseType : undefined;
            resourceInputs["featureSettings"] = args ? args.featureSettings : undefined;
            resourceInputs["iap"] = args ? args.iap : undefined;
            resourceInputs["locationId"] = args ? args.locationId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["servingStatus"] = args ? args.servingStatus : undefined;
            resourceInputs["appId"] = undefined /*out*/;
            resourceInputs["codeBucket"] = undefined /*out*/;
            resourceInputs["defaultBucket"] = undefined /*out*/;
            resourceInputs["defaultHostname"] = undefined /*out*/;
            resourceInputs["gcrDomain"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["urlDispatchRules"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Application.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Application resources.
 */
export interface ApplicationState {
    /**
     * Identifier of the app, usually `{PROJECT_ID}`
     */
    appId?: pulumi.Input<string>;
    /**
     * The domain to authenticate users with when using App Engine's User API.
     */
    authDomain?: pulumi.Input<string>;
    /**
     * The GCS bucket code is being stored in for this app.
     */
    codeBucket?: pulumi.Input<string>;
    /**
     * The type of the Cloud Firestore or Cloud Datastore database associated with this application.
     * Can be `CLOUD_FIRESTORE` or `CLOUD_DATASTORE_COMPATIBILITY` for new
     * instances.  To support old instances, the value `CLOUD_DATASTORE` is accepted by the provider, but will be rejected by the API.
     * To create a Cloud Firestore database without creating an App Engine application, use the
     * `gcp.firestore.Database`
     * resource instead.
     */
    databaseType?: pulumi.Input<string>;
    /**
     * The GCS bucket content is being stored in for this app.
     */
    defaultBucket?: pulumi.Input<string>;
    /**
     * The default hostname for this app.
     */
    defaultHostname?: pulumi.Input<string>;
    /**
     * A block of optional settings to configure specific App Engine features:
     */
    featureSettings?: pulumi.Input<inputs.appengine.ApplicationFeatureSettings>;
    /**
     * The GCR domain used for storing managed Docker images for this app.
     */
    gcrDomain?: pulumi.Input<string>;
    /**
     * Settings for enabling Cloud Identity Aware Proxy
     */
    iap?: pulumi.Input<inputs.appengine.ApplicationIap>;
    /**
     * The [location](https://cloud.google.com/appengine/docs/locations)
     * to serve the app from.
     */
    locationId?: pulumi.Input<string>;
    /**
     * Unique name of the app, usually `apps/{PROJECT_ID}`
     */
    name?: pulumi.Input<string>;
    /**
     * The project ID to create the application under.
     * ~>**NOTE:** GCP only accepts project ID, not project number. If you are using number,
     * you may get a "Permission denied" error.
     */
    project?: pulumi.Input<string>;
    /**
     * The serving status of the app.
     */
    servingStatus?: pulumi.Input<string>;
    /**
     * A list of dispatch rule blocks. Each block has a `domain`, `path`, and `service` field.
     */
    urlDispatchRules?: pulumi.Input<pulumi.Input<inputs.appengine.ApplicationUrlDispatchRule>[]>;
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    /**
     * The domain to authenticate users with when using App Engine's User API.
     */
    authDomain?: pulumi.Input<string>;
    /**
     * The type of the Cloud Firestore or Cloud Datastore database associated with this application.
     * Can be `CLOUD_FIRESTORE` or `CLOUD_DATASTORE_COMPATIBILITY` for new
     * instances.  To support old instances, the value `CLOUD_DATASTORE` is accepted by the provider, but will be rejected by the API.
     * To create a Cloud Firestore database without creating an App Engine application, use the
     * `gcp.firestore.Database`
     * resource instead.
     */
    databaseType?: pulumi.Input<string>;
    /**
     * A block of optional settings to configure specific App Engine features:
     */
    featureSettings?: pulumi.Input<inputs.appengine.ApplicationFeatureSettings>;
    /**
     * Settings for enabling Cloud Identity Aware Proxy
     */
    iap?: pulumi.Input<inputs.appengine.ApplicationIap>;
    /**
     * The [location](https://cloud.google.com/appengine/docs/locations)
     * to serve the app from.
     */
    locationId: pulumi.Input<string>;
    /**
     * The project ID to create the application under.
     * ~>**NOTE:** GCP only accepts project ID, not project number. If you are using number,
     * you may get a "Permission denied" error.
     */
    project?: pulumi.Input<string>;
    /**
     * The serving status of the app.
     */
    servingStatus?: pulumi.Input<string>;
}
