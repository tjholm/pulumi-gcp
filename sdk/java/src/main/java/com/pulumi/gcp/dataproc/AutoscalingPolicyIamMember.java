// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataproc;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.dataproc.AutoscalingPolicyIamMemberArgs;
import com.pulumi.gcp.dataproc.inputs.AutoscalingPolicyIamMemberState;
import com.pulumi.gcp.dataproc.outputs.AutoscalingPolicyIamMemberCondition;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for Dataproc AutoscalingPolicy. Each of these resources serves a different use case:
 * 
 * * `gcp.dataproc.AutoscalingPolicyIamPolicy`: Authoritative. Sets the IAM policy for the autoscalingpolicy and replaces any existing policy already attached.
 * * `gcp.dataproc.AutoscalingPolicyIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the autoscalingpolicy are preserved.
 * * `gcp.dataproc.AutoscalingPolicyIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the autoscalingpolicy are preserved.
 * 
 * &gt; **Note:** `gcp.dataproc.AutoscalingPolicyIamPolicy` **cannot** be used in conjunction with `gcp.dataproc.AutoscalingPolicyIamBinding` and `gcp.dataproc.AutoscalingPolicyIamMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.dataproc.AutoscalingPolicyIamBinding` resources **can be** used in conjunction with `gcp.dataproc.AutoscalingPolicyIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_dataproc\_autoscaling\_policy\_iam\_policy
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
 *                 .role(&#34;roles/viewer&#34;)
 *                 .members(&#34;user:jane@example.com&#34;)
 *                 .build())
 *             .build()));
 * 
 *         var policy = new AutoscalingPolicyIamPolicy(&#34;policy&#34;, AutoscalingPolicyIamPolicyArgs.builder()        
 *             .project(google_dataproc_autoscaling_policy.basic().project())
 *             .location(google_dataproc_autoscaling_policy.basic().location())
 *             .policyId(google_dataproc_autoscaling_policy.basic().policy_id())
 *             .policyData(admin.apply(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_dataproc\_autoscaling\_policy\_iam\_binding
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
 *         var binding = new AutoscalingPolicyIamBinding(&#34;binding&#34;, AutoscalingPolicyIamBindingArgs.builder()        
 *             .project(google_dataproc_autoscaling_policy.basic().project())
 *             .location(google_dataproc_autoscaling_policy.basic().location())
 *             .policyId(google_dataproc_autoscaling_policy.basic().policy_id())
 *             .role(&#34;roles/viewer&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_dataproc\_autoscaling\_policy\_iam\_member
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
 *         var member = new AutoscalingPolicyIamMember(&#34;member&#34;, AutoscalingPolicyIamMemberArgs.builder()        
 *             .project(google_dataproc_autoscaling_policy.basic().project())
 *             .location(google_dataproc_autoscaling_policy.basic().location())
 *             .policyId(google_dataproc_autoscaling_policy.basic().policy_id())
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
 * For all import syntaxes, the &#34;resource in question&#34; can take any of the following forms* projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}} * {{project}}/{{location}}/{{policy_id}} * {{location}}/{{policy_id}} * {{policy_id}} Any variables not passed in the import command will be taken from the provider configuration. Dataproc autoscalingpolicy IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:dataproc/autoscalingPolicyIamMember:AutoscalingPolicyIamMember editor &#34;projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}} roles/viewer user:jane@example.com&#34;
 * ```
 * 
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:dataproc/autoscalingPolicyIamMember:AutoscalingPolicyIamMember editor &#34;projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}} roles/viewer&#34;
 * ```
 * 
 *  IAM policy imports use the identifier of the resource in question, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:dataproc/autoscalingPolicyIamMember:AutoscalingPolicyIamMember editor projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}
 * ```
 * 
 *  -&gt; **Custom Roles**If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:dataproc/autoscalingPolicyIamMember:AutoscalingPolicyIamMember")
public class AutoscalingPolicyIamMember extends com.pulumi.resources.CustomResource {
    @Export(name="condition", type=AutoscalingPolicyIamMemberCondition.class, parameters={})
    private Output</* @Nullable */ AutoscalingPolicyIamMemberCondition> condition;

    public Output<Optional<AutoscalingPolicyIamMemberCondition>> condition() {
        return Codegen.optional(this.condition);
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
    /**
     * The  location where the autoscaling policy should reside.
     * The default value is `global`.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="location", type=String.class, parameters={})
    private Output<String> location;

    /**
     * @return The  location where the autoscaling policy should reside.
     * The default value is `global`.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    @Export(name="member", type=String.class, parameters={})
    private Output<String> member;

    public Output<String> member() {
        return this.member;
    }
    /**
     * The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
     * and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
     * 3 and 50 characters.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="policyId", type=String.class, parameters={})
    private Output<String> policyId;

    /**
     * @return The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
     * and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
     * 3 and 50 characters.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> policyId() {
        return this.policyId;
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The role that should be applied. Only one
     * `gcp.dataproc.AutoscalingPolicyIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Export(name="role", type=String.class, parameters={})
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.dataproc.AutoscalingPolicyIamBinding` can be used per role. Note that custom roles must be of the format
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
    public AutoscalingPolicyIamMember(String name) {
        this(name, AutoscalingPolicyIamMemberArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AutoscalingPolicyIamMember(String name, AutoscalingPolicyIamMemberArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AutoscalingPolicyIamMember(String name, AutoscalingPolicyIamMemberArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataproc/autoscalingPolicyIamMember:AutoscalingPolicyIamMember", name, args == null ? AutoscalingPolicyIamMemberArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AutoscalingPolicyIamMember(String name, Output<String> id, @Nullable AutoscalingPolicyIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataproc/autoscalingPolicyIamMember:AutoscalingPolicyIamMember", name, state, makeResourceOptions(options, id));
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
    public static AutoscalingPolicyIamMember get(String name, Output<String> id, @Nullable AutoscalingPolicyIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AutoscalingPolicyIamMember(name, id, state, options);
    }
}
