// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigquery;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.bigquery.ConnectionIamMemberArgs;
import com.pulumi.gcp.bigquery.inputs.ConnectionIamMemberState;
import com.pulumi.gcp.bigquery.outputs.ConnectionIamMemberCondition;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for BigQuery Connection Connection. Each of these resources serves a different use case:
 * 
 * * `gcp.bigquery.ConnectionIamPolicy`: Authoritative. Sets the IAM policy for the connection and replaces any existing policy already attached.
 * * `gcp.bigquery.ConnectionIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the connection are preserved.
 * * `gcp.bigquery.ConnectionIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the connection are preserved.
 * 
 * &gt; **Note:** `gcp.bigquery.ConnectionIamPolicy` **cannot** be used in conjunction with `gcp.bigquery.ConnectionIamBinding` and `gcp.bigquery.ConnectionIamMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.bigquery.ConnectionIamBinding` resources **can be** used in conjunction with `gcp.bigquery.ConnectionIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_bigquery\_connection\_iam\_policy
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetIAMPolicyArgs;
 * import com.pulumi.gcp.bigquery.ConnectionIamPolicy;
 * import com.pulumi.gcp.bigquery.ConnectionIamPolicyArgs;
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
 *         var policy = new ConnectionIamPolicy(&#34;policy&#34;, ConnectionIamPolicyArgs.builder()        
 *             .project(google_bigquery_connection.connection().project())
 *             .location(google_bigquery_connection.connection().location())
 *             .connectionId(google_bigquery_connection.connection().connection_id())
 *             .policyData(admin.applyValue(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_bigquery\_connection\_iam\_binding
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigquery.ConnectionIamBinding;
 * import com.pulumi.gcp.bigquery.ConnectionIamBindingArgs;
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
 *         var binding = new ConnectionIamBinding(&#34;binding&#34;, ConnectionIamBindingArgs.builder()        
 *             .project(google_bigquery_connection.connection().project())
 *             .location(google_bigquery_connection.connection().location())
 *             .connectionId(google_bigquery_connection.connection().connection_id())
 *             .role(&#34;roles/viewer&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_bigquery\_connection\_iam\_member
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigquery.ConnectionIamMember;
 * import com.pulumi.gcp.bigquery.ConnectionIamMemberArgs;
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
 *         var member = new ConnectionIamMember(&#34;member&#34;, ConnectionIamMemberArgs.builder()        
 *             .project(google_bigquery_connection.connection().project())
 *             .location(google_bigquery_connection.connection().location())
 *             .connectionId(google_bigquery_connection.connection().connection_id())
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
 * For all import syntaxes, the &#34;resource in question&#34; can take any of the following forms* projects/{{project}}/locations/{{location}}/connections/{{connection_id}} * {{project}}/{{location}}/{{connection_id}} * {{location}}/{{connection_id}} * {{connection_id}} Any variables not passed in the import command will be taken from the provider configuration. BigQuery Connection connection IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:bigquery/connectionIamMember:ConnectionIamMember editor &#34;projects/{{project}}/locations/{{location}}/connections/{{connection_id}} roles/viewer user:jane@example.com&#34;
 * ```
 * 
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:bigquery/connectionIamMember:ConnectionIamMember editor &#34;projects/{{project}}/locations/{{location}}/connections/{{connection_id}} roles/viewer&#34;
 * ```
 * 
 *  IAM policy imports use the identifier of the resource in question, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:bigquery/connectionIamMember:ConnectionIamMember editor projects/{{project}}/locations/{{location}}/connections/{{connection_id}}
 * ```
 * 
 *  -&gt; **Custom Roles**If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:bigquery/connectionIamMember:ConnectionIamMember")
public class ConnectionIamMember extends com.pulumi.resources.CustomResource {
    @Export(name="condition", type=ConnectionIamMemberCondition.class, parameters={})
    private Output</* @Nullable */ ConnectionIamMemberCondition> condition;

    public Output<Optional<ConnectionIamMemberCondition>> condition() {
        return Codegen.optional(this.condition);
    }
    /**
     * Optional connection id that should be assigned to the created connection.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="connectionId", type=String.class, parameters={})
    private Output<String> connectionId;

    /**
     * @return Optional connection id that should be assigned to the created connection.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> connectionId() {
        return this.connectionId;
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
     * The geographic location where the connection should reside.
     * Cloud SQL instance must be in the same location as the connection
     * with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
     * Examples: US, EU, asia-northeast1, us-central1, europe-west1.
     * Spanner Connections same as spanner region
     * AWS allowed regions are aws-us-east-1
     * Azure allowed regions are azure-eastus2 Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="location", type=String.class, parameters={})
    private Output<String> location;

    /**
     * @return The geographic location where the connection should reside.
     * Cloud SQL instance must be in the same location as the connection
     * with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
     * Examples: US, EU, asia-northeast1, us-central1, europe-west1.
     * Spanner Connections same as spanner region
     * AWS allowed regions are aws-us-east-1
     * Azure allowed regions are azure-eastus2 Used to find the parent resource to bind the IAM policy to
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
     * `gcp.bigquery.ConnectionIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Export(name="role", type=String.class, parameters={})
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.bigquery.ConnectionIamBinding` can be used per role. Note that custom roles must be of the format
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
    public ConnectionIamMember(String name) {
        this(name, ConnectionIamMemberArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ConnectionIamMember(String name, ConnectionIamMemberArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ConnectionIamMember(String name, ConnectionIamMemberArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigquery/connectionIamMember:ConnectionIamMember", name, args == null ? ConnectionIamMemberArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ConnectionIamMember(String name, Output<String> id, @Nullable ConnectionIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigquery/connectionIamMember:ConnectionIamMember", name, state, makeResourceOptions(options, id));
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
    public static ConnectionIamMember get(String name, Output<String> id, @Nullable ConnectionIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ConnectionIamMember(name, id, state, options);
    }
}
