// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.FirewallPolicyAssociationArgs;
import com.pulumi.gcp.compute.inputs.FirewallPolicyAssociationState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Allows associating hierarchical firewall policies with the target where they are applied. This allows creating policies and rules in a different location than they are applied.
 * 
 * For more information on applying hierarchical firewall policies see the [official documentation](https://cloud.google.com/vpc/docs/firewall-policies#managing_hierarchical_firewall_policy_resources)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.FirewallPolicy;
 * import com.pulumi.gcp.compute.FirewallPolicyArgs;
 * import com.pulumi.gcp.compute.FirewallPolicyAssociation;
 * import com.pulumi.gcp.compute.FirewallPolicyAssociationArgs;
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
 *         var defaultFirewallPolicy = new FirewallPolicy(&#34;defaultFirewallPolicy&#34;, FirewallPolicyArgs.builder()        
 *             .parent(&#34;organizations/12345&#34;)
 *             .shortName(&#34;my-policy&#34;)
 *             .description(&#34;Example Resource&#34;)
 *             .build());
 * 
 *         var defaultFirewallPolicyAssociation = new FirewallPolicyAssociation(&#34;defaultFirewallPolicyAssociation&#34;, FirewallPolicyAssociationArgs.builder()        
 *             .firewallPolicy(defaultFirewallPolicy.id())
 *             .attachmentTarget(google_folder.folder().name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * FirewallPolicyAssociation can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation default locations/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation default {{firewall_policy}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation")
public class FirewallPolicyAssociation extends com.pulumi.resources.CustomResource {
    /**
     * The target that the firewall policy is attached to.
     * 
     */
    @Export(name="attachmentTarget", type=String.class, parameters={})
    private Output<String> attachmentTarget;

    /**
     * @return The target that the firewall policy is attached to.
     * 
     */
    public Output<String> attachmentTarget() {
        return this.attachmentTarget;
    }
    /**
     * The firewall policy ID of the association.
     * 
     */
    @Export(name="firewallPolicy", type=String.class, parameters={})
    private Output<String> firewallPolicy;

    /**
     * @return The firewall policy ID of the association.
     * 
     */
    public Output<String> firewallPolicy() {
        return this.firewallPolicy;
    }
    /**
     * The name for an association.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name for an association.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The short name of the firewall policy of the association.
     * 
     */
    @Export(name="shortName", type=String.class, parameters={})
    private Output<String> shortName;

    /**
     * @return The short name of the firewall policy of the association.
     * 
     */
    public Output<String> shortName() {
        return this.shortName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FirewallPolicyAssociation(String name) {
        this(name, FirewallPolicyAssociationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FirewallPolicyAssociation(String name, FirewallPolicyAssociationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FirewallPolicyAssociation(String name, FirewallPolicyAssociationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation", name, args == null ? FirewallPolicyAssociationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private FirewallPolicyAssociation(String name, Output<String> id, @Nullable FirewallPolicyAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation", name, state, makeResourceOptions(options, id));
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
    public static FirewallPolicyAssociation get(String name, Output<String> id, @Nullable FirewallPolicyAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FirewallPolicyAssociation(name, id, state, options);
    }
}
