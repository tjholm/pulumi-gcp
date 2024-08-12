// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.RegionDiskIamPolicyArgs;
import com.pulumi.gcp.compute.inputs.RegionDiskIamPolicyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for Compute Engine RegionDisk. Each of these resources serves a different use case:
 * 
 * * `gcp.compute.RegionDiskIamPolicy`: Authoritative. Sets the IAM policy for the regiondisk and replaces any existing policy already attached.
 * * `gcp.compute.RegionDiskIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the regiondisk are preserved.
 * * `gcp.compute.RegionDiskIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the regiondisk are preserved.
 * 
 * A data source can be used to retrieve policy data in advent you do not need creation
 * 
 * * `gcp.compute.RegionDiskIamPolicy`: Retrieves the IAM policy for the regiondisk
 * 
 * &gt; **Note:** `gcp.compute.RegionDiskIamPolicy` **cannot** be used in conjunction with `gcp.compute.RegionDiskIamBinding` and `gcp.compute.RegionDiskIamMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.compute.RegionDiskIamBinding` resources **can be** used in conjunction with `gcp.compute.RegionDiskIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## gcp.compute.RegionDiskIamPolicy
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
 * import com.pulumi.gcp.compute.RegionDiskIamPolicy;
 * import com.pulumi.gcp.compute.RegionDiskIamPolicyArgs;
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
 *         var policy = new RegionDiskIamPolicy("policy", RegionDiskIamPolicyArgs.builder()
 *             .project(regiondisk.project())
 *             .region(regiondisk.region())
 *             .name(regiondisk.name())
 *             .policyData(admin.applyValue(getIAMPolicyResult -> getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## gcp.compute.RegionDiskIamBinding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.RegionDiskIamBinding;
 * import com.pulumi.gcp.compute.RegionDiskIamBindingArgs;
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
 *         var binding = new RegionDiskIamBinding("binding", RegionDiskIamBindingArgs.builder()
 *             .project(regiondisk.project())
 *             .region(regiondisk.region())
 *             .name(regiondisk.name())
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
 * ## gcp.compute.RegionDiskIamMember
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.RegionDiskIamMember;
 * import com.pulumi.gcp.compute.RegionDiskIamMemberArgs;
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
 *         var member = new RegionDiskIamMember("member", RegionDiskIamMemberArgs.builder()
 *             .project(regiondisk.project())
 *             .region(regiondisk.region())
 *             .name(regiondisk.name())
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
 * # IAM policy for Compute Engine RegionDisk
 * Three different resources help you manage your IAM policy for Compute Engine RegionDisk. Each of these resources serves a different use case:
 * 
 * * `gcp.compute.RegionDiskIamPolicy`: Authoritative. Sets the IAM policy for the regiondisk and replaces any existing policy already attached.
 * * `gcp.compute.RegionDiskIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the regiondisk are preserved.
 * * `gcp.compute.RegionDiskIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the regiondisk are preserved.
 * 
 * A data source can be used to retrieve policy data in advent you do not need creation
 * 
 * * `gcp.compute.RegionDiskIamPolicy`: Retrieves the IAM policy for the regiondisk
 * 
 * &gt; **Note:** `gcp.compute.RegionDiskIamPolicy` **cannot** be used in conjunction with `gcp.compute.RegionDiskIamBinding` and `gcp.compute.RegionDiskIamMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.compute.RegionDiskIamBinding` resources **can be** used in conjunction with `gcp.compute.RegionDiskIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## gcp.compute.RegionDiskIamPolicy
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
 * import com.pulumi.gcp.compute.RegionDiskIamPolicy;
 * import com.pulumi.gcp.compute.RegionDiskIamPolicyArgs;
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
 *         var policy = new RegionDiskIamPolicy("policy", RegionDiskIamPolicyArgs.builder()
 *             .project(regiondisk.project())
 *             .region(regiondisk.region())
 *             .name(regiondisk.name())
 *             .policyData(admin.applyValue(getIAMPolicyResult -> getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## gcp.compute.RegionDiskIamBinding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.RegionDiskIamBinding;
 * import com.pulumi.gcp.compute.RegionDiskIamBindingArgs;
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
 *         var binding = new RegionDiskIamBinding("binding", RegionDiskIamBindingArgs.builder()
 *             .project(regiondisk.project())
 *             .region(regiondisk.region())
 *             .name(regiondisk.name())
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
 * ## gcp.compute.RegionDiskIamMember
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.RegionDiskIamMember;
 * import com.pulumi.gcp.compute.RegionDiskIamMemberArgs;
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
 *         var member = new RegionDiskIamMember("member", RegionDiskIamMemberArgs.builder()
 *             .project(regiondisk.project())
 *             .region(regiondisk.region())
 *             .name(regiondisk.name())
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
 * * projects/{{project}}/regions/{{region}}/disks/{{name}}
 * 
 * * {{project}}/{{region}}/{{name}}
 * 
 * * {{region}}/{{name}}
 * 
 * * {{name}}
 * 
 * Any variables not passed in the import command will be taken from the provider configuration.
 * 
 * Compute Engine regiondisk IAM resources can be imported using the resource identifiers, role, and member.
 * 
 * IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:compute/regionDiskIamPolicy:RegionDiskIamPolicy editor &#34;projects/{{project}}/regions/{{region}}/disks/{{region_disk}} roles/viewer user:jane{@literal @}example.com&#34;
 * ```
 * 
 * IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:compute/regionDiskIamPolicy:RegionDiskIamPolicy editor &#34;projects/{{project}}/regions/{{region}}/disks/{{region_disk}} roles/viewer&#34;
 * ```
 * 
 * IAM policy imports use the identifier of the resource in question, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:compute/regionDiskIamPolicy:RegionDiskIamPolicy editor projects/{{project}}/regions/{{region}}/disks/{{region_disk}}
 * ```
 * 
 * -&gt; **Custom Roles**: If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 *  full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:compute/regionDiskIamPolicy:RegionDiskIamPolicy")
public class RegionDiskIamPolicy extends com.pulumi.resources.CustomResource {
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
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> name() {
        return this.name;
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
     * A reference to the region where the disk resides. Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
     * region is specified, it is taken from the provider configuration.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return A reference to the region where the disk resides. Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
     * region is specified, it is taken from the provider configuration.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RegionDiskIamPolicy(java.lang.String name) {
        this(name, RegionDiskIamPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RegionDiskIamPolicy(java.lang.String name, RegionDiskIamPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RegionDiskIamPolicy(java.lang.String name, RegionDiskIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/regionDiskIamPolicy:RegionDiskIamPolicy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RegionDiskIamPolicy(java.lang.String name, Output<java.lang.String> id, @Nullable RegionDiskIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/regionDiskIamPolicy:RegionDiskIamPolicy", name, state, makeResourceOptions(options, id), false);
    }

    private static RegionDiskIamPolicyArgs makeArgs(RegionDiskIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RegionDiskIamPolicyArgs.Empty : args;
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
    public static RegionDiskIamPolicy get(java.lang.String name, Output<java.lang.String> id, @Nullable RegionDiskIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RegionDiskIamPolicy(name, id, state, options);
    }
}
