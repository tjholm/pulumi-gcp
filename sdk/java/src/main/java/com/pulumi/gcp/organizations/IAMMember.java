// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.organizations;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.organizations.IAMMemberArgs;
import com.pulumi.gcp.organizations.inputs.IAMMemberState;
import com.pulumi.gcp.organizations.outputs.IAMMemberCondition;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Four different resources help you manage your IAM policy for a organization. Each of these resources serves a different use case:
 * 
 * * `gcp.organizations.IAMPolicy`: Authoritative. Sets the IAM policy for the organization and replaces any existing policy already attached.
 * * `gcp.organizations.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the organization are preserved.
 * * `gcp.organizations.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the organization are preserved.
 * * `gcp.organizations.IamAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.
 * 
 * &gt; **Note:** `gcp.organizations.IAMPolicy` **cannot** be used in conjunction with `gcp.organizations.IAMBinding`, `gcp.organizations.IAMMember`, or `gcp.organizations.IamAuditConfig` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.organizations.IAMBinding` resources **can be** used in conjunction with `gcp.organizations.IAMMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_organization\_iam\_policy
 * 
 * !&gt; **Warning:** New organizations have several default policies which will,
 *    without extreme caution, be **overwritten** by use of this resource.
 *    The safest alternative is to use multiple `gcp.organizations.IAMBinding`
 *    resources. This resource makes it easy to remove your own access to
 *    an organization, which will require a call to Google Support to have
 *    fixed, and can take multiple days to resolve.
 * 
 *    In general, this resource should only be used with organizations
 *    fully managed by this provider.I f you do use this resource,
 *    the best way to be sure that you are not making dangerous changes is to start
 *    by **importing** your existing policy, and examining the diff very closely.
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
 *         var organization = new IAMPolicy(&#34;organization&#34;, IAMPolicyArgs.builder()        
 *             .orgId(&#34;your-organization-id&#34;)
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
 *                 .role(&#34;roles/editor&#34;)
 *                 .build())
 *             .build()));
 * 
 *         var organization = new IAMPolicy(&#34;organization&#34;, IAMPolicyArgs.builder()        
 *             .orgId(&#34;your-organization-id&#34;)
 *             .policyData(admin.apply(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_organization\_iam\_binding
 * 
 * &gt; **Note:** If `role` is set to `roles/owner` and you don&#39;t specify a user or service account you have access to in `members`, you can lock yourself out of your organization.
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
 *         var organization = new IAMBinding(&#34;organization&#34;, IAMBindingArgs.builder()        
 *             .members(&#34;user:jane@example.com&#34;)
 *             .orgId(&#34;your-organization-id&#34;)
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
 *         var organization = new IAMBinding(&#34;organization&#34;, IAMBindingArgs.builder()        
 *             .condition(IAMBindingConditionArgs.builder()
 *                 .description(&#34;Expiring at midnight of 2019-12-31&#34;)
 *                 .expression(&#34;request.time &lt; timestamp(\&#34;2020-01-01T00:00:00Z\&#34;)&#34;)
 *                 .title(&#34;expires_after_2019_12_31&#34;)
 *                 .build())
 *             .members(&#34;user:jane@example.com&#34;)
 *             .orgId(&#34;your-organization-id&#34;)
 *             .role(&#34;roles/editor&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_organization\_iam\_member
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
 *         var organization = new IAMMember(&#34;organization&#34;, IAMMemberArgs.builder()        
 *             .member(&#34;user:jane@example.com&#34;)
 *             .orgId(&#34;your-organization-id&#34;)
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
 *         var organization = new IAMMember(&#34;organization&#34;, IAMMemberArgs.builder()        
 *             .condition(IAMMemberConditionArgs.builder()
 *                 .description(&#34;Expiring at midnight of 2019-12-31&#34;)
 *                 .expression(&#34;request.time &lt; timestamp(\&#34;2020-01-01T00:00:00Z\&#34;)&#34;)
 *                 .title(&#34;expires_after_2019_12_31&#34;)
 *                 .build())
 *             .member(&#34;user:jane@example.com&#34;)
 *             .orgId(&#34;your-organization-id&#34;)
 *             .role(&#34;roles/editor&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_organization\_iam\_audit\_config
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
 *         var organization = new IamAuditConfig(&#34;organization&#34;, IamAuditConfigArgs.builder()        
 *             .auditLogConfigs(            
 *                 IamAuditConfigAuditLogConfigArgs.builder()
 *                     .logType(&#34;ADMIN_READ&#34;)
 *                     .build(),
 *                 IamAuditConfigAuditLogConfigArgs.builder()
 *                     .exemptedMembers(&#34;user:joebloggs@hashicorp.com&#34;)
 *                     .logType(&#34;DATA_READ&#34;)
 *                     .build())
 *             .orgId(&#34;your-organization-id&#34;)
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
 * This member resource can be imported using the `org_id`, role, and member e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:organizations/iAMMember:IAMMember my_organization &#34;your-orgid roles/viewer user:foo@example.com&#34;
 * ```
 * 
 *  IAM binding imports use space-delimited identifiers; the resource in question and the role.
 * 
 * This binding resource can be imported using the `org_id` and role, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:organizations/iAMMember:IAMMember my_organization &#34;your-org-id roles/viewer&#34;
 * ```
 * 
 *  IAM policy imports use the identifier of the resource in question.
 * 
 * This policy resource can be imported using the `org_id`.
 * 
 * ```sh
 *  $ pulumi import gcp:organizations/iAMMember:IAMMember my_organization your-org-id
 * ```
 * 
 *  IAM audit config imports use the identifier of the resource in question and the service, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:organizations/iAMMember:IAMMember my_organization &#34;your-organization-id foo.googleapis.com&#34;
 * ```
 * 
 *  -&gt; **Custom Roles**If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 * full name of the custom role, e.g. `organizations/{{org_id}}/roles/{{role_id}}`. -&gt; **Conditional IAM Bindings**If you&#39;re importing a IAM binding with a condition block, make sure
 * 
 * ```sh
 *  $ pulumi import gcp:organizations/iAMMember:IAMMember to include the title of condition, e.g. `google_organization_iam_binding.my_organization &#34;your-org-id roles/{{role_id}} condition-title&#34;`
 * ```
 * 
 */
@ResourceType(type="gcp:organizations/iAMMember:IAMMember")
public class IAMMember extends com.pulumi.resources.CustomResource {
    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     * 
     */
    @Export(name="condition", type=IAMMemberCondition.class, parameters={})
    private Output</* @Nullable */ IAMMemberCondition> condition;

    /**
     * @return An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     * 
     */
    public Output<Optional<IAMMemberCondition>> condition() {
        return Codegen.optional(this.condition);
    }
    /**
     * (Computed) The etag of the organization&#39;s IAM policy.
     * 
     */
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    /**
     * @return (Computed) The etag of the organization&#39;s IAM policy.
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
     * The organization ID. If not specified for `gcp.organizations.IAMBinding`, `gcp.organizations.IAMMember`, or `gcp.organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
     * Required for `gcp.organizations.IAMPolicy` - you must explicitly set the organization, and it
     * will not be inferred from the provider.
     * 
     */
    @Export(name="orgId", type=String.class, parameters={})
    private Output<String> orgId;

    /**
     * @return The organization ID. If not specified for `gcp.organizations.IAMBinding`, `gcp.organizations.IAMMember`, or `gcp.organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
     * Required for `gcp.organizations.IAMPolicy` - you must explicitly set the organization, and it
     * will not be inferred from the provider.
     * 
     */
    public Output<String> orgId() {
        return this.orgId;
    }
    /**
     * The role that should be applied. Only one
     * `gcp.organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
     * `organizations/{{org_id}}/roles/{{role_id}}`.
     * 
     */
    @Export(name="role", type=String.class, parameters={})
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
     * `organizations/{{org_id}}/roles/{{role_id}}`.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IAMMember(String name) {
        this(name, IAMMemberArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IAMMember(String name, IAMMemberArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IAMMember(String name, IAMMemberArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:organizations/iAMMember:IAMMember", name, args == null ? IAMMemberArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IAMMember(String name, Output<String> id, @Nullable IAMMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:organizations/iAMMember:IAMMember", name, state, makeResourceOptions(options, id));
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
    public static IAMMember get(String name, Output<String> id, @Nullable IAMMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IAMMember(name, id, state, options);
    }
}
