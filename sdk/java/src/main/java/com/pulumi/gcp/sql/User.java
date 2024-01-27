// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.sql;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.sql.UserArgs;
import com.pulumi.gcp.sql.inputs.UserState;
import com.pulumi.gcp.sql.outputs.UserPasswordPolicy;
import com.pulumi.gcp.sql.outputs.UserSqlServerUserDetail;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a new Google SQL User on a Google SQL User Instance. For more information, see the [official documentation](https://cloud.google.com/sql/), or the [JSON API](https://cloud.google.com/sql/docs/admin-api/v1beta4/users).
 * 
 * ## Import
 * 
 * SQL users for MySQL databases can be imported using the `project`, `instance`, `host` and `name`, e.g. * `{{project_id}}/{{instance}}/{{host}}/{{name}}` SQL users for PostgreSQL databases can be imported using the `project`, `instance` and `name`, e.g. * `{{project_id}}/{{instance}}/{{name}}` When using the `pulumi import` command, NAME_HERE can be imported using one of the formats above. For exampleMySQL database
 * 
 * ```sh
 *  $ pulumi import gcp:sql/user:User default {{project_id}}/{{instance}}/{{host}}/{{name}}
 * ```
 * 
 *  PostgreSQL database
 * 
 * ```sh
 *  $ pulumi import gcp:sql/user:User default {{project_id}}/{{instance}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:sql/user:User")
public class User extends com.pulumi.resources.CustomResource {
    /**
     * The deletion policy for the user.
     * Setting `ABANDON` allows the resource to be abandoned rather than deleted. This is useful
     * for Postgres, where users cannot be deleted from the API if they have been granted SQL roles.
     * 
     * Possible values are: `ABANDON`.
     * 
     * ***
     * 
     */
    @Export(name="deletionPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> deletionPolicy;

    /**
     * @return The deletion policy for the user.
     * Setting `ABANDON` allows the resource to be abandoned rather than deleted. This is useful
     * for Postgres, where users cannot be deleted from the API if they have been granted SQL roles.
     * 
     * Possible values are: `ABANDON`.
     * 
     * ***
     * 
     */
    public Output<Optional<String>> deletionPolicy() {
        return Codegen.optional(this.deletionPolicy);
    }
    /**
     * The host the user can connect from. This is only supported
     * for BUILT_IN users in MySQL instances. Don&#39;t set this field for PostgreSQL and SQL Server instances.
     * Can be an IP address. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="host", refs={String.class}, tree="[0]")
    private Output<String> host;

    /**
     * @return The host the user can connect from. This is only supported
     * for BUILT_IN users in MySQL instances. Don&#39;t set this field for PostgreSQL and SQL Server instances.
     * Can be an IP address. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> host() {
        return this.host;
    }
    /**
     * The name of the Cloud SQL instance. Changing this
     * forces a new resource to be created.
     * 
     */
    @Export(name="instance", refs={String.class}, tree="[0]")
    private Output<String> instance;

    /**
     * @return The name of the Cloud SQL instance. Changing this
     * forces a new resource to be created.
     * 
     */
    public Output<String> instance() {
        return this.instance;
    }
    /**
     * The name of the user. Changing this forces a new resource
     * to be created.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the user. Changing this forces a new resource
     * to be created.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The password for the user. Can be updated. For Postgres
     * instances this is a Required field, unless type is set to either CLOUD_IAM_USER
     * or CLOUD_IAM_SERVICE_ACCOUNT. Don&#39;t set this field for CLOUD_IAM_USER
     * and CLOUD_IAM_SERVICE_ACCOUNT user types for any Cloud SQL instance.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> password;

    /**
     * @return The password for the user. Can be updated. For Postgres
     * instances this is a Required field, unless type is set to either CLOUD_IAM_USER
     * or CLOUD_IAM_SERVICE_ACCOUNT. Don&#39;t set this field for CLOUD_IAM_USER
     * and CLOUD_IAM_SERVICE_ACCOUNT user types for any Cloud SQL instance.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    @Export(name="passwordPolicy", refs={UserPasswordPolicy.class}, tree="[0]")
    private Output</* @Nullable */ UserPasswordPolicy> passwordPolicy;

    public Output<Optional<UserPasswordPolicy>> passwordPolicy() {
        return Codegen.optional(this.passwordPolicy);
    }
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    @Export(name="sqlServerUserDetails", refs={List.class,UserSqlServerUserDetail.class}, tree="[0,1]")
    private Output<List<UserSqlServerUserDetail>> sqlServerUserDetails;

    public Output<List<UserSqlServerUserDetail>> sqlServerUserDetails() {
        return this.sqlServerUserDetails;
    }
    /**
     * The user type. It determines the method to authenticate the
     * user during login. The default is the database&#39;s built-in user type. Flags
     * include &#34;BUILT_IN&#34;, &#34;CLOUD_IAM_USER&#34;, &#34;CLOUD_IAM_GROUP&#34; or &#34;CLOUD_IAM_SERVICE_ACCOUNT&#34;.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return The user type. It determines the method to authenticate the
     * user during login. The default is the database&#39;s built-in user type. Flags
     * include &#34;BUILT_IN&#34;, &#34;CLOUD_IAM_USER&#34;, &#34;CLOUD_IAM_GROUP&#34; or &#34;CLOUD_IAM_SERVICE_ACCOUNT&#34;.
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
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
        super("gcp:sql/user:User", name, args == null ? UserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private User(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:sql/user:User", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
            ))
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
