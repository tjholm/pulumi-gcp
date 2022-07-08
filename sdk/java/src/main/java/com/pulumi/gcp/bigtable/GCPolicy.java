// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigtable;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.bigtable.GCPolicyArgs;
import com.pulumi.gcp.bigtable.inputs.GCPolicyState;
import com.pulumi.gcp.bigtable.outputs.GCPolicyMaxAge;
import com.pulumi.gcp.bigtable.outputs.GCPolicyMaxVersion;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a Google Cloud Bigtable GC Policy inside a family. For more information see
 * [the official documentation](https://cloud.google.com/bigtable/) and
 * [API](https://cloud.google.com/bigtable/docs/go/reference).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var instance = new Instance(&#34;instance&#34;, InstanceArgs.builder()        
 *             .clusters(InstanceClusterArgs.builder()
 *                 .clusterId(&#34;tf-instance-cluster&#34;)
 *                 .numNodes(3)
 *                 .storageType(&#34;HDD&#34;)
 *                 .build())
 *             .build());
 * 
 *         var table = new Table(&#34;table&#34;, TableArgs.builder()        
 *             .instanceName(instance.name())
 *             .columnFamilies(TableColumnFamilyArgs.builder()
 *                 .family(&#34;name&#34;)
 *                 .build())
 *             .build());
 * 
 *         var policy = new GCPolicy(&#34;policy&#34;, GCPolicyArgs.builder()        
 *             .instanceName(instance.name())
 *             .table(table.name())
 *             .columnFamily(&#34;name&#34;)
 *             .maxAge(GCPolicyMaxAgeArgs.builder()
 *                 .duration(&#34;168h&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Multiple conditions is also supported. `UNION` when any of its sub-policies apply (OR). `INTERSECTION` when all its sub-policies apply (AND)
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var policy = new GCPolicy(&#34;policy&#34;, GCPolicyArgs.builder()        
 *             .instanceName(google_bigtable_instance.instance().name())
 *             .table(google_bigtable_table.table().name())
 *             .columnFamily(&#34;name&#34;)
 *             .mode(&#34;UNION&#34;)
 *             .maxAge(GCPolicyMaxAgeArgs.builder()
 *                 .duration(&#34;168h&#34;)
 *                 .build())
 *             .maxVersions(GCPolicyMaxVersionArgs.builder()
 *                 .number(10)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * For complex, nested policies, an optional `gc_rules` field are supported. This field
 * conflicts with `mode`, `max_age` and `max_version`. This field is a serialized JSON
 * string. Example:
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var instance = new Instance(&#34;instance&#34;, InstanceArgs.builder()        
 *             .clusters(InstanceClusterArgs.builder()
 *                 .clusterId(&#34;cid&#34;)
 *                 .zone(&#34;us-central1-b&#34;)
 *                 .build())
 *             .instanceType(&#34;DEVELOPMENT&#34;)
 *             .deletionProtection(false)
 *             .build());
 * 
 *         var table = new Table(&#34;table&#34;, TableArgs.builder()        
 *             .instanceName(instance.id())
 *             .columnFamilies(TableColumnFamilyArgs.builder()
 *                 .family(&#34;cf1&#34;)
 *                 .build())
 *             .build());
 * 
 *         var policy = new GCPolicy(&#34;policy&#34;, GCPolicyArgs.builder()        
 *             .instanceName(instance.id())
 *             .table(table.name())
 *             .columnFamily(&#34;cf1&#34;)
 *             .gcRules(&#34;&#34;&#34;
 * {
 *   &#34;mode&#34;: &#34;union&#34;,
 *   &#34;rules&#34;: [
 *     {
 *       &#34;max_age&#34;: &#34;10h&#34;
 *     },
 *     {
 *       &#34;mode&#34;: &#34;intersection&#34;,
 *       &#34;rules&#34;: [
 *         {
 *           &#34;max_age&#34;: &#34;2h&#34;
 *         },
 *         {
 *           &#34;max_version&#34;: 2
 *         }
 *       ]
 *     }
 *   ]
 * }
 *             &#34;&#34;&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * This is equivalent to running the following `cbt` command:
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * This resource does not support import.
 * 
 */
@ResourceType(type="gcp:bigtable/gCPolicy:GCPolicy")
public class GCPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The name of the column family.
     * 
     */
    @Export(name="columnFamily", type=String.class, parameters={})
    private Output<String> columnFamily;

    /**
     * @return The name of the column family.
     * 
     */
    public Output<String> columnFamily() {
        return this.columnFamily;
    }
    /**
     * Serialized JSON object to represent a more complex GC policy. Conflicts with `mode`, `max_age` and `max_version`. Conflicts with `mode`, `max_age` and `max_version`.
     * 
     */
    @Export(name="gcRules", type=String.class, parameters={})
    private Output</* @Nullable */ String> gcRules;

    /**
     * @return Serialized JSON object to represent a more complex GC policy. Conflicts with `mode`, `max_age` and `max_version`. Conflicts with `mode`, `max_age` and `max_version`.
     * 
     */
    public Output<Optional<String>> gcRules() {
        return Codegen.optional(this.gcRules);
    }
    /**
     * The name of the Bigtable instance.
     * 
     */
    @Export(name="instanceName", type=String.class, parameters={})
    private Output<String> instanceName;

    /**
     * @return The name of the Bigtable instance.
     * 
     */
    public Output<String> instanceName() {
        return this.instanceName;
    }
    /**
     * GC policy that applies to all cells older than the given age.
     * 
     */
    @Export(name="maxAge", type=GCPolicyMaxAge.class, parameters={})
    private Output</* @Nullable */ GCPolicyMaxAge> maxAge;

    /**
     * @return GC policy that applies to all cells older than the given age.
     * 
     */
    public Output<Optional<GCPolicyMaxAge>> maxAge() {
        return Codegen.optional(this.maxAge);
    }
    /**
     * GC policy that applies to all versions of a cell except for the most recent.
     * 
     */
    @Export(name="maxVersions", type=List.class, parameters={GCPolicyMaxVersion.class})
    private Output</* @Nullable */ List<GCPolicyMaxVersion>> maxVersions;

    /**
     * @return GC policy that applies to all versions of a cell except for the most recent.
     * 
     */
    public Output<Optional<List<GCPolicyMaxVersion>>> maxVersions() {
        return Codegen.optional(this.maxVersions);
    }
    /**
     * If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
     * 
     */
    @Export(name="mode", type=String.class, parameters={})
    private Output</* @Nullable */ String> mode;

    /**
     * @return If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
     * 
     */
    public Output<Optional<String>> mode() {
        return Codegen.optional(this.mode);
    }
    /**
     * The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The name of the table.
     * 
     */
    @Export(name="table", type=String.class, parameters={})
    private Output<String> table;

    /**
     * @return The name of the table.
     * 
     */
    public Output<String> table() {
        return this.table;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GCPolicy(String name) {
        this(name, GCPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GCPolicy(String name, GCPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GCPolicy(String name, GCPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigtable/gCPolicy:GCPolicy", name, args == null ? GCPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GCPolicy(String name, Output<String> id, @Nullable GCPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigtable/gCPolicy:GCPolicy", name, state, makeResourceOptions(options, id));
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
    public static GCPolicy get(String name, Output<String> id, @Nullable GCPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GCPolicy(name, id, state, options);
    }
}
