// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.deploymentmanager;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.deploymentmanager.DeploymentArgs;
import com.pulumi.gcp.deploymentmanager.inputs.DeploymentState;
import com.pulumi.gcp.deploymentmanager.outputs.DeploymentLabel;
import com.pulumi.gcp.deploymentmanager.outputs.DeploymentTarget;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A collection of resources that are deployed and managed together using
 * a configuration file
 * 
 * &gt; **Warning:** This resource is intended only to manage a Deployment resource,
 * and attempts to manage the Deployment&#39;s resources in the provider as well
 * will likely result in errors or unexpected behavior as the two tools
 * fight over ownership. We strongly discourage doing so unless you are an
 * experienced user of both tools.
 * 
 * In addition, due to limitations of the API, the provider will treat
 * deployments in preview as recreate-only for any update operation other
 * than actually deploying an in-preview deployment (i.e. `preview=true` to
 * `preview=false`).
 * 
 * ## Example Usage
 * 
 * ### Deployment Manager Deployment Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.deploymentmanager.Deployment;
 * import com.pulumi.gcp.deploymentmanager.DeploymentArgs;
 * import com.pulumi.gcp.deploymentmanager.inputs.DeploymentTargetArgs;
 * import com.pulumi.gcp.deploymentmanager.inputs.DeploymentTargetConfigArgs;
 * import com.pulumi.gcp.deploymentmanager.inputs.DeploymentLabelArgs;
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
 *         var deployment = new Deployment("deployment", DeploymentArgs.builder()
 *             .name("my-deployment")
 *             .target(DeploymentTargetArgs.builder()
 *                 .config(DeploymentTargetConfigArgs.builder()
 *                     .content(StdFunctions.file(FileArgs.builder()
 *                         .input("path/to/config.yml")
 *                         .build()).result())
 *                     .build())
 *                 .build())
 *             .labels(DeploymentLabelArgs.builder()
 *                 .key("foo")
 *                 .value("bar")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ## Import
 * 
 * Deployment can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/deployments/{{name}}`
 * 
 * * `{{project}}/{{name}}`
 * 
 * * `{{name}}`
 * 
 * When using the `pulumi import` command, Deployment can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:deploymentmanager/deployment:Deployment default projects/{{project}}/deployments/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:deploymentmanager/deployment:Deployment default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:deploymentmanager/deployment:Deployment default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:deploymentmanager/deployment:Deployment")
public class Deployment extends com.pulumi.resources.CustomResource {
    /**
     * Set the policy to use for creating new resources. Only used on create and update. Valid values are &#39;CREATE_OR_ACQUIRE&#39;
     * (default) or &#39;ACQUIRE&#39;. If set to &#39;ACQUIRE&#39; and resources do not already exist, the deployment will fail. Note that
     * updating this field does not actually affect the deployment, just how it is updated. Default value: &#34;CREATE_OR_ACQUIRE&#34;
     * Possible values: [&#34;ACQUIRE&#34;, &#34;CREATE_OR_ACQUIRE&#34;]
     * 
     */
    @Export(name="createPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> createPolicy;

    /**
     * @return Set the policy to use for creating new resources. Only used on create and update. Valid values are &#39;CREATE_OR_ACQUIRE&#39;
     * (default) or &#39;ACQUIRE&#39;. If set to &#39;ACQUIRE&#39; and resources do not already exist, the deployment will fail. Note that
     * updating this field does not actually affect the deployment, just how it is updated. Default value: &#34;CREATE_OR_ACQUIRE&#34;
     * Possible values: [&#34;ACQUIRE&#34;, &#34;CREATE_OR_ACQUIRE&#34;]
     * 
     */
    public Output<Optional<String>> createPolicy() {
        return Codegen.optional(this.createPolicy);
    }
    /**
     * Set the policy to use for deleting new resources on update/delete. Valid values are &#39;DELETE&#39; (default) or &#39;ABANDON&#39;. If
     * &#39;DELETE&#39;, resource is deleted after removal from Deployment Manager. If &#39;ABANDON&#39;, the resource is only removed from
     * Deployment Manager and is not actually deleted. Note that updating this field does not actually change the deployment,
     * just how it is updated. Default value: &#34;DELETE&#34; Possible values: [&#34;ABANDON&#34;, &#34;DELETE&#34;]
     * 
     */
    @Export(name="deletePolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> deletePolicy;

    /**
     * @return Set the policy to use for deleting new resources on update/delete. Valid values are &#39;DELETE&#39; (default) or &#39;ABANDON&#39;. If
     * &#39;DELETE&#39;, resource is deleted after removal from Deployment Manager. If &#39;ABANDON&#39;, the resource is only removed from
     * Deployment Manager and is not actually deleted. Note that updating this field does not actually change the deployment,
     * just how it is updated. Default value: &#34;DELETE&#34; Possible values: [&#34;ABANDON&#34;, &#34;DELETE&#34;]
     * 
     */
    public Output<Optional<String>> deletePolicy() {
        return Codegen.optional(this.deletePolicy);
    }
    /**
     * Unique identifier for deployment. Output only.
     * 
     */
    @Export(name="deploymentId", refs={String.class}, tree="[0]")
    private Output<String> deploymentId;

    /**
     * @return Unique identifier for deployment. Output only.
     * 
     */
    public Output<String> deploymentId() {
        return this.deploymentId;
    }
    /**
     * Optional user-provided description of deployment.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Optional user-provided description of deployment.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Key-value pairs to apply to this labels.
     * 
     */
    @Export(name="labels", refs={List.class,DeploymentLabel.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DeploymentLabel>> labels;

    /**
     * @return Key-value pairs to apply to this labels.
     * 
     */
    public Output<Optional<List<DeploymentLabel>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * Output only. URL of the manifest representing the last manifest that
     * was successfully deployed.
     * 
     */
    @Export(name="manifest", refs={String.class}, tree="[0]")
    private Output<String> manifest;

    /**
     * @return Output only. URL of the manifest representing the last manifest that
     * was successfully deployed.
     * 
     */
    public Output<String> manifest() {
        return this.manifest;
    }
    /**
     * Unique name for the deployment
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Unique name for the deployment
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="preview", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> preview;

    public Output<Optional<Boolean>> preview() {
        return Codegen.optional(this.preview);
    }
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    public Output<String> project() {
        return this.project;
    }
    /**
     * Output only. Server defined URL for the resource.
     * 
     */
    @Export(name="selfLink", refs={String.class}, tree="[0]")
    private Output<String> selfLink;

    /**
     * @return Output only. Server defined URL for the resource.
     * 
     */
    public Output<String> selfLink() {
        return this.selfLink;
    }
    /**
     * Parameters that define your deployment, including the deployment
     * configuration and relevant templates.
     * Structure is documented below.
     * 
     */
    @Export(name="target", refs={DeploymentTarget.class}, tree="[0]")
    private Output<DeploymentTarget> target;

    /**
     * @return Parameters that define your deployment, including the deployment
     * configuration and relevant templates.
     * Structure is documented below.
     * 
     */
    public Output<DeploymentTarget> target() {
        return this.target;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Deployment(java.lang.String name) {
        this(name, DeploymentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Deployment(java.lang.String name, DeploymentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Deployment(java.lang.String name, DeploymentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:deploymentmanager/deployment:Deployment", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Deployment(java.lang.String name, Output<java.lang.String> id, @Nullable DeploymentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:deploymentmanager/deployment:Deployment", name, state, makeResourceOptions(options, id), false);
    }

    private static DeploymentArgs makeArgs(DeploymentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DeploymentArgs.Empty : args;
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
    public static Deployment get(java.lang.String name, Output<java.lang.String> id, @Nullable DeploymentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Deployment(name, id, state, options);
    }
}
