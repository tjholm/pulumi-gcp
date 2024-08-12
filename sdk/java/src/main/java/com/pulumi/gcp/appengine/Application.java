// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.appengine;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.appengine.ApplicationArgs;
import com.pulumi.gcp.appengine.inputs.ApplicationState;
import com.pulumi.gcp.appengine.outputs.ApplicationFeatureSettings;
import com.pulumi.gcp.appengine.outputs.ApplicationIap;
import com.pulumi.gcp.appengine.outputs.ApplicationUrlDispatchRule;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Allows creation and management of an App Engine application.
 * 
 * &gt; App Engine applications cannot be deleted once they&#39;re created; you have to delete the
 *    entire project to delete the application. This provider will report the application has been
 *    successfully deleted; this is a limitation of the provider, and will go away in the future.
 *    This provider is not able to delete App Engine applications.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.Project;
 * import com.pulumi.gcp.organizations.ProjectArgs;
 * import com.pulumi.gcp.appengine.Application;
 * import com.pulumi.gcp.appengine.ApplicationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var myProject = new Project("myProject", ProjectArgs.builder()
 *             .name("My Project")
 *             .projectId("your-project-id")
 *             .orgId("1234567")
 *             .build());
 * 
 *         var app = new Application("app", ApplicationArgs.builder()
 *             .project(myProject.projectId())
 *             .locationId("us-central")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Applications can be imported using the ID of the project the application belongs to, e.g.
 * 
 * * `{{project-id}}`
 * 
 * When using the `pulumi import` command, Applications can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:appengine/application:Application default {{project-id}}
 * ```
 * 
 */
@ResourceType(type="gcp:appengine/application:Application")
public class Application extends com.pulumi.resources.CustomResource {
    /**
     * Identifier of the app, usually `{PROJECT_ID}`
     * 
     */
    @Export(name="appId", refs={String.class}, tree="[0]")
    private Output<String> appId;

    /**
     * @return Identifier of the app, usually `{PROJECT_ID}`
     * 
     */
    public Output<String> appId() {
        return this.appId;
    }
    /**
     * The domain to authenticate users with when using App Engine&#39;s User API.
     * 
     */
    @Export(name="authDomain", refs={String.class}, tree="[0]")
    private Output<String> authDomain;

    /**
     * @return The domain to authenticate users with when using App Engine&#39;s User API.
     * 
     */
    public Output<String> authDomain() {
        return this.authDomain;
    }
    /**
     * The GCS bucket code is being stored in for this app.
     * 
     */
    @Export(name="codeBucket", refs={String.class}, tree="[0]")
    private Output<String> codeBucket;

    /**
     * @return The GCS bucket code is being stored in for this app.
     * 
     */
    public Output<String> codeBucket() {
        return this.codeBucket;
    }
    /**
     * The type of the Cloud Firestore or Cloud Datastore database associated with this application.
     * Can be `CLOUD_FIRESTORE` or `CLOUD_DATASTORE_COMPATIBILITY` for new
     * instances.  To support old instances, the value `CLOUD_DATASTORE` is accepted by the provider, but will be rejected by the API.
     * To create a Cloud Firestore database without creating an App Engine application, use the
     * `gcp.firestore.Database`
     * resource instead.
     * 
     */
    @Export(name="databaseType", refs={String.class}, tree="[0]")
    private Output<String> databaseType;

    /**
     * @return The type of the Cloud Firestore or Cloud Datastore database associated with this application.
     * Can be `CLOUD_FIRESTORE` or `CLOUD_DATASTORE_COMPATIBILITY` for new
     * instances.  To support old instances, the value `CLOUD_DATASTORE` is accepted by the provider, but will be rejected by the API.
     * To create a Cloud Firestore database without creating an App Engine application, use the
     * `gcp.firestore.Database`
     * resource instead.
     * 
     */
    public Output<String> databaseType() {
        return this.databaseType;
    }
    /**
     * The GCS bucket content is being stored in for this app.
     * 
     */
    @Export(name="defaultBucket", refs={String.class}, tree="[0]")
    private Output<String> defaultBucket;

    /**
     * @return The GCS bucket content is being stored in for this app.
     * 
     */
    public Output<String> defaultBucket() {
        return this.defaultBucket;
    }
    /**
     * The default hostname for this app.
     * 
     */
    @Export(name="defaultHostname", refs={String.class}, tree="[0]")
    private Output<String> defaultHostname;

    /**
     * @return The default hostname for this app.
     * 
     */
    public Output<String> defaultHostname() {
        return this.defaultHostname;
    }
    /**
     * A block of optional settings to configure specific App Engine features:
     * 
     */
    @Export(name="featureSettings", refs={ApplicationFeatureSettings.class}, tree="[0]")
    private Output<ApplicationFeatureSettings> featureSettings;

    /**
     * @return A block of optional settings to configure specific App Engine features:
     * 
     */
    public Output<ApplicationFeatureSettings> featureSettings() {
        return this.featureSettings;
    }
    /**
     * The GCR domain used for storing managed Docker images for this app.
     * 
     */
    @Export(name="gcrDomain", refs={String.class}, tree="[0]")
    private Output<String> gcrDomain;

    /**
     * @return The GCR domain used for storing managed Docker images for this app.
     * 
     */
    public Output<String> gcrDomain() {
        return this.gcrDomain;
    }
    /**
     * Settings for enabling Cloud Identity Aware Proxy
     * 
     */
    @Export(name="iap", refs={ApplicationIap.class}, tree="[0]")
    private Output<ApplicationIap> iap;

    /**
     * @return Settings for enabling Cloud Identity Aware Proxy
     * 
     */
    public Output<ApplicationIap> iap() {
        return this.iap;
    }
    /**
     * The [location](https://cloud.google.com/appengine/docs/locations)
     * to serve the app from.
     * 
     */
    @Export(name="locationId", refs={String.class}, tree="[0]")
    private Output<String> locationId;

    /**
     * @return The [location](https://cloud.google.com/appengine/docs/locations)
     * to serve the app from.
     * 
     */
    public Output<String> locationId() {
        return this.locationId;
    }
    /**
     * Unique name of the app, usually `apps/{PROJECT_ID}`
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Unique name of the app, usually `apps/{PROJECT_ID}`
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The project ID to create the application under.
     * ~&gt;**NOTE:** GCP only accepts project ID, not project number. If you are using number,
     * you may get a &#34;Permission denied&#34; error.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The project ID to create the application under.
     * ~&gt;**NOTE:** GCP only accepts project ID, not project number. If you are using number,
     * you may get a &#34;Permission denied&#34; error.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The serving status of the app.
     * 
     */
    @Export(name="servingStatus", refs={String.class}, tree="[0]")
    private Output<String> servingStatus;

    /**
     * @return The serving status of the app.
     * 
     */
    public Output<String> servingStatus() {
        return this.servingStatus;
    }
    /**
     * A list of dispatch rule blocks. Each block has a `domain`, `path`, and `service` field.
     * 
     */
    @Export(name="urlDispatchRules", refs={List.class,ApplicationUrlDispatchRule.class}, tree="[0,1]")
    private Output<List<ApplicationUrlDispatchRule>> urlDispatchRules;

    /**
     * @return A list of dispatch rule blocks. Each block has a `domain`, `path`, and `service` field.
     * 
     */
    public Output<List<ApplicationUrlDispatchRule>> urlDispatchRules() {
        return this.urlDispatchRules;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Application(java.lang.String name) {
        this(name, ApplicationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Application(java.lang.String name, ApplicationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Application(java.lang.String name, ApplicationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:appengine/application:Application", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Application(java.lang.String name, Output<java.lang.String> id, @Nullable ApplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:appengine/application:Application", name, state, makeResourceOptions(options, id), false);
    }

    private static ApplicationArgs makeArgs(ApplicationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ApplicationArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Application get(java.lang.String name, Output<java.lang.String> id, @Nullable ApplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Application(name, id, state, options);
    }
}
