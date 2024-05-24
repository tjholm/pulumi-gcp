// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.billing;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.billing.ProjectInfoArgs;
import com.pulumi.gcp.billing.inputs.ProjectInfoState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Billing information for a project.
 * 
 * To get more information about ProjectInfo, see:
 * 
 * * [API documentation](https://cloud.google.com/billing/docs/reference/rest/v1/projects)
 * * How-to Guides
 *     * [Enable, disable, or change billing for a project](https://cloud.google.com/billing/docs/how-to/modify-project)
 * 
 * ## Example Usage
 * 
 * ### Billing Project Info Basic
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
 * import com.pulumi.gcp.billing.ProjectInfo;
 * import com.pulumi.gcp.billing.ProjectInfoArgs;
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
 *             .projectId("tf-test_15222")
 *             .name("tf-test_81126")
 *             .orgId("123456789")
 *             .build());
 * 
 *         var default_ = new ProjectInfo("default", ProjectInfoArgs.builder()
 *             .project(project.projectId())
 *             .billingAccount("000000-0000000-0000000-000000")
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
 * ProjectInfo can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}`
 * 
 * * `{{project}}`
 * 
 * When using the `pulumi import` command, ProjectInfo can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:billing/projectInfo:ProjectInfo default projects/{{project}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:billing/projectInfo:ProjectInfo default {{project}}
 * ```
 * 
 */
@ResourceType(type="gcp:billing/projectInfo:ProjectInfo")
public class ProjectInfo extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the billing account associated with the project, if
     * any. Set to empty string to disable billing for the project.
     * For example, `&#34;012345-567890-ABCDEF&#34;` or `&#34;&#34;`.
     * 
     * ***
     * 
     */
    @Export(name="billingAccount", refs={String.class}, tree="[0]")
    private Output<String> billingAccount;

    /**
     * @return The ID of the billing account associated with the project, if
     * any. Set to empty string to disable billing for the project.
     * For example, `&#34;012345-567890-ABCDEF&#34;` or `&#34;&#34;`.
     * 
     * ***
     * 
     */
    public Output<String> billingAccount() {
        return this.billingAccount;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectInfo(String name) {
        this(name, ProjectInfoArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectInfo(String name, ProjectInfoArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectInfo(String name, ProjectInfoArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:billing/projectInfo:ProjectInfo", name, args == null ? ProjectInfoArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProjectInfo(String name, Output<String> id, @Nullable ProjectInfoState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:billing/projectInfo:ProjectInfo", name, state, makeResourceOptions(options, id));
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
    public static ProjectInfo get(String name, Output<String> id, @Nullable ProjectInfoState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectInfo(name, id, state, options);
    }
}
