// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.healthcare;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.healthcare.DicomStoreIamPolicyArgs;
import com.pulumi.gcp.healthcare.inputs.DicomStoreIamPolicyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
 * 
 * This member resource can be imported using the `dicom_store_id`, role, and account e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy dicom_store_iam &#34;your-project-id/location-name/dataset-name/dicom-store-name roles/viewer user:foo@example.com&#34;
 * ```
 * 
 *  IAM binding imports use space-delimited identifiers; the resource in question and the role.
 * 
 * This binding resource can be imported using the `dicom_store_id` and role, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy dicom_store_iam &#34;your-project-id/location-name/dataset-name/dicom-store-name roles/viewer&#34;
 * ```
 * 
 *  IAM policy imports use the identifier of the resource in question.
 * 
 * This policy resource can be imported using the `dicom_store_id`, role, and account e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy dicom_store_iam your-project-id/location-name/dataset-name/dicom-store-name
 * ```
 * 
 */
@ResourceType(type="gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy")
public class DicomStoreIamPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The DICOM store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
     * `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider&#39;s
     * project setting will be used as a fallback.
     * 
     */
    @Export(name="dicomStoreId", type=String.class, parameters={})
    private Output<String> dicomStoreId;

    /**
     * @return The DICOM store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
     * `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider&#39;s
     * project setting will be used as a fallback.
     * 
     */
    public Output<String> dicomStoreId() {
        return this.dicomStoreId;
    }
    /**
     * (Computed) The etag of the DICOM store&#39;s IAM policy.
     * 
     */
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    /**
     * @return (Computed) The etag of the DICOM store&#39;s IAM policy.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     * 
     */
    @Export(name="policyData", type=String.class, parameters={})
    private Output<String> policyData;

    /**
     * @return The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     * 
     */
    public Output<String> policyData() {
        return this.policyData;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DicomStoreIamPolicy(String name) {
        this(name, DicomStoreIamPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DicomStoreIamPolicy(String name, DicomStoreIamPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DicomStoreIamPolicy(String name, DicomStoreIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy", name, args == null ? DicomStoreIamPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DicomStoreIamPolicy(String name, Output<String> id, @Nullable DicomStoreIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy", name, state, makeResourceOptions(options, id));
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
    public static DicomStoreIamPolicy get(String name, Output<String> id, @Nullable DicomStoreIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DicomStoreIamPolicy(name, id, state, options);
    }
}
