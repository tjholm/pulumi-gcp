// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.folder;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.folder.OrganizationPolicyArgs;
import com.pulumi.gcp.folder.inputs.OrganizationPolicyState;
import com.pulumi.gcp.folder.outputs.OrganizationPolicyBooleanPolicy;
import com.pulumi.gcp.folder.outputs.OrganizationPolicyListPolicy;
import com.pulumi.gcp.folder.outputs.OrganizationPolicyRestorePolicy;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows management of Organization Policies for a Google Cloud Folder.
 * 
 * &gt; **Warning:** This resource has been superseded by `gcp.orgpolicy.Policy`. `gcp.orgpolicy.Policy` uses Organization Policy API V2 instead of Cloud Resource Manager API V1 and it supports additional features such as tags and conditions.
 * 
 * To get more information about Organization Policies, see:
 * 
 * * [API documentation](https://cloud.google.com/resource-manager/reference/rest/v1/folders/setOrgPolicy)
 * * How-to Guides
 *     * [Introduction to the Organization Policy Service](https://cloud.google.com/resource-manager/docs/organization-policy/overview)
 * 
 * ## Example Usage
 * 
 * To set policy with a [boolean constraint](https://cloud.google.com/resource-manager/docs/organization-policy/quickstart-boolean-constraints):
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.folder.OrganizationPolicy;
 * import com.pulumi.gcp.folder.OrganizationPolicyArgs;
 * import com.pulumi.gcp.folder.inputs.OrganizationPolicyBooleanPolicyArgs;
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
 *         var serialPortPolicy = new OrganizationPolicy("serialPortPolicy", OrganizationPolicyArgs.builder()
 *             .folder("folders/123456789")
 *             .constraint("compute.disableSerialPortAccess")
 *             .booleanPolicy(OrganizationPolicyBooleanPolicyArgs.builder()
 *                 .enforced(true)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * To set a policy with a [list constraint](https://cloud.google.com/resource-manager/docs/organization-policy/quickstart-list-constraints):
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.folder.OrganizationPolicy;
 * import com.pulumi.gcp.folder.OrganizationPolicyArgs;
 * import com.pulumi.gcp.folder.inputs.OrganizationPolicyListPolicyArgs;
 * import com.pulumi.gcp.folder.inputs.OrganizationPolicyListPolicyAllowArgs;
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
 *         var servicesPolicy = new OrganizationPolicy("servicesPolicy", OrganizationPolicyArgs.builder()
 *             .folder("folders/123456789")
 *             .constraint("serviceuser.services")
 *             .listPolicy(OrganizationPolicyListPolicyArgs.builder()
 *                 .allow(OrganizationPolicyListPolicyAllowArgs.builder()
 *                     .all(true)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * Or to deny some services, use the following instead:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.folder.OrganizationPolicy;
 * import com.pulumi.gcp.folder.OrganizationPolicyArgs;
 * import com.pulumi.gcp.folder.inputs.OrganizationPolicyListPolicyArgs;
 * import com.pulumi.gcp.folder.inputs.OrganizationPolicyListPolicyDenyArgs;
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
 *         var servicesPolicy = new OrganizationPolicy("servicesPolicy", OrganizationPolicyArgs.builder()
 *             .folder("folders/123456789")
 *             .constraint("serviceuser.services")
 *             .listPolicy(OrganizationPolicyListPolicyArgs.builder()
 *                 .suggestedValue("compute.googleapis.com")
 *                 .deny(OrganizationPolicyListPolicyDenyArgs.builder()
 *                     .values("cloudresourcemanager.googleapis.com")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * To restore the default folder organization policy, use the following instead:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.folder.OrganizationPolicy;
 * import com.pulumi.gcp.folder.OrganizationPolicyArgs;
 * import com.pulumi.gcp.folder.inputs.OrganizationPolicyRestorePolicyArgs;
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
 *         var servicesPolicy = new OrganizationPolicy("servicesPolicy", OrganizationPolicyArgs.builder()
 *             .folder("folders/123456789")
 *             .constraint("serviceuser.services")
 *             .restorePolicy(OrganizationPolicyRestorePolicyArgs.builder()
 *                 .default_(true)
 *                 .build())
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
 * Folder organization policies can be imported using any of the follow formats:
 * 
 * * `folders/{{folder_id}}/constraints/serviceuser.services`
 * 
 * * `{{folder_id}}/serviceuser.services`
 * 
 * When using the `pulumi import` command, folder organization policies can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:folder/organizationPolicy:OrganizationPolicy * `google_folder_organization_policy.default folders/* ``{{folder_id}}/constraints/serviceuser.services`
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:folder/organizationPolicy:OrganizationPolicy * `* `google_folder_organization_policy.default {{folder_id}}/``serviceuser.services
 * ```
 * 
 */
@ResourceType(type="gcp:folder/organizationPolicy:OrganizationPolicy")
public class OrganizationPolicy extends com.pulumi.resources.CustomResource {
    /**
     * A boolean policy is a constraint that is either enforced or not. Structure is documented below.
     * 
     */
    @Export(name="booleanPolicy", refs={OrganizationPolicyBooleanPolicy.class}, tree="[0]")
    private Output</* @Nullable */ OrganizationPolicyBooleanPolicy> booleanPolicy;

    /**
     * @return A boolean policy is a constraint that is either enforced or not. Structure is documented below.
     * 
     */
    public Output<Optional<OrganizationPolicyBooleanPolicy>> booleanPolicy() {
        return Codegen.optional(this.booleanPolicy);
    }
    /**
     * The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
     * 
     * ***
     * 
     */
    @Export(name="constraint", refs={String.class}, tree="[0]")
    private Output<String> constraint;

    /**
     * @return The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
     * 
     * ***
     * 
     */
    public Output<String> constraint() {
        return this.constraint;
    }
    /**
     * (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
     * 
     */
    @Export(name="folder", refs={String.class}, tree="[0]")
    private Output<String> folder;

    /**
     * @return The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
     * 
     */
    public Output<String> folder() {
        return this.folder;
    }
    /**
     * A policy that can define specific values that are allowed or denied for the given constraint. It
     * can also be used to allow or deny all values. Structure is documented below.
     * 
     */
    @Export(name="listPolicy", refs={OrganizationPolicyListPolicy.class}, tree="[0]")
    private Output</* @Nullable */ OrganizationPolicyListPolicy> listPolicy;

    /**
     * @return A policy that can define specific values that are allowed or denied for the given constraint. It
     * can also be used to allow or deny all values. Structure is documented below.
     * 
     */
    public Output<Optional<OrganizationPolicyListPolicy>> listPolicy() {
        return Codegen.optional(this.listPolicy);
    }
    /**
     * A restore policy is a constraint to restore the default policy. Structure is documented below.
     * 
     * &gt; **Note:** If none of [`boolean_policy`, `list_policy`, `restore_policy`] are defined the policy for a given constraint will
     * effectively be unset. This is represented in the UI as the constraint being &#39;Inherited&#39;.
     * 
     * ***
     * 
     */
    @Export(name="restorePolicy", refs={OrganizationPolicyRestorePolicy.class}, tree="[0]")
    private Output</* @Nullable */ OrganizationPolicyRestorePolicy> restorePolicy;

    /**
     * @return A restore policy is a constraint to restore the default policy. Structure is documented below.
     * 
     * &gt; **Note:** If none of [`boolean_policy`, `list_policy`, `restore_policy`] are defined the policy for a given constraint will
     * effectively be unset. This is represented in the UI as the constraint being &#39;Inherited&#39;.
     * 
     * ***
     * 
     */
    public Output<Optional<OrganizationPolicyRestorePolicy>> restorePolicy() {
        return Codegen.optional(this.restorePolicy);
    }
    /**
     * (Computed) The timestamp in RFC3339 UTC &#34;Zulu&#34; format, accurate to nanoseconds, representing when the variable was last updated. Example: &#34;2016-10-09T12:33:37.578138407Z&#34;.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return (Computed) The timestamp in RFC3339 UTC &#34;Zulu&#34; format, accurate to nanoseconds, representing when the variable was last updated. Example: &#34;2016-10-09T12:33:37.578138407Z&#34;.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }
    /**
     * Version of the Policy. Default version is 0.
     * 
     */
    @Export(name="version", refs={Integer.class}, tree="[0]")
    private Output<Integer> version;

    /**
     * @return Version of the Policy. Default version is 0.
     * 
     */
    public Output<Integer> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public OrganizationPolicy(String name) {
        this(name, OrganizationPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OrganizationPolicy(String name, OrganizationPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OrganizationPolicy(String name, OrganizationPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:folder/organizationPolicy:OrganizationPolicy", name, args == null ? OrganizationPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private OrganizationPolicy(String name, Output<String> id, @Nullable OrganizationPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:folder/organizationPolicy:OrganizationPolicy", name, state, makeResourceOptions(options, id));
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
    public static OrganizationPolicy get(String name, Output<String> id, @Nullable OrganizationPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OrganizationPolicy(name, id, state, options);
    }
}
