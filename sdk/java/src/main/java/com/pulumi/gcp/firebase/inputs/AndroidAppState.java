// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.firebase.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AndroidAppState extends com.pulumi.resources.ResourceArgs {

    public static final AndroidAppState Empty = new AndroidAppState();

    /**
     * Immutable. The globally unique, Firebase-assigned identifier of the App. This identifier should be treated as an opaque
     * token, as the data format is not specified.
     * 
     */
    @Import(name="appId")
    private @Nullable Output<String> appId;

    /**
     * @return Immutable. The globally unique, Firebase-assigned identifier of the App. This identifier should be treated as an opaque
     * token, as the data format is not specified.
     * 
     */
    public Optional<Output<String>> appId() {
        return Optional.ofNullable(this.appId);
    }

    /**
     * (Optional) Set to &#39;ABANDON&#39; to allow the AndroidApp to be untracked from terraform state rather than deleted upon
     * &#39;terraform destroy&#39;. This is useful because the AndroidApp may be serving traffic. Set to &#39;DELETE&#39; to delete the
     * AndroidApp. Default to &#39;DELETE&#39;.
     * 
     */
    @Import(name="deletionPolicy")
    private @Nullable Output<String> deletionPolicy;

    /**
     * @return (Optional) Set to &#39;ABANDON&#39; to allow the AndroidApp to be untracked from terraform state rather than deleted upon
     * &#39;terraform destroy&#39;. This is useful because the AndroidApp may be serving traffic. Set to &#39;DELETE&#39; to delete the
     * AndroidApp. Default to &#39;DELETE&#39;.
     * 
     */
    public Optional<Output<String>> deletionPolicy() {
        return Optional.ofNullable(this.deletionPolicy);
    }

    /**
     * The user-assigned display name of the App.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return The user-assigned display name of the App.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * The fully qualified resource name of the App, for example: projects/projectId/androidApps/appId
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The fully qualified resource name of the App, for example: projects/projectId/androidApps/appId
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Immutable. The canonical package name of the Android app as would appear in the Google Play
     * Developer Console.
     * 
     */
    @Import(name="packageName")
    private @Nullable Output<String> packageName;

    /**
     * @return Immutable. The canonical package name of the Android app as would appear in the Google Play
     * Developer Console.
     * 
     */
    public Optional<Output<String>> packageName() {
        return Optional.ofNullable(this.packageName);
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

    private AndroidAppState() {}

    private AndroidAppState(AndroidAppState $) {
        this.appId = $.appId;
        this.deletionPolicy = $.deletionPolicy;
        this.displayName = $.displayName;
        this.name = $.name;
        this.packageName = $.packageName;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AndroidAppState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AndroidAppState $;

        public Builder() {
            $ = new AndroidAppState();
        }

        public Builder(AndroidAppState defaults) {
            $ = new AndroidAppState(Objects.requireNonNull(defaults));
        }

        /**
         * @param appId Immutable. The globally unique, Firebase-assigned identifier of the App. This identifier should be treated as an opaque
         * token, as the data format is not specified.
         * 
         * @return builder
         * 
         */
        public Builder appId(@Nullable Output<String> appId) {
            $.appId = appId;
            return this;
        }

        /**
         * @param appId Immutable. The globally unique, Firebase-assigned identifier of the App. This identifier should be treated as an opaque
         * token, as the data format is not specified.
         * 
         * @return builder
         * 
         */
        public Builder appId(String appId) {
            return appId(Output.of(appId));
        }

        /**
         * @param deletionPolicy (Optional) Set to &#39;ABANDON&#39; to allow the AndroidApp to be untracked from terraform state rather than deleted upon
         * &#39;terraform destroy&#39;. This is useful because the AndroidApp may be serving traffic. Set to &#39;DELETE&#39; to delete the
         * AndroidApp. Default to &#39;DELETE&#39;.
         * 
         * @return builder
         * 
         */
        public Builder deletionPolicy(@Nullable Output<String> deletionPolicy) {
            $.deletionPolicy = deletionPolicy;
            return this;
        }

        /**
         * @param deletionPolicy (Optional) Set to &#39;ABANDON&#39; to allow the AndroidApp to be untracked from terraform state rather than deleted upon
         * &#39;terraform destroy&#39;. This is useful because the AndroidApp may be serving traffic. Set to &#39;DELETE&#39; to delete the
         * AndroidApp. Default to &#39;DELETE&#39;.
         * 
         * @return builder
         * 
         */
        public Builder deletionPolicy(String deletionPolicy) {
            return deletionPolicy(Output.of(deletionPolicy));
        }

        /**
         * @param displayName The user-assigned display name of the App.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The user-assigned display name of the App.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param name The fully qualified resource name of the App, for example: projects/projectId/androidApps/appId
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The fully qualified resource name of the App, for example: projects/projectId/androidApps/appId
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param packageName Immutable. The canonical package name of the Android app as would appear in the Google Play
         * Developer Console.
         * 
         * @return builder
         * 
         */
        public Builder packageName(@Nullable Output<String> packageName) {
            $.packageName = packageName;
            return this;
        }

        /**
         * @param packageName Immutable. The canonical package name of the Android app as would appear in the Google Play
         * Developer Console.
         * 
         * @return builder
         * 
         */
        public Builder packageName(String packageName) {
            return packageName(Output.of(packageName));
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

        public AndroidAppState build() {
            return $;
        }
    }

}
