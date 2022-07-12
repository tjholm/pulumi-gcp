// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.spanner;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.spanner.DatabaseArgs;
import com.pulumi.gcp.spanner.inputs.DatabaseState;
import com.pulumi.gcp.spanner.outputs.DatabaseEncryptionConfig;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A Cloud Spanner Database which is hosted on a Spanner instance.
 * 
 * To get more information about Database, see:
 * 
 * * [API documentation](https://cloud.google.com/spanner/docs/reference/rest/v1/projects.instances.databases)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/spanner/)
 * 
 * &gt; **Warning:** It is strongly recommended to set `lifecycle { prevent_destroy = true }` on databases in order to prevent accidental data loss.
 * 
 * ## Example Usage
 * ### Spanner Database Basic
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
 *         var main = new Instance(&#34;main&#34;, InstanceArgs.builder()        
 *             .config(&#34;regional-europe-west1&#34;)
 *             .displayName(&#34;main-instance&#34;)
 *             .numNodes(1)
 *             .build());
 * 
 *         var database = new Database(&#34;database&#34;, DatabaseArgs.builder()        
 *             .instance(main.name())
 *             .versionRetentionPeriod(&#34;3d&#34;)
 *             .ddls(            
 *                 &#34;CREATE TABLE t1 (t1 INT64 NOT NULL,) PRIMARY KEY(t1)&#34;,
 *                 &#34;CREATE TABLE t2 (t2 INT64 NOT NULL,) PRIMARY KEY(t2)&#34;)
 *             .deletionProtection(false)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Database can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:spanner/database:Database default projects/{{project}}/instances/{{instance}}/databases/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:spanner/database:Database default instances/{{instance}}/databases/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:spanner/database:Database default {{project}}/{{instance}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:spanner/database:Database default {{instance}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:spanner/database:Database")
public class Database extends com.pulumi.resources.CustomResource {
    /**
     * The dialect of the Cloud Spanner Database.
     * If it is not provided, &#34;GOOGLE_STANDARD_SQL&#34; will be used.
     * Possible values are `GOOGLE_STANDARD_SQL` and `POSTGRESQL`.
     * 
     */
    @Export(name="databaseDialect", type=String.class, parameters={})
    private Output<String> databaseDialect;

    /**
     * @return The dialect of the Cloud Spanner Database.
     * If it is not provided, &#34;GOOGLE_STANDARD_SQL&#34; will be used.
     * Possible values are `GOOGLE_STANDARD_SQL` and `POSTGRESQL`.
     * 
     */
    public Output<String> databaseDialect() {
        return this.databaseDialect;
    }
    /**
     * An optional list of DDL statements to run inside the newly created
     * database. Statements can create tables, indexes, etc. These statements
     * execute atomically with the creation of the database: if there is an
     * error in any statement, the database is not created.
     * 
     */
    @Export(name="ddls", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> ddls;

    /**
     * @return An optional list of DDL statements to run inside the newly created
     * database. Statements can create tables, indexes, etc. These statements
     * execute atomically with the creation of the database: if there is an
     * error in any statement, the database is not created.
     * 
     */
    public Output<Optional<List<String>>> ddls() {
        return Codegen.optional(this.ddls);
    }
    /**
     * Whether or not to allow the provider to destroy the instance. Unless this field is set to false
     * in state, a `destroy` or `update` that would delete the instance will fail.
     * 
     */
    @Export(name="deletionProtection", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> deletionProtection;

    /**
     * @return Whether or not to allow the provider to destroy the instance. Unless this field is set to false
     * in state, a `destroy` or `update` that would delete the instance will fail.
     * 
     */
    public Output<Optional<Boolean>> deletionProtection() {
        return Codegen.optional(this.deletionProtection);
    }
    /**
     * Encryption configuration for the database
     * Structure is documented below.
     * 
     */
    @Export(name="encryptionConfig", type=DatabaseEncryptionConfig.class, parameters={})
    private Output</* @Nullable */ DatabaseEncryptionConfig> encryptionConfig;

    /**
     * @return Encryption configuration for the database
     * Structure is documented below.
     * 
     */
    public Output<Optional<DatabaseEncryptionConfig>> encryptionConfig() {
        return Codegen.optional(this.encryptionConfig);
    }
    /**
     * The instance to create the database on.
     * 
     */
    @Export(name="instance", type=String.class, parameters={})
    private Output<String> instance;

    /**
     * @return The instance to create the database on.
     * 
     */
    public Output<String> instance() {
        return this.instance;
    }
    /**
     * A unique identifier for the database, which cannot be changed after
     * the instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return A unique identifier for the database, which cannot be changed after
     * the instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].
     * 
     */
    public Output<String> name() {
        return this.name;
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
     * An explanation of the status of the database.
     * 
     */
    @Export(name="state", type=String.class, parameters={})
    private Output<String> state;

    /**
     * @return An explanation of the status of the database.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The retention period for the database. The retention period must be between 1 hour
     * and 7 days, and can be specified in days, hours, minutes, or seconds. For example,
     * the values 1d, 24h, 1440m, and 86400s are equivalent. Default value is 1h.
     * If this property is used, you must avoid adding new DDL statements to `ddl` that
     * update the database&#39;s version_retention_period.
     * 
     */
    @Export(name="versionRetentionPeriod", type=String.class, parameters={})
    private Output<String> versionRetentionPeriod;

    /**
     * @return The retention period for the database. The retention period must be between 1 hour
     * and 7 days, and can be specified in days, hours, minutes, or seconds. For example,
     * the values 1d, 24h, 1440m, and 86400s are equivalent. Default value is 1h.
     * If this property is used, you must avoid adding new DDL statements to `ddl` that
     * update the database&#39;s version_retention_period.
     * 
     */
    public Output<String> versionRetentionPeriod() {
        return this.versionRetentionPeriod;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Database(String name) {
        this(name, DatabaseArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Database(String name, DatabaseArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Database(String name, DatabaseArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:spanner/database:Database", name, args == null ? DatabaseArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Database(String name, Output<String> id, @Nullable DatabaseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:spanner/database:Database", name, state, makeResourceOptions(options, id));
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
    public static Database get(String name, Output<String> id, @Nullable DatabaseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Database(name, id, state, options);
    }
}
