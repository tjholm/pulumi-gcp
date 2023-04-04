// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.firestore;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.firestore.DatabaseArgs;
import com.pulumi.gcp.firestore.inputs.DatabaseState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * ## Import
 * 
 * Database can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:firestore/database:Database default projects/{{project}}/databases/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:firestore/database:Database default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:firestore/database:Database default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:firestore/database:Database")
public class Database extends com.pulumi.resources.CustomResource {
    /**
     * The App Engine integration mode to use for this database.
     * Possible values are `ENABLED` and `DISABLED`.
     * 
     */
    @Export(name="appEngineIntegrationMode", type=String.class, parameters={})
    private Output<String> appEngineIntegrationMode;

    /**
     * @return The App Engine integration mode to use for this database.
     * Possible values are `ENABLED` and `DISABLED`.
     * 
     */
    public Output<String> appEngineIntegrationMode() {
        return this.appEngineIntegrationMode;
    }
    /**
     * The concurrency control mode to use for this database.
     * Possible values are `OPTIMISTIC`, `PESSIMISTIC`, and `OPTIMISTIC_WITH_ENTITY_GROUPS`.
     * 
     */
    @Export(name="concurrencyMode", type=String.class, parameters={})
    private Output<String> concurrencyMode;

    /**
     * @return The concurrency control mode to use for this database.
     * Possible values are `OPTIMISTIC`, `PESSIMISTIC`, and `OPTIMISTIC_WITH_ENTITY_GROUPS`.
     * 
     */
    public Output<String> concurrencyMode() {
        return this.concurrencyMode;
    }
    /**
     * The timestamp at which this database was created.
     * 
     */
    @Export(name="createTime", type=String.class, parameters={})
    private Output<String> createTime;

    /**
     * @return The timestamp at which this database was created.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * This checksum is computed by the server based on the value of other fields,
     * and may be sent on update and delete requests to ensure the client has an
     * up-to-date value before proceeding.
     * 
     */
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    /**
     * @return This checksum is computed by the server based on the value of other fields,
     * and may be sent on update and delete requests to ensure the client has an
     * up-to-date value before proceeding.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * Output only. The keyPrefix for this database.
     * This keyPrefix is used, in combination with the project id (&#34;~&#34;) to construct the application id
     * that is returned from the Cloud Datastore APIs in Google App Engine first generation runtimes.
     * This value may be empty in which case the appid to use for URL-encoded keys is the project_id (eg: foo instead of v~foo).
     * 
     */
    @Export(name="keyPrefix", type=String.class, parameters={})
    private Output<String> keyPrefix;

    /**
     * @return Output only. The keyPrefix for this database.
     * This keyPrefix is used, in combination with the project id (&#34;~&#34;) to construct the application id
     * that is returned from the Cloud Datastore APIs in Google App Engine first generation runtimes.
     * This value may be empty in which case the appid to use for URL-encoded keys is the project_id (eg: foo instead of v~foo).
     * 
     */
    public Output<String> keyPrefix() {
        return this.keyPrefix;
    }
    /**
     * The location of the database. Available databases are listed at
     * https://cloud.google.com/firestore/docs/locations.
     * 
     */
    @Export(name="locationId", type=String.class, parameters={})
    private Output<String> locationId;

    /**
     * @return The location of the database. Available databases are listed at
     * https://cloud.google.com/firestore/docs/locations.
     * 
     */
    public Output<String> locationId() {
        return this.locationId;
    }
    /**
     * The ID to use for the database, which will become the final
     * component of the database&#39;s resource name. This value should be 4-63
     * characters. Valid characters are /[a-z][0-9]-/ with first character
     * a letter and the last a letter or a number. Must not be
     * UUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/.
     * &#34;(default)&#34; database id is also valid.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The ID to use for the database, which will become the final
     * component of the database&#39;s resource name. This value should be 4-63
     * characters. Valid characters are /[a-z][0-9]-/ with first character
     * a letter and the last a letter or a number. Must not be
     * UUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/.
     * &#34;(default)&#34; database id is also valid.
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
     * The type of the database.
     * See https://cloud.google.com/datastore/docs/firestore-or-datastore
     * for information about how to choose.
     * Possible values are `FIRESTORE_NATIVE` and `DATASTORE_MODE`.
     * 
     */
    @Export(name="type", type=String.class, parameters={})
    private Output<String> type;

    /**
     * @return The type of the database.
     * See https://cloud.google.com/datastore/docs/firestore-or-datastore
     * for information about how to choose.
     * Possible values are `FIRESTORE_NATIVE` and `DATASTORE_MODE`.
     * 
     */
    public Output<String> type() {
        return this.type;
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
        super("gcp:firestore/database:Database", name, args == null ? DatabaseArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Database(String name, Output<String> id, @Nullable DatabaseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:firestore/database:Database", name, state, makeResourceOptions(options, id));
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
