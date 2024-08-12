// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.essentialcontacts;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.essentialcontacts.DocumentAiWarehouseLocationArgs;
import com.pulumi.gcp.essentialcontacts.inputs.DocumentAiWarehouseLocationState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A location is used to initialize a project.
 * 
 * To get more information about Location, see:
 * 
 * * [API documentation](https://cloud.google.com/document-warehouse/docs/reference/rest/v1/projects.locations)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/document-warehouse/docs/overview)
 * 
 * ## Example Usage
 * 
 * ### Document Ai Warehouse Location
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
 * import com.pulumi.gcp.essentialcontacts.DocumentAiWarehouseLocation;
 * import com.pulumi.gcp.essentialcontacts.DocumentAiWarehouseLocationArgs;
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
 *         final var project = OrganizationsFunctions.getProject();
 * 
 *         var example = new DocumentAiWarehouseLocation("example", DocumentAiWarehouseLocationArgs.builder()
 *             .location("us")
 *             .projectNumber(project.applyValue(getProjectResult -> getProjectResult.number()))
 *             .accessControlMode("ACL_MODE_DOCUMENT_LEVEL_ACCESS_CONTROL_GCI")
 *             .databaseType("DB_INFRA_SPANNER")
 *             .kmsKey("dummy_key")
 *             .documentCreatorDefaultRole("DOCUMENT_ADMIN")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * This resource does not support import.
 * 
 */
@ResourceType(type="gcp:essentialcontacts/documentAiWarehouseLocation:DocumentAiWarehouseLocation")
public class DocumentAiWarehouseLocation extends com.pulumi.resources.CustomResource {
    /**
     * The access control mode for accessing the customer data.
     * Possible values are: `ACL_MODE_DOCUMENT_LEVEL_ACCESS_CONTROL_GCI`, `ACL_MODE_DOCUMENT_LEVEL_ACCESS_CONTROL_BYOID`, `ACL_MODE_UNIVERSAL_ACCESS`.
     * 
     */
    @Export(name="accessControlMode", refs={String.class}, tree="[0]")
    private Output<String> accessControlMode;

    /**
     * @return The access control mode for accessing the customer data.
     * Possible values are: `ACL_MODE_DOCUMENT_LEVEL_ACCESS_CONTROL_GCI`, `ACL_MODE_DOCUMENT_LEVEL_ACCESS_CONTROL_BYOID`, `ACL_MODE_UNIVERSAL_ACCESS`.
     * 
     */
    public Output<String> accessControlMode() {
        return this.accessControlMode;
    }
    /**
     * The type of database used to store customer data.
     * Possible values are: `DB_INFRA_SPANNER`, `DB_CLOUD_SQL_POSTGRES`.
     * 
     */
    @Export(name="databaseType", refs={String.class}, tree="[0]")
    private Output<String> databaseType;

    /**
     * @return The type of database used to store customer data.
     * Possible values are: `DB_INFRA_SPANNER`, `DB_CLOUD_SQL_POSTGRES`.
     * 
     */
    public Output<String> databaseType() {
        return this.databaseType;
    }
    /**
     * The default role for the person who create a document.
     * Possible values are: `DOCUMENT_ADMIN`, `DOCUMENT_EDITOR`, `DOCUMENT_VIEWER`.
     * 
     */
    @Export(name="documentCreatorDefaultRole", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> documentCreatorDefaultRole;

    /**
     * @return The default role for the person who create a document.
     * Possible values are: `DOCUMENT_ADMIN`, `DOCUMENT_EDITOR`, `DOCUMENT_VIEWER`.
     * 
     */
    public Output<Optional<String>> documentCreatorDefaultRole() {
        return Codegen.optional(this.documentCreatorDefaultRole);
    }
    /**
     * The KMS key used for CMEK encryption. It is required that
     * the kms key is in the same region as the endpoint. The
     * same key will be used for all provisioned resources, if
     * encryption is available. If the kmsKey is left empty, no
     * encryption will be enforced.
     * 
     */
    @Export(name="kmsKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsKey;

    /**
     * @return The KMS key used for CMEK encryption. It is required that
     * the kms key is in the same region as the endpoint. The
     * same key will be used for all provisioned resources, if
     * encryption is available. If the kmsKey is left empty, no
     * encryption will be enforced.
     * 
     */
    public Output<Optional<String>> kmsKey() {
        return Codegen.optional(this.kmsKey);
    }
    /**
     * The location in which the instance is to be provisioned. It takes the form projects/{projectNumber}/locations/{location}.
     * 
     * ***
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The location in which the instance is to be provisioned. It takes the form projects/{projectNumber}/locations/{location}.
     * 
     * ***
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The unique identifier of the project.
     * 
     */
    @Export(name="projectNumber", refs={String.class}, tree="[0]")
    private Output<String> projectNumber;

    /**
     * @return The unique identifier of the project.
     * 
     */
    public Output<String> projectNumber() {
        return this.projectNumber;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DocumentAiWarehouseLocation(java.lang.String name) {
        this(name, DocumentAiWarehouseLocationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DocumentAiWarehouseLocation(java.lang.String name, DocumentAiWarehouseLocationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DocumentAiWarehouseLocation(java.lang.String name, DocumentAiWarehouseLocationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:essentialcontacts/documentAiWarehouseLocation:DocumentAiWarehouseLocation", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private DocumentAiWarehouseLocation(java.lang.String name, Output<java.lang.String> id, @Nullable DocumentAiWarehouseLocationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:essentialcontacts/documentAiWarehouseLocation:DocumentAiWarehouseLocation", name, state, makeResourceOptions(options, id), false);
    }

    private static DocumentAiWarehouseLocationArgs makeArgs(DocumentAiWarehouseLocationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DocumentAiWarehouseLocationArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static DocumentAiWarehouseLocation get(java.lang.String name, Output<java.lang.String> id, @Nullable DocumentAiWarehouseLocationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DocumentAiWarehouseLocation(name, id, state, options);
    }
}
