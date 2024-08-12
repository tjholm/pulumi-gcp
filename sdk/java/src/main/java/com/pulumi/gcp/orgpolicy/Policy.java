// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.orgpolicy;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.orgpolicy.PolicyArgs;
import com.pulumi.gcp.orgpolicy.inputs.PolicyState;
import com.pulumi.gcp.orgpolicy.outputs.PolicyDryRunSpec;
import com.pulumi.gcp.orgpolicy.outputs.PolicySpec;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Defines an organization policy which is used to specify constraints for configurations of Google Cloud resources.
 * 
 * To get more information about Policy, see:
 * 
 * * [API documentation](https://cloud.google.com/resource-manager/docs/reference/orgpolicy/rest/v2/organizations.policies)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/resource-manager/docs/organization-policy/creating-managing-custom-constraints)
 *     * [Supported Services](https://cloud.google.com/resource-manager/docs/organization-policy/custom-constraint-supported-services)
 * 
 * ## Example Usage
 * 
 * ### Org Policy Policy Enforce
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.Project;
 * import com.pulumi.gcp.organizations.ProjectArgs;
 * import com.pulumi.gcp.orgpolicy.Policy;
 * import com.pulumi.gcp.orgpolicy.PolicyArgs;
 * import com.pulumi.gcp.orgpolicy.inputs.PolicySpecArgs;
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
 *         var basic = new Project("basic", ProjectArgs.builder()
 *             .projectId("id")
 *             .name("id")
 *             .orgId("123456789")
 *             .build());
 * 
 *         var primary = new Policy("primary", PolicyArgs.builder()
 *             .name(basic.name().applyValue(name -> String.format("projects/%s/policies/iam.disableServiceAccountKeyUpload", name)))
 *             .parent(basic.name().applyValue(name -> String.format("projects/%s", name)))
 *             .spec(PolicySpecArgs.builder()
 *                 .rules(PolicySpecRuleArgs.builder()
 *                     .enforce("FALSE")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Org Policy Policy Folder
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.Folder;
 * import com.pulumi.gcp.organizations.FolderArgs;
 * import com.pulumi.gcp.orgpolicy.Policy;
 * import com.pulumi.gcp.orgpolicy.PolicyArgs;
 * import com.pulumi.gcp.orgpolicy.inputs.PolicySpecArgs;
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
 *         var basic = new Folder("basic", FolderArgs.builder()
 *             .parent("organizations/123456789")
 *             .displayName("folder")
 *             .build());
 * 
 *         var primary = new Policy("primary", PolicyArgs.builder()
 *             .name(basic.name().applyValue(name -> String.format("%s/policies/gcp.resourceLocations", name)))
 *             .parent(basic.name())
 *             .spec(PolicySpecArgs.builder()
 *                 .inheritFromParent(true)
 *                 .rules(PolicySpecRuleArgs.builder()
 *                     .denyAll("TRUE")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Org Policy Policy Organization
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.orgpolicy.Policy;
 * import com.pulumi.gcp.orgpolicy.PolicyArgs;
 * import com.pulumi.gcp.orgpolicy.inputs.PolicySpecArgs;
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
 *         var primary = new Policy("primary", PolicyArgs.builder()
 *             .name("organizations/123456789/policies/gcp.detailedAuditLoggingMode")
 *             .parent("organizations/123456789")
 *             .spec(PolicySpecArgs.builder()
 *                 .reset(true)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Org Policy Policy Project
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.Project;
 * import com.pulumi.gcp.organizations.ProjectArgs;
 * import com.pulumi.gcp.orgpolicy.Policy;
 * import com.pulumi.gcp.orgpolicy.PolicyArgs;
 * import com.pulumi.gcp.orgpolicy.inputs.PolicySpecArgs;
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
 *         var basic = new Project("basic", ProjectArgs.builder()
 *             .projectId("id")
 *             .name("id")
 *             .orgId("123456789")
 *             .build());
 * 
 *         var primary = new Policy("primary", PolicyArgs.builder()
 *             .name(basic.name().applyValue(name -> String.format("projects/%s/policies/gcp.resourceLocations", name)))
 *             .parent(basic.name().applyValue(name -> String.format("projects/%s", name)))
 *             .spec(PolicySpecArgs.builder()
 *                 .rules(                
 *                     PolicySpecRuleArgs.builder()
 *                         .condition(PolicySpecRuleConditionArgs.builder()
 *                             .description("A sample condition for the policy")
 *                             .expression("resource.matchTagId('tagKeys/123', 'tagValues/345')")
 *                             .location("sample-location.log")
 *                             .title("sample-condition")
 *                             .build())
 *                         .values(PolicySpecRuleValuesArgs.builder()
 *                             .allowedValues("projects/allowed-project")
 *                             .deniedValues("projects/denied-project")
 *                             .build())
 *                         .build(),
 *                     PolicySpecRuleArgs.builder()
 *                         .allowAll("TRUE")
 *                         .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Org Policy Policy Dry Run Spec
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.orgpolicy.CustomConstraint;
 * import com.pulumi.gcp.orgpolicy.CustomConstraintArgs;
 * import com.pulumi.gcp.orgpolicy.Policy;
 * import com.pulumi.gcp.orgpolicy.PolicyArgs;
 * import com.pulumi.gcp.orgpolicy.inputs.PolicySpecArgs;
 * import com.pulumi.gcp.orgpolicy.inputs.PolicyDryRunSpecArgs;
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
 *         var constraint = new CustomConstraint("constraint", CustomConstraintArgs.builder()
 *             .name("custom.disableGkeAutoUpgrade_37559")
 *             .parent("organizations/123456789")
 *             .displayName("Disable GKE auto upgrade")
 *             .description("Only allow GKE NodePool resource to be created or updated if AutoUpgrade is not enabled where this custom constraint is enforced.")
 *             .actionType("ALLOW")
 *             .condition("resource.management.autoUpgrade == false")
 *             .methodTypes("CREATE")
 *             .resourceTypes("container.googleapis.com/NodePool")
 *             .build());
 * 
 *         var primary = new Policy("primary", PolicyArgs.builder()
 *             .name(constraint.name().applyValue(name -> String.format("organizations/123456789/policies/%s", name)))
 *             .parent("organizations/123456789")
 *             .spec(PolicySpecArgs.builder()
 *                 .rules(PolicySpecRuleArgs.builder()
 *                     .enforce("FALSE")
 *                     .build())
 *                 .build())
 *             .dryRunSpec(PolicyDryRunSpecArgs.builder()
 *                 .inheritFromParent(false)
 *                 .reset(false)
 *                 .rules(PolicyDryRunSpecRuleArgs.builder()
 *                     .enforce("FALSE")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Policy can be imported using any of these accepted formats:
 * 
 * * `{{parent}}/policies/{{name}}`
 * 
 * When using the `pulumi import` command, Policy can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:orgpolicy/policy:Policy default {{parent}}/policies/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:orgpolicy/policy:Policy")
public class Policy extends com.pulumi.resources.CustomResource {
    /**
     * Dry-run policy. Audit-only policy, can be used to monitor how the policy would have impacted the existing and future resources if it&#39;s enforced.
     * Structure is documented below.
     * 
     */
    @Export(name="dryRunSpec", refs={PolicyDryRunSpec.class}, tree="[0]")
    private Output</* @Nullable */ PolicyDryRunSpec> dryRunSpec;

    /**
     * @return Dry-run policy. Audit-only policy, can be used to monitor how the policy would have impacted the existing and future resources if it&#39;s enforced.
     * Structure is documented below.
     * 
     */
    public Output<Optional<PolicyDryRunSpec>> dryRunSpec() {
        return Codegen.optional(this.dryRunSpec);
    }
    /**
     * Optional. An opaque tag indicating the current state of the policy, used for concurrency control. This &#39;etag&#39; is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return Optional. An opaque tag indicating the current state of the policy, used for concurrency control. This &#39;etag&#39; is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, &#34;projects/123/policies/compute.disableSerialPortAccess&#34;. Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, &#34;projects/123/policies/compute.disableSerialPortAccess&#34;. Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The parent of the resource.
     * 
     * ***
     * 
     */
    @Export(name="parent", refs={String.class}, tree="[0]")
    private Output<String> parent;

    /**
     * @return The parent of the resource.
     * 
     * ***
     * 
     */
    public Output<String> parent() {
        return this.parent;
    }
    /**
     * Basic information about the Organization Policy.
     * Structure is documented below.
     * 
     */
    @Export(name="spec", refs={PolicySpec.class}, tree="[0]")
    private Output</* @Nullable */ PolicySpec> spec;

    /**
     * @return Basic information about the Organization Policy.
     * Structure is documented below.
     * 
     */
    public Output<Optional<PolicySpec>> spec() {
        return Codegen.optional(this.spec);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Policy(java.lang.String name) {
        this(name, PolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Policy(java.lang.String name, PolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Policy(java.lang.String name, PolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:orgpolicy/policy:Policy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Policy(java.lang.String name, Output<java.lang.String> id, @Nullable PolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:orgpolicy/policy:Policy", name, state, makeResourceOptions(options, id), false);
    }

    private static PolicyArgs makeArgs(PolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PolicyArgs.Empty : args;
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
    public static Policy get(java.lang.String name, Output<java.lang.String> id, @Nullable PolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Policy(name, id, state, options);
    }
}
