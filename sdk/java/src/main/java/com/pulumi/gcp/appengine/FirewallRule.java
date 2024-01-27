// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.appengine;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.appengine.FirewallRuleArgs;
import com.pulumi.gcp.appengine.inputs.FirewallRuleState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A single firewall rule that is evaluated against incoming traffic
 * and provides an action to take on matched requests.
 * 
 * To get more information about FirewallRule, see:
 * 
 * * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.firewall.ingressRules)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/appengine/docs/standard/python/creating-firewalls#creating_firewall_rules)
 * 
 * ## Example Usage
 * ### App Engine Firewall Rule Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.Project;
 * import com.pulumi.gcp.organizations.ProjectArgs;
 * import com.pulumi.gcp.appengine.Application;
 * import com.pulumi.gcp.appengine.ApplicationArgs;
 * import com.pulumi.gcp.appengine.FirewallRule;
 * import com.pulumi.gcp.appengine.FirewallRuleArgs;
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
 *         var myProject = new Project(&#34;myProject&#34;, ProjectArgs.builder()        
 *             .orgId(&#34;123456789&#34;)
 *             .billingAccount(&#34;000000-0000000-0000000-000000&#34;)
 *             .build());
 * 
 *         var app = new Application(&#34;app&#34;, ApplicationArgs.builder()        
 *             .project(myProject.projectId())
 *             .locationId(&#34;us-central&#34;)
 *             .build());
 * 
 *         var rule = new FirewallRule(&#34;rule&#34;, FirewallRuleArgs.builder()        
 *             .project(app.project())
 *             .priority(1000)
 *             .action(&#34;ALLOW&#34;)
 *             .sourceRange(&#34;*&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * FirewallRule can be imported using any of these accepted formats* `apps/{{project}}/firewall/ingressRules/{{priority}}` * `{{project}}/{{priority}}` * `{{priority}}` When using the `pulumi import` command, FirewallRule can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:appengine/firewallRule:FirewallRule default apps/{{project}}/firewall/ingressRules/{{priority}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:appengine/firewallRule:FirewallRule default {{project}}/{{priority}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:appengine/firewallRule:FirewallRule default {{priority}}
 * ```
 * 
 */
@ResourceType(type="gcp:appengine/firewallRule:FirewallRule")
public class FirewallRule extends com.pulumi.resources.CustomResource {
    /**
     * The action to take if this rule matches.
     * Possible values are: `UNSPECIFIED_ACTION`, `ALLOW`, `DENY`.
     * 
     * ***
     * 
     */
    @Export(name="action", refs={String.class}, tree="[0]")
    private Output<String> action;

    /**
     * @return The action to take if this rule matches.
     * Possible values are: `UNSPECIFIED_ACTION`, `ALLOW`, `DENY`.
     * 
     * ***
     * 
     */
    public Output<String> action() {
        return this.action;
    }
    /**
     * An optional string description of this rule.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return An optional string description of this rule.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * A positive integer that defines the order of rule evaluation.
     * Rules with the lowest priority are evaluated first.
     * A default rule at priority Int32.MaxValue matches all IPv4 and
     * IPv6 traffic when no previous rule matches. Only the action of
     * this rule can be modified by the user.
     * 
     */
    @Export(name="priority", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> priority;

    /**
     * @return A positive integer that defines the order of rule evaluation.
     * Rules with the lowest priority are evaluated first.
     * A default rule at priority Int32.MaxValue matches all IPv4 and
     * IPv6 traffic when no previous rule matches. Only the action of
     * this rule can be modified by the user.
     * 
     */
    public Output<Optional<Integer>> priority() {
        return Codegen.optional(this.priority);
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * IP address or range, defined using CIDR notation, of requests that this rule applies to.
     * 
     */
    @Export(name="sourceRange", refs={String.class}, tree="[0]")
    private Output<String> sourceRange;

    /**
     * @return IP address or range, defined using CIDR notation, of requests that this rule applies to.
     * 
     */
    public Output<String> sourceRange() {
        return this.sourceRange;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FirewallRule(String name) {
        this(name, FirewallRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FirewallRule(String name, FirewallRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FirewallRule(String name, FirewallRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:appengine/firewallRule:FirewallRule", name, args == null ? FirewallRuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private FirewallRule(String name, Output<String> id, @Nullable FirewallRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:appengine/firewallRule:FirewallRule", name, state, makeResourceOptions(options, id));
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
    public static FirewallRule get(String name, Output<String> id, @Nullable FirewallRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FirewallRule(name, id, state, options);
    }
}
