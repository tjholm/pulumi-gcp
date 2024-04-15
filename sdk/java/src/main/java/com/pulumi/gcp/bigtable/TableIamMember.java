// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigtable;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.bigtable.TableIamMemberArgs;
import com.pulumi.gcp.bigtable.inputs.TableIamMemberState;
import com.pulumi.gcp.bigtable.outputs.TableIamMemberCondition;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage IAM policies on bigtable tables. Each of these resources serves a different use case:
 * 
 * * `gcp.bigtable.TableIamPolicy`: Authoritative. Sets the IAM policy for the tables and replaces any existing policy already attached.
 * * `gcp.bigtable.TableIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
 * * `gcp.bigtable.TableIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the table are preserved.
 * 
 * &gt; **Note:** `gcp.bigtable.TableIamPolicy` **cannot** be used in conjunction with `gcp.bigtable.TableIamBinding` and `gcp.bigtable.TableIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the table as `gcp.bigtable.TableIamPolicy` replaces the entire policy.
 * 
 * &gt; **Note:** `gcp.bigtable.TableIamBinding` resources **can be** used in conjunction with `gcp.bigtable.TableIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_bigtable\_table\_iam\_policy
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetIAMPolicyArgs;
 * import com.pulumi.gcp.bigtable.TableIamPolicy;
 * import com.pulumi.gcp.bigtable.TableIamPolicyArgs;
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
 *                 .role(&#34;roles/bigtable.user&#34;)
 *                 .members(&#34;user:jane@example.com&#34;)
 *                 .build())
 *             .build());
 * 
 *         var editor = new TableIamPolicy(&#34;editor&#34;, TableIamPolicyArgs.builder()        
 *             .project(&#34;your-project&#34;)
 *             .instance(&#34;your-bigtable-instance&#34;)
 *             .table(&#34;your-bigtable-table&#34;)
 *             .policyData(admin.applyValue(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## google\_bigtable\_table\_iam\_binding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigtable.TableIamBinding;
 * import com.pulumi.gcp.bigtable.TableIamBindingArgs;
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
 *         var editor = new TableIamBinding(&#34;editor&#34;, TableIamBindingArgs.builder()        
 *             .table(&#34;your-bigtable-table&#34;)
 *             .instance(&#34;your-bigtable-instance&#34;)
 *             .role(&#34;roles/bigtable.user&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## google\_bigtable\_table\_iam\_member
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigtable.TableIamMember;
 * import com.pulumi.gcp.bigtable.TableIamMemberArgs;
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
 *         var editor = new TableIamMember(&#34;editor&#34;, TableIamMemberArgs.builder()        
 *             .table(&#34;your-bigtable-table&#34;)
 *             .instance(&#34;your-bigtable-instance&#34;)
 *             .role(&#34;roles/bigtable.user&#34;)
 *             .member(&#34;user:jane@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## google\_bigtable\_table\_iam\_policy
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetIAMPolicyArgs;
 * import com.pulumi.gcp.bigtable.TableIamPolicy;
 * import com.pulumi.gcp.bigtable.TableIamPolicyArgs;
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
 *                 .role(&#34;roles/bigtable.user&#34;)
 *                 .members(&#34;user:jane@example.com&#34;)
 *                 .build())
 *             .build());
 * 
 *         var editor = new TableIamPolicy(&#34;editor&#34;, TableIamPolicyArgs.builder()        
 *             .project(&#34;your-project&#34;)
 *             .instance(&#34;your-bigtable-instance&#34;)
 *             .table(&#34;your-bigtable-table&#34;)
 *             .policyData(admin.applyValue(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## google\_bigtable\_table\_iam\_binding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigtable.TableIamBinding;
 * import com.pulumi.gcp.bigtable.TableIamBindingArgs;
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
 *         var editor = new TableIamBinding(&#34;editor&#34;, TableIamBindingArgs.builder()        
 *             .table(&#34;your-bigtable-table&#34;)
 *             .instance(&#34;your-bigtable-instance&#34;)
 *             .role(&#34;roles/bigtable.user&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## google\_bigtable\_table\_iam\_member
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigtable.TableIamMember;
 * import com.pulumi.gcp.bigtable.TableIamMemberArgs;
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
 *         var editor = new TableIamMember(&#34;editor&#34;, TableIamMemberArgs.builder()        
 *             .table(&#34;your-bigtable-table&#34;)
 *             .instance(&#34;your-bigtable-instance&#34;)
 *             .role(&#34;roles/bigtable.user&#34;)
 *             .member(&#34;user:jane@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ### Importing IAM policies
 * 
 * IAM policy imports use the `table` identifier of the Bigtable Table resource only. For example:
 * 
 * * `&#34;projects/{project}/instances/{instance}/tables/{table}&#34;`
 * 
 * An `import` block (Terraform v1.5.0 and later) can be used to import IAM policies:
 * 
 * tf
 * 
 * import {
 * 
 *   id = &#34;projects/{project}/instances/{instance}/tables/{table}&#34;
 * 
 *   to = google_bigtable_table_iam_policy.default
 * 
 * }
 * 
 * The `pulumi import` command can also be used:
 * 
 * ```sh
 * $ pulumi import gcp:bigtable/tableIamMember:TableIamMember default projects/{project}/instances/{instance}/tables/{table}
 * ```
 * 
 */
@ResourceType(type="gcp:bigtable/tableIamMember:TableIamMember")
public class TableIamMember extends com.pulumi.resources.CustomResource {
    @Export(name="condition", refs={TableIamMemberCondition.class}, tree="[0]")
    private Output</* @Nullable */ TableIamMemberCondition> condition;

    public Output<Optional<TableIamMemberCondition>> condition() {
        return Codegen.optional(this.condition);
    }
    /**
     * (Computed) The etag of the tables&#39;s IAM policy.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return (Computed) The etag of the tables&#39;s IAM policy.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The name or relative resource id of the instance that owns the table.
     * 
     */
    @Export(name="instance", refs={String.class}, tree="[0]")
    private Output<String> instance;

    /**
     * @return The name or relative resource id of the instance that owns the table.
     * 
     */
    public Output<String> instance() {
        return this.instance;
    }
    /**
     * Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     * 
     */
    @Export(name="member", refs={String.class}, tree="[0]")
    private Output<String> member;

    /**
     * @return Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     * 
     */
    public Output<String> member() {
        return this.member;
    }
    /**
     * The project in which the table belongs. If it
     * is not provided, this provider will use the provider default.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The project in which the table belongs. If it
     * is not provided, this provider will use the provider default.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The role that should be applied. Only one
     * `gcp.bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
     * 
     * `gcp.bigtable.TableIamPolicy` only:
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
     * 
     * `gcp.bigtable.TableIamPolicy` only:
     * 
     */
    public Output<String> role() {
        return this.role;
    }
    /**
     * The name or relative resource id of the table to manage IAM policies for.
     * 
     * For `gcp.bigtable.TableIamMember` or `gcp.bigtable.TableIamBinding`:
     * 
     */
    @Export(name="table", refs={String.class}, tree="[0]")
    private Output<String> table;

    /**
     * @return The name or relative resource id of the table to manage IAM policies for.
     * 
     * For `gcp.bigtable.TableIamMember` or `gcp.bigtable.TableIamBinding`:
     * 
     */
    public Output<String> table() {
        return this.table;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TableIamMember(String name) {
        this(name, TableIamMemberArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TableIamMember(String name, TableIamMemberArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TableIamMember(String name, TableIamMemberArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigtable/tableIamMember:TableIamMember", name, args == null ? TableIamMemberArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TableIamMember(String name, Output<String> id, @Nullable TableIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigtable/tableIamMember:TableIamMember", name, state, makeResourceOptions(options, id));
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
    public static TableIamMember get(String name, Output<String> id, @Nullable TableIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TableIamMember(name, id, state, options);
    }
}
