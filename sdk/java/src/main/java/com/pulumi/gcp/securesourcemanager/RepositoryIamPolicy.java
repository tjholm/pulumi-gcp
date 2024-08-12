// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.securesourcemanager;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.securesourcemanager.RepositoryIamPolicyArgs;
import com.pulumi.gcp.securesourcemanager.inputs.RepositoryIamPolicyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for Secure Source Manager Repository. Each of these resources serves a different use case:
 * 
 * * `gcp.securesourcemanager.RepositoryIamPolicy`: Authoritative. Sets the IAM policy for the repository and replaces any existing policy already attached.
 * * `gcp.securesourcemanager.RepositoryIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the repository are preserved.
 * * `gcp.securesourcemanager.RepositoryIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the repository are preserved.
 * 
 * A data source can be used to retrieve policy data in advent you do not need creation
 * 
 * * `gcp.securesourcemanager.RepositoryIamPolicy`: Retrieves the IAM policy for the repository
 * 
 * &gt; **Note:** `gcp.securesourcemanager.RepositoryIamPolicy` **cannot** be used in conjunction with `gcp.securesourcemanager.RepositoryIamBinding` and `gcp.securesourcemanager.RepositoryIamMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.securesourcemanager.RepositoryIamBinding` resources **can be** used in conjunction with `gcp.securesourcemanager.RepositoryIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## gcp.securesourcemanager.RepositoryIamPolicy
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetIAMPolicyArgs;
 * import com.pulumi.gcp.securesourcemanager.RepositoryIamPolicy;
 * import com.pulumi.gcp.securesourcemanager.RepositoryIamPolicyArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         final var admin = OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role("roles/securesourcemanager.repoAdmin")
 *                 .members("user:jane}{@literal @}{@code example.com")
 *                 .build())
 *             .build());
 * 
 *         var policy = new RepositoryIamPolicy("policy", RepositoryIamPolicyArgs.builder()
 *             .project(default_.project())
 *             .location(default_.location())
 *             .repositoryId(default_.repositoryId())
 *             .policyData(admin.applyValue(getIAMPolicyResult -> getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## gcp.securesourcemanager.RepositoryIamBinding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.securesourcemanager.RepositoryIamBinding;
 * import com.pulumi.gcp.securesourcemanager.RepositoryIamBindingArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var binding = new RepositoryIamBinding("binding", RepositoryIamBindingArgs.builder()
 *             .project(default_.project())
 *             .location(default_.location())
 *             .repositoryId(default_.repositoryId())
 *             .role("roles/securesourcemanager.repoAdmin")
 *             .members("user:jane}{@literal @}{@code example.com")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## gcp.securesourcemanager.RepositoryIamMember
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.securesourcemanager.RepositoryIamMember;
 * import com.pulumi.gcp.securesourcemanager.RepositoryIamMemberArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var member = new RepositoryIamMember("member", RepositoryIamMemberArgs.builder()
 *             .project(default_.project())
 *             .location(default_.location())
 *             .repositoryId(default_.repositoryId())
 *             .role("roles/securesourcemanager.repoAdmin")
 *             .member("user:jane}{@literal @}{@code example.com")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## This resource supports User Project Overrides.
 * 
 * - 
 * 
 * # IAM policy for Secure Source Manager Repository
 * Three different resources help you manage your IAM policy for Secure Source Manager Repository. Each of these resources serves a different use case:
 * 
 * * `gcp.securesourcemanager.RepositoryIamPolicy`: Authoritative. Sets the IAM policy for the repository and replaces any existing policy already attached.
 * * `gcp.securesourcemanager.RepositoryIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the repository are preserved.
 * * `gcp.securesourcemanager.RepositoryIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the repository are preserved.
 * 
 * A data source can be used to retrieve policy data in advent you do not need creation
 * 
 * * `gcp.securesourcemanager.RepositoryIamPolicy`: Retrieves the IAM policy for the repository
 * 
 * &gt; **Note:** `gcp.securesourcemanager.RepositoryIamPolicy` **cannot** be used in conjunction with `gcp.securesourcemanager.RepositoryIamBinding` and `gcp.securesourcemanager.RepositoryIamMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.securesourcemanager.RepositoryIamBinding` resources **can be** used in conjunction with `gcp.securesourcemanager.RepositoryIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## gcp.securesourcemanager.RepositoryIamPolicy
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetIAMPolicyArgs;
 * import com.pulumi.gcp.securesourcemanager.RepositoryIamPolicy;
 * import com.pulumi.gcp.securesourcemanager.RepositoryIamPolicyArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         final var admin = OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role("roles/securesourcemanager.repoAdmin")
 *                 .members("user:jane}{@literal @}{@code example.com")
 *                 .build())
 *             .build());
 * 
 *         var policy = new RepositoryIamPolicy("policy", RepositoryIamPolicyArgs.builder()
 *             .project(default_.project())
 *             .location(default_.location())
 *             .repositoryId(default_.repositoryId())
 *             .policyData(admin.applyValue(getIAMPolicyResult -> getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## gcp.securesourcemanager.RepositoryIamBinding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.securesourcemanager.RepositoryIamBinding;
 * import com.pulumi.gcp.securesourcemanager.RepositoryIamBindingArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var binding = new RepositoryIamBinding("binding", RepositoryIamBindingArgs.builder()
 *             .project(default_.project())
 *             .location(default_.location())
 *             .repositoryId(default_.repositoryId())
 *             .role("roles/securesourcemanager.repoAdmin")
 *             .members("user:jane}{@literal @}{@code example.com")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## gcp.securesourcemanager.RepositoryIamMember
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.securesourcemanager.RepositoryIamMember;
 * import com.pulumi.gcp.securesourcemanager.RepositoryIamMemberArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var member = new RepositoryIamMember("member", RepositoryIamMemberArgs.builder()
 *             .project(default_.project())
 *             .location(default_.location())
 *             .repositoryId(default_.repositoryId())
 *             .role("roles/securesourcemanager.repoAdmin")
 *             .member("user:jane}{@literal @}{@code example.com")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * For all import syntaxes, the &#34;resource in question&#34; can take any of the following forms:
 * 
 * * projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}
 * 
 * * {{project}}/{{location}}/{{repository_id}}
 * 
 * * {{location}}/{{repository_id}}
 * 
 * * {{repository_id}}
 * 
 * Any variables not passed in the import command will be taken from the provider configuration.
 * 
 * Secure Source Manager repository IAM resources can be imported using the resource identifiers, role, and member.
 * 
 * IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:securesourcemanager/repositoryIamPolicy:RepositoryIamPolicy editor &#34;projects/{{project}}/locations/{{location}}/repositories/{{repository_id}} roles/securesourcemanager.repoAdmin user:jane{@literal @}example.com&#34;
 * ```
 * 
 * IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:securesourcemanager/repositoryIamPolicy:RepositoryIamPolicy editor &#34;projects/{{project}}/locations/{{location}}/repositories/{{repository_id}} roles/securesourcemanager.repoAdmin&#34;
 * ```
 * 
 * IAM policy imports use the identifier of the resource in question, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:securesourcemanager/repositoryIamPolicy:RepositoryIamPolicy editor projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}
 * ```
 * 
 * -&gt; **Custom Roles**: If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 *  full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:securesourcemanager/repositoryIamPolicy:RepositoryIamPolicy")
public class RepositoryIamPolicy extends com.pulumi.resources.CustomResource {
    /**
     * (Computed) The etag of the IAM policy.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return (Computed) The etag of the IAM policy.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The location for the Repository.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The location for the Repository.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     * 
     */
    @Export(name="policyData", refs={String.class}, tree="[0]")
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
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
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
     * The ID for the Repository.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="repositoryId", refs={String.class}, tree="[0]")
    private Output<String> repositoryId;

    /**
     * @return The ID for the Repository.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> repositoryId() {
        return this.repositoryId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RepositoryIamPolicy(java.lang.String name) {
        this(name, RepositoryIamPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RepositoryIamPolicy(java.lang.String name, RepositoryIamPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RepositoryIamPolicy(java.lang.String name, RepositoryIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:securesourcemanager/repositoryIamPolicy:RepositoryIamPolicy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RepositoryIamPolicy(java.lang.String name, Output<java.lang.String> id, @Nullable RepositoryIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:securesourcemanager/repositoryIamPolicy:RepositoryIamPolicy", name, state, makeResourceOptions(options, id), false);
    }

    private static RepositoryIamPolicyArgs makeArgs(RepositoryIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RepositoryIamPolicyArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static RepositoryIamPolicy get(java.lang.String name, Output<java.lang.String> id, @Nullable RepositoryIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RepositoryIamPolicy(name, id, state, options);
    }
}
