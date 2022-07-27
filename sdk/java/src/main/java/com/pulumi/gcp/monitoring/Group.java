// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.monitoring;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.monitoring.GroupArgs;
import com.pulumi.gcp.monitoring.inputs.GroupState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The description of a dynamic collection of monitored resources. Each group
 * has a filter that is matched against monitored resources and their
 * associated metadata. If a group&#39;s filter matches an available monitored
 * resource, then that resource is a member of that group.
 * 
 * To get more information about Group, see:
 * 
 * * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.groups)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/monitoring/groups/)
 * 
 * ## Example Usage
 * ### Monitoring Group Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.monitoring.Group;
 * import com.pulumi.gcp.monitoring.GroupArgs;
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
 *         var basic = new Group(&#34;basic&#34;, GroupArgs.builder()        
 *             .displayName(&#34;tf-test MonitoringGroup&#34;)
 *             .filter(&#34;resource.metadata.region=\&#34;europe-west2\&#34;&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Monitoring Group Subgroup
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.monitoring.Group;
 * import com.pulumi.gcp.monitoring.GroupArgs;
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
 *         var parent = new Group(&#34;parent&#34;, GroupArgs.builder()        
 *             .displayName(&#34;tf-test MonitoringParentGroup&#34;)
 *             .filter(&#34;resource.metadata.region=\&#34;europe-west2\&#34;&#34;)
 *             .build());
 * 
 *         var subgroup = new Group(&#34;subgroup&#34;, GroupArgs.builder()        
 *             .displayName(&#34;tf-test MonitoringSubGroup&#34;)
 *             .filter(&#34;resource.metadata.region=\&#34;europe-west2\&#34;&#34;)
 *             .parentName(parent.name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Group can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:monitoring/group:Group default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:monitoring/group:Group")
public class Group extends com.pulumi.resources.CustomResource {
    /**
     * A user-assigned name for this group, used only for display
     * purposes.
     * 
     */
    @Export(name="displayName", type=String.class, parameters={})
    private Output<String> displayName;

    /**
     * @return A user-assigned name for this group, used only for display
     * purposes.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * The filter used to determine which monitored resources
     * belong to this group.
     * 
     */
    @Export(name="filter", type=String.class, parameters={})
    private Output<String> filter;

    /**
     * @return The filter used to determine which monitored resources
     * belong to this group.
     * 
     */
    public Output<String> filter() {
        return this.filter;
    }
    /**
     * If true, the members of this group are considered to be a
     * cluster. The system can perform additional analysis on
     * groups that are clusters.
     * 
     */
    @Export(name="isCluster", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> isCluster;

    /**
     * @return If true, the members of this group are considered to be a
     * cluster. The system can perform additional analysis on
     * groups that are clusters.
     * 
     */
    public Output<Optional<Boolean>> isCluster() {
        return Codegen.optional(this.isCluster);
    }
    /**
     * A unique identifier for this group. The format is &#34;projects/{project_id_or_number}/groups/{group_id}&#34;.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return A unique identifier for this group. The format is &#34;projects/{project_id_or_number}/groups/{group_id}&#34;.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The name of the group&#39;s parent, if it has one. The format is
     * &#34;projects/{project_id_or_number}/groups/{group_id}&#34;. For
     * groups with no parent, parentName is the empty string, &#34;&#34;.
     * 
     */
    @Export(name="parentName", type=String.class, parameters={})
    private Output</* @Nullable */ String> parentName;

    /**
     * @return The name of the group&#39;s parent, if it has one. The format is
     * &#34;projects/{project_id_or_number}/groups/{group_id}&#34;. For
     * groups with no parent, parentName is the empty string, &#34;&#34;.
     * 
     */
    public Output<Optional<String>> parentName() {
        return Codegen.optional(this.parentName);
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Group(String name) {
        this(name, GroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Group(String name, GroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Group(String name, GroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:monitoring/group:Group", name, args == null ? GroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Group(String name, Output<String> id, @Nullable GroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:monitoring/group:Group", name, state, makeResourceOptions(options, id));
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
    public static Group get(String name, Output<String> id, @Nullable GroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Group(name, id, state, options);
    }
}
