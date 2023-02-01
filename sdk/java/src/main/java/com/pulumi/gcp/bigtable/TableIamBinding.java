// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigtable;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.bigtable.TableIamBindingArgs;
import com.pulumi.gcp.bigtable.inputs.TableIamBindingState;
import com.pulumi.gcp.bigtable.outputs.TableIamBindingCondition;
import java.lang.String;
import java.util.List;
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
 * 
 * ## google\_bigtable\_table\_iam\_binding
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
 *             .instance(&#34;your-bigtable-instance&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
 *             .role(&#34;roles/bigtable.user&#34;)
 *             .table(&#34;your-bigtable-table&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_bigtable\_table\_iam\_member
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
 *             .instance(&#34;your-bigtable-instance&#34;)
 *             .member(&#34;user:jane@example.com&#34;)
 *             .role(&#34;roles/bigtable.user&#34;)
 *             .table(&#34;your-bigtable-table&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Table IAM resources can be imported using the project, table name, role and/or member.
 * 
 * ```sh
 *  $ pulumi import gcp:bigtable/tableIamBinding:TableIamBinding editor &#34;projects/{project}/tables/{table}&#34;
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:bigtable/tableIamBinding:TableIamBinding editor &#34;projects/{project}/tables/{table} roles/editor&#34;
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:bigtable/tableIamBinding:TableIamBinding editor &#34;projects/{project}/tables/{table} roles/editor user:jane@example.com&#34;
 * ```
 * 
 *  -&gt; **Custom Roles**If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:bigtable/tableIamBinding:TableIamBinding")
public class TableIamBinding extends com.pulumi.resources.CustomResource {
    @Export(name="condition", type=TableIamBindingCondition.class, parameters={})
    private Output</* @Nullable */ TableIamBindingCondition> condition;

    public Output<Optional<TableIamBindingCondition>> condition() {
        return Codegen.optional(this.condition);
    }
    /**
     * (Computed) The etag of the tables&#39;s IAM policy.
     * 
     */
    @Export(name="etag", type=String.class, parameters={})
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
    @Export(name="instance", type=String.class, parameters={})
    private Output<String> instance;

    /**
     * @return The name or relative resource id of the instance that owns the table.
     * 
     */
    public Output<String> instance() {
        return this.instance;
    }
    @Export(name="members", type=List.class, parameters={String.class})
    private Output<List<String>> members;

    public Output<List<String>> members() {
        return this.members;
    }
    @Export(name="project", type=String.class, parameters={})
    private Output<String> project;

    public Output<String> project() {
        return this.project;
    }
    /**
     * The role that should be applied. Only one
     * `gcp.bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
     * 
     */
    @Export(name="role", type=String.class, parameters={})
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
     * 
     */
    public Output<String> role() {
        return this.role;
    }
    /**
     * The name or relative resource id of the table to manage IAM policies for.
     * 
     */
    @Export(name="table", type=String.class, parameters={})
    private Output<String> table;

    /**
     * @return The name or relative resource id of the table to manage IAM policies for.
     * 
     */
    public Output<String> table() {
        return this.table;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TableIamBinding(String name) {
        this(name, TableIamBindingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TableIamBinding(String name, TableIamBindingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TableIamBinding(String name, TableIamBindingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigtable/tableIamBinding:TableIamBinding", name, args == null ? TableIamBindingArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TableIamBinding(String name, Output<String> id, @Nullable TableIamBindingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigtable/tableIamBinding:TableIamBinding", name, state, makeResourceOptions(options, id));
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
    public static TableIamBinding get(String name, Output<String> id, @Nullable TableIamBindingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TableIamBinding(name, id, state, options);
    }
}
