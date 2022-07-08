// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.healthcare;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.healthcare.inputs.DicomStoreIamBindingConditionArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DicomStoreIamBindingArgs extends com.pulumi.resources.ResourceArgs {

    public static final DicomStoreIamBindingArgs Empty = new DicomStoreIamBindingArgs();

    @Import(name="condition")
    private @Nullable Output<DicomStoreIamBindingConditionArgs> condition;

    public Optional<Output<DicomStoreIamBindingConditionArgs>> condition() {
        return Optional.ofNullable(this.condition);
    }

    /**
     * The DICOM store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
     * `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider&#39;s
     * project setting will be used as a fallback.
     * 
     */
    @Import(name="dicomStoreId", required=true)
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

    @Import(name="members", required=true)
    private Output<List<String>> members;

    public Output<List<String>> members() {
        return this.members;
    }

    /**
     * The role that should be applied. Only one
     * `gcp.healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Import(name="role", required=true)
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    private DicomStoreIamBindingArgs() {}

    private DicomStoreIamBindingArgs(DicomStoreIamBindingArgs $) {
        this.condition = $.condition;
        this.dicomStoreId = $.dicomStoreId;
        this.members = $.members;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DicomStoreIamBindingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DicomStoreIamBindingArgs $;

        public Builder() {
            $ = new DicomStoreIamBindingArgs();
        }

        public Builder(DicomStoreIamBindingArgs defaults) {
            $ = new DicomStoreIamBindingArgs(Objects.requireNonNull(defaults));
        }

        public Builder condition(@Nullable Output<DicomStoreIamBindingConditionArgs> condition) {
            $.condition = condition;
            return this;
        }

        public Builder condition(DicomStoreIamBindingConditionArgs condition) {
            return condition(Output.of(condition));
        }

        /**
         * @param dicomStoreId The DICOM store ID, in the form
         * `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
         * `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider&#39;s
         * project setting will be used as a fallback.
         * 
         * @return builder
         * 
         */
        public Builder dicomStoreId(Output<String> dicomStoreId) {
            $.dicomStoreId = dicomStoreId;
            return this;
        }

        /**
         * @param dicomStoreId The DICOM store ID, in the form
         * `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
         * `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider&#39;s
         * project setting will be used as a fallback.
         * 
         * @return builder
         * 
         */
        public Builder dicomStoreId(String dicomStoreId) {
            return dicomStoreId(Output.of(dicomStoreId));
        }

        public Builder members(Output<List<String>> members) {
            $.members = members;
            return this;
        }

        public Builder members(List<String> members) {
            return members(Output.of(members));
        }

        public Builder members(String... members) {
            return members(List.of(members));
        }

        /**
         * @param role The role that should be applied. Only one
         * `gcp.healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
         * `[projects|organizations]/{parent-name}/roles/{role-name}`.
         * 
         * @return builder
         * 
         */
        public Builder role(Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The role that should be applied. Only one
         * `gcp.healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
         * `[projects|organizations]/{parent-name}/roles/{role-name}`.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public DicomStoreIamBindingArgs build() {
            $.dicomStoreId = Objects.requireNonNull($.dicomStoreId, "expected parameter 'dicomStoreId' to be non-null");
            $.members = Objects.requireNonNull($.members, "expected parameter 'members' to be non-null");
            $.role = Objects.requireNonNull($.role, "expected parameter 'role' to be non-null");
            return $;
        }
    }

}
