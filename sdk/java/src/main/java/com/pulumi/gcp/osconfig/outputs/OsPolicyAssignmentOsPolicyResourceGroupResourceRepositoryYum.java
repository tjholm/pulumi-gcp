// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.osconfig.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class OsPolicyAssignmentOsPolicyResourceGroupResourceRepositoryYum {
    /**
     * @return Required. The location of the repository directory.
     * 
     */
    private final String baseUrl;
    /**
     * @return The display name of the repository.
     * 
     */
    private final @Nullable String displayName;
    /**
     * @return URIs of GPG keys.
     * 
     */
    private final @Nullable List<String> gpgKeys;
    /**
     * @return Required. A one word, unique name for this repository. This is the `repo id` in the zypper config file and also the `display_name` if `display_name` is omitted. This id is also used as the unique identifier when checking for GuestPolicy conflicts.
     * 
     */
    private final String id;

    @CustomType.Constructor
    private OsPolicyAssignmentOsPolicyResourceGroupResourceRepositoryYum(
        @CustomType.Parameter("baseUrl") String baseUrl,
        @CustomType.Parameter("displayName") @Nullable String displayName,
        @CustomType.Parameter("gpgKeys") @Nullable List<String> gpgKeys,
        @CustomType.Parameter("id") String id) {
        this.baseUrl = baseUrl;
        this.displayName = displayName;
        this.gpgKeys = gpgKeys;
        this.id = id;
    }

    /**
     * @return Required. The location of the repository directory.
     * 
     */
    public String baseUrl() {
        return this.baseUrl;
    }
    /**
     * @return The display name of the repository.
     * 
     */
    public Optional<String> displayName() {
        return Optional.ofNullable(this.displayName);
    }
    /**
     * @return URIs of GPG keys.
     * 
     */
    public List<String> gpgKeys() {
        return this.gpgKeys == null ? List.of() : this.gpgKeys;
    }
    /**
     * @return Required. A one word, unique name for this repository. This is the `repo id` in the zypper config file and also the `display_name` if `display_name` is omitted. This id is also used as the unique identifier when checking for GuestPolicy conflicts.
     * 
     */
    public String id() {
        return this.id;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(OsPolicyAssignmentOsPolicyResourceGroupResourceRepositoryYum defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String baseUrl;
        private @Nullable String displayName;
        private @Nullable List<String> gpgKeys;
        private String id;

        public Builder() {
    	      // Empty
        }

        public Builder(OsPolicyAssignmentOsPolicyResourceGroupResourceRepositoryYum defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.baseUrl = defaults.baseUrl;
    	      this.displayName = defaults.displayName;
    	      this.gpgKeys = defaults.gpgKeys;
    	      this.id = defaults.id;
        }

        public Builder baseUrl(String baseUrl) {
            this.baseUrl = Objects.requireNonNull(baseUrl);
            return this;
        }
        public Builder displayName(@Nullable String displayName) {
            this.displayName = displayName;
            return this;
        }
        public Builder gpgKeys(@Nullable List<String> gpgKeys) {
            this.gpgKeys = gpgKeys;
            return this;
        }
        public Builder gpgKeys(String... gpgKeys) {
            return gpgKeys(List.of(gpgKeys));
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }        public OsPolicyAssignmentOsPolicyResourceGroupResourceRepositoryYum build() {
            return new OsPolicyAssignmentOsPolicyResourceGroupResourceRepositoryYum(baseUrl, displayName, gpgKeys, id);
        }
    }
}
