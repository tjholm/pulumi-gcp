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
import com.pulumi.gcp.orgpolicy.outputs.PolicySpec;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * An organization policy gives you programmatic control over your organization&#39;s cloud resources.  Using Organization Policies, you will be able to configure constraints across your entire resource hierarchy.
 * 
 * For more information, see:
 * * [Understanding Org Policy concepts](https://cloud.google.com/resource-manager/docs/organization-policy/overview)
 * * [The resource hierarchy](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy)
 * * [All valid constraints](https://cloud.google.com/resource-manager/docs/organization-policy/org-policy-constraints)
 * ## Example Usage
 * ### Enforce_policy
 * A test of an enforce orgpolicy policy for a project
 * ```java
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
 *         var basic = new Project(&#34;basic&#34;, ProjectArgs.builder()        
 *             .orgId(&#34;123456789&#34;)
 *             .projectId(&#34;id&#34;)
 *             .build());
 * 
 *         var primary = new Policy(&#34;primary&#34;, PolicyArgs.builder()        
 *             .parent(basic.name().applyValue(name -&gt; String.format(&#34;projects/%s&#34;, name)))
 *             .spec(PolicySpecArgs.builder()
 *                 .rules(PolicySpecRuleArgs.builder()
 *                     .enforce(&#34;FALSE&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Folder_policy
 * A test of an orgpolicy policy for a folder
 * ```java
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
 *         var basic = new Folder(&#34;basic&#34;, FolderArgs.builder()        
 *             .parent(&#34;organizations/123456789&#34;)
 *             .displayName(&#34;folder&#34;)
 *             .build());
 * 
 *         var primary = new Policy(&#34;primary&#34;, PolicyArgs.builder()        
 *             .parent(basic.name())
 *             .spec(PolicySpecArgs.builder()
 *                 .inheritFromParent(true)
 *                 .rules(PolicySpecRuleArgs.builder()
 *                     .denyAll(&#34;TRUE&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Organization_policy
 * A test of an orgpolicy policy for an organization
 * ```java
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
 *         var primary = new Policy(&#34;primary&#34;, PolicyArgs.builder()        
 *             .parent(&#34;organizations/123456789&#34;)
 *             .spec(PolicySpecArgs.builder()
 *                 .reset(true)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Project_policy
 * A test of an orgpolicy policy for a project
 * ```java
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
 *         var basic = new Project(&#34;basic&#34;, ProjectArgs.builder()        
 *             .orgId(&#34;123456789&#34;)
 *             .projectId(&#34;id&#34;)
 *             .build());
 * 
 *         var primary = new Policy(&#34;primary&#34;, PolicyArgs.builder()        
 *             .parent(basic.name().applyValue(name -&gt; String.format(&#34;projects/%s&#34;, name)))
 *             .spec(PolicySpecArgs.builder()
 *                 .rules(                
 *                     PolicySpecRuleArgs.builder()
 *                         .condition(PolicySpecRuleConditionArgs.builder()
 *                             .description(&#34;A sample condition for the policy&#34;)
 *                             .expression(&#34;resource.matchLabels(&#39;labelKeys/123&#39;, &#39;labelValues/345&#39;)&#34;)
 *                             .location(&#34;sample-location.log&#34;)
 *                             .title(&#34;sample-condition&#34;)
 *                             .build())
 *                         .values(PolicySpecRuleValuesArgs.builder()
 *                             .allowedValues(&#34;projects/allowed-project&#34;)
 *                             .deniedValues(&#34;projects/denied-project&#34;)
 *                             .build())
 *                         .build(),
 *                     PolicySpecRuleArgs.builder()
 *                         .allowAll(&#34;TRUE&#34;)
 *                         .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Policy can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:orgpolicy/policy:Policy default {{parent}}/policies/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:orgpolicy/policy:Policy")
public class Policy extends com.pulumi.resources.CustomResource {
    /**
     * Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, &#34;projects/123/policies/compute.disableSerialPortAccess&#34;. Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
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
     */
    @Export(name="parent", type=String.class, parameters={})
    private Output<String> parent;

    /**
     * @return The parent of the resource.
     * 
     */
    public Output<String> parent() {
        return this.parent;
    }
    /**
     * Basic information about the Organization Policy.
     * 
     */
    @Export(name="spec", type=PolicySpec.class, parameters={})
    private Output</* @Nullable */ PolicySpec> spec;

    /**
     * @return Basic information about the Organization Policy.
     * 
     */
    public Output<Optional<PolicySpec>> spec() {
        return Codegen.optional(this.spec);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Policy(String name) {
        this(name, PolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Policy(String name, PolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Policy(String name, PolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:orgpolicy/policy:Policy", name, args == null ? PolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Policy(String name, Output<String> id, @Nullable PolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:orgpolicy/policy:Policy", name, state, makeResourceOptions(options, id));
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
    public static Policy get(String name, Output<String> id, @Nullable PolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Policy(name, id, state, options);
    }
}
