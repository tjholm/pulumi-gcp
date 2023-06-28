// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudrun;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.cloudrun.inputs.ServiceMetadataArgs;
import com.pulumi.gcp.cloudrun.inputs.ServiceTemplateArgs;
import com.pulumi.gcp.cloudrun.inputs.ServiceTrafficArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceArgs Empty = new ServiceArgs();

    /**
     * If set to `true`, the revision name (template.metadata.name) will be omitted and
     * autogenerated by Cloud Run. This cannot be set to `true` while `template.metadata.name`
     * is also set.
     * (For legacy support, if `template.metadata.name` is unset in state while
     * this field is set to false, the revision name will still autogenerate.)
     * 
     */
    @Import(name="autogenerateRevisionName")
    private @Nullable Output<Boolean> autogenerateRevisionName;

    /**
     * @return If set to `true`, the revision name (template.metadata.name) will be omitted and
     * autogenerated by Cloud Run. This cannot be set to `true` while `template.metadata.name`
     * is also set.
     * (For legacy support, if `template.metadata.name` is unset in state while
     * this field is set to false, the revision name will still autogenerate.)
     * 
     */
    public Optional<Output<Boolean>> autogenerateRevisionName() {
        return Optional.ofNullable(this.autogenerateRevisionName);
    }

    /**
     * The location of the cloud run instance. eg us-central1
     * 
     */
    @Import(name="location", required=true)
    private Output<String> location;

    /**
     * @return The location of the cloud run instance. eg us-central1
     * 
     */
    public Output<String> location() {
        return this.location;
    }

    /**
     * Optional metadata for this Revision, including labels and annotations.
     * Name will be generated by the Configuration. To set minimum instances
     * for this revision, use the &#34;autoscaling.knative.dev/minScale&#34; annotation
     * key. To set maximum instances for this revision, use the
     * &#34;autoscaling.knative.dev/maxScale&#34; annotation key. To set Cloud SQL
     * connections for the revision, use the &#34;run.googleapis.com/cloudsql-instances&#34;
     * annotation key.
     * Structure is documented below.
     * 
     * (Optional)
     * Metadata associated with this Service, including name, namespace, labels,
     * and annotations.
     * Structure is documented below.
     * 
     */
    @Import(name="metadata")
    private @Nullable Output<ServiceMetadataArgs> metadata;

    /**
     * @return Optional metadata for this Revision, including labels and annotations.
     * Name will be generated by the Configuration. To set minimum instances
     * for this revision, use the &#34;autoscaling.knative.dev/minScale&#34; annotation
     * key. To set maximum instances for this revision, use the
     * &#34;autoscaling.knative.dev/maxScale&#34; annotation key. To set Cloud SQL
     * connections for the revision, use the &#34;run.googleapis.com/cloudsql-instances&#34;
     * annotation key.
     * Structure is documented below.
     * 
     * (Optional)
     * Metadata associated with this Service, including name, namespace, labels,
     * and annotations.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ServiceMetadataArgs>> metadata() {
        return Optional.ofNullable(this.metadata);
    }

    /**
     * Name must be unique within a Google Cloud project and region.
     * Is required when creating resources. Name is primarily intended
     * for creation idempotence and configuration definition. Cannot be updated.
     * More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name must be unique within a Google Cloud project and region.
     * Is required when creating resources. Name is primarily intended
     * for creation idempotence and configuration definition. Cannot be updated.
     * More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * template holds the latest specification for the Revision to be stamped out. The template references the container image,
     * and may also include labels and annotations that should be attached to the Revision. To correlate a Revision, and/or to
     * force a Revision to be created when the spec doesn&#39;t otherwise change, a nonce label may be provided in the template
     * metadata. For more details, see:
     * https://github.com/knative/serving/blob/main/docs/client-conventions.md#associate-modifications-with-revisions Cloud Run
     * does not currently support referencing a build that is responsible for materializing the container image from source.
     * 
     */
    @Import(name="template")
    private @Nullable Output<ServiceTemplateArgs> template;

    /**
     * @return template holds the latest specification for the Revision to be stamped out. The template references the container image,
     * and may also include labels and annotations that should be attached to the Revision. To correlate a Revision, and/or to
     * force a Revision to be created when the spec doesn&#39;t otherwise change, a nonce label may be provided in the template
     * metadata. For more details, see:
     * https://github.com/knative/serving/blob/main/docs/client-conventions.md#associate-modifications-with-revisions Cloud Run
     * does not currently support referencing a build that is responsible for materializing the container image from source.
     * 
     */
    public Optional<Output<ServiceTemplateArgs>> template() {
        return Optional.ofNullable(this.template);
    }

    /**
     * Traffic specifies how to distribute traffic over a collection of Knative Revisions and Configurations
     * 
     */
    @Import(name="traffics")
    private @Nullable Output<List<ServiceTrafficArgs>> traffics;

    /**
     * @return Traffic specifies how to distribute traffic over a collection of Knative Revisions and Configurations
     * 
     */
    public Optional<Output<List<ServiceTrafficArgs>>> traffics() {
        return Optional.ofNullable(this.traffics);
    }

    private ServiceArgs() {}

    private ServiceArgs(ServiceArgs $) {
        this.autogenerateRevisionName = $.autogenerateRevisionName;
        this.location = $.location;
        this.metadata = $.metadata;
        this.name = $.name;
        this.project = $.project;
        this.template = $.template;
        this.traffics = $.traffics;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceArgs $;

        public Builder() {
            $ = new ServiceArgs();
        }

        public Builder(ServiceArgs defaults) {
            $ = new ServiceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param autogenerateRevisionName If set to `true`, the revision name (template.metadata.name) will be omitted and
         * autogenerated by Cloud Run. This cannot be set to `true` while `template.metadata.name`
         * is also set.
         * (For legacy support, if `template.metadata.name` is unset in state while
         * this field is set to false, the revision name will still autogenerate.)
         * 
         * @return builder
         * 
         */
        public Builder autogenerateRevisionName(@Nullable Output<Boolean> autogenerateRevisionName) {
            $.autogenerateRevisionName = autogenerateRevisionName;
            return this;
        }

        /**
         * @param autogenerateRevisionName If set to `true`, the revision name (template.metadata.name) will be omitted and
         * autogenerated by Cloud Run. This cannot be set to `true` while `template.metadata.name`
         * is also set.
         * (For legacy support, if `template.metadata.name` is unset in state while
         * this field is set to false, the revision name will still autogenerate.)
         * 
         * @return builder
         * 
         */
        public Builder autogenerateRevisionName(Boolean autogenerateRevisionName) {
            return autogenerateRevisionName(Output.of(autogenerateRevisionName));
        }

        /**
         * @param location The location of the cloud run instance. eg us-central1
         * 
         * @return builder
         * 
         */
        public Builder location(Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The location of the cloud run instance. eg us-central1
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param metadata Optional metadata for this Revision, including labels and annotations.
         * Name will be generated by the Configuration. To set minimum instances
         * for this revision, use the &#34;autoscaling.knative.dev/minScale&#34; annotation
         * key. To set maximum instances for this revision, use the
         * &#34;autoscaling.knative.dev/maxScale&#34; annotation key. To set Cloud SQL
         * connections for the revision, use the &#34;run.googleapis.com/cloudsql-instances&#34;
         * annotation key.
         * Structure is documented below.
         * 
         * (Optional)
         * Metadata associated with this Service, including name, namespace, labels,
         * and annotations.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder metadata(@Nullable Output<ServiceMetadataArgs> metadata) {
            $.metadata = metadata;
            return this;
        }

        /**
         * @param metadata Optional metadata for this Revision, including labels and annotations.
         * Name will be generated by the Configuration. To set minimum instances
         * for this revision, use the &#34;autoscaling.knative.dev/minScale&#34; annotation
         * key. To set maximum instances for this revision, use the
         * &#34;autoscaling.knative.dev/maxScale&#34; annotation key. To set Cloud SQL
         * connections for the revision, use the &#34;run.googleapis.com/cloudsql-instances&#34;
         * annotation key.
         * Structure is documented below.
         * 
         * (Optional)
         * Metadata associated with this Service, including name, namespace, labels,
         * and annotations.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder metadata(ServiceMetadataArgs metadata) {
            return metadata(Output.of(metadata));
        }

        /**
         * @param name Name must be unique within a Google Cloud project and region.
         * Is required when creating resources. Name is primarily intended
         * for creation idempotence and configuration definition. Cannot be updated.
         * More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name must be unique within a Google Cloud project and region.
         * Is required when creating resources. Name is primarily intended
         * for creation idempotence and configuration definition. Cannot be updated.
         * More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
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
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param template template holds the latest specification for the Revision to be stamped out. The template references the container image,
         * and may also include labels and annotations that should be attached to the Revision. To correlate a Revision, and/or to
         * force a Revision to be created when the spec doesn&#39;t otherwise change, a nonce label may be provided in the template
         * metadata. For more details, see:
         * https://github.com/knative/serving/blob/main/docs/client-conventions.md#associate-modifications-with-revisions Cloud Run
         * does not currently support referencing a build that is responsible for materializing the container image from source.
         * 
         * @return builder
         * 
         */
        public Builder template(@Nullable Output<ServiceTemplateArgs> template) {
            $.template = template;
            return this;
        }

        /**
         * @param template template holds the latest specification for the Revision to be stamped out. The template references the container image,
         * and may also include labels and annotations that should be attached to the Revision. To correlate a Revision, and/or to
         * force a Revision to be created when the spec doesn&#39;t otherwise change, a nonce label may be provided in the template
         * metadata. For more details, see:
         * https://github.com/knative/serving/blob/main/docs/client-conventions.md#associate-modifications-with-revisions Cloud Run
         * does not currently support referencing a build that is responsible for materializing the container image from source.
         * 
         * @return builder
         * 
         */
        public Builder template(ServiceTemplateArgs template) {
            return template(Output.of(template));
        }

        /**
         * @param traffics Traffic specifies how to distribute traffic over a collection of Knative Revisions and Configurations
         * 
         * @return builder
         * 
         */
        public Builder traffics(@Nullable Output<List<ServiceTrafficArgs>> traffics) {
            $.traffics = traffics;
            return this;
        }

        /**
         * @param traffics Traffic specifies how to distribute traffic over a collection of Knative Revisions and Configurations
         * 
         * @return builder
         * 
         */
        public Builder traffics(List<ServiceTrafficArgs> traffics) {
            return traffics(Output.of(traffics));
        }

        /**
         * @param traffics Traffic specifies how to distribute traffic over a collection of Knative Revisions and Configurations
         * 
         * @return builder
         * 
         */
        public Builder traffics(ServiceTrafficArgs... traffics) {
            return traffics(List.of(traffics));
        }

        public ServiceArgs build() {
            $.location = Objects.requireNonNull($.location, "expected parameter 'location' to be non-null");
            return $;
        }
    }

}
