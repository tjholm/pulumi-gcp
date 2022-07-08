// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.artifactregistry;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.artifactregistry.RepositoryIamMemberArgs;
import com.pulumi.gcp.artifactregistry.inputs.RepositoryIamMemberState;
import com.pulumi.gcp.artifactregistry.outputs.RepositoryIamMemberCondition;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for Artifact Registry Repository. Each of these resources serves a different use case:
 * 
 * * `gcp.artifactregistry.RepositoryIamPolicy`: Authoritative. Sets the IAM policy for the repository and replaces any existing policy already attached.
 * * `gcp.artifactregistry.RepositoryIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the repository are preserved.
 * * `gcp.artifactregistry.RepositoryIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the repository are preserved.
 * 
 * &gt; **Note:** `gcp.artifactregistry.RepositoryIamPolicy` **cannot** be used in conjunction with `gcp.artifactregistry.RepositoryIamBinding` and `gcp.artifactregistry.RepositoryIamMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.artifactregistry.RepositoryIamBinding` resources **can be** used in conjunction with `gcp.artifactregistry.RepositoryIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_artifact\_registry\_repository\_iam\_policy
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var policy = new RepositoryIamPolicy(&#34;policy&#34;, RepositoryIamPolicyArgs.builder()        
 *             .project(google_artifact_registry_repository.my-repo().project())
 *             .location(google_artifact_registry_repository.my-repo().location())
 *             .repository(google_artifact_registry_repository.my-repo().name())
 *             .policyData(admin.apply(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_artifact\_registry\_repository\_iam\_binding
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * import com.pulumi.resources.CustomResourceOptions;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var binding = new RepositoryIamBinding(&#34;binding&#34;, RepositoryIamBindingArgs.builder()        
 *             .project(google_artifact_registry_repository.my-repo().project())
 *             .location(google_artifact_registry_repository.my-repo().location())
 *             .repository(google_artifact_registry_repository.my-repo().name())
 *             .role(&#34;roles/viewer&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_artifact\_registry\_repository\_iam\_member
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * import com.pulumi.resources.CustomResourceOptions;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var member = new RepositoryIamMember(&#34;member&#34;, RepositoryIamMemberArgs.builder()        
 *             .project(google_artifact_registry_repository.my-repo().project())
 *             .location(google_artifact_registry_repository.my-repo().location())
 *             .repository(google_artifact_registry_repository.my-repo().name())
 *             .role(&#34;roles/viewer&#34;)
 *             .member(&#34;user:jane@example.com&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * For all import syntaxes, the &#34;resource in question&#34; can take any of the following forms* projects/{{project}}/locations/{{location}}/repositories/{{repository}} * {{project}}/{{location}}/{{repository}} * {{location}}/{{repository}} * {{repository}} Any variables not passed in the import command will be taken from the provider configuration. Artifact Registry repository IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:artifactregistry/repositoryIamMember:RepositoryIamMember editor &#34;projects/{{project}}/locations/{{location}}/repositories/{{repository}} roles/viewer user:jane@example.com&#34;
 * ```
 * 
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:artifactregistry/repositoryIamMember:RepositoryIamMember editor &#34;projects/{{project}}/locations/{{location}}/repositories/{{repository}} roles/viewer&#34;
 * ```
 * 
 *  IAM policy imports use the identifier of the resource in question, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:artifactregistry/repositoryIamMember:RepositoryIamMember editor projects/{{project}}/locations/{{location}}/repositories/{{repository}}
 * ```
 * 
 *  -&gt; **Custom Roles**If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:artifactregistry/repositoryIamMember:RepositoryIamMember")
public class RepositoryIamMember extends com.pulumi.resources.CustomResource {
    @Export(name="condition", type=RepositoryIamMemberCondition.class, parameters={})
    private Output</* @Nullable */ RepositoryIamMemberCondition> condition;

    public Output<Optional<RepositoryIamMemberCondition>> condition() {
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
     * The name of the location this repository is located in.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="location", type=String.class, parameters={})
    private Output<String> location;

    /**
     * @return The name of the location this repository is located in.
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
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="repository", type=String.class, parameters={})
    private Output<String> repository;

    /**
     * @return Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }
    /**
     * The role that should be applied. Only one
     * `gcp.artifactregistry.RepositoryIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Export(name="role", type=String.class, parameters={})
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.artifactregistry.RepositoryIamBinding` can be used per role. Note that custom roles must be of the format
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
    public RepositoryIamMember(String name) {
        this(name, RepositoryIamMemberArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RepositoryIamMember(String name, RepositoryIamMemberArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RepositoryIamMember(String name, RepositoryIamMemberArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:artifactregistry/repositoryIamMember:RepositoryIamMember", name, args == null ? RepositoryIamMemberArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RepositoryIamMember(String name, Output<String> id, @Nullable RepositoryIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:artifactregistry/repositoryIamMember:RepositoryIamMember", name, state, makeResourceOptions(options, id));
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
    public static RepositoryIamMember get(String name, Output<String> id, @Nullable RepositoryIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RepositoryIamMember(name, id, state, options);
    }
}
