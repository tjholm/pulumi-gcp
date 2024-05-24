// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.iap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.iap.BrandArgs;
import com.pulumi.gcp.iap.inputs.BrandState;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * ### Iap Brand
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
 * import com.pulumi.gcp.projects.Service;
 * import com.pulumi.gcp.projects.ServiceArgs;
 * import com.pulumi.gcp.iap.Brand;
 * import com.pulumi.gcp.iap.BrandArgs;
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
 *         var project = new Project("project", ProjectArgs.builder()
 *             .projectId("my-project")
 *             .name("my-project")
 *             .orgId("123456789")
 *             .build());
 * 
 *         var projectService = new Service("projectService", ServiceArgs.builder()
 *             .project(project.projectId())
 *             .service("iap.googleapis.com")
 *             .build());
 * 
 *         var projectBrand = new Brand("projectBrand", BrandArgs.builder()
 *             .supportEmail("support{@literal @}example.com")
 *             .applicationTitle("Cloud IAP protected Application")
 *             .project(projectService.project())
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
 * Brand can be imported using any of these accepted formats:
 * 
 * * `projects/{{project_id}}/brands/{{brand_id}}`
 * 
 * * `projects/{{project_number}}/brands/{{brand_id}}`
 * 
 * * `{{project_number}}/{{brand_id}}`
 * 
 * When using the `pulumi import` command, Brand can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:iap/brand:Brand default projects/{{project_id}}/brands/{{brand_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:iap/brand:Brand default projects/{{project_number}}/brands/{{brand_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:iap/brand:Brand default {{project_number}}/{{brand_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:iap/brand:Brand")
public class Brand extends com.pulumi.resources.CustomResource {
    /**
     * Application name displayed on OAuth consent screen.
     * 
     * ***
     * 
     */
    @Export(name="applicationTitle", refs={String.class}, tree="[0]")
    private Output<String> applicationTitle;

    /**
     * @return Application name displayed on OAuth consent screen.
     * 
     * ***
     * 
     */
    public Output<String> applicationTitle() {
        return this.applicationTitle;
    }
    /**
     * Output only. Identifier of the brand, in the format `projects/{project_number}/brands/{brand_id}`
     * NOTE: The name can also be expressed as `projects/{project_id}/brands/{brand_id}`, e.g. when importing.
     * NOTE: The brand identification corresponds to the project number as only one
     * brand can be created per project.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Output only. Identifier of the brand, in the format `projects/{project_number}/brands/{brand_id}`
     * NOTE: The name can also be expressed as `projects/{project_id}/brands/{brand_id}`, e.g. when importing.
     * NOTE: The brand identification corresponds to the project number as only one
     * brand can be created per project.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Whether the brand is only intended for usage inside the GSuite organization only.
     * 
     */
    @Export(name="orgInternalOnly", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> orgInternalOnly;

    /**
     * @return Whether the brand is only intended for usage inside the GSuite organization only.
     * 
     */
    public Output<Boolean> orgInternalOnly() {
        return this.orgInternalOnly;
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
     * Support email displayed on the OAuth consent screen. Can be either a
     * user or group email. When a user email is specified, the caller must
     * be the user with the associated email address. When a group email is
     * specified, the caller can be either a user or a service account which
     * is an owner of the specified group in Cloud Identity.
     * 
     */
    @Export(name="supportEmail", refs={String.class}, tree="[0]")
    private Output<String> supportEmail;

    /**
     * @return Support email displayed on the OAuth consent screen. Can be either a
     * user or group email. When a user email is specified, the caller must
     * be the user with the associated email address. When a group email is
     * specified, the caller can be either a user or a service account which
     * is an owner of the specified group in Cloud Identity.
     * 
     */
    public Output<String> supportEmail() {
        return this.supportEmail;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Brand(String name) {
        this(name, BrandArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Brand(String name, BrandArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Brand(String name, BrandArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:iap/brand:Brand", name, args == null ? BrandArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Brand(String name, Output<String> id, @Nullable BrandState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:iap/brand:Brand", name, state, makeResourceOptions(options, id));
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
    public static Brand get(String name, Output<String> id, @Nullable BrandState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Brand(name, id, state, options);
    }
}
