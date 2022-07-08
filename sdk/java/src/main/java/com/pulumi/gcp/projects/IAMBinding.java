// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.projects;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.projects.IAMBindingArgs;
import com.pulumi.gcp.projects.inputs.IAMBindingState;
import com.pulumi.gcp.projects.outputs.IAMBindingCondition;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Four different resources help you manage your IAM policy for a project. Each of these resources serves a different use case:
 * 
 * * `gcp.projects.IAMPolicy`: Authoritative. Sets the IAM policy for the project and replaces any existing policy already attached.
 * * `gcp.projects.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the project are preserved.
 * * `gcp.projects.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the project are preserved.
 * * `gcp.projects.IAMAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.
 * 
 * &gt; **Note:** `gcp.projects.IAMPolicy` **cannot** be used in conjunction with `gcp.projects.IAMBinding`, `gcp.projects.IAMMember`, or `gcp.projects.IAMAuditConfig` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.projects.IAMBinding` resources **can be** used in conjunction with `gcp.projects.IAMMember` resources **only if** they do not grant privilege to the same role.
 * 
 * &gt; **Note:** The underlying API method `projects.setIamPolicy` has a lot of constraints which are documented [here](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy). In addition to these constraints,
 *    IAM Conditions cannot be used with Basic Roles such as Owner. Violating these constraints will result in the API returning 400 error code so please review these if you encounter errors with this resource.
 * 
 * ## google\_project\_iam\_policy
 * 
 * !&gt; **Be careful!** You can accidentally lock yourself out of your project
 *    using this resource. Deleting a `gcp.projects.IAMPolicy` removes access
 *    from anyone without organization-level access to the project. Proceed with caution.
 *    It&#39;s not recommended to use `gcp.projects.IAMPolicy` with your provider project
 *    to avoid locking yourself out, and it should generally only be used with projects
 *    fully managed by this provider. If you do use this resource, it is recommended to **import** the policy before
 *    applying the change.
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
 *         final var admin = Output.of(OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role(&#34;roles/editor&#34;)
 *                 .members(&#34;user:jane@example.com&#34;)
 *                 .build())
 *             .build()));
 * 
 *         var project = new IAMPolicy(&#34;project&#34;, IAMPolicyArgs.builder()        
 *             .project(&#34;your-project-id&#34;)
 *             .policyData(admin.apply(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * With IAM Conditions:
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
 *         final var admin = Output.of(OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .condition(GetIAMPolicyBindingConditionArgs.builder()
 *                     .description(&#34;Expiring at midnight of 2019-12-31&#34;)
 *                     .expression(&#34;request.time &lt; timestamp(\&#34;2020-01-01T00:00:00Z\&#34;)&#34;)
 *                     .title(&#34;expires_after_2019_12_31&#34;)
 *                     .build())
 *                 .members(&#34;user:jane@example.com&#34;)
 *                 .role(&#34;roles/compute.admin&#34;)
 *                 .build())
 *             .build()));
 * 
 *         var project = new IAMPolicy(&#34;project&#34;, IAMPolicyArgs.builder()        
 *             .policyData(admin.apply(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .project(&#34;your-project-id&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_project\_iam\_binding
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
 *         var project = new IAMBinding(&#34;project&#34;, IAMBindingArgs.builder()        
 *             .members(&#34;user:jane@example.com&#34;)
 *             .project(&#34;your-project-id&#34;)
 *             .role(&#34;roles/editor&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * With IAM Conditions:
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
 *         var project = new IAMBinding(&#34;project&#34;, IAMBindingArgs.builder()        
 *             .condition(IAMBindingConditionArgs.builder()
 *                 .description(&#34;Expiring at midnight of 2019-12-31&#34;)
 *                 .expression(&#34;request.time &lt; timestamp(\&#34;2020-01-01T00:00:00Z\&#34;)&#34;)
 *                 .title(&#34;expires_after_2019_12_31&#34;)
 *                 .build())
 *             .members(&#34;user:jane@example.com&#34;)
 *             .project(&#34;your-project-id&#34;)
 *             .role(&#34;roles/container.admin&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_project\_iam\_member
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
 *         var project = new IAMMember(&#34;project&#34;, IAMMemberArgs.builder()        
 *             .member(&#34;user:jane@example.com&#34;)
 *             .project(&#34;your-project-id&#34;)
 *             .role(&#34;roles/editor&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * With IAM Conditions:
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
 *         var project = new IAMMember(&#34;project&#34;, IAMMemberArgs.builder()        
 *             .condition(IAMMemberConditionArgs.builder()
 *                 .description(&#34;Expiring at midnight of 2019-12-31&#34;)
 *                 .expression(&#34;request.time &lt; timestamp(\&#34;2020-01-01T00:00:00Z\&#34;)&#34;)
 *                 .title(&#34;expires_after_2019_12_31&#34;)
 *                 .build())
 *             .member(&#34;user:jane@example.com&#34;)
 *             .project(&#34;your-project-id&#34;)
 *             .role(&#34;roles/firebase.admin&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_project\_iam\_audit\_config
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
 *         var project = new IAMAuditConfig(&#34;project&#34;, IAMAuditConfigArgs.builder()        
 *             .auditLogConfigs(            
 *                 IAMAuditConfigAuditLogConfigArgs.builder()
 *                     .logType(&#34;ADMIN_READ&#34;)
 *                     .build(),
 *                 IAMAuditConfigAuditLogConfigArgs.builder()
 *                     .exemptedMembers(&#34;user:joebloggs@hashicorp.com&#34;)
 *                     .logType(&#34;DATA_READ&#34;)
 *                     .build())
 *             .project(&#34;your-project-id&#34;)
 *             .service(&#34;allServices&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
 * 
 * This member resource can be imported using the `project_id`, role, and member e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:projects/iAMBinding:IAMBinding my_project &#34;your-project-id roles/viewer user:foo@example.com&#34;
 * ```
 * 
 *  IAM binding imports use space-delimited identifiers; the resource in question and the role.
 * 
 * This binding resource can be imported using the `project_id` and role, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:projects/iAMBinding:IAMBinding my_project &#34;your-project-id roles/viewer&#34;
 * ```
 * 
 *  IAM policy imports use the identifier of the resource in question.
 * 
 * This policy resource can be imported using the `project_id`.
 * 
 * ```sh
 *  $ pulumi import gcp:projects/iAMBinding:IAMBinding my_project your-project-id
 * ```
 * 
 *  IAM audit config imports use the identifier of the resource in question and the service, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:projects/iAMBinding:IAMBinding my_project &#34;your-project-id foo.googleapis.com&#34;
 * ```
 * 
 *  -&gt; **Custom Roles**If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`. -&gt; **Conditional IAM Bindings**If you&#39;re importing a IAM binding with a condition block, make sure
 * 
 * ```sh
 *  $ pulumi import gcp:projects/iAMBinding:IAMBinding to include the title of condition, e.g. `google_project_iam_binding.my_project &#34;{{your-project-id}} roles/{{role_id}} condition-title&#34;`
 * ```
 * 
 */
@ResourceType(type="gcp:projects/iAMBinding:IAMBinding")
public class IAMBinding extends com.pulumi.resources.CustomResource {
    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     * 
     */
    @Export(name="condition", type=IAMBindingCondition.class, parameters={})
    private Output</* @Nullable */ IAMBindingCondition> condition;

    /**
     * @return An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     * 
     */
    public Output<Optional<IAMBindingCondition>> condition() {
        return Codegen.optional(this.condition);
    }
    /**
     * (Computed) The etag of the project&#39;s IAM policy.
     * 
     */
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    /**
     * @return (Computed) The etag of the project&#39;s IAM policy.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    @Export(name="members", type=List.class, parameters={String.class})
    private Output<List<String>> members;

    public Output<List<String>> members() {
        return this.members;
    }
    /**
     * The project id of the target project. This is not
     * inferred from the provider.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
    private Output<String> project;

    /**
     * @return The project id of the target project. This is not
     * inferred from the provider.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The role that should be applied. Only one
     * `gcp.projects.IAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Export(name="role", type=String.class, parameters={})
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.projects.IAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IAMBinding(String name) {
        this(name, IAMBindingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IAMBinding(String name, IAMBindingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IAMBinding(String name, IAMBindingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:projects/iAMBinding:IAMBinding", name, args == null ? IAMBindingArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IAMBinding(String name, Output<String> id, @Nullable IAMBindingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:projects/iAMBinding:IAMBinding", name, state, makeResourceOptions(options, id));
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
    public static IAMBinding get(String name, Output<String> id, @Nullable IAMBindingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IAMBinding(name, id, state, options);
    }
}
