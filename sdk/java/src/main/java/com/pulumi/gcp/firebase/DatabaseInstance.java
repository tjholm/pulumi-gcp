// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.firebase;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.firebase.DatabaseInstanceArgs;
import com.pulumi.gcp.firebase.inputs.DatabaseInstanceState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * ### Firebase Database Instance Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.firebase.DatabaseInstance;
 * import com.pulumi.gcp.firebase.DatabaseInstanceArgs;
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
 *         var basic = new DatabaseInstance("basic", DatabaseInstanceArgs.builder()
 *             .project("my-project-name")
 *             .region("us-central1")
 *             .instanceId("active-db")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Firebase Database Instance Full
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.firebase.DatabaseInstance;
 * import com.pulumi.gcp.firebase.DatabaseInstanceArgs;
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
 *         var full = new DatabaseInstance("full", DatabaseInstanceArgs.builder()
 *             .project("my-project-name")
 *             .region("europe-west1")
 *             .instanceId("disabled-db")
 *             .type("USER_DATABASE")
 *             .desiredState("DISABLED")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Firebase Database Instance Default Database
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
 * import com.pulumi.gcp.firebase.Project;
 * import com.pulumi.gcp.firebase.ProjectArgs;
 * import com.pulumi.gcp.projects.Service;
 * import com.pulumi.gcp.projects.ServiceArgs;
 * import com.pulumi.gcp.firebase.DatabaseInstance;
 * import com.pulumi.gcp.firebase.DatabaseInstanceArgs;
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
 *         var default_ = new Project("default", ProjectArgs.builder()
 *             .projectId("rtdb-project")
 *             .name("rtdb-project")
 *             .orgId("123456789")
 *             .labels(Map.of("firebase", "enabled"))
 *             .build());
 * 
 *         var defaultProject = new Project("defaultProject", ProjectArgs.builder()
 *             .project(default_.projectId())
 *             .build());
 * 
 *         var firebaseDatabase = new Service("firebaseDatabase", ServiceArgs.builder()
 *             .project(defaultProject.project())
 *             .service("firebasedatabase.googleapis.com")
 *             .build());
 * 
 *         var defaultDatabaseInstance = new DatabaseInstance("defaultDatabaseInstance", DatabaseInstanceArgs.builder()
 *             .project(defaultProject.project())
 *             .region("us-central1")
 *             .instanceId("rtdb-project-default-rtdb")
 *             .type("DEFAULT_DATABASE")
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
 * Instance can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{region}}/instances/{{instance_id}}`
 * 
 * * `{{project}}/{{region}}/{{instance_id}}`
 * 
 * * `{{region}}/{{instance_id}}`
 * 
 * * `{{instance_id}}`
 * 
 * When using the `pulumi import` command, Instance can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:firebase/databaseInstance:DatabaseInstance default projects/{{project}}/locations/{{region}}/instances/{{instance_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:firebase/databaseInstance:DatabaseInstance default {{project}}/{{region}}/{{instance_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:firebase/databaseInstance:DatabaseInstance default {{region}}/{{instance_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:firebase/databaseInstance:DatabaseInstance default {{instance_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:firebase/databaseInstance:DatabaseInstance")
public class DatabaseInstance extends com.pulumi.resources.CustomResource {
    /**
     * The database URL in the form of https://{instance-id}.firebaseio.com for us-central1 instances
     * or https://{instance-id}.{region}.firebasedatabase.app in other regions.
     * 
     */
    @Export(name="databaseUrl", refs={String.class}, tree="[0]")
    private Output<String> databaseUrl;

    /**
     * @return The database URL in the form of https://{instance-id}.firebaseio.com for us-central1 instances
     * or https://{instance-id}.{region}.firebasedatabase.app in other regions.
     * 
     */
    public Output<String> databaseUrl() {
        return this.databaseUrl;
    }
    /**
     * The intended database state.
     * 
     */
    @Export(name="desiredState", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> desiredState;

    /**
     * @return The intended database state.
     * 
     */
    public Output<Optional<String>> desiredState() {
        return Codegen.optional(this.desiredState);
    }
    /**
     * The globally unique identifier of the Firebase Realtime Database instance.
     * Instance IDs cannot be reused after deletion.
     * 
     * ***
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The globally unique identifier of the Firebase Realtime Database instance.
     * Instance IDs cannot be reused after deletion.
     * 
     * ***
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The fully-qualified resource name of the Firebase Realtime Database, in the
     * format: projects/PROJECT_NUMBER/locations/REGION_IDENTIFIER/instances/INSTANCE_ID
     * PROJECT_NUMBER: The Firebase project&#39;s [`ProjectNumber`](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_number)
     * Learn more about using project identifiers in Google&#39;s [AIP 2510 standard](https://google.aip.dev/cloud/2510).
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The fully-qualified resource name of the Firebase Realtime Database, in the
     * format: projects/PROJECT_NUMBER/locations/REGION_IDENTIFIER/instances/INSTANCE_ID
     * PROJECT_NUMBER: The Firebase project&#39;s [`ProjectNumber`](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_number)
     * Learn more about using project identifiers in Google&#39;s [AIP 2510 standard](https://google.aip.dev/cloud/2510).
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * A reference to the region where the Firebase Realtime database resides.
     * Check all [available regions](https://firebase.google.com/docs/projects/locations#rtdb-locations)
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return A reference to the region where the Firebase Realtime database resides.
     * Check all [available regions](https://firebase.google.com/docs/projects/locations#rtdb-locations)
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The current database state. Set desired_state to :DISABLED to disable the database and :ACTIVE to reenable the database
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return The current database state. Set desired_state to :DISABLED to disable the database and :ACTIVE to reenable the database
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The database type.
     * Each project can create one default Firebase Realtime Database, which cannot be deleted once created.
     * Creating user Databases is only available for projects on the Blaze plan.
     * Projects can be upgraded using the Cloud Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo.
     * Default value is `USER_DATABASE`.
     * Possible values are: `DEFAULT_DATABASE`, `USER_DATABASE`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return The database type.
     * Each project can create one default Firebase Realtime Database, which cannot be deleted once created.
     * Creating user Databases is only available for projects on the Blaze plan.
     * Projects can be upgraded using the Cloud Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo.
     * Default value is `USER_DATABASE`.
     * Possible values are: `DEFAULT_DATABASE`, `USER_DATABASE`.
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DatabaseInstance(String name) {
        this(name, DatabaseInstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DatabaseInstance(String name, DatabaseInstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DatabaseInstance(String name, DatabaseInstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:firebase/databaseInstance:DatabaseInstance", name, args == null ? DatabaseInstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DatabaseInstance(String name, Output<String> id, @Nullable DatabaseInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:firebase/databaseInstance:DatabaseInstance", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static DatabaseInstance get(String name, Output<String> id, @Nullable DatabaseInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DatabaseInstance(name, id, state, options);
    }
}
