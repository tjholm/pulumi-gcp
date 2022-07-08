// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.spanner;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.spanner.DatabaseIAMPolicyArgs;
import com.pulumi.gcp.spanner.inputs.DatabaseIAMPolicyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for a Spanner database. Each of these resources serves a different use case:
 * 
 * * `gcp.spanner.DatabaseIAMPolicy`: Authoritative. Sets the IAM policy for the database and replaces any existing policy already attached.
 * 
 * &gt; **Warning:** It&#39;s entirely possibly to lock yourself out of your database using `gcp.spanner.DatabaseIAMPolicy`. Any permissions granted by default will be removed unless you include them in your config.
 * 
 * * `gcp.spanner.DatabaseIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the database are preserved.
 * * `gcp.spanner.DatabaseIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the database are preserved.
 * 
 * &gt; **Note:** `gcp.spanner.DatabaseIAMPolicy` **cannot** be used in conjunction with `gcp.spanner.DatabaseIAMBinding` and `gcp.spanner.DatabaseIAMMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.spanner.DatabaseIAMBinding` resources **can be** used in conjunction with `gcp.spanner.DatabaseIAMMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_spanner\_database\_iam\_policy
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
 *         var database = new DatabaseIAMPolicy(&#34;database&#34;, DatabaseIAMPolicyArgs.builder()        
 *             .instance(&#34;your-instance-name&#34;)
 *             .database(&#34;your-database-name&#34;)
 *             .policyData(admin.apply(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_spanner\_database\_iam\_binding
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
 *         var database = new DatabaseIAMBinding(&#34;database&#34;, DatabaseIAMBindingArgs.builder()        
 *             .database(&#34;your-database-name&#34;)
 *             .instance(&#34;your-instance-name&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
 *             .role(&#34;roles/compute.networkUser&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_spanner\_database\_iam\_member
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
 *         var database = new DatabaseIAMMember(&#34;database&#34;, DatabaseIAMMemberArgs.builder()        
 *             .database(&#34;your-database-name&#34;)
 *             .instance(&#34;your-instance-name&#34;)
 *             .member(&#34;user:jane@example.com&#34;)
 *             .role(&#34;roles/compute.networkUser&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * For all import syntaxes, the &#34;resource in question&#34; can take any of the following forms* {{project}}/{{instance}}/{{database}} * {{instance}}/{{database}} (project is taken from provider project) IAM member imports use space-delimited identifiers; the resource in question, the role, and the member identity, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:spanner/databaseIAMPolicy:DatabaseIAMPolicy database &#34;project-name/instance-name/database-name roles/viewer user:foo@example.com&#34;
 * ```
 * 
 *  IAM binding imports use space-delimited identifiers; the resource in question and the role, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:spanner/databaseIAMPolicy:DatabaseIAMPolicy database &#34;project-name/instance-name/database-name roles/viewer&#34;
 * ```
 * 
 *  IAM policy imports use the identifier of the resource in question, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:spanner/databaseIAMPolicy:DatabaseIAMPolicy database project-name/instance-name/database-name
 * ```
 * 
 *  -&gt; **Custom Roles:** If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:spanner/databaseIAMPolicy:DatabaseIAMPolicy")
public class DatabaseIAMPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The name of the Spanner database.
     * 
     */
    @Export(name="database", type=String.class, parameters={})
    private Output<String> database;

    /**
     * @return The name of the Spanner database.
     * 
     */
    public Output<String> database() {
        return this.database;
    }
    /**
     * (Computed) The etag of the database&#39;s IAM policy.
     * 
     */
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    /**
     * @return (Computed) The etag of the database&#39;s IAM policy.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The name of the Spanner instance the database belongs to.
     * 
     */
    @Export(name="instance", type=String.class, parameters={})
    private Output<String> instance;

    /**
     * @return The name of the Spanner instance the database belongs to.
     * 
     */
    public Output<String> instance() {
        return this.instance;
    }
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     * 
     */
    @Export(name="policyData", type=String.class, parameters={})
    private Output<String> policyData;

    /**
     * @return The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     * 
     */
    public Output<String> policyData() {
        return this.policyData;
    }
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DatabaseIAMPolicy(String name) {
        this(name, DatabaseIAMPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DatabaseIAMPolicy(String name, DatabaseIAMPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DatabaseIAMPolicy(String name, DatabaseIAMPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:spanner/databaseIAMPolicy:DatabaseIAMPolicy", name, args == null ? DatabaseIAMPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DatabaseIAMPolicy(String name, Output<String> id, @Nullable DatabaseIAMPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:spanner/databaseIAMPolicy:DatabaseIAMPolicy", name, state, makeResourceOptions(options, id));
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
    public static DatabaseIAMPolicy get(String name, Output<String> id, @Nullable DatabaseIAMPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DatabaseIAMPolicy(name, id, state, options);
    }
}
