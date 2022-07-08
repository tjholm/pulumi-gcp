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
 * Firestore enabled. If you haven&#39;t already enabled it, you can create a
 * `gcp.appengine.Application` resource with `database_type` set to
 * `&#34;CLOUD_FIRESTORE&#34;` to do so. Your Firestore location will be the same as
 * the App Engine location specified.
 * 
 * ## Example Usage
 * ### Firestore Document Basic
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
 *         var mydoc = new Document(&#34;mydoc&#34;, DocumentArgs.builder()        
 *             .collection(&#34;somenewcollection&#34;)
 *             .documentId(&#34;my-doc&#34;)
 *             .fields(&#34;{\&#34;something\&#34;:{\&#34;mapValue\&#34;:{\&#34;fields\&#34;:{\&#34;akey\&#34;:{\&#34;stringValue\&#34;:\&#34;avalue\&#34;}}}}}&#34;)
 *             .project(&#34;my-project-name&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Firestore Document Nested Document
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
 *         var mydoc = new Document(&#34;mydoc&#34;, DocumentArgs.builder()        
 *             .collection(&#34;somenewcollection&#34;)
 *             .documentId(&#34;my-doc&#34;)
 *             .fields(&#34;{\&#34;something\&#34;:{\&#34;mapValue\&#34;:{\&#34;fields\&#34;:{\&#34;akey\&#34;:{\&#34;stringValue\&#34;:\&#34;avalue\&#34;}}}}}&#34;)
 *             .project(&#34;my-project-name&#34;)
 *             .build());
 * 
 *         var subDocument = new Document(&#34;subDocument&#34;, DocumentArgs.builder()        
 *             .collection(mydoc.path().apply(path -&gt; String.format(&#34;%s/subdocs&#34;, path)))
 *             .documentId(&#34;bitcoinkey&#34;)
 *             .fields(&#34;{\&#34;something\&#34;:{\&#34;mapValue\&#34;:{\&#34;fields\&#34;:{\&#34;ayo\&#34;:{\&#34;stringValue\&#34;:\&#34;val2\&#34;}}}}}&#34;)
 *             .project(&#34;my-project-name&#34;)
 *             .build());
 * 
 *         var subSubDocument = new Document(&#34;subSubDocument&#34;, DocumentArgs.builder()        
 *             .collection(subDocument.path().apply(path -&gt; String.format(&#34;%s/subsubdocs&#34;, path)))
 *             .documentId(&#34;asecret&#34;)
 *             .fields(&#34;{\&#34;something\&#34;:{\&#34;mapValue\&#34;:{\&#34;fields\&#34;:{\&#34;secret\&#34;:{\&#34;stringValue\&#34;:\&#34;hithere\&#34;}}}}}&#34;)
 *             .project(&#34;my-project-name&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Document can be imported using any of these accepted formats
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
    @Export(name="collection", type=String.class, parameters={})
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
    @Export(name="createTime", type=String.class, parameters={})
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
    @Export(name="database", type=String.class, parameters={})
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
     */
    @Export(name="documentId", type=String.class, parameters={})
    private Output<String> documentId;

    /**
     * @return The client-assigned document ID to use for this document during creation.
     * 
     */
    public Output<String> documentId() {
        return this.documentId;
    }
    /**
     * The document&#39;s [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
     * 
     */
    @Export(name="fields", type=String.class, parameters={})
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
     * &#39;projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}&#39;
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return A server defined name for this index. Format:
     * &#39;projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}&#39;
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A relative path to the collection this document exists within
     * 
     */
    @Export(name="path", type=String.class, parameters={})
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
     * Last update timestamp in RFC3339 format.
     * 
     */
    @Export(name="updateTime", type=String.class, parameters={})
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
