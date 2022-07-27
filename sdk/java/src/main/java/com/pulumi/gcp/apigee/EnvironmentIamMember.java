// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.apigee;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.apigee.EnvironmentIamMemberArgs;
import com.pulumi.gcp.apigee.inputs.EnvironmentIamMemberState;
import com.pulumi.gcp.apigee.outputs.EnvironmentIamMemberCondition;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for Apigee Environment. Each of these resources serves a different use case:
 * 
 * * `gcp.apigee.EnvironmentIamPolicy`: Authoritative. Sets the IAM policy for the environment and replaces any existing policy already attached.
 * * `gcp.apigee.EnvironmentIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the environment are preserved.
 * * `gcp.apigee.EnvironmentIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the environment are preserved.
 * 
 * &gt; **Note:** `gcp.apigee.EnvironmentIamPolicy` **cannot** be used in conjunction with `gcp.apigee.EnvironmentIamBinding` and `gcp.apigee.EnvironmentIamMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.apigee.EnvironmentIamBinding` resources **can be** used in conjunction with `gcp.apigee.EnvironmentIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_apigee\_environment\_iam\_policy
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetIAMPolicyArgs;
 * import com.pulumi.gcp.apigee.EnvironmentIamPolicy;
 * import com.pulumi.gcp.apigee.EnvironmentIamPolicyArgs;
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
 *         final var admin = OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role(&#34;roles/viewer&#34;)
 *                 .members(&#34;user:jane@example.com&#34;)
 *                 .build())
 *             .build());
 * 
 *         var policy = new EnvironmentIamPolicy(&#34;policy&#34;, EnvironmentIamPolicyArgs.builder()        
 *             .orgId(google_apigee_environment.apigee_environment().org_id())
 *             .envId(google_apigee_environment.apigee_environment().name())
 *             .policyData(admin.applyValue(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_apigee\_environment\_iam\_binding
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.apigee.EnvironmentIamBinding;
 * import com.pulumi.gcp.apigee.EnvironmentIamBindingArgs;
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
 *         var binding = new EnvironmentIamBinding(&#34;binding&#34;, EnvironmentIamBindingArgs.builder()        
 *             .orgId(google_apigee_environment.apigee_environment().org_id())
 *             .envId(google_apigee_environment.apigee_environment().name())
 *             .role(&#34;roles/viewer&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_apigee\_environment\_iam\_member
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.apigee.EnvironmentIamMember;
 * import com.pulumi.gcp.apigee.EnvironmentIamMemberArgs;
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
 *         var member = new EnvironmentIamMember(&#34;member&#34;, EnvironmentIamMemberArgs.builder()        
 *             .orgId(google_apigee_environment.apigee_environment().org_id())
 *             .envId(google_apigee_environment.apigee_environment().name())
 *             .role(&#34;roles/viewer&#34;)
 *             .member(&#34;user:jane@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * For all import syntaxes, the &#34;resource in question&#34; can take any of the following forms* {{org_id}}/environments/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Apigee environment IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:apigee/environmentIamMember:EnvironmentIamMember editor &#34;{{org_id}}/environments/{{environment}} roles/viewer user:jane@example.com&#34;
 * ```
 * 
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:apigee/environmentIamMember:EnvironmentIamMember editor &#34;{{org_id}}/environments/{{environment}} roles/viewer&#34;
 * ```
 * 
 *  IAM policy imports use the identifier of the resource in question, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:apigee/environmentIamMember:EnvironmentIamMember editor {{org_id}}/environments/{{environment}}
 * ```
 * 
 *  -&gt; **Custom Roles**If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:apigee/environmentIamMember:EnvironmentIamMember")
public class EnvironmentIamMember extends com.pulumi.resources.CustomResource {
    @Export(name="condition", type=EnvironmentIamMemberCondition.class, parameters={})
    private Output</* @Nullable */ EnvironmentIamMemberCondition> condition;

    public Output<Optional<EnvironmentIamMemberCondition>> condition() {
        return Codegen.optional(this.condition);
    }
    /**
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="envId", type=String.class, parameters={})
    private Output<String> envId;

    /**
     * @return Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> envId() {
        return this.envId;
    }
    /**
     * (Computed) The etag of the IAM policy.
     * 
     */
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    /**
     * @return (Computed) The etag of the IAM policy.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    @Export(name="member", type=String.class, parameters={})
    private Output<String> member;

    public Output<String> member() {
        return this.member;
    }
    @Export(name="orgId", type=String.class, parameters={})
    private Output<String> orgId;

    public Output<String> orgId() {
        return this.orgId;
    }
    /**
     * The role that should be applied. Only one
     * `gcp.apigee.EnvironmentIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Export(name="role", type=String.class, parameters={})
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.apigee.EnvironmentIamBinding` can be used per role. Note that custom roles must be of the format
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
    public EnvironmentIamMember(String name) {
        this(name, EnvironmentIamMemberArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EnvironmentIamMember(String name, EnvironmentIamMemberArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EnvironmentIamMember(String name, EnvironmentIamMemberArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigee/environmentIamMember:EnvironmentIamMember", name, args == null ? EnvironmentIamMemberArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EnvironmentIamMember(String name, Output<String> id, @Nullable EnvironmentIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigee/environmentIamMember:EnvironmentIamMember", name, state, makeResourceOptions(options, id));
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
    public static EnvironmentIamMember get(String name, Output<String> id, @Nullable EnvironmentIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EnvironmentIamMember(name, id, state, options);
    }
}
