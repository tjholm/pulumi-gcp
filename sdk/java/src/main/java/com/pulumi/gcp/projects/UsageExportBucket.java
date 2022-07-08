// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.projects;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.projects.UsageExportBucketArgs;
import com.pulumi.gcp.projects.inputs.UsageExportBucketState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows creation and management of a Google Cloud Platform project.
 * 
 * Projects created with this resource must be associated with an Organization.
 * See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.
 * 
 * The user or service account that is running this provider when creating a `gcp.organizations.Project`
 * resource must have `roles/resourcemanager.projectCreator` on the specified organization. See the
 * [Access Control for Organizations Using IAM](https://cloud.google.com/resource-manager/docs/access-control-org)
 * doc for more information.
 * 
 * &gt; This resource reads the specified billing account on every provider apply and plan operation so you must have permissions on the specified billing account.
 * 
 * To get more information about projects, see:
 * 
 * * [API documentation](https://cloud.google.com/resource-manager/reference/rest/v1/projects)
 * * How-to Guides
 *     * [Creating and managing projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var myProject = new Project(&#34;myProject&#34;, ProjectArgs.builder()        
 *             .orgId(&#34;1234567&#34;)
 *             .projectId(&#34;your-project-id&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * To create a project under a specific folder
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var department1 = new Folder(&#34;department1&#34;, FolderArgs.builder()        
 *             .displayName(&#34;Department 1&#34;)
 *             .parent(&#34;organizations/1234567&#34;)
 *             .build());
 * 
 *         var myProject_in_a_folder = new Project(&#34;myProject-in-a-folder&#34;, ProjectArgs.builder()        
 *             .projectId(&#34;your-project-id&#34;)
 *             .folderId(department1.name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Projects can be imported using the `project_id`, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:projects/usageExportBucket:UsageExportBucket my_project your-project-id
 * ```
 * 
 */
@ResourceType(type="gcp:projects/usageExportBucket:UsageExportBucket")
public class UsageExportBucket extends com.pulumi.resources.CustomResource {
    /**
     * The bucket to store reports in.
     * 
     */
    @Export(name="bucketName", type=String.class, parameters={})
    private Output<String> bucketName;

    /**
     * @return The bucket to store reports in.
     * 
     */
    public Output<String> bucketName() {
        return this.bucketName;
    }
    /**
     * A prefix for the reports, for instance, the project name.
     * 
     */
    @Export(name="prefix", type=String.class, parameters={})
    private Output</* @Nullable */ String> prefix;

    /**
     * @return A prefix for the reports, for instance, the project name.
     * 
     */
    public Output<Optional<String>> prefix() {
        return Codegen.optional(this.prefix);
    }
    /**
     * The project to set the export bucket on. If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
    private Output<String> project;

    /**
     * @return The project to set the export bucket on. If it is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UsageExportBucket(String name) {
        this(name, UsageExportBucketArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UsageExportBucket(String name, UsageExportBucketArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UsageExportBucket(String name, UsageExportBucketArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:projects/usageExportBucket:UsageExportBucket", name, args == null ? UsageExportBucketArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private UsageExportBucket(String name, Output<String> id, @Nullable UsageExportBucketState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:projects/usageExportBucket:UsageExportBucket", name, state, makeResourceOptions(options, id));
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
    public static UsageExportBucket get(String name, Output<String> id, @Nullable UsageExportBucketState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UsageExportBucket(name, id, state, options);
    }
}
