// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.iap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.iap.TunnelIamMemberArgs;
import com.pulumi.gcp.iap.inputs.TunnelIamMemberState;
import com.pulumi.gcp.iap.outputs.TunnelIamMemberCondition;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for Identity-Aware Proxy Tunnel. Each of these resources serves a different use case:
 * 
 * * `gcp.iap.TunnelIamPolicy`: Authoritative. Sets the IAM policy for the tunnel and replaces any existing policy already attached.
 * * `gcp.iap.TunnelIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the tunnel are preserved.
 * * `gcp.iap.TunnelIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the tunnel are preserved.
 * 
 * &gt; **Note:** `gcp.iap.TunnelIamPolicy` **cannot** be used in conjunction with `gcp.iap.TunnelIamBinding` and `gcp.iap.TunnelIamMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.iap.TunnelIamBinding` resources **can be** used in conjunction with `gcp.iap.TunnelIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_iap\_tunnel\_iam\_policy
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
 *                 .role(&#34;roles/iap.tunnelResourceAccessor&#34;)
 *                 .members(&#34;user:jane@example.com&#34;)
 *                 .build())
 *             .build()));
 * 
 *         var policy = new TunnelIamPolicy(&#34;policy&#34;, TunnelIamPolicyArgs.builder()        
 *             .project(google_project_service.project_service().project())
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
 *                 .role(&#34;roles/iap.tunnelResourceAccessor&#34;)
 *                 .members(&#34;user:jane@example.com&#34;)
 *                 .condition(GetIAMPolicyBindingConditionArgs.builder()
 *                     .title(&#34;expires_after_2019_12_31&#34;)
 *                     .description(&#34;Expiring at midnight of 2019-12-31&#34;)
 *                     .expression(&#34;request.time &lt; timestamp(\&#34;2020-01-01T00:00:00Z\&#34;)&#34;)
 *                     .build())
 *                 .build())
 *             .build()));
 * 
 *         var policy = new TunnelIamPolicy(&#34;policy&#34;, TunnelIamPolicyArgs.builder()        
 *             .project(google_project_service.project_service().project())
 *             .policyData(admin.apply(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## google\_iap\_tunnel\_iam\_binding
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
 *         var binding = new TunnelIamBinding(&#34;binding&#34;, TunnelIamBindingArgs.builder()        
 *             .project(google_project_service.project_service().project())
 *             .role(&#34;roles/iap.tunnelResourceAccessor&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
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
 *         var binding = new TunnelIamBinding(&#34;binding&#34;, TunnelIamBindingArgs.builder()        
 *             .project(google_project_service.project_service().project())
 *             .role(&#34;roles/iap.tunnelResourceAccessor&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
 *             .condition(TunnelIamBindingConditionArgs.builder()
 *                 .title(&#34;expires_after_2019_12_31&#34;)
 *                 .description(&#34;Expiring at midnight of 2019-12-31&#34;)
 *                 .expression(&#34;request.time &lt; timestamp(\&#34;2020-01-01T00:00:00Z\&#34;)&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## google\_iap\_tunnel\_iam\_member
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
 *         var member = new TunnelIamMember(&#34;member&#34;, TunnelIamMemberArgs.builder()        
 *             .project(google_project_service.project_service().project())
 *             .role(&#34;roles/iap.tunnelResourceAccessor&#34;)
 *             .member(&#34;user:jane@example.com&#34;)
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
 *         var member = new TunnelIamMember(&#34;member&#34;, TunnelIamMemberArgs.builder()        
 *             .project(google_project_service.project_service().project())
 *             .role(&#34;roles/iap.tunnelResourceAccessor&#34;)
 *             .member(&#34;user:jane@example.com&#34;)
 *             .condition(TunnelIamMemberConditionArgs.builder()
 *                 .title(&#34;expires_after_2019_12_31&#34;)
 *                 .description(&#34;Expiring at midnight of 2019-12-31&#34;)
 *                 .expression(&#34;request.time &lt; timestamp(\&#34;2020-01-01T00:00:00Z\&#34;)&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * For all import syntaxes, the &#34;resource in question&#34; can take any of the following forms* projects/{{project}}/iap_tunnel * {{project}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy tunnel IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:iap/tunnelIamMember:TunnelIamMember editor &#34;projects/{{project}}/iap_tunnel roles/iap.tunnelResourceAccessor user:jane@example.com&#34;
 * ```
 * 
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:iap/tunnelIamMember:TunnelIamMember editor &#34;projects/{{project}}/iap_tunnel roles/iap.tunnelResourceAccessor&#34;
 * ```
 * 
 *  IAM policy imports use the identifier of the resource in question, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:iap/tunnelIamMember:TunnelIamMember editor projects/{{project}}/iap_tunnel
 * ```
 * 
 *  -&gt; **Custom Roles**If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:iap/tunnelIamMember:TunnelIamMember")
public class TunnelIamMember extends com.pulumi.resources.CustomResource {
    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     * 
     */
    @Export(name="condition", type=TunnelIamMemberCondition.class, parameters={})
    private Output</* @Nullable */ TunnelIamMemberCondition> condition;

    /**
     * @return An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     * 
     */
    public Output<Optional<TunnelIamMemberCondition>> condition() {
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
    @Export(name="member", type=String.class, parameters={})
    private Output<String> member;

    public Output<String> member() {
        return this.member;
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
     * `gcp.iap.TunnelIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Export(name="role", type=String.class, parameters={})
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.iap.TunnelIamBinding` can be used per role. Note that custom roles must be of the format
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
    public TunnelIamMember(String name) {
        this(name, TunnelIamMemberArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TunnelIamMember(String name, TunnelIamMemberArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TunnelIamMember(String name, TunnelIamMemberArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:iap/tunnelIamMember:TunnelIamMember", name, args == null ? TunnelIamMemberArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TunnelIamMember(String name, Output<String> id, @Nullable TunnelIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:iap/tunnelIamMember:TunnelIamMember", name, state, makeResourceOptions(options, id));
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
    public static TunnelIamMember get(String name, Output<String> id, @Nullable TunnelIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TunnelIamMember(name, id, state, options);
    }
}
