// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkehub;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.gkehub.MembershipIamPolicyArgs;
import com.pulumi.gcp.gkehub.inputs.MembershipIamPolicyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for GKEHub Membership. Each of these resources serves a different use case:
 * 
 * * `gcp.gkehub.MembershipIamPolicy`: Authoritative. Sets the IAM policy for the membership and replaces any existing policy already attached.
 * * `gcp.gkehub.MembershipIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the membership are preserved.
 * * `gcp.gkehub.MembershipIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the membership are preserved.
 * 
 * A data source can be used to retrieve policy data in advent you do not need creation
 * 
 * * `gcp.gkehub.MembershipIamPolicy`: Retrieves the IAM policy for the membership
 * 
 * &gt; **Note:** `gcp.gkehub.MembershipIamPolicy` **cannot** be used in conjunction with `gcp.gkehub.MembershipIamBinding` and `gcp.gkehub.MembershipIamMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.gkehub.MembershipIamBinding` resources **can be** used in conjunction with `gcp.gkehub.MembershipIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## gcp.gkehub.MembershipIamPolicy
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
 * import com.pulumi.gcp.gkehub.MembershipIamPolicy;
 * import com.pulumi.gcp.gkehub.MembershipIamPolicyArgs;
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
 *                 .role("roles/viewer")
 *                 .members("user:jane}{@literal @}{@code example.com")
 *                 .build())
 *             .build());
 * 
 *         var policy = new MembershipIamPolicy("policy", MembershipIamPolicyArgs.builder()
 *             .project(membership.project())
 *             .location(membership.location())
 *             .membershipId(membership.membershipId())
 *             .policyData(admin.applyValue(getIAMPolicyResult -> getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## gcp.gkehub.MembershipIamBinding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.gkehub.MembershipIamBinding;
 * import com.pulumi.gcp.gkehub.MembershipIamBindingArgs;
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
 *         var binding = new MembershipIamBinding("binding", MembershipIamBindingArgs.builder()
 *             .project(membership.project())
 *             .location(membership.location())
 *             .membershipId(membership.membershipId())
 *             .role("roles/viewer")
 *             .members("user:jane}{@literal @}{@code example.com")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## gcp.gkehub.MembershipIamMember
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.gkehub.MembershipIamMember;
 * import com.pulumi.gcp.gkehub.MembershipIamMemberArgs;
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
 *         var member = new MembershipIamMember("member", MembershipIamMemberArgs.builder()
 *             .project(membership.project())
 *             .location(membership.location())
 *             .membershipId(membership.membershipId())
 *             .role("roles/viewer")
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
 * # IAM policy for GKEHub Membership
 * Three different resources help you manage your IAM policy for GKEHub Membership. Each of these resources serves a different use case:
 * 
 * * `gcp.gkehub.MembershipIamPolicy`: Authoritative. Sets the IAM policy for the membership and replaces any existing policy already attached.
 * * `gcp.gkehub.MembershipIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the membership are preserved.
 * * `gcp.gkehub.MembershipIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the membership are preserved.
 * 
 * A data source can be used to retrieve policy data in advent you do not need creation
 * 
 * * `gcp.gkehub.MembershipIamPolicy`: Retrieves the IAM policy for the membership
 * 
 * &gt; **Note:** `gcp.gkehub.MembershipIamPolicy` **cannot** be used in conjunction with `gcp.gkehub.MembershipIamBinding` and `gcp.gkehub.MembershipIamMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.gkehub.MembershipIamBinding` resources **can be** used in conjunction with `gcp.gkehub.MembershipIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## gcp.gkehub.MembershipIamPolicy
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
 * import com.pulumi.gcp.gkehub.MembershipIamPolicy;
 * import com.pulumi.gcp.gkehub.MembershipIamPolicyArgs;
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
 *                 .role("roles/viewer")
 *                 .members("user:jane}{@literal @}{@code example.com")
 *                 .build())
 *             .build());
 * 
 *         var policy = new MembershipIamPolicy("policy", MembershipIamPolicyArgs.builder()
 *             .project(membership.project())
 *             .location(membership.location())
 *             .membershipId(membership.membershipId())
 *             .policyData(admin.applyValue(getIAMPolicyResult -> getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## gcp.gkehub.MembershipIamBinding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.gkehub.MembershipIamBinding;
 * import com.pulumi.gcp.gkehub.MembershipIamBindingArgs;
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
 *         var binding = new MembershipIamBinding("binding", MembershipIamBindingArgs.builder()
 *             .project(membership.project())
 *             .location(membership.location())
 *             .membershipId(membership.membershipId())
 *             .role("roles/viewer")
 *             .members("user:jane}{@literal @}{@code example.com")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## gcp.gkehub.MembershipIamMember
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.gkehub.MembershipIamMember;
 * import com.pulumi.gcp.gkehub.MembershipIamMemberArgs;
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
 *         var member = new MembershipIamMember("member", MembershipIamMemberArgs.builder()
 *             .project(membership.project())
 *             .location(membership.location())
 *             .membershipId(membership.membershipId())
 *             .role("roles/viewer")
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
 * * projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}
 * 
 * * {{project}}/{{location}}/{{membership_id}}
 * 
 * * {{location}}/{{membership_id}}
 * 
 * * {{membership_id}}
 * 
 * Any variables not passed in the import command will be taken from the provider configuration.
 * 
 * GKEHub membership IAM resources can be imported using the resource identifiers, role, and member.
 * 
 * IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:gkehub/membershipIamPolicy:MembershipIamPolicy editor &#34;projects/{{project}}/locations/{{location}}/memberships/{{membership_id}} roles/viewer user:jane{@literal @}example.com&#34;
 * ```
 * 
 * IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:gkehub/membershipIamPolicy:MembershipIamPolicy editor &#34;projects/{{project}}/locations/{{location}}/memberships/{{membership_id}} roles/viewer&#34;
 * ```
 * 
 * IAM policy imports use the identifier of the resource in question, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:gkehub/membershipIamPolicy:MembershipIamPolicy editor projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}
 * ```
 * 
 * -&gt; **Custom Roles**: If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 *  full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:gkehub/membershipIamPolicy:MembershipIamPolicy")
public class MembershipIamPolicy extends com.pulumi.resources.CustomResource {
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
     * Location of the membership.
     * The default value is `global`.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return Location of the membership.
     * The default value is `global`.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    @Export(name="membershipId", refs={String.class}, tree="[0]")
    private Output<String> membershipId;

    public Output<String> membershipId() {
        return this.membershipId;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MembershipIamPolicy(java.lang.String name) {
        this(name, MembershipIamPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MembershipIamPolicy(java.lang.String name, MembershipIamPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MembershipIamPolicy(java.lang.String name, MembershipIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:gkehub/membershipIamPolicy:MembershipIamPolicy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private MembershipIamPolicy(java.lang.String name, Output<java.lang.String> id, @Nullable MembershipIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:gkehub/membershipIamPolicy:MembershipIamPolicy", name, state, makeResourceOptions(options, id), false);
    }

    private static MembershipIamPolicyArgs makeArgs(MembershipIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? MembershipIamPolicyArgs.Empty : args;
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
    public static MembershipIamPolicy get(java.lang.String name, Output<java.lang.String> id, @Nullable MembershipIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MembershipIamPolicy(name, id, state, options);
    }
}
