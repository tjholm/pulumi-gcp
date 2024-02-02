// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.firestore;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.firestore.DocumentArgs;
import com.pulumi.gcp.firestore.inputs.DocumentState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * In Cloud Firestore, the unit of storage is the document. A document is a lightweight record
 * that contains fields, which map to values. Each document is identified by a name.
 * 
 * To get more information about Document, see:
 * 
 * * [API documentation](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/firestore/docs/manage-data/add-data)
 * 
 * &gt; **Warning:** This resource creates a Firestore Document on a project that already has
 * a Firestore database. If you haven&#39;t already created it, you may
 * create a `gcp.firestore.Database` resource with `type` set to
 * `&#34;FIRESTORE_NATIVE&#34;` and `location_id` set to your chosen location.
 * If you wish to use App Engine, you may instead create a
 * `gcp.appengine.Application` resource with `database_type` set to
 * `&#34;CLOUD_FIRESTORE&#34;`. Your Firestore location will be the same as
 * the App Engine location specified.
 * Note: The surface does not support configurable database id. Only `(default)`
 * is allowed for the database parameter.
 * 
 * ## Example Usage
 * 
 * ## Import
 * 
 * Document can be imported using any of these accepted formats* `{{name}}` When using the `pulumi import` command, Document can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:firestore/document:Document default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:firestore/document:Document")
public class Document extends com.pulumi.resources.CustomResource {
    /**
     * The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
     * 
     */
    @Export(name="collection", refs={String.class}, tree="[0]")
    private Output<String> collection;

    /**
     * @return The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
     * 
     */
    public Output<String> collection() {
        return this.collection;
    }
    /**
     * Creation timestamp in RFC3339 format.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Creation timestamp in RFC3339 format.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * The Firestore database id. Defaults to `&#34;(default)&#34;`.
     * 
     */
    @Export(name="database", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> database;

    /**
     * @return The Firestore database id. Defaults to `&#34;(default)&#34;`.
     * 
     */
    public Output<Optional<String>> database() {
        return Codegen.optional(this.database);
    }
    /**
     * The client-assigned document ID to use for this document during creation.
     * 
     * ***
     * 
     */
    @Export(name="documentId", refs={String.class}, tree="[0]")
    private Output<String> documentId;

    /**
     * @return The client-assigned document ID to use for this document during creation.
     * 
     * ***
     * 
     */
    public Output<String> documentId() {
        return this.documentId;
    }
    /**
     * The document&#39;s [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
     * 
     */
    @Export(name="fields", refs={String.class}, tree="[0]")
    private Output<String> fields;

    /**
     * @return The document&#39;s [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
     * 
     */
    public Output<String> fields() {
        return this.fields;
    }
    /**
     * A server defined name for this index. Format:
     * `projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}`
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A server defined name for this index. Format:
     * `projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}`
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A relative path to the collection this document exists within
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return A relative path to the collection this document exists within
     * 
     */
    public Output<String> path() {
        return this.path;
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
     * Last update timestamp in RFC3339 format.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Last update timestamp in RFC3339 format.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Document(String name) {
        this(name, DocumentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Document(String name, DocumentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Document(String name, DocumentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:firestore/document:Document", name, args == null ? DocumentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Document(String name, Output<String> id, @Nullable DocumentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:firestore/document:Document", name, state, makeResourceOptions(options, id));
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
    public static Document get(String name, Output<String> id, @Nullable DocumentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Document(name, id, state, options);
    }
}
