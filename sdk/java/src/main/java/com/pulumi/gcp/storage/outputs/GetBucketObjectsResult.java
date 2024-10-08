// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.storage.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gcp.storage.outputs.GetBucketObjectsBucketObject;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBucketObjectsResult {
    private String bucket;
    /**
     * @return A list of retrieved objects contained in the provided GCS bucket. Structure is defined below.
     * 
     */
    private List<GetBucketObjectsBucketObject> bucketObjects;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String matchGlob;
    private @Nullable String prefix;

    private GetBucketObjectsResult() {}
    public String bucket() {
        return this.bucket;
    }
    /**
     * @return A list of retrieved objects contained in the provided GCS bucket. Structure is defined below.
     * 
     */
    public List<GetBucketObjectsBucketObject> bucketObjects() {
        return this.bucketObjects;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> matchGlob() {
        return Optional.ofNullable(this.matchGlob);
    }
    public Optional<String> prefix() {
        return Optional.ofNullable(this.prefix);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBucketObjectsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String bucket;
        private List<GetBucketObjectsBucketObject> bucketObjects;
        private String id;
        private @Nullable String matchGlob;
        private @Nullable String prefix;
        public Builder() {}
        public Builder(GetBucketObjectsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bucket = defaults.bucket;
    	      this.bucketObjects = defaults.bucketObjects;
    	      this.id = defaults.id;
    	      this.matchGlob = defaults.matchGlob;
    	      this.prefix = defaults.prefix;
        }

        @CustomType.Setter
        public Builder bucket(String bucket) {
            if (bucket == null) {
              throw new MissingRequiredPropertyException("GetBucketObjectsResult", "bucket");
            }
            this.bucket = bucket;
            return this;
        }
        @CustomType.Setter
        public Builder bucketObjects(List<GetBucketObjectsBucketObject> bucketObjects) {
            if (bucketObjects == null) {
              throw new MissingRequiredPropertyException("GetBucketObjectsResult", "bucketObjects");
            }
            this.bucketObjects = bucketObjects;
            return this;
        }
        public Builder bucketObjects(GetBucketObjectsBucketObject... bucketObjects) {
            return bucketObjects(List.of(bucketObjects));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetBucketObjectsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder matchGlob(@Nullable String matchGlob) {

            this.matchGlob = matchGlob;
            return this;
        }
        @CustomType.Setter
        public Builder prefix(@Nullable String prefix) {

            this.prefix = prefix;
            return this;
        }
        public GetBucketObjectsResult build() {
            final var _resultValue = new GetBucketObjectsResult();
            _resultValue.bucket = bucket;
            _resultValue.bucketObjects = bucketObjects;
            _resultValue.id = id;
            _resultValue.matchGlob = matchGlob;
            _resultValue.prefix = prefix;
            return _resultValue;
        }
    }
}
