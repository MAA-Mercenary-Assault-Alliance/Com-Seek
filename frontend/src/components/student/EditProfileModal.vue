<template>
  <div
    class="fixed inset-0 z-50 bg-black/40 flex items-center justify-center p-4"
    @keydown.esc="$emit('close')"
  >
    <div class="bg-white w-full max-w-2xl rounded-xl shadow-lg p-6">
      <div class="flex items-start justify-between mb-4">
        <h3 class="text-xl font-semibold">Edit Profile</h3>
        <button
          class="p-2 rounded-full hover:bg-gray-100"
          @click="$emit('close')"
          aria-label="Close"
        >
          <span class="material-icons">close</span>
        </button>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <!-- Names + Bio -->
        <div class="space-y-4">
          <label class="block">
            <span class="text-sm font-medium text-gray-700">First Name</span>
            <input
              v-model="editable.firstName"
              type="text"
              class="mt-1 w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent"
              placeholder="Enter first name"
            />
          </label>

          <label class="block">
            <span class="text-sm font-medium text-gray-700">Last Name</span>
            <input
              v-model="editable.lastName"
              type="text"
              class="mt-1 w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent"
              placeholder="Enter last name"
            />
          </label>

          <label class="block">
            <span class="text-sm font-medium text-gray-700"
              >About / Description</span
            >
            <textarea
              v-model="editable.description"
              rows="6"
              class="mt-1 w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent"
              placeholder="Tell people a bit about yourself…"
            ></textarea>
          </label>
        </div>

        <!-- Socials + Image uploads -->
        <div class="space-y-4">
          <div class="grid grid-cols-1 gap-3">
            <label class="block">
              <span class="text-sm font-medium text-gray-700"
                >Facebook URL</span
              >
              <input
                v-model="editable.socials.facebook"
                type="url"
                inputmode="url"
                class="mt-1 w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent"
                placeholder="https://facebook.com/your-profile"
              />
            </label>
            <label class="block">
              <span class="text-sm font-medium text-gray-700"
                >Twitter / X URL</span
              >
              <input
                v-model="editable.socials.twitter"
                type="url"
                inputmode="url"
                class="mt-1 w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent"
                placeholder="https://x.com/your-handle"
              />
            </label>
            <label class="block">
              <span class="text-sm font-medium text-gray-700"
                >Instagram URL</span
              >
              <input
                v-model="editable.socials.instagram"
                type="url"
                inputmode="url"
                class="mt-1 w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent"
                placeholder="https://instagram.com/your-handle"
              />
            </label>
            <label class="block">
              <span class="text-sm font-medium text-gray-700">GitHub URL</span>
              <input
                v-model="editable.socials.github"
                type="url"
                inputmode="url"
                class="mt-1 w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent"
                placeholder="https://github.com/your-username"
              />
            </label>
          </div>

          <!-- Avatar -->
          <div class="border rounded-lg p-3">
            <p class="text-sm font-medium text-gray-700 mb-2">Profile Image</p>
            <div class="flex items-center gap-4">
              <div
                class="w-20 h-20 rounded-full overflow-hidden bg-gray-100 flex items-center justify-center"
              >
                <img
                  v-if="avatarPreview || editable.avatar"
                  :src="avatarPreview || editable.avatar"
                  alt="avatar preview"
                  class="object-cover w-20 h-20"
                />
                <span v-else class="material-icons text-gray-400">person</span>
              </div>
              <div class="flex-1">
                <input
                  ref="avatarFile"
                  type="file"
                  accept="image/*"
                  class="block w-full text-sm text-gray-700 file:mr-4 file:py-2 file:px-4 file:rounded-md file:border-0 file:bg-primary file:text-white hover:file:opacity-90"
                  @change="onAvatarFileChange"
                />
                <div class="flex gap-2 mt-2">
                  <button
                    v-if="avatarPreview"
                    class="btn btn-ghost btn-sm"
                    @click="clearAvatarPreview"
                  >
                    Remove
                  </button>
                </div>
              </div>
            </div>
            <p class="text-xs text-gray-500 mt-2">
              Max ~5MB. Square images look best.
            </p>
          </div>

          <!-- Cover -->
          <div class="border rounded-lg p-3">
            <p class="text-sm font-medium text-gray-700 mb-2">Cover Image</p>
            <div class="flex items-center gap-4">
              <div
                class="w-28 h-14 overflow-hidden rounded-md bg-gray-100 flex items-center justify-center"
              >
                <img
                  v-if="coverPreview || editable.cover"
                  :src="coverPreview || editable.cover"
                  alt="cover preview"
                  class="object-cover w-28 h-14"
                />
                <span v-else class="material-icons text-gray-400">image</span>
              </div>
              <div class="flex-1">
                <input
                  ref="coverFile"
                  type="file"
                  accept="image/*"
                  class="block w-full text-sm text-gray-700 file:mr-4 file:py-2 file:px-4 file:rounded-md file:border-0 file:bg-primary file:text-white hover:file:opacity-90"
                  @change="onCoverFileChange"
                />
                <div class="flex gap-2 mt-2">
                  <button
                    v-if="coverPreview"
                    class="btn btn-ghost btn-sm"
                    @click="clearCoverPreview"
                  >
                    Remove
                  </button>
                </div>
              </div>
            </div>
            <p class="text-xs text-gray-500 mt-2">
              Widescreen (e.g., 1200×300) recommended.
            </p>
          </div>
        </div>
      </div>

      <!-- Footer -->
      <div
        class="mt-6 flex flex-col md:flex-row items-center gap-3 md:justify-end"
      >
        <p v-if="validationError" class="text-sm text-red-600 mr-auto">
          {{ validationError }}
        </p>
        <div class="flex gap-3 w-full md:w-auto">
          <button
            class="btn btn-ghost w-full md:w-auto"
            @click="$emit('close')"
          >
            Cancel
          </button>
          <button class="btn btn-primary w-full md:w-auto" @click="save">
            Save
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "EditProfileModal",
  props: {
    profile: { type: Object, required: true },
  },
  emits: ["close", "save"],
  data() {
    return {
      editable: {
        firstName: "",
        lastName: "",
        description: "",
        socials: { facebook: "", twitter: "", instagram: "", github: "" },
        avatar: "",
        cover: "",
        newAvatarFile: null,
        newCoverFile: null,
      },
      avatarPreview: "",
      coverPreview: "",
      validationError: "",
    };
  },
  created() {
    // preload from props
    this.editable.firstName = this.profile.firstName || "";
    this.editable.lastName = this.profile.lastName || "";
    this.editable.description = this.profile.description || "";
    this.editable.socials = {
      facebook: this.profile.socials?.facebook || "",
      twitter: this.profile.socials?.twitter || "",
      instagram: this.profile.socials?.instagram || "",
      github: this.profile.socials?.github || "",
    };
    this.editable.avatar = this.profile.avatar || "";
    this.editable.cover = this.profile.cover || "";
  },
  beforeUnmount() {
    this.revokePreview("avatarPreview");
    this.revokePreview("coverPreview");
  },
  methods: {
    onAvatarFileChange(e) {
      const file = e.target.files && e.target.files[0];
      if (!file) return;
      if (!file.type.startsWith("image/")) {
        this.validationError = "Profile image must be an image file.";
        return;
      }
      if (file.size > 5 * 1024 * 1024) {
        this.validationError = "Profile image exceeds 5MB.";
        return;
      }
      this.validationError = "";
      this.replacePreview("avatarPreview", file);
      this.editable.newAvatarFile = file;
    },
    onCoverFileChange(e) {
      const file = e.target.files && e.target.files[0];
      if (!file) return;
      if (!file.type.startsWith("image/")) {
        this.validationError = "Cover image must be an image file.";
        return;
      }
      if (file.size > 8 * 1024 * 1024) {
        this.validationError = "Cover image exceeds 8MB.";
        return;
      }
      this.validationError = "";
      this.replacePreview("coverPreview", file);
      this.editable.newCoverFile = file;
    },
    replacePreview(key, file) {
      this.revokePreview(key);
      const url = URL.createObjectURL(file);
      this[key] = url;
    },
    revokePreview(key) {
      if (this[key]) {
        URL.revokeObjectURL(this[key]);
        this[key] = "";
      }
    },
    clearAvatarPreview() {
      this.revokePreview("avatarPreview");
      if (this.$refs.avatarFile) this.$refs.avatarFile.value = "";
      this.editable.newAvatarFile = null;
    },
    clearCoverPreview() {
      this.revokePreview("coverPreview");
      if (this.$refs.coverFile) this.$refs.coverFile.value = "";
      this.editable.newCoverFile = null;
    },
    save() {
      // basic URL validation
      const urlFields = ["facebook", "twitter", "instagram", "github"];
      for (const f of urlFields) {
        const val = (this.editable.socials[f] || "").trim();
        if (val && !/^https?:\/\/.+/i.test(val)) {
          this.validationError = `The ${f} link must start with http:// or https://`;
          return;
        }
      }

      const payload = {
        firstName: this.editable.firstName.trim() || this.profile.firstName,
        lastName: this.editable.lastName.trim() || this.profile.lastName,
        description: this.editable.description.trim(),
        socials: {
          facebook: this.editable.socials.facebook.trim(),
          twitter: this.editable.socials.twitter.trim(),
          instagram: this.editable.socials.instagram.trim(),
          github: this.editable.socials.github.trim(),
        },
        // In production: upload files and replace with CDN URLs.
        // Here we commit previews (object URLs) if set; else keep existing.
        profileImageFile: this.editable.newAvatarFile,
        coverImageFile: this.editable.newCoverFile,
      };

      this.$emit("save", payload);
    },
  },
};
</script>
