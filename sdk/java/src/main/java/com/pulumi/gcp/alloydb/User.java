// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.alloydb;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.alloydb.UserArgs;
import com.pulumi.gcp.alloydb.inputs.UserState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A database user in an AlloyDB cluster.
 * 
 * To get more information about User, see:
 * 
 * * [API documentation](https://cloud.google.com/alloydb/docs/reference/rest/v1/projects.locations.clusters.users/create)
 * * How-to Guides
 *     * [AlloyDB](https://cloud.google.com/alloydb/docs/)
 * 
 * ## Example Usage
 * ### Alloydb User Builtin
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.alloydb.Cluster;
 * import com.pulumi.gcp.alloydb.ClusterArgs;
 * import com.pulumi.gcp.alloydb.inputs.ClusterInitialUserArgs;
 * import com.pulumi.gcp.compute.GlobalAddress;
 * import com.pulumi.gcp.compute.GlobalAddressArgs;
 * import com.pulumi.gcp.servicenetworking.Connection;
 * import com.pulumi.gcp.servicenetworking.ConnectionArgs;
 * import com.pulumi.gcp.alloydb.Instance;
 * import com.pulumi.gcp.alloydb.InstanceArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
 * import com.pulumi.gcp.alloydb.User;
 * import com.pulumi.gcp.alloydb.UserArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;);
 * 
 *         var defaultCluster = new Cluster(&#34;defaultCluster&#34;, ClusterArgs.builder()        
 *             .clusterId(&#34;alloydb-cluster&#34;)
 *             .location(&#34;us-central1&#34;)
 *             .network(defaultNetwork.id())
 *             .initialUser(ClusterInitialUserArgs.builder()
 *                 .password(&#34;cluster_secret&#34;)
 *                 .build())
 *             .build());
 * 
 *         var privateIpAlloc = new GlobalAddress(&#34;privateIpAlloc&#34;, GlobalAddressArgs.builder()        
 *             .addressType(&#34;INTERNAL&#34;)
 *             .purpose(&#34;VPC_PEERING&#34;)
 *             .prefixLength(16)
 *             .network(defaultNetwork.id())
 *             .build());
 * 
 *         var vpcConnection = new Connection(&#34;vpcConnection&#34;, ConnectionArgs.builder()        
 *             .network(defaultNetwork.id())
 *             .service(&#34;servicenetworking.googleapis.com&#34;)
 *             .reservedPeeringRanges(privateIpAlloc.name())
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .cluster(defaultCluster.name())
 *             .instanceId(&#34;alloydb-instance&#34;)
 *             .instanceType(&#34;PRIMARY&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(vpcConnection)
 *                 .build());
 * 
 *         final var project = OrganizationsFunctions.getProject();
 * 
 *         var user1 = new User(&#34;user1&#34;, UserArgs.builder()        
 *             .cluster(defaultCluster.name())
 *             .userId(&#34;user1&#34;)
 *             .userType(&#34;ALLOYDB_BUILT_IN&#34;)
 *             .password(&#34;user_secret&#34;)
 *             .databaseRoles(&#34;alloydbsuperuser&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(defaultInstance)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Alloydb User Iam
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.alloydb.Cluster;
 * import com.pulumi.gcp.alloydb.ClusterArgs;
 * import com.pulumi.gcp.alloydb.inputs.ClusterInitialUserArgs;
 * import com.pulumi.gcp.compute.GlobalAddress;
 * import com.pulumi.gcp.compute.GlobalAddressArgs;
 * import com.pulumi.gcp.servicenetworking.Connection;
 * import com.pulumi.gcp.servicenetworking.ConnectionArgs;
 * import com.pulumi.gcp.alloydb.Instance;
 * import com.pulumi.gcp.alloydb.InstanceArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
 * import com.pulumi.gcp.alloydb.User;
 * import com.pulumi.gcp.alloydb.UserArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;);
 * 
 *         var defaultCluster = new Cluster(&#34;defaultCluster&#34;, ClusterArgs.builder()        
 *             .clusterId(&#34;alloydb-cluster&#34;)
 *             .location(&#34;us-central1&#34;)
 *             .network(defaultNetwork.id())
 *             .initialUser(ClusterInitialUserArgs.builder()
 *                 .password(&#34;cluster_secret&#34;)
 *                 .build())
 *             .build());
 * 
 *         var privateIpAlloc = new GlobalAddress(&#34;privateIpAlloc&#34;, GlobalAddressArgs.builder()        
 *             .addressType(&#34;INTERNAL&#34;)
 *             .purpose(&#34;VPC_PEERING&#34;)
 *             .prefixLength(16)
 *             .network(defaultNetwork.id())
 *             .build());
 * 
 *         var vpcConnection = new Connection(&#34;vpcConnection&#34;, ConnectionArgs.builder()        
 *             .network(defaultNetwork.id())
 *             .service(&#34;servicenetworking.googleapis.com&#34;)
 *             .reservedPeeringRanges(privateIpAlloc.name())
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .cluster(defaultCluster.name())
 *             .instanceId(&#34;alloydb-instance&#34;)
 *             .instanceType(&#34;PRIMARY&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(vpcConnection)
 *                 .build());
 * 
 *         final var project = OrganizationsFunctions.getProject();
 * 
 *         var user2 = new User(&#34;user2&#34;, UserArgs.builder()        
 *             .cluster(defaultCluster.name())
 *             .userId(&#34;user2@foo.com&#34;)
 *             .userType(&#34;ALLOYDB_IAM_USER&#34;)
 *             .databaseRoles(&#34;alloydbiamuser&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(defaultInstance)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * User can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/users/{{user_id}}` * `{{project}}/{{location}}/{{cluster}}/{{user_id}}` * `{{location}}/{{cluster}}/{{user_id}}` When using the `pulumi import` command, User can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:alloydb/user:User default projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/users/{{user_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:alloydb/user:User default {{project}}/{{location}}/{{cluster}}/{{user_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:alloydb/user:User default {{location}}/{{cluster}}/{{user_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:alloydb/user:User")
public class User extends com.pulumi.resources.CustomResource {
    /**
     * Identifies the alloydb cluster. Must be in the format
     * &#39;projects/{project}/locations/{location}/clusters/{cluster_id}&#39;
     * 
     */
    @Export(name="cluster", refs={String.class}, tree="[0]")
    private Output<String> cluster;

    /**
     * @return Identifies the alloydb cluster. Must be in the format
     * &#39;projects/{project}/locations/{location}/clusters/{cluster_id}&#39;
     * 
     */
    public Output<String> cluster() {
        return this.cluster;
    }
    /**
     * List of database roles this database user has.
     * 
     */
    @Export(name="databaseRoles", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> databaseRoles;

    /**
     * @return List of database roles this database user has.
     * 
     */
    public Output<Optional<List<String>>> databaseRoles() {
        return Codegen.optional(this.databaseRoles);
    }
    /**
     * Name of the resource in the form of projects/{project}/locations/{location}/clusters/{cluster}/users/{user}.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the resource in the form of projects/{project}/locations/{location}/clusters/{cluster}/users/{user}.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Password for this database user.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> password;

    /**
     * @return Password for this database user.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * The database role name of the user.
     * 
     */
    @Export(name="userId", refs={String.class}, tree="[0]")
    private Output<String> userId;

    /**
     * @return The database role name of the user.
     * 
     */
    public Output<String> userId() {
        return this.userId;
    }
    /**
     * The type of this user.
     * Possible values are: `ALLOYDB_BUILT_IN`, `ALLOYDB_IAM_USER`.
     * 
     * ***
     * 
     */
    @Export(name="userType", refs={String.class}, tree="[0]")
    private Output<String> userType;

    /**
     * @return The type of this user.
     * Possible values are: `ALLOYDB_BUILT_IN`, `ALLOYDB_IAM_USER`.
     * 
     * ***
     * 
     */
    public Output<String> userType() {
        return this.userType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public User(String name) {
        this(name, UserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public User(String name, UserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public User(String name, UserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:alloydb/user:User", name, args == null ? UserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private User(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:alloydb/user:User", name, state, makeResourceOptions(options, id));
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
    public static User get(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new User(name, id, state, options);
    }
}
