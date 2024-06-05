// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificateauthority.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCertificateTemplateIamPolicyArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCertificateTemplateIamPolicyArgs Empty = new GetCertificateTemplateIamPolicyArgs();

    /**
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Import(name="certificateTemplate", required=true)
    private Output<String> certificateTemplate;

    /**
     * @return Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> certificateTemplate() {
        return this.certificateTemplate;
    }

    /**
     * The location for the resource Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The location for the resource Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    private GetCertificateTemplateIamPolicyArgs() {}

    private GetCertificateTemplateIamPolicyArgs(GetCertificateTemplateIamPolicyArgs $) {
        this.certificateTemplate = $.certificateTemplate;
        this.location = $.location;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCertificateTemplateIamPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCertificateTemplateIamPolicyArgs $;

        public Builder() {
            $ = new GetCertificateTemplateIamPolicyArgs();
        }

        public Builder(GetCertificateTemplateIamPolicyArgs defaults) {
            $ = new GetCertificateTemplateIamPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param certificateTemplate Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder certificateTemplate(Output<String> certificateTemplate) {
            $.certificateTemplate = certificateTemplate;
            return this;
        }

        /**
         * @param certificateTemplate Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder certificateTemplate(String certificateTemplate) {
            return certificateTemplate(Output.of(certificateTemplate));
        }

        /**
         * @param location The location for the resource Used to find the parent resource to bind the IAM policy to. If not specified,
         * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
         * location is specified, it is taken from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The location for the resource Used to find the parent resource to bind the IAM policy to. If not specified,
         * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
         * location is specified, it is taken from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        public GetCertificateTemplateIamPolicyArgs build() {
            if ($.certificateTemplate == null) {
                throw new MissingRequiredPropertyException("GetCertificateTemplateIamPolicyArgs", "certificateTemplate");
            }
            return $;
        }
    }

}
