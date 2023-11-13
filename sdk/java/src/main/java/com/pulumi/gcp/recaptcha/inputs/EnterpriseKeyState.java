// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.recaptcha.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.recaptcha.inputs.EnterpriseKeyAndroidSettingsArgs;
import com.pulumi.gcp.recaptcha.inputs.EnterpriseKeyIosSettingsArgs;
import com.pulumi.gcp.recaptcha.inputs.EnterpriseKeyTestingOptionsArgs;
import com.pulumi.gcp.recaptcha.inputs.EnterpriseKeyWebSettingsArgs;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EnterpriseKeyState extends com.pulumi.resources.ResourceArgs {

    public static final EnterpriseKeyState Empty = new EnterpriseKeyState();

    /**
     * Settings for keys that can be used by Android apps.
     * 
     */
    @Import(name="androidSettings")
    private @Nullable Output<EnterpriseKeyAndroidSettingsArgs> androidSettings;

    /**
     * @return Settings for keys that can be used by Android apps.
     * 
     */
    public Optional<Output<EnterpriseKeyAndroidSettingsArgs>> androidSettings() {
        return Optional.ofNullable(this.androidSettings);
    }

    /**
     * The timestamp corresponding to the creation of this Key.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return The timestamp corresponding to the creation of this Key.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * Human-readable display name of this key. Modifiable by user.
     * 
     * ***
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return Human-readable display name of this key. Modifiable by user.
     * 
     * ***
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    @Import(name="effectiveLabels")
    private @Nullable Output<Map<String,Object>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    public Optional<Output<Map<String,Object>>> effectiveLabels() {
        return Optional.ofNullable(this.effectiveLabels);
    }

    /**
     * Settings for keys that can be used by iOS apps.
     * 
     */
    @Import(name="iosSettings")
    private @Nullable Output<EnterpriseKeyIosSettingsArgs> iosSettings;

    /**
     * @return Settings for keys that can be used by iOS apps.
     * 
     */
    public Optional<Output<EnterpriseKeyIosSettingsArgs>> iosSettings() {
        return Optional.ofNullable(this.iosSettings);
    }

    /**
     * See [Creating and managing labels](https://cloud.google.com/recaptcha-enterprise/docs/labels).
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return See [Creating and managing labels](https://cloud.google.com/recaptcha-enterprise/docs/labels).
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The resource name for the Key in the format &#34;projects/{project}/keys/{key}&#34;.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The resource name for the Key in the format &#34;projects/{project}/keys/{key}&#34;.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The project for the resource
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The project for the resource
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * The combination of labels configured directly on the resource and default labels configured on the provider.
     * 
     */
    @Import(name="pulumiLabels")
    private @Nullable Output<Map<String,Object>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource and default labels configured on the provider.
     * 
     */
    public Optional<Output<Map<String,Object>>> pulumiLabels() {
        return Optional.ofNullable(this.pulumiLabels);
    }

    /**
     * Options for user acceptance testing.
     * 
     */
    @Import(name="testingOptions")
    private @Nullable Output<EnterpriseKeyTestingOptionsArgs> testingOptions;

    /**
     * @return Options for user acceptance testing.
     * 
     */
    public Optional<Output<EnterpriseKeyTestingOptionsArgs>> testingOptions() {
        return Optional.ofNullable(this.testingOptions);
    }

    /**
     * Settings for keys that can be used by websites.
     * 
     */
    @Import(name="webSettings")
    private @Nullable Output<EnterpriseKeyWebSettingsArgs> webSettings;

    /**
     * @return Settings for keys that can be used by websites.
     * 
     */
    public Optional<Output<EnterpriseKeyWebSettingsArgs>> webSettings() {
        return Optional.ofNullable(this.webSettings);
    }

    private EnterpriseKeyState() {}

    private EnterpriseKeyState(EnterpriseKeyState $) {
        this.androidSettings = $.androidSettings;
        this.createTime = $.createTime;
        this.displayName = $.displayName;
        this.effectiveLabels = $.effectiveLabels;
        this.iosSettings = $.iosSettings;
        this.labels = $.labels;
        this.name = $.name;
        this.project = $.project;
        this.pulumiLabels = $.pulumiLabels;
        this.testingOptions = $.testingOptions;
        this.webSettings = $.webSettings;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EnterpriseKeyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EnterpriseKeyState $;

        public Builder() {
            $ = new EnterpriseKeyState();
        }

        public Builder(EnterpriseKeyState defaults) {
            $ = new EnterpriseKeyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param androidSettings Settings for keys that can be used by Android apps.
         * 
         * @return builder
         * 
         */
        public Builder androidSettings(@Nullable Output<EnterpriseKeyAndroidSettingsArgs> androidSettings) {
            $.androidSettings = androidSettings;
            return this;
        }

        /**
         * @param androidSettings Settings for keys that can be used by Android apps.
         * 
         * @return builder
         * 
         */
        public Builder androidSettings(EnterpriseKeyAndroidSettingsArgs androidSettings) {
            return androidSettings(Output.of(androidSettings));
        }

        /**
         * @param createTime The timestamp corresponding to the creation of this Key.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime The timestamp corresponding to the creation of this Key.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param displayName Human-readable display name of this key. Modifiable by user.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Human-readable display name of this key. Modifiable by user.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param effectiveLabels All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
         * 
         * @return builder
         * 
         */
        public Builder effectiveLabels(@Nullable Output<Map<String,Object>> effectiveLabels) {
            $.effectiveLabels = effectiveLabels;
            return this;
        }

        /**
         * @param effectiveLabels All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
         * 
         * @return builder
         * 
         */
        public Builder effectiveLabels(Map<String,Object> effectiveLabels) {
            return effectiveLabels(Output.of(effectiveLabels));
        }

        /**
         * @param iosSettings Settings for keys that can be used by iOS apps.
         * 
         * @return builder
         * 
         */
        public Builder iosSettings(@Nullable Output<EnterpriseKeyIosSettingsArgs> iosSettings) {
            $.iosSettings = iosSettings;
            return this;
        }

        /**
         * @param iosSettings Settings for keys that can be used by iOS apps.
         * 
         * @return builder
         * 
         */
        public Builder iosSettings(EnterpriseKeyIosSettingsArgs iosSettings) {
            return iosSettings(Output.of(iosSettings));
        }

        /**
         * @param labels See [Creating and managing labels](https://cloud.google.com/recaptcha-enterprise/docs/labels).
         * 
         * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
         * Please refer to the field `effective_labels` for all of the labels present on the resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels See [Creating and managing labels](https://cloud.google.com/recaptcha-enterprise/docs/labels).
         * 
         * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
         * Please refer to the field `effective_labels` for all of the labels present on the resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param name The resource name for the Key in the format &#34;projects/{project}/keys/{key}&#34;.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The resource name for the Key in the format &#34;projects/{project}/keys/{key}&#34;.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param project The project for the resource
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The project for the resource
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param pulumiLabels The combination of labels configured directly on the resource and default labels configured on the provider.
         * 
         * @return builder
         * 
         */
        public Builder pulumiLabels(@Nullable Output<Map<String,Object>> pulumiLabels) {
            $.pulumiLabels = pulumiLabels;
            return this;
        }

        /**
         * @param pulumiLabels The combination of labels configured directly on the resource and default labels configured on the provider.
         * 
         * @return builder
         * 
         */
        public Builder pulumiLabels(Map<String,Object> pulumiLabels) {
            return pulumiLabels(Output.of(pulumiLabels));
        }

        /**
         * @param testingOptions Options for user acceptance testing.
         * 
         * @return builder
         * 
         */
        public Builder testingOptions(@Nullable Output<EnterpriseKeyTestingOptionsArgs> testingOptions) {
            $.testingOptions = testingOptions;
            return this;
        }

        /**
         * @param testingOptions Options for user acceptance testing.
         * 
         * @return builder
         * 
         */
        public Builder testingOptions(EnterpriseKeyTestingOptionsArgs testingOptions) {
            return testingOptions(Output.of(testingOptions));
        }

        /**
         * @param webSettings Settings for keys that can be used by websites.
         * 
         * @return builder
         * 
         */
        public Builder webSettings(@Nullable Output<EnterpriseKeyWebSettingsArgs> webSettings) {
            $.webSettings = webSettings;
            return this;
        }

        /**
         * @param webSettings Settings for keys that can be used by websites.
         * 
         * @return builder
         * 
         */
        public Builder webSettings(EnterpriseKeyWebSettingsArgs webSettings) {
            return webSettings(Output.of(webSettings));
        }

        public EnterpriseKeyState build() {
            return $;
        }
    }

}
